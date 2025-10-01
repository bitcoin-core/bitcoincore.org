---
name: walletprocesspsbt
btcversion: 24.0.0
btcgroup: wallet
permalink: en/doc/24.0.0/rpc/wallet/walletprocesspsbt/
---

walletprocesspsbt "psbt" ( sign "sighashtype" bip32derivs finalize )

Update a PSBT with input information from our wallet and then sign inputs
that we can sign for.
Requires wallet passphrase to be set with walletpassphrase call if wallet is encrypted.

Arguments:
1. psbt           (string, required) The transaction base64 string
2. sign           (boolean, optional, default=true) Also sign the transaction when updating (requires wallet to be unlocked)
3. sighashtype    (string, optional, default="DEFAULT for Taproot, ALL otherwise") The signature hash type to sign with if not specified by the PSBT. Must be one of
                  "DEFAULT"
                  "ALL"
                  "NONE"
                  "SINGLE"
                  "ALL|ANYONECANPAY"
                  "NONE|ANYONECANPAY"
                  "SINGLE|ANYONECANPAY"
4. bip32derivs    (boolean, optional, default=true) Include BIP 32 derivation paths for public keys if we know them
5. finalize       (boolean, optional, default=true) Also finalize inputs if possible

Result:
{                             (json object)
  "psbt" : "str",             (string) The base64-encoded partially signed transaction
  "complete" : true|false     (boolean) If the transaction has a complete set of signatures
}

Examples:
> bitcoin-cli walletprocesspsbt "psbt"


