---
title: CVE-2026-117348 - Node stall through duplicate orphan announcements
name: blog-disclose-orphan-inv-stall
id: en-blog-disclose-orphan-inv-stall
lang: en
type: advisory
layout: post
redirect_from:
  - /en/2026/11/20/disclose-orphan-inv-stall
  - /en/2026/11/20/disclose-cve-2026-117348

## If this is a new post, reset this counter to 1.
version: 1

## Only true if release announcement or security annoucement. English posts only
announcement: 1

excerpt: >
  A peer could stop a node from processing blocks and transactions by repeatedly announcing an orphan transaction with many missing parents. A fix was released in Bitcoin Core versions 30.5, 31.3 and 32.1.
---

A peer could stop a node from processing blocks and transactions temporarily by repeatedly announcing an orphan
transaction with many missing parents. A fix was released in Bitcoin Core versions 30.5, 31.3 and
32.1. Affected versions are 29.0 to 29.5, 30.0 to 30.4 and 31.0 to 31.2. The 29.x series has
reached end of life and will not receive a fix.

This issue is considered **Medium** severity and **easily discoverable**.

## Details

Bitcoin Core keeps transactions whose parents it has not yet seen ("orphans") in its orphanage, and
asks peers that announce an orphan for its missing parents. Since version 29.0, when a peer announces
a transaction by wtxid and that transaction is in the orphanage, Bitcoin Core builds the list of the
orphan's parents and checks each one against its mempool and its filters of recently seen
transactions, to decide whether to ask that peer for them.

A single `inv` message can carry up to 50,000 entries, and duplicate entries were not skipped. An
attacker could send a node an orphan with thousands of missing parents, then send `inv` messages
repeating its wtxid 50,000 times. Each such message stalls the node for over a dozen seconds, during
which it processes no messages from other peers, including blocks.

The issue was fixed by processing each transaction announced in an `inv` message only once.

## Attribution

This issue was discovered through an automated LLM scan, as part of Elias Rohrer's [Project
Loupe](https://github.com/project-loupe/loupe) at Spiral. The Loupe finding was validated and
surfaced to the Bitcoin Core security team by Gregory Sanders.

## Timeline

- 2026-07-08 - Project Loupe reports the issue
- 2026-08-04 - Stutxo independently reports the issue
- 2026-08-04 - Gregory Sanders opens PR [#35878](https://github.com/bitcoin/bitcoin/pull/35878) with
  a fix
- 2026-08-06 - PR #35878 is merged into master
- 2026-10-27 - Bitcoin Core 32.1 is released with a fix
- 2026-11-17 - The fix is backported to the 31.x branch in PR
  [#36892](https://github.com/bitcoin/bitcoin/pull/36892)
- 2026-11-17 - The fix is backported to the 30.x branch in PR
  [#36893](https://github.com/bitcoin/bitcoin/pull/36893)
- 2026-11-20 - Bitcoin Core 30.5 and 31.3 are released with a fix
- 2026-11-20 - Public disclosure

{% include references.md %}
