---
title: CVE-2024-52912 - Netsplit due to timestamp adjustment
name: blog-disclose-timestamp-overflow
id: en-blog-disclose-timestamp-overflow
lang: en
type: advisory
layout: post

## If this is a new post, reset this counter to 1.
version: 1

## Only true if release announcement or security annoucement. English posts only
announcement: 1

excerpt: >
  A node could be split from the network when attacked by its first 200 peers. A fix was released on January 15th, 2021 in Bitcoin Core version 0.21.0.
---

Disclosure of the details of an integer overflow bug which risked causing a network split, a fix for
which was released on January 15th, 2021 in Bitcoin Core version 0.21.0.

This issue is considered **Medium** severity.

## Technical details

A network split vulnerability resulted from two separate bugs in the processing code of `version`
messages:
* Signed-integer overflow when calculating the time offset for newly connecting peers.
* abs64 logic bug (`abs64(std::numeric_limits<int64_t>::min()) ==
  std::numeric_limits<int64_t>::min()`), resulting in a bypass of the maximum time adjustment limit.

The two bugs allow an attacker to force a victims adjusted time (`system time + network time
offset`) to be skewed such that any new blocks are rejected for having a timestamp that is dated too
far in the future. It should be noted that this attack assumes the attacker is among the first 200
peers to connect to the victim, as only the time offsets from those initial connections are factored
into adjusted time.

## Attribution

Credit goes to [practicalswift](https://github.com/practicalswift) for discovering and providing the
initial fix for the vulnerability, and Pieter Wuille for the fix as well as general cleanup to the
at-risk code.

## Timeline

* 2020-10-10 Initial report send to security@bitcoincore.org
* 2020-10-13 Fix merged into Bitcoin Core (https://github.com/bitcoin/bitcoin/pull/20141)
* 2021-01-15 v0.21.0 released
* 2022-04-25 The last vulnerable Bitcoin Core version (0.20.x) goes EOL
* 2024-07-03 Public disclosure

{% include references.md %}
