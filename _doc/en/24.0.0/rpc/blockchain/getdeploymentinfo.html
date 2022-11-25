---
name: getdeploymentinfo
btcversion: 24.0.0
btcgroup: blockchain
permalink: en/doc/24.0.0/rpc/blockchain/getdeploymentinfo/
---

getdeploymentinfo ( "blockhash" )

Returns an object containing various state info regarding deployments of consensus changes.

Arguments:
1. blockhash    (string, optional, default="hash of current chain tip") The block hash at which to query deployment state

Result:
{                                       (json object)
  "hash" : "str",                       (string) requested block hash (or tip)
  "height" : n,                         (numeric) requested block height (or tip)
  "deployments" : {                     (json object)
    "xxxx" : {                          (json object) name of the deployment
      "type" : "str",                   (string) one of "buried", "bip9"
      "height" : n,                     (numeric, optional) height of the first block which the rules are or will be enforced (only for "buried" type, or "bip9" type with "active" status)
      "active" : true|false,            (boolean) true if the rules are enforced for the mempool and the next block
      "bip9" : {                        (json object, optional) status of bip9 softforks (only for "bip9" type)
        "bit" : n,                      (numeric, optional) the bit (0-28) in the block version field used to signal this softfork (only for "started" and "locked_in" status)
        "start_time" : xxx,             (numeric) the minimum median time past of a block at which the bit gains its meaning
        "timeout" : xxx,                (numeric) the median time past of a block at which the deployment is considered failed if not yet locked in
        "min_activation_height" : n,    (numeric) minimum height of blocks for which the rules may be enforced
        "status" : "str",               (string) status of deployment at specified block (one of "defined", "started", "locked_in", "active", "failed")
        "since" : n,                    (numeric) height of the first block to which the status applies
        "status_next" : "str",          (string) status of deployment at the next block
        "statistics" : {                (json object, optional) numeric statistics about signalling for a softfork (only for "started" and "locked_in" status)
          "period" : n,                 (numeric) the length in blocks of the signalling period
          "threshold" : n,              (numeric, optional) the number of blocks with the version bit set required to activate the feature (only for "started" status)
          "elapsed" : n,                (numeric) the number of blocks elapsed since the beginning of the current period
          "count" : n,                  (numeric) the number of blocks with the version bit set in the current period
          "possible" : true|false       (boolean, optional) returns false if there are not enough blocks left in this period to pass activation threshold (only for "started" status)
        },
        "signalling" : "str"            (string, optional) indicates blocks that signalled with a # and blocks that did not with a -
      }
    },
    ...
  }
}

Examples:
> bitcoin-cli getdeploymentinfo 
> curl --user myusername --data-binary '{"jsonrpc": "1.0", "id": "curltest", "method": "getdeploymentinfo", "params": []}' -H 'content-type: text/plain;' http://127.0.0.1:8332/


