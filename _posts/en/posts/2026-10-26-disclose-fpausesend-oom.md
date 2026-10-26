---
title: Slow unbounded memory growth on paused pre-BIP31 peers
name: blog-disclose-fpausesend-oom
id: en-blog-disclose-fpausesend-oom
lang: en
type: advisory
layout: post

## If this is a new post, reset this counter to 1.
version: 1

## Only true if release announcement or security annoucement. English posts only
announcement: 1

excerpt: >
  Bitcoin Core queues announcements for all peers, including paused ones. Pre-BIP31 peers could stay paused forever and let the queue grow unboundedly.
---

Bitcoin Core queues announcements for all peers, including paused ones. Pre-BIP31 peers could stay
paused forever and let the queue grow unboundedly. A fix for this issue was released in Bitcoin Core
versions 29.5, 30.4, 31.2 and 32.1.

This issue is considered **Medium** severity and **easily discoverable**.

## Details

Bitcoin Core implements backpressure for slow to read peers. Once the queue of messages to send a
peer grows past a certain limit, it will mark that peer as paused for send purposes. However, it
will still add some announcements to the queue, in particular the frequent transaction `INV`s.

This is normally not an issue, as peers which don't respond to a (nonced) ping message within 20
minutes get disconnected. However, this disconnection is not enabled for peers that don't support
BIP 31 nonced ping messages. Therefore, a malicious peer could pretend to not support those and
strategically delay reading messages to cause the queue to slowly grow without limit.

Memory usage could grow at a rate of up to 2MB per hour per inbound peer, and 5MB per hour per
outbound peer. Nodes running with `-blocksonly` are much less affected, as transaction announcements
make up the bulk of the queued messages.

## Attribution

This issue was discovered through an automated LLM scan, as part of Elias Rohrer's [Project
Loupe](https://github.com/project-loupe/loupe) at Spiral. The Loupe finding was validated and
surfaced to the Bitcoin Core team by Gregory Sanders.

This issue was independently re-discovered by L0rinc.

## Timeline

- 2026-09-15 Project Loupe identifies the issue and notifies Greg Sanders, who validates the report
  and notifies the Bitcoin Core security team on the same day.
- 2026-09-21 L0rinc privately notifies Gregory Sanders, who replies that issue was already reported.
- 2026-10-26 Antoine Poinsot opens PR [#36667](https://github.com/bitcoin/bitcoin/pull/36667) to
  mitigate the vulnerability.
- 2026-10-26 PR #36667 is merged and backported to versions 29, 30, 31, and 32.
- 2026-10-27 Bitcoin Core 29.5, 30.4, 31.2 and 32.1 are released with a fix.
- 2026-10-27 Public disclosure.

{% include references.md %}
