---
name: enumeratesigners
btcversion: 22.0.0
btcgroup: signer
permalink: en/doc/22.0.0/rpc/signer/enumeratesigners/
---

enumeratesigners
Returns a list of external signers from -signer.
Result:
{                  (json object)
  "signers" : [    (json array)
    "hex",         (string) Master key fingerprint
    "str",         (string) Device name
    ...
  ]
}

Examples:
> bitcoin-cli enumeratesigners 
> curl --user myusername --data-binary '{"jsonrpc": "1.0", "id": "curltest", "method": "enumeratesigners", "params": []}' -H 'content-type: text/plain;' http://127.0.0.1:8332/


