---
name: listdescriptors
btcversion: 24.0.0
btcgroup: wallet
permalink: en/doc/24.0.0/rpc/wallet/listdescriptors/
---

listdescriptors ( private )

List descriptors imported into a descriptor-enabled wallet.

Arguments:
1. private    (boolean, optional, default=false) Show private descriptors.

Result:
{                                 (json object)
  "wallet_name" : "str",          (string) Name of wallet this operation was performed on
  "descriptors" : [               (json array) Array of descriptor objects
    {                             (json object)
      "desc" : "str",             (string) Descriptor string representation
      "timestamp" : n,            (numeric) The creation time of the descriptor
      "active" : true|false,      (boolean) Whether this descriptor is currently used to generate new addresses
      "internal" : true|false,    (boolean, optional) True if this descriptor is used to generate change addresses. False if this descriptor is used to generate receiving addresses; defined only for active descriptors
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
> bitcoin-cli listdescriptors true
> curl --user myusername --data-binary '{"jsonrpc": "1.0", "id": "curltest", "method": "listdescriptors", "params": [true]}' -H 'content-type: text/plain;' http://127.0.0.1:8332/


