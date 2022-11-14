---
name: importpubkey
btcversion: 24.0.0
btcgroup: wallet
permalink: en/doc/24.0.0/rpc/wallet/importpubkey/
---

importpubkey "pubkey" ( "label" rescan )

Adds a public key (in hex) that can be watched as if it were in your wallet but cannot be used to spend. Requires a new wallet backup.
Hint: use importmulti to import more than one public key.

Note: This call can take over an hour to complete if rescan is true, during that time, other rpc calls
may report that the imported pubkey exists but related transactions are still missing, leading to temporarily incorrect/bogus balances and unspent outputs until rescan completes.
The rescan parameter can be set to false if the key was never used to create transactions. If it is set to false,
but the key was used to create transactions, rescanblockchain needs to be called with the appropriate block range.
Note: Use "getwalletinfo" to query the scanning progress.

Arguments:
1. pubkey    (string, required) The hex-encoded public key
2. label     (string, optional, default="") An optional label
3. rescan    (boolean, optional, default=true) Scan the chain and mempool for wallet transactions.

Result:
null    (json null)

Examples:

Import a public key with rescan
> bitcoin-cli importpubkey "mypubkey"

Import using a label without rescan
> bitcoin-cli importpubkey "mypubkey" "testing" false

As a JSON-RPC call
> curl --user myusername --data-binary '{"jsonrpc": "1.0", "id": "curltest", "method": "importpubkey", "params": ["mypubkey", "testing", false]}' -H 'content-type: text/plain;' http://127.0.0.1:8332/


