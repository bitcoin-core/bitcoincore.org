---
name: getblockstats
btcversion: 0.21.0
btcgroup: blockchain
permalink: en/doc/0.21.0/rpc/blockchain/getblockstats/
---

getblockstats hash_or_height ( stats )

Compute per block statistics for a given window. All amounts are in satoshis.
It won't work for some heights with pruning.

Arguments:
1. hash_or_height    (string or numeric, required) The block hash or height of the target block
2. stats             (json array, optional, default=all values) Values to plot (see result below)
     [
       "height",     (string) Selected statistic
       "time",       (string) Selected statistic
       ...
     ]

Result:
{                              (json object)
  "avgfee" : n,                (numeric) Average fee in the block
  "avgfeerate" : n,            (numeric) Average feerate (in satoshis per virtual byte)
  "avgtxsize" : n,             (numeric) Average transaction size
  "blockhash" : "hex",         (string) The block hash (to check for potential reorgs)
  "feerate_percentiles" : [    (json array) Feerates at the 10th, 25th, 50th, 75th, and 90th percentile weight unit (in satoshis per virtual byte)
    n,                         (numeric) The 10th percentile feerate
    n,                         (numeric) The 25th percentile feerate
    n,                         (numeric) The 50th percentile feerate
    n,                         (numeric) The 75th percentile feerate
    n                          (numeric) The 90th percentile feerate
  ],
  "height" : n,                (numeric) The height of the block
  "ins" : n,                   (numeric) The number of inputs (excluding coinbase)
  "maxfee" : n,                (numeric) Maximum fee in the block
  "maxfeerate" : n,            (numeric) Maximum feerate (in satoshis per virtual byte)
  "maxtxsize" : n,             (numeric) Maximum transaction size
  "medianfee" : n,             (numeric) Truncated median fee in the block
  "mediantime" : n,            (numeric) The block median time past
  "mediantxsize" : n,          (numeric) Truncated median transaction size
  "minfee" : n,                (numeric) Minimum fee in the block
  "minfeerate" : n,            (numeric) Minimum feerate (in satoshis per virtual byte)
  "mintxsize" : n,             (numeric) Minimum transaction size
  "outs" : n,                  (numeric) The number of outputs
  "subsidy" : n,               (numeric) The block subsidy
  "swtotal_size" : n,          (numeric) Total size of all segwit transactions
  "swtotal_weight" : n,        (numeric) Total weight of all segwit transactions
  "swtxs" : n,                 (numeric) The number of segwit transactions
  "time" : n,                  (numeric) The block time
  "total_out" : n,             (numeric) Total amount in all outputs (excluding coinbase and thus reward [ie subsidy + totalfee])
  "total_size" : n,            (numeric) Total size of all non-coinbase transactions
  "total_weight" : n,          (numeric) Total weight of all non-coinbase transactions
  "totalfee" : n,              (numeric) The fee total
  "txs" : n,                   (numeric) The number of transactions (including coinbase)
  "utxo_increase" : n,         (numeric) The increase/decrease in the number of unspent outputs
  "utxo_size_inc" : n          (numeric) The increase/decrease in size for the utxo index (not discounting op_return and similar)
}

Examples:
> bitcoin-cli getblockstats '"00000000c937983704a73af28acdec37b049d214adbda81d7e2a3dd146f6ed09"' '["minfeerate","avgfeerate"]'
> bitcoin-cli getblockstats 1000 '["minfeerate","avgfeerate"]'
> curl --user myusername --data-binary '{"jsonrpc": "1.0", "id": "curltest", "method": "getblockstats", "params": ["00000000c937983704a73af28acdec37b049d214adbda81d7e2a3dd146f6ed09", ["minfeerate","avgfeerate"]]}' -H 'content-type: text/plain;' http://127.0.0.1:8332/
> curl --user myusername --data-binary '{"jsonrpc": "1.0", "id": "curltest", "method": "getblockstats", "params": [1000, ["minfeerate","avgfeerate"]]}' -H 'content-type: text/plain;' http://127.0.0.1:8332/


