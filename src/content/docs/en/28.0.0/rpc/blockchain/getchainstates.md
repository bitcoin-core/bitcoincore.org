---
name: getchainstates
btcversion: 28.0.0
btcgroup: blockchain
permalink: en/doc/28.0.0/rpc/blockchain/getchainstates/
---

getchainstates

Return information about chainstates.

Result:
{                                      (json object)
  "headers" : n,                       (numeric) the number of headers seen so far
  "chainstates" : [                    (json array) list of the chainstates ordered by work, with the most-work (active) chainstate last
    {                                  (json object)
      "blocks" : n,                    (numeric) number of blocks in this chainstate
      "bestblockhash" : "hex",         (string) blockhash of the tip
      "difficulty" : n,                (numeric) difficulty of the tip
      "verificationprogress" : n,      (numeric) progress towards the network tip
      "snapshot_blockhash" : "hex",    (string, optional) the base block of the snapshot this chainstate is based on, if any
      "coins_db_cache_bytes" : n,      (numeric) size of the coinsdb cache
      "coins_tip_cache_bytes" : n,     (numeric) size of the coinstip cache
      "validated" : true|false         (boolean) whether the chainstate is fully validated. True if all blocks in the chainstate were validated, false if the chain is based on a snapshot and the snapshot has not yet been validated.
    },
    ...
  ]
}

Examples:
> bitcoin-cli getchainstates 
> curl --user myusername --data-binary '{"jsonrpc": "2.0", "id": "curltest", "method": "getchainstates", "params": []}' -H 'content-type: application/json' http://127.0.0.1:8332/


