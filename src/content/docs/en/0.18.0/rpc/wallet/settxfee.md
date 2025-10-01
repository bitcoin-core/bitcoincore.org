---
name: settxfee
btcversion: 0.18.0
btcgroup: wallet
permalink: en/doc/0.18.0/rpc/wallet/settxfee/
---

settxfee amount

Set the transaction fee per kB for this wallet. Overrides the global -paytxfee command line parameter.

Arguments:
1. amount    (numeric or string, required) The transaction fee in BTC/kB

Result:
true|false        (boolean) Returns true if successful

Examples:
> bitcoin-cli settxfee 0.00001
> curl --user myusername --data-binary '{"jsonrpc": "1.0", "id":"curltest", "method": "settxfee", "params": [0.00001] }' -H 'content-type: text/plain;' http://127.0.0.1:8332/


