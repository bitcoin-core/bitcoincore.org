---
name: getnodeaddresses
btcversion: 24.0.0
btcgroup: network
permalink: en/doc/24.0.0/rpc/network/getnodeaddresses/
---

getnodeaddresses ( count "network" )

Return known addresses, after filtering for quality and recency.
These can potentially be used to find new peers in the network.
The total number of addresses known to the node may be higher.

Arguments:
1. count      (numeric, optional, default=1) The maximum number of addresses to return. Specify 0 to return all known addresses.
2. network    (string, optional, default=all networks) Return only addresses of the specified network. Can be one of: ipv4, ipv6, onion, i2p, cjdns.

Result:
[                         (json array)
  {                       (json object)
    "time" : xxx,         (numeric) The UNIX epoch time when the node was last seen
    "services" : n,       (numeric) The services offered by the node
    "address" : "str",    (string) The address of the node
    "port" : n,           (numeric) The port number of the node
    "network" : "str"     (string) The network (ipv4, ipv6, onion, i2p, cjdns) the node connected through
  },
  ...
]

Examples:
> bitcoin-cli getnodeaddresses 8
> bitcoin-cli getnodeaddresses 4 "i2p"
> bitcoin-cli -named getnodeaddresses network=onion count=12
> curl --user myusername --data-binary '{"jsonrpc": "1.0", "id": "curltest", "method": "getnodeaddresses", "params": [8]}' -H 'content-type: text/plain;' http://127.0.0.1:8332/
> curl --user myusername --data-binary '{"jsonrpc": "1.0", "id": "curltest", "method": "getnodeaddresses", "params": [4, "i2p"]}' -H 'content-type: text/plain;' http://127.0.0.1:8332/


