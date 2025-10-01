---
name: removeprunedfunds
btcversion: 29.0.0
btcgroup: wallet
permalink: en/doc/29.0.0/rpc/wallet/removeprunedfunds/
---

removeprunedfunds "txid"

Deletes the specified transaction from the wallet. Meant for use with pruned wallets and as a companion to importprunedfunds. This will affect wallet balances.

Arguments:
1. txid    (string, required) The hex-encoded id of the transaction you are deleting

Result:
null    (json null)

Examples:
> bitcoin-cli removeprunedfunds "a8d0c0184dde994a09ec054286f1ce581bebf46446a512166eae7628734ea0a5"

As a JSON-RPC call
> curl --user myusername --data-binary '{"jsonrpc": "2.0", "id": "curltest", "method": "removeprunedfunds", "params": ["a8d0c0184dde994a09ec054286f1ce581bebf46446a512166eae7628734ea0a5"]}' -H 'content-type: application/json' http://127.0.0.1:8332/


