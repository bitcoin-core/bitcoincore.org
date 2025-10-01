---
name: clearbanned
btcversion: 29.0.0
btcgroup: network
permalink: en/doc/29.0.0/rpc/network/clearbanned/
---

clearbanned

Clear all banned IPs.

Result:
null    (json null)

Examples:
> bitcoin-cli clearbanned 
> curl --user myusername --data-binary '{"jsonrpc": "2.0", "id": "curltest", "method": "clearbanned", "params": []}' -H 'content-type: application/json' http://127.0.0.1:8332/


