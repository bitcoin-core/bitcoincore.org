---
name: verifychain
btcversion: 24.0.0
btcgroup: blockchain
permalink: en/doc/24.0.0/rpc/blockchain/verifychain/
---

verifychain ( checklevel nblocks )

Verifies blockchain database.

Arguments:
1. checklevel    (numeric, optional, default=3, range=0-4) How thorough the block verification is:
                 - level 0 reads the blocks from disk
                 - level 1 verifies block validity
                 - level 2 verifies undo data
                 - level 3 checks disconnection of tip blocks
                 - level 4 tries to reconnect the blocks
                 - each level includes the checks of the previous levels
2. nblocks       (numeric, optional, default=6, 0=all) The number of blocks to check.

Result:
true|false    (boolean) Verified or not

Examples:
> bitcoin-cli verifychain 
> curl --user myusername --data-binary '{"jsonrpc": "1.0", "id": "curltest", "method": "verifychain", "params": []}' -H 'content-type: text/plain;' http://127.0.0.1:8332/


