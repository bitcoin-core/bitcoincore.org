---
name: setlabel
btcversion: 24.0.0
btcgroup: wallet
permalink: en/doc/24.0.0/rpc/wallet/setlabel/
---

setlabel "address" "label"

Sets the label associated with the given address.

Arguments:
1. address    (string, required) The bitcoin address to be associated with a label.
2. label      (string, required) The label to assign to the address.

Result:
null    (json null)

Examples:
> bitcoin-cli setlabel "bc1q09vm5lfy0j5reeulh4x5752q25uqqvz34hufdl" "tabby"
> curl --user myusername --data-binary '{"jsonrpc": "1.0", "id": "curltest", "method": "setlabel", "params": ["bc1q09vm5lfy0j5reeulh4x5752q25uqqvz34hufdl", "tabby"]}' -H 'content-type: text/plain;' http://127.0.0.1:8332/


