---
name: testmempoolaccept
btcversion: 0.18.0
btcgroup: rawtransactions
permalink: en/doc/0.18.0/rpc/rawtransactions/testmempoolaccept/
---

testmempoolaccept ["rawtx",...] ( allowhighfees )

Returns result of mempool acceptance tests indicating if raw transaction (serialized, hex-encoded) would be accepted by mempool.

This checks if the transaction violates the consensus or policy rules.

See sendrawtransaction call.

Arguments:
1. rawtxs           (json array, required) An array of hex strings of raw transactions.
                    Length must be one for now.
     [
       "rawtx",     (string)
       ...
     ]
2. allowhighfees    (boolean, optional, default=false) Allow high fees

Result:
[                   (array) The result of the mempool acceptance test for each raw transaction in the input array.
                            Length is exactly one for now.
 {
  "txid"           (string) The transaction hash in hex
  "allowed"        (boolean) If the mempool allows this tx to be inserted
  "reject-reason"  (string) Rejection string (only present when 'allowed' is false)
 }
]

Examples:

Create a transaction
> bitcoin-cli createrawtransaction "[{\"txid\" : \"mytxid\",\"vout\":0}]" "{\"myaddress\":0.01}"
Sign the transaction, and get back the hex
> bitcoin-cli signrawtransactionwithwallet "myhex"

Test acceptance of the transaction (signed hex)
> bitcoin-cli testmempoolaccept ["signedhex"]

As a JSON-RPC call
> curl --user myusername --data-binary '{"jsonrpc": "1.0", "id":"curltest", "method": "testmempoolaccept", "params": [["signedhex"]] }' -H 'content-type: text/plain;' http://127.0.0.1:8332/


