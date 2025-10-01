---
name: importmempool
btcversion: 29.0.0
btcgroup: blockchain
permalink: en/doc/29.0.0/rpc/blockchain/importmempool/
---

importmempool "filepath" ( options )

Import a mempool.dat file and attempt to add its contents to the mempool.
Warning: Importing untrusted files is dangerous, especially if metadata from the file is taken over.

Arguments:
1. filepath    (string, required) The mempool file
2. options     (json object, optional) Options object that can be used to pass named arguments, listed below.

Named Arguments:
use_current_time            (boolean, optional, default=true) Whether to use the current system time or use the entry time metadata from the mempool file.
                            Warning: Importing untrusted metadata may lead to unexpected issues and undesirable behavior.
apply_fee_delta_priority    (boolean, optional, default=false) Whether to apply the fee delta metadata from the mempool file.
                            It will be added to any existing fee deltas.
                            The fee delta can be set by the prioritisetransaction RPC.
                            Warning: Importing untrusted metadata may lead to unexpected issues and undesirable behavior.
                            Only set this bool if you understand what it does.
apply_unbroadcast_set       (boolean, optional, default=false) Whether to apply the unbroadcast set metadata from the mempool file.
                            Warning: Importing untrusted metadata may lead to unexpected issues and undesirable behavior.

Result:
{}    (empty JSON object)

Examples:
> bitcoin-cli importmempool /path/to/mempool.dat
> curl --user myusername --data-binary '{"jsonrpc": "2.0", "id": "curltest", "method": "importmempool", "params": [/path/to/mempool.dat]}' -H 'content-type: application/json' http://127.0.0.1:8332/


