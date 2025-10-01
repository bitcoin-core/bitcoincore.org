---
name: importwallet
btcversion: 24.0.0
btcgroup: wallet
permalink: en/doc/24.0.0/rpc/wallet/importwallet/
---

importwallet "filename"

Imports keys from a wallet dump file (see dumpwallet). Requires a new wallet backup to include imported keys.
Note: Blockchain and Mempool will be rescanned after a successful import. Use "getwalletinfo" to query the scanning progress.

Arguments:
1. filename    (string, required) The wallet file

Result:
null    (json null)

Examples:

Dump the wallet
> bitcoin-cli dumpwallet "test"

Import the wallet
> bitcoin-cli importwallet "test"

Import using the json rpc call
> curl --user myusername --data-binary '{"jsonrpc": "1.0", "id": "curltest", "method": "importwallet", "params": ["test"]}' -H 'content-type: text/plain;' http://127.0.0.1:8332/


