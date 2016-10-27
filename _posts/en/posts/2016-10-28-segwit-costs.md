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
