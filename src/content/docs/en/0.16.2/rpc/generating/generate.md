---
name: generate
btcversion: 0.16.2
btcgroup: generating
permalink: en/doc/0.16.2/rpc/generating/generate/
---

generate nblocks ( maxtries )

Mine up to nblocks blocks immediately (before the RPC call returns) to an address in the wallet.

Arguments:
1. nblocks      (numeric, required) How many blocks are generated immediately.
2. maxtries     (numeric, optional) How many iterations to try (default = 1000000).

Result:
[ blockhashes ]     (array) hashes of blocks generated

Examples:

Generate 11 blocks
> bitcoin-cli generate 11


