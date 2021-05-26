// Copyright Simcord LLC
// This file is part of the Bitbon System Blockchain Node library.
//
// The Bitbon System Blockchain Node library is free software: you can redistribute it and/or modify
// it under the terms of the GNU Lesser General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// The Bitbon System Blockchain Node library is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU Lesser General Public License for more details.
//
// You should have received a copy of the GNU Lesser General Public License
// along with the Bitbon System Blockchain Node library. If not, see <http://www.gnu.org/licenses/>.

// Package quorum implements the proof-of-concept community proof-of-stake consensus engine.
package quorum

import (
	"bytes"
	"fmt"
	"math/big"
	"sort"
	"sync"
	"time"

	lru "github.com/hashicorp/golang-lru"
	"github.com/pkg/errors"
	"github.com/simcord-llc/bitbon-system-blockchain/common"
	"github.com/simcord-llc/bitbon-system-blockchain/consensus"
	"github.com/simcord-llc/bitbon-system-blockchain/consensus/misc"
	"github.com/simcord-llc/bitbon-system-blockchain/core/rawdb"
	"github.com/simcord-llc/bitbon-system-blockchain/core/state"
	"github.com/simcord-llc/bitbon-system-blockchain/core/types"
	"github.com/simcord-llc/bitbon-system-blockchain/crypto"
	"github.com/simcord-llc/bitbon-system-blockchain/ethdb"
	"github.com/simcord-llc/bitbon-system-blockchain/log"
	"github.com/simcord-llc/bitbon-system-blockchain/node"
	"github.com/simcord-llc/bitbon-system-blockchain/params"
	"github.com/simcord-llc/bitbon-system-blockchain/rlp"
	"github.com/simcord-llc/bitbon-system-blockchain/rpc"
	"golang.org/x/crypto/sha3"
)

const (
	inMemorySnapshots  = 128  // Number of recent vote snapshots to keep in memory
	inMemorySignatures = 4096 // Number of recent block signatures to keep in memory
)

var (
	ethDifficulty = big.NewInt(2000)
)

// QuorumEngine proof-of-concept community proof-of-stake protocol constants.
var (
	extraSeal = 65 // Fixed number of extra-data suffix bytes reserved for signer seal

	uncleHash = types.CalcUncleHash(nil) // Always Keccak256(RLP([])) as uncles are meaningless outside of PoW.

	FrontierBlockReward       = big.NewInt(5e+18) // Block reward in wei for successfully mining a block
	ByzantiumBlockReward      = big.NewInt(3e+18) // Block reward in wei for successfully mining a block
	ConstantinopleBlockReward = big.NewInt(2e+18) // Block reward in wei for successfully mining a block
)

// Some weird constants to avoid constant memory allocs for them.
var (
	big8  = big.NewInt(8)
	big32 = big.NewInt(32)
)

// Various error messages to mark blocks invalid. These should be private to
// prevent engine specific errors from being referenced in the remainder of the
// codebase, inherently breaking if the engine is swapped out. Please put common
// error types into the consensus package.
var (
	// errUnknownBlock is returned when the list of signers is requested for a block
	// that is not part of the local blockchain.
	errUnknownBlock = errors.New("unknown block")

	// errMissingVanity is returned if a block's extra-data section is shorter than
	// 32 bytes, which is required to store the signer vanity.
	errMissingVanity = errors.New("extra-data 32 byte vanity prefix missing")

	// errMissingSignature is returned if a block's extra-data section doesn't seem
	// to contain a 65 byte secp256k1 signature.
	errMissingSignature = errors.New("extra-data 65 byte signature suffix missing")

	// errExtraBlockProducers is returned if non-checkpoint block contain block producer data in
	// their extra-data fields.
	errExtraBlockProducers = errors.New("non-checkpoint block contains extra block producer list")

	// errInvalidCheckpointSigners is returned if a checkpoint block contains an
	// invalid list of signers (i.e. non divisible by 20 bytes).
	errInvalidCheckpointSigners = errors.New("invalid block producer list on checkpoint block")

	// errInvalidMixDigest is returned if a block's mix digest is non-zero.
	errInvalidMixDigest = errors.New("non-zero mix digest")

	// errInvalidUncleHash is returned if a block contains an non-empty uncle list.
	errInvalidUncleHash = errors.New("non empty uncle hash")

	// errUnauthorizedSigner is returned if a header is signed by a non-authorized entity.
	errUnauthorizedSigner = errors.New("unauthorized producer")

	// errRecentlySigned is returned if a header is signed by an authorized entity
	// that already signed a header recently, thus is temporarily not allowed to.
	errRecentlySigned = errors.New("recently produced")

	// errOutOfTurn is returned if a header is signed by an authorized entity
	// that signed the header not in it's turn.
	errOutOfTurn = errors.New("out of turn signing")

	// errSeveralSigningsInPeriod is returned if a header is signed by an authorized entity
	// several times in one period
	errSeveralSigningsInPeriod = errors.New("several signings in period")

	// errCheckpointSignersMismatch is returned if a checkpoint block contains signers,
	// which does not match QuorumService.MinerListAtEpoch
	errCheckpointSignersMismatch = errors.New("checkpoint signers mismatch")

	errQuorumServiceNotPrepared = errors.New("quorum service is not prepared yet")
)

