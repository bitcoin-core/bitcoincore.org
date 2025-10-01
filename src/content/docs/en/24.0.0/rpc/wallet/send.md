---
name: send
btcversion: 24.0.0
btcgroup: wallet
permalink: en/doc/24.0.0/rpc/wallet/send/
---

send [{"address":amount,...},{"data":"hex"},...] ( conf_target "estimate_mode" fee_rate options )

EXPERIMENTAL warning: this call may be changed in future releases.

Send a transaction.

Arguments:
1. outputs                               (json array, required) The outputs (key-value pairs), where none of the keys are duplicated.
                                         That is, each address can only appear once and there can only be one 'data' object.
                                         For convenience, a dictionary, which holds the key-value pairs directly, is also accepted.
     [
       {                                 (json object)
         "address": amount,              (numeric or string, required) A key-value pair. The key (string) is the bitcoin address, the value (float or string) is the amount in BTC
         ...
       },
       {                                 (json object)
         "data": "hex",                  (string, required) A key-value pair. The key must be "data", the value is hex-encoded data
       },
       ...
     ]
2. conf_target                           (numeric, optional, default=wallet -txconfirmtarget) Confirmation target in blocks
3. estimate_mode                         (string, optional, default="unset") The fee estimate mode, must be one of (case insensitive):
                                         "unset"
                                         "economical"
                                         "conservative"
4. fee_rate                              (numeric or string, optional, default=not set, fall back to wallet fee estimation) Specify a fee rate in sat/vB.
5. options                               (json object, optional)
     {
       "add_inputs": bool,               (boolean, optional, default=false when "inputs" are specified, true otherwise) Automatically include coins from the wallet to cover the target amount.
                                         
       "include_unsafe": bool,           (boolean, optional, default=false) Include inputs that are not safe to spend (unconfirmed transactions from outside keys and unconfirmed replacement transactions).
                                         Warning: the resulting transaction may become invalid if one of the unsafe inputs disappears.
                                         If that happens, you will need to fund the transaction with different inputs and republish it.
       "add_to_wallet": bool,            (boolean, optional, default=true) When false, returns a serialized transaction which will not be added to the wallet or broadcast
       "change_address": "str",          (string, optional, default=automatic) The bitcoin address to receive the change
       "change_position": n,             (numeric, optional, default=random) The index of the change output
       "change_type": "str",             (string, optional, default=set by -changetype) The output type to use. Only valid if change_address is not specified. Options are "legacy", "p2sh-segwit", "bech32" and "bech32m".
       "fee_rate": amount,               (numeric or string, optional, default=not set, fall back to wallet fee estimation) Specify a fee rate in sat/vB.
       "include_watching": bool,         (boolean, optional, default=true for watch-only wallets, otherwise false) Also select inputs which are watch only.
                                         Only solvable inputs can be used. Watch-only destinations are solvable if the public key and/or output script was imported,
                                         e.g. with 'importpubkey' or 'importmulti' with the 'pubkeys' or 'desc' field.
       "inputs": [                       (json array, optional, default=[]) Specify inputs instead of adding them automatically. A JSON array of JSON objects
         "txid",                         (string, required) The transaction id
         vout,                           (numeric, required) The output number
         sequence,                       (numeric, required) The sequence number
         weight,                         (numeric, optional, default=Calculated from wallet and solving data) The maximum weight for this input, including the weight of the outpoint and sequence number. Note that signature sizes are not guaranteed to be consistent, so the maximum DER signatures size of 73 bytes should be used when considering ECDSA signatures.Remember to convert serialized sizes to weight units when necessary.
         ...
       ],
       "locktime": n,                    (numeric, optional, default=0) Raw locktime. Non-0 value also locktime-activates inputs
       "lock_unspents": bool,            (boolean, optional, default=false) Lock selected unspent outputs
       "psbt": bool,                     (boolean, optional, default=automatic) Always return a PSBT, implies add_to_wallet=false.
       "subtract_fee_from_outputs": [    (json array, optional, default=[]) Outputs to subtract the fee from, specified as integer indices.
                                         The fee will be equally deducted from the amount of each specified output.
                                         Those recipients will receive less bitcoins than you enter in their corresponding amount field.
                                         If no outputs are specified here, the sender pays the fee.
         vout_index,                     (numeric) The zero-based output index, before a change output is added.
         ...
       ],
       "conf_target": n,                 (numeric, optional, default=wallet -txconfirmtarget) Confirmation target in blocks
       "estimate_mode": "str",           (string, optional, default="unset") The fee estimate mode, must be one of (case insensitive):
                                         "unset"
                                         "economical"
                                         "conservative"
       "replaceable": bool,              (boolean, optional, default=wallet default) Marks this transaction as BIP125-replaceable.
                                         Allows this transaction to be replaced by a transaction with higher fees
       "solving_data": {                 (json object, optional) Keys and scripts needed for producing a final transaction with a dummy signature.
                                         Used for fee estimation during coin selection.
         "pubkeys": [                    (json array, optional, default=[]) Public keys involved in this transaction.
           "pubkey",                     (string) A public key
           ...
         ],
         "scripts": [                    (json array, optional, default=[]) Scripts involved in this transaction.
           "script",                     (string) A script
           ...
         ],
         "descriptors": [                (json array, optional, default=[]) Descriptors that provide solving data for this transaction.
           "descriptor",                 (string) A descriptor
           ...
         ],
       },
     }

Result:
{                             (json object)
  "complete" : true|false,    (boolean) If the transaction has a complete set of signatures
  "txid" : "hex",             (string, optional) The transaction id for the send. Only 1 transaction is created regardless of the number of addresses.
  "hex" : "hex",              (string, optional) If add_to_wallet is false, the hex-encoded raw transaction with signature(s)
  "psbt" : "str"              (string, optional) If more signatures are needed, or if add_to_wallet is false, the base64-encoded (partially) signed transaction
}

Examples:

Send 0.1 BTC with a confirmation target of 6 blocks in economical fee estimate mode
> bitcoin-cli send '{"bc1q09vm5lfy0j5reeulh4x5752q25uqqvz34hufdl": 0.1}' 6 economical

Send 0.2 BTC with a fee rate of 1.1 sat/vB using positional arguments
> bitcoin-cli send '{"bc1q09vm5lfy0j5reeulh4x5752q25uqqvz34hufdl": 0.2}' null "unset" 1.1

Send 0.2 BTC with a fee rate of 1 sat/vB using the options argument
> bitcoin-cli send '{"bc1q09vm5lfy0j5reeulh4x5752q25uqqvz34hufdl": 0.2}' null "unset" null '{"fee_rate": 1}'

Send 0.3 BTC with a fee rate of 25 sat/vB using named arguments
> bitcoin-cli -named send outputs='{"bc1q09vm5lfy0j5reeulh4x5752q25uqqvz34hufdl": 0.3}' fee_rate=25

Create a transaction that should confirm the next block, with a specific input, and return result without adding to wallet or broadcasting to the network
> bitcoin-cli send '{"bc1q09vm5lfy0j5reeulh4x5752q25uqqvz34hufdl": 0.1}' 1 economical '{"add_to_wallet": false, "inputs": [{"txid":"a08e6907dbbd3d809776dbfc5d82e371b764ed838b5655e72f463568df1aadf0", "vout":1}]}'


