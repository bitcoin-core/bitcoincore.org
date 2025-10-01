---
name: importdescriptors
btcversion: 24.0.0
btcgroup: wallet
permalink: en/doc/24.0.0/rpc/wallet/importdescriptors/
---

importdescriptors "requests"

Import descriptors. This will trigger a rescan of the blockchain based on the earliest timestamp of all descriptors being imported. Requires a new wallet backup.

Note: This call can take over an hour to complete if using an early timestamp; during that time, other rpc calls
may report that the imported keys, addresses or scripts exist but related transactions are still missing.

Arguments:
1. requests                                 (json array, required) Data to be imported
     [
       {                                    (json object)
         "desc": "str",                     (string, required) Descriptor to import.
         "active": bool,                    (boolean, optional, default=false) Set this descriptor to be the active descriptor for the corresponding output type/externality
         "range": n or [n,n],               (numeric or array) If a ranged descriptor is used, this specifies the end or the range (in the form [begin,end]) to import
         "next_index": n,                   (numeric) If a ranged descriptor is set to active, this specifies the next index to generate addresses from
         "timestamp": timestamp | "now",    (integer / string, required) Time from which to start rescanning the blockchain for this descriptor, in UNIX epoch time
                                            Use the string "now" to substitute the current synced blockchain time.
                                            "now" can be specified to bypass scanning, for outputs which are known to never have been used, and
                                            0 can be specified to scan the entire blockchain. Blocks up to 2 hours before the earliest timestamp
                                            of all descriptors being imported will be scanned as well as the mempool.
         "internal": bool,                  (boolean, optional, default=false) Whether matching outputs should be treated as not incoming payments (e.g. change)
         "label": "str",                    (string, optional, default="") Label to assign to the address, only allowed with internal=false. Disabled for ranged descriptors
       },
       ...
     ]

Result:
[                              (json array) Response is an array with the same size as the input that has the execution result
  {                            (json object)
    "success" : true|false,    (boolean)
    "warnings" : [             (json array, optional)
      "str",                   (string)
      ...
    ],
    "error" : {                (json object, optional)
      ...                      JSONRPC error
    }
  },
  ...
]

Examples:
> bitcoin-cli importdescriptors '[{ "desc": "<my descriptor>", "timestamp":1455191478, "internal": true }, { "desc": "<my descriptor 2>", "label": "example 2", "timestamp": 1455191480 }]'
> bitcoin-cli importdescriptors '[{ "desc": "<my descriptor>", "timestamp":1455191478, "active": true, "range": [0,100], "label": "<my bech32 wallet>" }]'


