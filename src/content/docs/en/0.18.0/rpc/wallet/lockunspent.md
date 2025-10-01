---
name: lockunspent
btcversion: 0.18.0
btcgroup: wallet
permalink: en/doc/0.18.0/rpc/wallet/lockunspent/
---

lockunspent unlock ( [{"txid":"hex","vout":n},...] )

Updates list of temporarily unspendable outputs.
Temporarily lock (unlock=false) or unlock (unlock=true) specified transaction outputs.
If no transaction outputs are specified when unlocking then all current locked transaction outputs are unlocked.
A locked transaction output will not be chosen by automatic coin selection, when spending bitcoins.
Locks are stored in memory only. Nodes start with zero locked outputs, and the locked output list
is always cleared (by virtue of process exit) when a node stops or fails.
Also see the listunspent call

Arguments:
1. unlock                  (boolean, required) Whether to unlock (true) or lock (false) the specified transactions
2. transactions            (json array, optional, default=empty array) A json array of objects. Each object the txid (string) vout (numeric).
     [
       {                   (json object)
         "txid": "hex",    (string, required) The transaction id
         "vout": n,        (numeric, required) The output number
       },
       ...
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

As a JSON-RPC call
> curl --user myusername --data-binary '{"jsonrpc": "1.0", "id":"curltest", "method": "lockunspent", "params": [false, "[{\"txid\":\"a08e6907dbbd3d809776dbfc5d82e371b764ed838b5655e72f463568df1aadf0\",\"vout\":1}]"] }' -H 'content-type: text/plain;' http://127.0.0.1:8332/


