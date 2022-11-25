---
name: getblock
btcversion: 24.0.0
btcgroup: blockchain
permalink: en/doc/24.0.0/rpc/blockchain/getblock/
---

getblock "blockhash" ( verbosity )

If verbosity is 0, returns a string that is serialized, hex-encoded data for block 'hash'.
If verbosity is 1, returns an Object with information about block <hash>.
If verbosity is 2, returns an Object with information about block <hash> and information about each transaction.
If verbosity is 3, returns an Object with information about block <hash> and information about each transaction, including prevout information for inputs (only for unpruned blocks in the current best chain).

Arguments:
1. blockhash    (string, required) The block hash
2. verbosity    (numeric, optional, default=1) 0 for hex-encoded data, 1 for a JSON object, 2 for JSON object with transaction data, and 3 for JSON object with transaction data including prevout information for inputs

Result (for verbosity = 0):
"hex"    (string) A string that is serialized, hex-encoded data for block 'hash'

Result (for verbosity = 1):
{                                 (json object)
  "hash" : "hex",                 (string) the block hash (same as provided)
  "confirmations" : n,            (numeric) The number of confirmations, or -1 if the block is not on the main chain
  "size" : n,                     (numeric) The block size
  "strippedsize" : n,             (numeric) The block size excluding witness data
  "weight" : n,                   (numeric) The block weight as defined in BIP 141
  "height" : n,                   (numeric) The block height or index
  "version" : n,                  (numeric) The block version
  "versionHex" : "hex",           (string) The block version formatted in hexadecimal
  "merkleroot" : "hex",           (string) The merkle root
  "tx" : [                        (json array) The transaction ids
    "hex",                        (string) The transaction id
    ...
  ],
  "time" : xxx,                   (numeric) The block time expressed in UNIX epoch time
  "mediantime" : xxx,             (numeric) The median block time expressed in UNIX epoch time
  "nonce" : n,                    (numeric) The nonce
  "bits" : "hex",                 (string) The bits
  "difficulty" : n,               (numeric) The difficulty
  "chainwork" : "hex",            (string) Expected number of hashes required to produce the chain up to this block (in hex)
  "nTx" : n,                      (numeric) The number of transactions in the block
  "previousblockhash" : "hex",    (string, optional) The hash of the previous block (if available)
  "nextblockhash" : "hex"         (string, optional) The hash of the next block (if available)
}

Result (for verbosity = 2):
{                   (json object)
  ...,              Same output as verbosity = 1
  "tx" : [          (json array)
    {               (json object)
      ...,          The transactions in the format of the getrawtransaction RPC. Different from verbosity = 1 "tx" result
      "fee" : n     (numeric) The transaction fee in BTC, omitted if block undo data is not available
    },
    ...
  ]
}

Result (for verbosity = 3):
{                                        (json object)
  ...,                                   Same output as verbosity = 2
  "tx" : [                               (json array)
    {                                    (json object)
      "vin" : [                          (json array)
        {                                (json object)
          ...,                           The same output as verbosity = 2
          "prevout" : {                  (json object) (Only if undo information is available)
            "generated" : true|false,    (boolean) Coinbase or not
            "height" : n,                (numeric) The height of the prevout
            "value" : n,                 (numeric) The value in BTC
            "scriptPubKey" : {           (json object)
              "asm" : "str",             (string) Disassembly of the public key script
              "desc" : "str",            (string) Inferred descriptor for the output
              "hex" : "hex",             (string) The raw public key script bytes, hex-encoded
              "address" : "str",         (string, optional) The Bitcoin address (only if a well-defined address exists)
              "type" : "str"             (string) The type (one of: nonstandard, pubkey, pubkeyhash, scripthash, multisig, nulldata, witness_v0_scripthash, witness_v0_keyhash, witness_v1_taproot, witness_unknown)
            }
          }
        },
        ...
      ]
    },
    ...
  ]
}

Examples:
> bitcoin-cli getblock "00000000c937983704a73af28acdec37b049d214adbda81d7e2a3dd146f6ed09"
> curl --user myusername --data-binary '{"jsonrpc": "1.0", "id": "curltest", "method": "getblock", "params": ["00000000c937983704a73af28acdec37b049d214adbda81d7e2a3dd146f6ed09"]}' -H 'content-type: text/plain;' http://127.0.0.1:8332/


