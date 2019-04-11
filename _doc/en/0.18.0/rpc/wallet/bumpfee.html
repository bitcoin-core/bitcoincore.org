---
name: bumpfee
btcversion: 0.18.0
btcgroup: wallet
permalink: en/doc/0.18.0/rpc/wallet/bumpfee/
---

bumpfee "txid" ( options )

Bumps the fee of an opt-in-RBF transaction T, replacing it with a new transaction B.
An opt-in RBF transaction with the given txid must be in the wallet.
The command will pay the additional fee by decreasing (or perhaps removing) its change output.
If the change output is not big enough to cover the increased fee, the command will currently fail
instead of adding new inputs to compensate. (A future implementation could improve this.)
The command will fail if the wallet or mempool contains a transaction that spends one of T's outputs.
By default, the new fee will be calculated automatically using estimatesmartfee.
The user can specify a confirmation target for estimatesmartfee.
Alternatively, the user can specify totalFee, or use RPC settxfee to set a higher fee rate.
At a minimum, the new fee rate must be high enough to pay an additional new relay fee (incrementalfee
returned by getnetworkinfo) to enter the node's mempool.

Arguments:
1. txid                           (string, required) The txid to be bumped
2. options                        (json object, optional)
     {
       "confTarget": n,           (numeric, optional, default=fallback to wallet's default) Confirmation target (in blocks)
       "totalFee": n,             (numeric, optional, default=fallback to 'confTarget') Total fee (NOT feerate) to pay, in satoshis.
                                  In rare cases, the actual fee paid might be slightly higher than the specified
                                  totalFee if the tx change output has to be removed because it is too close to
                                  the dust threshold.
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


