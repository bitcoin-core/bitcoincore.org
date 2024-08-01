---
name: getrawtransaction
btcversion: 27.0.0
btcgroup: rawtransactions
permalink: en/doc/27.0.0/rpc/rawtransactions/getrawtransaction/
---

getrawtransaction "txid" ( verbosity "blockhash" )

By default, this call only returns a transaction if it is in the mempool. If -txindex is enabled
and no blockhash argument is passed, it will return the transaction if it is in the mempool or any block.
If a blockhash argument is passed, it will return the transaction if
the specified block is available and the transaction is in that block.

Hint: Use gettransaction for wallet transactions.

If verbosity is 0 or omitted, returns the serialized transaction as a hex-encoded string.
If verbosity is 1, returns a JSON Object with information about the transaction.
If verbosity is 2, returns a JSON Object with information about the transaction, including fee and prevout information.

Arguments:
1. txid         (string, required) The transaction id
2. verbosity    (numeric, optional, default=0) 0 for hex-encoded data, 1 for a JSON object, and 2 for JSON object with fee and prevout
3. blockhash    (string, optional) The block in which to look for the transaction

Result (if verbosity is not set or set to 0):
"str"    (string) The serialized transaction as a hex-encoded string for 'txid'

Result (if verbosity is set to 1):
{                                    (json object)
  "in_active_chain" : true|false,    (boolean, optional) Whether specified block is in the active chain or not (only present with explicit "blockhash" argument)
  "blockhash" : "hex",               (string, optional) the block hash
  "confirmations" : n,               (numeric, optional) The confirmations
  "blocktime" : xxx,                 (numeric, optional) The block time expressed in UNIX epoch time
  "time" : n,                        (numeric, optional) Same as "blocktime"
  "hex" : "hex",                     (string) The serialized, hex-encoded data for 'txid'
  "txid" : "hex",                    (string) The transaction id (same as provided)
  "hash" : "hex",                    (string) The transaction hash (differs from txid for witness transactions)
  "size" : n,                        (numeric) The serialized transaction size
  "vsize" : n,                       (numeric) The virtual transaction size (differs from size for witness transactions)
  "weight" : n,                      (numeric) The transaction's weight (between vsize*4-3 and vsize*4)
  "version" : n,                     (numeric) The version
  "locktime" : xxx,                  (numeric) The lock time
  "vin" : [                          (json array)
    {                                (json object)
      "coinbase" : "hex",            (string, optional) The coinbase value (only if coinbase transaction)
      "txid" : "hex",                (string, optional) The transaction id (if not coinbase transaction)
      "vout" : n,                    (numeric, optional) The output number (if not coinbase transaction)
      "scriptSig" : {                (json object, optional) The script (if not coinbase transaction)
        "asm" : "str",               (string) Disassembly of the signature script
        "hex" : "hex"                (string) The raw signature script bytes, hex-encoded
      },
      "txinwitness" : [              (json array, optional)
        "hex",                       (string) hex-encoded witness data (if any)
        ...
      ],
      "sequence" : n                 (numeric) The script sequence number
    },
    ...
  ],
  "vout" : [                         (json array)
    {                                (json object)
      "value" : n,                   (numeric) The value in BTC
      "n" : n,                       (numeric) index
      "scriptPubKey" : {             (json object)
        "asm" : "str",               (string) Disassembly of the public key script
        "desc" : "str",              (string) Inferred descriptor for the output
        "hex" : "hex",               (string) The raw public key script bytes, hex-encoded
        "address" : "str",           (string, optional) The Bitcoin address (only if a well-defined address exists)
        "type" : "str"               (string) The type (one of: nonstandard, pubkey, pubkeyhash, scripthash, multisig, nulldata, witness_v0_scripthash, witness_v0_keyhash, witness_v1_taproot, witness_unknown)
      }
    },
    ...
  ]
}

Result (for verbosity = 2):
{                                    (json object)
  ...,                               Same output as verbosity = 1
  "fee" : n,                         (numeric, optional) transaction fee in BTC, omitted if block undo data is not available
  "vin" : [                          (json array)
    {                                (json object) utxo being spent
      ...,                           Same output as verbosity = 1
      "prevout" : {                  (json object, optional) The previous output, omitted if block undo data is not available
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
}

Examples:
> bitcoin-cli getrawtransaction "mytxid"
> bitcoin-cli getrawtransaction "mytxid" 1
> curl --user myusername --data-binary '{"jsonrpc": "1.0", "id": "curltest", "method": "getrawtransaction", "params": ["mytxid", 1]}' -H 'content-type: text/plain;' http://127.0.0.1:8332/
> bitcoin-cli getrawtransaction "mytxid" 0 "myblockhash"
> bitcoin-cli getrawtransaction "mytxid" 1 "myblockhash"
> bitcoin-cli getrawtransaction "mytxid" 2 "myblockhash"


