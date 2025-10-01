---
name: decodepsbt
btcversion: 23.0.0
btcgroup: rawtransactions
permalink: en/doc/23.0.0/rpc/rawtransactions/decodepsbt/
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
  "global_xpubs" : [                       (json array)
    {                                      (json object)
      "xpub" : "str",                      (string) The extended public key this path corresponds to
      "master_fingerprint" : "hex",        (string) The fingerprint of the master key
      "path" : "str"                       (string) The path
    },
    ...
  ],
  "psbt_version" : n,                      (numeric) The PSBT version number. Not to be confused with the unsigned transaction version
  "proprietary" : [                        (json array) The global proprietary map
    {                                      (json object)
      "identifier" : "hex",                (string) The hex string for the proprietary identifier
      "subtype" : n,                       (numeric) The number for the subtype
      "key" : "hex",                       (string) The hex for the key
      "value" : "hex"                      (string) The hex for the value
    },
    ...
  ],
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
          "desc" : "str",                  (string) Inferred descriptor for the output
          "hex" : "hex",                   (string) The hex
          "type" : "str",                  (string) The type, eg 'pubkeyhash'
          "address" : "str"                (string, optional) The Bitcoin address (only if a well-defined address exists)
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
        {                                  (json object)
          "pubkey" : "str",                (string) The public key with the derivation path as the value.
          "master_fingerprint" : "str",    (string) The fingerprint of the master key
          "path" : "str"                   (string) The path
        },
        ...
      ],
      "final_scriptSig" : {                (json object, optional)
        "asm" : "str",                     (string) The asm
        "hex" : "str"                      (string) The hex
      },
      "final_scriptwitness" : [            (json array, optional)
        "hex",                             (string) hex-encoded witness data (if any)
        ...
      ],
      "ripemd160_preimages" : {            (json object, optional)
        "hash" : "str",                    (string) The hash and preimage that corresponds to it.
        ...
      },
      "sha256_preimages" : {               (json object, optional)
        "hash" : "str",                    (string) The hash and preimage that corresponds to it.
        ...
      },
      "hash160_preimages" : {              (json object, optional)
        "hash" : "str",                    (string) The hash and preimage that corresponds to it.
        ...
      },
      "hash256_preimages" : {              (json object, optional)
        "hash" : "str",                    (string) The hash and preimage that corresponds to it.
        ...
      },
      "unknown" : {                        (json object, optional) The unknown input fields
        "key" : "hex",                     (string) (key-value pair) An unknown key-value pair
        ...
      },
      "proprietary" : [                    (json array, optional) The input proprietary map
        {                                  (json object)
          "identifier" : "hex",            (string) The hex string for the proprietary identifier
          "subtype" : n,                   (numeric) The number for the subtype
          "key" : "hex",                   (string) The hex for the key
          "value" : "hex"                  (string) The hex for the value
        },
        ...
      ]
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
      "unknown" : {                        (json object, optional) The unknown global fields
        "key" : "hex",                     (string) (key-value pair) An unknown key-value pair
        ...
      },
      "proprietary" : [                    (json array, optional) The output proprietary map
        {                                  (json object)
          "identifier" : "hex",            (string) The hex string for the proprietary identifier
          "subtype" : n,                   (numeric) The number for the subtype
          "key" : "hex",                   (string) The hex for the key
          "value" : "hex"                  (string) The hex for the value
        },
        ...
      ]
    },
    ...
  ],
  "fee" : n                                (numeric, optional) The transaction fee paid if all UTXOs slots in the PSBT have been filled.
}

Examples:
> bitcoin-cli decodepsbt "psbt"


