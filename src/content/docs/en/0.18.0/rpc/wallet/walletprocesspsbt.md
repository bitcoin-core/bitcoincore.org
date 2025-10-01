---
name: walletprocesspsbt
btcversion: 0.18.0
btcgroup: wallet
permalink: en/doc/0.18.0/rpc/wallet/walletprocesspsbt/
---

walletprocesspsbt "psbt" ( sign "sighashtype" bip32derivs )

Update a PSBT with input information from our wallet and then sign inputs
that we can sign for.

Arguments:
1. psbt           (string, required) The transaction base64 string
2. sign           (boolean, optional, default=true) Also sign the transaction when updating
3. sighashtype    (string, optional, default=ALL) The signature hash type to sign with if not specified by the PSBT. Must be one of
                  "ALL"
                  "NONE"
                  "SINGLE"
                  "ALL|ANYONECANPAY"
                  "NONE|ANYONECANPAY"
                  "SINGLE|ANYONECANPAY"
4. bip32derivs    (boolean, optional, default=false) If true, includes the BIP 32 derivation paths for public keys if we know them

Result:
{
  "psbt" : "value",          (string) The base64-encoded partially signed transaction
  "complete" : true|false,   (boolean) If the transaction has a complete set of signatures
  ]
}

Examples:
> bitcoin-cli walletprocesspsbt "psbt"


