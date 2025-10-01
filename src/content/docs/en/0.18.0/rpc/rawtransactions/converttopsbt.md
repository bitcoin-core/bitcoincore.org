---
name: converttopsbt
btcversion: 0.18.0
btcgroup: rawtransactions
permalink: en/doc/0.18.0/rpc/rawtransactions/converttopsbt/
---

converttopsbt "hexstring" ( permitsigdata iswitness )

Converts a network serialized transaction to a PSBT. This should be used only with createrawtransaction and fundrawtransaction
createpsbt and walletcreatefundedpsbt should be used for new applications.

Arguments:
1. hexstring        (string, required) The hex string of a raw transaction
2. permitsigdata    (boolean, optional, default=false) If true, any signatures in the input will be discarded and conversion.
                    will continue. If false, RPC will fail if any signatures are present.
3. iswitness        (boolean, optional, default=depends on heuristic tests) Whether the transaction hex is a serialized witness transaction.
                    If iswitness is not present, heuristic tests will be used in decoding. If true, only witness deserializaion
                    will be tried. If false, only non-witness deserialization will be tried. Only has an effect if
                    permitsigdata is true.

Result:
  "psbt"        (string)  The resulting raw transaction (base64-encoded string)

Examples:

Create a transaction
> bitcoin-cli createrawtransaction "[{\"txid\":\"myid\",\"vout\":0}]" "[{\"data\":\"00010203\"}]"

Convert the transaction to a PSBT
> bitcoin-cli converttopsbt "rawtransaction"


