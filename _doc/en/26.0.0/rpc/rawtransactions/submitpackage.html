---
name: submitpackage
btcversion: 26.0.0
btcgroup: rawtransactions
permalink: en/doc/26.0.0/rpc/rawtransactions/submitpackage/
---

submitpackage ["rawtx",...]

Submit a package of raw transactions (serialized, hex-encoded) to local node.
The package must consist of a child with its parents, and none of the parents may depend on one another.
The package will be validated according to consensus and mempool policy rules. If all transactions pass, they will be accepted to mempool.
This RPC is experimental and the interface may be unstable. Refer to doc/policy/packages.md for documentation on package policies.
Warning: successful submission does not mean the transactions will propagate throughout the network.

Arguments:
1. package         (json array, required) An array of raw transactions.
     [
       "rawtx",    (string)
       ...
     ]

Result:
{                                   (json object)
  "tx-results" : {                  (json object) transaction results keyed by wtxid
    "wtxid" : {                     (json object) transaction wtxid
      "txid" : "hex",               (string) The transaction hash in hex
      "other-wtxid" : "hex",        (string, optional) The wtxid of a different transaction with the same txid but different witness found in the mempool. This means the submitted transaction was ignored.
      "vsize" : n,                  (numeric) Virtual transaction size as defined in BIP 141.
      "fees" : {                    (json object) Transaction fees
        "base" : n,                 (numeric) transaction fee in BTC
        "effective-feerate" : n,    (numeric, optional) if the transaction was not already in the mempool, the effective feerate in BTC per KvB. For example, the package feerate and/or feerate with modified fees from prioritisetransaction.
        "effective-includes" : [    (json array, optional) if effective-feerate is provided, the wtxids of the transactions whose fees and vsizes are included in effective-feerate.
          "hex",                    (string) transaction wtxid in hex
          ...
        ]
      }
    },
    ...
  },
  "replaced-transactions" : [       (json array, optional) List of txids of replaced transactions
    "hex",                          (string) The transaction id
    ...
  ]
}

Examples:
> bitcoin-cli testmempoolaccept [rawtx1, rawtx2]
> bitcoin-cli submitpackage [rawtx1, rawtx2]


