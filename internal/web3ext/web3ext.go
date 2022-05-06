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

// package web3ext contains geth specific web3.js extensions.
package web3ext

var Modules = map[string]string{
	"accounting":   AccountingJs,
	"admin":        AdminJs,
	"bitbon":       BitbonJs,
	"bitbonclient": BitbonClientJs,
	"bitbonagent":  BitbonAgentJs,
	"quorum":       QuorumJs,
	"quorumengine": QuorumEngineJs,
	"chequebook":   ChequebookJs,
	"clique":       CliqueJs,
	"ethash":       EthashJs,
	"debug":        DebugJs,
	"eth":          EthJs,
	"miner":        MinerJs,
	"net":          NetJs,
	"personal":     PersonalJs,
	"rpc":          RpcJs,
	"shh":          ShhJs,
	"swarmfs":      SwarmfsJs,
	"txpool":       TxpoolJs,
	"les":          LESJs,
	"lespay":       LESPayJs,
}

const ChequebookJs = `
web3._extend({
	property: 'chequebook',
	methods: [
		new web3._extend.Method({
			name: 'deposit',
			call: 'chequebook_deposit',
			params: 1,
			inputFormatter: [null]
		}),
		new web3._extend.Property({
			name: 'balance',
			getter: 'chequebook_balance',
			outputFormatter: web3._extend.utils.toDecimal
		}),
		new web3._extend.Method({
			name: 'cash',
			call: 'chequebook_cash',
			params: 1,
			inputFormatter: [null]
		}),
		new web3._extend.Method({
			name: 'issue',
			call: 'chequebook_issue',
			params: 2,
			inputFormatter: [null, null]
		}),
	]
});
`

const CliqueJs = `
web3._extend({
	property: 'clique',
	methods: [
		new web3._extend.Method({
			name: 'getSnapshot',
			call: 'clique_getSnapshot',
			params: 1,
			inputFormatter: [web3._extend.formatters.inputBlockNumberFormatter]
		}),
		new web3._extend.Method({
			name: 'getSnapshotAtHash',
			call: 'clique_getSnapshotAtHash',
			params: 1
		}),
		new web3._extend.Method({
			name: 'getSigners',
			call: 'clique_getSigners',
			params: 1,
			inputFormatter: [web3._extend.formatters.inputBlockNumberFormatter]
		}),
		new web3._extend.Method({
			name: 'getSignersAtHash',
			call: 'clique_getSignersAtHash',
			params: 1
		}),
		new web3._extend.Method({
			name: 'propose',
			call: 'clique_propose',
			params: 2
		}),
		new web3._extend.Method({
			name: 'discard',
			call: 'clique_discard',
			params: 1
		}),
		new web3._extend.Method({
			name: 'status',
			call: 'clique_status',
			params: 0
		}),
	],
	properties: [
		new web3._extend.Property({
			name: 'proposals',
			getter: 'clique_proposals'
		}),
	]
});
`

const EthashJs = `
web3._extend({
	property: 'ethash',
	methods: [
		new web3._extend.Method({
			name: 'getWork',
			call: 'ethash_getWork',
			params: 0
		}),
		new web3._extend.Method({
			name: 'getHashrate',
			call: 'ethash_getHashrate',
			params: 0
		}),
		new web3._extend.Method({
			name: 'submitWork',
			call: 'ethash_submitWork',
			params: 3,
		}),
		new web3._extend.Method({
			name: 'submitHashRate',
			call: 'ethash_submitHashRate',
			params: 2,
		}),
	]
});
`

const AdminJs = `
web3._extend({
	property: 'admin',
	methods: [
		new web3._extend.Method({
			name: 'addPeer',
			call: 'admin_addPeer',
			params: 1
		}),
		new web3._extend.Method({
			name: 'removePeer',
			call: 'admin_removePeer',
			params: 1
		}),
		new web3._extend.Method({
			name: 'addTrustedPeer',
			call: 'admin_addTrustedPeer',
			params: 1
		}),
		new web3._extend.Method({
			name: 'removeTrustedPeer',
			call: 'admin_removeTrustedPeer',
			params: 1
		}),
		new web3._extend.Method({
			name: 'exportChain',
			call: 'admin_exportChain',
			params: 3,
			inputFormatter: [null, null, null]
		}),
		new web3._extend.Method({
			name: 'importChain',
			call: 'admin_importChain',
			params: 1
		}),
		new web3._extend.Method({
			name: 'sleepBlocks',
			call: 'admin_sleepBlocks',
			params: 2
		}),
		new web3._extend.Method({
			name: 'startRPC',
			call: 'admin_startRPC',
			params: 4,
			inputFormatter: [null, null, null, null]
		}),
		new web3._extend.Method({
			name: 'stopRPC',
			call: 'admin_stopRPC'
		}),
		new web3._extend.Method({
			name: 'startWS',
			call: 'admin_startWS',
			params: 4,
			inputFormatter: [null, null, null, null]
		}),
		new web3._extend.Method({
			name: 'stopWS',
			call: 'admin_stopWS'
		}),
	],
	properties: [
		new web3._extend.Property({
			name: 'nodeInfo',
			getter: 'admin_nodeInfo'
		}),
		new web3._extend.Property({
			name: 'peers',
			getter: 'admin_peers'
		}),
		new web3._extend.Property({
			name: 'datadir',
			getter: 'admin_datadir'
		}),
	]
});
`

