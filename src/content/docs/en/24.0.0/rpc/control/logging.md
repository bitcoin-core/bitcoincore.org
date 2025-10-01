---
name: logging
btcversion: 24.0.0
btcgroup: control
permalink: en/doc/24.0.0/rpc/control/logging/
---

logging ( ["include_category",...] ["exclude_category",...] )

Gets and sets the logging configuration.
When called without an argument, returns the list of categories with status that are currently being debug logged or not.
When called with arguments, adds or removes categories from debug logging and return the lists above.
The arguments are evaluated in order "include", "exclude".
If an item is both included and excluded, it will thus end up being excluded.
The valid logging categories are: addrman, bench, blockstorage, cmpctblock, coindb, estimatefee, http, i2p, ipc, leveldb, libevent, mempool, mempoolrej, net, proxy, prune, qt, rand, reindex, rpc, selectcoins, tor, util, validation, walletdb, zmq
In addition, the following are available as category names with special meanings:
  - "all",  "1" : represent all logging categories.
  - "none", "0" : even if other logging categories are specified, ignore all of them.

Arguments:
1. include                    (json array, optional) The categories to add to debug logging
     [
       "include_category",    (string) the valid logging category
       ...
     ]
2. exclude                    (json array, optional) The categories to remove from debug logging
     [
       "exclude_category",    (string) the valid logging category
       ...
     ]

Result:
{                             (json object) keys are the logging categories, and values indicates its status
  "category" : true|false,    (boolean) if being debug logged or not. false:inactive, true:active
  ...
}

Examples:
> bitcoin-cli logging "[\"all\"]" "[\"http\"]"
> curl --user myusername --data-binary '{"jsonrpc": "1.0", "id": "curltest", "method": "logging", "params": [["all"], ["libevent"]]}' -H 'content-type: text/plain;' http://127.0.0.1:8332/


