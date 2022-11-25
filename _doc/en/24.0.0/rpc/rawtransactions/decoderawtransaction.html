---
name: decoderawtransaction
btcversion: 24.0.0
btcgroup: rawtransactions
permalink: en/doc/24.0.0/rpc/rawtransactions/decoderawtransaction/
---

decoderawtransaction "hexstring" ( iswitness )

Return a JSON object representing the serialized, hex-encoded transaction.

Arguments:
1. hexstring    (string, required) The transaction hex string
2. iswitness    (boolean, optional, default=depends on heuristic tests) Whether the transaction hex is a serialized witness transaction.
                If iswitness is not present, heuristic tests will be used in decoding.
                If true, only witness deserialization will be tried.
                If false, only non-witness deserialization will be tried.
                This boolean should reflect whether the transaction has inputs
                (e.g. fully valid, or on-chain transactions), if known by the caller.

Result:
{                             (json object)
  "txid" : "hex",             (string) The transaction id
  "hash" : "hex",             (string) The transaction hash (differs from txid for witness transactions)
  "size" : n,                 (numeric) The serialized transaction size
  "vsize" : n,                (numeric) The virtual transaction size (differs from size for witness transactions)
  "weight" : n,               (numeric) The transaction's weight (between vsize*4-3 and vsize*4)
  "version" : n,              (numeric) The version
  "locktime" : xxx,           (numeric) The lock time
  "vin" : [                   (json array)
    {                         (json object)
      "coinbase" : "hex",     (string, optional) The coinbase value (only if coinbase transaction)
      "txid" : "hex",         (string, optional) The transaction id (if not coinbase transaction)
      "vout" : n,             (numeric, optional) The output number (if not coinbase transaction)
      "scriptSig" : {         (json object, optional) The script (if not coinbase transaction)
        "asm" : "str",        (string) Disassembly of the signature script
        "hex" : "hex"         (string) The raw signature script bytes, hex-encoded
      },
      "txinwitness" : [       (json array, optional)
        "hex",                (string) hex-encoded witness data (if any)
        ...
      ],
      "sequence" : n          (numeric) The script sequence number
    },
    ...
  ],
  "vout" : [                  (json array)
    {                         (json object)
      "value" : n,            (numeric) The value in BTC
      "n" : n,                (numeric) index
      "scriptPubKey" : {      (json object)
        "asm" : "str",        (string) Disassembly of the public key script
        "desc" : "str",       (string) Inferred descriptor for the output
        "hex" : "hex",        (string) The raw public key script bytes, hex-encoded
        "type" : "str",       (string) The type, eg 'pubkeyhash'
        "address" : "str"     (string, optional) The Bitcoin address (only if a well-defined address exists)
      }
    },
    ...
  ]
}

Examples:
> bitcoin-cli decoderawtransaction "hexstring"
> curl --user myusername --data-binary '{"jsonrpc": "1.0", "id": "curltest", "method": "decoderawtransaction", "params": ["hexstring"]}' -H 'content-type: text/plain;' http://127.0.0.1:8332/


