---
title: CVE-2015-20111 - Remote code execution due to bug in miniupnpc
name: blog-disclose-upnp-rce
id: en-blog-disclose-upnp-rce
lang: en
type: advisory
layout: post

## If this is a new post, reset this counter to 1.
version: 1

## Only true if release announcement or security annoucement. English posts only
announcement: 1

excerpt: >
  A bug in the miniupnpc library could have led to a remote code execution in Bitcoin Core. A fix was released on October 15th, 2015 in Bitcoin Core 0.11.1.
---

A buffer overflow enabling a significant data leak was discovered in `miniupnpc`. Combined with the then
recently-disclosed CVE-2015-6031 it enabled an RCE in `miniupnpc` which could have led to an RCE
in Bitcoin Core. This was fixed in [Bitcoin Core 0.12](https://bitcoincore.org/en/releases/0.12.0/),
released in February 2016.

This issue is considered **Medium** severity.

## Details

[CVE-2015-6031](https://nvd.nist.gov/vuln/detail/CVE-2015-6031), disclosed in September 2015, made
it possible for a malicious UPnP server to remotely crash a Bitcoin Core process on the local
network at startup. See [here](https://nvd.nist.gov/vuln/detail/CVE-2015-6031) for details. The fix
was [pulled in Bitcoin Core](https://github.com/bitcoin/bitcoin/pull/6789) and released in [version
0.11.1](https://bitcoincore.org/en/releases/0.11.1/), released in October 2015. UPnP was then
[turned off by default](https://github.com/bitcoin/bitcoin/pull/6795).

CVE-2015-6031 disclosed a buffer overflow, which in addition to enabling a remote crash could have
made it possible to remotely execute code on a victim's machine. While investigating this
possibility, Wladimir J. Van Der Laan found another buffer overflow in `miniupnpc` which enabled a
significant data leak. This was [fixed by Wladimir in
`miniupnpc`](https://github.com/miniupnp/miniupnp/pull/157) in commit
`4c90b87ce3d2517097880279e8c3daa7731100e6`. The fix was then [pulled into Bitcoin
Core](https://github.com/bitcoin/bitcoin/pull/6980) and released as part of version 0.12.

This data leak did not disclose secret information (such as the wallet's private keys) directly. But
combined with another stack overflow (such as the one disclosed in CVE-2015-6031) this made it
possible to trigger a remote code execution. Wladimir demonstrated this against Ubuntu's `miniupnpc`
version `1.6-precise`. The specific approach used in this exploit was however not directly portable
to Bitcoin Core.

## Attribution

Credits go to Aleksandar Nikolic for identifying CVE-2015-0035 and to Wladimir J. Van Der Laan for
investigating its impact and discovering the second buffer overflow.

## Timeline

- 2015-09-15 CVE-2015-0035 is
  [fixed](https://github.com/miniupnp/miniupnp/commit/79cca974a4c2ab1199786732a67ff6d898051b78) and
  [disclosed](https://talosintelligence.com/vulnerability_reports/TALOS-2015-0035/).
- 2015-10-09 [PR #6789](https://github.com/bitcoin/bitcoin/pull/6789) is merged in Bitcoin Core
- 2015-10-14 Wladimir's remote code execution by leveraging the second buffer overflow is disclosed
  to Ubuntu security and Bitcoin developers.
- 2015-10-15 Bitcoin Core 0.11.1 [is
  released](https://lists.linuxfoundation.org/pipermail/bitcoin-dev/2015-October/011545.html)
- 2015-10-26 The fix for the second buffer overflow [is
  merged](https://github.com/miniupnp/miniupnp/pull/157) into `miniupnpc`.
- 2015-12-18 The fix is [pulled into Bitcoin Core](https://github.com/bitcoin/bitcoin/pull/6980).
- 2016-02-23 Bitcoin Core version 0.12 [is
  released](https://lists.linuxfoundation.org/pipermail/bitcoin-dev/2016-February/012456.html).
- 2017-03-08 The last vulnerable Bitcoin Core Version (0.11.x) goes EOL
- 2024-07-03 Public disclosure

{% include references.md %}