// SignerFn is a signer callback function to request a hash to be signed by a
// backing account.
type SignerFn func([]byte) ([]byte, error)

// sigHash returns the hash which is used as input for the proof-of-concept community proof-of-stake
// signing. It is the hash of the entire header apart from the 65 byte signature
// contained at the end of the extra data.
//
// Note, the method requires the extra data to be at least 65 bytes, otherwise it
// panics. This is done to avoid accidentally using both forms (signature present
// or not), which could be abused to produce different hashes for the same header.
func sigHash(header *types.Header) (hash common.Hash, err error) {
	hasher := sha3.NewLegacyKeccak256()

	err = rlp.Encode(hasher, []interface{}{
		header.ParentHash,
		header.UncleHash,
		header.Coinbase,
		header.Root,
		header.TxHash,
		header.ReceiptHash,
		header.Bloom,
		header.Difficulty,
		header.Number,
		header.GasLimit,
		header.GasUsed,
		header.Time,
		header.Extra[:len(header.Extra)-65], // Yes, this will panic if extra is too short
		header.MixDigest,
		header.Nonce,
	})
	hasher.Sum(hash[:0])
	return hash, err
}

// ecrecover extracts the Ethereum account address from a signed header.
func ecrecover(header *types.Header, sigcache *lru.ARCCache) (common.Address, error) {
	// If the signature's already cached, return that
	hash := header.Hash()
	if address, known := sigcache.Get(hash); known {
		return address.(common.Address), nil
	}
	// Retrieve the signature from the header extra-data
	if len(header.Extra) < extraSeal {
		return common.Address{}, errMissingSignature
	}
	signature := header.Extra[len(header.Extra)-extraSeal:]

	// Recover the public key and the Ethereum address
	key, err := sigHash(header)
	if err != nil {
		return common.Address{}, err
	}
	pubkey, err := crypto.Ecrecover(key.Bytes(), signature)
	if err != nil {
		return common.Address{}, err
	}
	var signer common.Address
	copy(signer[:], crypto.Keccak256(pubkey[1:])[12:])

	sigcache.Add(hash, signer)
	return signer, nil
}

type QuorumService interface { // nolint
	node.Service
	Prepared() bool
	PrepareNodeListAtEpoch(uint64)
	MinerListAtEpoch(uint64) []common.Address
}

// QuorumEngine is the proof-of-concept community proof-of-stake consensus engine
type QuorumEngine struct { // nolint
	config *params.QuorumEngineConfig // Consensus engine configuration parameters
	quorum QuorumService              // Quorum Service for interactions with quorum functionality
	db     ethdb.Database             // Database to store and retrieve snapshot checkpoints

	recents    *lru.ARCCache // Snapshots for recent block to speed up reorgs
	signatures *lru.ARCCache // Signatures of recent blocks to speed up mining

	signer            common.Address // Ethereum address of the signing key
	signFn            SignerFn       // Signer function to authorize hashes with
	lastBlockProducer common.Address // Last producer of the previous epoch to exclude from producing in the current epoch
	lock              sync.RWMutex   // Protects the signer fields

	// nilUncleHash cash for types.CalcUncleHash(nil) function. Purpose do not calculate it in every Finalize method
	nilUncleHash common.Hash

	cposStartBlock *big.Int // number of block when start work cpos

	recentEpochBlocks *lru.ARCCache
}

// New creates a QuorumEngine proof-of-concept community proof-of-stake consensus engine with the initial
// signers set to the ones provided in the genesis.json
func New(config *params.QuorumEngineConfig, db ethdb.Database, q QuorumService, cposBlock *big.Int) *QuorumEngine {
	// Set any missing consensus parameters to their defaults
	conf := *config
	// Allocate the snapshot caches and create the engine
	recents, _ := lru.NewARC(inMemorySnapshots)
	signatures, _ := lru.NewARC(inMemorySignatures)
	recentEpochBlocks, _ := lru.NewARC(inMemorySignatures)

	// reset cposBlock if it is 0
	if cposBlock != nil && cposBlock.Cmp(big.NewInt(0)) == 0 {
		cposBlock = nil
	}

	return &QuorumEngine{
		config:            &conf,
		db:                db,
		recents:           recents,
		signatures:        signatures,
		quorum:            q,
		nilUncleHash:      types.CalcUncleHash(nil),
		cposStartBlock:    cposBlock,
		recentEpochBlocks: recentEpochBlocks,
	}
}

