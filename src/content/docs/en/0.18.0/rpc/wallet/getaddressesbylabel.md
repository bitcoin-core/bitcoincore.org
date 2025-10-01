---
name: getaddressesbylabel
btcversion: 0.18.0
btcgroup: wallet
permalink: en/doc/0.18.0/rpc/wallet/getaddressesbylabel/
---

getaddressesbylabel "label"

Returns the list of addresses assigned the specified label.

Arguments:
1. label    (string, required) The label.

Result:
{ (json object with addresses as keys)
  "address": { (json object with information about address)
    "purpose": "string" (string)  Purpose of address ("send" for sending address, "receive" for receiving address)
  },...
}

Examples:
> bitcoin-cli getaddressesbylabel "tabby"
> curl --user myusername --data-binary '{"jsonrpc": "1.0", "id":"curltest", "method": "getaddressesbylabel", "params": ["tabby"] }' -H 'content-type: text/plain;' http://127.0.0.1:8332/


