---
name: prioritisetransaction
btcversion: 24.0.0
btcgroup: mining
permalink: en/doc/24.0.0/rpc/mining/prioritisetransaction/
---

prioritisetransaction "txid" ( dummy ) fee_delta

Accepts the transaction into mined blocks at a higher (or lower) priority

Arguments:
1. txid         (string, required) The transaction id.
2. dummy        (numeric, optional) API-Compatibility for previous API. Must be zero or null.
                DEPRECATED. For forward compatibility use named arguments and omit this parameter.
3. fee_delta    (numeric, required) The fee value (in satoshis) to add (or subtract, if negative).
                Note, that this value is not a fee rate. It is a value to modify absolute fee of the TX.
                The fee is not actually paid, only the algorithm for selecting transactions into a block
                considers the transaction as it would have paid a higher (or lower) fee.

Result:
true|false    (boolean) Returns true

Examples:
> bitcoin-cli prioritisetransaction "txid" 0.0 10000
> curl --user myusername --data-binary '{"jsonrpc": "1.0", "id": "curltest", "method": "prioritisetransaction", "params": ["txid", 0.0, 10000]}' -H 'content-type: text/plain;' http://127.0.0.1:8332/


