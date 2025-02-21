---
title: CVE-2024-52919 - Remote crash due to addr message spam
name: blog-disclose-addrman-idcount-in-overflow
id: blog-disclose-addrman-idcount-in-overflow
lang: en
type: advisory
layout: post

## If this is a new post, reset this counter to 1.
version: 1

## Only true if release announcement or security annoucement. English posts only
announcement: 1

excerpt: >
  Nodes could be spammed with addr messsages, which could be used to crash them. A fix was released on September 14th, 2021 in Bitcoin Core v22.0.
---

Disclosure of the details of an integer overflow bug which causes an assertion
crash, a fix for which was released on September 14th, 2021 in Bitcoin Core
version v22.0.

This issue is considered **High** severity.

## Details

`CAddrMan` has a 32-bit `nIdCount` field that is incremented on every insertion
into addrman, and which then becomes the identifier for the new entry. By
getting the victim to insert 2<sup>32</sup> entries (through e.g. spamming addr
messages), this identifier overflows, which leads to an assertion crash.

## Attribution

Credit goes to Eugene Siegel for discovering and disclosing the vulnerability,
and to Pieter Wuille for fixing the issue in
https://github.com/bitcoin/bitcoin/pull/22387.

## Timeline

* 2021-06-21 - Initial report sent to security@bitcoincore.org by Eugene Siegel
* 2021-07-19 - Fix is merged (https://github.com/bitcoin/bitcoin/pull/22387)
* 2021-09-13 - v22.0 is released
* 2024-07-31 - Public disclosure

{% include references.md %}
