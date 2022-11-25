---
name: getrawtransaction
btcversion: 24.0.0
btcgroup: rawtransactions
permalink: en/doc/24.0.0/rpc/rawtransactions/getrawtransaction/
---

getrawtransaction "txid" ( verbose "blockhash" )

Return the raw transaction data.

By default, this call only returns a transaction if it is in the mempool. If -txindex is enabled
and no blockhash argument is passed, it will return the transaction if it is in the mempool or any block.
If a blockhash argument is passed, it will return the transaction if
the specified block is available and the transaction is in that block.

Hint: Use gettransaction for wallet transactions.

If verbose is 'true', returns an Object with information about 'txid'.
If verbose is 'false' or omitted, returns a string that is serialized, hex-encoded data for 'txid'.

Arguments:
1. txid         (string, required) The transaction id
2. verbose      (boolean, optional, default=false) If false, return a string, otherwise return a json object
3. blockhash    (string, optional) The block in which to look for the transaction

Result (if verbose is not set or set to false):
"str"    (string) The serialized, hex-encoded data for 'txid'

Result (if verbose is set to true):
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
        "type" : "str",              (string) The type, eg 'pubkeyhash'
        "address" : "str"            (string, optional) The Bitcoin address (only if a well-defined address exists)
      }
    },
    ...
  ]
}

Examples:
> bitcoin-cli getrawtransaction "mytxid"
> bitcoin-cli getrawtransaction "mytxid" true
> curl --user myusername --data-binary '{"jsonrpc": "1.0", "id": "curltest", "method": "getrawtransaction", "params": ["mytxid", true]}' -H 'content-type: text/plain;' http://127.0.0.1:8332/
> bitcoin-cli getrawtransaction "mytxid" false "myblockhash"
> bitcoin-cli getrawtransaction "mytxid" true "myblockhash"


