---
title: CVE-2024-52916 - Memory DoS using low-difficulty headers
name: blog-disclose-header-spam-checkpoint-bypass
id: en-blog-disclose-header-spam-checkpoint-bypass
lang: en
type: advisory
layout: post

## If this is a new post, reset this counter to 1.
version: 1

## Only true if release announcement or security annoucement. English posts only
announcement: 1

excerpt: >
  Nodes could be spammed with low-difficulty headers, which could be used to crash it. A fix was released on September 14th, 2017 in Bitcoin Core 0.15.0.
---

After Bitcoin Core 0.12.0 and before Bitcoin Core 0.15.0 a node could be spammed with minimum
difficulty headers, which could possibly be leveraged to crash it by OOM.

This issue is considered **Medium** severity.

## Details

Before the introduction of [headers
pre-synchronisation](https://github.com/bitcoin/bitcoin/pull/25717), nodes relied exclusively on
checkpoints to avoid getting spammed by low-difficulty headers.

In Bitcoin Core 0.12.0 a check for headers forking before the last checkpoint's height was moved to
after storing the header in `mapBlockIndex`. This allowed an attacker to grow the map unboundedly by
spamming headers whose parent is the genesis block (which only need difficulty 1 to create), as such
blocks bypassed the checkpoint logic.

## Attribution

Credits to Cory Fields for finding and responsibly disclosing the bug.

## Timeline

- 2017-08-08 Cory Fields privately reports the bug
- 2017-08-11 Pieter Wuille opens [PR #11028](https://github.com/bitcoin/bitcoin/pull/11028) to fix it
- 2017-08-14 PR #11028 is merged
- 2017-09-14 Bitcoin Core version 0.15.0 is released with a fix
- 2018-10-03 The last vulnerable version of Bitcoin Core (0.14.3) goes end of life
- 2024-07-03 Public disclosure.

{% include references.md %}
