---
# This file is licensed under the MIT License (MIT) available on
# http://opensource.org/licenses/MIT.

layout: post
lang: en
columns: 1
id: details-chain-hijacking
title: Details about chain hijacking
breadcrumbs:
  - features
  - verification
---
> Alice believes that there should never be more than 21 million
> bitcoins---but one day she's tricked into accepting "bitcoins" that
> are only valid on a block chain with permanent 10% inflation.

- **Wallet service**{:.fgred} users have to use whatever block chain the
  service uses.  Services can even profit from switching their users to
  a new chain and selling their users' bitcoins from the old chain.

- **Light wallet**{:.fgred} users accept the block chain they know about
  with the most proof of work. This lets the hash rate majority of
  miners force light wallet users off of Bitcoin.

- **Bitcoin Core**{:.fggreen} users don't have to worry about chain
  hijacking because Bitcoin Core verifies every block using *all* of
  Bitcoin's consensus rules.

<div class="callout" markdown="block">
Preventing chain hijacking is one of Bitcoin Core's most important jobs.
The alternative is to allow miners to do whatever they want.
</div>

### Real Example

In July 2015, several large Bitcoin miners accidentally produced an
invalid block chain several blocks longer than the correct block chain.
Some wallet services and many light wallets accepted this longer chain,
putting their users' bitcoins at risk.

Recent versions of Bitcoin Core never accepted any of the blocks from
the invalid chain and never put any bitcoins at risk.

It is believed that the miners at fault controlled more than 50% of the
network hash rate, so they could have continued to fool light wallets
indefinitely.  It was only their desire to remain compatible with
Bitcoin Core users that forced them to abandon over $37,500 USD worth of
mining income.

**Learn more:** [July 2015 chain forks](https://en.bitcoin.it/wiki/July_2015_chain_forks)
