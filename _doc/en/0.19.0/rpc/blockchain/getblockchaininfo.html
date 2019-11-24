---
name: getblockchaininfo
btcversion: 0.19.0
btcgroup: blockchain
permalink: en/doc/0.19.0/rpc/blockchain/getblockchaininfo/
---

getblockchaininfo
Returns an object containing various state info regarding blockchain processing.

Result:
{
  "chain": "xxxx",              (string) current network name as defined in BIP70 (main, test, regtest)
  "blocks": xxxxxx,             (numeric) the height of the most-work fully-validated chain. The genesis block has height 0
  "headers": xxxxxx,            (numeric) the current number of headers we have validated
  "bestblockhash": "...",       (string) the hash of the currently best block
  "difficulty": xxxxxx,         (numeric) the current difficulty
  "mediantime": xxxxxx,         (numeric) median time for the current best block
  "verificationprogress": xxxx, (numeric) estimate of verification progress [0..1]
  "initialblockdownload": xxxx, (bool) (debug information) estimate of whether this node is in Initial Block Download mode.
  "chainwork": "xxxx"           (string) total amount of work in active chain, in hexadecimal
  "size_on_disk": xxxxxx,       (numeric) the estimated size of the block and undo files on disk
  "pruned": xx,                 (boolean) if the blocks are subject to pruning
  "pruneheight": xxxxxx,        (numeric) lowest-height complete block stored (only present if pruning is enabled)
  "automatic_pruning": xx,      (boolean) whether automatic pruning is enabled (only present if pruning is enabled)
  "prune_target_size": xxxxxx,  (numeric) the target size used by pruning (only present if automatic pruning is enabled)
  "softforks": {                (object) status of softforks
     "xxxx" : {                 (string) name of the softfork
        "type": "xxxx",         (string) one of "buried", "bip9"
        "bip9": {               (object) status of bip9 softforks (only for "bip9" type)
           "status": "xxxx",    (string) one of "defined", "started", "locked_in", "active", "failed"
           "bit": xx,           (numeric) the bit (0-28) in the block version field used to signal this softfork (only for "started" status)
           "start_time": xx,     (numeric) the minimum median time past of a block at which the bit gains its meaning
           "timeout": xx,       (numeric) the median time past of a block at which the deployment is considered failed if not yet locked in
           "since": xx,         (numeric) height of the first block to which the status applies
           "statistics": {      (object) numeric statistics about BIP9 signalling for a softfork
              "period": xx,     (numeric) the length in blocks of the BIP9 signalling period 
              "threshold": xx,  (numeric) the number of blocks with the version bit set required to activate the feature 
              "elapsed": xx,    (numeric) the number of blocks elapsed since the beginning of the current period 
              "count": xx,      (numeric) the number of blocks with the version bit set in the current period 
              "possible": xx    (boolean) returns false if there are not enough blocks left in this period to pass activation threshold 
           }
        },
        "height": "xxxxxx",     (numeric) height of the first block which the rules are or will be enforced (only for "buried" type, or "bip9" type with "active" status)
        "active": xx,           (boolean) true if the rules are enforced for the mempool and the next block
     }
  }
  "warnings" : "...",           (string) any network and blockchain warnings.
}

Examples:
> bitcoin-cli getblockchaininfo 
> curl --user myusername --data-binary '{"jsonrpc": "1.0", "id":"curltest", "method": "getblockchaininfo", "params": [] }' -H 'content-type: text/plain;' http://127.0.0.1:8332/


