---
title: Disclosure of DoS due to inv-to-send sets growing too large
name: blog-disclose-large-inv-to-send
id: en-blog-disclose-large-inv-to-send
lang: en
type: advisory
layout: post

## If this is a new post, reset this counter to 1.
version: 1

## Only true if release announcement or security announcement. English posts only
announcement: 1

excerpt: >
    The inv-to-send sets could grow too large to a point where the time spent sorting the sets would affect the node's ability to communicate with its peers.
---

Before Bitcoin Core v25.0, the per-peer `m_tx_inventory_to_send` sets could grow
too large to a point where sorting these sets when constructing inventory
messages would affect the node's ability to communicate with its peers. Network
conditions in early May 2023 triggered this DoS and affected block and transaction
propagation.

This issue is considered **Medium** severity.

## Details

As part of transaction relay, Bitcoin Core maintains a per-peer
`m_tx_inventory_to_send` set with transactions that should be announced to the
peer. When constructing an inventory message for a peer, the set is sorted by
transaction dependencies and feerate to prioritize high-feerate transactions and
to avoid leaking the order the node learned about the transactions. Before
Bitcoin Core v25.0, when constructing inventory messages, relevant (still in
mempool, not yet announced to us by the peer, above the fee filter) transactions
were being drained at a rate of 7 transactions per second.

In early May 2023, increased network activity caused the sets to grow faster
than they were being drained resulting in significant time spent sorting the
sets in the P2P communication thread. Additionally, peers that only listen for
transaction announcements but never announce any themselves (commonly referred
to as "spy nodes"), amplified this by having huge sets (with transactions they
already know about) that take a long time to sort. It was observed that sorting
took up nearly the complete time spent in the P2P communication thread, which
significantly affected block and transaction propagation as well as keeping
connection with peers alive.

This was fixed in [#27610](https://github.com/bitcoin/bitcoin/pull/27610) by 1)
earlier removing transactions that aren't in the mempool anymore and 2) by
dynamically increasing the set drainage rate depending on the set size.

## Attribution

Credit goes to Anthony Towns for working on a fix and to b10c for initially
reporting and narrowing the problem down to the slow inv-to-send sorting.

## Timeline

- 2023-05-02 - Problem first observed and reported
- 2023-05-11 - Fix is merged ([#27610](https://github.com/bitcoin/bitcoin/pull/27610))
- 2023-05-25 - v25.0 is released
- 2024-10-09 - Public disclosure

{% include references.md %}
