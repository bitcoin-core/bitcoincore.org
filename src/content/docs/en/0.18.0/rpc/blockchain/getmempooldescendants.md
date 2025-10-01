---
name: getmempooldescendants
btcversion: 0.18.0
btcgroup: blockchain
permalink: en/doc/0.18.0/rpc/blockchain/getmempooldescendants/
---

getmempooldescendants "txid" ( verbose )

If txid is in the mempool, returns all in-mempool descendants.

Arguments:
1. txid       (string, required) The transaction id (must be in mempool)
2. verbose    (boolean, optional, default=false) True for a json object, false for array of transaction ids

Result (for verbose = false):
[                       (json array of strings)
  "transactionid"           (string) The transaction id of an in-mempool descendant transaction
  ,...
]

Result (for verbose = true):
{                           (json object)
  "transactionid" : {       (json object)
    "size" : n,             (numeric) virtual transaction size as defined in BIP 141. This is different from actual serialized size for witness transactions as witness data is discounted.
    "fee" : n,              (numeric) transaction fee in BTC (DEPRECATED)
    "modifiedfee" : n,      (numeric) transaction fee with fee deltas used for mining priority (DEPRECATED)
    "time" : n,             (numeric) local time transaction entered pool in seconds since 1 Jan 1970 GMT
    "height" : n,           (numeric) block height when transaction entered pool
    "descendantcount" : n,  (numeric) number of in-mempool descendant transactions (including this one)
    "descendantsize" : n,   (numeric) virtual transaction size of in-mempool descendants (including this one)
    "descendantfees" : n,   (numeric) modified fees (see above) of in-mempool descendants (including this one) (DEPRECATED)
    "ancestorcount" : n,    (numeric) number of in-mempool ancestor transactions (including this one)
    "ancestorsize" : n,     (numeric) virtual transaction size of in-mempool ancestors (including this one)
    "ancestorfees" : n,     (numeric) modified fees (see above) of in-mempool ancestors (including this one) (DEPRECATED)
    "wtxid" : hash,         (string) hash of serialized transaction, including witness data
    "fees" : {
        "base" : n,         (numeric) transaction fee in BTC
        "modified" : n,     (numeric) transaction fee with fee deltas used for mining priority in BTC
        "ancestor" : n,     (numeric) modified fees (see above) of in-mempool ancestors (including this one) in BTC
        "descendant" : n,   (numeric) modified fees (see above) of in-mempool descendants (including this one) in BTC
    }
    "depends" : [           (array) unconfirmed transactions used as inputs for this transaction
        "transactionid",    (string) parent transaction id
       ... ]
    "spentby" : [           (array) unconfirmed transactions spending outputs from this transaction
        "transactionid",    (string) child transaction id
       ... ]
    "bip125-replaceable" : true|false,  (boolean) Whether this transaction could be replaced due to BIP125 (replace-by-fee)
  }, ...
}

Examples:
> bitcoin-cli getmempooldescendants "mytxid"
> curl --user myusername --data-binary '{"jsonrpc": "1.0", "id":"curltest", "method": "getmempooldescendants", "params": ["mytxid"] }' -H 'content-type: text/plain;' http://127.0.0.1:8332/


