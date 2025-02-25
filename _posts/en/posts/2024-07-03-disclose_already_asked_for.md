---
title: CVE-2024-52913 - Censorship due to transaction re-request handling
name: blog-disclose-already-asked-for
id: en-blog-disclose-already-asked-for
lang: en
type: advisory
layout: post

## If this is a new post, reset this counter to 1.
version: 2

## Only true if release announcement or security annoucement. English posts only
announcement: 1

excerpt: >
  Nodes could be prevented from seeing specific unconfirmed transactions by a malicious peer. A fix was released on January 14th, 2021 in Bitcoin Core 0.21.0.
---

An attacker could prevent a node from seeing a specific unconfirmed transaction.

This issue is considered **Medium** severity.

## Details

Before this issue was fixed in PR 19988, the "g_already_asked_for" mechanism was used to schedule `GETDATA` requests for transactions. The `SendMessages()` function would send out `GETDATA`s for transactions recently announced by peers, remembering when that request was sent out in `g_already_asked_for`. However, this `g_already_asked_for` was a "limitedmap" data structure, with a bounded size that would forget the oldest entries if it reaches 50000 entries. This makes the following attack possible:
* The attacker is the first to announce a legitimate transaction T to the victim.
* The victim requests T from the attacker using `GETDATA`.
* The attacker does not respond to `GETDATA` until close to the time when the victim would request T from other peers (~60 seconds).
* Then, the attacker carefully spams the victim with bogus announcements, causing the victim's `g_already_asked_for` to evict T.
* The attacker announces T again to the victim (due to how the queueing works in `m_tx_process_time`, this does not need to be timed particularly accurately).
* The victim, not finding T in `g_already_asked_for` will treat it as a new announcement, sending a new `GETDATA` for it to the attacker.
* The attacker again does not respond to `GETDATA`.
* etc.

This way, the attacker can prevent the victim from ever requesting the transaction from anyone but the attacker.

## Attribution

Responsibly disclosed by John Newbery, claiming discovery by Amiti Uttarwar and him.

## Timeline

- 2020-04-03 John Newbery reports the bug in an email to Suhas Daftuar and others
- 2020-05-08 John Newbery suggests an approach to fixing the bug
- 2020-09-21 Pieter Wuille opens [PR #19988](https://github.com/bitcoin/bitcoin/pull/19988) as a comprehensive approach to fixing this and other bugs
- 2020-10-14 Pieter's PR is merged
- 2021-01-14 Bitcoin Core version 0.21.0 is released with a fix
- 2022-04-25 The last vulnerable Bitcoin Core version (0.20.x) goes EOL
- 2024-07-03 Public disclosure

{% include references.md %}
