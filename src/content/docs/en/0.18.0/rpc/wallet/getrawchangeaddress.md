---
name: getrawchangeaddress
btcversion: 0.18.0
btcgroup: wallet
permalink: en/doc/0.18.0/rpc/wallet/getrawchangeaddress/
---

getrawchangeaddress ( "address_type" )

Returns a new Bitcoin address, for receiving change.
This is for use with raw transactions, NOT normal use.

Arguments:
1. address_type    (string, optional, default=set by -changetype) The address type to use. Options are "legacy", "p2sh-segwit", and "bech32".

Result:
"address"    (string) The address

Examples:
> bitcoin-cli getrawchangeaddress 
> curl --user myusername --data-binary '{"jsonrpc": "1.0", "id":"curltest", "method": "getrawchangeaddress", "params": [] }' -H 'content-type: text/plain;' http://127.0.0.1:8332/


