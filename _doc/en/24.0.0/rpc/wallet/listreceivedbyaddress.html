---
name: listreceivedbyaddress
btcversion: 24.0.0
btcgroup: wallet
permalink: en/doc/24.0.0/rpc/wallet/listreceivedbyaddress/
---

listreceivedbyaddress ( minconf include_empty include_watchonly "address_filter" include_immature_coinbase )

List balances by receiving address.

Arguments:
1. minconf                      (numeric, optional, default=1) The minimum number of confirmations before payments are included.
2. include_empty                (boolean, optional, default=false) Whether to include addresses that haven't received any payments.
3. include_watchonly            (boolean, optional, default=true for watch-only wallets, otherwise false) Whether to include watch-only addresses (see 'importaddress')
4. address_filter               (string, optional) If present and non-empty, only return information on this address.
5. include_immature_coinbase    (boolean, optional, default=false) Include immature coinbase transactions.

Result:
[                                        (json array)
  {                                      (json object)
    "involvesWatchonly" : true|false,    (boolean, optional) Only returns true if imported addresses were involved in transaction
    "address" : "str",                   (string) The receiving address
    "amount" : n,                        (numeric) The total amount in BTC received by the address
    "confirmations" : n,                 (numeric) The number of confirmations of the most recent transaction included
    "label" : "str",                     (string) The label of the receiving address. The default label is ""
    "txids" : [                          (json array)
      "hex",                             (string) The ids of transactions received with the address
      ...
    ]
  },
  ...
]

Examples:
> bitcoin-cli listreceivedbyaddress 
> bitcoin-cli listreceivedbyaddress 6 true
> bitcoin-cli listreceivedbyaddress 6 true true "" true
> curl --user myusername --data-binary '{"jsonrpc": "1.0", "id": "curltest", "method": "listreceivedbyaddress", "params": [6, true, true]}' -H 'content-type: text/plain;' http://127.0.0.1:8332/
> curl --user myusername --data-binary '{"jsonrpc": "1.0", "id": "curltest", "method": "listreceivedbyaddress", "params": [6, true, true, "bc1q09vm5lfy0j5reeulh4x5752q25uqqvz34hufdl", true]}' -H 'content-type: text/plain;' http://127.0.0.1:8332/


