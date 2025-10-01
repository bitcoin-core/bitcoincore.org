---
name: importwallet
btcversion: 28.0.0
btcgroup: wallet
permalink: en/doc/28.0.0/rpc/wallet/importwallet/
---

importwallet "filename"

Imports keys from a wallet dump file (see dumpwallet). Requires a new wallet backup to include imported keys.
Note: Blockchain and Mempool will be rescanned after a successful import. Use "getwalletinfo" to query the scanning progress.
Note: This command is only compatible with legacy wallets.

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
> curl --user myusername --data-binary '{"jsonrpc": "2.0", "id": "curltest", "method": "importwallet", "params": ["test"]}' -H 'content-type: application/json' http://127.0.0.1:8332/


