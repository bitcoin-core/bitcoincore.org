---
name: dumpprivkey
btcversion: 24.0.0
btcgroup: wallet
permalink: en/doc/24.0.0/rpc/wallet/dumpprivkey/
---

dumpprivkey "address"

Reveals the private key corresponding to 'address'.
Then the importprivkey can be used with this output

Arguments:
1. address    (string, required) The bitcoin address for the private key

Result:
"str"    (string) The private key

Examples:
> bitcoin-cli dumpprivkey "myaddress"
> bitcoin-cli importprivkey "mykey"
> curl --user myusername --data-binary '{"jsonrpc": "1.0", "id": "curltest", "method": "dumpprivkey", "params": ["myaddress"]}' -H 'content-type: text/plain;' http://127.0.0.1:8332/


