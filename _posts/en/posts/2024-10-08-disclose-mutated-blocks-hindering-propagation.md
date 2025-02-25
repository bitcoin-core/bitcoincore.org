---
title: CVE-2024-52921 - Hindered block propagation due to mutated blocks
name: blog-disclose-mutated-blocks-hindering-propagation
id: en-blog-disclose-mutated-blocks-hindering-propagation
lang: en
type: advisory
layout: post

## If this is a new post, reset this counter to 1.
version: 1

## Only true if release announcement or security announcement. English posts only
announcement: 1

excerpt: >
    A peer could hinder block propagation by sending mutated blocks.
---

Before Bitcoin Core v25.0, a peer sending mutated blocks could clear the
download state of other peers that also announced the block to us, which would
hinder block propagation.

This issue is considered **Medium** severity.

## Details

Bitcoin Core treats a block as mutated when, for example, the Merkle root in the
header or the witness commitment in the coinbase transaction doesn't match the
transactions in the block.

Before Bitcoin Core v25.0, a peer could clear the block download state of
other peers by sending an unrequested mutated block. This was a problem for, for
example, compact block relay. After receiving a compact block and while waiting
for a response to a `getblocktxn` request to reconstruct the full block,
receiving the mutated block would let Bitcoin Core forget about the compact
block reconstruction state. A `blocktxn` response arriving after the mutated
block couldn't be used to reconstruct the block. This hindered block propagation.

This was fixed in [#27608](https://github.com/bitcoin/bitcoin/pull/27608) by
making sure that a peer can only affect its own block download state and not the
download state of other peers.

## Attribution

Credit goes to Suhas Daftuar for noticing the problem and working on a fix.

## Timeline

- 2023-05-08 - A problem with mutated blocks is first reported in the [#bitcoin-core-dev IRC channel](https://bitcoin-irc.chaincode.com/bitcoin-core-dev/2023-05-08).
- 2023-05-10 - Fix is merged ([#27608](https://github.com/bitcoin/bitcoin/pull/27608))
- 2023-05-25 - v25.0 is released
- 2024-10-09 - Public disclosure

{% include references.md %}
