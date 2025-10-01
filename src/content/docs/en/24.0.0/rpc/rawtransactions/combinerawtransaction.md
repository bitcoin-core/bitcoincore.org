---
name: combinerawtransaction
btcversion: 24.0.0
btcgroup: rawtransactions
permalink: en/doc/24.0.0/rpc/rawtransactions/combinerawtransaction/
---

combinerawtransaction ["hexstring",...]

Combine multiple partially signed transactions into one transaction.
The combined transaction may be another partially signed transaction or a 
fully signed transaction.

Arguments:
1. txs                 (json array, required) The hex strings of partially signed transactions
     [
       "hexstring",    (string) A hex-encoded raw transaction
       ...
     ]

Result:
"str"    (string) The hex-encoded raw transaction with signature(s)

Examples:
> bitcoin-cli combinerawtransaction '["myhex1", "myhex2", "myhex3"]'


