---
name: getdescriptorinfo
btcversion: 24.0.0
btcgroup: util
permalink: en/doc/24.0.0/rpc/util/getdescriptorinfo/
---

getdescriptorinfo "descriptor"

Analyses a descriptor.

Arguments:
1. descriptor    (string, required) The descriptor.

Result:
{                                   (json object)
  "descriptor" : "str",             (string) The descriptor in canonical form, without private keys
  "checksum" : "str",               (string) The checksum for the input descriptor
  "isrange" : true|false,           (boolean) Whether the descriptor is ranged
  "issolvable" : true|false,        (boolean) Whether the descriptor is solvable
  "hasprivatekeys" : true|false     (boolean) Whether the input descriptor contained at least one private key
}

Examples:
Analyse a descriptor
> bitcoin-cli getdescriptorinfo "wpkh([d34db33f/84h/0h/0h]0279be667ef9dcbbac55a06295Ce870b07029Bfcdb2dce28d959f2815b16f81798)"
> curl --user myusername --data-binary '{"jsonrpc": "1.0", "id": "curltest", "method": "getdescriptorinfo", "params": ["wpkh([d34db33f/84h/0h/0h]0279be667ef9dcbbac55a06295Ce870b07029Bfcdb2dce28d959f2815b16f81798)"]}' -H 'content-type: text/plain;' http://127.0.0.1:8332/


