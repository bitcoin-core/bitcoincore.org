---
name: loadtxoutset
btcversion: 27.0.0
btcgroup: blockchain
permalink: en/doc/27.0.0/rpc/blockchain/loadtxoutset/
---

loadtxoutset "path"

Load the serialized UTXO set from a file.
Once this snapshot is loaded, its contents will be deserialized into a second chainstate data structure, which is then used to sync to the network's tip. Meanwhile, the original chainstate will complete the initial block download process in the background, eventually validating up to the block that the snapshot is based upon.

The result is a usable bitcoind instance that is current with the network tip in a matter of minutes rather than hours. UTXO snapshot are typically obtained from third-party sources (HTTP, torrent, etc.) which is reasonable since their contents are always checked by hash.

You can find more information on this process in the `assumeutxo` design document (<https://github.com/bitcoin/bitcoin/blob/master/doc/design/assumeutxo.md>).

Arguments:
1. path    (string, required) path to the snapshot file. If relative, will be prefixed by datadir.

Result:
{                        (json object)
  "coins_loaded" : n,    (numeric) the number of coins loaded from the snapshot
  "tip_hash" : "hex",    (string) the hash of the base of the snapshot
  "base_height" : n,     (numeric) the height of the base of the snapshot
  "path" : "str"         (string) the absolute path that the snapshot was loaded from
}

Examples:
> bitcoin-cli loadtxoutset utxo.dat


