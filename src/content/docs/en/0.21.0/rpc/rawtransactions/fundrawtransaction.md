---
name: fundrawtransaction
btcversion: 0.21.0
btcgroup: rawtransactions
permalink: en/doc/0.21.0/rpc/rawtransactions/fundrawtransaction/
---

fundrawtransaction "hexstring" ( options iswitness )

If the transaction has no inputs, they will be automatically selected to meet its out value.
It will add at most one change output to the outputs.
No existing outputs will be modified unless "subtractFeeFromOutputs" is specified.
Note that inputs which were signed may need to be resigned after completion since in/outputs have been added.
The inputs added will not be signed, use signrawtransactionwithkey
 or signrawtransactionwithwallet for that.
Note that all existing inputs must have their previous output transaction be in the wallet.
Note that all inputs selected must be of standard form and P2SH scripts must be
in the wallet using importaddress or addmultisigaddress (to calculate fees).
You can see whether this is the case by checking the "solvable" field in the listunspent output.
Only pay-to-pubkey, multisig, and P2SH versions thereof are currently supported for watch-only

Arguments:
1. hexstring                          (string, required) The hex string of the raw transaction
2. options                            (json object, optional) for backward compatibility: passing in a true instead of an object will result in {"includeWatching":true}
     {
       "add_inputs": bool,            (boolean, optional, default=true) For a transaction with existing inputs, automatically include more if they are not enough.
       "changeAddress": "str",        (string, optional, default=pool address) The bitcoin address to receive the change
       "changePosition": n,           (numeric, optional, default=random) The index of the change output
       "change_type": "str",          (string, optional, default=set by -changetype) The output type to use. Only valid if changeAddress is not specified. Options are "legacy", "p2sh-segwit", and "bech32".
       "includeWatching": bool,       (boolean, optional, default=true for watch-only wallets, otherwise false) Also select inputs which are watch only.
                                      Only solvable inputs can be used. Watch-only destinations are solvable if the public key and/or output script was imported,
                                      e.g. with 'importpubkey' or 'importmulti' with the 'pubkeys' or 'desc' field.
       "lockUnspents": bool,          (boolean, optional, default=false) Lock selected unspent outputs
       "fee_rate": amount,            (numeric or string, optional, default=not set, fall back to wallet fee estimation) Specify a fee rate in sat/vB.
       "feeRate": amount,             (numeric or string, optional, default=not set, fall back to wallet fee estimation) Specify a fee rate in BTC/kvB.
       "subtractFeeFromOutputs": [    (json array, optional, default=empty array) The integers.
                                      The fee will be equally deducted from the amount of each specified output.
                                      Those recipients will receive less bitcoins than you enter in their corresponding amount field.
                                      If no outputs are specified here, the sender pays the fee.
         vout_index,                  (numeric) The zero-based output index, before a change output is added.
         ...
       ],
       "replaceable": bool,           (boolean, optional, default=wallet default) Marks this transaction as BIP125 replaceable.
                                      Allows this transaction to be replaced by a transaction with higher fees
       "conf_target": n,              (numeric, optional, default=wallet -txconfirmtarget) Confirmation target in blocks
       "estimate_mode": "str",        (string, optional, default=unset) The fee estimate mode, must be one of (case insensitive):
                                      "unset"
                                      "economical"
                                      "conservative"
     }
3. iswitness                          (boolean, optional, default=depends on heuristic tests) Whether the transaction hex is a serialized witness transaction.
                                      If iswitness is not present, heuristic tests will be used in decoding.
                                      If true, only witness deserialization will be tried.
                                      If false, only non-witness deserialization will be tried.
                                      This boolean should reflect whether the transaction has inputs
                                      (e.g. fully valid, or on-chain transactions), if known by the caller.

Result:
{                     (json object)
  "hex" : "hex",      (string) The resulting raw transaction (hex-encoded string)
  "fee" : n,          (numeric) Fee in BTC the resulting transaction pays
  "changepos" : n     (numeric) The position of the added change output, or -1
}

Examples:

Create a transaction with no inputs
> bitcoin-cli createrawtransaction "[]" "{\"myaddress\":0.01}"

Add sufficient unsigned inputs to meet the output value
> bitcoin-cli fundrawtransaction "rawtransactionhex"

Sign the transaction
> bitcoin-cli signrawtransactionwithwallet "fundedtransactionhex"

Send the transaction
> bitcoin-cli sendrawtransaction "signedtransactionhex"


