---
name: getblockchaininfo
btcversion: 29.0.0
btcgroup: blockchain
permalink: en/doc/29.0.0/rpc/blockchain/getblockchaininfo/
---

getblockchaininfo

Returns an object containing various state info regarding blockchain processing.

Result:
{                                         (json object)
  "chain" : "str",                        (string) current network name (main, test, testnet4, signet, regtest)
  "blocks" : n,                           (numeric) the height of the most-work fully-validated chain. The genesis block has height 0
  "headers" : n,                          (numeric) the current number of headers we have validated
  "bestblockhash" : "str",                (string) the hash of the currently best block
  "bits" : "hex",                         (string) nBits: compact representation of the block difficulty target
  "target" : "hex",                       (string) The difficulty target
  "difficulty" : n,                       (numeric) the current difficulty
  "time" : xxx,                           (numeric) The block time expressed in UNIX epoch time
  "mediantime" : xxx,                     (numeric) The median block time expressed in UNIX epoch time
  "verificationprogress" : n,             (numeric) estimate of verification progress [0..1]
  "initialblockdownload" : true|false,    (boolean) (debug information) estimate of whether this node is in Initial Block Download mode
  "chainwork" : "hex",                    (string) total amount of work in active chain, in hexadecimal
  "size_on_disk" : n,                     (numeric) the estimated size of the block and undo files on disk
  "pruned" : true|false,                  (boolean) if the blocks are subject to pruning
  "pruneheight" : n,                      (numeric, optional) height of the last block pruned, plus one (only present if pruning is enabled)
  "automatic_pruning" : true|false,       (boolean, optional) whether automatic pruning is enabled (only present if pruning is enabled)
  "prune_target_size" : n,                (numeric, optional) the target size used by pruning (only present if automatic pruning is enabled)
  "signet_challenge" : "hex",             (string, optional) the block challenge (aka. block script), in hexadecimal (only present if the current network is a signet)
  "warnings" : [                          (json array) any network and blockchain warnings (run with `-deprecatedrpc=warnings` to return the latest warning as a single string)
    "str",                                (string) warning
    ...
  ]
}

Examples:
> bitcoin-cli getblockchaininfo 
> curl --user myusername --data-binary '{"jsonrpc": "2.0", "id": "curltest", "method": "getblockchaininfo", "params": []}' -H 'content-type: application/json' http://127.0.0.1:8332/


