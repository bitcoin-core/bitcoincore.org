---
name: listreceivedbyaddress
btcversion: 0.16.0
btcgroup: wallet
permalink: en/doc/0.16.0/rpc/wallet/listreceivedbyaddress/
---

listreceivedbyaddress ( minconf include_empty include_watchonly)

List balances by receiving address.

Arguments:
1. minconf           (numeric, optional, default=1) The minimum number of confirmations before payments are included.
2. include_empty     (bool, optional, default=false) Whether to include addresses that haven't received any payments.
3. include_watchonly (bool, optional, default=false) Whether to include watch-only addresses (see 'importaddress').

Result:
[
  {
    "involvesWatchonly" : true,        (bool) Only returned if imported addresses were involved in transaction
    "address" : "receivingaddress",  (string) The receiving address
    "account" : "accountname",       (string) DEPRECATED. The account of the receiving address. The default account is "".
    "amount" : x.xxx,                  (numeric) The total amount in BTC received by the address
    "confirmations" : n,               (numeric) The number of confirmations of the most recent transaction included
    "label" : "label",               (string) A comment for the address/transaction, if any
    "txids": [
       n,                                (numeric) The ids of transactions received with the address 
       ...
    ]
  }
  ,...
]

Examples:
> bitcoin-cli listreceivedbyaddress 
> bitcoin-cli listreceivedbyaddress 6 true
> curl --user myusername --data-binary '{"jsonrpc": "1.0", "id":"curltest", "method": "listreceivedbyaddress", "params": [6, true, true] }' -H 'content-type: text/plain;' http://127.0.0.1:8332/


