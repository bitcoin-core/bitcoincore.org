---
name: listunspent
btcversion: 24.0.0
btcgroup: wallet
permalink: en/doc/24.0.0/rpc/wallet/listunspent/
---

listunspent ( minconf maxconf ["address",...] include_unsafe query_options )

Returns array of unspent transaction outputs
with between minconf and maxconf (inclusive) confirmations.
Optionally filter to only include txouts paid to specified addresses.

Arguments:
1. minconf                            (numeric, optional, default=1) The minimum confirmations to filter
2. maxconf                            (numeric, optional, default=9999999) The maximum confirmations to filter
3. addresses                          (json array, optional, default=[]) The bitcoin addresses to filter
     [
       "address",                     (string) bitcoin address
       ...
     ]
4. include_unsafe                     (boolean, optional, default=true) Include outputs that are not safe to spend
                                      See description of "safe" attribute below.
5. query_options                      (json object, optional) JSON with query options
     {
       "minimumAmount": amount,       (numeric or string, optional, default="0.00") Minimum value of each UTXO in BTC
       "maximumAmount": amount,       (numeric or string, optional, default=unlimited) Maximum value of each UTXO in BTC
       "maximumCount": n,             (numeric, optional, default=unlimited) Maximum number of UTXOs
       "minimumSumAmount": amount,    (numeric or string, optional, default=unlimited) Minimum sum value of all UTXOs in BTC
     }

Result:
[                                (json array)
  {                              (json object)
    "txid" : "hex",              (string) the transaction id
    "vout" : n,                  (numeric) the vout value
    "address" : "str",           (string, optional) the bitcoin address
    "label" : "str",             (string, optional) The associated label, or "" for the default label
    "scriptPubKey" : "str",      (string) the script key
    "amount" : n,                (numeric) the transaction output amount in BTC
    "confirmations" : n,         (numeric) The number of confirmations
    "ancestorcount" : n,         (numeric, optional) The number of in-mempool ancestor transactions, including this one (if transaction is in the mempool)
    "ancestorsize" : n,          (numeric, optional) The virtual transaction size of in-mempool ancestors, including this one (if transaction is in the mempool)
    "ancestorfees" : n,          (numeric, optional) The total fees of in-mempool ancestors (including this one) with fee deltas used for mining priority in sat (if transaction is in the mempool)
    "redeemScript" : "hex",      (string, optional) The redeemScript if scriptPubKey is P2SH
    "witnessScript" : "str",     (string, optional) witnessScript if the scriptPubKey is P2WSH or P2SH-P2WSH
    "spendable" : true|false,    (boolean) Whether we have the private keys to spend this output
    "solvable" : true|false,     (boolean) Whether we know how to spend this output, ignoring the lack of keys
    "reused" : true|false,       (boolean, optional) (only present if avoid_reuse is set) Whether this output is reused/dirty (sent to an address that was previously spent from)
    "desc" : "str",              (string, optional) (only when solvable) A descriptor for spending this output
    "parent_descs" : [           (json array) List of parent descriptors for the scriptPubKey of this coin.
      "str",                     (string) The descriptor string.
      ...
    ],
    "safe" : true|false          (boolean) Whether this output is considered safe to spend. Unconfirmed transactions
                                 from outside keys and unconfirmed replacement transactions are considered unsafe
                                 and are not eligible for spending by fundrawtransaction and sendtoaddress.
  },
  ...
]

Examples:
> bitcoin-cli listunspent 
> bitcoin-cli listunspent 6 9999999 "[\"bc1q09vm5lfy0j5reeulh4x5752q25uqqvz34hufdl\",\"bc1q02ad21edsxd23d32dfgqqsz4vv4nmtfzuklhy3\"]"
> curl --user myusername --data-binary '{"jsonrpc": "1.0", "id": "curltest", "method": "listunspent", "params": [6, 9999999 "[\"bc1q09vm5lfy0j5reeulh4x5752q25uqqvz34hufdl\",\"bc1q02ad21edsxd23d32dfgqqsz4vv4nmtfzuklhy3\"]"]}' -H 'content-type: text/plain;' http://127.0.0.1:8332/
> bitcoin-cli listunspent 6 9999999 '[]' true '{ "minimumAmount": 0.005 }'
> curl --user myusername --data-binary '{"jsonrpc": "1.0", "id": "curltest", "method": "listunspent", "params": [6, 9999999, [] , true, { "minimumAmount": 0.005 } ]}' -H 'content-type: text/plain;' http://127.0.0.1:8332/


