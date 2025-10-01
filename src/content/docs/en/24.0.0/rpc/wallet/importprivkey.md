---
name: importprivkey
btcversion: 24.0.0
btcgroup: wallet
permalink: en/doc/24.0.0/rpc/wallet/importprivkey/
---

importprivkey "privkey" ( "label" rescan )

Adds a private key (as returned by dumpprivkey) to your wallet. Requires a new wallet backup.
Hint: use importmulti to import more than one private key.

Note: This call can take over an hour to complete if rescan is true, during that time, other rpc calls
may report that the imported key exists but related transactions are still missing, leading to temporarily incorrect/bogus balances and unspent outputs until rescan completes.
The rescan parameter can be set to false if the key was never used to create transactions. If it is set to false,
but the key was used to create transactions, rescanblockchain needs to be called with the appropriate block range.
Note: Use "getwalletinfo" to query the scanning progress.

Arguments:
1. privkey    (string, required) The private key (see dumpprivkey)
2. label      (string, optional, default=current label if address exists, otherwise "") An optional label
3. rescan     (boolean, optional, default=true) Scan the chain and mempool for wallet transactions.

Result:
null    (json null)

Examples:

Dump a private key
> bitcoin-cli dumpprivkey "myaddress"

Import the private key with rescan
> bitcoin-cli importprivkey "mykey"

Import using a label and without rescan
> bitcoin-cli importprivkey "mykey" "testing" false

Import using default blank label and without rescan
> bitcoin-cli importprivkey "mykey" "" false

As a JSON-RPC call
> curl --user myusername --data-binary '{"jsonrpc": "1.0", "id": "curltest", "method": "importprivkey", "params": ["mykey", "testing", false]}' -H 'content-type: text/plain;' http://127.0.0.1:8332/


