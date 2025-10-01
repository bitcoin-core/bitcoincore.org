---
name: keypoolrefill
btcversion: 29.0.0
btcgroup: wallet
permalink: en/doc/29.0.0/rpc/wallet/keypoolrefill/
---

keypoolrefill ( newsize )

Fills the keypool.
Requires wallet passphrase to be set with walletpassphrase call if wallet is encrypted.

Arguments:
1. newsize    (numeric, optional, default=1000, or as set by -keypool) The new keypool size

Result:
null    (json null)

Examples:
> bitcoin-cli keypoolrefill 
> curl --user myusername --data-binary '{"jsonrpc": "2.0", "id": "curltest", "method": "keypoolrefill", "params": []}' -H 'content-type: application/json' http://127.0.0.1:8332/


