---
name: getblockfrompeer
btcversion: 24.0.0
btcgroup: blockchain
permalink: en/doc/24.0.0/rpc/blockchain/getblockfrompeer/
---

getblockfrompeer "blockhash" peer_id

Attempt to fetch block from a given peer.

We must have the header for this block, e.g. using submitheader.
Subsequent calls for the same block and a new peer will cause the response from the previous peer to be ignored.

Returns an empty JSON object if the request was successfully scheduled.

Arguments:
1. blockhash    (string, required) The block hash to try to fetch
2. peer_id      (numeric, required) The peer to fetch it from (see getpeerinfo for peer IDs)

Result:
{}    (empty JSON object)

Examples:
> bitcoin-cli getblockfrompeer "00000000c937983704a73af28acdec37b049d214adbda81d7e2a3dd146f6ed09" 0
> curl --user myusername --data-binary '{"jsonrpc": "1.0", "id": "curltest", "method": "getblockfrompeer", "params": ["00000000c937983704a73af28acdec37b049d214adbda81d7e2a3dd146f6ed09" 0]}' -H 'content-type: text/plain;' http://127.0.0.1:8332/


