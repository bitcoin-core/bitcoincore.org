---
name: getdifficulty
btcversion: 29.0.0
btcgroup: blockchain
permalink: en/doc/29.0.0/rpc/blockchain/getdifficulty/
---

getdifficulty

Returns the proof-of-work difficulty as a multiple of the minimum difficulty.

Result:
n    (numeric) the proof-of-work difficulty as a multiple of the minimum difficulty.

Examples:
> bitcoin-cli getdifficulty 
> curl --user myusername --data-binary '{"jsonrpc": "2.0", "id": "curltest", "method": "getdifficulty", "params": []}' -H 'content-type: application/json' http://127.0.0.1:8332/


