---
name: signrawtransactionwithwallet
btcversion: 24.0.0
btcgroup: wallet
permalink: en/doc/24.0.0/rpc/wallet/signrawtransactionwithwallet/
---

signrawtransactionwithwallet "hexstring" ( [{"txid":"hex","vout":n,"scriptPubKey":"hex","redeemScript":"hex","witnessScript":"hex","amount":amount},...] "sighashtype" )

Sign inputs for raw transaction (serialized, hex-encoded).
The second optional argument (may be null) is an array of previous transaction outputs that
this transaction depends on but may not yet be in the block chain.
Requires wallet passphrase to be set with walletpassphrase call if wallet is encrypted.

Arguments:
1. hexstring                        (string, required) The transaction hex string
2. prevtxs                          (json array, optional) The previous dependent transaction outputs
     [
       {                            (json object)
         "txid": "hex",             (string, required) The transaction id
         "vout": n,                 (numeric, required) The output number
         "scriptPubKey": "hex",     (string, required) script key
         "redeemScript": "hex",     (string) (required for P2SH) redeem script
         "witnessScript": "hex",    (string) (required for P2WSH or P2SH-P2WSH) witness script
         "amount": amount,          (numeric or string) (required for Segwit inputs) the amount spent
       },
       ...
     ]
3. sighashtype                      (string, optional, default="DEFAULT for Taproot, ALL otherwise") The signature hash type. Must be one of
                                    "DEFAULT"
                                    "ALL"
                                    "NONE"
                                    "SINGLE"
                                    "ALL|ANYONECANPAY"
                                    "NONE|ANYONECANPAY"
                                    "SINGLE|ANYONECANPAY"

Result:
{                             (json object)
  "hex" : "hex",              (string) The hex-encoded raw transaction with signature(s)
  "complete" : true|false,    (boolean) If the transaction has a complete set of signatures
  "errors" : [                (json array, optional) Script verification errors (if there are any)
    {                         (json object)
      "txid" : "hex",         (string) The hash of the referenced, previous transaction
      "vout" : n,             (numeric) The index of the output to spent and used as input
      "witness" : [           (json array)
        "hex",                (string)
        ...
      ],
      "scriptSig" : "hex",    (string) The hex-encoded signature script
      "sequence" : n,         (numeric) Script sequence number
      "error" : "str"         (string) Verification or signing error related to the input
    },
    ...
  ]
}

Examples:
> bitcoin-cli signrawtransactionwithwallet "myhex"
> curl --user myusername --data-binary '{"jsonrpc": "1.0", "id": "curltest", "method": "signrawtransactionwithwallet", "params": ["myhex"]}' -H 'content-type: text/plain;' http://127.0.0.1:8332/


