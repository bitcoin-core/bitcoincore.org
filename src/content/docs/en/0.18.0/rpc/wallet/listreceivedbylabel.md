---
name: listreceivedbylabel
btcversion: 0.18.0
btcgroup: wallet
permalink: en/doc/0.18.0/rpc/wallet/listreceivedbylabel/
---

listreceivedbylabel ( minconf include_empty include_watchonly )

List received transactions by label.

Arguments:
1. minconf              (numeric, optional, default=1) The minimum number of confirmations before payments are included.
2. include_empty        (boolean, optional, default=false) Whether to include labels that haven't received any payments.
3. include_watchonly    (boolean, optional, default=false) Whether to include watch-only addresses (see 'importaddress').

Result:
[
  {
    "involvesWatchonly" : true,   (bool) Only returned if imported addresses were involved in transaction
    "amount" : x.xxx,             (numeric) The total amount received by addresses with this label
    "confirmations" : n,          (numeric) The number of confirmations of the most recent transaction included
    "label" : "label"           (string) The label of the receiving address. The default label is "".
  }
  ,...
]

Examples:
> bitcoin-cli listreceivedbylabel 
> bitcoin-cli listreceivedbylabel 6 true
> curl --user myusername --data-binary '{"jsonrpc": "1.0", "id":"curltest", "method": "listreceivedbylabel", "params": [6, true, true] }' -H 'content-type: text/plain;' http://127.0.0.1:8332/


