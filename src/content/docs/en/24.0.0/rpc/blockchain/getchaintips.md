---
name: getchaintips
btcversion: 24.0.0
btcgroup: blockchain
permalink: en/doc/24.0.0/rpc/blockchain/getchaintips/
---

getchaintips

Return information about all known tips in the block tree, including the main chain as well as orphaned branches.

Result:
[                        (json array)
  {                      (json object)
    "height" : n,        (numeric) height of the chain tip
    "hash" : "hex",      (string) block hash of the tip
    "branchlen" : n,     (numeric) zero for main chain, otherwise length of branch connecting the tip to the main chain
    "status" : "str"     (string) status of the chain, "active" for the main chain
                         Possible values for status:
                         1.  "invalid"               This branch contains at least one invalid block
                         2.  "headers-only"          Not all blocks for this branch are available, but the headers are valid
                         3.  "valid-headers"         All blocks are available for this branch, but they were never fully validated
                         4.  "valid-fork"            This branch is not part of the active chain, but is fully validated
                         5.  "active"                This is the tip of the active main chain, which is certainly valid
  },
  ...
]

Examples:
> bitcoin-cli getchaintips 
> curl --user myusername --data-binary '{"jsonrpc": "1.0", "id": "curltest", "method": "getchaintips", "params": []}' -H 'content-type: text/plain;' http://127.0.0.1:8332/


