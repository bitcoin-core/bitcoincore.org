---
name: sendall
btcversion: 24.0.0
btcgroup: wallet
permalink: en/doc/24.0.0/rpc/wallet/sendall/
---

sendall ["address",{"address":amount,...},...] ( conf_target "estimate_mode" fee_rate options )

EXPERIMENTAL warning: this call may be changed in future releases.

Spend the value of all (or specific) confirmed UTXOs in the wallet to one or more recipients.
Unconfirmed inbound UTXOs and locked UTXOs will not be spent. Sendall will respect the avoid_reuse wallet flag.
If your wallet contains many small inputs, either because it received tiny payments or as a result of accumulating change, consider using `send_max` to exclude inputs that are worth less than the fees needed to spend them.

Arguments:
1. recipients                       (json array, required) The sendall destinations. Each address may only appear once.
                                    Optionally some recipients can be specified with an amount to perform payments, but at least one address must appear without a specified amount.
                                    
     [
       "address",                   (string, required) A bitcoin address which receives an equal share of the unspecified amount.
       {                            (json object)
         "address": amount,         (numeric or string, required) A key-value pair. The key (string) is the bitcoin address, the value (float or string) is the amount in BTC
         ...
       },
       ...
     ]
2. conf_target                      (numeric, optional, default=wallet -txconfirmtarget) Confirmation target in blocks
3. estimate_mode                    (string, optional, default="unset") The fee estimate mode, must be one of (case insensitive):
                                    "unset"
                                    "economical"
                                    "conservative"
4. fee_rate                         (numeric or string, optional, default=not set, fall back to wallet fee estimation) Specify a fee rate in sat/vB.
5. options                          (json object, optional)
     {
       "add_to_wallet": bool,       (boolean, optional, default=true) When false, returns the serialized transaction without broadcasting or adding it to the wallet
       "fee_rate": amount,          (numeric or string, optional, default=not set, fall back to wallet fee estimation) Specify a fee rate in sat/vB.
       "include_watching": bool,    (boolean, optional, default=true for watch-only wallets, otherwise false) Also select inputs which are watch-only.
                                    Only solvable inputs can be used. Watch-only destinations are solvable if the public key and/or output script was imported,
                                    e.g. with 'importpubkey' or 'importmulti' with the 'pubkeys' or 'desc' field.
       "inputs": [                  (json array, optional, default=[]) Use exactly the specified inputs to build the transaction. Specifying inputs is incompatible with send_max.
         {                          (json object)
           "txid": "hex",           (string, required) The transaction id
           "vout": n,               (numeric, required) The output number
           "sequence": n,           (numeric, optional, default=depends on the value of the 'replaceable' and 'locktime' arguments) The sequence number
         },
         ...
       ],
       "locktime": n,               (numeric, optional, default=0) Raw locktime. Non-0 value also locktime-activates inputs
       "lock_unspents": bool,       (boolean, optional, default=false) Lock selected unspent outputs
       "psbt": bool,                (boolean, optional, default=automatic) Always return a PSBT, implies add_to_wallet=false.
       "send_max": bool,            (boolean, optional, default=false) When true, only use UTXOs that can pay for their own fees to maximize the output amount. When 'false' (default), no UTXO is left behind. send_max is incompatible with providing specific inputs.
       "conf_target": n,            (numeric, optional, default=wallet -txconfirmtarget) Confirmation target in blocks
       "estimate_mode": "str",      (string, optional, default="unset") The fee estimate mode, must be one of (case insensitive):
                                    "unset"
                                    "economical"
                                    "conservative"
       "replaceable": bool,         (boolean, optional, default=wallet default) Marks this transaction as BIP125-replaceable.
                                    Allows this transaction to be replaced by a transaction with higher fees
       "solving_data": {            (json object, optional) Keys and scripts needed for producing a final transaction with a dummy signature.
                                    Used for fee estimation during coin selection.
         "pubkeys": [               (json array, optional, default=[]) Public keys involved in this transaction.
           "pubkey",                (string) A public key
           ...
         ],
         "scripts": [               (json array, optional, default=[]) Scripts involved in this transaction.
           "script",                (string) A script
           ...
         ],
         "descriptors": [           (json array, optional, default=[]) Descriptors that provide solving data for this transaction.
           "descriptor",            (string) A descriptor
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

Spend all UTXOs from the wallet with a fee rate of 1 sat/vB using named arguments
> bitcoin-cli -named sendall recipients='["bc1q09vm5lfy0j5reeulh4x5752q25uqqvz34hufdl"]' fee_rate=1

Spend all UTXOs with a fee rate of 1.1 sat/vB using positional arguments
> bitcoin-cli sendall '["bc1q09vm5lfy0j5reeulh4x5752q25uqqvz34hufdl"]' null "unset" 1.1

Spend all UTXOs split into equal amounts to two addresses with a fee rate of 1.5 sat/vB using the options argument
> bitcoin-cli sendall '["bc1q09vm5lfy0j5reeulh4x5752q25uqqvz34hufdl", "bc1q02ad21edsxd23d32dfgqqsz4vv4nmtfzuklhy3"]' null "unset" null '{"fee_rate": 1.5}'

Leave dust UTXOs in wallet, spend only UTXOs with positive effective value with a fee rate of 10 sat/vB using the options argument
> bitcoin-cli sendall '["bc1q09vm5lfy0j5reeulh4x5752q25uqqvz34hufdl"]' null "unset" null '{"fee_rate": 10, "send_max": true}'

Spend all UTXOs with a fee rate of 1.3 sat/vB using named arguments and sending a 0.25 BTC to another recipient
> bitcoin-cli -named sendall recipients='[{"bc1q02ad21edsxd23d32dfgqqsz4vv4nmtfzuklhy3": 0.25}, "bc1q09vm5lfy0j5reeulh4x5752q25uqqvz34hufdl"]' fee_rate=1.3



