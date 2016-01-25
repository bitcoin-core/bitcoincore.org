---
layout: post
name: segwit-benefits
id: en-segwit-benefits
title: Segregated Witness Benefits
permalink: /en/2016/01/25/segwit-benefits
version: 1
excerpt: >
  This page summarises some of the benefits of segregated witness that
  go beyond simply increasing the capacity of the block chain.
---
{% include _toc.html %}

The Segregated Witness soft-fork (segwit) includes a wide range of
features, many of which are highly technical. This page summarises
some of the benefits of those features that go beyond simply increasing
the capacity of the block chain.

## Malleability Fixes

Bitcoin transactions are identified by a 64-digit hexadecimal hash
called a transaction identifier (txid)
which is based on both the coins being spent and on who
will be able to spend the results of the transaction.  Unfortunately,
the way the hash is calculated allows anyone to make small
modifications to the transaction that will not change its meaning,
but will change the hash.  This is called third-party malleability.

For example, you could submit a transaction with hash
ef74...c309 to the
network, but instead find that the network confirms that transaction under
the hash 683f...8bfa
instead.

Segwit prevents third-party malleability by allowing Bitcoin users to move the
malleable parts of the transaction into something called a *transaction witness,* and
segregating that witness so that changes to the witness does not affect
calculation of the transaction's hash.

Previous attempts to reduce malleability have included BIP62 ("dealing
with malleability", which was partially implemented as standardness
checks but is now withdrawn) and BIP140 ("normalized txid").

### Who benefits?

- **Anyone monitoring unconfirmed transactions:** it's easiest to
  monitor the status of a transaction by looking it up by txid.  But in
  a system with third-party malleability, wallets must implement extra
  code to be able to deal with changed txids (hashes).

- **Anyone spending unconfirmed transactions:** if Alice pays Bob in
  transaction 1, Bob pays Charlie in transaction 2, and then Alice's
  payment gets malleated and confirmed with a different txid, Bob's
  payment is now invalid and Bob gets all of the bitcoins he paid
  Charlie back.

    If Bob is honest, he will reissue the payment to Charlie; but if he
    isn't, he can keep those bitcoins for himself.

- **Anyone using the block chain:** smart contracts today, such as
  micropayment channels, and anticipated new smart contracts, such as
  Lightning network channels, can use smaller transactions if they don't
  need to worry about third-party malleability, which saves block space
  for other users.  These smart contracts also become less complicated
  to design, understand, and monitor.

Note: segwit transactions only avoid malleability if all their inputs are
segwit spends (either directly, or via a backwards compatible segwit
P2SH address).

