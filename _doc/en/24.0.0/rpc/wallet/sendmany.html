---
name: sendmany
btcversion: 24.0.0
btcgroup: wallet
permalink: en/doc/24.0.0/rpc/wallet/sendmany/
---

sendmany "" {"address":amount,...} ( minconf "comment" ["address",...] replaceable conf_target "estimate_mode" fee_rate verbose )

Send multiple times. Amounts are double-precision floating point numbers.
Requires wallet passphrase to be set with walletpassphrase call if wallet is encrypted.

Arguments:
1. dummy                     (string, required) Must be set to "" for backwards compatibility.
2. amounts                   (json object, required) The addresses and amounts
     {
       "address": amount,    (numeric or string, required) The bitcoin address is the key, the numeric amount (can be string) in BTC is the value
       ...
     }
3. minconf                   (numeric, optional) Ignored dummy value
4. comment                   (string, optional) A comment
5. subtractfeefrom           (json array, optional) The addresses.
                             The fee will be equally deducted from the amount of each selected address.
                             Those recipients will receive less bitcoins than you enter in their corresponding amount field.
                             If no addresses are specified here, the sender pays the fee.
     [
       "address",            (string) Subtract fee from this address
       ...
     ]
6. replaceable               (boolean, optional, default=wallet default) Signal that this transaction can be replaced by a transaction (BIP 125)
7. conf_target               (numeric, optional, default=wallet -txconfirmtarget) Confirmation target in blocks
8. estimate_mode             (string, optional, default="unset") The fee estimate mode, must be one of (case insensitive):
                             "unset"
                             "economical"
                             "conservative"
9. fee_rate                  (numeric or string, optional, default=not set, fall back to wallet fee estimation) Specify a fee rate in sat/vB.
10. verbose                  (boolean, optional, default=false) If true, return extra infomration about the transaction.

Result (if verbose is not set or set to false):
"hex"    (string) The transaction id for the send. Only 1 transaction is created regardless of
         the number of addresses.

Result (if verbose is set to true):
{                          (json object)
  "txid" : "hex",          (string) The transaction id for the send. Only 1 transaction is created regardless of
                           the number of addresses.
  "fee_reason" : "str"     (string) The transaction fee reason.
}

Examples:

Send two amounts to two different addresses:
> bitcoin-cli sendmany "" "{\"bc1q09vm5lfy0j5reeulh4x5752q25uqqvz34hufdl\":0.01,\"bc1q02ad21edsxd23d32dfgqqsz4vv4nmtfzuklhy3\":0.02}"

Send two amounts to two different addresses setting the confirmation and comment:
> bitcoin-cli sendmany "" "{\"bc1q09vm5lfy0j5reeulh4x5752q25uqqvz34hufdl\":0.01,\"bc1q02ad21edsxd23d32dfgqqsz4vv4nmtfzuklhy3\":0.02}" 6 "testing"

Send two amounts to two different addresses, subtract fee from amount:
> bitcoin-cli sendmany "" "{\"bc1q09vm5lfy0j5reeulh4x5752q25uqqvz34hufdl\":0.01,\"bc1q02ad21edsxd23d32dfgqqsz4vv4nmtfzuklhy3\":0.02}" 1 "" "[\"bc1q09vm5lfy0j5reeulh4x5752q25uqqvz34hufdl\",\"bc1q02ad21edsxd23d32dfgqqsz4vv4nmtfzuklhy3\"]"

As a JSON-RPC call
> curl --user myusername --data-binary '{"jsonrpc": "1.0", "id": "curltest", "method": "sendmany", "params": ["", {"bc1q09vm5lfy0j5reeulh4x5752q25uqqvz34hufdl":0.01,"bc1q02ad21edsxd23d32dfgqqsz4vv4nmtfzuklhy3":0.02}, 6, "testing"]}' -H 'content-type: text/plain;' http://127.0.0.1:8332/


