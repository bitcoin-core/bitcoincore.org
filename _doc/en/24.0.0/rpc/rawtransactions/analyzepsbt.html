---
name: analyzepsbt
btcversion: 24.0.0
btcgroup: rawtransactions
permalink: en/doc/24.0.0/rpc/rawtransactions/analyzepsbt/
---

analyzepsbt "psbt"

Analyzes and provides information about the current status of a PSBT and its inputs

Arguments:
1. psbt    (string, required) A base64 string of a PSBT

Result:
{                                   (json object)
  "inputs" : [                      (json array, optional)
    {                               (json object)
      "has_utxo" : true|false,      (boolean) Whether a UTXO is provided
      "is_final" : true|false,      (boolean) Whether the input is finalized
      "missing" : {                 (json object, optional) Things that are missing that are required to complete this input
        "pubkeys" : [               (json array, optional)
          "hex",                    (string) Public key ID, hash160 of the public key, of a public key whose BIP 32 derivation path is missing
          ...
        ],
        "signatures" : [            (json array, optional)
          "hex",                    (string) Public key ID, hash160 of the public key, of a public key whose signature is missing
          ...
        ],
        "redeemscript" : "hex",     (string, optional) Hash160 of the redeemScript that is missing
        "witnessscript" : "hex"     (string, optional) SHA256 of the witnessScript that is missing
      },
      "next" : "str"                (string, optional) Role of the next person that this input needs to go to
    },
    ...
  ],
  "estimated_vsize" : n,            (numeric, optional) Estimated vsize of the final signed transaction
  "estimated_feerate" : n,          (numeric, optional) Estimated feerate of the final signed transaction in BTC/kvB. Shown only if all UTXO slots in the PSBT have been filled
  "fee" : n,                        (numeric, optional) The transaction fee paid. Shown only if all UTXO slots in the PSBT have been filled
  "next" : "str",                   (string) Role of the next person that this psbt needs to go to
  "error" : "str"                   (string, optional) Error message (if there is one)
}

Examples:
> bitcoin-cli analyzepsbt "psbt"


