---
name: descriptorprocesspsbt
btcversion: 27.0.0
btcgroup: rawtransactions
permalink: en/doc/27.0.0/rpc/rawtransactions/descriptorprocesspsbt/
---

descriptorprocesspsbt "psbt" ["",{"desc":"str","range":n or [n,n]},...] ( "sighashtype" bip32derivs finalize )

Update all segwit inputs in a PSBT with information from output descriptors, the UTXO set or the mempool. 
Then, sign the inputs we are able to with information from the output descriptors.

Arguments:
1. psbt                          (string, required) The transaction base64 string
2. descriptors                   (json array, required) An array of either strings or objects
     [
       "",                       (string) An output descriptor
       {                         (json object) An object with an output descriptor and extra information
         "desc": "str",          (string, required) An output descriptor
         "range": n or [n,n],    (numeric or array, optional, default=1000) Up to what index HD chains should be explored (either end or [begin,end])
       },
       ...
     ]
3. sighashtype                   (string, optional, default="DEFAULT for Taproot, ALL otherwise") The signature hash type to sign with if not specified by the PSBT. Must be one of
                                 "DEFAULT"
                                 "ALL"
                                 "NONE"
                                 "SINGLE"
                                 "ALL|ANYONECANPAY"
                                 "NONE|ANYONECANPAY"
                                 "SINGLE|ANYONECANPAY"
4. bip32derivs                   (boolean, optional, default=true) Include BIP 32 derivation paths for public keys if we know them
5. finalize                      (boolean, optional, default=true) Also finalize inputs if possible

Result:
{                             (json object)
  "psbt" : "str",             (string) The base64-encoded partially signed transaction
  "complete" : true|false,    (boolean) If the transaction has a complete set of signatures
  "hex" : "hex"               (string, optional) The hex-encoded network transaction if complete
}

Examples:
> bitcoin-cli descriptorprocesspsbt "psbt" "[\"descriptor1\", \"descriptor2\"]"
> bitcoin-cli descriptorprocesspsbt "psbt" "[{\"desc\":\"mydescriptor\", \"range\":21}]"


