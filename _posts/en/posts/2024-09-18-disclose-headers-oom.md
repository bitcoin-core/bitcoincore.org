---
title: CVE-2019-25220 - Memory DoS due to headers spam
name: blog-disclose-headers-oom
id: en-blog-disclose-headers-oom
lang: en
type: advisory
layout: post

## If this is a new post, reset this counter to 1.
version: 1

## Only true if release announcement or security annoucement. English posts only
announcement: 1

excerpt: >
  An attacker could spam a Bitcoin Core node with low-difficulty headers chains, which could be used to remotely crash it.
---

Before Bitcoin Core v24.0.1, attackers could spam nodes with low-difficulty headers chains, which
could be used to remotely crash peers.

This issue is considered **High** severity.

## Details

Bitcoin Core stores the blockchain headers in memory. This makes it susceptible to being DoSed, by
having it download and store extremely long chains of headers, even if they are of low difficulty.
It is important to note that once crafted, an attack chain could be reused to crash any node on the
network.

The possibility of using this to attack nodes has long been known, and was the primary reason why
the checkpoint system was still in place: making an attacker start an attack at the last checkpoint
makes it far more costly than starting at the genesis block. However, over time, with decreasing
hashrate costs, even this mitigation became less effective.

This attack was independently discovered and reported to the Bitcoin Core project in January 2019 by
David Jaenson, who suggested introducing newer checkpoints as a practical mitigation. However:
1. This still leaves nodes performing IBD with no protection before they receive checkpoint blocks.
2. It relies on the ecosystem semi-regularly adopting updated software with new checkpoints, a
   practice which Bitcoin Core contributors have long been uncomfortable with.

It later got increased attention when Braydon Fuller [posted his "Chain width expansion"
writeup](https://lists.linuxfoundation.org/pipermail/bitcoin-dev/2019-October/017354.html) to the
bitcoin-dev mailing list in October 2019. He had previously responsibly reported it to the Bitcoin
Core security list. The suggested approach was not adopted in Bitcoin Core due to concerns about
network convergence when limiting the number of parallel chains.

At the time, the computational cost of creating a huge low-difficulty headers chain was equal to
about 32.28% of mining one block at the tip. That is a cost of about 4.12 BTC since the block reward
then was about 12.77 BTC.

By February 2022, the cost of the attack had dropped further to around 14.73% of the cost of mining
a block, and this prompted investigation of alternative solutions. If unaddressed, the cost today
(September 2024) would just be 4.44% of a block. These figures translate to a cost of about 1.07 BTC
and 0.14 BTC respectively, given the block reward at these dates.

A protection against this DoS was implemented in Bitcoin Core PR
[#25717](https://github.com/bitcoin/bitcoin/pull/25717), whereby the node will first verify a
presented chain has enough work before committing to store it. With that, Bitcoin Core no longer
relies on having checkpoints to protect against any known attacks.

## Attribution

Credit goes to David Jaenson and Braydon Fuller for independently re-discovering the attack,
estimating its cost and suggesting modifications.

Credit goes to Suhas Daftuar and Pieter Wuille for researching a satisfying fix and implementing it.

## Timeline

* 2010-07-17 - Bitcoin 0.3.2 is released, which introduces checkpoints. They protect among other
  things against low-difficulty block spam.
* 2011-11-21 - Bitcoin 0.5.0 is released, which skips script validation for blocks before the last
  checkpoint. This makes the role of checkpoints even more security-critical.
* 2014-04-09 - Block 295000 is mined, which becomes the last Bitcoin Core checkpoint. The protection
  offered by checkpoints against block spam starts eroding from this point on as hashrate costs
  decrease.
* 2015-02-16 - Bitcoin Core 0.10.0 is released, with headers-first synchronization. This weakens the
  low-difficulty block spam attack to a block *header* spam attack.
* 2017-03-08 - Bitcoin Core 0.14.0 is released, which disentangles the skipping of script validation
  from checkpoints, leaving them only relevant for protecting against block header spam.
* 2019-01-28 - David Jaenson reports this issue to the Bitcoin Core security mailing list.
* 2019-09-18 - Braydon Fuller emails the Bitcoin Core security list with a paper titled "[Bitcoin
  Chain Width Expansion Denial-of-Service
  Attacks](https://bcoin.io/papers/bitcoin-chain-expansion.pdf)", which discusses the dangers of
  block and block header spam, a cost analysis, and a proposed solution.
* 2019-09-26 - Suhas Daftuar replies to Braydon Fuller that it's a known issue, and invites him to
  post his writeup to the bitcoin-dev mailing list.
* 2019-10-04 - Braydon Fuller sends his paper to the bitcoin-dev mailing list.
* 2019-10-31 - In response to the above events, Suhas Daftuar opens PR
  [#17332](https://github.com/bitcoin/bitcoin/pull/17332) with an earlier but impractical proof of
  concept he worked on to improve the situation, with the hope of causing more discussion on the
  topic.
* 2022-02    - Suhas Daftuar and Pieter Wuille discuss this issue and estimate that the cost of this
  attack has now actually become so low that it warrants immediate action, and the need to avoid
  talking about it publicly.
* 2022-06-22 - Suhas Daftuar opens PR [#25454](https://github.com/bitcoin/bitcoin/pull/25454) as
  preparatory work toward implementing a fix.
* 2022-06-22 - Suhas Daftuar messages a group of long-term contributors detailing the attack, its
  cost and the fix Pieter Wuille and him have been working on.
* 2022-07-26 - Suhas Daftuar opens PR [#25717](https://github.com/bitcoin/bitcoin/pull/25717),
  co-authored with Pieter Wuille, which implements the fix.
* 2022-08-30 - PR #25717 is merged.
* 2022-10-21 - Niklas Gögge's PR [#26355](https://github.com/bitcoin/bitcoin/pull/26355) is merged,
  which fixes a bug in the headers pre-synchronization step that was introduced in PR #25717.
  Without this, it would still have been possible to spam block headers.  The discovery of this bug,
  and the possibility of potential undiscovered ones, is the reason why the old checkpoints have not
  been removed entirely yet.
* 2022-12-12 - Bitcoin Core 24.0.1 is released with the fix.
* 2023-12-07 - The last vulnerable version of Bitcoin Core (23.2) goes end of life.
* 2024-09-18 - Public disclosure.

{% include references.md %}
