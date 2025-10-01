---
name: importaddress
btcversion: 0.16.1
btcgroup: wallet
permalink: en/doc/0.16.1/rpc/wallet/importaddress/
---

importaddress "address" ( "label" rescan p2sh )

Adds a script (in hex) or address that can be watched as if it were in your wallet but cannot be used to spend. Requires a new wallet backup.

Arguments:
1. "script"           (string, required) The hex-encoded script (or address)
2. "label"            (string, optional, default="") An optional label
3. rescan               (boolean, optional, default=true) Rescan the wallet for transactions
4. p2sh                 (boolean, optional, default=false) Add the P2SH version of the script as well

Note: This call can take minutes to complete if rescan is true, during that time, other rpc calls
may report that the imported address exists but related transactions are still missing, leading to temporarily incorrect/bogus balances and unspent outputs until rescan completes.
If you have the full public key, you should call importpubkey instead of this.

Note: If you import a non-standard raw script in hex form, outputs sending to it will be treated
as change, and not show up in many RPCs.

Examples:

Import a script with rescan
> bitcoin-cli importaddress "myscript"

Import using a label without rescan
> bitcoin-cli importaddress "myscript" "testing" false

As a JSON-RPC call
> curl --user myusername --data-binary '{"jsonrpc": "1.0", "id":"curltest", "method": "importaddress", "params": ["myscript", "testing", false] }' -H 'content-type: text/plain;' http://127.0.0.1:8332/


