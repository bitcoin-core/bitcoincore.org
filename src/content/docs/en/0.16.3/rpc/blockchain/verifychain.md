---
name: verifychain
btcversion: 0.16.3
btcgroup: blockchain
permalink: en/doc/0.16.3/rpc/blockchain/verifychain/
---

verifychain ( checklevel nblocks )

Verifies blockchain database.

Arguments:
1. checklevel   (numeric, optional, 0-4, default=3) How thorough the block verification is.
2. nblocks      (numeric, optional, default=6, 0=all) The number of blocks to check.

Result:
true|false       (boolean) Verified or not

Examples:
> bitcoin-cli verifychain 
> curl --user myusername --data-binary '{"jsonrpc": "1.0", "id":"curltest", "method": "verifychain", "params": [] }' -H 'content-type: text/plain;' http://127.0.0.1:8332/


