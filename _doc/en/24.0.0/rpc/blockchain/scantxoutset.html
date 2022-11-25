---
name: scantxoutset
btcversion: 24.0.0
btcgroup: blockchain
permalink: en/doc/24.0.0/rpc/blockchain/scantxoutset/
---

scantxoutset "action" ( [scanobjects,...] )

Scans the unspent transaction output set for entries that match certain output descriptors.
Examples of output descriptors are:
    addr(<address>)                      Outputs whose scriptPubKey corresponds to the specified address (does not include P2PK)
    raw(<hex script>)                    Outputs whose scriptPubKey equals the specified hex scripts
    combo(<pubkey>)                      P2PK, P2PKH, P2WPKH, and P2SH-P2WPKH outputs for the given pubkey
    pkh(<pubkey>)                        P2PKH outputs for the given pubkey
    sh(multi(<n>,<pubkey>,<pubkey>,...)) P2SH-multisig outputs for the given threshold and pubkeys

In the above, <pubkey> either refers to a fixed public key in hexadecimal notation, or to an xpub/xprv optionally followed by one
or more path elements separated by "/", and optionally ending in "/*" (unhardened), or "/*'" or "/*h" (hardened) to specify all
unhardened or hardened child keys.
In the latter case, a range needs to be specified by below if different from 1000.
For more information on output descriptors, see the documentation in the doc/descriptors.md file.

Arguments:
1. action                        (string, required) The action to execute
                                 "start" for starting a scan
                                 "abort" for aborting the current scan (returns true when abort was successful)
                                 "status" for progress report (in %) of the current scan
2. scanobjects                   (json array) Array of scan objects. Required for "start" action
                                 Every scan object is either a string descriptor or an object:
     [
       "descriptor",             (string) An output descriptor
       {                         (json object) An object with output descriptor and metadata
         "desc": "str",          (string, required) An output descriptor
         "range": n or [n,n],    (numeric or array, optional, default=1000) The range of HD chain indexes to explore (either end or [begin,end])
       },
       ...
     ]

Result (when action=='start'; only returns after scan completes):
{                                (json object)
  "success" : true|false,        (boolean) Whether the scan was completed
  "txouts" : n,                  (numeric) The number of unspent transaction outputs scanned
  "height" : n,                  (numeric) The current block height (index)
  "bestblock" : "hex",           (string) The hash of the block at the tip of the chain
  "unspents" : [                 (json array)
    {                            (json object)
      "txid" : "hex",            (string) The transaction id
      "vout" : n,                (numeric) The vout value
      "scriptPubKey" : "hex",    (string) The script key
      "desc" : "str",            (string) A specialized descriptor for the matched scriptPubKey
      "amount" : n,              (numeric) The total amount in BTC of the unspent output
      "height" : n               (numeric) Height of the unspent transaction output
    },
    ...
  ],
  "total_amount" : n             (numeric) The total amount of all found unspent outputs in BTC
}

Result (when action=='abort'):
true|false    (boolean) True if scan will be aborted (not necessarily before this RPC returns), or false if there is no scan to abort

Result (when action=='status' and a scan is currently in progress):
{                    (json object)
  "progress" : n     (numeric) Approximate percent complete
}

Result (when action=='status' and no scan is in progress - possibly already completed):
null    (json null)

Examples:
> bitcoin-cli scantxoutset start '["raw(76a91411b366edfc0a8b66feebae5c2e25a7b6a5d1cf3188ac)#fm24fxxy"]'
> bitcoin-cli scantxoutset status
> bitcoin-cli scantxoutset abort
> curl --user myusername --data-binary '{"jsonrpc": "1.0", "id": "curltest", "method": "scantxoutset", "params": ["start", ["raw(76a91411b366edfc0a8b66feebae5c2e25a7b6a5d1cf3188ac)#fm24fxxy"]]}' -H 'content-type: text/plain;' http://127.0.0.1:8332/
> curl --user myusername --data-binary '{"jsonrpc": "1.0", "id": "curltest", "method": "scantxoutset", "params": ["status"]}' -H 'content-type: text/plain;' http://127.0.0.1:8332/
> curl --user myusername --data-binary '{"jsonrpc": "1.0", "id": "curltest", "method": "scantxoutset", "params": ["abort"]}' -H 'content-type: text/plain;' http://127.0.0.1:8332/


