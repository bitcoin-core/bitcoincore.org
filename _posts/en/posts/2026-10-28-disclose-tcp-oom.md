---
title: CVE-2026-117350 - Unbounded memory growth through TCP urgent data on non-Linux systems
name: blog-disclose-tcp-oom
id: en-blog-disclose-tcp-oom
lang: en
type: advisory
layout: post
redirect_from:
  - /en/2026/10/28/disclose-tcp-oom
  - /en/2026/10/28/disclose-cve-2026-117350

## If this is a new post, reset this counter to 1.
version: 1

## Only true if release announcement or security annoucement. English posts only
announcement: 1

excerpt: >
  On platforms other than Linux, a malicious peer could cause a node to run out of memory by sending TCP urgent (out-of-band) data.
---

On platforms other than Linux, a malicious peer could cause a node to run out of memory by sending TCP
urgent (out-of-band) data. A fix for this issue was released in Bitcoin Core versions 29.5, 30.4,
31.2 and 32.1.

This issue is considered **Medium** severity and **easily discoverable**.

## Details

Bitcoin Core uses a different backend (based on `select()`) to monitor socket events on platforms
other than Linux (where it uses `poll()`). This issue only affects the `select()` code path. Bitcoin
Core users on **Linux** are therefore **not affected** by this issue.

Bitcoin Core limits the amount of data it will queue for processing per peer, by temporarily pausing
reads when the queue exceeds a certain size. The `select()` socket monitoring backend watched
sockets for exceptional conditions regardless of whether reads were paused. In handling that event,
a read on the socket would be triggered, which could queue up more data for processing for that
peer, without limit.

That event can be triggered by sending TCP urgent data. This allows a paused peer to bypass the
backpressure mechanism and queue an unbounded amount of data for processing. In practice, this means
a malicious peer can trivially crash any Windows, macOS or BSD node it is connected to.

## Attribution

This issue was discovered through an automated LLM scan, as part of Elias Rohrer's [Project
Loupe](https://github.com/project-loupe/loupe) at Spiral. The Loupe finding was validated and
surfaced to the Bitcoin Core team by Gregory Sanders.

## Timeline

- 2026-10-23 Project Loupe identifies the issue and notifies Greg Sanders, who validates the report
  and notifies the Bitcoin Core security team on the same day.
- 2026-10-26 Antoine Poinsot opens PR [#36665](https://github.com/bitcoin/bitcoin/pull/36665) to fix
  this issue.
- 2026-10-26 PR #36665 is merged and backported to versions 29, 30, 31, and 32.
- 2026-10-27 Bitcoin Core 31.2 and 32.1 are released with a fix
- 2026-10-28 Bitcoin Core 29.5, 30.4 are released with a fix
- 2026-10-28 Public disclosure.

{% include references.md %}
