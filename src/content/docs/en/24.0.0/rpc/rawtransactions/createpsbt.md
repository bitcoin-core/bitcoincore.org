---
name: createpsbt
btcversion: 24.0.0
btcgroup: rawtransactions
permalink: en/doc/24.0.0/rpc/rawtransactions/createpsbt/
---

createpsbt [{"txid":"hex","vout":n,"sequence":n},...] [{"address":amount,...},{"data":"hex"},...] ( locktime replaceable )

Creates a transaction in the Partially Signed Transaction format.
Implements the Creator role.

Arguments:
1. inputs                      (json array, required) The inputs
     [
       {                       (json object)
         "txid": "hex",        (string, required) The transaction id
         "vout": n,            (numeric, required) The output number
         "sequence": n,        (numeric, optional, default=depends on the value of the 'replaceable' and 'locktime' arguments) The sequence number
       },
       ...
     ]
2. outputs                     (json array, required) The outputs (key-value pairs), where none of the keys are duplicated.
                               That is, each address can only appear once and there can only be one 'data' object.
                               For compatibility reasons, a dictionary, which holds the key-value pairs directly, is also
                               accepted as second parameter.
     [
       {                       (json object)
         "address": amount,    (numeric or string, required) A key-value pair. The key (string) is the bitcoin address, the value (float or string) is the amount in BTC
         ...
       },
       {                       (json object)
         "data": "hex",        (string, required) A key-value pair. The key must be "data", the value is hex-encoded data
       },
       ...
     ]
3. locktime                    (numeric, optional, default=0) Raw locktime. Non-0 value also locktime-activates inputs
4. replaceable                 (boolean, optional, default=true) Marks this transaction as BIP125-replaceable.
                               Allows this transaction to be replaced by a transaction with higher fees. If provided, it is an error if explicit sequence numbers are incompatible.

Result:
"str"    (string) The resulting raw transaction (base64-encoded string)

Examples:
> bitcoin-cli createpsbt "[{\"txid\":\"myid\",\"vout\":0}]" "[{\"data\":\"00010203\"}]"


