---
name: setaccount
btcversion: 0.16.2
btcgroup: wallet
permalink: en/doc/0.16.2/rpc/wallet/setaccount/
---

setaccount "address" "account"

DEPRECATED. Sets the account associated with the given address.

Arguments:
1. "address"         (string, required) The bitcoin address to be associated with an account.
2. "account"         (string, required) The account to assign the address to.

Examples:
> bitcoin-cli setaccount "1D1ZrZNe3JUo7ZycKEYQQiQAWd9y54F4XX" "tabby"
> curl --user myusername --data-binary '{"jsonrpc": "1.0", "id":"curltest", "method": "setaccount", "params": ["1D1ZrZNe3JUo7ZycKEYQQiQAWd9y54F4XX", "tabby"] }' -H 'content-type: text/plain;' http://127.0.0.1:8332/


