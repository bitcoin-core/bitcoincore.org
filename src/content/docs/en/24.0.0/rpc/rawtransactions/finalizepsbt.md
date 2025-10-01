---
name: finalizepsbt
btcversion: 24.0.0
btcgroup: rawtransactions
permalink: en/doc/24.0.0/rpc/rawtransactions/finalizepsbt/
---

finalizepsbt "psbt" ( extract )

Finalize the inputs of a PSBT. If the transaction is fully signed, it will produce a
network serialized transaction which can be broadcast with sendrawtransaction. Otherwise a PSBT will be
created which has the final_scriptSig and final_scriptWitness fields filled for inputs that are complete.
Implements the Finalizer and Extractor roles.

Arguments:
1. psbt       (string, required) A base64 string of a PSBT
2. extract    (boolean, optional, default=true) If true and the transaction is complete,
              extract and return the complete transaction in normal network serialization instead of the PSBT.

Result:
{                             (json object)
  "psbt" : "str",             (string, optional) The base64-encoded partially signed transaction if not extracted
  "hex" : "hex",              (string, optional) The hex-encoded network transaction if extracted
  "complete" : true|false     (boolean) If the transaction has a complete set of signatures
}

Examples:
> bitcoin-cli finalizepsbt "psbt"


