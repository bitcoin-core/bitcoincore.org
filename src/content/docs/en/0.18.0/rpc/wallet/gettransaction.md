---
name: gettransaction
btcversion: 0.18.0
btcgroup: wallet
permalink: en/doc/0.18.0/rpc/wallet/gettransaction/
---

gettransaction "txid" ( include_watchonly )

Get detailed information about in-wallet transaction <txid>

Arguments:
1. txid                 (string, required) The transaction id
2. include_watchonly    (boolean, optional, default=false) Whether to include watch-only addresses in balance calculation and details[]

Result:
{
  "amount" : x.xxx,        (numeric) The transaction amount in BTC
  "fee": x.xxx,            (numeric) The amount of the fee in BTC. This is negative and only available for the 
                              'send' category of transactions.
  "confirmations" : n,     (numeric) The number of confirmations
  "blockhash" : "hash",  (string) The block hash
  "blockindex" : xx,       (numeric) The index of the transaction in the block that includes it
  "blocktime" : ttt,       (numeric) The time in seconds since epoch (1 Jan 1970 GMT)
  "txid" : "transactionid",   (string) The transaction id.
  "time" : ttt,            (numeric) The transaction time in seconds since epoch (1 Jan 1970 GMT)
  "timereceived" : ttt,    (numeric) The time received in seconds since epoch (1 Jan 1970 GMT)
  "bip125-replaceable": "yes|no|unknown",  (string) Whether this transaction could be replaced due to BIP125 (replace-by-fee);
                                                   may be unknown for unconfirmed transactions not in the mempool
  "details" : [
    {
      "address" : "address",          (string) The bitcoin address involved in the transaction
      "category" :                      (string) The transaction category.
                   "send"                  Transactions sent.
                   "receive"               Non-coinbase transactions received.
                   "generate"              Coinbase transactions received with more than 100 confirmations.
                   "immature"              Coinbase transactions received with 100 or fewer confirmations.
                   "orphan"                Orphaned coinbase transactions received.
      "amount" : x.xxx,                 (numeric) The amount in BTC
      "label" : "label",              (string) A comment for the address/transaction, if any
      "vout" : n,                       (numeric) the vout value
      "fee": x.xxx,                     (numeric) The amount of the fee in BTC. This is negative and only available for the 
                                           'send' category of transactions.
      "abandoned": xxx                  (bool) 'true' if the transaction has been abandoned (inputs are respendable). Only available for the 
                                           'send' category of transactions.
    }
    ,...
  ],
  "hex" : "data"         (string) Raw data for transaction
}

Examples:
> bitcoin-cli gettransaction "1075db55d416d3ca199f55b6084e2115b9345e16c5cf302fc80e9d5fbf5d48d"
> bitcoin-cli gettransaction "1075db55d416d3ca199f55b6084e2115b9345e16c5cf302fc80e9d5fbf5d48d" true
> curl --user myusername --data-binary '{"jsonrpc": "1.0", "id":"curltest", "method": "gettransaction", "params": ["1075db55d416d3ca199f55b6084e2115b9345e16c5cf302fc80e9d5fbf5d48d"] }' -H 'content-type: text/plain;' http://127.0.0.1:8332/