func (q *QuorumEngine) GetSigner() common.Address {
	return q.signer
}

// Author implements consensus.Engine, returning the Ethereum address recovered
// from the signature in the header's extra-data section.
func (q *QuorumEngine) Author(header *types.Header) (common.Address, error) {
	if q.isPOWBlock(header) {
		return header.Coinbase, nil
	}
	return ecrecover(header, q.signatures)
}

// VerifyHeader checks whether a header conforms to the consensus rules.
func (q *QuorumEngine) VerifyHeader(chain consensus.ChainReader, header *types.Header, seal bool) error {
	if q.isPOWBlock(header) {
		return nil
	}
	err := q.verifyHeader(chain, header, nil)
	return err
}

// VerifyHeaders is similar to VerifyHeader, but verifies a batch of headers. The
// method returns a quit channel to abort the operations and a results channel to
// retrieve the async verifications (the order is that of the input slice).
// nolint:gocritic
func (q *QuorumEngine) VerifyHeaders(chain consensus.ChainReader,
	headers []*types.Header, seals []bool) (chan<- struct{}, <-chan error) {
	var cposHeaders []*types.Header

	if len(headers) > 0 {
		if !q.isPOWBlock(headers[0]) {
			// all blocks are cpos
			cposHeaders = headers
		} else {
			// finding last pow header and splitting headers by engine type: eth (pow) or quorum (cpos)
			for i := len(headers) - 1; i >= 0; i-- {
				if q.isPOWBlock(headers[i]) {
					cposHeaders = headers[i-1:]
				}
			}
		}
	}

	abort := make(chan struct{})
	results := make(chan error, len(headers))

	go func() {
		if len(cposHeaders) > 0 {
			for i, header := range cposHeaders {
				err := q.verifyHeader(chain, header, cposHeaders[:i])

				select {
				case <-abort:
					return
				case results <- err:
				}
			}
		}
	}()

	return abort, results
}

// VerifyUncles implements consensus.Engine, always returning an error for any
// uncles as this consensus mechanism doesn't permit uncles.
func (q *QuorumEngine) VerifyUncles(chain consensus.ChainReader, block *types.Block) error {
	if q.isPOWBlock(block.Header()) {
		return nil
	}
	if len(block.Uncles()) > 0 {
		return errors.New("uncles not allowed")
	}
	return nil
}

// VerifySeal implements consensus.Engine, checking whether the signature contained
// in the header satisfies the consensus protocol requirements.
func (q *QuorumEngine) VerifySeal(chain consensus.ChainReader, header *types.Header) error {
	if q.isPOWBlock(header) {
		return nil
	}
	return q.verifySeal(chain, header, nil)
}

// Prepare implements consensus.Engine, preparing all the consensus fields of the
// header for running the transactions on top.
func (q *QuorumEngine) Prepare(chain consensus.ChainReader, header *types.Header) (err error) {
	if q.isPOWBlock(header) {
		return nil
	}
	if !q.quorum.Prepared() {
		return errQuorumServiceNotPrepared
	}

	// If the block isn't a checkpoint, cast a random vote (good enough for now)
	header.Nonce = types.BlockNonce{}
	number := header.Number.Uint64()
	q.lock.RLock()
	header.Coinbase = q.signer
	q.lock.RUnlock()
	// Ensure the extra data has all it's components
	if len(header.Extra) < rawdb.ExtraSnapshotHash {
		header.Extra = append(header.Extra, bytes.Repeat([]byte{0x00}, rawdb.ExtraSnapshotHash-len(header.Extra))...)
	}
	header.Extra = header.Extra[:rawdb.ExtraSnapshotHash]
	now := uint64(time.Now().Unix())
	header.Time = now - now%q.config.Period
	checkpoint, err := q.isCheckpoint(chain, header, nil)
	if err != nil {
		return err
	}

	epoch := q.getEpoch(header)
	if number == 1 || checkpoint || (q.cposStartBlock != nil && q.cposStartBlock.Cmp(header.Number) == 0) {
		// prepare miners list for the next epoch
		q.quorum.PrepareNodeListAtEpoch(epoch + 1)
	}

	// header for getting snapshot
	// if current block is a checkpoint block - use it for getting snapshot
	// otherwise use parent block
	var headerForSnapshot *types.Header
	if checkpoint {
		minerAddresses := q.quorum.MinerListAtEpoch(epoch)
		if len(minerAddresses) == 0 {
			return errors.Errorf("Miners for epoch %d are not yet gathered", epoch)
		}
		for _, sig := range minerAddresses {
			header.Extra = append(header.Extra, sig[:]...)
		}
		// if header is a snapshot block create new snapshot based on the header
		headerForSnapshot = header
	} else {
		// if header is not a checkpoint block, get parent's snapshot
		parent := chain.GetHeader(header.ParentHash, number-1)
		if parent == nil {
			return errors.Errorf("unable to get parent block %s for non checkpoint block", header.ParentHash.Hex())
		}
		headerForSnapshot = parent
	}
	// add empty data for extraSeal
	header.Extra = append(header.Extra, make([]byte, extraSeal)...)

	// snapshot method validates header ExtraData length -
	// so it is required to allocate memory for extraSeal
	// before calling snapshot method
	snap, err := q.snapshot(headerForSnapshot)
	if err != nil {
		return errors.Wrap(err, "unable to get snapshot")
	}

	// add snapshot hash to Header.Extra
	copy(header.Extra[:rawdb.ExtraSnapshotHash], snap.Hash().Bytes())

	// Set the correct difficulty
	header.Difficulty = CalcDifficulty(header, snap)

	// Mix digest is reserved for now, set to empty
	header.MixDigest = common.Hash{}

	return nil
}

