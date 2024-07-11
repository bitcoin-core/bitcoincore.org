---
title: Disclosure of CVE-2020-14198
name: blog-disclose-unbounded-banlist
id: en-blog-disclose-unbounded-banlist
lang: en
type: advisory
layout: post

## If this is a new post, reset this counter to 1.
version: 1

## Only true if release announcement or security annoucement. English posts only
announcement: 1

excerpt: >
  Nodes could be subject to CPU and memory DoS when attacked by lots of distinct IPs. A fix was released on August 1st, 2020 in Bitcoin Core 0.20.1.
---

Bitcoin Core maintained an unlimited list of banned IP addresses and performed a quadratic operation
on it. This could lead to an OOM crash and a CPU Dos.

This issue is considered **High** severity.

## Details

Bitcoin Core maintained a list of banned IP addresses. This list was not bounded and could be
manipulated by an adversary. Adding new entries to this list was particularly cheap for an attacker
when considering IPV6. In addition, when receiving a `GETADDR` message, Bitcoin Core would scan the
entire ban list for every single address to be returned (up to 2500).

## Attribution

Calin Culianu first responsibly disclosed it. Calin later publicly disclosed the bug in [a PR
comment](https://github.com/bitcoin/bitcoin/pull/15617#issuecomment-640898523).

On the same day Jason Cox from Bitcoin ABC emailed the Bitcoin Core project to share this same
report they also received.

## Timeline

- 2020-06-08 Calin Culianu privately reports the bug to the Bitcoin Core project
- 2020-06-08 Jason Cox privately shares the (same) report sent to Bitcoin ABC with Bitcoin Core
- 2020-06-08 Calin Culianu publicly discloses the vulnerability on the original PR which introduced the quadratic behaviour
- 2020-06-09 Pieter Wuille opens PR [#19219](https://github.com/bitcoin/bitcoin/pull/19219) which fixes both the unbounded memory usage and the quadratic behaviour
- 2020-06-16 Luke Dashjr gets assigned CVE-2020-14198 for this vulnerability after his request
- 2020-07-07 Pieter's PR is merged
- 2020-08-01 Bitcoin Core 0.20.1 is released with the fix
- 2021-01-14 Bitcoin Core 0.21.0 is released with the fix
- 2022-04-25 The last vulnerable Bitcoin Core version (0.20.0) goes EOL
- 2024-07-03 (Official) Public Disclosure

{% include references.md %}
