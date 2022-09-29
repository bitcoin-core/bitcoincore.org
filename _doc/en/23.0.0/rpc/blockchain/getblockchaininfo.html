---
name: getblockchaininfo
btcversion: 23.0.0
btcgroup: blockchain
permalink: en/doc/23.0.0/rpc/blockchain/getblockchaininfo/
---

getblockchaininfo

Returns an object containing various state info regarding blockchain processing.

Result:
{                                         (json object)
  "chain" : "str",                        (string) current network name (main, test, signet, regtest)
  "blocks" : n,                           (numeric) the height of the most-work fully-validated chain. The genesis block has height 0
  "headers" : n,                          (numeric) the current number of headers we have validated
  "bestblockhash" : "str",                (string) the hash of the currently best block
  "difficulty" : n,                       (numeric) the current difficulty
  "time" : xxx,                           (numeric) The block time expressed in UNIX epoch time
  "mediantime" : xxx,                     (numeric) The median block time expressed in UNIX epoch time
  "verificationprogress" : n,             (numeric) estimate of verification progress [0..1]
  "initialblockdownload" : true|false,    (boolean) (debug information) estimate of whether this node is in Initial Block Download mode
  "chainwork" : "hex",                    (string) total amount of work in active chain, in hexadecimal
  "size_on_disk" : n,                     (numeric) the estimated size of the block and undo files on disk
  "pruned" : true|false,                  (boolean) if the blocks are subject to pruning
  "pruneheight" : n,                      (numeric, optional) lowest-height complete block stored (only present if pruning is enabled)
  "automatic_pruning" : true|false,       (boolean, optional) whether automatic pruning is enabled (only present if pruning is enabled)
  "prune_target_size" : n,                (numeric, optional) the target size used by pruning (only present if automatic pruning is enabled)
  "softforks" : {                         (json object) (DEPRECATED, returned only if config option -deprecatedrpc=softforks is passed) status of softforks
    "xxxx" : {                            (json object) name of the softfork
      "type" : "str",                     (string) one of "buried", "bip9"
      "height" : n,                       (numeric, optional) height of the first block which the rules are or will be enforced (only for "buried" type, or "bip9" type with "active" status)
      "active" : true|false,              (boolean) true if the rules are enforced for the mempool and the next block
      "bip9" : {                          (json object, optional) status of bip9 softforks (only for "bip9" type)
        "bit" : n,                        (numeric, optional) the bit (0-28) in the block version field used to signal this softfork (only for "started" and "locked_in" status)
        "start_time" : xxx,               (numeric) the minimum median time past of a block at which the bit gains its meaning
        "timeout" : xxx,                  (numeric) the median time past of a block at which the deployment is considered failed if not yet locked in
        "min_activation_height" : n,      (numeric) minimum height of blocks for which the rules may be enforced
        "status" : "str",                 (string) status of deployment at specified block (one of "defined", "started", "locked_in", "active", "failed")
        "since" : n,                      (numeric) height of the first block to which the status applies
        "status_next" : "str",            (string) status of deployment at the next block
        "statistics" : {                  (json object, optional) numeric statistics about signalling for a softfork (only for "started" and "locked_in" status)
          "period" : n,                   (numeric) the length in blocks of the signalling period
          "threshold" : n,                (numeric, optional) the number of blocks with the version bit set required to activate the feature (only for "started" status)
          "elapsed" : n,                  (numeric) the number of blocks elapsed since the beginning of the current period
          "count" : n,                    (numeric) the number of blocks with the version bit set in the current period
          "possible" : true|false         (boolean, optional) returns false if there are not enough blocks left in this period to pass activation threshold (only for "started" status)
        },
        "signalling" : "str"              (string) indicates blocks that signalled with a # and blocks that did not with a -
      }
    },
    ...
  },
  "warnings" : "str"                      (string) any network and blockchain warnings
}

Examples:
> bitcoin-cli getblockchaininfo 
> curl --user myusername --data-binary '{"jsonrpc": "1.0", "id": "curltest", "method": "getblockchaininfo", "params": []}' -H 'content-type: text/plain;' http://127.0.0.1:8332/


