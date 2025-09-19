---
title: Bitcoin Core development and transaction relay policy
permalink: /en/2025/06/06/relay-statement/
lang: en
id: en-relay-statement
name: relay-statement
type: posts
layout: post
version: 1
---

We'd like to share our view on the relationship between Bitcoin Core development and transaction relay
policy on the network.

Bitcoin is a network that is defined by its users, who have ultimate freedom in choosing what
software they use (fully-validating or not) and implementing whatever policies they desire. Bitcoin
Core contributors are not in a position to mandate what those are. One way this is reflected is by
our long-running practice of avoiding auto-updating in the software. This means that no entity can
unilaterally push out changes to Bitcoin Core users: changes must be made by users choosing to
adopt new software releases themselves, or if they so desire, different software. Being free to run
any software is the network's primary safeguard against coercion.

As Bitcoin Core developers we also consider it our responsibility to make our software work as
efficiently and reliably as possible for its purpose, namely validating and relaying blocks and
transactions in the Bitcoin peer-to-peer network, so that Bitcoin succeeds as a decentralized digital
currency. With regards to transaction relay, this may include adding policies for denial of service (DoS)
protection and fee assessment, but not blocking relay of transactions that have sustained economic
demand and reliably make it into blocks. The goals of transaction relay include:

* predicting what transactions will be mined (for example for fee estimation or fee bumping, but it
  is also the basis for many DoS protection strategies inside of node software);
* speeding up block propagation for the transactions we expect to be mined. Reduced latency helps
  prevent large miners from gaining unfair advantages;
* helping miners learn about fee-paying transactions (so they do not need to rely on out-of-band
  transaction submission schemes that undermine mining decentralization).

**Knowingly refusing to relay transactions that miners would include in blocks anyway forces users into
alternate communication channels, undermining the above goals.**

It is the case that transaction acceptance rules have been used effectively in the past to
discourage the development of use cases that used block space inefficiently while doing so was very
cheap. However this can only be effective while both users and miners are satisfied with whatever
alternatives exist. When that is no longer the case, and an economically viable use case develops
that would conflict with policy rules, users and miners can directly collaborate to avoid any
external attempt to impose restrictions on their activities. In fact, the ability to do precisely
that is an important aspect of Bitcoin's censorship resistance, and other node software with
preferential peering has also shown that circumventing filters of the vast majority of the nodes
is relatively easy. Given that, we believe it is better for Bitcoin node software to aim to have a
realistic idea of what will end up in the next block, rather than attempting to intervene between
consenting transaction creators and miners in order to discourage activity that is largely harmless
at a technical level.

**This is not endorsing or condoning non-financial data usage, but accepting
that as a censorship-resistant system, Bitcoin can and will be used for use cases not everyone
agrees on.**

While we recognize that this view isn't held universally by all users and developers, it is our
sincere belief that it is in the best interest of Bitcoin and its users, and we hope our users agree.
We will continue to apply our best judgment as developers in aligning transaction acceptance rules
with Bitcoin's long-term health and miners' rational self-interest, including specific
technical reasons such as upgrade safety, efficient block building, and node DoS attacks.

Signed,

(List of contributors who support this letter)

* Andrew Toth
* Antoine Poinsot
* Anthony Towns
* Ava Chow
* b10c
* Bruno Garcia
* David Gumberg
* fjahr
* Gloria Zhao
* Gregory Sanders
* hodlinator
* ismaelsadeeq
* Josie Baker
* kevkevinpal
* l0rinc
* Marco De Leon
* Martin Zumsande
* Matthew Zipkin
* Michael Ford
* Murch
* Niklas Gögge
* pablomartin4btc
* Pieter Wuille
* Pol Espinasa
* Sebastian Falbesoner
* Sergi Delgado
* Stephan Vuylsteke
* TheCharlatan
* Vasil Dimov
* Will Clark
* w0xlt
