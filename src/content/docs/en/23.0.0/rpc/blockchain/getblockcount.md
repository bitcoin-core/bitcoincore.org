---
name: getblockcount
btcversion: 23.0.0
btcgroup: blockchain
permalink: en/doc/23.0.0/rpc/blockchain/getblockcount/
---

getblockcount

Returns the height of the most-work fully-validated chain.
The genesis block has height 0.

Result:
n    (numeric) The current block count

Examples:
> bitcoin-cli getblockcount 
> curl --user myusername --data-binary '{"jsonrpc": "1.0", "id": "curltest", "method": "getblockcount", "params": []}' -H 'content-type: text/plain;' http://127.0.0.1:8332/


