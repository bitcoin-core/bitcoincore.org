---
name: getreceivedbyaddress
btcversion: 0.21.0
btcgroup: wallet
permalink: en/doc/0.21.0/rpc/wallet/getreceivedbyaddress/
---

getreceivedbyaddress "address" ( minconf )

Returns the total amount received by the given address in transactions with at least minconf confirmations.

Arguments:
1. address    (string, required) The bitcoin address for transactions.
2. minconf    (numeric, optional, default=1) Only include transactions confirmed at least this many times.

Result:
n    (numeric) The total amount in BTC received at this address.

Examples:

The amount from transactions with at least 1 confirmation
> bitcoin-cli getreceivedbyaddress "bc1q09vm5lfy0j5reeulh4x5752q25uqqvz34hufdl"

The amount including unconfirmed transactions, zero confirmations
> bitcoin-cli getreceivedbyaddress "bc1q09vm5lfy0j5reeulh4x5752q25uqqvz34hufdl" 0

The amount with at least 6 confirmations
> bitcoin-cli getreceivedbyaddress "bc1q09vm5lfy0j5reeulh4x5752q25uqqvz34hufdl" 6

As a JSON-RPC call
> curl --user myusername --data-binary '{"jsonrpc": "1.0", "id": "curltest", "method": "getreceivedbyaddress", "params": ["bc1q09vm5lfy0j5reeulh4x5752q25uqqvz34hufdl", 6]}' -H 'content-type: text/plain;' http://127.0.0.1:8332/