// Finalize implements consensus.Engine, ensuring no uncles are set, nor block
// rewards given, and returns the final block.
func (q *QuorumEngine) Finalize(chain consensus.ChainReader,
	header *types.Header, stateDB *state.DB, _ []*types.Transaction, _ []*types.Header) {
	if q.isPOWBlock(header) {
		return
	}
	header.Root = stateDB.IntermediateRoot(chain.Config().IsEIP158(header.Number))
	header.UncleHash = q.nilUncleHash
}

func (q *QuorumEngine) FinalizeAndAssemble(chain consensus.ChainReader, header *types.Header,
	stateDB *state.DB, txs []*types.Transaction,
	uncles []*types.Header, receipts []*types.Receipt) (*types.Block, error) {
	if q.isPOWBlock(header) {
		// Accumulate any block and uncle rewards and commit the final state root
		accumulateRewards(chain.Config(), stateDB, header, uncles)
		header.Root = stateDB.IntermediateRoot(chain.Config().IsEIP158(header.Number))
		// Header seems complete, assemble into a block and return
		return types.NewBlock(header, txs, uncles, receipts), nil
	}

	// Accumulate any block and uncle rewards and commit the final state root
	header.Root = stateDB.IntermediateRoot(chain.Config().IsEIP158(header.Number))
	// Header seems complete, assemble into a block and return
	return types.NewBlock(header, txs, nil, receipts), nil
}

// Authorize injects a private key into the consensus engine to mint new blocks
// with.
func (q *QuorumEngine) Authorize(signer common.Address, signFn SignerFn) {
	q.lock.Lock()
	defer q.lock.Unlock()

	q.signer = signer
	q.signFn = signFn
}

func (q *QuorumEngine) blockTimestampHash(blockNumber, blockTimestamp *big.Int) []byte {
	return crypto.Keccak256(blockNumber.Bytes(), blockTimestamp.Bytes())
}

// Seal implements consensus.Engine, attempting to create a sealed block using
// the local signing credentials.
// nolint:funlen
func (q *QuorumEngine) Seal(chain consensus.ChainReader, block *types.Block, results chan<- *types.Block,
	stop <-chan struct{}) (err error) {
	if q.isPOWBlock(block.Header()) {
		return nil
	}

	header := block.Header()
	// Sealing the genesis block is not supported
	number := header.Number.Uint64()
	if number == 0 {
		return errUnknownBlock
	}

	parent := chain.GetHeader(header.ParentHash, number-1)
	if parent == nil {
		return consensus.ErrUnknownAncestor
	}

	// Don't hold the signer fields for the entire sealing procedure
	q.lock.RLock()
	signer, signFn := q.signer, q.signFn
	q.lock.RUnlock()

	// Bail out if we're unauthorized to sign a block
	snap, err := q.snapshot(header)
	if err != nil {
		return err
	}

	checkpoint, err := q.isCheckpoint(chain, header, nil)
	if err != nil {
		return err
	}

	if _, authorized := snap.Signers[signer]; !authorized {
		return errUnauthorizedSigner
	}

	if !snap.inturn(header, signer) {
		return errOutOfTurn
	}

	// check to prevent several blocks by one signer in one period
	if snap.signersElapsed(header) == snap.signersElapsed(parent) {
		return errSeveralSigningsInPeriod
	}

	// skip block by default ethereum timer if there are no any txs and it's not a checkpoint block
	if len(block.Transactions()) == 0 && !checkpoint {
		return nil
	}

	if len(snap.Signers) > 1 && parent.Coinbase == header.Coinbase && !checkpoint {
		return errRecentlySigned
	}

	// Sign all the things!
	key, err := sigHash(header)
	if err != nil {
		return err
	}
	sighash, err := signFn(key.Bytes())
	if err != nil {
		return err
	}
	copy(header.Extra[len(header.Extra)-extraSeal:], sighash)

	// Wait until sealing is terminated or delay timeout.
	go func() {
		select {
		case <-stop:
			return
		default:
		}

		select {
		case results <- block.WithSeal(header):
		default:
		}
	}()

	return nil
}

