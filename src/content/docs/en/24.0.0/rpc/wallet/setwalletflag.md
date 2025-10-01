---
name: setwalletflag
btcversion: 24.0.0
btcgroup: wallet
permalink: en/doc/24.0.0/rpc/wallet/setwalletflag/
---

setwalletflag "flag" ( value )

Change the state of the given wallet flag for a wallet.

Arguments:
1. flag     (string, required) The name of the flag to change. Current available flags: avoid_reuse
2. value    (boolean, optional, default=true) The new state.

Result:
{                               (json object)
  "flag_name" : "str",          (string) The name of the flag that was modified
  "flag_state" : true|false,    (boolean) The new state of the flag
  "warnings" : "str"            (string, optional) Any warnings associated with the change
}

Examples:
> bitcoin-cli setwalletflag avoid_reuse
> curl --user myusername --data-binary '{"jsonrpc": "1.0", "id": "curltest", "method": "setwalletflag", "params": ["avoid_reuse"]}' -H 'content-type: text/plain;' http://127.0.0.1:8332/


