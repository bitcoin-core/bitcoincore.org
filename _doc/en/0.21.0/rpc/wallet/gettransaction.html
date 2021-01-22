---
name: gettransaction
btcversion: 0.21.0
btcgroup: wallet
permalink: en/doc/0.21.0/rpc/wallet/gettransaction/
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
  "fee" : n,                               (numeric) The amount of the fee in BTC. This is negative and only available for the
                                           'send' category of transactions.
  "confirmations" : n,                     (numeric) The number of confirmations for the transaction. Negative confirmations means the
                                           transaction conflicted that many blocks ago.
  "generated" : true|false,                (boolean) Only present if transaction only input is a coinbase one.
  "trusted" : true|false,                  (boolean) Only present if we consider transaction to be trusted and so safe to spend from.
  "blockhash" : "hex",                     (string) The block hash containing the transaction.
  "blockheight" : n,                       (numeric) The block height containing the transaction.
  "blockindex" : n,                        (numeric) The index of the transaction in the block that includes it.
  "blocktime" : xxx,                       (numeric) The block time expressed in UNIX epoch time.
  "txid" : "hex",                          (string) The transaction id.
  "walletconflicts" : [                    (json array) Conflicting transaction ids.
    "hex",                                 (string) The transaction id.
    ...
  ],
  "time" : xxx,                            (numeric) The transaction time expressed in UNIX epoch time.
  "timereceived" : xxx,                    (numeric) The time received expressed in UNIX epoch time.
  "comment" : "str",                       (string) If a comment is associated with the transaction, only present if not empty.
  "bip125-replaceable" : "str",            (string) ("yes|no|unknown") Whether this transaction could be replaced due to BIP125 (replace-by-fee);
                                           may be unknown for unconfirmed transactions not in the mempool
  "details" : [                            (json array)
    {                                      (json object)
      "involvesWatchonly" : true|false,    (boolean) Only returns true if imported addresses were involved in transaction.
      "address" : "str",                   (string) The bitcoin address involved in the transaction.
      "category" : "str",                  (string) The transaction category.
                                           "send"                  Transactions sent.
                                           "receive"               Non-coinbase transactions received.
                                           "generate"              Coinbase transactions received with more than 100 confirmations.
                                           "immature"              Coinbase transactions received with 100 or fewer confirmations.
                                           "orphan"                Orphaned coinbase transactions received.
      "amount" : n,                        (numeric) The amount in BTC
      "label" : "str",                     (string) A comment for the address/transaction, if any
      "vout" : n,                          (numeric) the vout value
      "fee" : n,                           (numeric) The amount of the fee in BTC. This is negative and only available for the 
                                           'send' category of transactions.
      "abandoned" : true|false             (boolean) 'true' if the transaction has been abandoned (inputs are respendable). Only available for the 
                                           'send' category of transactions.
    },
    ...
  ],
  "hex" : "hex",                           (string) Raw data for transaction
  "decoded" : {                            (json object) Optional, the decoded transaction (only present when `verbose` is passed)
    ...                                    Equivalent to the RPC decoderawtransaction method, or the RPC getrawtransaction method when `verbose` is passed.
  }
}

Examples:
> bitcoin-cli gettransaction "1075db55d416d3ca199f55b6084e2115b9345e16c5cf302fc80e9d5fbf5d48d"
> bitcoin-cli gettransaction "1075db55d416d3ca199f55b6084e2115b9345e16c5cf302fc80e9d5fbf5d48d" true
> bitcoin-cli gettransaction "1075db55d416d3ca199f55b6084e2115b9345e16c5cf302fc80e9d5fbf5d48d" false true
> curl --user myusername --data-binary '{"jsonrpc": "1.0", "id": "curltest", "method": "gettransaction", "params": ["1075db55d416d3ca199f55b6084e2115b9345e16c5cf302fc80e9d5fbf5d48d"]}' -H 'content-type: text/plain;' http://127.0.0.1:8332/


