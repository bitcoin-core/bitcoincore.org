---
name: restorewallet
btcversion: 25.0.0
btcgroup: wallet
permalink: en/doc/25.0.0/rpc/wallet/restorewallet/
---

restorewallet "wallet_name" "backup_file" ( load_on_startup )

Restore and loads a wallet from backup.

The rescan is significantly faster if a descriptor wallet is restored
and block filters are available (using startup option "-blockfilterindex=1").

Arguments:
1. wallet_name        (string, required) The name that will be applied to the restored wallet
2. backup_file        (string, required) The backup file that will be used to restore the wallet.
3. load_on_startup    (boolean, optional) Save wallet name to persistent settings and load on startup. True to add wallet to startup list, false to remove, null to leave unchanged.

Result:
{                       (json object)
  "name" : "str",       (string) The wallet name if restored successfully.
  "warning" : "str",    (string, optional) Warning messages, if any, related to restoring the wallet. Multiple messages will be delimited by newlines. (DEPRECATED, returned only if config option -deprecatedrpc=walletwarningfield is passed.)
  "warnings" : [        (json array, optional) Warning messages, if any, related to restoring the wallet.
    "str",              (string)
    ...
  ]
}

Examples:
> bitcoin-cli restorewallet "testwallet" "home\backups\backup-file.bak"
> curl --user myusername --data-binary '{"jsonrpc": "1.0", "id": "curltest", "method": "restorewallet", "params": ["testwallet" "home\backups\backup-file.bak"]}' -H 'content-type: text/plain;' http://127.0.0.1:8332/
> bitcoin-cli -named restorewallet wallet_name=testwallet backup_file='home\backups\backup-file.bak"' load_on_startup=true
> curl --user myusername --data-binary '{"jsonrpc": "1.0", "id": "curltest", "method": "restorewallet", "params": {"wallet_name":"testwallet","backup_file":"home\\backups\\backup-file.bak\"","load_on_startup":true}}' -H 'content-type: text/plain;' http://127.0.0.1:8332/