### Further information

 * [Bitcoin Wiki on Malleability](https://en.bitcoin.it/wiki/Transaction_Malleability)
 * [MtGox Malleability issues](http://www.coindesk.com/bitcoin-bug-guide-transaction-malleability/)
 * [Coin Telegraph article on 2015 Malleability attack](http://cointelegraph.com/news/115374/the-ongoing-bitcoin-malleability-attack)
 * [Bitcoin Magazine article on 2015 Malleability attack](https://bitcoinmagazine.com/articles/the-who-what-why-and-how-of-the-ongoing-transaction-malleability-attack-1444253640)
 * ["Overview of BIPs necessary for Lightning" transcript](http://diyhpl.us/wiki/transcripts/scalingbitcoin/hong-kong/overview-of-bips-necessary-for-lightning/)
 * [BIP 62](https://github.com/bitcoin/bips/blob/master/bip-0062.mediawiki)
 * [BIP 140](https://github.com/bitcoin/bips/blob/master/bip-0140.mediawiki)

## Linear scaling of sighash operations

A major problem with simple approaches to increasing the Bitcoin blocksize
is that for certain transactions, signature-hashing scales quadratically
rather than linearly.

![Linear versus quadratic](/assets/images/linear-quad-scale.png)

In essence, doubling the size of a transaction
increases can double both the number of signature operations, and the
amount of data that has to be hashed for each of those signatures to
be verified. This has been seen in the wild, where an individual block
required 25 seconds to validate, and maliciously designed transactions
could take over 3 minutes.

Segwit resolves this by changing the calculation of the transaction hash
for signatures so that each byte of a transaction only needs to be hashed
at most twice. This provides the same functionality more efficiently,
so that large transactions can still be generated without running into
problems due to signature hashing, even if they are generated maliciously
or much larger blocks (and therefore larger transactions) are supported.

### Who benefits?

Removing the quadratic scaling of hashed data for verifying signatures
makes increasing the block size safer. Doing that without also limiting
transaction sizes allows Bitcoin to continue to support payments that go to or
come from large groups, such as payments of mining rewards or crowdfunding
services.

The modified hash only applies to signature operations initiated from
witness data, so signature operations from the base block will continue
to require lower limits.

### Further information

 * [BIP 143](https://github.com/bitcoin/bips/blob/master/bip-0143.mediawiki)
 * [Blog post by Rusty Russell on the 25s transaction](http://rusty.ozlabs.org/?p=522)
 * [CVE 2013-2292 on Bitcoin wiki](https://en.bitcoin.it/wiki/Common_Vulnerabilities_and_Exposures#CVE-2013-2292)
 * [Proposal to limit transactions to 100kB](https://lists.linuxfoundation.org/pipermail/bitcoin-dev/2015-July/009494.html)
 * [Bitcoin Classic commit adding additional consensus limit on sighash bytes](https://github.com/bitcoinclassic/bitcoinclassic/commit/842dc24b23ad9551c67672660c4cba882c4c840a)

## Signing of input values

When a hardware wallet signs a transaction, it can easily verify the total
amount being spent, but can only safely determine the fee by having a
full copy of all the input transactions being spent, and must hash each
of those to ensure it is not being fed false data. Since individual
transactions can be up to 1MB in size, this is not necessarily a cheap
operation, even if the transaction being signed is itself quite small.

Segwit resolves this by explicitly hashing the input value. This means
that a hardware wallet can simply be given the transaction hash, index,
and value (and told what public key was used), and can safely sign the
spending transaction, no matter how large or complicated the transaction
being spent was.

### Who benefits?

Manufacturers and users of hardware wallets are the obvious beneficiaries;
however this likely also makes it much easier to safely use Bitcoin in
small embedded devices for "Internet of things" applications.

This benefit is only available when spending transactions sent to segwit
enabled addresses (or segwit-via-P2SH addresses).

### Further information

 * [BIP 143](https://github.com/bitcoin/bips/blob/master/bip-0143.mediawiki)

## Increased security for multisig via pay-to-script-hash (P2SH)

Multisig payments currently use P2SH which is secured by the 160-bit
HASH160 algorithm (RIPEMD of SHA256). However, if one of the signers
wishes to steal all the funds, they can find a collision between a valid
address as part of a multisig script and a script that simply pays them
all the funds with only 80-bits (2<sup>80</sup>) worth of work, which is
already within the realm of possibility for an extremely well-resourced
attacker. (For comparison, at a sustained 1 exahash/second, the Bitcoin
mining network does 80-bits worth of work every two weeks)

Segwit resolves this by using HASH160 only for payments direct to a
single public key (where this sort of attack is useless), while using
256-bit SHA256 hashes for payments to a script hash.

### Who benefits?

Everyone paying to multisig or smart contracts via segwit benefits from
the extra security provided for scripts.

### Further information

 * [Gavin Andresen asking if 80-bit attacks are worth worrying about](https://lists.linuxfoundation.org/pipermail/bitcoin-dev/2016-January/012198.html)
 * [Ethan Heilman describing a cycle finding algorithm](https://lists.linuxfoundation.org/pipermail/bitcoin-dev/2016-January/012202.html)
 * [Rusty Russell calculating costs of performing an attack](https://lists.linuxfoundation.org/pipermail/bitcoin-dev/2016-January/012227.html)
 * [Anthony Towns applying the cycle finding algorithm to exploit transactions](https://lists.linuxfoundation.org/pipermail/bitcoin-dev/2016-January/012218.html)
 * [Gavin Andresen summarising the thread](https://lists.linuxfoundation.org/pipermail/bitcoin-dev/2016-January/012234.html)

## Script versioning

Changes to Bitcoin's script allow for both improved security and
improved functionality. However, the design of script only allows
backwards-compatible (soft-forking) changes to be implemented by replacing
one of the ten extra OP\_NOP opcodes with a new opcode that can conditionally fail the script,
but which otherwise does nothing. This is sufficient for many changes --
such as introducing a new signature method or a feature like OP\_CLTV,
but it is both slightly hacky (for example, OP\_CLTV usually has to be accompanied by
an OP\_DROP) and cannot be used to enable even features as simple
as joining two strings.

Segwit resolves this by including a version number for scripts, so that
additional opcodes that would have required a hard-fork to be used
in non-segwit transactions can instead be supported by simply increasing the
script version.

### Who benefits?

Easier changes to script opcodes will make advanced scripting in Bitcoin
easier. This includes changes such as introducing Schnorr signatures,
using key recovery to shrink signature sizes, supporting sidechains,
or creating even smarter contracts by using Merklized Abstract Syntax
Trees (MAST) and other research-level ideas.

## Reducing UTXO growth

The Unspent Transaction Output (UTXO) database is maintained by each
validating Bitcoin node in order to determine whether new transactions are
valid or fraudulent. For efficient operation of the network, this database
needs to be very quick to query and modify, and should ideally be able
to fit in main memory (RAM), so keeping the database's
byte size as small as possible is valuable.

This becomes more difficult as Bitcoin grows, as each new user must have
at least one UTXO entry of their own and will prefer having multiple
entries so that they can use techniques that help improve their privacy
and flexibility, as well as to use techniques such as Lightning that
don't require putting all of their data on the block chain.

Segwit improves the situation here by providing a 75% byte-size discount
to people who create their transactions in a way that reduces the amount
of UTXO space they use so that we store a greater number of UTXO entries
in each megabyte of data.

### Who benefits?

Reduced UTXO growth will benefit miners, businesses, and users who run full nodes,
which in turn helps maintain the current security of the Bitcoin network
as more users enter the system.

## Compact fraud proofs

As the Bitcoin userbase expands, validating the entire blockchain
naturally becomes more expensive. To maintain the decentralised, trustless
nature of Bitcoin, it is important to allow those who cannot afford to
validate the entire blockchain to at least be able to cheaply validate
as much of it as they can afford.

Segwit improves the situation here by allowing a future soft-fork to
extend the witness structure to include commitment data, which will allow lightweight (SPV)
clients to enforce consensus rules such such as the number of bitcoins
introduced in a block, the size of a block, and the number of sigops
used in a block.

### Who benefits?

Fraud proofs allow SPV users to help enforce Bitcoin's consensus rules,
which will potentially greatly increase the security of the Bitcoin
network as a whole, as well as reduce the ways in which individual users
can be attacked.

These fraud proofs can be added to the witness data structure as part
of a future soft-fork, and they'll help SPV clients enforce the rules
even on transactions that don't make use of the segwit features.

## Signature pruning

For old transactions, signatures may be less interesting -- for example,
some SPV clients simply don't check signatures, trusting that miners
have already done that.  Even Bitcoin Core does not check signatures for
transactions prior to the most recent checkpoint. At present, however,
signature data is an integral part of the transaction and must be
present in order to calculate the transaction hash.

Segregating the signature data allows nodes that aren't interested in
signature data to prune it from the disk, or never download it in the
first place, saving resources.

### Who benefits?

As more transactions use segwit addresses, people running pruned or SPV
nodes will be able to operate with less bandwidth and disk space.

## Moving towards a single combined block limit

Currently there are two consensus-enforced limits on blocksize: the
block can be no larger than 1MB and, independently, there can be no
more than 20,000 signature checks performed across the transactions in
the block.

Finding the most profitable set of transactions to include
in a block given a single limit is an instance of the knapsack
problem, which can be easily solved almost perfectly with a simple
greedy algorithm. However adding the second constraint makes finding a
good solution very hard in some cases, and this theoretical problem has
been exploited in practice to force blocks to be mined at a size well
below capacity.

It is not possible to solve this problem without either a hardfork,
or substantially decreasing the block size.  Since segwit can't fix the
problem, it settles on not making it worse: in particular, segwit's newly introduced limit
on the new concept of witness data is implemented as a weighted sum
of the existing block size and the new witness data.

## Who benefits?

Ultimately miners will benefit if a future hardfork that changes the block capacity limit to be a single
weighted sum of parameters. For example:

    50*sigops + 4*basedata + 1*witnessdata < 10M

This lets miners easily and accurately fill blocks
while maximising fee income, and that will benefit users by allowing them
to more reliably calculate the appropriate fee needed for their transaction
to be mined.

### Further information

 * [Knapsack problem](https://en.wikipedia.org/wiki/Knapsack_problem)
 * [Sigop attack discussion on bitcointalk in Aug 2015](https://bitcointalk.org/index.php?topic=1166928.0;all)
 * [Gregory Maxwell on bitcoin-dev on witness limits](https://lists.linuxfoundation.org/pipermail/bitcoin-dev/2015-December/011870.html)
 * ["Validation Cost Metric" transcript](http://diyhpl.us/wiki/transcripts/scalingbitcoin/hong-kong/validation-cost-metric/)

