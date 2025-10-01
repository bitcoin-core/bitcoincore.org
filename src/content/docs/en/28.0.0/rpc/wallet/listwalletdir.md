---
name: listwalletdir
btcversion: 28.0.0
btcgroup: wallet
permalink: en/doc/28.0.0/rpc/wallet/listwalletdir/
---

listwalletdir

Returns a list of wallets in the wallet directory.

Result:
{                        (json object)
  "wallets" : [          (json array)
    {                    (json object)
      "name" : "str"     (string) The wallet name
    },
    ...
  ]
}

Examples:
> bitcoin-cli listwalletdir 
> curl --user myusername --data-binary '{"jsonrpc": "2.0", "id": "curltest", "method": "listwalletdir", "params": []}' -H 'content-type: application/json' http://127.0.0.1:8332/


