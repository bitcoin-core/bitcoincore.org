---
name: listlockunspent
btcversion: 24.0.0
btcgroup: wallet
permalink: en/doc/24.0.0/rpc/wallet/listlockunspent/
---

listlockunspent

Returns list of temporarily unspendable outputs.
See the lockunspent call to lock and unlock transactions for spending.

Result:
[                      (json array)
  {                    (json object)
    "txid" : "hex",    (string) The transaction id locked
    "vout" : n         (numeric) The vout value
  },
  ...
]

Examples:

List the unspent transactions
> bitcoin-cli listunspent 

Lock an unspent transaction
> bitcoin-cli lockunspent false "[{\"txid\":\"a08e6907dbbd3d809776dbfc5d82e371b764ed838b5655e72f463568df1aadf0\",\"vout\":1}]"

List the locked transactions
> bitcoin-cli listlockunspent 

Unlock the transaction again
> bitcoin-cli lockunspent true "[{\"txid\":\"a08e6907dbbd3d809776dbfc5d82e371b764ed838b5655e72f463568df1aadf0\",\"vout\":1}]"

As a JSON-RPC call
> curl --user myusername --data-binary '{"jsonrpc": "1.0", "id": "curltest", "method": "listlockunspent", "params": []}' -H 'content-type: text/plain;' http://127.0.0.1:8332/


