// Copyright Simcord LLC
// This file is part of the Bitbon System Blockchain Node library.
// This file has been modified, you can find the original version by following the link
// <https://github.com/ethereum/go-ethereum>
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

// nolint:gocritic
package zksnark

// #cgo !darwin LDFLAGS: -L${SRCDIR} -lzksnark -lstdc++ -lgmp -lgomp
// #cgo darwin LDFLAGS: -L${SRCDIR} -lzksnark -lstdc++ -lgmp -lomp
// #include <zksnark.h>
import "C"
import (
	"sync"
	"unsafe"

	"github.com/pkg/errors"
)

//go:generate make all

// Init() is only ever called once
var onceInit sync.Once

const pkLen = 344733
const vkLen = 2882
const proofLen = 259
const hashLen = 32

// nolint:gochecknoinits
func init() {
	Init()
}

func Init() {
	onceInit.Do(func() {
		C.initialize(C.bool(false))
	})
}

// GenerateKeys returns newly generated PK and VK.
func GenerateKeys(constraintsLen uint64) (pk []byte, vk []byte) {
	pk = make([]byte, pkLen)
	vk = make([]byte, vkLen)
	C.keypair_generator_v2(unsafe.Pointer(&pk[0]), unsafe.Pointer(&vk[0]), C.uint64_t(hashLen), C.uint64_t(constraintsLen))
	return pk, vk
}

// Prove generates proof depended on hashInput and PK.
func Prove(hashInput []byte, pk []byte, constraintsLen uint64) (proof []byte, err error) { // nolint:gocritic
	if len(hashInput) != hashLen {
		return nil, errors.Errorf("inappropriate hashInput length. Have %d, want %d", len(hashInput), hashLen)
	}
	if len(pk) != pkLen {
		return nil, errors.Errorf("inappropriate PK length. Have %d, want %d", len(pk), pkLen)
	}
	proof = make([]byte, proofLen)
	hashInputPtr := C.CBytes(hashInput[:])
	proveKeywordPtr := C.CBytes(pk[:])

	C.prove_v2(unsafe.Pointer(&proof[0]), proveKeywordPtr, hashInputPtr, C.uint64_t(hashLen), C.uint64_t(constraintsLen))

	C.free(hashInputPtr)
	C.free(proveKeywordPtr)
	return proof, nil
}

// Verify verifies hashInput depended on proof and VK.
func Verify(proof []byte, vk []byte, hashInput []byte) (bool, error) { // nolint:gocritic
	if len(proof) != proofLen {
		return false, errors.Errorf("inappropriate proof length. Have %d, want %d", len(proof), proofLen)
	}
	if len(vk) != vkLen {
		return false, errors.Errorf("inappropriate VK length. Have %d, want %d", len(vk), vkLen)
	}
	if len(hashInput) != hashLen {
		return false, errors.Errorf("inappropriate hashInput length. Have %d, want %d", len(hashInput), hashLen)
	}
	proofPtr := C.CBytes(proof[:])
	vkPtr := C.CBytes(vk[:])
	inputValuePtr := C.CBytes(hashInput[:])

	ret := C.verify_v2(proofPtr, vkPtr, inputValuePtr, C.uint64_t(hashLen))

	C.free(proofPtr)
	C.free(vkPtr)
	C.free(inputValuePtr)

	if ret {
		return true, nil
	}
	return false, nil
}

// Old version
func ProveTransfer(hashInput [64]byte, auxiliaryPerConstraint uint64) ([259]byte, [5058]byte) { // nolint:gocritic
	var proof [259]byte
	var vk [5058]byte

	inputValuePtr := C.CBytes(hashInput[:])

	C.prove(unsafe.Pointer(&proof[0]), unsafe.Pointer(&vk[0]), inputValuePtr, C.uint64_t(auxiliaryPerConstraint))

	C.free(inputValuePtr)

	return proof, vk
}

// Old version
func VerifyTransfer(proof [259]byte, vk [5058]byte, hashInput [64]byte) bool {
	proofPtr := C.CBytes(proof[:])
	vkPtr := C.CBytes(vk[:])
	inputValuePtr := C.CBytes(hashInput[:])

	ret := C.verify(proofPtr, vkPtr, inputValuePtr)

	C.free(proofPtr)
	C.free(vkPtr)
	C.free(inputValuePtr)

	if ret { // nolint
		return true
	}
	return false
}

// Old version
func Circuit(inputValue [64]byte) [32]byte {
	var hash [32]byte
	inputValuePtr := C.CBytes(inputValue[:])

	C.circuit(unsafe.Pointer(&hash[0]), inputValuePtr)
	C.free(inputValuePtr)

	return hash
}

// Old version
func VerifyTransferByHash(proof [259]byte, vk [5058]byte, inputHash [32]byte) bool {
	proofPtr := C.CBytes(proof[:])
	vkPtr := C.CBytes(vk[:])
	inputhashPtr := C.CBytes(inputHash[:])

	ret := C.verify_by_hash(proofPtr, vkPtr, inputhashPtr)

	C.free(proofPtr)
	C.free(vkPtr)
	C.free(inputhashPtr)

	if ret { // nolint
		return true
	}
	return false
}
