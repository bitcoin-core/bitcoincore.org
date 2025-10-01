---
name: rescanblockchain
btcversion: 24.0.0
btcgroup: wallet
permalink: en/doc/24.0.0/rpc/wallet/rescanblockchain/
---

rescanblockchain ( start_height stop_height )

Rescan the local blockchain for wallet related transactions.
Note: Use "getwalletinfo" to query the scanning progress.

Arguments:
1. start_height    (numeric, optional, default=0) block height where the rescan should start
2. stop_height     (numeric, optional) the last block height that should be scanned. If none is provided it will rescan up to the tip at return time of this call.

Result:
{                        (json object)
  "start_height" : n,    (numeric) The block height where the rescan started (the requested height or 0)
  "stop_height" : n      (numeric) The height of the last rescanned block. May be null in rare cases if there was a reorg and the call didn't scan any blocks because they were already scanned in the background.
}

Examples:
> bitcoin-cli rescanblockchain 100000 120000
> curl --user myusername --data-binary '{"jsonrpc": "1.0", "id": "curltest", "method": "rescanblockchain", "params": [100000, 120000]}' -H 'content-type: text/plain;' http://127.0.0.1:8332/


