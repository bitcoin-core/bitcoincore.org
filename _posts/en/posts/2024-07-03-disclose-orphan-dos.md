---
title: CVE-2024-52914 - Significant DoS due to orphan handling
name: blog-disclose-orphan-dos
id: en-blog-disclose-orphan-dos
lang: en
type: advisory
layout: post

## If this is a new post, reset this counter to 1.
version: 1

## Only true if release announcement or security annoucement. English posts only
announcement: 1

excerpt: >
  A node could be stalled for hours when receiving a specially crafted unconfirmed transaction. A fix was released on May 18th, 2019 in Bitcoin Core 0.18.0.
---

A node could be stalled for hours when processing the orphans of a specially crafted unconfirmed
transaction.

This issue is considered **High** severity.

## Details

After accepting a transaction into its mempool, the node would go through its cache of orphan
transactions to find if this new accepted transaction makes it possible to accept any. This search
was quadratic: for each output in the newly accepted transaction it would go through all cached
orphan transactions (limited to 100). By specially crafting the orphan transactions to be invalid
yet expensive to validate a node could be stalled for several hours.

The stall was fixed by Pieter Wuille in [PR #15644](https://github.com/bitcoin/bitcoin/pull/15644)
by interrupting the orphan resolution to process new messages when a match is found (whether the
orphan turns out to be valid or not).

## Attribution

Credits to sec.eine for responsibly disclosing the bug and providing feedback on the fix.

## Timeline

- 2019-03-19 sec.eine reports the issue to Greg Maxwell by email
- 2019-03-21 Greg Maxwell responds with information about the proposed patch
- 2019-03-22 sec.eine gives feedback on the patch ("seems solid and [..] doesn't attract attention")
- 2019-03-22 Pieter Wuille opens PR #15644
- 2019-04-01 PR #15644 is merged
- 2019-05-18 Bitcoin Core version 0.18.0 is released with a fix
- 2020-07-22 The issue is [partially disclosed](https://bitcoincore.reviews/15644#l-285) during a PR review club
- 2020-08-01 The last vulnerable Bitcoin Core version (0.17.x) goes EOL
- 2024-07-03 Public disclosure.

{% include references.md %}
