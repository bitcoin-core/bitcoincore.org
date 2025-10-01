---
name: analyzepsbt
btcversion: 0.18.0
btcgroup: rawtransactions
permalink: en/doc/0.18.0/rpc/rawtransactions/analyzepsbt/
---

analyzepsbt "psbt"

Analyzes and provides information about the current status of a PSBT and its inputs

Arguments:
1. psbt    (string, required) A base64 string of a PSBT

Result:
{
  "inputs" : [                      (array of json objects)
    {
      "has_utxo" : true|false     (boolean) Whether a UTXO is provided
      "is_final" : true|false     (boolean) Whether the input is finalized
      "missing" : {               (json object, optional) Things that are missing that are required to complete this input
        "pubkeys" : [             (array, optional)
          "keyid"                 (string) Public key ID, hash160 of the public key, of a public key whose BIP 32 derivation path is missing
        ]
        "signatures" : [          (array, optional)
          "keyid"                 (string) Public key ID, hash160 of the public key, of a public key whose signature is missing
        ]
        "redeemscript" : "hash"   (string, optional) Hash160 of the redeemScript that is missing
        "witnessscript" : "hash"  (string, optional) SHA256 of the witnessScript that is missing
      }
      "next" : "role"             (string, optional) Role of the next person that this input needs to go to
    }
    ,...
  ]
  "estimated_vsize" : vsize       (numeric, optional) Estimated vsize of the final signed transaction
  "estimated_feerate" : feerate   (numeric, optional) Estimated feerate of the final signed transaction in BTC/kB. Shown only if all UTXO slots in the PSBT have been filled.
  "fee" : fee                     (numeric, optional) The transaction fee paid. Shown only if all UTXO slots in the PSBT have been filled.
  "next" : "role"                 (string) Role of the next person that this psbt needs to go to
}

Examples:
> bitcoin-cli analyzepsbt "psbt"


