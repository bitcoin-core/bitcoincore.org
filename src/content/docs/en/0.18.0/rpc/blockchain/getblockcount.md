---
name: getblockcount
btcversion: 0.18.0
btcgroup: blockchain
permalink: en/doc/0.18.0/rpc/blockchain/getblockcount/
---

getblockcount

Returns the number of blocks in the longest blockchain.

Result:
n    (numeric) The current block count

Examples:
> bitcoin-cli getblockcount 
> curl --user myusername --data-binary '{"jsonrpc": "1.0", "id":"curltest", "method": "getblockcount", "params": [] }' -H 'content-type: text/plain;' http://127.0.0.1:8332/


