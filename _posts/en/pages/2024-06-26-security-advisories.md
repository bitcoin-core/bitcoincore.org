---
title: Security Advisories
name: Security Advisories
id: en-security-advisories
permalink: /en/security-advisories/
layout: page
type: pages
lang: en
version: 1
---
{% include toc.html %}

This page summarizes policies in relation to disclosing vulnerabilities in Bitcoin Core,
as well as provides a summary of historical Security Advisories.

## Policy

All vulnerabilities should be reported to security@bitcoincore.org (see
[SECURITY.md](https://github.com/bitcoin/bitcoin/blob/master/SECURITY.md) for
details). When reported, a vulnerability will be assigned a severity category.
We differentiate between 4 classes of vulnerabilities:
* **Critical**: Bugs that threaten the fundamental security and integrity of the entire Bitcoin
  network. These are bugs that allow for coin theft at the protocol level, the
  creation of coins outside of the specified issuance schedule, or permanent,
  network-wide chain splits.
  <details markdown="1">

  <summary>
  Examples
  </summary>

  * A bug allowing inflating the money supply by spending the same transaction
    output twice within a block ([CVE-2018-17144](/en/2018/09/20/notice/)).
  * A consensus failure where nodes running older software rejected a block that
    newer software accepted due to an underlying database limit, causing a
    network-wide chain split ([BIP
    50](https://github.com/bitcoin/bips/blob/master/bip-0050.mediawiki)).

  </details>
* **High**: Bugs with a significant impact on affected nodes or the network. These are
  typically exploitable remotely under default configurations and can cause
  widespread disruption.
  <details markdown="1">

  <summary>
  Examples
  </summary>

  * A remotely triggerable crash that could take many nodes offline
    ([CVE-2024-35202](/en/2024/10/08/disclose-blocktxn-crash/)).
  * A denial-of-service attack that causes a node to stall for an extended
    period, preventing it from processing new transactions and blocks
    ([CVE-2024-52914](/en/2024/07/03/disclose-orphan-dos/)).
  * A memory exhaustion vulnerability that could be triggered remotely to crash
    nodes by having them store an excessive amount of block headers
    ([CVE-2019-25220](/en/2024/09/18/disclose-headers-oom/)).

  </details>
* **Medium**: Bugs that can noticeably degrade the network's or a node's performance or
  functionality, but are limited in their scope or exploitability. These might
  require special conditions to trigger, such as non-default settings, or result
  in service degradation rather than a complete node failure.
  <details markdown="1">

  <summary>
  Examples
  </summary>

  * A potential Remote Code Execution (RCE) vulnerability on the local network
    that is only exploitable if a non-default feature like UPnP is enabled
    ([CVE-2015-20111](/en/2024/07/03/disclose_upnp_rce/)).
  * A peer could hinder block propagation by sending mutated blocks, delaying the
    reception of new blocks for a node
    ([CVE-2024-52921](/en/2024/10/08/disclose-mutated-blocks-hindering-propagation/)).
  * An attacker announcing a block to a node and then failing to provide it,
    causing the victim node to wait for up to 10 minutes before being able to
    fetch it from another peer
    ([CVE-2024-52922](/en/2024/11/05/cb-stall-hindering-propagation/)).

  </details>
* **Low**: Bugs that are challenging to exploit or have a minor impact on a node's
  operation. They might only be triggerable under non-default configurations or
  from the local network, and do not pose an immediate or widespread threat.
  <details markdown="1">

  <summary>
  Examples
  </summary>

  * A malformed `getdata` message could cause a peer connection to enter an
    infinite loop, consuming CPU but not affecting the node's ability to process
    blocks or handle other peer connections
    ([CVE-2024-52920](/en/2024/07/03/disclose-getdata-cpu/)).
  * A bug in a dependency that could crash a node, but only if a non-default
    feature like UPnP is enabled
    ([CVE-2024-52917](/en/2024/07/31/disclose-upnp-oom/)).
  * A bug that could crash a node, but is extremely difficult to exploit
    ([CVE-2024-52919](/en/2025/04/28/disclose-cve-2024-52919/)).

  </details>

**Low** severity vulnerabilities will be disclosed 2 weeks after the release of a major version
containing the fix. **Medium** and **High** severity vulnerabilities will be disclosed 2 weeks after
the last affected release goes [End of Life](/en/lifecycle/) (approximately a year after a major
version containing the fix was first released).

A pre-announcement will be made two weeks prior to releasing the details of a vulnerability. This
pre-announcement will coincide with the release of a new major version and contain the number of
fixed vulnerabilities and their severity levels.

**Critical** bugs are not considered in the standard policy, as they would most likely require an
  ad-hoc procedure. Also, a bug may not be considered a vulnerability at all. Any reported issue
  may also be considered serious, yet not require embargo.

## Past Security Advisories

{% assign advisories=site.posts | where:"lang", 'en' | where:"type", 'advisory' | sort: "date" | reverse %}
{% for advisory in advisories %}
{% assign post=advisory %}
  <article>
    <h2><a href="{{ post.url }}" title="{{ post.title | xml_escape }}">{{ post.title }}</a></h2>
    <p>{{ post.excerpt | markdownify | strip_html | truncate: 200 }}</p>
  </article>
{% endfor %}
