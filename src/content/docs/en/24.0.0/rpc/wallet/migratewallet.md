---
name: migratewallet
btcversion: 24.0.0
btcgroup: wallet
permalink: en/doc/24.0.0/rpc/wallet/migratewallet/
---

migratewallet

EXPERIMENTAL warning: This call may not work as expected and may be changed in future releases

Migrate the wallet to a descriptor wallet.
A new wallet backup will need to be made.

The migration process will create a backup of the wallet before migrating. This backup
file will be named <wallet name>-<timestamp>.legacy.bak and can be found in the directory
for this wallet. In the event of an incorrect migration, the backup can be restored using restorewallet.
Requires wallet passphrase to be set with walletpassphrase call if wallet is encrypted.

Result:
{                              (json object)
  "wallet_name" : "str",       (string) The name of the primary migrated wallet
  "watchonly_name" : "str",    (string, optional) The name of the migrated wallet containing the watchonly scripts
  "solvables_name" : "str",    (string, optional) The name of the migrated wallet containing solvable but not watched scripts
  "backup_path" : "str"        (string) The location of the backup of the original wallet
}

Examples:
> bitcoin-cli migratewallet 
> curl --user myusername --data-binary '{"jsonrpc": "1.0", "id": "curltest", "method": "migratewallet", "params": []}' -H 'content-type: text/plain;' http://127.0.0.1:8332/


