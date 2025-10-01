---
name: gettxoutsetinfo
btcversion: 0.18.0
btcgroup: blockchain
permalink: en/doc/0.18.0/rpc/blockchain/gettxoutsetinfo/
---

gettxoutsetinfo

Returns statistics about the unspent transaction output set.
Note this call may take some time.

Result:
{
  "height":n,     (numeric) The current block height (index)
  "bestblock": "hex",   (string) The hash of the block at the tip of the chain
  "transactions": n,      (numeric) The number of transactions with unspent outputs
  "txouts": n,            (numeric) The number of unspent transaction outputs
  "bogosize": n,          (numeric) A meaningless metric for UTXO set size
  "hash_serialized_2": "hash", (string) The serialized hash
  "disk_size": n,         (numeric) The estimated size of the chainstate on disk
  "total_amount": x.xxx          (numeric) The total amount
}

Examples:
> bitcoin-cli gettxoutsetinfo 
> curl --user myusername --data-binary '{"jsonrpc": "1.0", "id":"curltest", "method": "gettxoutsetinfo", "params": [] }' -H 'content-type: text/plain;' http://127.0.0.1:8332/


