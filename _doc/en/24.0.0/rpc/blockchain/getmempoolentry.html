---
name: getmempoolentry
btcversion: 24.0.0
btcgroup: blockchain
permalink: en/doc/24.0.0/rpc/blockchain/getmempoolentry/
---

getmempoolentry "txid"

Returns mempool data for given transaction

Arguments:
1. txid    (string, required) The transaction id (must be in mempool)

Result:
{                                       (json object)
  "vsize" : n,                          (numeric) virtual transaction size as defined in BIP 141. This is different from actual serialized size for witness transactions as witness data is discounted.
  "weight" : n,                         (numeric) transaction weight as defined in BIP 141.
  "time" : xxx,                         (numeric) local time transaction entered pool in seconds since 1 Jan 1970 GMT
  "height" : n,                         (numeric) block height when transaction entered pool
  "descendantcount" : n,                (numeric) number of in-mempool descendant transactions (including this one)
  "descendantsize" : n,                 (numeric) virtual transaction size of in-mempool descendants (including this one)
  "ancestorcount" : n,                  (numeric) number of in-mempool ancestor transactions (including this one)
  "ancestorsize" : n,                   (numeric) virtual transaction size of in-mempool ancestors (including this one)
  "wtxid" : "hex",                      (string) hash of serialized transaction, including witness data
  "fees" : {                            (json object)
    "base" : n,                         (numeric) transaction fee, denominated in BTC
    "modified" : n,                     (numeric) transaction fee with fee deltas used for mining priority, denominated in BTC
    "ancestor" : n,                     (numeric) transaction fees of in-mempool ancestors (including this one) with fee deltas used for mining priority, denominated in BTC
    "descendant" : n                    (numeric) transaction fees of in-mempool descendants (including this one) with fee deltas used for mining priority, denominated in BTC
  },
  "depends" : [                         (json array) unconfirmed transactions used as inputs for this transaction
    "hex",                              (string) parent transaction id
    ...
  ],
  "spentby" : [                         (json array) unconfirmed transactions spending outputs from this transaction
    "hex",                              (string) child transaction id
    ...
  ],
  "bip125-replaceable" : true|false,    (boolean) Whether this transaction signals BIP125 replaceability or has an unconfirmed ancestor signaling BIP125 replaceability.
                                        
  "unbroadcast" : true|false            (boolean) Whether this transaction is currently unbroadcast (initial broadcast not yet acknowledged by any peers)
}

Examples:
> bitcoin-cli getmempoolentry "mytxid"
> curl --user myusername --data-binary '{"jsonrpc": "1.0", "id": "curltest", "method": "getmempoolentry", "params": ["mytxid"]}' -H 'content-type: text/plain;' http://127.0.0.1:8332/


