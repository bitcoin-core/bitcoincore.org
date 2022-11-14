---
name: gettransaction
btcversion: 24.0.0
btcgroup: wallet
permalink: en/doc/24.0.0/rpc/wallet/gettransaction/
---

gettransaction "txid" ( include_watchonly verbose )

Get detailed information about in-wallet transaction <txid>

Arguments:
1. txid                 (string, required) The transaction id
2. include_watchonly    (boolean, optional, default=true for watch-only wallets, otherwise false) Whether to include watch-only addresses in balance calculation and details[]
3. verbose              (boolean, optional, default=false) Whether to include a `decoded` field containing the decoded transaction (equivalent to RPC decoderawtransaction)

Result:
{                                          (json object)
  "amount" : n,                            (numeric) The amount in BTC
  "fee" : n,                               (numeric, optional) The amount of the fee in BTC. This is negative and only available for the
                                           'send' category of transactions.
  "confirmations" : n,                     (numeric) The number of confirmations for the transaction. Negative confirmations means the
                                           transaction conflicted that many blocks ago.
  "generated" : true|false,                (boolean, optional) Only present if the transaction's only input is a coinbase one.
  "trusted" : true|false,                  (boolean, optional) Whether we consider the transaction to be trusted and safe to spend from.
                                           Only present when the transaction has 0 confirmations (or negative confirmations, if conflicted).
  "blockhash" : "hex",                     (string, optional) The block hash containing the transaction.
  "blockheight" : n,                       (numeric, optional) The block height containing the transaction.
  "blockindex" : n,                        (numeric, optional) The index of the transaction in the block that includes it.
  "blocktime" : xxx,                       (numeric, optional) The block time expressed in UNIX epoch time.
  "txid" : "hex",                          (string) The transaction id.
  "wtxid" : "hex",                         (string) The hash of serialized transaction, including witness data.
  "walletconflicts" : [                    (json array) Conflicting transaction ids.
    "hex",                                 (string) The transaction id.
    ...
  ],
  "replaced_by_txid" : "hex",              (string, optional) The txid if this tx was replaced.
  "replaces_txid" : "hex",                 (string, optional) The txid if the tx replaces one.
  "comment" : "str",                       (string, optional)
  "to" : "str",                            (string, optional) If a comment to is associated with the transaction.
  "time" : xxx,                            (numeric) The transaction time expressed in UNIX epoch time.
  "timereceived" : xxx,                    (numeric) The time received expressed in UNIX epoch time.
  "comment" : "str",                       (string, optional) If a comment is associated with the transaction, only present if not empty.
  "bip125-replaceable" : "str",            (string) ("yes|no|unknown") Whether this transaction signals BIP125 replaceability or has an unconfirmed ancestor signaling BIP125 replaceability.
                                           May be unknown for unconfirmed transactions not in the mempool because their unconfirmed ancestors are unknown.
  "parent_descs" : [                       (json array, optional) Only if 'category' is 'received'. List of parent descriptors for the scriptPubKey of this coin.
    "str",                                 (string) The descriptor string.
    ...
  ],
  "details" : [                            (json array)
    {                                      (json object)
      "involvesWatchonly" : true|false,    (boolean, optional) Only returns true if imported addresses were involved in transaction.
      "address" : "str",                   (string, optional) The bitcoin address involved in the transaction.
      "category" : "str",                  (string) The transaction category.
                                           "send"                  Transactions sent.
                                           "receive"               Non-coinbase transactions received.
                                           "generate"              Coinbase transactions received with more than 100 confirmations.
                                           "immature"              Coinbase transactions received with 100 or fewer confirmations.
                                           "orphan"                Orphaned coinbase transactions received.
      "amount" : n,                        (numeric) The amount in BTC
      "label" : "str",                     (string, optional) A comment for the address/transaction, if any
      "vout" : n,                          (numeric) the vout value
      "fee" : n,                           (numeric, optional) The amount of the fee in BTC. This is negative and only available for the 
                                           'send' category of transactions.
      "abandoned" : true|false,            (boolean, optional) 'true' if the transaction has been abandoned (inputs are respendable). Only available for the 
                                           'send' category of transactions.
      "parent_descs" : [                   (json array, optional) Only if 'category' is 'received'. List of parent descriptors for the scriptPubKey of this coin.
        "str",                             (string) The descriptor string.
        ...
      ]
    },
    ...
  ],
  "hex" : "hex",                           (string) Raw data for transaction
  "decoded" : {                            (json object, optional) The decoded transaction (only present when `verbose` is passed)
    ...                                    Equivalent to the RPC decoderawtransaction method, or the RPC getrawtransaction method when `verbose` is passed.
  }
}

Examples:
> bitcoin-cli gettransaction "1075db55d416d3ca199f55b6084e2115b9345e16c5cf302fc80e9d5fbf5d48d"
> bitcoin-cli gettransaction "1075db55d416d3ca199f55b6084e2115b9345e16c5cf302fc80e9d5fbf5d48d" true
> bitcoin-cli gettransaction "1075db55d416d3ca199f55b6084e2115b9345e16c5cf302fc80e9d5fbf5d48d" false true
> curl --user myusername --data-binary '{"jsonrpc": "1.0", "id": "curltest", "method": "gettransaction", "params": ["1075db55d416d3ca199f55b6084e2115b9345e16c5cf302fc80e9d5fbf5d48d"]}' -H 'content-type: text/plain;' http://127.0.0.1:8332/


