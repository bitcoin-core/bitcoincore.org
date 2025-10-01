---
name: listreceivedbyaddress
btcversion: 0.18.0
btcgroup: wallet
permalink: en/doc/0.18.0/rpc/wallet/listreceivedbyaddress/
---

listreceivedbyaddress ( minconf include_empty include_watchonly "address_filter" )

List balances by receiving address.

Arguments:
1. minconf              (numeric, optional, default=1) The minimum number of confirmations before payments are included.
2. include_empty        (boolean, optional, default=false) Whether to include addresses that haven't received any payments.
3. include_watchonly    (boolean, optional, default=false) Whether to include watch-only addresses (see 'importaddress').
4. address_filter       (string, optional) If present, only return information on this address.

Result:
[
  {
    "involvesWatchonly" : true,        (bool) Only returned if imported addresses were involved in transaction
    "address" : "receivingaddress",  (string) The receiving address
    "amount" : x.xxx,                  (numeric) The total amount in BTC received by the address
    "confirmations" : n,               (numeric) The number of confirmations of the most recent transaction included
    "label" : "label",               (string) The label of the receiving address. The default label is "".
    "txids": [
       "txid",                         (string) The ids of transactions received with the address 
       ...
    ]
  }
  ,...
]

Examples:
> bitcoin-cli listreceivedbyaddress 
> bitcoin-cli listreceivedbyaddress 6 true
> curl --user myusername --data-binary '{"jsonrpc": "1.0", "id":"curltest", "method": "listreceivedbyaddress", "params": [6, true, true] }' -H 'content-type: text/plain;' http://127.0.0.1:8332/
> curl --user myusername --data-binary '{"jsonrpc": "1.0", "id":"curltest", "method": "listreceivedbyaddress", "params": [6, true, true, "1M72Sfpbz1BPpXFHz9m3CdqATR44Jvaydd"] }' -H 'content-type: text/plain;' http://127.0.0.1:8332/


