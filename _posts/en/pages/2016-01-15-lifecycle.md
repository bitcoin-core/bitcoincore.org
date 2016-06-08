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

Bitcoin Core versions releases as follows 0.MAJOR.MINOR. Release candidates are suffixed with rc1, rc2 etc.

## Major releases

We aim to make a major release every 6-7 months.

These will be numbered 0.11.0, 0.12.0 etc.

## Maintenance releases

We will provide maintenance "minor releases" that fix bugs within the major releases. As a general rule we do not introduce major new features in a maintenance release (consensus rule are exempt from this rule). However, we may add minor features where necessary, and we will back-port consensus rule changes such as soft forks.

Minor releases will be numbered 0.11.1, 0.11.2, 0.12.1, 0.12.2 etc.

## Soft forks

Soft fork consensus rule proposals are always shipped first in maintenance versions such as 0.11.2, 0.12.1 etc. This make it easier for enterprise users to assess and test the proposal due to smaller changeset than in a major release. 

## Maintenance period

We maintain the major versions until their "Maintenance End". We generally maintain the current and previous major release. So if the current release is 0.12, then 0.11 is also considered maintained. Once 0.13 is released, then 0.11 would be considered at it's "Maintenance End". The older the major release, the more critical issues have to be to get backported to it, and the more to warrant a new minor release. Once software has reached the "Maintenance End" period it will only receive critical security fixes until the EOL date. After EOL, users must upgrade to a later version to receive security updates.

Please note that minor versions get bugfixes, translation updates, and soft forks. Translation on [Transifex][bitcoin-transifex-link] is only open for the last two major releases.

For example, major version 0.9 was released on 2014-03-19 and we provided maintenance fixes (point releases) until 2015-02-16. Critical security issues would still be continue to be fixed until the End-Of-Life "EOL" date of 2016-02-31. However, to take advantage of bug fixes, you would have to upgrade to a later major version.

## Schedule

Once EOL is reached, you will need to upgrade to a newer version.

| Version | Release Date | Maintenance End | End of Life |
|---------|--------------|-----------------|-------------|
{% include posts/_maintenance-table.md %}

\* _We aim to make a major release every 6-7 months_

_TBA: to be announced_

[bitcoin-transifex-link]: https://www.transifex.com/bitcoin/bitcoin/
