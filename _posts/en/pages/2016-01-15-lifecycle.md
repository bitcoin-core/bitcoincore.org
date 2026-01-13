---
title: Software Life Cycle
name: software-life-cycle
id: en-software-life-cycle
permalink: /en/lifecycle/
layout: page
type: pages
lang: en
version: 2
---
{% include toc.html %}

This document describes the life-cycle of the Bitcoin Core software package released by the Bitcoin Core project. It is in line with standard maintenance policy across commercial software.  

## Versioning

Bitcoin Core releases are versioned as follows: MAJOR.MINOR, and release candidates are suffixed with rc1, rc2 etc.

We aim to make a major release every 6 months. These will be numbered 29.0, 30.0 etc.

We will provide minor ("maintenance") releases that fix bugs (security and otherwise) for each major release. These
will be numbered 29.3, 30.1, etc. We will not introduce major new features in maintenance releases (besides consensus
rules change, see below).

## Consensus rules

Proposals to change consensus rules are always shipped first in maintenance versions such as 22.2, 23.1 etc. This makes it easier for enterprise users to assess and test the proposal because of its smaller changeset compared to a major release. It also allows users who follow a more conservative upgrade path to adopt consensus rule changes in a more timely manner.

## Maintenance period

We always maintain the latest three major versions. When a new major version is released, the oldest one falls out of
the maintenance window and becomes "End of Life". For example, if the last major release is 30.0, then 29.x and 28.x are
also considered maintained. Once 31.0 is released, 28.x becomes "End of Life". The threshold for backporting a change
to an older major version increases as it ages.

Major versions that are "End of Life" do not generally receive security fixes. For more about our policy on security
fixes, see our [security advisories][] page. We recommend running the latest maintenance release of the most recent
major version you are able to upgrade to.

## Schedule

Once EOL is reached, you will need to upgrade to a newer version.

| Version | Release Date | End of Life |
|---------|--------------|-------------|
{% include posts/maintenance-table.md %}

\* _We aim to make a major release every 6-7 months_

_TBA: to be announced_

## Protocol versioning

The description above only describes Bitcoin Core software releases. Many other parts of the Bitcoin system contain their own versions.  A few examples:

- Every **transaction** contains a version number.
- The **P2P network protocol** uses version numbers to allow nodes to announce what features they support.
- Bitcoin Core's **built-in wallet** has its own internal version number.

These version numbers are deliberately decoupled from Bitcoin Core's version number as the Bitcoin Core project either has no direct control over them (as is the case with blocks and transactions), or tries to maintain compatibility with other projects (as is the case with the network protocol), or allows for the possibility that no major changes will be made in some releases (as is sometimes the case with the built-in wallet).

The consensus protocol itself doesn't have a version number.

## Relationship to SemVer

Bitcoin Core software versioning does not follow the [SemVer][] optional versioning standard, but its release versioning is superficially similar.  SemVer was designed for use in normal software libraries where individuals can choose to upgrade the library at their own pace, or even stay behind on an older release if they don't like the changes.

Parts of Bitcoin, most notably the consensus rules, don't work that way.  In order for a new consensus rule to go into effect, it must be enforced by some number of miners, full nodes, or both; and once it has gone into effect, software that doesn't know about the new rule may generate or accept invalid transactions (although upgrades are designed to prevent this from happening when possible).

For this reason, Bitcoin Core deviates from SemVer for changes to consensus rules and other updates where network-wide adoption is necessary or desirable.  Bitcoin Core releases these changes as maintenance releases (`x.y`) instead of as major releases (`x.0`); this minimizes the size of the patch in order to make it easy for as many people as possible to inspect it, test it, and deploy it.  It also makes it possible to backport the same patch to multiple previous major releases, further increasing the number of users who can easily upgrade, although there are not always enough volunteers to manage this.

[SemVer]: https://semver.org/
[bitcoin-transifex-link]: https://explore.transifex.com/bitcoin/bitcoin/
[security advisories]: /en/security-advisories