const DebugJs = `
web3._extend({
	property: 'debug',
	methods: [
		new web3._extend.Method({
			name: 'accountRange',
			call: 'debug_accountRange',
			params: 2
		}),
		new web3._extend.Method({
			name: 'printBlock',
			call: 'debug_printBlock',
			params: 1
		}),
		new web3._extend.Method({
			name: 'getBlockRlp',
			call: 'debug_getBlockRlp',
			params: 1
		}),
		new web3._extend.Method({
			name: 'testSignCliqueBlock',
			call: 'debug_testSignCliqueBlock',
			params: 2,
			inputFormatters: [web3._extend.formatters.inputAddressFormatter, null],
		}),
		new web3._extend.Method({
			name: 'setHead',
			call: 'debug_setHead',
			params: 1
		}),
		new web3._extend.Method({
			name: 'seedHash',
			call: 'debug_seedHash',
			params: 1
		}),
		new web3._extend.Method({
			name: 'dumpBlock',
			call: 'debug_dumpBlock',
			params: 1
		}),
		new web3._extend.Method({
			name: 'chaindbProperty',
			call: 'debug_chaindbProperty',
			params: 1,
			outputFormatter: console.log
		}),
		new web3._extend.Method({
			name: 'chaindbCompact',
			call: 'debug_chaindbCompact',
		}),
		new web3._extend.Method({
			name: 'verbosity',
			call: 'debug_verbosity',
			params: 1
		}),
		new web3._extend.Method({
			name: 'vmodule',
			call: 'debug_vmodule',
			params: 1
		}),
		new web3._extend.Method({
			name: 'backtraceAt',
			call: 'debug_backtraceAt',
			params: 1,
		}),
		new web3._extend.Method({
			name: 'stacks',
			call: 'debug_stacks',
			params: 0,
			outputFormatter: console.log
		}),
		new web3._extend.Method({
			name: 'freeOSMemory',
			call: 'debug_freeOSMemory',
			params: 0,
		}),
		new web3._extend.Method({
			name: 'setGCPercent',
			call: 'debug_setGCPercent',
			params: 1,
		}),
		new web3._extend.Method({
			name: 'memStats',
			call: 'debug_memStats',
			params: 0,
		}),
		new web3._extend.Method({
			name: 'gcStats',
			call: 'debug_gcStats',
			params: 0,
		}),
		new web3._extend.Method({
			name: 'cpuProfile',
			call: 'debug_cpuProfile',
			params: 2
		}),
		new web3._extend.Method({
			name: 'startCPUProfile',
			call: 'debug_startCPUProfile',
			params: 1
		}),
		new web3._extend.Method({
			name: 'stopCPUProfile',
			call: 'debug_stopCPUProfile',
			params: 0
		}),
		new web3._extend.Method({
			name: 'goTrace',
			call: 'debug_goTrace',
			params: 2
		}),
		new web3._extend.Method({
			name: 'startGoTrace',
			call: 'debug_startGoTrace',
			params: 1
		}),
		new web3._extend.Method({
			name: 'stopGoTrace',
			call: 'debug_stopGoTrace',
			params: 0
		}),
		new web3._extend.Method({
			name: 'blockProfile',
			call: 'debug_blockProfile',
			params: 2
		}),
		new web3._extend.Method({
			name: 'setBlockProfileRate',
			call: 'debug_setBlockProfileRate',
			params: 1
		}),
		new web3._extend.Method({
			name: 'writeBlockProfile',
			call: 'debug_writeBlockProfile',
			params: 1
		}),
		new web3._extend.Method({
			name: 'mutexProfile',
			call: 'debug_mutexProfile',
			params: 2
		}),
		new web3._extend.Method({
			name: 'setMutexProfileFraction',
			call: 'debug_setMutexProfileFraction',
			params: 1
		}),
		new web3._extend.Method({
			name: 'writeMutexProfile',
			call: 'debug_writeMutexProfile',
			params: 1
		}),
		new web3._extend.Method({
			name: 'writeMemProfile',
			call: 'debug_writeMemProfile',
			params: 1
		}),
		new web3._extend.Method({
			name: 'traceBlock',
			call: 'debug_traceBlock',
			params: 2,
			inputFormatter: [null, null]
		}),
		new web3._extend.Method({
			name: 'traceBlockFromFile',
			call: 'debug_traceBlockFromFile',
			params: 2,
			inputFormatter: [null, null]
		}),
		new web3._extend.Method({
			name: 'traceBadBlock',
			call: 'debug_traceBadBlock',
			params: 1,
			inputFormatter: [null]
		}),
		new web3._extend.Method({
			name: 'standardTraceBadBlockToFile',
			call: 'debug_standardTraceBadBlockToFile',
			params: 2,
			inputFormatter: [null, null]
		}),
		new web3._extend.Method({
			name: 'standardTraceBlockToFile',
			call: 'debug_standardTraceBlockToFile',
			params: 2,
			inputFormatter: [null, null]
		}),
		new web3._extend.Method({
			name: 'traceBlockByNumber',
			call: 'debug_traceBlockByNumber',
			params: 2,
			inputFormatter: [null, null]
		}),
		new web3._extend.Method({
			name: 'traceBlockByHash',
			call: 'debug_traceBlockByHash',
			params: 2,
			inputFormatter: [null, null]
		}),
		new web3._extend.Method({
			name: 'traceTransaction',
			call: 'debug_traceTransaction',
			params: 2,
			inputFormatter: [null, null]
		}),
		new web3._extend.Method({
			name: 'preimage',
			call: 'debug_preimage',
			params: 1,
			inputFormatter: [null]
		}),
		new web3._extend.Method({
			name: 'getBadBlocks',
			call: 'debug_getBadBlocks',
			params: 0,
		}),
		new web3._extend.Method({
			name: 'storageRangeAt',
			call: 'debug_storageRangeAt',
			params: 5,
		}),
		new web3._extend.Method({
			name: 'getModifiedAccountsByNumber',
			call: 'debug_getModifiedAccountsByNumber',
			params: 2,
			inputFormatter: [null, null],
		}),
		new web3._extend.Method({
			name: 'getModifiedAccountsByHash',
			call: 'debug_getModifiedAccountsByHash',
			params: 2,
			inputFormatter:[null, null],
		}),
		new web3._extend.Method({
			name: 'freezeClient',
			call: 'debug_freezeClient',
			params: 1,
		}),
	],
	properties: []
});
`

