---
name: lockunspent
btcversion: 0.16.2
btcgroup: wallet
permalink: en/doc/0.16.2/rpc/wallet/lockunspent/
---

lockunspent unlock ([{"txid":"txid","vout":n},...])

Updates list of temporarily unspendable outputs.
Temporarily lock (unlock=false) or unlock (unlock=true) specified transaction outputs.
If no transaction outputs are specified when unlocking then all current locked transaction outputs are unlocked.
A locked transaction output will not be chosen by automatic coin selection, when spending bitcoins.
Locks are stored in memory only. Nodes start with zero locked outputs, and the locked output list
is always cleared (by virtue of process exit) when a node stops or fails.
Also see the listunspent call

Arguments:
1. unlock            (boolean, required) Whether to unlock (true) or lock (false) the specified transactions
2. "transactions"  (string, optional) A json array of objects. Each object the txid (string) vout (numeric)
     [           (json array of json objects)
       {
         "txid":"id",    (string) The transaction id
         "vout": n         (numeric) The output number
       }
       ,...
     ]

Result:
true|false    (boolean) Whether the command was successful or not

Examples:

List the unspent transactions
> bitcoin-cli listunspent 

Lock an unspent transaction
> bitcoin-cli lockunspent false "[{\"txid\":\"a08e6907dbbd3d809776dbfc5d82e371b764ed838b5655e72f463568df1aadf0\",\"vout\":1}]"

List the locked transactions
> bitcoin-cli listlockunspent 

Unlock the transaction again
> bitcoin-cli lockunspent true "[{\"txid\":\"a08e6907dbbd3d809776dbfc5d82e371b764ed838b5655e72f463568df1aadf0\",\"vout\":1}]"

As a json rpc call
> curl --user myusername --data-binary '{"jsonrpc": "1.0", "id":"curltest", "method": "lockunspent", "params": [false, "[{\"txid\":\"a08e6907dbbd3d809776dbfc5d82e371b764ed838b5655e72f463568df1aadf0\",\"vout\":1}]"] }' -H 'content-type: text/plain;' http://127.0.0.1:8332/


