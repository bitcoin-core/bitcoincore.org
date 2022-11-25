---
name: psbtbumpfee
btcversion: 24.0.0
btcgroup: wallet
permalink: en/doc/24.0.0/rpc/wallet/psbtbumpfee/
---

psbtbumpfee "txid" ( options )

Bumps the fee of an opt-in-RBF transaction T, replacing it with a new transaction B.
Returns a PSBT instead of creating and signing a new transaction.
An opt-in RBF transaction with the given txid must be in the wallet.
The command will pay the additional fee by reducing change outputs or adding inputs when necessary.
It may add a new change output if one does not already exist.
All inputs in the original transaction will be included in the replacement transaction.
The command will fail if the wallet or mempool contains a transaction that spends one of T's outputs.
By default, the new fee will be calculated automatically using the estimatesmartfee RPC.
The user can specify a confirmation target for estimatesmartfee.
Alternatively, the user can specify a fee rate in sat/vB for the new transaction.
At a minimum, the new fee rate must be high enough to pay an additional new relay fee (incrementalfee
returned by getnetworkinfo) to enter the node's mempool.
* WARNING: before version 0.21, fee_rate was in BTC/kvB. As of 0.21, fee_rate is in sat/vB. *

Arguments:
1. txid                           (string, required) The txid to be bumped
2. options                        (json object, optional)
     {
       "conf_target": n,          (numeric, optional, default=wallet -txconfirmtarget) Confirmation target in blocks
                                  
       "fee_rate": amount,        (numeric or string, optional, default=not set, fall back to wallet fee estimation) 
                                  Specify a fee rate in sat/vB instead of relying on the built-in fee estimator.
                                  Must be at least 1.000 sat/vB higher than the current transaction fee rate.
                                  WARNING: before version 0.21, fee_rate was in BTC/kvB. As of 0.21, fee_rate is in sat/vB.
                                  
       "replaceable": bool,       (boolean, optional, default=true) Whether the new transaction should still be
                                  marked bip-125 replaceable. If true, the sequence numbers in the transaction will
                                  be left unchanged from the original. If false, any input sequence numbers in the
                                  original transaction that were less than 0xfffffffe will be increased to 0xfffffffe
                                  so the new transaction will not be explicitly bip-125 replaceable (though it may
                                  still be replaceable in practice, for example if it has unconfirmed ancestors which
                                  are replaceable).
                                  
       "estimate_mode": "str",    (string, optional, default="unset") The fee estimate mode, must be one of (case insensitive):
                                  "unset"
                                  "economical"
                                  "conservative"
     }

Result:
{                    (json object)
  "psbt" : "str",    (string) The base64-encoded unsigned PSBT of the new transaction.
  "origfee" : n,     (numeric) The fee of the replaced transaction.
  "fee" : n,         (numeric) The fee of the new transaction.
  "errors" : [       (json array) Errors encountered during processing (may be empty).
    "str",           (string)
    ...
  ]
}

Examples:

Bump the fee, get the new transaction's psbt
> bitcoin-cli psbtbumpfee <txid>


