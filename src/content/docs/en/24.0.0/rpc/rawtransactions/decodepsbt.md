---
name: decodepsbt
btcversion: 24.0.0
btcgroup: rawtransactions
permalink: en/doc/24.0.0/rpc/rawtransactions/decodepsbt/
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
          "asm" : "str",                   (string) Disassembly of the public key script
          "desc" : "str",                  (string) Inferred descriptor for the output
          "hex" : "hex",                   (string) The raw public key script bytes, hex-encoded
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
        "asm" : "str",                     (string) Disassembly of the redeem script
        "hex" : "hex",                     (string) The raw redeem script bytes, hex-encoded
        "type" : "str"                     (string) The type, eg 'pubkeyhash'
      },
      "witness_script" : {                 (json object, optional)
        "asm" : "str",                     (string) Disassembly of the witness script
        "hex" : "hex",                     (string) The raw witness script bytes, hex-encoded
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
        "asm" : "str",                     (string) Disassembly of the final signature script
        "hex" : "hex"                      (string) The raw final signature script bytes, hex-encoded
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
      "taproot_key_path_sig" : "hex",      (string, optional) hex-encoded signature for the Taproot key path spend
      "taproot_script_path_sigs" : [       (json array, optional)
        {                                  (json object, optional) The signature for the pubkey and leaf hash combination
          "pubkey" : "str",                (string) The x-only pubkey for this signature
          "leaf_hash" : "str",             (string) The leaf hash for this signature
          "sig" : "str"                    (string) The signature itself
        },
        ...
      ],
      "taproot_scripts" : [                (json array, optional)
        {                                  (json object)
          "script" : "hex",                (string) A leaf script
          "leaf_ver" : n,                  (numeric) The version number for the leaf script
          "control_blocks" : [             (json array) The control blocks for this script
            "hex",                         (string) A hex-encoded control block for this script
            ...
          ]
        },
        ...
      ],
      "taproot_bip32_derivs" : [           (json array, optional)
        {                                  (json object)
          "pubkey" : "str",                (string) The x-only public key this path corresponds to
          "master_fingerprint" : "str",    (string) The fingerprint of the master key
          "path" : "str",                  (string) The path
          "leaf_hashes" : [                (json array) The hashes of the leaves this pubkey appears in
            "hex",                         (string) The hash of a leaf this pubkey appears in
            ...
          ]
        },
        ...
      ],
      "taproot_internal_key" : "hex",      (string, optional) The hex-encoded Taproot x-only internal key
      "taproot_merkle_root" : "hex",       (string, optional) The hex-encoded Taproot merkle root
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
        "asm" : "str",                     (string) Disassembly of the redeem script
        "hex" : "hex",                     (string) The raw redeem script bytes, hex-encoded
        "type" : "str"                     (string) The type, eg 'pubkeyhash'
      },
      "witness_script" : {                 (json object, optional)
        "asm" : "str",                     (string) Disassembly of the witness script
        "hex" : "hex",                     (string) The raw witness script bytes, hex-encoded
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
      "taproot_internal_key" : "hex",      (string, optional) The hex-encoded Taproot x-only internal key
      "taproot_tree" : [                   (json array, optional) The tuples that make up the Taproot tree, in depth first search order
        {                                  (json object, optional) A single leaf script in the taproot tree
          "depth" : n,                     (numeric) The depth of this element in the tree
          "leaf_ver" : n,                  (numeric) The version of this leaf
          "script" : "str"                 (string) The hex-encoded script itself
        },
        ...
      ],
      "taproot_bip32_derivs" : [           (json array, optional)
        {                                  (json object)
          "pubkey" : "str",                (string) The x-only public key this path corresponds to
          "master_fingerprint" : "str",    (string) The fingerprint of the master key
          "path" : "str",                  (string) The path
          "leaf_hashes" : [                (json array) The hashes of the leaves this pubkey appears in
            "hex",                         (string) The hash of a leaf this pubkey appears in
            ...
          ]
        },
        ...
      ],
      "unknown" : {                        (json object, optional) The unknown output fields
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


