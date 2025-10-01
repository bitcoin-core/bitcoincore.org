---
name: estimatesmartfee
btcversion: 28.0.0
btcgroup: util
permalink: en/doc/28.0.0/rpc/util/estimatesmartfee/
---

estimatesmartfee conf_target ( "estimate_mode" )

Estimates the approximate fee per kilobyte needed for a transaction to begin
confirmation within conf_target blocks if possible and return the number of blocks
for which the estimate is valid. Uses virtual transaction size as defined
in BIP 141 (witness data is discounted).

Arguments:
1. conf_target      (numeric, required) Confirmation target in blocks (1 - 1008)
2. estimate_mode    (string, optional, default="economical") The fee estimate mode.
                    unset, economical, conservative 
                    unset means no mode set (default mode will be used). 
                    economical estimates use a shorter time horizon, making them more
                    responsive to short-term drops in the prevailing fee market. This mode
                    potentially returns a lower fee rate estimate.
                    conservative estimates use a longer time horizon, making them
                    less responsive to short-term drops in the prevailing fee market. This mode
                    potentially returns a higher fee rate estimate.
                    

Result:
{                   (json object)
  "feerate" : n,    (numeric, optional) estimate fee rate in BTC/kvB (only present if no errors were encountered)
  "errors" : [      (json array, optional) Errors encountered during processing (if there are any)
    "str",          (string) error
    ...
  ],
  "blocks" : n      (numeric) block number where estimate was found
                    The request target will be clamped between 2 and the highest target
                    fee estimation is able to return based on how long it has been running.
                    An error is returned if not enough transactions and blocks
                    have been observed to make an estimate for any number of blocks.
}

Examples:
> bitcoin-cli estimatesmartfee 6
> curl --user myusername --data-binary '{"jsonrpc": "2.0", "id": "curltest", "method": "estimatesmartfee", "params": [6]}' -H 'content-type: application/json' http://127.0.0.1:8332/