// CalcDifficulty is the difficulty adjustment algorithm. It returns the difficulty
// that a new block should have based on the previous blocks in the chain and the
// current signer.
func (q *QuorumEngine) CalcDifficulty(chain consensus.ChainReader, _ uint64, parent *types.Header) (res *big.Int) {
	if q.isPOWBlock(parent) {
		return ethDifficulty
	}
	snap, err := q.snapshot(parent)
	if err != nil {
		return
	}
	return CalcDifficulty(parent, snap)
}

// CalcDifficulty is the difficulty adjustment algorithm. It returns the difficulty
// that a new block should have based on the previous blocks in the chain and the
// current signer.
func CalcDifficulty(header *types.Header, snap *Snapshot) *big.Int {
	signersPerEpoch := snap.config.Epoch / snap.config.Period
	return new(big.Int).SetUint64(signersPerEpoch - snap.signersElapsed(header))
}

// SealHash returns the hash of a block prior to it being sealed.
func (q *QuorumEngine) SealHash(header *types.Header) common.Hash {
	if q.isPOWBlock(header) {
		return sealEthHash(header)
	}
	addr, err := sigHash(header)
	if err != nil {
		return common.Hash{}
	}
	return addr
}

// Close implements consensus.Engine. It's a noop for quorum as there are no background threads.
func (q *QuorumEngine) Close() error {
	return nil
}

// APIs implements consensus.Engine, returning the user facing RPC API to allow
// controlling the signer voting.
func (q *QuorumEngine) APIs(chain consensus.ChainReader) []rpc.API {
	return []rpc.API{{
		Namespace: "quorumengine",
		Version:   "1.0",
		Service:   &API{chain: chain, quorum: q},
		Public:    false,
	}}
}

// AuthorTimestamp extends consensus.Engine, returning the Ethereum address recovered
// from the signature in the timestamp message's signature field.
func (q *QuorumEngine) AuthorTimestamp(blockNumber, blockTimestamp *big.Int, signature []byte) (common.Address, error) {
	return q.recoverFromTimestamp(blockNumber, blockTimestamp, signature)
}

// Inturn returns true if a it's a signer's turn to produce a block at a given timestamp.
func (q *QuorumEngine) Inturn(blockNumber, blockTimestamp *big.Int, signature []byte, snapshotHash common.Hash) bool {
	address, err := q.recoverFromTimestamp(blockNumber, blockTimestamp, signature)
	if err != nil {
		log.Error("unable to recover public key from the signature")
		return false
	}

	fakeHeader := &types.Header{Number: blockNumber, Time: blockTimestamp.Uint64(),
		Coinbase: address, Extra: make([]byte, rawdb.ExtraSnapshotHash)}
	copy(fakeHeader.Extra, snapshotHash.Bytes())

	snap, err := q.snapshot(fakeHeader)
	if err != nil {
		log.Error("unable to load or create snapshot", "error", err)
		return false
	}

	return snap.inturn(fakeHeader, address)
}

// GetBlockProducers shuffles the miners list using hash as shuffle seed,
// compacts the list is the form of map[common.Address]BlockProducer
func (q *QuorumEngine) GetBlockProducerSequence(miners []common.Address,
	hash common.Hash) map[common.Address]BlockProducer {
	res := make(map[common.Address]BlockProducer)

	sort.SliceStable(miners, func(i, j int) bool {
		return bytes.Compare(miners[i].Bytes(), miners[j].Bytes()) > 0
	})
	sequence := generateSequence(hash, len(miners))
	lastAddress := q.lastBlockProducer
	for index, address := range miners {
		lastAddress = address
		res[address] = BlockProducer{SequenceNumber: sequence[index]}
	}

	// save last address to global variable
	q.lastBlockProducer = lastAddress
	return res
}

func (q *QuorumEngine) Config() *params.QuorumEngineConfig {
	return q.config
}

