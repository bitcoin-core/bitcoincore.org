---
name: listtransactions
btcversion: 24.0.0
btcgroup: wallet
permalink: en/doc/24.0.0/rpc/wallet/listtransactions/
---

listtransactions ( "label" count skip include_watchonly )

If a label name is provided, this will return only incoming transactions paying to addresses with the specified label.

Returns up to 'count' most recent transactions skipping the first 'from' transactions.

Arguments:
1. label                (string, optional) If set, should be a valid label name to return only incoming transactions
                        with the specified label, or "*" to disable filtering and return all transactions.
2. count                (numeric, optional, default=10) The number of transactions to return
3. skip                 (numeric, optional, default=0) The number of transactions to skip
4. include_watchonly    (boolean, optional, default=true for watch-only wallets, otherwise false) Include transactions to watch-only addresses (see 'importaddress')

Result:
[                                        (json array)
  {                                      (json object)
    "involvesWatchonly" : true|false,    (boolean, optional) Only returns true if imported addresses were involved in transaction.
    "address" : "str",                   (string, optional) The bitcoin address of the transaction (not returned if the output does not have an address, e.g. OP_RETURN null data).
    "category" : "str",                  (string) The transaction category.
                                         "send"                  Transactions sent.
                                         "receive"               Non-coinbase transactions received.
                                         "generate"              Coinbase transactions received with more than 100 confirmations.
                                         "immature"              Coinbase transactions received with 100 or fewer confirmations.
                                         "orphan"                Orphaned coinbase transactions received.
    "amount" : n,                        (numeric) The amount in BTC. This is negative for the 'send' category, and is positive
                                         for all other categories
    "label" : "str",                     (string, optional) A comment for the address/transaction, if any
    "vout" : n,                          (numeric) the vout value
    "fee" : n,                           (numeric, optional) The amount of the fee in BTC. This is negative and only available for the
                                         'send' category of transactions.
    "confirmations" : n,                 (numeric) The number of confirmations for the transaction. Negative confirmations means the
                                         transaction conflicted that many blocks ago.
    "generated" : true|false,            (boolean, optional) Only present if the transaction's only input is a coinbase one.
    "trusted" : true|false,              (boolean, optional) Whether we consider the transaction to be trusted and safe to spend from.
                                         Only present when the transaction has 0 confirmations (or negative confirmations, if conflicted).
    "blockhash" : "hex",                 (string, optional) The block hash containing the transaction.
    "blockheight" : n,                   (numeric, optional) The block height containing the transaction.
    "blockindex" : n,                    (numeric, optional) The index of the transaction in the block that includes it.
    "blocktime" : xxx,                   (numeric, optional) The block time expressed in UNIX epoch time.
    "txid" : "hex",                      (string) The transaction id.
    "wtxid" : "hex",                     (string) The hash of serialized transaction, including witness data.
    "walletconflicts" : [                (json array) Conflicting transaction ids.
      "hex",                             (string) The transaction id.
      ...
    ],
    "replaced_by_txid" : "hex",          (string, optional) The txid if this tx was replaced.
    "replaces_txid" : "hex",             (string, optional) The txid if the tx replaces one.
    "comment" : "str",                   (string, optional)
    "to" : "str",                        (string, optional) If a comment to is associated with the transaction.
    "time" : xxx,                        (numeric) The transaction time expressed in UNIX epoch time.
    "timereceived" : xxx,                (numeric) The time received expressed in UNIX epoch time.
    "comment" : "str",                   (string, optional) If a comment is associated with the transaction, only present if not empty.
    "bip125-replaceable" : "str",        (string) ("yes|no|unknown") Whether this transaction signals BIP125 replaceability or has an unconfirmed ancestor signaling BIP125 replaceability.
                                         May be unknown for unconfirmed transactions not in the mempool because their unconfirmed ancestors are unknown.
    "parent_descs" : [                   (json array, optional) Only if 'category' is 'received'. List of parent descriptors for the scriptPubKey of this coin.
      "str",                             (string) The descriptor string.
      ...
    ],
    "abandoned" : true|false             (boolean, optional) 'true' if the transaction has been abandoned (inputs are respendable). Only available for the 
                                         'send' category of transactions.
  },
  ...
]

Examples:

List the most recent 10 transactions in the systems
> bitcoin-cli listtransactions 

List transactions 100 to 120
> bitcoin-cli listtransactions "*" 20 100

As a JSON-RPC call
> curl --user myusername --data-binary '{"jsonrpc": "1.0", "id": "curltest", "method": "listtransactions", "params": ["*", 20, 100]}' -H 'content-type: text/plain;' http://127.0.0.1:8332/


