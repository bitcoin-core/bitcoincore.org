---
name: getdescriptoractivity
btcversion: 29.0.0
btcgroup: blockchain
permalink: en/doc/29.0.0/rpc/blockchain/getdescriptoractivity/
---

getdescriptoractivity ( ["blockhash",...] [scanobjects,...] include_mempool )

Get spend and receive activity associated with a set of descriptors for a set of blocks. This command pairs well with the `relevant_blocks` output of `scanblocks()`.
This call may take several minutes. If you encounter timeouts, try specifying no RPC timeout (bitcoin-cli -rpcclienttimeout=0)

Arguments:
1. blockhashes                   (json array, optional) The list of blockhashes to examine for activity. Order doesn't matter. Must be along main chain or an error is thrown.
                                 
     [
       "blockhash",              (string) A valid blockhash
       ...
     ]
2. scanobjects                   (json array, optional) Array of scan objects. Required for "start" action
                                 Every scan object is either a string descriptor or an object:
     [
       "descriptor",             (string) An output descriptor
       {                         (json object) An object with output descriptor and metadata
         "desc": "str",          (string, required) An output descriptor
         "range": n or [n,n],    (numeric or array, optional, default=1000) The range of HD chain indexes to explore (either end or [begin,end])
       },
       ...
     ]
3. include_mempool               (boolean, optional, default=true) Whether to include unconfirmed activity

Result:
{                                (json object)
  "activity" : [                 (json array) events
    {                            (json object)
      "type" : "str",            (string) always 'spend'
      "amount" : n,              (numeric) The total amount in BTC of the spent output
      "blockhash" : "hex",       (string, optional) The blockhash this spend appears in (omitted if unconfirmed)
      "height" : n,              (numeric, optional) Height of the spend (omitted if unconfirmed)
      "spend_txid" : "hex",      (string) The txid of the spending transaction
      "spend_vout" : n,          (numeric) The vout of the spend
      "prevout_txid" : "hex",    (string) The txid of the prevout
      "prevout_vout" : n,        (numeric) The vout of the prevout
      "prevout_spk" : {          (json object)
        "asm" : "str",           (string) Disassembly of the output script
        "desc" : "str",          (string) Inferred descriptor for the output
        "hex" : "hex",           (string) The raw output script bytes, hex-encoded
        "address" : "str",       (string, optional) The Bitcoin address (only if a well-defined address exists)
        "type" : "str"           (string) The type (one of: nonstandard, anchor, pubkey, pubkeyhash, scripthash, multisig, nulldata, witness_v0_scripthash, witness_v0_keyhash, witness_v1_taproot, witness_unknown)
      }
    },
    {                            (json object)
      "type" : "str",            (string) always 'receive'
      "amount" : n,              (numeric) The total amount in BTC of the new output
      "blockhash" : "hex",       (string, optional) The block that this receive is in (omitted if unconfirmed)
      "height" : n,              (numeric, optional) The height of the receive (omitted if unconfirmed)
      "txid" : "hex",            (string) The txid of the receiving transaction
      "vout" : n,                (numeric) The vout of the receiving output
      "output_spk" : {           (json object)
        "asm" : "str",           (string) Disassembly of the output script
        "desc" : "str",          (string) Inferred descriptor for the output
        "hex" : "hex",           (string) The raw output script bytes, hex-encoded
        "address" : "str",       (string, optional) The Bitcoin address (only if a well-defined address exists)
        "type" : "str"           (string) The type (one of: nonstandard, anchor, pubkey, pubkeyhash, scripthash, multisig, nulldata, witness_v0_scripthash, witness_v0_keyhash, witness_v1_taproot, witness_unknown)
      }
    },
    ...
  ]
}

Examples:
> bitcoin-cli getdescriptoractivity '["000000000000000000001347062c12fded7c528943c8ce133987e2e2f5a840ee"]' '["addr(bc1qzl6nsgqzu89a66l50cvwapnkw5shh23zarqkw9)"]'


