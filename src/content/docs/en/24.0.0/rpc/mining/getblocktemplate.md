---
name: getblocktemplate
btcversion: 24.0.0
btcgroup: mining
permalink: en/doc/24.0.0/rpc/mining/getblocktemplate/
---

getblocktemplate ( "template_request" )

If the request parameters include a 'mode' key, that is used to explicitly select between the default 'template' request or a 'proposal'.
It returns data needed to construct a block to work on.
For full specification, see BIPs 22, 23, 9, and 145:
    https://github.com/bitcoin/bips/blob/master/bip-0022.mediawiki
    https://github.com/bitcoin/bips/blob/master/bip-0023.mediawiki
    https://github.com/bitcoin/bips/blob/master/bip-0009.mediawiki#getblocktemplate_changes
    https://github.com/bitcoin/bips/blob/master/bip-0145.mediawiki

Arguments:
1. template_request         (json object, optional, default={}) Format of the template
     {
       "mode": "str",       (string, optional) This must be set to "template", "proposal" (see BIP 23), or omitted
       "capabilities": [    (json array, optional) A list of strings
         "str",             (string) client side supported feature, 'longpoll', 'coinbasevalue', 'proposal', 'serverlist', 'workid'
         ...
       ],
       "rules": [           (json array, required) A list of strings
         "segwit",          (string, required) (literal) indicates client side segwit support
         "str",             (string) other client side supported softfork deployment
         ...
       ],
     }

Result (If the proposal was accepted with mode=='proposal'):
null    (json null)

Result (If the proposal was not accepted with mode=='proposal'):
"str"    (string) According to BIP22

Result (Otherwise):
{                                          (json object)
  "version" : n,                           (numeric) The preferred block version
  "rules" : [                              (json array) specific block rules that are to be enforced
    "str",                                 (string) name of a rule the client must understand to some extent; see BIP 9 for format
    ...
  ],
  "vbavailable" : {                        (json object) set of pending, supported versionbit (BIP 9) softfork deployments
    "rulename" : n,                        (numeric) identifies the bit number as indicating acceptance and readiness for the named softfork rule
    ...
  },
  "capabilities" : [                       (json array)
    "str",                                 (string) A supported feature, for example 'proposal'
    ...
  ],
  "vbrequired" : n,                        (numeric) bit mask of versionbits the server requires set in submissions
  "previousblockhash" : "str",             (string) The hash of current highest block
  "transactions" : [                       (json array) contents of non-coinbase transactions that should be included in the next block
    {                                      (json object)
      "data" : "hex",                      (string) transaction data encoded in hexadecimal (byte-for-byte)
      "txid" : "hex",                      (string) transaction id encoded in little-endian hexadecimal
      "hash" : "hex",                      (string) hash encoded in little-endian hexadecimal (including witness data)
      "depends" : [                        (json array) array of numbers
        n,                                 (numeric) transactions before this one (by 1-based index in 'transactions' list) that must be present in the final block if this one is
        ...
      ],
      "fee" : n,                           (numeric) difference in value between transaction inputs and outputs (in satoshis); for coinbase transactions, this is a negative Number of the total collected block fees (ie, not including the block subsidy); if key is not present, fee is unknown and clients MUST NOT assume there isn't one
      "sigops" : n,                        (numeric) total SigOps cost, as counted for purposes of block limits; if key is not present, sigop cost is unknown and clients MUST NOT assume it is zero
      "weight" : n                         (numeric) total transaction weight, as counted for purposes of block limits
    },
    ...
  ],
  "coinbaseaux" : {                        (json object) data that should be included in the coinbase's scriptSig content
    "key" : "hex",                         (string) values must be in the coinbase (keys may be ignored)
    ...
  },
  "coinbasevalue" : n,                     (numeric) maximum allowable input to coinbase transaction, including the generation award and transaction fees (in satoshis)
  "longpollid" : "str",                    (string) an id to include with a request to longpoll on an update to this template
  "target" : "str",                        (string) The hash target
  "mintime" : xxx,                         (numeric) The minimum timestamp appropriate for the next block time, expressed in UNIX epoch time
  "mutable" : [                            (json array) list of ways the block template may be changed
    "str",                                 (string) A way the block template may be changed, e.g. 'time', 'transactions', 'prevblock'
    ...
  ],
  "noncerange" : "hex",                    (string) A range of valid nonces
  "sigoplimit" : n,                        (numeric) limit of sigops in blocks
  "sizelimit" : n,                         (numeric) limit of block size
  "weightlimit" : n,                       (numeric, optional) limit of block weight
  "curtime" : xxx,                         (numeric) current timestamp in UNIX epoch time
  "bits" : "str",                          (string) compressed target of next block
  "height" : n,                            (numeric) The height of the next block
  "signet_challenge" : "hex",              (string, optional) Only on signet
  "default_witness_commitment" : "hex"     (string, optional) a valid witness commitment for the unmodified block template
}

Examples:
> bitcoin-cli getblocktemplate '{"rules": ["segwit"]}'
> curl --user myusername --data-binary '{"jsonrpc": "1.0", "id": "curltest", "method": "getblocktemplate", "params": [{"rules": ["segwit"]}]}' -H 'content-type: text/plain;' http://127.0.0.1:8332/


