---
title: Tor Onion-Location meta tag
name: tor-onion-location
id: en-tor-onion-location
lang: en
type: posts
layout: post

## If this is a new post, reset this counter to 1.
version: 1

## Only true if release announcement or security annoucement.  English posts only
announcement: 0

excerpt: >
  Bitcoincore.org now advertises it's own .onion website and prompts visitors to
switch to the version delivered using the Tor Onion service.
---
Bitcoincore.org now advertises it's own [.onion website] and prompts visitors to
switch to the version delivered using the Tor Onion service for improved security and privacy.

Thanks to the new [Onion-Location meta tag], when a user visits Bitcoincore.org, Tor Browser
will display an information pill telling them that there's a more secure version of the website,
and the user will be asked to opt-in to upgrade to the onion service on their first use.
If the user already opted-in for their network security upgrade, they will get directly to the onion site.

[.onion website]: /en/2020/03/27/hidden-service/
[Onion-Location meta tag]: https://github.com/bitcoin-core/bitcoincore.org/pull/706
