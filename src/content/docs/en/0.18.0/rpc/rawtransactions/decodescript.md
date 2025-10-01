---
name: decodescript
btcversion: 0.18.0
btcgroup: rawtransactions
permalink: en/doc/0.18.0/rpc/rawtransactions/decodescript/
---

decodescript "hexstring"

Decode a hex-encoded script.

Arguments:
1. hexstring    (string, required) the hex-encoded script

Result:
{
  "asm":"asm",   (string) Script public key
  "hex":"hex",   (string) hex-encoded public key
  "type":"type", (string) The output type
  "reqSigs": n,    (numeric) The required signatures
  "addresses": [   (json array of string)
     "address"     (string) bitcoin address
     ,...
  ],
  "p2sh","address" (string) address of P2SH script wrapping this redeem script (not returned if the script is already a P2SH).
}

Examples:
> bitcoin-cli decodescript "hexstring"
> curl --user myusername --data-binary '{"jsonrpc": "1.0", "id":"curltest", "method": "decodescript", "params": ["hexstring"] }' -H 'content-type: text/plain;' http://127.0.0.1:8332/


