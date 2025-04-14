---
name: submitpackage
btcversion: 29.0.0
btcgroup: rawtransactions
permalink: en/doc/29.0.0/rpc/rawtransactions/submitpackage/
---

submitpackage ["rawtx",...] ( maxfeerate maxburnamount )

Submit a package of raw transactions (serialized, hex-encoded) to local node.
The package will be validated according to consensus and mempool policy rules. If any transaction passes, it will be accepted to mempool.
This RPC is experimental and the interface may be unstable. Refer to doc/policy/packages.md for documentation on package policies.
Warning: successful submission does not mean the transactions will propagate throughout the network.

Arguments:
1. package          (json array, required) An array of raw transactions.
                    The package must solely consist of a child transaction and all of its unconfirmed parents, if any. None of the parents may depend on each other.
                    The package must be topologically sorted, with the child being the last element in the array.
     [
       "rawtx",     (string)
       ...
     ]
2. maxfeerate       (numeric or string, optional, default="0.10") Reject transactions whose fee rate is higher than the specified value, expressed in BTC/kvB.
                    Fee rates larger than 1BTC/kvB are rejected.
                    Set to 0 to accept any fee rate.
3. maxburnamount    (numeric or string, optional, default="0.00") Reject transactions with provably unspendable outputs (e.g. 'datacarrier' outputs that use the OP_RETURN opcode) greater than the specified value, expressed in BTC.
                    If burning funds through unspendable outputs is desired, increase this value.
                    This check is based on heuristics and does not guarantee spendability of outputs.
                    

Result:
{                                   (json object)
  "package_msg" : "str",            (string) The transaction package result message. "success" indicates all transactions were accepted into or are already in the mempool.
  "tx-results" : {                  (json object) transaction results keyed by wtxid
    "wtxid" : {                     (json object) transaction wtxid
      "txid" : "hex",               (string) The transaction hash in hex
      "other-wtxid" : "hex",        (string, optional) The wtxid of a different transaction with the same txid but different witness found in the mempool. This means the submitted transaction was ignored.
      "vsize" : n,                  (numeric, optional) Sigops-adjusted virtual transaction size.
      "fees" : {                    (json object, optional) Transaction fees
        "base" : n,                 (numeric) transaction fee in BTC
        "effective-feerate" : n,    (numeric, optional) if the transaction was not already in the mempool, the effective feerate in BTC per KvB. For example, the package feerate and/or feerate with modified fees from prioritisetransaction.
        "effective-includes" : [    (json array, optional) if effective-feerate is provided, the wtxids of the transactions whose fees and vsizes are included in effective-feerate.
          "hex",                    (string) transaction wtxid in hex
          ...
        ]
      },
      "error" : "str"               (string, optional) The transaction error string, if it was rejected by the mempool
    },
    ...
  },
  "replaced-transactions" : [       (json array, optional) List of txids of replaced transactions
    "hex",                          (string) The transaction id
    ...
  ]
}

Examples:
> curl --user myusername --data-binary '{"jsonrpc": "2.0", "id": "curltest", "method": "submitpackage", "params": [["raw-parent-tx-1", "raw-parent-tx-2", "raw-child-tx"]]}' -H 'content-type: application/json' http://127.0.0.1:8332/
> bitcoin-cli submitpackage '["raw-tx-without-unconfirmed-parents"]'


