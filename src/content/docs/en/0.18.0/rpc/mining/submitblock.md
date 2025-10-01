---
name: submitblock
btcversion: 0.18.0
btcgroup: mining
permalink: en/doc/0.18.0/rpc/mining/submitblock/
---

submitblock "hexdata" ( "dummy" )

Attempts to submit new block to network.
See https://en.bitcoin.it/wiki/BIP_0022 for full specification.

Arguments:
1. hexdata    (string, required) the hex-encoded block data to submit
2. dummy      (string, optional, default=ignored) dummy value, for compatibility with BIP22. This value is ignored.

Examples:
> bitcoin-cli submitblock "mydata"
> curl --user myusername --data-binary '{"jsonrpc": "1.0", "id":"curltest", "method": "submitblock", "params": ["mydata"] }' -H 'content-type: text/plain;' http://127.0.0.1:8332/


