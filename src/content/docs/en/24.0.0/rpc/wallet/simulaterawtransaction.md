---
name: simulaterawtransaction
btcversion: 24.0.0
btcgroup: wallet
permalink: en/doc/24.0.0/rpc/wallet/simulaterawtransaction/
---

simulaterawtransaction ( ["rawtx",...] {"include_watchonly":bool,...} )

Calculate the balance change resulting in the signing and broadcasting of the given transaction(s).

Arguments:
1. rawtxs                            (json array, optional) An array of hex strings of raw transactions.
                                     
     [
       "rawtx",                      (string)
       ...
     ]
2. options                           (json object, optional) Options
     {
       "include_watchonly": bool,    (boolean, optional, default=true for watch-only wallets, otherwise false) Whether to include watch-only addresses (see RPC importaddress)
       ...
     }

Result:
{                          (json object)
  "balance_change" : n     (numeric) The wallet balance change (negative means decrease).
}

Examples:
> bitcoin-cli simulaterawtransaction ["myhex"]
> curl --user myusername --data-binary '{"jsonrpc": "1.0", "id": "curltest", "method": "simulaterawtransaction", "params": [["myhex"]]}' -H 'content-type: text/plain;' http://127.0.0.1:8332/


