---
name: deriveaddresses
btcversion: 24.0.0
btcgroup: util
permalink: en/doc/24.0.0/rpc/util/deriveaddresses/
---

deriveaddresses "descriptor" ( range )

Derives one or more addresses corresponding to an output descriptor.
Examples of output descriptors are:
    pkh(<pubkey>)                        P2PKH outputs for the given pubkey
    wpkh(<pubkey>)                       Native segwit P2PKH outputs for the given pubkey
    sh(multi(<n>,<pubkey>,<pubkey>,...)) P2SH-multisig outputs for the given threshold and pubkeys
    raw(<hex script>)                    Outputs whose scriptPubKey equals the specified hex scripts

In the above, <pubkey> either refers to a fixed public key in hexadecimal notation, or to an xpub/xprv optionally followed by one
or more path elements separated by "/", where "h" represents a hardened child key.
For more information on output descriptors, see the documentation in the doc/descriptors.md file.

Arguments:
1. descriptor    (string, required) The descriptor.
2. range         (numeric or array, optional) If a ranged descriptor is used, this specifies the end or the range (in [begin,end] notation) to derive.

Result:
[           (json array)
  "str",    (string) the derived addresses
  ...
]

Examples:
First three native segwit receive addresses
> bitcoin-cli deriveaddresses "wpkh([d34db33f/84h/0h/0h]xpub6DJ2dNUysrn5Vt36jH2KLBT2i1auw1tTSSomg8PhqNiUtx8QX2SvC9nrHu81fT41fvDUnhMjEzQgXnQjKEu3oaqMSzhSrHMxyyoEAmUHQbY/0/*)#cjjspncu" "[0,2]"
> curl --user myusername --data-binary '{"jsonrpc": "1.0", "id": "curltest", "method": "deriveaddresses", "params": ["wpkh([d34db33f/84h/0h/0h]xpub6DJ2dNUysrn5Vt36jH2KLBT2i1auw1tTSSomg8PhqNiUtx8QX2SvC9nrHu81fT41fvDUnhMjEzQgXnQjKEu3oaqMSzhSrHMxyyoEAmUHQbY/0/*)#cjjspncu", "[0,2]"]}' -H 'content-type: text/plain;' http://127.0.0.1:8332/


