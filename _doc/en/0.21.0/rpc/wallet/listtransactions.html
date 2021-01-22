---
name: listtransactions
btcversion: 0.21.0
btcgroup: wallet
permalink: en/doc/0.21.0/rpc/wallet/listtransactions/
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
    "involvesWatchonly" : true|false,    (boolean) Only returns true if imported addresses were involved in transaction.
    "address" : "str",                   (string) The bitcoin address of the transaction.
    "category" : "str",                  (string) The transaction category.
                                         "send"                  Transactions sent.
                                         "receive"               Non-coinbase transactions received.
                                         "generate"              Coinbase transactions received with more than 100 confirmations.
                                         "immature"              Coinbase transactions received with 100 or fewer confirmations.
                                         "orphan"                Orphaned coinbase transactions received.
    "amount" : n,                        (numeric) The amount in BTC. This is negative for the 'send' category, and is positive
                                         for all other categories
    "label" : "str",                     (string) A comment for the address/transaction, if any
    "vout" : n,                          (numeric) the vout value
    "fee" : n,                           (numeric) The amount of the fee in BTC. This is negative and only available for the
                                         'send' category of transactions.
    "confirmations" : n,                 (numeric) The number of confirmations for the transaction. Negative confirmations means the
                                         transaction conflicted that many blocks ago.
    "generated" : true|false,            (boolean) Only present if transaction only input is a coinbase one.
    "trusted" : true|false,              (boolean) Only present if we consider transaction to be trusted and so safe to spend from.
    "blockhash" : "hex",                 (string) The block hash containing the transaction.
    "blockheight" : n,                   (numeric) The block height containing the transaction.
    "blockindex" : n,                    (numeric) The index of the transaction in the block that includes it.
    "blocktime" : xxx,                   (numeric) The block time expressed in UNIX epoch time.
    "txid" : "hex",                      (string) The transaction id.
    "walletconflicts" : [                (json array) Conflicting transaction ids.
      "hex",                             (string) The transaction id.
      ...
    ],
    "time" : xxx,                        (numeric) The transaction time expressed in UNIX epoch time.
    "timereceived" : xxx,                (numeric) The time received expressed in UNIX epoch time.
    "comment" : "str",                   (string) If a comment is associated with the transaction, only present if not empty.
    "bip125-replaceable" : "str",        (string) ("yes|no|unknown") Whether this transaction could be replaced due to BIP125 (replace-by-fee);
                                         may be unknown for unconfirmed transactions not in the mempool
    "abandoned" : true|false             (boolean) 'true' if the transaction has been abandoned (inputs are respendable). Only available for the 
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


