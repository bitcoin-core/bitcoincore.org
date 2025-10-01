---
name: dumptxoutset
btcversion: 29.0.0
btcgroup: blockchain
permalink: en/doc/29.0.0/rpc/blockchain/dumptxoutset/
---

dumptxoutset "path" ( "type" {"rollback":n,...} )

Write the serialized UTXO set to a file. This can be used in loadtxoutset afterwards if this snapshot height is supported in the chainparams as well.

Unless the "latest" type is requested, the node will roll back to the requested height and network activity will be suspended during this process. Because of this it is discouraged to interact with the node in any other way during the execution of this call to avoid inconsistent results and race conditions, particularly RPCs that interact with blockstorage.

This call may take several minutes. Make sure to use no RPC timeout (bitcoin-cli -rpcclienttimeout=0)

Arguments:
1. path       (string, required) Path to the output file. If relative, will be prefixed by datadir.
2. type       (string, optional, default="") The type of snapshot to create. Can be "latest" to create a snapshot of the current UTXO set or "rollback" to temporarily roll back the state of the node to a historical block before creating the snapshot of a historical UTXO set. This parameter can be omitted if a separate "rollback" named parameter is specified indicating the height or hash of a specific historical block. If "rollback" is specified and separate "rollback" named parameter is not specified, this will roll back to the latest valid snapshot block that can currently be loaded with loadtxoutset.
3. options    (json object, optional) Options object that can be used to pass named arguments, listed below.

Named Arguments:
rollback    (string or numeric, optional) Height or hash of the block to roll back to before creating the snapshot. Note: The further this number is from the tip, the longer this process will take. Consider setting a higher -rpcclienttimeout value in this case.

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
> bitcoin-cli -rpcclienttimeout=0 dumptxoutset utxo.dat latest
> bitcoin-cli -rpcclienttimeout=0 dumptxoutset utxo.dat rollback
> bitcoin-cli -rpcclienttimeout=0 -named dumptxoutset utxo.dat rollback=853456


