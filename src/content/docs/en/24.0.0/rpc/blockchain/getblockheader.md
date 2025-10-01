---
name: getblockheader
btcversion: 24.0.0
btcgroup: blockchain
permalink: en/doc/24.0.0/rpc/blockchain/getblockheader/
---

getblockheader "blockhash" ( verbose )

If verbose is false, returns a string that is serialized, hex-encoded data for blockheader 'hash'.
If verbose is true, returns an Object with information about blockheader <hash>.

Arguments:
1. blockhash    (string, required) The block hash
2. verbose      (boolean, optional, default=true) true for a json object, false for the hex-encoded data

Result (for verbose = true):
{                                 (json object)
  "hash" : "hex",                 (string) the block hash (same as provided)
  "confirmations" : n,            (numeric) The number of confirmations, or -1 if the block is not on the main chain
  "height" : n,                   (numeric) The block height or index
  "version" : n,                  (numeric) The block version
  "versionHex" : "hex",           (string) The block version formatted in hexadecimal
  "merkleroot" : "hex",           (string) The merkle root
  "time" : xxx,                   (numeric) The block time expressed in UNIX epoch time
  "mediantime" : xxx,             (numeric) The median block time expressed in UNIX epoch time
  "nonce" : n,                    (numeric) The nonce
  "bits" : "hex",                 (string) The bits
  "difficulty" : n,               (numeric) The difficulty
  "chainwork" : "hex",            (string) Expected number of hashes required to produce the current chain
  "nTx" : n,                      (numeric) The number of transactions in the block
  "previousblockhash" : "hex",    (string, optional) The hash of the previous block (if available)
  "nextblockhash" : "hex"         (string, optional) The hash of the next block (if available)
}

Result (for verbose=false):
"hex"    (string) A string that is serialized, hex-encoded data for block 'hash'

Examples:
> bitcoin-cli getblockheader "00000000c937983704a73af28acdec37b049d214adbda81d7e2a3dd146f6ed09"
> curl --user myusername --data-binary '{"jsonrpc": "1.0", "id": "curltest", "method": "getblockheader", "params": ["00000000c937983704a73af28acdec37b049d214adbda81d7e2a3dd146f6ed09"]}' -H 'content-type: text/plain;' http://127.0.0.1:8332/


