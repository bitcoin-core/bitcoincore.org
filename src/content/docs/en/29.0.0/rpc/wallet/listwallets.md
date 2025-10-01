---
name: listwallets
btcversion: 29.0.0
btcgroup: wallet
permalink: en/doc/29.0.0/rpc/wallet/listwallets/
---

listwallets

Returns a list of currently loaded wallets.
For full information on the wallet, use "getwalletinfo"

Result:
[           (json array)
  "str",    (string) the wallet name
  ...
]

Examples:
> bitcoin-cli listwallets 
> curl --user myusername --data-binary '{"jsonrpc": "2.0", "id": "curltest", "method": "listwallets", "params": []}' -H 'content-type: application/json' http://127.0.0.1:8332/


