---
title: Software Life Cycle
name: software-life-cycle
id: en-software-life-cycle
permalink: /en/lifecycle/
layout: page
type: pages
lang: en
share: false
version: 2
---
{% include _toc.html %}

This document describes the life-cycle of the Bitcoin Core software package released by the Bitcoin Core project. It is in line with standard maintenance policy across commercial software.  

## Versioning

Bitcoin Core releases are versioned as follows: 0.MAJOR.MINOR, and release candidates are suffixed with rc1, rc2 etc.

## Major releases

We aim to make a major release every 6-7 months.

These will be numbered 0.11.0, 0.12.0 etc.

## Maintenance releases

We will provide maintenance "minor releases" that fix bugs within the major releases. As a general rule we do not introduce major new features in a maintenance release (except for consensus rules). However, we may add minor features where necessary, and we will back-port consensus rule changes such as soft forks.

Minor releases will be numbered 0.11.1, 0.11.2, 0.12.1, 0.12.2 etc.

## Consensus rules

Proposals to change consensus rules are always shipped first in maintenance versions such as 0.11.2, 0.12.1 etc. This makes it easier for enterprise users to assess and test the proposal because of its smaller changeset compared to a major release. It also allows users who follow a more conservative upgrade path to adopt consensus rule changes in a more timely manner.

## Maintenance period

We maintain the major versions until their "Maintenance End". We generally maintain the current and previous major release.
So if the current release is 0.13, then 0.12 is also considered maintained.
Once 0.14 is released, then 0.12 would be considered at it's "Maintenance End".
The older the major release, the more critical issues have to be to get backported to it, and the more to warrant a new minor release.
Once software has reached the "Maintenance End" period it will only receive critical security fixes until the EOL date.
After EOL, users must upgrade to a later version to receive security updates, even though the community may provide fixes for critical issues on a best effort basis.
Generally, it is recommended to run the latest maintanance release (point release) of the current or previous major version.

Please note that minor versions get bugfixes, translation updates, and soft forks. Translation on [Transifex][bitcoin-transifex-link] is only open for the last two major releases.

For example, major version 0.9 was released on 2014-03-19 and we provided maintenance fixes (point releases) until 2015-06-16.
Critical security issues would still be continue to be fixed until the End-Of-Life "EOL" date of 2016-02-28.
However, to take advantage of bug fixes, you would have to upgrade to a later major version.

## Schedule

Once EOL is reached, you will need to upgrade to a newer version.

| Version | Release Date | Maintenance End | End of Life |
|---------|--------------|-----------------|-------------|
{% include posts/_maintenance-table.md %}

\* _We aim to make a major release every 6-7 months_

_TBA: to be announced_

## Protocol versioning

The description above only describes Bitcoin Core software releases. Many other parts of the Bitcoin system contain their own versions.  A few examples:

- Every **transaction** contains a version number.
- The **P2P network protocol** uses version numbers to allow nodes to announce what features they support.
- Bitcoin Core's **built-in wallet** has its own internal version number.

These versions numbers are deliberately decoupled from Bitcoin Core's version number as the Bitcoin Core project either has no direct control over them (as is the case with blocks and transactions), or tries to maintain compatibility with other projects (as is the case with the network protocol), or allows for the possibility that no major changes will be made in some releases (as is sometimes the case with the built-in wallet).

The consensus protocol itself doesn't have a version number.

## Relationship to SemVer

Bitcoin Core software versioning does not follow the [SemVer][] optional versioning standard, but its release versioning is superficially similar.  SemVer was designed for use in normal software libraries where individuals can choose to upgrade the library at their own pace, or even stay behind on an older release if they don't like the changes.

Parts of Bitcoin, most notably the consensus rules, don't work that way.  In order for a new consensus rule to go into effect, it must be enforced by some number of miners, full nodes, or both; and once it has gone into effect, software that doesn't know about the new rule may generate or accept invalid transactions (although upgrades are designed to prevent this from happening when possible).

For this reason, Bitcoin Core deviates from SemVer for changes to consensus rules and other updates where network-wide adoption is necessary or desirable.  Bitcoin Core releases these changes as maintenance releases (`0.x.y`) instead of as major releases(`0.x.0`); this minimizes the size of the patch in order to make it easy for as many people as possible to inspect it, test it, and deploy it.  It also makes it possible to backport the same patch to multiple previous major releases, further increasing the number of users who can easily upgrade, although there are not always enough volunteers to manage this.

[SemVer]: https://semver.org/
[bitcoin-transifex-link]: https://www.transifex.com/bitcoin/bitcoin/
