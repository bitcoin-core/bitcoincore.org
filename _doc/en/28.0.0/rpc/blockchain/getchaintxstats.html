---
name: getchaintxstats
btcversion: 28.0.0
btcgroup: blockchain
permalink: en/doc/28.0.0/rpc/blockchain/getchaintxstats/
---

getchaintxstats ( nblocks "blockhash" )

Compute statistics about the total number and rate of transactions in the chain.

Arguments:
1. nblocks      (numeric, optional, default=one month) Size of the window in number of blocks
2. blockhash    (string, optional, default=chain tip) The hash of the block that ends the window.

Result:
{                                       (json object)
  "time" : xxx,                         (numeric) The timestamp for the final block in the window, expressed in UNIX epoch time
  "txcount" : n,                        (numeric, optional) The total number of transactions in the chain up to that point, if known. It may be unknown when using assumeutxo.
  "window_final_block_hash" : "hex",    (string) The hash of the final block in the window
  "window_final_block_height" : n,      (numeric) The height of the final block in the window.
  "window_block_count" : n,             (numeric) Size of the window in number of blocks
  "window_interval" : n,                (numeric, optional) The elapsed time in the window in seconds. Only returned if "window_block_count" is > 0
  "window_tx_count" : n,                (numeric, optional) The number of transactions in the window. Only returned if "window_block_count" is > 0 and if txcount exists for the start and end of the window.
  "txrate" : n                          (numeric, optional) The average rate of transactions per second in the window. Only returned if "window_interval" is > 0 and if window_tx_count exists.
}

Examples:
> bitcoin-cli getchaintxstats 
> curl --user myusername --data-binary '{"jsonrpc": "2.0", "id": "curltest", "method": "getchaintxstats", "params": [2016]}' -H 'content-type: application/json' http://127.0.0.1:8332/


