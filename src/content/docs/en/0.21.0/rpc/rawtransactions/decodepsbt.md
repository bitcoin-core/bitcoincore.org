---
name: decodepsbt
btcversion: 0.21.0
btcgroup: rawtransactions
permalink: en/doc/0.21.0/rpc/rawtransactions/decodepsbt/
---

decodepsbt "psbt"

Return a JSON object representing the serialized, base64-encoded partially signed Bitcoin transaction.

Arguments:
1. psbt    (string, required) The PSBT base64 string

Result:
{                                          (json object)
  "tx" : {                                 (json object) The decoded network-serialized unsigned transaction.
    ...                                    The layout is the same as the output of decoderawtransaction.
  },
  "unknown" : {                            (json object) The unknown global fields
    "key" : "hex",                         (string) (key-value pair) An unknown key-value pair
    ...
  },
  "inputs" : [                             (json array)
    {                                      (json object)
      "non_witness_utxo" : {               (json object, optional) Decoded network transaction for non-witness UTXOs
        ...
      },
      "witness_utxo" : {                   (json object, optional) Transaction output for witness UTXOs
        "amount" : n,                      (numeric) The value in BTC
        "scriptPubKey" : {                 (json object)
          "asm" : "str",                   (string) The asm
          "hex" : "hex",                   (string) The hex
          "type" : "str",                  (string) The type, eg 'pubkeyhash'
          "address" : "str"                (string)  Bitcoin address if there is one
        }
      },
      "partial_signatures" : {             (json object, optional)
        "pubkey" : "str",                  (string) The public key and signature that corresponds to it.
        ...
      },
      "sighash" : "str",                   (string, optional) The sighash type to be used
      "redeem_script" : {                  (json object, optional)
        "asm" : "str",                     (string) The asm
        "hex" : "hex",                     (string) The hex
        "type" : "str"                     (string) The type, eg 'pubkeyhash'
      },
      "witness_script" : {                 (json object, optional)
        "asm" : "str",                     (string) The asm
        "hex" : "hex",                     (string) The hex
        "type" : "str"                     (string) The type, eg 'pubkeyhash'
      },
      "bip32_derivs" : [                   (json array, optional)
        {                                  (json object, optional) The public key with the derivation path as the value.
          "master_fingerprint" : "str",    (string) The fingerprint of the master key
          "path" : "str"                   (string) The path
        },
        ...
      ],
      "final_scriptsig" : {                (json object, optional)
        "asm" : "str",                     (string) The asm
        "hex" : "str"                      (string) The hex
      },
      "final_scriptwitness" : [            (json array)
        "hex",                             (string) hex-encoded witness data (if any)
        ...
      ],
      "unknown" : {                        (json object) The unknown global fields
        "key" : "hex",                     (string) (key-value pair) An unknown key-value pair
        ...
      }
    },
    ...
  ],
  "outputs" : [                            (json array)
    {                                      (json object)
      "redeem_script" : {                  (json object, optional)
        "asm" : "str",                     (string) The asm
        "hex" : "hex",                     (string) The hex
        "type" : "str"                     (string) The type, eg 'pubkeyhash'
      },
      "witness_script" : {                 (json object, optional)
        "asm" : "str",                     (string) The asm
        "hex" : "hex",                     (string) The hex
        "type" : "str"                     (string) The type, eg 'pubkeyhash'
      },
      "bip32_derivs" : [                   (json array, optional)
        {                                  (json object)
          "pubkey" : "str",                (string) The public key this path corresponds to
          "master_fingerprint" : "str",    (string) The fingerprint of the master key
          "path" : "str"                   (string) The path
        },
        ...
      ],
      "unknown" : {                        (json object) The unknown global fields
        "key" : "hex",                     (string) (key-value pair) An unknown key-value pair
        ...
      }
    },
    ...
  ],
  "fee" : n                                (numeric, optional) The transaction fee paid if all UTXOs slots in the PSBT have been filled.
}

Examples:
> bitcoin-cli decodepsbt "psbt"


