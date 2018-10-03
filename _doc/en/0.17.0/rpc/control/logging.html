---
name: logging
btcversion: 0.17.0
btcgroup: control
permalink: en/doc/0.17.0/rpc/control/logging/
---

logging ( <include> <exclude> )
Gets and sets the logging configuration.
When called without an argument, returns the list of categories with status that are currently being debug logged or not.
When called with arguments, adds or removes categories from debug logging and return the lists above.
The arguments are evaluated in order "include", "exclude".
If an item is both included and excluded, it will thus end up being excluded.
The valid logging categories are: net, tor, mempool, http, bench, zmq, db, rpc, estimatefee, addrman, selectcoins, reindex, cmpctblock, rand, prune, proxy, mempoolrej, libevent, coindb, qt, leveldb
In addition, the following are available as category names with special meanings:
  - "all",  "1" : represent all logging categories.
  - "none", "0" : even if other logging categories are specified, ignore all of them.

Arguments:
1. "include"        (array of strings, optional) A json array of categories to add debug logging
     [
       "category"   (string) the valid logging category
       ,...
     ]
2. "exclude"        (array of strings, optional) A json array of categories to remove debug logging
     [
       "category"   (string) the valid logging category
       ,...
     ]

Result:
{                   (json object where keys are the logging categories, and values indicates its status
  "category": 0|1,  (numeric) if being debug logged or not. 0:inactive, 1:active
  ...
}

Examples:
> bitcoin-cli logging "[\"all\"]" "[\"http\"]"
> curl --user myusername --data-binary '{"jsonrpc": "1.0", "id":"curltest", "method": "logging", "params": [["all"], "[libevent]"] }' -H 'content-type: text/plain;' http://127.0.0.1:8332/


