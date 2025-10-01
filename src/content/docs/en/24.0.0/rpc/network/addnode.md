---
name: addnode
btcversion: 24.0.0
btcgroup: network
permalink: en/doc/24.0.0/rpc/network/addnode/
---

addnode "node" "command"

Attempts to add or remove a node from the addnode list.
Or try a connection to a node once.
Nodes added using addnode (or -connect) are protected from DoS disconnection and are not required to be
full nodes/support SegWit as other outbound peers are (though such peers will not be synced from).
Addnode connections are limited to 8 at a time and are counted separately from the -maxconnections limit.

Arguments:
1. node       (string, required) The node (see getpeerinfo for nodes)
2. command    (string, required) 'add' to add a node to the list, 'remove' to remove a node from the list, 'onetry' to try a connection to the node once

Result:
null    (json null)

Examples:
> bitcoin-cli addnode "192.168.0.6:8333" "onetry"
> curl --user myusername --data-binary '{"jsonrpc": "1.0", "id": "curltest", "method": "addnode", "params": ["192.168.0.6:8333", "onetry"]}' -H 'content-type: text/plain;' http://127.0.0.1:8332/


