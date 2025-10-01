---
name: dumptxoutset
btcversion: 28.0.0
btcgroup: blockchain
permalink: en/doc/28.0.0/rpc/blockchain/dumptxoutset/
---

dumptxoutset "path"

Write the serialized UTXO set to a file.

Arguments:
1. path    (string, required) Path to the output file. If relative, will be prefixed by datadir.

Result:
{                             (json object)
  "coins_written" : n,        (numeric) the number of coins written in the snapshot
  "base_hash" : "hex",        (string) the hash of the base of the snapshot
  "base_height" : n,          (numeric) the height of the base of the snapshot
  "path" : "str",             (string) the absolute path that the snapshot was written to
  "txoutset_hash" : "hex",    (string) the hash of the UTXO set contents
  "nchaintx" : n              (numeric) the number of transactions in the chain up to and including the base block
}

Examples:
> bitcoin-cli dumptxoutset utxo.dat


