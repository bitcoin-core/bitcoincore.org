---
name: generatetoaddress
btcversion: 0.18.0
btcgroup: generating
permalink: en/doc/0.18.0/rpc/generating/generatetoaddress/
---

generatetoaddress nblocks "address" ( maxtries )

Mine blocks immediately to a specified address (before the RPC call returns)

Arguments:
1. nblocks     (numeric, required) How many blocks are generated immediately.
2. address     (string, required) The address to send the newly generated bitcoin to.
3. maxtries    (numeric, optional, default=1000000) How many iterations to try.

Result:
[ blockhashes ]     (array) hashes of blocks generated

Examples:

Generate 11 blocks to myaddress
> bitcoin-cli generatetoaddress 11 "myaddress"
If you are running the bitcoin core wallet, you can get a new address to send the newly generated bitcoin to with:
> bitcoin-cli getnewaddress 