const EthJs = `
web3._extend({
	property: 'eth',
	methods: [
		new web3._extend.Method({
			name: 'chainId',
			call: 'eth_chainId',
			params: 0
		}),
		new web3._extend.Method({
			name: 'sign',
			call: 'eth_sign',
			params: 2,
			inputFormatter: [web3._extend.formatters.inputAddressFormatter, null]
		}),
		new web3._extend.Method({
			name: 'resend',
			call: 'eth_resend',
			params: 3,
			inputFormatter: [web3._extend.formatters.inputTransactionFormatter, web3._extend.utils.fromDecimal, web3._extend.utils.fromDecimal]
		}),
		new web3._extend.Method({
			name: 'signTransaction',
			call: 'eth_signTransaction',
			params: 1,
			inputFormatter: [web3._extend.formatters.inputTransactionFormatter]
		}),
		new web3._extend.Method({
			name: 'submitTransaction',
			call: 'eth_submitTransaction',
			params: 1,
			inputFormatter: [web3._extend.formatters.inputTransactionFormatter]
		}),
		new web3._extend.Method({
			name: 'fillTransaction',
			call: 'eth_fillTransaction',
			params: 1,
			inputFormatter: [web3._extend.formatters.inputTransactionFormatter]
		}),
		new web3._extend.Method({
			name: 'getHeaderByNumber',
			call: 'eth_getHeaderByNumber',
			params: 1
		}),
		new web3._extend.Method({
			name: 'getHeaderByHash',
			call: 'eth_getHeaderByHash',
			params: 1
		}),
		new web3._extend.Method({
			name: 'getBlockByNumber',
			call: 'eth_getBlockByNumber',
			params: 2
		}),
		new web3._extend.Method({
			name: 'getBlockByHash',
			call: 'eth_getBlockByHash',
			params: 2
		}),
		new web3._extend.Method({
			name: 'getRawTransaction',
			call: 'eth_getRawTransactionByHash',
			params: 1
		}),
		new web3._extend.Method({
			name: 'getRawTransactionFromBlock',
			call: function(args) {
				return (web3._extend.utils.isString(args[0]) && args[0].indexOf('0x') === 0) ? 'eth_getRawTransactionByBlockHashAndIndex' : 'eth_getRawTransactionByBlockNumberAndIndex';
			},
			params: 2,
			inputFormatter: [web3._extend.formatters.inputBlockNumberFormatter, web3._extend.utils.toHex]
		}),
		new web3._extend.Method({
			name: 'getProof',
			call: 'eth_getProof',
			params: 3,
			inputFormatter: [web3._extend.formatters.inputAddressFormatter, null, web3._extend.formatters.inputBlockNumberFormatter]
		}),
	],
	properties: [
		new web3._extend.Property({
			name: 'pendingTransactions',
			getter: 'eth_pendingTransactions',
			outputFormatter: function(txs) {
				var formatted = [];
				for (var i = 0; i < txs.length; i++) {
					formatted.push(web3._extend.formatters.outputTransactionFormatter(txs[i]));
					formatted[i].blockHash = null;
				}
				return formatted;
			}
		}),
	]
});
`
const BitbonJs = `
web3._extend({
	property: 'bitbon',
	methods: [
		new web3._extend.Method({
			name: 'version',
			call: 'bitbon_version'
		}),
		new web3._extend.Method({
			name: 'setContractStorageAddress',
			call: 'bitbon_setContractStorageAddress',
			params: 1,
			inputFormatter: [web3._extend.formatters.inputAddressFormatter]
		}),
		new web3._extend.Method({
			name: 'getContractAddresses',
			call: 'bitbon_getContractAddresses',
			params: 0,
		}),
		new web3._extend.Method({
			name: 'getServiceAddress',
			call: 'bitbon_getServiceAddress',
			params: 0,
		}),
		new web3._extend.Method({
			name: 'prepareAssetboxes',
			call: 'bitbon_prepareAssetboxes',
			params: 1,
			inputFormatter: [null]
		}),
		new web3._extend.Method({
			name: 'deleteAssetbox',
			call: 'bitbon_deleteAssetbox',
			params: 1,
			inputFormatter: [function (options){
				options.address = web3._extend.formatters.inputAddressFormatter(options.address);
    			return options;
			}]
		}),
		new web3._extend.Method({
			name: 'setPublicAssetboxInfo',
			call: 'bitbon_setPublicAssetboxInfo',
			params: 1,
			inputFormatter: [function (options){
				options.address = web3._extend.formatters.inputAddressFormatter(options.address);
    			return options;
			}]
		}),
		new web3._extend.Method({
			name: 'getPublicAssetboxInfo',
			call: 'bitbon_getPublicAssetboxInfo',
			params: 1,
			inputFormatter: [web3._extend.formatters.inputAddressFormatter]
		}),
		new web3._extend.Method({
			name: 'getAssetboxBalance',
			call: 'bitbon_getAssetboxBalance',
			params: 1,
			inputFormatter: [web3._extend.formatters.inputAddressFormatter],        
			outputFormatter: web3._extend.formatters.outputBigNumberFormatter
		}),
		new web3._extend.Method({
			name: 'balanceOfLocked',
			call: 'bitbon_balanceOfLocked',
			params: 1,
			inputFormatter: [web3._extend.formatters.inputAddressFormatter],        
			outputFormatter: web3._extend.formatters.outputBigNumberFormatter
		}),
		new web3._extend.Method({
			name: 'getAssetboxBalances',
			call: 'bitbon_getAssetboxBalances',
			params: 1,
			inputFormatter: [null],
			outputFormatter : function(result){
				for (var key in result) {
  					result[key] = web3._extend.formatters.outputBigNumberFormatter(result[key]);
				}
				return result;
			}
		}),
		new web3._extend.Method({
			name: 'getAssetboxBalancesSum',
			call: 'bitbon_getAssetboxBalancesSum',
			params: 1,
			inputFormatter: [null],        
			outputFormatter: web3._extend.formatters.outputBigNumberFormatter
		}),
		new web3._extend.Method({
			name: 'createSafeTransfer',
			call: 'bitbon_createSafeTransfer',
			params: 1,
			inputFormatter: [function (options){
    			options.from = web3._extend.formatters.inputAddressFormatter(options.from);
    			options.to = web3._extend.formatters.inputAddressFormatter(options.to);
    			['value'].filter(function (key) {
    			    return options[key] !== undefined;
    			}).forEach(function(key){
    			    options[key] = web3._extend.utils.fromDecimal(options[key]);
    			});
    			return options;
			}]			
		}),
		new web3._extend.Method({
			name: 'approveSafeTransfer',
			call: 'bitbon_approveSafeTransfer',
			params: 1,
			inputFormatter: [function (options){
				options.address = web3._extend.formatters.inputAddressFormatter(options.address);
				return options;
			}]
		}),
		new web3._extend.Method({
			name: 'cancelSafeTransfer',
			call: 'bitbon_cancelSafeTransfer',
			params: 1,
			inputFormatter: [function (options){
				options.address = web3._extend.formatters.inputAddressFormatter(options.address);
				return options;
			}]
		}),
		new web3._extend.Method({
			name: 'createFullBalanceSafeTransfer',
			call: 'bitbon_createFullBalanceSafeTransfer',
			params: 1,
			inputFormatter: [function (options){
    			options.from = web3._extend.formatters.inputAddressFormatter(options.from);
    			options.to = web3._extend.formatters.inputAddressFormatter(options.to);
    			return options;
			}]			
		}),
		new web3._extend.Method({
			name: 'approveFullBalanceSafeTransfer',
			call: 'bitbon_approveFullBalanceSafeTransfer',
			params: 1,
			inputFormatter: [function (options){
				options.address = web3._extend.formatters.inputAddressFormatter(options.address);
				return options;
			}]
		}),
		new web3._extend.Method({
			name: 'cancelFullBalanceSafeTransfer',
			call: 'bitbon_cancelFullBalanceSafeTransfer',
			params: 1,
			inputFormatter: [function (options){
				options.address = web3._extend.formatters.inputAddressFormatter(options.address);
				return options;
			}]
		}),
		new web3._extend.Method({
			name: 'expireTransfersAsync',
			call: 'bitbon_expireTransfersAsync',
		}),
		new web3._extend.Method({
			name: 'getTransfer',
			call: 'bitbon_getTransfer',
			params: 1
		}),
		new web3._extend.Method({
			name: 'getTransferByIndex',
			call: 'bitbon_getTransferByIndex',
			params: 1
		}),
		new web3._extend.Method({
			name: 'transferExists',
			call: 'bitbon_transferExists',
			params: 1
		}),
		new web3._extend.Method({
			name: 'directTransfer',
			call: 'bitbon_directTransfer',
			params: 1,
			inputFormatter: [function (options){
		    	options.from = web3._extend.formatters.inputAddressFormatter(options.from);
		    	options.to = web3._extend.formatters.inputAddressFormatter(options.to);
		    	['value'].filter(function (key) {
		    	    return options[key] !== undefined;
		    	}).forEach(function(key){
		    	    options[key] = web3._extend.utils.fromDecimal(options[key]);
		    	});
		    	return options;
			}]
		}),
		new web3._extend.Method({
			name: 'getNonces',
			call: 'bitbon_getNonces',
			params: 1,
			inputFormatter: [web3._extend.formatters.inputAddressFormatter]
		}),
		new web3._extend.Method({
			name: 'restoreServiceAccountNonce',
			call: 'bitbon_restoreServiceAccountNonce',
			params: 3,
		}),
		new web3._extend.Method({
			name: 'forceNoncerAccountNonce',
			call: 'bitbon_forceNoncerServiceAccountNonce',
			params: 3,
		}),
		new web3._extend.Method({
			name: 'forceNoncerNonceFromBlockChain',
			call: 'bitbon_forceNoncerNonceFromBlockChain',
			params: 1,
			inputFormatter: [web3._extend.formatters.inputAddressFormatter]
		}),
		new web3._extend.Method({
			name: 'removeAssetboxesFromNoncer',
			call: 'bitbon_removeAssetboxesFromNoncer',
			params: 1,
		}),
		new web3._extend.Method({
			name: 'getNoncerAssetboxes',
			call: 'bitbon_getNoncerAssetboxes',
			params: 0,
		}),
		new web3._extend.Method({
			name: 'blockByHash',
			call: 'bitbon_blockByHash',
			params: 1,
		
		}),
		new web3._extend.Method({
			name: 'blockByNumber',
			call: 'bitbon_blockByNumber',
			params: 1,
		}),
		new web3._extend.Method({
			name: 'blockByTime',
			call: 'bitbon_blockByTime',
			params: 1,
		}),
		new web3._extend.Method({
			name: 'blockByTime',
			call: 'bitbon_blockByTime',
			params: 1,
		}),
		new web3._extend.Method({
			name: 'rangeBlockByNumber',
			call: 'bitbon_rangeBlockByNumber',
			params: 1,
		}),
        new web3._extend.Method({
			name: 'blocksByTimePeriod',
			call: 'bitbon_blocksByTimePeriod',
			params: 1,
		}),
        new web3._extend.Method({
			name: 'blockTransactionCountByBlockHash',
			call: 'bitbon_blockTransactionCountByBlockHash',
			params: 1,
		}),
        new web3._extend.Method({
			name: 'blockTransactionCountByBlockNumber',
			call: 'bitbon_blockTransactionCountByBlockNumber',
			params: 1,
		}),
        new web3._extend.Method({
			name: 'transactionByBlockHashAndIndex',
			call: 'bitbon_transactionByBlockHashAndIndex',
			params: 1,
		}),
        new web3._extend.Method({
			name: 'transactionByBlockNumberAndIndex',
			call: 'bitbon_transactionByBlockNumberAndIndex',
			params: 1,
		}),
        new web3._extend.Method({
			name: 'transactionByHash',
			call: 'bitbon_transactionByHash',
			params: 1,
		}), 
		new web3._extend.Method({
			name: 'transactionsByTimePeriod',
			call: 'bitbon_transactionsByTimePeriod',
			params: 1,
		}),
		new web3._extend.Method({
			name: 'checkTransactionsByTimePeriod',
			call: 'bitbon_checkTransactionsByTimePeriod',
			params: 1,
		}),
        new web3._extend.Method({
			name: 'saveAbiInfo',
			call: 'bitbon_saveAbiInfo',
			params: 1,
		}),
        new web3._extend.Method({
			name: 'editAbiInfo',
			call: 'bitbon_editAbiInfo',
			params: 1,
		}),
        new web3._extend.Method({
			name: 'deleteAbiInfoByAddress',
			call: 'bitbon_deleteAbiInfoByAddress',
			params: 1,
		}),
        new web3._extend.Method({
			name: 'getAbiJson',
			call: 'bitbon_getAbiJson',
			params: 1,
		}),
        new web3._extend.Method({
			name: 'getAbiVersion',
			call: 'bitbon_getAbiVersion',
			params: 1,
		}),
        new web3._extend.Method({
			name: 'getFullAbiInfo',
			call: 'bitbon_getFullAbiInfo',
			params: 1,
		}),
        new web3._extend.Method({
			name: 'getAbiAddressesByVersion',
			call: 'bitbon_getAbiAddressesByVersion',
			params: 1,
		}),
        new web3._extend.Method({
			name: 'getFullAbiInfoByVersionAndType',
			call: 'bitbon_getFullAbiInfoByVersionAndType',
			params: 2,
		}),
        new web3._extend.Method({
			name: 'getAbiInfosByVersion',
			call: 'bitbon_getAbiInfosByVersion',
			params: 1,
		}),
		new web3._extend.Method({
			name: 'getCurrentDistribution',
			call: 'bitbon_getCurrentDistribution',
			params: 0,
		}),
		new web3._extend.Method({
			name: 'proposeDistribution',
			call: 'bitbon_proposeDistribution',
			params: 2,
			inputFormatter: [web3._extend.formatters.inputAddressFormatter, null]
		}),
		new web3._extend.Method({
			name: 'getAllMiningAgents',
			call: 'bitbon_getAllMiningAgents',
			params: 0,
		}),
		new web3._extend.Method({
			name: 'isMiningAgent',
			call: 'bitbon_isMiningAgent',
			params: 1,
			inputFormatter: [web3._extend.formatters.inputAddressFormatter]
		}),
		new web3._extend.Method({
			name: 'addMiningAgent',
			call: 'bitbon_addMiningAgent',
			params: 1,
			inputFormatter: [web3._extend.formatters.inputAddressFormatter]
		}),
		new web3._extend.Method({
			name: 'removeMiningAgent',
			call: 'bitbon_removeMiningAgent',
			params: 1,
			inputFormatter: [web3._extend.formatters.inputAddressFormatter]
		}),
		new web3._extend.Method({
			name: 'fillCurrentAbiInfo',
			call: 'bitbon_fillCurrentAbiInfo',
			params: 1,
		}),
		new web3._extend.Method({
			name: 'getFeeDistributionAccounts',
			call: 'bitbon_getFeeDistributionAccounts',
			params: 1,
		}),
		new web3._extend.Method({
			name: 'getFeeDistributionAmounts',
			call: 'bitbon_getFeeDistributionAmounts',
			params: 1,
		}),
		new web3._extend.Method({
			name: 'createWPCSafeTransfer',
			call: 'bitbon_createWPCSafeTransfer',
			params: 1,
			inputFormatter: [function (options){
    			options.from = web3._extend.formatters.inputAddressFormatter(options.from);
    			options.to = web3._extend.formatters.inputAddressFormatter(options.to);
    			['value'].filter(function (key) {
    			    return options[key] !== undefined;
    			}).forEach(function(key){
    			    options[key] = web3._extend.utils.fromDecimal(options[key]);
    			});
    			return options;
			}]			
		}),
		new web3._extend.Method({
			name: 'approveWPCSafeTransfer',
			call: 'bitbon_approveWPCSafeTransfer',
			params: 1,
			inputFormatter: [function (options){
				options.address = web3._extend.formatters.inputAddressFormatter(options.address);
				return options;
			}]
		}),
		new web3._extend.Method({
			name: 'cancelWPCSafeTransfer',
			call: 'bitbon_cancelWPCSafeTransfer',
			params: 1,
			inputFormatter: [function (options){
				options.address = web3._extend.formatters.inputAddressFormatter(options.address);
				return options;
			}]
		}),
		new web3._extend.Method({
			name: 'createFullBalanceWPCSafeTransfer',
			call: 'bitbon_createFullBalanceWPCSafeTransfer',
			params: 1,
			inputFormatter: [function (options){
    			options.from = web3._extend.formatters.inputAddressFormatter(options.from);
    			options.to = web3._extend.formatters.inputAddressFormatter(options.to);
    			return options;
			}]			
		}),
		new web3._extend.Method({
			name: 'approveFullBalanceWPCSafeTransfer',
			call: 'bitbon_approveFullBalanceWPCSafeTransfer',
			params: 1,
			inputFormatter: [function (options){
				options.address = web3._extend.formatters.inputAddressFormatter(options.address);
				return options;
			}]
		}),
		new web3._extend.Method({
			name: 'cancelFullBalanceWPCSafeTransfer',
			call: 'bitbon_cancelFullBalanceWPCSafeTransfer',
			params: 1,
			inputFormatter: [function (options){
				options.address = web3._extend.formatters.inputAddressFormatter(options.address);
				return options;
			}]
		}),
		new web3._extend.Method({
			name: 'quickTransfer',
			call: 'bitbon_quickTransfer',
			params: 1,
			inputFormatter: [function (options){
		    	options.from = web3._extend.formatters.inputAddressFormatter(options.from);
		    	options.to = web3._extend.formatters.inputAddressFormatter(options.to);
		    	['value'].filter(function (key) {
		    	    return options[key] !== undefined;
		    	}).forEach(function(key){
		    	    options[key] = web3._extend.utils.fromDecimal(options[key]);
		    	});
		    	return options;
			}]
		}),
		new web3._extend.Method({
			name: 'frameTransfer',
			call: 'bitbon_frameTransfer',
			params: 1,
			inputFormatter: [function (options){
		    	options.from = web3._extend.formatters.inputAddressFormatter(options.from);
		    	options.to = web3._extend.formatters.inputAddressFormatter(options.to);
		    	['value'].filter(function (key) {
		    	    return options[key] !== undefined;
		    	}).forEach(function(key){
		    	    options[key] = web3._extend.utils.fromDecimal(options[key]);
		    	});
		    	return options;
			}]
		}),
		new web3._extend.Method({
			name: 'getAsseboxesMiningState',
			call: 'bitbon_getAsseboxesMiningState',
			params: 1,
			inputFormatter: [null],
		}),
	],
	properties: []
});
`

