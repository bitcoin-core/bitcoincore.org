---
name: getbalance
btcversion: 0.16.0
btcgroup: wallet
permalink: en/doc/0.16.0/rpc/wallet/getbalance/
---

getbalance ( "account" minconf include_watchonly )

If account is not specified, returns the server's total available balance.
The available balance is what the wallet considers currently spendable, and is
thus affected by options which limit spendability such as -spendzeroconfchange.
If account is specified (DEPRECATED), returns the balance in the account.
Note that the account "" is not the same as leaving the parameter out.
The server total may be different to the balance in the default "" account.

Arguments:
1. "account"         (string, optional) DEPRECATED. The account string may be given as a
                     specific account name to find the balance associated with wallet keys in
                     a named account, or as the empty string ("") to find the balance
                     associated with wallet keys not in any named account, or as "*" to find
                     the balance associated with all wallet keys regardless of account.
                     When this option is specified, it calculates the balance in a different
                     way than when it is not specified, and which can count spends twice when
                     there are conflicting pending transactions (such as those created by
                     the bumpfee command), temporarily resulting in low or even negative
                     balances. In general, account balance calculation is not considered
                     reliable and has resulted in confusing outcomes, so it is recommended to
                     avoid passing this argument.
2. minconf           (numeric, optional, default=1) Only include transactions confirmed at least this many times.
3. include_watchonly (bool, optional, default=false) Also include balance in watch-only addresses (see 'importaddress')

Result:
amount              (numeric) The total amount in BTC received for this account.

Examples:

The total amount in the wallet with 1 or more confirmations
> bitcoin-cli getbalance 

The total amount in the wallet at least 6 blocks confirmed
> bitcoin-cli getbalance "*" 6

As a json rpc call
> curl --user myusername --data-binary '{"jsonrpc": "1.0", "id":"curltest", "method": "getbalance", "params": ["*", 6] }' -H 'content-type: text/plain;' http://127.0.0.1:8332/


