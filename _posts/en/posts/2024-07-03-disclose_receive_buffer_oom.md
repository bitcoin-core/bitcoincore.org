---
title: Disclosure of CVE-2015-3641
name: blog-disclose-receive-buffer-oom
id: en-blog-disclose-receive-buffer-oom
lang: en
type: advisory
layout: post

## If this is a new post, reset this counter to 1.
version: 1

## Only true if release announcement or security annoucement. English posts only
announcement: 1

excerpt: >
  Attackers sending large incomplete messages would cause high memory usage. A fix was released on April 27th, 2015 in Bitcoin Core 0.10.1.
---

A node could be forced to allocate large buffers when receiving a message, which could be leveraged to remotely crash it by OOM.

This issue is considered **Medium** severity.

## Details

Without a tighter bound, received messages' size was limited by the maximum serialized message size
of 32 MiB. An attacker could force a node to allocate this much RAM per connection, which may lead
to an OOM.

[PR #5843](https://github.com/bitcoin/bitcoin/pull/5843) reduced the size P2P messages can have
before receiving the payload. This reduces the per-peer receive buffer memory size a malicious peer
can cause. The PR reduced the number from 32 MiB to 2 MiB, which was later increased back to 4 MB as
part of the Segwit BIP144 changes.

## Attribution

Reported to Greg Maxwell by bitcointalk user Evil-Knievel. Fixed by Pieter Wuille.

## Timeline

- 2015-02-05 Evil-Knievel reports the vulnerability to Greg Maxwell through bitcointalk private messages.
- 2015-??-?? `CVE-2015-3641` is registered for it.
- 2015-03-01 [PR #5843](https://github.com/bitcoin/bitcoin/pull/5843) is opened to fix it.
- 2015-03-06 PR #5843 is merged.
- 2015-03-09 The fix is backported to version 0.10.1.
- 2015-04-27 Bitcoin Core version [0.10.1 is released](https://lists.linuxfoundation.org/pipermail/bitcoin-dev/2015-April/007828.html) with a fix.
- 2015-06-25 A disclosure is [pre-announced](https://lists.linuxfoundation.org/pipermail/bitcoin-dev/2015-June/009135.html).
- 2015-07-07 Disclosure is [postponed](https://lists.linuxfoundation.org/pipermail/bitcoin-dev/2015-July/009362.html).
- 2016-08-23 The last vulnerable Bitcoin Core Version (0.10.x) goes EOL
- 2024-07-03 Public disclosure.

{% include references.md %}