const BitbonClientJs = `
web3._extend({
	property: 'bitbonclient',
	methods: [
		new web3._extend.Method({
			name: 'version',
			call: 'bitbonclient_version'
		}),
		new web3._extend.Method({
			name: 'getPoolSizes',
			call: 'bitbonclient_getPoolSizes',
			params: 0
		}),
		new web3._extend.Method({
			name: 'getPoolSize',
			call: 'bitbonclient_getPoolSize',
			params: 1
		}),
		new web3._extend.Method({
			name: 'setPoolSize',
			call: 'bitbonclient_setPoolSize',
			params: 2
		}),
		new web3._extend.Method({
			name: 'getJournal',
			call: 'bitbonclient_getJournal',
			params: 0
		}),
		new web3._extend.Method({
			name: 'getPoolErrors',
			call: 'bitbonclient_getPoolErrors',
			params: 0
		}),
		new web3._extend.Method({
			name: 'getPoolError',
			call: 'bitbonclient_getPoolError',
			params: 1
		}),
		new web3._extend.Method({
			name: 'getClientModes',
			call: 'bitbonclient_getClientModes',
			params: 0
		}),
		new web3._extend.Method({
			name: 'getCurrentMode',
			call: 'bitbonclient_getCurrentMode',
			params: 0
		}),
	],
	properties: [
		new web3._extend.Property({
			name: 'handlerKeys',
			getter: 'bitbonclient_getHandlerKeys'
		}),
	]
})
`

