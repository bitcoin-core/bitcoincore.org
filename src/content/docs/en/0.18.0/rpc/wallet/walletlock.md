---
name: walletlock
btcversion: 0.18.0
btcgroup: wallet
permalink: en/doc/0.18.0/rpc/wallet/walletlock/
---

walletlock

Removes the wallet encryption key from memory, locking the wallet.
After calling this method, you will need to call walletpassphrase again
before being able to call any methods which require the wallet to be unlocked.

Examples:

Set the passphrase for 2 minutes to perform a transaction
> bitcoin-cli walletpassphrase "my pass phrase" 120

Perform a send (requires passphrase set)
> bitcoin-cli sendtoaddress "1M72Sfpbz1BPpXFHz9m3CdqATR44Jvaydd" 1.0

Clear the passphrase since we are done before 2 minutes is up
> bitcoin-cli walletlock 

As a JSON-RPC call
> curl --user myusername --data-binary '{"jsonrpc": "1.0", "id":"curltest", "method": "walletlock", "params": [] }' -H 'content-type: text/plain;' http://127.0.0.1:8332/


