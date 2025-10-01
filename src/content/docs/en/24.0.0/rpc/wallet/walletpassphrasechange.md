---
name: walletpassphrasechange
btcversion: 24.0.0
btcgroup: wallet
permalink: en/doc/24.0.0/rpc/wallet/walletpassphrasechange/
---

walletpassphrasechange "oldpassphrase" "newpassphrase"

Changes the wallet passphrase from 'oldpassphrase' to 'newpassphrase'.

Arguments:
1. oldpassphrase    (string, required) The current passphrase
2. newpassphrase    (string, required) The new passphrase

Result:
null    (json null)

Examples:
> bitcoin-cli walletpassphrasechange "old one" "new one"
> curl --user myusername --data-binary '{"jsonrpc": "1.0", "id": "curltest", "method": "walletpassphrasechange", "params": ["old one", "new one"]}' -H 'content-type: text/plain;' http://127.0.0.1:8332/