const BitbonAgentJs = `
web3._extend({
	property: 'bitbonagent',
	methods: [
		new web3._extend.Method({
			name: 'version',
			call: 'bitbonagent_version'
		}),
		new web3._extend.Method({
			name: 'getJournal',
			call: 'bitbonagent_getJournal',
			params: 0
		}),
	]
})
`

const QuorumJs = `
web3._extend({
	property: 'quorum',
	methods: [
		new web3._extend.Method({
			name: 'version',
			call: 'quorum_version'
		}),
		new web3._extend.Method({
			name: 'allNodes',
			call: 'quorum_allNodes'
		}),
		new web3._extend.Method({
			name: 'minerListAtRecentEpoch',
			call: 'quorum_minerListAtRecentEpoch'
		}),
		new web3._extend.Method({
			name: 'pingList',
			call: 'quorum_pingList'
		}),
		new web3._extend.Method({
			name: 'peerList',
			call: 'quorum_peerList'
		}),
		new web3._extend.Method({
			name: 'snapshot',
			call: 'quorum_snapshot',
			params: 1
		}),
		new web3._extend.Method({
			name: 'snapshotSystem',
			call: 'quorum_snapshotSystem',
		}),
		new web3._extend.Method({
			name: 'snapshotSystemEpochs',
			call: 'quorum_snapshotSystemEpochs',
		}),
		new web3._extend.Method({
			name: 'snapshotSystemRecentEpoch',
			call: 'quorum_snapshotSystemRecentEpoch',
		}),
		new web3._extend.Method({
			name: 'rateList',
 	 	 	call: 'quorum_rateList'
		}),
		new web3._extend.Method({
			name: 'changeRateRequest',
			call: 'quorum_changeRateRequest',
			params: 3
		}),
	]
})
`

