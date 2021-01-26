---
name: scantxoutset
btcversion: 0.21.0
btcgroup: blockchain
permalink: en/doc/0.21.0/rpc/blockchain/scantxoutset/
---

scantxoutset "action" ( [scanobjects,...] )

EXPERIMENTAL warning: this call may be removed or changed in future releases.

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

Result:
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


