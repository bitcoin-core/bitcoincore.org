---
name: getaddressesbylabel
btcversion: 29.0.0
btcgroup: wallet
permalink: en/doc/29.0.0/rpc/wallet/getaddressesbylabel/
---

getaddressesbylabel "label"

Returns the list of addresses assigned the specified label.

Arguments:
1. label    (string, required) The label.

Result:
{                         (json object) json object with addresses as keys
  "address" : {           (json object) json object with information about address
    "purpose" : "str"     (string) Purpose of address ("send" for sending address, "receive" for receiving address)
  },
  ...
}

Examples:
> bitcoin-cli getaddressesbylabel "tabby"
> curl --user myusername --data-binary '{"jsonrpc": "2.0", "id": "curltest", "method": "getaddressesbylabel", "params": ["tabby"]}' -H 'content-type: application/json' http://127.0.0.1:8332/