const QuorumEngineJs = `
web3._extend({
	property: 'quorumengine',
	methods: [
		new web3._extend.Method({
			name: 'getSnapshot',
			call: 'quorumengine_getSnapshot',
			params: 1,
			inputFormatter: [null]
		}),
		new web3._extend.Method({
			name: 'getSnapshotAtHash',
			call: 'quorumengine_getSnapshotAtHash',
			params: 1
		}),
		new web3._extend.Method({
			name: 'getSigners',
			call: 'quorumengine_getSigners',
			params: 1,
			inputFormatter: [null]
		}),
		new web3._extend.Method({
			name: 'getSigner',
			call: 'quorumengine_getSigner',
		}),
	]
})
`

const MinerJs = `
web3._extend({
	property: 'miner',
	methods: [
		new web3._extend.Method({
			name: 'start',
			call: 'miner_start',
			params: 1,
			inputFormatter: [null]
		}),
		new web3._extend.Method({
			name: 'stop',
			call: 'miner_stop'
		}),
		new web3._extend.Method({
			name: 'setEtherbase',
			call: 'miner_setEtherbase',
			params: 1,
			inputFormatter: [web3._extend.formatters.inputAddressFormatter]
		}),
		new web3._extend.Method({
			name: 'setExtra',
			call: 'miner_setExtra',
			params: 1
		}),
		new web3._extend.Method({
			name: 'setGasPrice',
			call: 'miner_setGasPrice',
			params: 1,
			inputFormatter: [web3._extend.utils.fromDecimal]
		}),
		new web3._extend.Method({
			name: 'setRecommitInterval',
			call: 'miner_setRecommitInterval',
			params: 1,
		}),
		new web3._extend.Method({
			name: 'getHashrate',
			call: 'miner_getHashrate'
		}),
	],
	properties: []
});
`

