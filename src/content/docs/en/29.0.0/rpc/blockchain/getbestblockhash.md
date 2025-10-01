---
name: getbestblockhash
btcversion: 29.0.0
btcgroup: blockchain
permalink: en/doc/29.0.0/rpc/blockchain/getbestblockhash/
---

getbestblockhash

Returns the hash of the best (tip) block in the most-work fully-validated chain.

Result:
"hex"    (string) the block hash, hex-encoded

Examples:
> bitcoin-cli getbestblockhash 
> curl --user myusername --data-binary '{"jsonrpc": "2.0", "id": "curltest", "method": "getbestblockhash", "params": []}' -H 'content-type: application/json' http://127.0.0.1:8332/