// verifyCascadingFields verifies all the header fields that are not standalone,
// rather depend on a batch of previous headers. The caller may optionally pass
// in a batch of parents (ascending order) to avoid looking those up from the
// database. This is useful for concurrently verifying a batch of new headers.
func (q *QuorumEngine) verifyCascadingFields(chain consensus.ChainReader,
	header *types.Header, parents []*types.Header) error {
	// The genesis block is the always valid dead-end
	number := header.Number.Uint64()
	if number == 0 {
		return nil
	}

	return q.verifySeal(chain, header, parents)
}

// verifySeal checks whether the signature contained in the header satisfies the
// consensus protocol requirements. The method accepts an optional list of parent
// headers that aren't yet part of the local blockchain to generate the snapshots
// from.
// nolint:gocyclo
func (q *QuorumEngine) verifySeal(chain consensus.ChainReader,
	header *types.Header, parents []*types.Header) (err error) {
	// Verifying the genesis block is not supported
	number := header.Number.Uint64()
	if number == 0 {
		return errUnknownBlock
	}

	// Ensure that the block's timestamp isn't too close to it's parent
	var parent *types.Header
	if len(parents) > 0 {
		parent = parents[len(parents)-1]
	} else {
		parent = chain.GetHeader(header.ParentHash, number-1)
	}
	if parent == nil || parent.Number.Uint64() != number-1 || parent.Hash() != header.ParentHash {
		return consensus.ErrUnknownAncestor
	}

	checkpoint, err := q.isCheckpoint(chain, header, parents)
	if err != nil {
		return err
	}

	// Retrieve the snapshot needed to verify this header and cache it
	snap, err := q.snapshot(header)
	if err != nil && checkpoint {
		return err
	}

	if snap == nil {
		return nil
	}

	// try to validate signers in checkpoint block
	if checkpoint {
		curEpoch := q.getEpoch(header)
		if curEpoch != 0 {
			if minerList := q.quorum.MinerListAtEpoch(curEpoch); len(minerList) != 0 {
				if !snap.checkSigners(minerList) {
					return errCheckpointSignersMismatch
				}
			}
		}
	}

	if len(snap.Signers) > 1 && parent.Coinbase == header.Coinbase && !checkpoint {
		return errRecentlySigned
	}

	// check to prevent several blocks by one signer in one period
	if snap.signersElapsed(header) == snap.signersElapsed(parent) {
		return errSeveralSigningsInPeriod
	}

	// Resolve the authorization key and check against signers
	signer, err := ecrecover(header, q.signatures)
	if err != nil {
		return err
	}
	if _, ok := snap.Signers[signer]; !ok {
		return errUnauthorizedSigner
	}

	if !snap.inturn(header, signer) {
		return errOutOfTurn
	}

	return nil
}

// verifyHeader checks whether a header conforms to the consensus rules.The
// caller may optionally pass in a batch of parents (ascending order) to avoid
// looking those up from the database. This is useful for concurrently verifying
// a batch of new headers.
func (q *QuorumEngine) verifyHeader(chain consensus.ChainReader,
	header *types.Header, parents []*types.Header) (err error) {
	defer func() {
		epoch := q.getEpoch(header)
		if err != nil {
			log.Error("verifyHeader ended", "block", header.Number, "hash", header.Hash().Hex(),
				"epoch", epoch, "error", err, "extradata", common.Bytes2Hex(header.Extra))
		}
	}()

	if header.Number == nil {
		return errUnknownBlock
	}

	// Don't waste time checking blocks from the future
	if header.Time > uint64(time.Now().Unix()) {
		return consensus.ErrFutureBlock
	}
	// Checkpoint blocks need to enforce zero beneficiary
	checkpoint, err := q.isCheckpoint(chain, header, parents)
	if err != nil {
		return err
	}

	// Check that the extra-data contains both the vanity and signature
	if len(header.Extra) < rawdb.ExtraSnapshotHash {
		return errMissingVanity
	}
	if len(header.Extra) < rawdb.ExtraSnapshotHash+extraSeal {
		return errMissingSignature
	}
	// Ensure that the extra-data contains a signer list on checkpoint, but none otherwise
	signersBytes := len(header.Extra) - rawdb.ExtraSnapshotHash - extraSeal
	if !checkpoint && signersBytes != 0 {
		return errExtraBlockProducers
	}
	if checkpoint && signersBytes%common.AddressLength != 0 {
		return errInvalidCheckpointSigners
	}
	// Ensure that the mix digest is zero as we don't have fork protection currently
	if header.MixDigest != (common.Hash{}) {
		return errInvalidMixDigest
	}
	// Ensure that the block doesn't contain any uncles which are meaningless in PoA
	if header.UncleHash != uncleHash {
		return errInvalidUncleHash
	}
	// If all checks passed, validate any special fields for hard forks
	if err := misc.VerifyForkHashes(chain.Config(), header, false); err != nil {
		return err
	}
	// All basic checks passed, verify cascading fields
	return q.verifyCascadingFields(chain, header, parents)
}