const NetJs = `
web3._extend({
	property: 'net',
	methods: [],
	properties: [
		new web3._extend.Property({
			name: 'version',
			getter: 'net_version'
		}),
	]
});
`

const PersonalJs = `
web3._extend({
	property: 'personal',
	methods: [
		new web3._extend.Method({
			name: 'importRawKey',
			call: 'personal_importRawKey',
			params: 2
		}),
		new web3._extend.Method({
			name: 'sign',
			call: 'personal_sign',
			params: 3,
			inputFormatter: [null, web3._extend.formatters.inputAddressFormatter, null]
		}),
		new web3._extend.Method({
			name: 'ecRecover',
			call: 'personal_ecRecover',
			params: 2
		}),
		new web3._extend.Method({
			name: 'openWallet',
			call: 'personal_openWallet',
			params: 2
		}),
		new web3._extend.Method({
			name: 'deriveAccount',
			call: 'personal_deriveAccount',
			params: 3
		}),
		new web3._extend.Method({
			name: 'signTransaction',
			call: 'personal_signTransaction',
			params: 2,
			inputFormatter: [web3._extend.formatters.inputTransactionFormatter, null]
		}),
		new web3._extend.Method({
			name: 'unpair',
			call: 'personal_unpair',
			params: 2
		}),
		new web3._extend.Method({
			name: 'initializeWallet',
			call: 'personal_initializeWallet',
			params: 1
		})
	],
	properties: [
		new web3._extend.Property({
			name: 'listWallets',
			getter: 'personal_listWallets'
		}),
	]
})
`

