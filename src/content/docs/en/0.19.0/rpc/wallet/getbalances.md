---
name: getbalances
btcversion: 0.19.0
btcgroup: wallet
permalink: en/doc/0.19.0/rpc/wallet/getbalances/
---

getbalances
Returns an object with all balances in BTC.

Result:
{
    "mine": {                        (object) balances from outputs that the wallet can sign
      "trusted": xxx                 (numeric) trusted balance (outputs created by the wallet or confirmed outputs)
      "untrusted_pending": xxx       (numeric) untrusted pending balance (outputs created by others that are in the mempool)
      "immature": xxx                (numeric) balance from immature coinbase outputs
      "used": xxx                    (numeric) (only present if avoid_reuse is set) balance from coins sent to addresses that were previously spent from (potentially privacy violating)
    },
    "watchonly": {                   (object) watchonly balances (not present if wallet does not watch anything)
      "trusted": xxx                 (numeric) trusted balance (outputs created by the wallet or confirmed outputs)
      "untrusted_pending": xxx       (numeric) untrusted pending balance (outputs created by others that are in the mempool)
      "immature": xxx                (numeric) balance from immature coinbase outputs
    },
}

Examples:
> bitcoin-cli getbalances 
> curl --user myusername --data-binary '{"jsonrpc": "1.0", "id":"curltest", "method": "getbalances", "params": [] }' -H 'content-type: text/plain;' http://127.0.0.1:8332/


