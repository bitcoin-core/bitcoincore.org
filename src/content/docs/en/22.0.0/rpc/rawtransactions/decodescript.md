---
name: decodescript
btcversion: 22.0.0
btcgroup: rawtransactions
permalink: en/doc/22.0.0/rpc/rawtransactions/decodescript/
---

decodescript "hexstring"

Decode a hex-encoded script.

Arguments:
1. hexstring    (string, required) the hex-encoded script

Result:
{                             (json object)
  "asm" : "str",              (string) Script public key
  "type" : "str",             (string) The output type (e.g. nonstandard, pubkey, pubkeyhash, scripthash, multisig, nulldata, witness_v0_scripthash, witness_v0_keyhash, witness_v1_taproot, witness_unknown)
  "address" : "str",          (string, optional) bitcoin address (only if a well-defined address exists)
  "reqSigs" : n,              (numeric, optional) (DEPRECATED, returned only if config option -deprecatedrpc=addresses is passed) Number of required signatures
  "addresses" : [             (json array, optional) (DEPRECATED, returned only if config option -deprecatedrpc=addresses is passed) Array of bitcoin addresses
    "str",                    (string) bitcoin address
    ...
  ],
  "p2sh" : "str",             (string) address of P2SH script wrapping this redeem script (not returned if the script is already a P2SH)
  "segwit" : {                (json object) Result of a witness script public key wrapping this redeem script (not returned if the script is a P2SH or witness)
    "asm" : "str",            (string) String representation of the script public key
    "hex" : "hex",            (string) Hex string of the script public key
    "type" : "str",           (string) The type of the script public key (e.g. witness_v0_keyhash or witness_v0_scripthash)
    "address" : "str",        (string, optional) bitcoin address (only if a well-defined address exists)
    "reqSigs" : n,            (numeric, optional) (DEPRECATED, returned only if config option -deprecatedrpc=addresses is passed) Number of required signatures
    "addresses" : [           (json array, optional) (DEPRECATED, returned only if config option -deprecatedrpc=addresses is passed) Array of bitcoin addresses
      "str",                  (string) segwit address
      ...
    ],
    "p2sh-segwit" : "str"     (string) address of the P2SH script wrapping this witness redeem script
  }
}

Examples:
> bitcoin-cli decodescript "hexstring"
> curl --user myusername --data-binary '{"jsonrpc": "1.0", "id": "curltest", "method": "decodescript", "params": ["hexstring"]}' -H 'content-type: text/plain;' http://127.0.0.1:8332/