const RpcJs = `
web3._extend({
	property: 'rpc',
	methods: [],
	properties: [
		new web3._extend.Property({
			name: 'modules',
			getter: 'rpc_modules'
		}),
	]
});
`

const ShhJs = `
web3._extend({
	property: 'shh',
	methods: [
	],
	properties:
	[
		new web3._extend.Property({
			name: 'version',
			getter: 'shh_version',
			outputFormatter: web3._extend.utils.toDecimal
		}),
		new web3._extend.Property({
			name: 'info',
			getter: 'shh_info'
		}),
	]
});
`

const SwarmfsJs = `
web3._extend({
	property: 'swarmfs',
	methods:
	[
		new web3._extend.Method({
			name: 'mount',
			call: 'swarmfs_mount',
			params: 2
		}),
		new web3._extend.Method({
			name: 'unmount',
			call: 'swarmfs_unmount',
			params: 1
		}),
		new web3._extend.Method({
			name: 'listmounts',
			call: 'swarmfs_listmounts',
			params: 0
		}),
	]
});
`

const TxpoolJs = `
web3._extend({
	property: 'txpool',
	methods: [],
	properties:
	[
		new web3._extend.Property({
			name: 'content',
			getter: 'txpool_content'
		}),
		new web3._extend.Property({
			name: 'inspect',
			getter: 'txpool_inspect'
		}),
		new web3._extend.Property({
			name: 'status',
			getter: 'txpool_status',
			outputFormatter: function(status) {
				status.pending = web3._extend.utils.toDecimal(status.pending);
				status.queued = web3._extend.utils.toDecimal(status.queued);
				return status;
			}
		}),
	]
});
`

const AccountingJs = `
web3._extend({
	property: 'accounting',
	methods: [
		new web3._extend.Property({
			name: 'balance',
			getter: 'account_balance'
		}),
		new web3._extend.Property({
			name: 'balanceCredit',
			getter: 'account_balanceCredit'
		}),
		new web3._extend.Property({
			name: 'balanceDebit',
			getter: 'account_balanceDebit'
		}),
		new web3._extend.Property({
			name: 'bytesCredit',
			getter: 'account_bytesCredit'
		}),
		new web3._extend.Property({
			name: 'bytesDebit',
			getter: 'account_bytesDebit'
		}),
		new web3._extend.Property({
			name: 'msgCredit',
			getter: 'account_msgCredit'
		}),
		new web3._extend.Property({
			name: 'msgDebit',
			getter: 'account_msgDebit'
		}),
		new web3._extend.Property({
			name: 'peerDrops',
			getter: 'account_peerDrops'
		}),
		new web3._extend.Property({
			name: 'selfDrops',
			getter: 'account_selfDrops'
		}),
	]
});
`

const LESJs = `
web3._extend({
	property: 'les',
	methods:
	[
		new web3._extend.Method({
			name: 'getCheckpoint',
			call: 'les_getCheckpoint',
			params: 1
		}),
		new web3._extend.Method({
			name: 'clientInfo',
			call: 'les_clientInfo',
			params: 1
		}),
		new web3._extend.Method({
			name: 'priorityClientInfo',
			call: 'les_priorityClientInfo',
			params: 3
		}),
		new web3._extend.Method({
			name: 'setClientParams',
			call: 'les_setClientParams',
			params: 2
		}),
		new web3._extend.Method({
			name: 'setDefaultParams',
			call: 'les_setDefaultParams',
			params: 1
		}),
		new web3._extend.Method({
			name: 'addBalance',
			call: 'les_addBalance',
			params: 3
		}),
	],
	properties:
	[
		new web3._extend.Property({
			name: 'latestCheckpoint',
			getter: 'les_latestCheckpoint'
		}),
		new web3._extend.Property({
			name: 'checkpointContractAddress',
			getter: 'les_getCheckpointContractAddress'
		}),
		new web3._extend.Property({
			name: 'serverInfo',
			getter: 'les_serverInfo'
		}),
	]
});
`

const LESPayJs = `
web3._extend({
	property: 'lespay',
	methods:
	[
		new web3._extend.Method({
			name: 'distribution',
			call: 'lespay_distribution',
			params: 2
		}),
		new web3._extend.Method({
			name: 'timeout',
			call: 'lespay_timeout',
			params: 2
		}),
		new web3._extend.Method({
			name: 'value',
			call: 'lespay_value',
			params: 2
		}),
	],
	properties:
	[
		new web3._extend.Property({
			name: 'requestStats',
			getter: 'lespay_requestStats'
		}),
	]
});
`
