---
name: getprioritisedtransactions
btcversion: 26.0.0
btcgroup: mining
permalink: en/doc/26.0.0/rpc/mining/getprioritisedtransactions/
---

getprioritisedtransactions

Returns a map of all user-created (see prioritisetransaction) fee deltas by txid, and whether the tx is present in mempool.

Result:
{                                 (json object) prioritisation keyed by txid
  "txid" : {                      (json object)
    "fee_delta" : n,              (numeric) transaction fee delta in satoshis
    "in_mempool" : true|false     (boolean) whether this transaction is currently in mempool
  },
  ...
}

Examples:
> bitcoin-cli getprioritisedtransactions 
> curl --user myusername --data-binary '{"jsonrpc": "1.0", "id": "curltest", "method": "getprioritisedtransactions", "params": []}' -H 'content-type: text/plain;' http://127.0.0.1:8332/


