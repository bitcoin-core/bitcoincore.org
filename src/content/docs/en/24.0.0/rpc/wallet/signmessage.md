---
name: signmessage
btcversion: 24.0.0
btcgroup: wallet
permalink: en/doc/24.0.0/rpc/wallet/signmessage/
---

signmessage "address" "message"

Sign a message with the private key of an address
Requires wallet passphrase to be set with walletpassphrase call if wallet is encrypted.

Arguments:
1. address    (string, required) The bitcoin address to use for the private key.
2. message    (string, required) The message to create a signature of.

Result:
"str"    (string) The signature of the message encoded in base 64

Examples:

Unlock the wallet for 30 seconds
> bitcoin-cli walletpassphrase "mypassphrase" 30

Create the signature
> bitcoin-cli signmessage "1D1ZrZNe3JUo7ZycKEYQQiQAWd9y54F4XX" "my message"

Verify the signature
> bitcoin-cli verifymessage "1D1ZrZNe3JUo7ZycKEYQQiQAWd9y54F4XX" "signature" "my message"

As a JSON-RPC call
> curl --user myusername --data-binary '{"jsonrpc": "1.0", "id": "curltest", "method": "signmessage", "params": ["1D1ZrZNe3JUo7ZycKEYQQiQAWd9y54F4XX", "my message"]}' -H 'content-type: text/plain;' http://127.0.0.1:8332/


