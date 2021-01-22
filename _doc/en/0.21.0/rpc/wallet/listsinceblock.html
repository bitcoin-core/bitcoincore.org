---
name: listsinceblock
btcversion: 0.21.0
btcgroup: wallet
permalink: en/doc/0.21.0/rpc/wallet/listsinceblock/
---

listsinceblock ( "blockhash" target_confirmations include_watchonly include_removed )

Get all transactions in blocks since block [blockhash], or all transactions if omitted.
If "blockhash" is no longer a part of the main chain, transactions from the fork point onward are included.
Additionally, if include_removed is set, transactions affecting the wallet which were removed are returned in the "removed" array.

Arguments:
1. blockhash               (string, optional) If set, the block hash to list transactions since, otherwise list all transactions.
2. target_confirmations    (numeric, optional, default=1) Return the nth block hash from the main chain. e.g. 1 would mean the best block hash. Note: this is not used as a filter, but only affects [lastblock] in the return value
3. include_watchonly       (boolean, optional, default=true for watch-only wallets, otherwise false) Include transactions to watch-only addresses (see 'importaddress')
4. include_removed         (boolean, optional, default=true) Show transactions that were removed due to a reorg in the "removed" array
                           (not guaranteed to work on pruned nodes)

Result:
{                                          (json object)
  "transactions" : [                       (json array)
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
      "abandoned" : true|false,            (boolean) 'true' if the transaction has been abandoned (inputs are respendable). Only available for the 
                                           'send' category of transactions.
      "label" : "str",                     (string) A comment for the address/transaction, if any
      "to" : "str"                         (string) If a comment to is associated with the transaction.
    },
    ...
  ],
  "removed" : [                            (json array) <structure is the same as "transactions" above, only present if include_removed=true>
                                           Note: transactions that were re-added in the active chain will appear as-is in this array, and may thus have a positive confirmation count.
    ...
  ],
  "lastblock" : "hex"                      (string) The hash of the block (target_confirmations-1) from the best block on the main chain, or the genesis hash if the referenced block does not exist yet. This is typically used to feed back into listsinceblock the next time you call it. So you would generally use a target_confirmations of say 6, so you will be continually re-notified of transactions until they've reached 6 confirmations plus any new ones
}

Examples:
> bitcoin-cli listsinceblock 
> bitcoin-cli listsinceblock "000000000000000bacf66f7497b7dc45ef753ee9a7d38571037cdb1a57f663ad" 6
> curl --user myusername --data-binary '{"jsonrpc": "1.0", "id": "curltest", "method": "listsinceblock", "params": ["000000000000000bacf66f7497b7dc45ef753ee9a7d38571037cdb1a57f663ad", 6]}' -H 'content-type: text/plain;' http://127.0.0.1:8332/


