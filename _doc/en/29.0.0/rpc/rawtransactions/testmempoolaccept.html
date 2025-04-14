---
name: testmempoolaccept
btcversion: 29.0.0
btcgroup: rawtransactions
permalink: en/doc/29.0.0/rpc/rawtransactions/testmempoolaccept/
---

testmempoolaccept ["rawtx",...] ( maxfeerate )

Returns result of mempool acceptance tests indicating if raw transaction(s) (serialized, hex-encoded) would be accepted by mempool.

If multiple transactions are passed in, parents must come before children and package policies apply: the transactions cannot conflict with any mempool transactions or each other.

If one transaction fails, other transactions may not be fully validated (the 'allowed' key will be blank).

The maximum number of transactions allowed is 25.

This checks if transactions violate the consensus or policy rules.

See sendrawtransaction call.

Arguments:
1. rawtxs          (json array, required) An array of hex strings of raw transactions.
     [
       "rawtx",    (string)
       ...
     ]
2. maxfeerate      (numeric or string, optional, default="0.10") Reject transactions whose fee rate is higher than the specified value, expressed in BTC/kvB.
                   Fee rates larger than 1BTC/kvB are rejected.
                   Set to 0 to accept any fee rate.

Result:
[                                 (json array) The result of the mempool acceptance test for each raw transaction in the input array.
                                  Returns results for each transaction in the same order they were passed in.
                                  Transactions that cannot be fully validated due to failures in other transactions will not contain an 'allowed' result.
                                  
  {                               (json object)
    "txid" : "hex",               (string) The transaction hash in hex
    "wtxid" : "hex",              (string) The transaction witness hash in hex
    "package-error" : "str",      (string, optional) Package validation error, if any (only possible if rawtxs had more than 1 transaction).
    "allowed" : true|false,       (boolean, optional) Whether this tx would be accepted to the mempool and pass client-specified maxfeerate. If not present, the tx was not fully validated due to a failure in another tx in the list.
    "vsize" : n,                  (numeric, optional) Virtual transaction size as defined in BIP 141. This is different from actual serialized size for witness transactions as witness data is discounted (only present when 'allowed' is true)
    "fees" : {                    (json object, optional) Transaction fees (only present if 'allowed' is true)
      "base" : n,                 (numeric) transaction fee in BTC
      "effective-feerate" : n,    (numeric) the effective feerate in BTC per KvB. May differ from the base feerate if, for example, there are modified fees from prioritisetransaction or a package feerate was used.
      "effective-includes" : [    (json array) transactions whose fees and vsizes are included in effective-feerate.
        "hex",                    (string) transaction wtxid in hex
        ...
      ]
    },
    "reject-reason" : "str",      (string, optional) Rejection reason (only present when 'allowed' is false)
    "reject-details" : "str"      (string, optional) Rejection details (only present when 'allowed' is false and rejection details exist)
  },
  ...
]

Examples:

Create a transaction
> bitcoin-cli createrawtransaction "[{\"txid\" : \"mytxid\",\"vout\":0}]" "{\"myaddress\":0.01}"
Sign the transaction, and get back the hex
> bitcoin-cli signrawtransactionwithwallet "myhex"

Test acceptance of the transaction (signed hex)
> bitcoin-cli testmempoolaccept '["signedhex"]'

As a JSON-RPC call
> curl --user myusername --data-binary '{"jsonrpc": "2.0", "id": "curltest", "method": "testmempoolaccept", "params": [["signedhex"]]}' -H 'content-type: application/json' http://127.0.0.1:8332/


