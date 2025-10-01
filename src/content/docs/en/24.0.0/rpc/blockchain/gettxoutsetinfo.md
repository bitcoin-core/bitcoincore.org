---
name: gettxoutsetinfo
btcversion: 24.0.0
btcgroup: blockchain
permalink: en/doc/24.0.0/rpc/blockchain/gettxoutsetinfo/
---

gettxoutsetinfo ( "hash_type" hash_or_height use_index )

Returns statistics about the unspent transaction output set.
Note this call may take some time if you are not using coinstatsindex.

Arguments:
1. hash_type         (string, optional, default="hash_serialized_2") Which UTXO set hash should be calculated. Options: 'hash_serialized_2' (the legacy algorithm), 'muhash', 'none'.
2. hash_or_height    (string or numeric, optional, default=the current best block) The block hash or height of the target height (only available with coinstatsindex).
3. use_index         (boolean, optional, default=true) Use coinstatsindex, if available.

Result:
{                                     (json object)
  "height" : n,                       (numeric) The block height (index) of the returned statistics
  "bestblock" : "hex",                (string) The hash of the block at which these statistics are calculated
  "txouts" : n,                       (numeric) The number of unspent transaction outputs
  "bogosize" : n,                     (numeric) Database-independent, meaningless metric indicating the UTXO set size
  "hash_serialized_2" : "hex",        (string, optional) The serialized hash (only present if 'hash_serialized_2' hash_type is chosen)
  "muhash" : "hex",                   (string, optional) The serialized hash (only present if 'muhash' hash_type is chosen)
  "transactions" : n,                 (numeric, optional) The number of transactions with unspent outputs (not available when coinstatsindex is used)
  "disk_size" : n,                    (numeric, optional) The estimated size of the chainstate on disk (not available when coinstatsindex is used)
  "total_amount" : n,                 (numeric) The total amount of coins in the UTXO set
  "total_unspendable_amount" : n,     (numeric, optional) The total amount of coins permanently excluded from the UTXO set (only available if coinstatsindex is used)
  "block_info" : {                    (json object, optional) Info on amounts in the block at this block height (only available if coinstatsindex is used)
    "prevout_spent" : n,              (numeric) Total amount of all prevouts spent in this block
    "coinbase" : n,                   (numeric) Coinbase subsidy amount of this block
    "new_outputs_ex_coinbase" : n,    (numeric) Total amount of new outputs created by this block
    "unspendable" : n,                (numeric) Total amount of unspendable outputs created in this block
    "unspendables" : {                (json object) Detailed view of the unspendable categories
      "genesis_block" : n,            (numeric) The unspendable amount of the Genesis block subsidy
      "bip30" : n,                    (numeric) Transactions overridden by duplicates (no longer possible with BIP30)
      "scripts" : n,                  (numeric) Amounts sent to scripts that are unspendable (for example OP_RETURN outputs)
      "unclaimed_rewards" : n         (numeric) Fee rewards that miners did not claim in their coinbase transaction
    }
  }
}

Examples:
> bitcoin-cli gettxoutsetinfo 
> bitcoin-cli gettxoutsetinfo "none"
> bitcoin-cli gettxoutsetinfo "none" 1000
> bitcoin-cli gettxoutsetinfo "none" '"00000000c937983704a73af28acdec37b049d214adbda81d7e2a3dd146f6ed09"'
> bitcoin-cli -named gettxoutsetinfo hash_type='muhash' use_index='false'
> curl --user myusername --data-binary '{"jsonrpc": "1.0", "id": "curltest", "method": "gettxoutsetinfo", "params": []}' -H 'content-type: text/plain;' http://127.0.0.1:8332/
> curl --user myusername --data-binary '{"jsonrpc": "1.0", "id": "curltest", "method": "gettxoutsetinfo", "params": ["none"]}' -H 'content-type: text/plain;' http://127.0.0.1:8332/
> curl --user myusername --data-binary '{"jsonrpc": "1.0", "id": "curltest", "method": "gettxoutsetinfo", "params": ["none", 1000]}' -H 'content-type: text/plain;' http://127.0.0.1:8332/
> curl --user myusername --data-binary '{"jsonrpc": "1.0", "id": "curltest", "method": "gettxoutsetinfo", "params": ["none", "00000000c937983704a73af28acdec37b049d214adbda81d7e2a3dd146f6ed09"]}' -H 'content-type: text/plain;' http://127.0.0.1:8332/


