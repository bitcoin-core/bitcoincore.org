---
name: listbanned
btcversion: 24.0.0
btcgroup: network
permalink: en/doc/24.0.0/rpc/network/listbanned/
---

listbanned

List all manually banned IPs/Subnets.

Result:
[                              (json array)
  {                            (json object)
    "address" : "str",         (string) The IP/Subnet of the banned node
    "ban_created" : xxx,       (numeric) The UNIX epoch time the ban was created
    "banned_until" : xxx,      (numeric) The UNIX epoch time the ban expires
    "ban_duration" : xxx,      (numeric) The ban duration, in seconds
    "time_remaining" : xxx     (numeric) The time remaining until the ban expires, in seconds
  },
  ...
]

Examples:
> bitcoin-cli listbanned 
> curl --user myusername --data-binary '{"jsonrpc": "1.0", "id": "curltest", "method": "listbanned", "params": []}' -H 'content-type: text/plain;' http://127.0.0.1:8332/


