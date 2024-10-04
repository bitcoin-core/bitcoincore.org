---
name: gettxout
btcversion: 28.0.0
btcgroup: blockchain
permalink: en/doc/28.0.0/rpc/blockchain/gettxout/
---

gettxout "txid" n ( include_mempool )

Returns details about an unspent transaction output.

Arguments:
1. txid               (string, required) The transaction id
2. n                  (numeric, required) vout number
3. include_mempool    (boolean, optional, default=true) Whether to include the mempool. Note that an unspent output that is spent in the mempool won't appear.

Result (If the UTXO was not found):
null    (json null)

Result (Otherwise):
{                             (json object)
  "bestblock" : "hex",        (string) The hash of the block at the tip of the chain
  "confirmations" : n,        (numeric) The number of confirmations
  "value" : n,                (numeric) The transaction value in BTC
  "scriptPubKey" : {          (json object)
    "asm" : "str",            (string) Disassembly of the output script
    "desc" : "str",           (string) Inferred descriptor for the output
    "hex" : "hex",            (string) The raw output script bytes, hex-encoded
    "type" : "str",           (string) The type, eg pubkeyhash
    "address" : "str"         (string, optional) The Bitcoin address (only if a well-defined address exists)
  },
  "coinbase" : true|false     (boolean) Coinbase or not
}

Examples:

Get unspent transactions
> bitcoin-cli listunspent 

View the details
> bitcoin-cli gettxout "txid" 1

As a JSON-RPC call
> curl --user myusername --data-binary '{"jsonrpc": "2.0", "id": "curltest", "method": "gettxout", "params": ["txid", 1]}' -H 'content-type: application/json' http://127.0.0.1:8332/


