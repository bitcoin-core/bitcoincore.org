---
name: walletcreatefundedpsbt
btcversion: 0.17.0
btcgroup: wallet
permalink: en/doc/0.17.0/rpc/wallet/walletcreatefundedpsbt/
---

walletcreatefundedpsbt [{"txid":"id","vout":n},...] [{"address":amount},{"data":"hex"},...] ( locktime ) ( replaceable ) ( options bip32derivs )

Creates and funds a transaction in the Partially Signed Transaction format. Inputs will be added if supplied inputs are not enough
Implements the Creator and Updater roles.

Arguments:
1. "inputs"                (array, required) A json array of json objects
     [
       {
         "txid":"id",      (string, required) The transaction id
         "vout":n,         (numeric, required) The output number
         "sequence":n      (numeric, optional) The sequence number
       } 
       ,...
     ]
2. "outputs"               (array, required) a json array with outputs (key-value pairs)
   [
    {
      "address": x.xxx,    (obj, optional) A key-value pair. The key (string) is the bitcoin address, the value (float or string) is the amount in BTC
    },
    {
      "data": "hex"        (obj, optional) A key-value pair. The key must be "data", the value is hex encoded data
    }
    ,...                     More key-value pairs of the above form. For compatibility reasons, a dictionary, which holds the key-value pairs directly, is also
                             accepted as second parameter.
   ]
3. locktime                  (numeric, optional, default=0) Raw locktime. Non-0 value also locktime-activates inputs
                             Allows this transaction to be replaced by a transaction with higher fees. If provided, it is an error if explicit sequence numbers are incompatible.
4. options                 (object, optional)
   {
     "changeAddress"          (string, optional, default pool address) The bitcoin address to receive the change
     "changePosition"         (numeric, optional, default random) The index of the change output
     "change_type"            (string, optional) The output type to use. Only valid if changeAddress is not specified. Options are "legacy", "p2sh-segwit", and "bech32". Default is set by -changetype.
     "includeWatching"        (boolean, optional, default false) Also select inputs which are watch only
     "lockUnspents"           (boolean, optional, default false) Lock selected unspent outputs
     "feeRate"                (numeric, optional, default not set: makes wallet determine the fee) Set a specific fee rate in BTC/kB
     "subtractFeeFromOutputs" (array, optional) A json array of integers.
                              The fee will be equally deducted from the amount of each specified output.
                              The outputs are specified by their zero-based index, before any change output is added.
                              Those recipients will receive less bitcoins than you enter in their corresponding amount field.
                              If no outputs are specified here, the sender pays the fee.
                                  [vout_index,...]
     "replaceable"            (boolean, optional) Marks this transaction as BIP125 replaceable.
                              Allows this transaction to be replaced by a transaction with higher fees
     "conf_target"            (numeric, optional) Confirmation target (in blocks)
     "estimate_mode"          (string, optional, default=UNSET) The fee estimate mode, must be one of:
         "UNSET"
         "ECONOMICAL"
         "CONSERVATIVE"
   }
5. bip32derivs                    (boolean, optional, default=false) If true, includes the BIP 32 derivation paths for public keys if we know them

Result:
{
  "psbt": "value",        (string)  The resulting raw transaction (base64-encoded string)
  "fee":       n,         (numeric) Fee in BTC the resulting transaction pays
  "changepos": n          (numeric) The position of the added change output, or -1
}

Examples:

Create a transaction with no inputs
> bitcoin-cli walletcreatefundedpsbt "[{\"txid\":\"myid\",\"vout\":0}]" "[{\"data\":\"00010203\"}]"


