---
title: CVE-2024-52917 - Infinite loop bug in the miniupnp dependency
name: blog-disclose-miniupnp-bug-impact
id: en-blog-disclose-miniupnp-bug-impact
lang: en
type: advisory
layout: post

## If this is a new post, reset this counter to 1.
version: 1

## Only true if release announcement or security annoucement. English posts only
announcement: 1

excerpt: >
  Nodes could be crashed by a malicious UPnP device on the local network. A fix was released on September 14th, 2021 in Bitcoin Core v22.0.
---

Disclosure of the impact of an infinite loop bug in the miniupnp dependency on
Bitcoin Core, a fix for which was released on September 14th, 2021 in Bitcoin
Core version v22.0.

This issue is considered **Low** severity.

## Details

Miniupnp, the UPnP library used by Bitcoin Core, would be waiting upon
discovery for as long as it receives random data from a device on the network.
In addition it would allocate memory for every new device information. An
attacker on the local network could pretend to be a UPnP device and keep
sending bloated M-SEARCH replies to the Bitcoin Core node until it runs out of
memory.

Only users running with the <code>-miniupnp</code> option would have been
affected by this bug as Miniupnp is otherwise turned off by default.

## Attribution

Credit goes to Ronald Huveneers for reporting the infinite loop bug to the
miniupnp project, and to Michael Ford (Fanquake) for the report to the Bitcoin
Core project along with a PoC exploit to trigger an OOM and a pull request to
bump the dependency (containing the fix).

## Timeline

* 2020-09-17 - Initial report of infinite loop bug to miniupnp by Ronald Huveneers
* 2020-10-13 - Initial report sent to security@bitcoincore.org by Michael Ford
* 2021-03-23 - Fix is merged (https://github.com/bitcoin/bitcoin/pull/20421)
* 2021-09-13 - v22.0 is released
* 2024-07-31 - Public disclosure

{% include references.md %}
