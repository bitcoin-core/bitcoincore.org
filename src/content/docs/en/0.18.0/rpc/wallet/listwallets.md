---
name: listwallets
btcversion: 0.18.0
btcgroup: wallet
permalink: en/doc/0.18.0/rpc/wallet/listwallets/
---

listwallets
Returns a list of currently loaded wallets.
For full information on the wallet, use "getwalletinfo"

Result:
[                         (json array of strings)
  "walletname"            (string) the wallet name
   ...
]

Examples:
> bitcoin-cli listwallets 
> curl --user myusername --data-binary '{"jsonrpc": "1.0", "id":"curltest", "method": "listwallets", "params": [] }' -H 'content-type: text/plain;' http://127.0.0.1:8332/