// isCheckpoint gets epochs of header and header's parent and compares them
// if the header epoch is greater than the parent's epoch isCheckpoint returns true, otherwise - false
func (q *QuorumEngine) isCheckpoint(chain consensus.ChainReader,
	header *types.Header, parents []*types.Header) (bool, error) {
	curEpoch := q.getEpoch(header)
	if curEpoch == 0 {
		return false, nil
	}
	number := header.Number.Uint64()
	// Ensure that the block's timestamp isn't too close to it's parent
	var parent *types.Header
	if len(parents) > 0 {
		parent = parents[len(parents)-1]
	} else {
		parent = chain.GetHeader(header.ParentHash, number-1)
	}
	if parent == nil || parent.Number.Uint64() != number-1 || parent.Hash() != header.ParentHash {
		return false, nil
	}

	parentEpoch := q.getEpoch(parent)
	checkpoint := curEpoch > parentEpoch

	return checkpoint, nil
}

func (q *QuorumEngine) GetEpochByTime(t uint64) uint64 {
	if t <= q.config.Start {
		// for excluding overhead uint64 value
		return 0
	}
	epoch := (t - q.config.Start) / q.config.Epoch
	return epoch
}

// getEpoch returns epoch of the header calculated from it's timestamp and QuorumEngine config
func (q *QuorumEngine) getEpoch(header *types.Header) uint64 {
	return q.GetEpochByTime(header.Time)
}

// Checkpoint gets epochs of header and header's parent and compares them
// if the header epoch is greater than the parent's epoch isCheckpoint returns true, otherwise - false
func (q *QuorumEngine) Checkpoint(current, parent *types.Header, chain consensus.ChainReader) (bool, error) {
	return q.isCheckpoint(chain, current, []*types.Header{parent})
}

// Snapshot loads the authorization snapshot at a given point in time.
func (q *QuorumEngine) Snapshot(header *types.Header) (*Snapshot, error) {
	snap := q.tryLoadSnapshot(header)
	if snap != nil {
		return snap, nil
	}
	return nil, errors.New("snapshot not found")
}

// snapshot retrieves the authorization snapshot at a given point in time.
// In this version, snapshotting is based on epochs, although there is no "canonical" snapshot for an epoch.
// Snapshots are stored under a key (snapshot's hash) that is comprised of snapshot's epoch and signers.
// This allows GHOST rule out reorgs as all the possible snapshots for an epoch are stored and every block
// states the snapshot hash that it is meant to be verified against.
// Before this, snapshot used to stay the same and it would cause get in the way of GHOST by returning BAD BLOCK.
func (q *QuorumEngine) snapshot(header *types.Header) (*Snapshot, error) {
	if header == nil {
		return nil, errors.New("empty block header")
	}
	if header.Time == 0 {
		return nil, errors.New("empty block header time")
	}

	// First, look up snapshot in cache and on disk
	snap := q.tryLoadSnapshot(header)
	if snap != nil {
		return snap, nil
	}

	// If the snapshot could not be found, create new snapshot if it's possible (if the block is a snapshot block)
	snap, err := q.snapshotFromHeader(header)
	if err != nil {
		return nil, err
	}

	// Try to persist snapshot to cache and on disk under it's hash
	err = q.tryStoreSnapshot(snap)
	if err != nil {
		return nil, err
	}

	return snap, nil
}

func (q *QuorumEngine) tryStoreSnapshot(snap *Snapshot) error {
	if err := snap.store(q.db); err != nil {
		return err
	}
	q.recents.Add(snap.Hash(), snap)

	return nil
}

func (q *QuorumEngine) tryLoadSnapshot(header *types.Header) *Snapshot {
	if len(header.Extra) < rawdb.ExtraSnapshotHash {
		return nil
	}
	snapHash := common.BytesToHash(header.Extra[:rawdb.ExtraSnapshotHash])
	var snap *Snapshot

	// If an in-memory snapshot was found, use that
	if s, ok := q.recents.Get(snapHash); ok {
		snap = s.(*Snapshot)
		return snap
	}

	if s, err := loadSnapshot(q.config, q.signatures, q.db, snapHash, header.Number.Uint64()); err == nil {
		snap = s
		return snap
	}

	return nil
}

