---
title: CVE-2024-52922 - Hindered block propagation due to stalling peers
name: blog-disclose-stalling-peers-hindering-propagation
id: en-blog-disclose-stalling-peers-hindering-propagation
lang: en
type: advisory
layout: post

## If this is a new post, reset this counter to 1.
version: 1

## Only true if release announcement or security announcement. English posts only
announcement: 1

excerpt: >
    A peer could hinder block propagation by announcing blocks first and then simply withholding the block.
---

Before Bitcoin Core v25.1, an attacker can cause a node to not
download the latest block.

This issue is considered **Medium** severity.

## Details

When receiving a new block announcement via a headers or compact
blocks message, the delivering peer is requested either the full
block or missing transaction details by the receiving node. If
the announcing peer then doesn't respond as the peer to peer
protocol requires, the affected Bitcoin Core node will wait
up to 10 minutes before disconnecting the peer and making another
block download attempt. If the attacker is able to
make multiple incoming or outgoing connections, this process
can be repeated.

Delaying block delivery can cause network degradation by slowing down network convergence,
making mining payouts less fair, and causing liveliness issues.

This issue was further exacerbated by other issues disclosed recently (for
instance the [inventory build-up](https://bitcoincore.org/en/2024/10/08/disclose-large-inv-to-send/)),
when mempools were relatively heterogeneous, disallowing
opportunistic reconstruction of compact blocks by honest peers.

A mitigation was introduced in [#27626](https://github.com/bitcoin/bitcoin/pull/27626),
introduced in Bitcoin Core v26.0 and backported to v25.1.
It ensures that blocks can be requested concurrently from up to 3
high-bandwidth compact block peers, one of which is required
to be an outbound connection.

## Attribution

Reported and fixed by Greg Sanders.

## Timeline

- 2023-05-08 - Users reporting block timeouts in the [#bitcoin-core-dev IRC channel](https://bitcoin-irc.chaincode.com/bitcoin-core-dev/2023-05-08)
- 2023-05-09 - First github issues describing the issue https://github.com/bitcoin/bitcoin/issues/25258#issuecomment-1540028533
- 2023-05-11 - Mitigation PR opened https://github.com/bitcoin/bitcoin/pull/27626
- 2023-05-24 - PR merged prior to Bitcoin Core v26.0
- 2023-05-25 - Backport to Bitcoin Core v25.1 merged https://github.com/bitcoin/bitcoin/pull/27752
- 2023-10-19 - Bitcoin Core v25.1 Released
- 2024-11-05 - Public disclosure

{% include references.md %}
