---
name: combinepsbt
btcversion: 0.17.0
btcgroup: rawtransactions
permalink: en/doc/0.17.0/rpc/rawtransactions/combinepsbt/
---

combinepsbt ["psbt",...]

Combine multiple partially signed Bitcoin transactions into one transaction.
Implements the Combiner role.

Arguments:
1. "txs"                   (string) A json array of base64 strings of partially signed transactions
    [
      "psbt"             (string) A base64 string of a PSBT
      ,...
    ]

Result:
  "psbt"          (string) The base64-encoded partially signed transaction

Examples:
> bitcoin-cli combinepsbt ["mybase64_1", "mybase64_2", "mybase64_3"]


