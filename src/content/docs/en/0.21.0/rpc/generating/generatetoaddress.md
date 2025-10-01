---
name: generatetoaddress
btcversion: 0.21.0
btcgroup: generating
permalink: en/doc/0.21.0/rpc/generating/generatetoaddress/
---

generatetoaddress nblocks "address" ( maxtries )

Mine blocks immediately to a specified address (before the RPC call returns)

Arguments:
1. nblocks     (numeric, required) How many blocks are generated immediately.
2. address     (string, required) The address to send the newly generated bitcoin to.
3. maxtries    (numeric, optional, default=1000000) How many iterations to try.

Result:
[           (json array) hashes of blocks generated
  "hex",    (string) blockhash
  ...
]

Examples:

Generate 11 blocks to myaddress
> bitcoin-cli generatetoaddress 11 "myaddress"
If you are using the Bitcoin Core wallet, you can get a new address to send the newly generated bitcoin to with:
> bitcoin-cli getnewaddress 


