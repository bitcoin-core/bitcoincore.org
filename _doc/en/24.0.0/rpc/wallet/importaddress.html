---
name: importaddress
btcversion: 24.0.0
btcgroup: wallet
permalink: en/doc/24.0.0/rpc/wallet/importaddress/
---

importaddress "address" ( "label" rescan p2sh )

Adds an address or script (in hex) that can be watched as if it were in your wallet but cannot be used to spend. Requires a new wallet backup.

Note: This call can take over an hour to complete if rescan is true, during that time, other rpc calls
may report that the imported address exists but related transactions are still missing, leading to temporarily incorrect/bogus balances and unspent outputs until rescan completes.
The rescan parameter can be set to false if the key was never used to create transactions. If it is set to false,
but the key was used to create transactions, rescanblockchain needs to be called with the appropriate block range.
If you have the full public key, you should call importpubkey instead of this.
Hint: use importmulti to import more than one address.

Note: If you import a non-standard raw script in hex form, outputs sending to it will be treated
as change, and not show up in many RPCs.
Note: Use "getwalletinfo" to query the scanning progress.
Note: This command is only compatible with legacy wallets. Use "importdescriptors" with "addr(X)" for descriptor wallets.

Arguments:
1. address    (string, required) The Bitcoin address (or hex-encoded script)
2. label      (string, optional, default="") An optional label
3. rescan     (boolean, optional, default=true) Scan the chain and mempool for wallet transactions.
4. p2sh       (boolean, optional, default=false) Add the P2SH version of the script as well

Result:
null    (json null)

Examples:

Import an address with rescan
> bitcoin-cli importaddress "myaddress"

Import using a label without rescan
> bitcoin-cli importaddress "myaddress" "testing" false

As a JSON-RPC call
> curl --user myusername --data-binary '{"jsonrpc": "1.0", "id": "curltest", "method": "importaddress", "params": ["myaddress", "testing", false]}' -H 'content-type: text/plain;' http://127.0.0.1:8332/


