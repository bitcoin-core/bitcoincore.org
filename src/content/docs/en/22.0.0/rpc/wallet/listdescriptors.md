---
name: listdescriptors
btcversion: 22.0.0
btcgroup: wallet
permalink: en/doc/22.0.0/rpc/wallet/listdescriptors/
---

listdescriptors

List descriptors imported into a descriptor-enabled wallet.
Result:
{                                 (json object)
  "wallet_name" : "str",          (string) Name of wallet this operation was performed on
  "descriptors" : [               (json array) Array of descriptor objects
    {                             (json object)
      "desc" : "str",             (string) Descriptor string representation
      "timestamp" : n,            (numeric) The creation time of the descriptor
      "active" : true|false,      (boolean) Activeness flag
      "internal" : true|false,    (boolean, optional) Whether this is an internal or external descriptor; defined only for active descriptors
      "range" : [                 (json array, optional) Defined only for ranged descriptors
        n,                        (numeric) Range start inclusive
        n                         (numeric) Range end inclusive
      ],
      "next" : n                  (numeric, optional) The next index to generate addresses from; defined only for ranged descriptors
    },
    ...
  ]
}

Examples:
> bitcoin-cli listdescriptors 
> curl --user myusername --data-binary '{"jsonrpc": "1.0", "id": "curltest", "method": "listdescriptors", "params": []}' -H 'content-type: text/plain;' http://127.0.0.1:8332/


