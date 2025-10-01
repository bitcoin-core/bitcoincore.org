---
name: getnetworkhashps
btcversion: 24.0.0
btcgroup: mining
permalink: en/doc/24.0.0/rpc/mining/getnetworkhashps/
---

getnetworkhashps ( nblocks height )

Returns the estimated network hashes per second based on the last n blocks.
Pass in [blocks] to override # of blocks, -1 specifies since last difficulty change.
Pass in [height] to estimate the network speed at the time when a certain block was found.

Arguments:
1. nblocks    (numeric, optional, default=120) The number of blocks, or -1 for blocks since last difficulty change.
2. height     (numeric, optional, default=-1) To estimate at the time of the given height.

Result:
n    (numeric) Hashes per second estimated

Examples:
> bitcoin-cli getnetworkhashps 
> curl --user myusername --data-binary '{"jsonrpc": "1.0", "id": "curltest", "method": "getnetworkhashps", "params": []}' -H 'content-type: text/plain;' http://127.0.0.1:8332/


