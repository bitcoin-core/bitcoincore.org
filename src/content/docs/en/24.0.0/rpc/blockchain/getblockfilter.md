---
name: getblockfilter
btcversion: 24.0.0
btcgroup: blockchain
permalink: en/doc/24.0.0/rpc/blockchain/getblockfilter/
---

getblockfilter "blockhash" ( "filtertype" )

Retrieve a BIP 157 content filter for a particular block.

Arguments:
1. blockhash     (string, required) The hash of the block
2. filtertype    (string, optional, default="basic") The type name of the filter

Result:
{                      (json object)
  "filter" : "hex",    (string) the hex-encoded filter data
  "header" : "hex"     (string) the hex-encoded filter header
}

Examples:
> bitcoin-cli getblockfilter "00000000c937983704a73af28acdec37b049d214adbda81d7e2a3dd146f6ed09" "basic"
> curl --user myusername --data-binary '{"jsonrpc": "1.0", "id": "curltest", "method": "getblockfilter", "params": ["00000000c937983704a73af28acdec37b049d214adbda81d7e2a3dd146f6ed09", "basic"]}' -H 'content-type: text/plain;' http://127.0.0.1:8332/


