---
name: decodepsbt
btcversion: 0.18.0
btcgroup: rawtransactions
permalink: en/doc/0.18.0/rpc/rawtransactions/decodepsbt/
---

decodepsbt "psbt"

Return a JSON object representing the serialized, base64-encoded partially signed Bitcoin transaction.

Arguments:
1. psbt    (string, required) The PSBT base64 string

Result:
{
  "tx" : {                   (json object) The decoded network-serialized unsigned transaction.
    ...                                      The layout is the same as the output of decoderawtransaction.
  },
  "unknown" : {                (json object) The unknown global fields
    "key" : "value"            (key-value pair) An unknown key-value pair
     ...
  },
  "inputs" : [                 (array of json objects)
    {
      "non_witness_utxo" : {   (json object, optional) Decoded network transaction for non-witness UTXOs
        ...
      },
      "witness_utxo" : {            (json object, optional) Transaction output for witness UTXOs
        "amount" : x.xxx,           (numeric) The value in BTC
        "scriptPubKey" : {          (json object)
          "asm" : "asm",            (string) The asm
          "hex" : "hex",            (string) The hex
          "type" : "pubkeyhash",    (string) The type, eg 'pubkeyhash'
          "address" : "address"     (string) Bitcoin address if there is one
        }
      },
      "partial_signatures" : {             (json object, optional)
        "pubkey" : "signature",           (string) The public key and signature that corresponds to it.
        ,...
      }
      "sighash" : "type",                  (string, optional) The sighash type to be used
      "redeem_script" : {       (json object, optional)
          "asm" : "asm",            (string) The asm
          "hex" : "hex",            (string) The hex
          "type" : "pubkeyhash",    (string) The type, eg 'pubkeyhash'
        }
      "witness_script" : {       (json object, optional)
          "asm" : "asm",            (string) The asm
          "hex" : "hex",            (string) The hex
          "type" : "pubkeyhash",    (string) The type, eg 'pubkeyhash'
        }
      "bip32_derivs" : {          (json object, optional)
        "pubkey" : {                     (json object, optional) The public key with the derivation path as the value.
          "master_fingerprint" : "fingerprint"     (string) The fingerprint of the master key
          "path" : "path",                         (string) The path
        }
        ,...
      }
      "final_scriptsig" : {       (json object, optional)
          "asm" : "asm",            (string) The asm
          "hex" : "hex",            (string) The hex
        }
       "final_scriptwitness": ["hex", ...] (array of string) hex-encoded witness data (if any)
      "unknown" : {                (json object) The unknown global fields
        "key" : "value"            (key-value pair) An unknown key-value pair
         ...
      },
    }
    ,...
  ]
  "outputs" : [                 (array of json objects)
    {
      "redeem_script" : {       (json object, optional)
          "asm" : "asm",            (string) The asm
          "hex" : "hex",            (string) The hex
          "type" : "pubkeyhash",    (string) The type, eg 'pubkeyhash'
        }
      "witness_script" : {       (json object, optional)
          "asm" : "asm",            (string) The asm
          "hex" : "hex",            (string) The hex
          "type" : "pubkeyhash",    (string) The type, eg 'pubkeyhash'
      }
      "bip32_derivs" : [          (array of json objects, optional)
        {
          "pubkey" : "pubkey",                     (string) The public key this path corresponds to
          "master_fingerprint" : "fingerprint"     (string) The fingerprint of the master key
          "path" : "path",                         (string) The path
          }
        }
        ,...
      ],
      "unknown" : {                (json object) The unknown global fields
        "key" : "value"            (key-value pair) An unknown key-value pair
         ...
      },
    }
    ,...
  ]
  "fee" : fee                      (numeric, optional) The transaction fee paid if all UTXOs slots in the PSBT have been filled.
}

Examples:
> bitcoin-cli decodepsbt "psbt"