func (q *QuorumEngine) snapshotFromHeader(header *types.Header) (*Snapshot, error) {
	var signers []common.Address
	var sequenceHash common.Hash
	epoch := q.getEpoch(header)

	if epoch == 0 || (q.cposStartBlock != nil && q.cposStartBlock.Cmp(header.Number) == 0) {
		signers = make([]common.Address, len(q.config.InitialSigners))
		copy(signers, q.config.InitialSigners)
		sequenceHash = q.config.InitialHash
	} else {
		if header.Extra == nil {
			return nil, errors.New(fmt.Sprintf("Empty block extra data. Header number %v, header hash %s",
				header.Number, header.Hash().Hex()))
		}

		if len(header.Extra) <= extraSeal+rawdb.ExtraSnapshotHash {
			return nil, errors.New(fmt.Sprintf("Block extra data does not contain signers. "+
				"Header number %v, header hash %s, header.Extra len %d, must be greater then %d",
				header.Number, header.Hash().Hex(), len(header.Extra), extraSeal+rawdb.ExtraSnapshotHash))
		}

		signers = make([]common.Address, (len(header.Extra)-rawdb.ExtraSnapshotHash-extraSeal)/common.AddressLength)
		for i := 0; i < len(signers); i++ {
			copy(signers[i][:], header.Extra[rawdb.ExtraSnapshotHash+i*common.AddressLength:])
		}
		sequenceHash = header.ParentHash
	}

	blockProducers := q.GetBlockProducerSequence(signers, sequenceHash)

	return newSnapshot(q.config, q.signatures, blockProducers, epoch), nil
}

// recoverFromTimestamp returns address recovered from the signature in the timestamp message's signature field.
func (q *QuorumEngine) recoverFromTimestamp(blockNumber, blockTimestamp *big.Int,
	signature []byte) (common.Address, error) {
	hash := q.blockTimestampHash(blockNumber, blockTimestamp)
	pubKey, err := crypto.SigToPub(hash, signature)
	if err != nil {
		return common.Address{}, err
	}
	return crypto.PubkeyToAddress(*pubKey), nil
}

func (q *QuorumEngine) isPOWBlock(header *types.Header) bool {
	if q.cposStartBlock == nil {
		return false
	}
	return q.cposStartBlock.Cmp(header.Number) == 1
}

// generateSequence returns slice of uint64 of the passed size
// uint64s in the slice represent turns of the block producers with the same index as uint64s
func generateSequence(hash common.Hash, size int) []uint64 {
	slice := generateIndexes(size)
	return shuffle2(hash, slice)
}

func shuffle2(hash common.Hash, slice []uint64) []uint64 {
	l := len(slice)
	if l == 0 {
		return nil
	}

	for i := 0; i < len(hash)/2; i++ {
		i1 := int(hash[i]) % l
		i2 := int(hash[len(hash)-i-1]) % l
		slice[i1], slice[i2] = slice[i2], slice[i1]
	}
	return slice
}

func shuffle3(hash common.Hash, slice []uint64) []uint64 {
	l := uint64(len(slice))
	if l == 0 {
		return nil
	}

	seed1 := new(big.Int).SetBytes(hash[:len(hash)/2]).Uint64()
	seed2 := new(big.Int).SetBytes(hash[len(hash)/2:]).Uint64()

	for i := uint64(0); i < l; i++ {
		i1 := (seed1 + i) % l
		i2 := (seed2 + i) % l
		slice[i1], slice[i2] = slice[i2], slice[i1]
	}
	return slice
}

func generateIndexes(n int) []uint64 {
	res := make([]uint64, n)
	for i := 0; i < n; i++ {
		res[i] = uint64(i)
	}
	return res
}

func accumulateRewards(config *params.ChainConfig, stateDB *state.DB,
	header *types.Header, uncles []*types.Header) {
	// Select the correct block reward based on chain progression
	blockReward := FrontierBlockReward
	if config.IsByzantium(header.Number) {
		blockReward = ByzantiumBlockReward
	}
	if config.IsConstantinople(header.Number) {
		blockReward = ConstantinopleBlockReward
	}
	// Accumulate the rewards for the miner and any included uncles
	reward := new(big.Int).Set(blockReward)
	r := new(big.Int)
	for _, uncle := range uncles {
		r.Add(uncle.Number, big8)
		r.Sub(r, header.Number)
		r.Mul(r, blockReward)
		r.Div(r, big8)
		stateDB.AddBalance(uncle.Coinbase, r)

		r.Div(blockReward, big32)
		reward.Add(reward, r)
	}
	stateDB.AddBalance(header.Coinbase, reward)
}

func sealEthHash(header *types.Header) (hash common.Hash) {
	hasher := sha3.NewLegacyKeccak256()

	rlp.Encode(hasher, []interface{}{ // nolint
		header.ParentHash,
		header.UncleHash,
		header.Coinbase,
		header.Root,
		header.TxHash,
		header.ReceiptHash,
		header.Bloom,
		header.Difficulty,
		header.Number,
		header.GasLimit,
		header.GasUsed,
		header.Time,
		header.Extra,
	})
	hasher.Sum(hash[:0])
	return hash
}
