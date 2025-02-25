---
title: CVE-2024-52915 - Memory DoS using huge INV messages
name: blog-disclose-inv-buffer-blowup
id: en-blog-disclose-inv-buffer-blowup
lang: en
type: advisory
layout: post

## If this is a new post, reset this counter to 1.
version: 1

## Only true if release announcement or security annoucement. English posts only
announcement: 1

excerpt: >
  Nodes would allocate up to 50 MB of memory per attacker sending a malicious `INV` message. A fix was released on June 3rd, 2020 in Bitcoin Core 0.20.0.
---

A node could be forced to allocate a significant amount of memory upon receiving a specially crafted
`INV` message. This was particularly an issue for nodes with little available memory or a large
number of connections.

This issue is considered **Medium** severity.

## Details

An `INV` message filled with 50,000 block items could cause 50,000 `getheaders` responses to be sent
in a single `ProcessMessages()` call. Each response contains a locator and is around 1 kB. All would
be put into the send buffer at once. The attacker could just refuse to receive data to prevent the
50 MB buffer from draining.

John Newbery opened [PR #18962](https://github.com/bitcoin/bitcoin/pull/18962) to fix this issue
pretexting a bandwidth gain from sending a single `GETHEADERS` per received `INV`.

## Attribution

Credits to John Newbery for finding this bug, responsibly disclosing it and fixing it.

## Timeline

- 2020-05-08 John Newbery reports his finding by email
- 2020-05-12 John Newbery opens #18962
- 2020-05-14 #18962 is merged
- 2020-06-03 Bitcoin Core version 0.20.0 is released with a fix
- 2021-09-13 The last vulnerable Bitcoin Core version (0.19.x) goes EOL
- 2024-07-03 Public disclosure.

{% include references.md %}
