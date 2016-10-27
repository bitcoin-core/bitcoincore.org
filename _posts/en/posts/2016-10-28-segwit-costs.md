---
title: Segregated Witness Costs and Risks
name: segwit-costs
id: en-segwit-costs
lang: en
permalink: /en/2016/10/28/segwit-costs/
type: posts
layout: post
share: yes
version: 1
excerpt: This page summarises some of the costs and expected risks of deploying segregated witness.
---
{% include _toc.html %}

# Introduction

This post is a companion to the earlier post on [Segregated
Witness Benefits](/en/2016/01/26/segwit-benefits/),
giving an overview of the technical costs and risks that
may be incurred by activating segregated witness via [BIP
141](https://github.com/bitcoin/bips/blob/master/bip-0141.mediawiki).

## Aims

For the purpose of this post, we will use *costs* to describe negative
results that are certain to occur if segwit is deployed and activated,
and *risks* to describe negative impacts that may not eventuate, or
changes that are only subjectively negative -- that is a reasonable
person may not consider them to be negative.

When analysing risks, we consider steps undertaken to *avoid* the risk
(that is, to minimise the chance of it occuring), and steps undertaken to
*mitigate* the risk (that is, if it does occur, how the negative impact
can be minimised).

This post does not attempt to produce a conclusion as to whether the
benefits outweigh the costs or whether segwit should be deployed or
activated, but rather to assist by providing background information to
assist stakeholders in making informed decisions.

# Serialisation costs

Transactions and block information are serialised for three main purposes:

 1. Transmision across the peer-to-peer network;

 2. Storage of the blockchain on disk; and

 3. Evaluation of block limits.

Segwit affects serialisation in two ways:

 * A witness commitment is included in the coinbase transaction,
   adding between 38 and 47 bytes, or about 0.005% of a block.
   (See [BIP 141 - commitment structure](https://github.com/bitcoin/bips/blob/master/bip-0141.mediawiki#commitment-structure))

 * A new transaction serialisation that includes the
   segregated witness data is defined (see [BIP 141](https://github.com/bitcoin/bips/blob/master/bip-0141.mediawiki#Transaction_ID), or [BIP 144](https://github.com/bitcoin/bips/blob/master/bip-0144.mediawiki#Serialization)).
   This adds an overhead of 2 bytes per transaction to allow the
   serialisation formats to be easily distinguished, and an overhead of
   1 byte per input for the count of witness items for each input. These
   combine to about 1% per transaction.

The segwit transaction formats (see [BIP 141 - witness
program](https://github.com/bitcoin/bips/blob/master/bip-0141.mediawiki#witness-program)) have the following impact when serialised:

 * Compared to P2PKH, P2WPKH uses 3 *fewer* bytes (-1%) in the
   scriptPubKey, and the same number of witness bytes as P2PKH scriptSig.

 * Compared to P2SH, P2WSH uses 11 additional bytes (6%) in the
   scriptPubKey, and the same number of witness bytes as P2SH scriptSig.

 * Compared to P2PKH, P2WPKH/P2SH uses 21 additional bytes (11%),
   due to using 24 bytes in scriptPubKey, 3 fewer bytes in scriptSig
   than in P2PKH scriptPubKey, and the same number of witness bytes as
   P2PKH scriptSig.

 * Compared to P2SH, P2WSH/P2SH uses 35 additional bytes (19%),
   due to using 24 bytes in scriptPubKey, 11 additional bytes in scriptSig
   compared to P2SH scriptPubKey, and the same number of witness bytes
   as P2SH scriptSig.

The percentages above are based on a transaction of 180 bytes with one
input and one output. These proportions remain roughly the same as the
number of inputs/outputs increases, but decreases if more complicated
transaction scripts (such as multisig) are in use.

## Rationale

The transaction size overhead is due to two factors:

 1. using a 256 bit hash for P2WSH rather than the 160 bit hash
    for P2SH; and
 2. encoding via P2SH so that old wallets without segwit support
    can send funds that can be spent using segwit, allowing the receiver
    to gain segwit's benefits.

Without these two factors, the overhead would be negligible at -3 bytes
for P2WPKH and +1 bytes for P2WSH.

The motivation behind the first factor is discussed under
[Increased security for multisig via pay-to-script-hash
(P2SH)](/en/2016/01/26/segwit-benefits/#increased-security-for-multisig-via-pay-to-script-hash-p2sh).

The second factor is a tradeoff that individual users can make when
publishing a receiving address, and users who choose P2WPKH/P2SH
or P2WSH/P2SH will pay higher fees in proportion to the overhead.
This should naturally limit the impact of this overhead in the long term.

## Future reductions

It is possible to make most of this overhead disappear via
changes to the network and storage serialisation formats: the
full serialisation can be recovered from a simple flag indicating
which format is in use (P2PKH, P2WPKH, P2WPKH/P2SH, etc) along
with the actual data (the pubkey and signature).  (For example,
[compacted_txn.txt](https://people.xiph.org/~greg/compacted_txn.txt))

# Block validation costs

With segwit, additional processing is introduced when validating
a block in order both to check the witness merkle tree, and to
deal with P2SH-encoded witness transactions. This requires about
five additional SHA256 hashes per transaction, and an additional
SHA256 per P2SH-encoded-P2WSH input, and an additional HASH160 per
P2SH-encoded-P2WPKH output. This however only amounts to about six SHA256
runs over at most 4MB of data, or roughly about 24MB of SHA256 data in
total, which should translate to at most an additional 15s per block on
a Raspberry Pi, or under a tenth of a second on more capable hardware.

# Risk of introducing bugs

The segwit patch set is a major change to Bitcoin, and was rolled out,
though not activated on the main bitcoin network, in Bitcoin Core 0.13.0.
Any major change like this runs a variety of risks, including:

 * Outright bugs: mistakes can be made in design or implementation
   giving unexpected or harmful results. For example
   [PR#8525](https://github.com/bitcoin/bitcoin/pull/8525).

 * User errors: changes to the system can result in user confusion,
   resulting in incorrect use of the system, which in turn may lead to
   harmful results.

 * Ecosystem interactions: different parts of the Bitcoin ecosystem may
   have hard-coded assumptions that will be violated with the update. For
   example, applications that parse bitcoind's block store will need to
   be updated to understand the new serialisation.

## Avoidance

In order to reduce the chances of these risks eventuating when segwit
is activated, the following steps have been undertaken:

 * Peer review: all the changes in segwit, both design and implementation
   have been presented publicly and reviewed by multiple independent
   experts; often resulting in suggested improvements being undertaken.

   Public presentations include:
   - [Elements Project](https://www.elementsproject.org/elements/segregated-witness/)
   - [Scaling Bitcoin Hong Kong](http://diyhpl.us/wiki/transcripts/scalingbitcoin/hong-kong/segregated-witness-and-its-impact-on-scalability/)
   - [SF Bitcoin Devs](https://www.youtube.com/watch?v=NOYNZB5BCHM)
   - [Scaling Bitcoin Milan](http://diyhpl.us/wiki/transcripts/scalingbitcoin/milan/segwit-lessons-learned/)
   - [BIP 141](https://github.com/bitcoin/bips/blob/master/bip-0141.mediawiki),
     [BIP 142](https://github.com/bitcoin/bips/blob/master/bip-0142.mediawiki),
     [BIP 143](https://github.com/bitcoin/bips/blob/master/bip-0143.mediawiki),
     [BIP 144](https://github.com/bitcoin/bips/blob/master/bip-0144.mediawiki),
     [BIP 145](https://github.com/bitcoin/bips/blob/master/bip-0145.mediawiki)

   Technical reviews include:
   - [PR#7910](https://github.com/bitcoin/bitcoin/pull/7910)
   - [PR#8149](https://github.com/bitcoin/bitcoin/pull/8149)
   - [Development branch pull requests](https://github.com/sipa/bitcoin/pulls?utf8=%E2%9C%93&q=is%3Apr%20)
   - [Bitcoin Core Zurich Meeting](https://bitcoincore.org/logs/2016-05-zurich-meeting-notes.html)
   - [Peter Todd's review](https://petertodd.org/2016/segwit-consensus-critical-code-review)

 * Test cases: as described in the [Next Steps](https://bitcoincore.org/en/2016/06/24/segwit-next-steps/#how-segwit-was-tested)
   post, "The combined changes to the consensus rules and the P2P
   networking code consist of 1,486 lines of added or modified code. The
   segwit patch also includes an additional 3,338 lines of added or
   modified code in the unit and integration tests that help ensure
   segwit is functioning as expected on every full build of the Bitcoin
   Core program."

 * Test networks: during development, segregated witness has been deployed
   on multiple test nets, allowing the code to be vetted, and developers
   from the wider ecosystem, such as block explorers and wallets,
   to ensure their software interoperates correctly with segregated
   witness. These test networks have included:
   - Elements Project -- tested the concept of segregated witness
     implemented as a hard-fork, along with many other changes
   - segnet1 through segnet4 -- tested implementation of segwit as a
     soft-fork, between January and May 2016
   - testnet3 -- segwit activated on the standard testnet in May 2016

 * Alternative implementations: the segwit BIPs have been reimplemented in
   [btcd](https://github.com/btcsuite/btcd/pull/656) (Go) and
   [Bcoin](https://medium.com/purse-essays/introducing-bcoin-fdfcb22dfa34) (Javascript),
   as well as in [various wallets and libraries](/en/segwit_adoption/).
   Independent reimplementation helps shake out unstated assumptions, and
   ambiguities in the design, and avoid bugs that may result from them.

## Mitigation

A major factor in mitigating the impact of any bugs is that segwit is
implemented as a soft-fork. This means:

 * Users of bitcoin can simply avoid newly introduced features until
   they are personally confident they are implemented correctly,
   without losing any functionality.

 * All valid segwit blocks are also valid blocks to pre-segwit software,
   so earlier versions of bitcoin that don't include the segwit changes,
   and hence don't include any bugs introduced in those changes, can be
   used to verify blocks to provide a second level of assurance against
   the possibility of consensus regressions.

In addition, the possibility of versioning the "script" language
introduces the possibility to fix bugs in the bitcoin script language,
including both pre-existing bugs as well as any potential new bugs that
segwit may introduce.

# Risks related to complexity and technical debt

The concept of *technical debt* is that an easy fix now might cause
enough difficulty and problems in the long term, that spending more time
and effort now will turn out to be more economical.

In the context of Bitcoin, there are two types of technical debt:

 * ugly or complicated code, that can be refactored without impacting users or
   consensus; and
 * poor design decisions, many of which have to be retained indefinitely,
   as otherwise Bitcoin users would lose access to their existing coins.

## Avoidance

As noted above, the segwit code has been heavily reviewed, which
helps resist the introduction of technical debt at both a code and
design level.

As also noted above, segwit has multiple independent reimplementations,
which helps discover any unnecessary complexity and technical debt at
the point that it can still be avoided.

In support of existing efforts to pay down technical debt
by refactoring and improving the bitcoin codebase, segwit was merged as
a code-only update as part of [work towards the 0.13.0 release](https://bitcoincore.org/en/meetings/2016/05/26/).

## Mitigation

Bitcoin already suffers from some significant design debt, and segwit
is specifically designed to reduce the impact of some of this debt
(notably transaction malleability, linear scaling of signature hashing,
and signing of input values).

The script versioning method provided by segwit provides an elegant way of
allowing future soft-fork updates to further reduce design debt, including
by fixing bugs in existing opcodes (such as CHECKMULTISIG), re-enabling
disabled opcodes (such as CAT), or switching to superior verification
methods (such as Schnorr signatures, or aggregate signatures).

Generally, design debt in Bitcoin script cannot be fully paid off, as
it is always possible that there are some unspent transactions paying
to P2SH addresses that make use of the "ugly" functionality. Disabling
those features would render those transactions unspendable, effectively
stealing funds from users. Script versioning allows the "cost" of this
design debt to be reduced, by partitioning such "ugly" functionality as
only applicable to "old" script versions, thus allowing new development
work to largely ignore the old code.
