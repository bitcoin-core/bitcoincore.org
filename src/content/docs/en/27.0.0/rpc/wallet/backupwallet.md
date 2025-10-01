---
name: backupwallet
btcversion: 27.0.0
btcgroup: wallet
permalink: en/doc/27.0.0/rpc/wallet/backupwallet/
---

backupwallet "destination"

Safely copies the current wallet file to the specified destination, which can either be a directory or a path with a filename.

Arguments:
1. destination    (string, required) The destination directory or file

Result:
null    (json null)

Examples:
> bitcoin-cli backupwallet "backup.dat"
> curl --user myusername --data-binary '{"jsonrpc": "1.0", "id": "curltest", "method": "backupwallet", "params": ["backup.dat"]}' -H 'content-type: text/plain;' http://127.0.0.1:8332/


