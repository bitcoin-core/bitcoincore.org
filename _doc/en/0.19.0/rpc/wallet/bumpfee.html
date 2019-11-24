---
name: bumpfee
btcversion: 0.19.0
btcgroup: wallet
permalink: en/doc/0.19.0/rpc/wallet/bumpfee/
---

bumpfee "txid" ( options )

Bumps the fee of an opt-in-RBF transaction T, replacing it with a new transaction B.
An opt-in RBF transaction with the given txid must be in the wallet.
The command will pay the additional fee by reducing change outputs or adding inputs when necessary. It may add a new change output if one does not already exist.
If `totalFee` (DEPRECATED) is given, adding inputs is not supported, so there must be a single change output that is big enough or it will fail.
All inputs in the original transaction will be included in the replacement transaction.
The command will fail if the wallet or mempool contains a transaction that spends one of T's outputs.
By default, the new fee will be calculated automatically using estimatesmartfee.
The user can specify a confirmation target for estimatesmartfee.
Alternatively, the user can specify totalFee (DEPRECATED), or fee_rate (BTC per kB) for the new transaction .
At a minimum, the new fee rate must be high enough to pay an additional new relay fee (incrementalfee
returned by getnetworkinfo) to enter the node's mempool.

Arguments:
1. txid                           (string, required) The txid to be bumped
2. options                        (json object, optional)
     {
       "confTarget": n,           (numeric, optional, default=wallet default) Confirmation target (in blocks)
       "totalFee": n,             (numeric, optional, default=fallback to 'confTarget') Total fee (NOT feerate) to pay, in satoshis. (DEPRECATED)
                                  In rare cases, the actual fee paid might be slightly higher than the specified
                                  totalFee if the tx change output has to be removed because it is too close to
                                  the dust threshold.
       "fee_rate": n,             (numeric, optional, default=fallback to 'confTarget') FeeRate (NOT total fee) to pay, in BTC per kB
                                  Specify a fee rate instead of relying on the built-in fee estimator.
                                  Must be at least 0.0001 BTC per kB higher than the current transaction fee rate.
                                  
       "replaceable": bool,       (boolean, optional, default=true) Whether the new transaction should still be
                                  marked bip-125 replaceable. If true, the sequence numbers in the transaction will
                                  be left unchanged from the original. If false, any input sequence numbers in the
                                  original transaction that were less than 0xfffffffe will be increased to 0xfffffffe
                                  so the new transaction will not be explicitly bip-125 replaceable (though it may
                                  still be replaceable in practice, for example if it has unconfirmed ancestors which
                                  are replaceable).
       "estimate_mode": "str",    (string, optional, default=UNSET) The fee estimate mode, must be one of:
                                  "UNSET"
                                  "ECONOMICAL"
                                  "CONSERVATIVE"
     }

Result:
{
  "txid":    "value",   (string)  The id of the new transaction
  "origfee":  n,         (numeric) Fee of the replaced transaction
  "fee":      n,         (numeric) Fee of the new transaction
  "errors":  [ str... ] (json array of strings) Errors encountered during processing (may be empty)
}

Examples:

Bump the fee, get the new transaction's txid
> bitcoin-cli bumpfee <txid>


