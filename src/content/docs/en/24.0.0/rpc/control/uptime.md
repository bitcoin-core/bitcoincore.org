---
name: uptime
btcversion: 24.0.0
btcgroup: control
permalink: en/doc/24.0.0/rpc/control/uptime/
---

uptime

Returns the total uptime of the server.

Result:
n    (numeric) The number of seconds that the server has been running

Examples:
> bitcoin-cli uptime 
> curl --user myusername --data-binary '{"jsonrpc": "1.0", "id": "curltest", "method": "uptime", "params": []}' -H 'content-type: text/plain;' http://127.0.0.1:8332/


