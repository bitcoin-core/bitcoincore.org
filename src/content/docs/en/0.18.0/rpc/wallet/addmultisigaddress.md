---
name: addmultisigaddress
btcversion: 0.18.0
btcgroup: wallet
permalink: en/doc/0.18.0/rpc/wallet/addmultisigaddress/
---

addmultisigaddress nrequired ["key",...] ( "label" "address_type" )

Add a nrequired-to-sign multisignature address to the wallet. Requires a new wallet backup.
Each key is a Bitcoin address or hex-encoded public key.
This functionality is only intended for use with non-watchonly addresses.
See `importaddress` for watchonly p2sh address support.
If 'label' is specified, assign address to that label.

Arguments:
1. nrequired       (numeric, required) The number of required signatures out of the n keys or addresses.
2. keys            (json array, required) A json array of bitcoin addresses or hex-encoded public keys
     [
       "key",      (string) bitcoin address or hex-encoded public key
       ...
     ]
3. label           (string, optional) A label to assign the addresses to.
4. address_type    (string, optional, default=set by -addresstype) The address type to use. Options are "legacy", "p2sh-segwit", and "bech32".

Result:
{
  "address":"multisigaddress",    (string) The value of the new multisig address.
  "redeemScript":"script"         (string) The string value of the hex-encoded redemption script.
}

Examples:

Add a multisig address from 2 addresses
> bitcoin-cli addmultisigaddress 2 "[\"16sSauSf5pF2UkUwvKGq4qjNRzBZYqgEL5\",\"171sgjn4YtPu27adkKGrdDwzRTxnRkBfKV\"]"

As a JSON-RPC call
> curl --user myusername --data-binary '{"jsonrpc": "1.0", "id":"curltest", "method": "addmultisigaddress", "params": [2, "[\"16sSauSf5pF2UkUwvKGq4qjNRzBZYqgEL5\",\"171sgjn4YtPu27adkKGrdDwzRTxnRkBfKV\"]"] }' -H 'content-type: text/plain;' http://127.0.0.1:8332/


