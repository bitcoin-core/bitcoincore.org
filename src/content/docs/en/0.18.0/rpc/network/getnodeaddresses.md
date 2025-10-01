---
name: getnodeaddresses
btcversion: 0.18.0
btcgroup: network
permalink: en/doc/0.18.0/rpc/network/getnodeaddresses/
---

getnodeaddresses ( count )

Return known addresses which can potentially be used to find new nodes in the network

Arguments:
1. count    (numeric, optional, default=1) How many addresses to return. Limited to the smaller of 2500 or 23% of all known addresses.

Result:
[
  {
    "time": ttt,                (numeric) Timestamp in seconds since epoch (Jan 1 1970 GMT) keeping track of when the node was last seen
    "services": n,              (numeric) The services offered
    "address": "host",          (string) The address of the node
    "port": n                   (numeric) The port of the node
  }
  ,....
]

Examples:
> bitcoin-cli getnodeaddresses 8
> curl --user myusername --data-binary '{"jsonrpc": "1.0", "id":"curltest", "method": "getnodeaddresses", "params": [8] }' -H 'content-type: text/plain;' http://127.0.0.1:8332/


