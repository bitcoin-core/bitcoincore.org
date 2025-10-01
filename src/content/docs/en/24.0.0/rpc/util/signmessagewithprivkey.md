---
name: signmessagewithprivkey
btcversion: 24.0.0
btcgroup: util
permalink: en/doc/24.0.0/rpc/util/signmessagewithprivkey/
---

signmessagewithprivkey "privkey" "message"

Sign a message with the private key of an address

Arguments:
1. privkey    (string, required) The private key to sign the message with.
2. message    (string, required) The message to create a signature of.

Result:
"str"    (string) The signature of the message encoded in base 64

Examples:

Create the signature
> bitcoin-cli signmessagewithprivkey "privkey" "my message"

Verify the signature
> bitcoin-cli verifymessage "1D1ZrZNe3JUo7ZycKEYQQiQAWd9y54F4XX" "signature" "my message"

As a JSON-RPC call
> curl --user myusername --data-binary '{"jsonrpc": "1.0", "id": "curltest", "method": "signmessagewithprivkey", "params": ["privkey", "my message"]}' -H 'content-type: text/plain;' http://127.0.0.1:8332/


