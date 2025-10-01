---
name: clearbanned
btcversion: 0.18.0
btcgroup: network
permalink: en/doc/0.18.0/rpc/network/clearbanned/
---

clearbanned

Clear all banned IPs.

Examples:
> bitcoin-cli clearbanned 
> curl --user myusername --data-binary '{"jsonrpc": "1.0", "id":"curltest", "method": "clearbanned", "params": [] }' -H 'content-type: text/plain;' http://127.0.0.1:8332/


