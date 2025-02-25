---
title: CVE-2024-52920 - DoS using huge GETDATA messages
name: blog-disclose-getdata-cpu
id: en-blog-disclose-getdata-cpu
lang: en
type: advisory
layout: post

## If this is a new post, reset this counter to 1.
version: 1

## Only true if release announcement or security annoucement. English posts only
announcement: 1

excerpt: >
  A malformed `GETDATA` message could trigger 100% CPU usage on the receiving node. A fix was released on June 3rd, 2020 in Bitcoin Core 0.20.0.
---

A malformed `GETDATA` message could trigger an infinite loop on the receiving node, using 100% of
the CPU allocated to this thread and not making further progress on this connection.

This issue is considered **Low** severity.

## Details

Before Bitcoin Core 0.20.0, an attacker (or buggy client, even) could send us a `GETDATA` message
that would cause our net_processing thread to start spinning at 100%, and not make progress
processing messages for the attacker peer anymore. It would still make progress processing messages
from other peers, so it is just a CPU DoS with low impact beyond that (not making progress for
attacker peers is a non-issue). It also increases per-peer long-term memory usage up by 1.5 MB per
attacker peer.

John Newbery opened [PR #18808](https://github.com/bitcoin/bitcoin/pull/18808) to fix this issue by
only disclosing the lack of progress.

## Attribution

Credits to John Newbery for finding this bug, responsibly disclosing it and fixing it.

## Timeline

- 2020-04-29 John Newbery opens #18808
- 2020-05-08 John Newbery reports his finding by email
- 2020-05-12 #18808 is merged
- 2020-06-03 Bitcoin Core version 0.20.0 is released with a fix
- 2021-09-13 The last vulnerable Bitcoin Core version (0.19.x) goes EOL
- 2024-07-03 Public disclosure.

{% include references.md %}
