---
title: CVE-2024-52918 - Crash using malicious BIP72 URI
name: blog-disclose-bip70-crash
id: en-blog-disclose-bip70-crash
lang: en
type: advisory
layout: post

## If this is a new post, reset this counter to 1.
version: 2

## Only true if release announcement or security annoucement. English posts only
announcement: 1

excerpt: >
  The BIP70 implementation in Bitcoin-Qt could silently crash when opening a BIP72 URI. A fix was released on June 3rd, 2020 in Bitcoin Core 0.20.0.
---

Bitcoin-Qt could crash upon opening a [BIP72](https://github.com/bitcoin/bips/blob/master/bip-0072.mediawiki) URI.

This issue is considered **Medium** severity.

## Details

[BIP72](https://github.com/bitcoin/bips/blob/master/bip-0072.mediawiki) extends the BIP21 URI scheme
with an `r` parameter to fetch a payment request from. An attacker could simply point the URL
contained in the `r` parameter to a very large file, for which Bitcoin-Qt would try to allocate
enough memory and crash.

The victim could get tricked into opening a rogue payment request. The large download would happen
in the background with little to no output in the GUI until the application runs out of memory.

## Attribution

Credits go to Michael Ford (Fanquake) for responsibly disclosing the issue and providing a PoC.

## Timeline

- 2019-08-12 Michael Ford reports the bug to Cory Fields and Wladimir Van Der Laan
- 2019-10-16 Michael Ford opens PR [#17165](https://github.com/bitcoin/bitcoin/pull/17165) to get rid of BIP70 support entirely
- 2019-10-26 Michael's PR is merged into Bitcoin Core
- 2020-06-03 Bitcoin Core version 0.20.0 is released
- 2021-09-13 The last vulnerable Bitcoin Core version (0.19.x) goes EOL
- 2024-07-03 Public disclosure

{% include references.md %}
