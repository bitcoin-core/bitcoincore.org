---
name: createwallet
btcversion: 24.0.0
btcgroup: wallet
permalink: en/doc/24.0.0/rpc/wallet/createwallet/
---

createwallet "wallet_name" ( disable_private_keys blank "passphrase" avoid_reuse descriptors load_on_startup external_signer )

Creates and loads a new wallet.

Arguments:
1. wallet_name             (string, required) The name for the new wallet. If this is a path, the wallet will be created at the path location.
2. disable_private_keys    (boolean, optional, default=false) Disable the possibility of private keys (only watchonlys are possible in this mode).
3. blank                   (boolean, optional, default=false) Create a blank wallet. A blank wallet has no keys or HD seed. One can be set using sethdseed.
4. passphrase              (string, optional) Encrypt the wallet with this passphrase.
5. avoid_reuse             (boolean, optional, default=false) Keep track of coin reuse, and treat dirty and clean coins differently with privacy considerations in mind.
6. descriptors             (boolean, optional, default=true) Create a native descriptor wallet. The wallet will use descriptors internally to handle address creation. Setting to "false" will create a legacy wallet; however, the legacy wallet type is being deprecated and support for creating and opening legacy wallets will be removed in the future.
7. load_on_startup         (boolean, optional) Save wallet name to persistent settings and load on startup. True to add wallet to startup list, false to remove, null to leave unchanged.
8. external_signer         (boolean, optional, default=false) Use an external signer such as a hardware wallet. Requires -signer to be configured. Wallet creation will fail if keys cannot be fetched. Requires disable_private_keys and descriptors set to true.

Result:
{                       (json object)
  "name" : "str",       (string) The wallet name if created successfully. If the wallet was created using a full path, the wallet_name will be the full path.
  "warning" : "str"     (string) Warning message if wallet was not loaded cleanly.
}

Examples:
> bitcoin-cli createwallet "testwallet"
> curl --user myusername --data-binary '{"jsonrpc": "1.0", "id": "curltest", "method": "createwallet", "params": ["testwallet"]}' -H 'content-type: text/plain;' http://127.0.0.1:8332/
> bitcoin-cli -named createwallet wallet_name=descriptors avoid_reuse=true descriptors=true load_on_startup=true
> curl --user myusername --data-binary '{"jsonrpc": "1.0", "id": "curltest", "method": "createwallet", "params": {"wallet_name":"descriptors","avoid_reuse":true,"descriptors":true,"load_on_startup":true}}' -H 'content-type: text/plain;' http://127.0.0.1:8332/


