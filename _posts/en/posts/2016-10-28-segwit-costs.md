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
excerpt: A summary of some of the costs and expected risks of deploying segregated witness.
---
{% include _toc.html %}
{% include _references.md %}

## Introduction

This post is a companion to the earlier post on [Segregated Witness Benefits](/en/2016/01/26/segwit-benefits/), giving an overview of the technical costs and risks that may be incurred by activating segregated witness via [BIP141][].

### Aims

For the purpose of this post, we will use *costs* to describe negative results that are certain to occur if segwit is deployed and activated, and *risks* to describe negative impacts that may not happen, or changes that not everyone may consider negative.

When analysing risks, we consider steps undertaken to *avoid* the risk (that is, to minimise the chance of it occurring), and steps undertaken to *mitigate* the risk (that is, if it does occur, how the negative impact can be minimised).

This post does not attempt to produce a conclusion as to whether the benefits outweigh the costs or whether segwit should be deployed or activated, but rather to assist by providing background information to assist stakeholders in making informed decisions.

## Serialisation costs

Transactions and block information are serialised for three main purposes:

 1. Transmission across the peer-to-peer network;

 2. Storage of the blockchain on disk; and

 3. Evaluation of block limits.

Segwit affects serialisation in two ways:

 * A witness commitment is included in the coinbase transaction, adding between 38 and 47 bytes, or about 0.005% of a block.  (See [BIP 141 - commitment structure](https://github.com/bitcoin/bips/blob/master/bip-0141.mediawiki#commitment-structure))

 * A new transaction serialisation that includes the segregated witness data is defined (see [BIP 141](https://github.com/bitcoin/bips/blob/master/bip-0141.mediawiki#Transaction_ID), or [BIP 144](https://github.com/bitcoin/bips/blob/master/bip-0144.mediawiki#Serialization)).  This adds an overhead of 2 bytes per transaction to allow the serialisation formats to be easily distinguished, and an overhead of 1 byte per input for the count of witness items for each input. These combine to about 1% per transaction.

The segwit transaction formats (see [BIP 141 - witness program](https://github.com/bitcoin/bips/blob/master/bip-0141.mediawiki#witness-program)) have the following impact when serialised:

 * Compared to P2PKH, P2WPKH uses 3 *fewer* bytes (-1%) in the scriptPubKey, and the same number of witness bytes as P2PKH scriptSig.

 * Compared to P2SH, P2WSH uses 11 additional bytes (6%) in the scriptPubKey, and the same number of witness bytes as P2SH scriptSig.

 * Compared to P2PKH, P2WPKH/P2SH uses 21 additional bytes (11%), due to using 24 bytes in scriptPubKey, 3 fewer bytes in scriptSig than in P2PKH scriptPubKey, and the same number of witness bytes as P2PKH scriptSig.

 * Compared to P2SH, P2WSH/P2SH uses 35 additional bytes (19%), due to using 24 bytes in scriptPubKey, 11 additional bytes in scriptSig compared to P2SH scriptPubKey, and the same number of witness bytes as P2SH scriptSig.

The percentages above are based on a transaction of 180 bytes with one input and one output. These proportions remain roughly the same as the number of inputs/outputs increases, but decreases if more complicated transaction scripts (such as multisig) are in use.

### Rationale

The transaction size overhead is due to two factors:

 1. using a 256 bit hash for P2WSH rather than the 160 bit hash for P2SH; and
 2. encoding via P2SH so that old wallets that don't support segwit can send funds that can be spent using segwit, allowing the receiver to gain segwit's benefits.

Without these two factors, the overhead would be negligible at -3 bytes for P2WPKH and +1 bytes for P2WSH.

The motivation behind the first factor is discussed under [Increased security for multisig via pay-to-script-hash (P2SH)](/en/2016/01/26/segwit-benefits/#increased-security-for-multisig-via-pay-to-script-hash-p2sh).

The second factor is a tradeoff that individual users can make when publishing a receiving address, and users who choose P2WPKH/P2SH or P2WSH/P2SH will pay higher fees in proportion to the overhead.  This should naturally limit the impact of this overhead in the long term.

### Future reductions

It is possible to make most of this overhead disappear via changes to the network and storage serialisation formats: the full serialisation can be recovered from a simple flag indicating which format is in use (P2PKH, P2WPKH, P2WPKH/P2SH, etc) along with the actual data (the pubkey and signature).  (For example, [compacted_txn.txt](https://people.xiph.org/~greg/compacted_txn.txt))

## Block validation costs

With segwit, additional processing is introduced when validating a block in order both to check the witness merkle tree, and to deal with P2SH-encoded witness transactions. This requires about five additional SHA256 hashes per transaction, an additional SHA256 per P2SH-encoded-P2WSH input, and an additional HASH160 per P2SH-encoded-P2WPKH output. This however only amounts to about six SHA256 runs over at most 4MB of data, or roughly about 24MB of SHA256 data in total, which should translate to at most an additional 15
seconds per block on a Raspberry Pi v1, or under 1/10th of a second on more capable hardware.

## Risk of introducing bugs

The segwit patch set is a major change to Bitcoin, and was rolled out, though not activated on the main Bitcoin network, in Bitcoin Core 0.13.0.  Any major change like this runs a variety of risks, including:

 * Outright bugs: mistakes can be made in design or implementation giving unexpected or harmful results. For example [PR#8525](https://github.com/bitcoin/bitcoin/pull/8525).

 * User errors: changes to the system can result in user confusion, resulting in incorrect use of the system, which in turn may lead to harmful results.

 * Ecosystem interactions: different parts of the Bitcoin ecosystem may have hard-coded assumptions that will be violated with the update. For example, applications that parse bitcoind's block store will need to be updated to understand the new serialisation.

### Avoidance

In order to reduce the chances of these risks occurring when segwit is activated, the following steps have been undertaken:

 * Peer review: all the changes in segwit, both design and implementation, have been presented publicly and reviewed by multiple independent experts; often resulting in suggested improvements being undertaken.

    Public presentations include:

    - [Elements Project](https://www.elementsproject.org/elements/segregated-witness/)
    - [Scaling Bitcoin Hong Kong](http://diyhpl.us/wiki/transcripts/scalingbitcoin/hong-kong/segregated-witness-and-its-impact-on-scalability/)
    - [SF Bitcoin Devs](https://www.youtube.com/watch?v=NOYNZB5BCHM)
    - [Scaling Bitcoin Milan](http://diyhpl.us/wiki/transcripts/scalingbitcoin/milan/segwit-lessons-learned/)
    - [BIP141][],
      [BIP142][],
      [BIP143][],
      [BIP144][], and
      [BIP145][]

    Technical reviews include:

    - [PR#7910](https://github.com/bitcoin/bitcoin/pull/7910)
    - [PR#8149](https://github.com/bitcoin/bitcoin/pull/8149)
    - [Development branch pull requests](https://github.com/sipa/bitcoin/pulls?utf8=%E2%9C%93&q=is%3Apr%20)
    - [Bitcoin Core Zurich Meeting](https://bitcoincore.org/logs/2016-05-zurich-meeting-notes.html)
    - [Peter Todd's review](https://petertodd.org/2016/segwit-consensus-critical-code-review)

 * Test cases: as described in the [Next Steps](https://bitcoincore.org/en/2016/06/24/segwit-next-steps/#how-segwit-was-tested) post, "The combined changes to the consensus rules and the P2P networking code consist of 1,486 lines of added or modified code. The segwit patch also includes an additional 3,338 lines of added or modified code in the unit and integration tests that help ensure segwit is functioning as expected on every full build of the Bitcoin Core program."

 * Test networks: during development, segregated witness has been deployed on multiple test nets, allowing the code to be vetted, and developers from the wider ecosystem, such as block explorers and wallets, to ensure their software interoperates correctly with segregated witness. These test networks have included:
   - Elements Project -- tested the concept of segregated witness implemented as a hard-fork, along with many other changes
   - segnet1 through segnet4 -- tested implementation of segwit as a soft-fork, between January and May 2016
   - testnet3 -- segwit activated on the standard testnet in May 2016

 * Alternative implementations: the segwit BIPs have been reimplemented in [btcd](https://github.com/btcsuite/btcd/pull/656) (Go) and [Bcoin](https://medium.com/purse-essays/introducing-bcoin-fdfcb22dfa34) (Javascript), as well as in [various wallets and libraries](/en/segwit_adoption/).  Independent reimplementation helps shake out unstated assumptions and ambiguities in the design, and avoid bugs that may result from them.

### Mitigation

A major factor in mitigating the impact of any bugs is that segwit is implemented as a soft-fork. This means:

 * Users of Bitcoin can simply avoid newly introduced features until they are personally confident they are implemented correctly, without losing any functionality.

 * All valid segwit blocks are also valid blocks to pre-segwit software, so earlier versions of Bitcoin that don't include the segwit changes, and hence don't include any bugs introduced in those changes, can be used to verify blocks to provide a second level of assurance against the possibility of consensus regressions.

In addition, the possibility of versioning the "script" language introduces the possibility to fix bugs in the Bitcoin script language, including both pre-existing bugs as well as any potential new bugs that segwit may introduce.

## Risks related to complexity and technical debt

The concept of *technical debt* is that an easy fix now might cause enough difficulty and problems in the long term, that spending more time and effort now will turn out to be more economical.

In the context of Bitcoin, there are two types of technical debt:

 * ugly or complicated code, that can be refactored without impacting users or consensus; and
 * poor design decisions, many of which have to be retained indefinitely, as otherwise Bitcoin users would lose access to their existing coins.

### Avoidance

As noted above, the segwit code has been heavily reviewed, which helps resist the introduction of technical debt at both a code and design level.

Also as noted above, segwit has multiple independent reimplementations, which helps discover any unnecessary complexity and technical debt at the point that it can still be avoided.

In support of existing efforts to pay down technical debt by refactoring and improving the Bitcoin codebase, segwit was merged as a code-only update as part of [work towards the 0.13.0 release](https://bitcoincore.org/en/meetings/2016/05/26/).

### Mitigation

Bitcoin already suffers from some significant design debt, and segwit is specifically designed to reduce the impact of some of this debt (notably transaction malleability, linear scaling of signature hashing, and signing of input values).

The script versioning method provided by segwit provides an elegant way of allowing future soft-fork updates to further reduce design debt, including by fixing bugs in existing opcodes (such as CHECKMULTISIG), re-enabling disabled opcodes (such as CAT), or switching to superior verification methods (such as Schnorr signatures, or aggregate signatures).

Generally, design debt in Bitcoin script cannot be fully paid off, as it is always possible that there are some unspent transactions paying to P2SH addresses that make use of the "ugly" functionality. Disabling those features would render those transactions unspendable, effectively stealing funds from users. Script versioning allows the "cost" of this design debt to be reduced, by partitioning such "ugly" functionality as only applicable to "old" script versions, thus allowing new development work to largely ignore the old code.

## Risks related to soft-fork deployment

A soft-fork is any change to Bitcoin consensus rules that invalidates some set of previously valid transaction. A poorly handled soft-fork can cause a number of problems in the Bitcoin ecosystem, and, because segwit makes the additional witness data critical to establishing Bitcoin's distributed consensus, a poorly handled upgrade could cause the system to fail in additional ways. The primary potential failure modes include:

 1. making it impossible for some Bitcoin holders to spend their money
 2. causing old nodes and upgraded nodes to have a different view of which unconfirmed transactions are valid and likely to be mined
 3. having miners mistakenly mine blocks that are not valid under the new rules
 4. being activated, with some actual use, then backed out
 5. allowing large blockchain forks, due to the p2p network being effectively disconnected as a result of connections via old nodes that are unable to forward witness data

### Avoidance

Numerous soft-forks have already been activated in Bitcoin (including BIPs [16][BIP16], [34][BIP34], [65][BIP65], [66][BIP66], [68][BIP68], [112][BIP112], and [113][BIP113]), and this experience has been codified in the [BIP9][] process for activating soft-forks. The BIP9 process was used for deploying the CSV soft-fork (BIPs 68, 112, and 113), and resulted in a fast and unproblematic upgrade to the consensus rules for that change.

The segwit design and BIP9 deployment avoids the problems listed above in the following ways:

 1. The new restrictions imposed by segwit only affect transactions that no one would currently make use of because:

    - The affected transactions would be non-standard, and thus not relayed by the vast majority of nodes or mined by most miners.

    - Any transactions that were affected would currently be considered obvious "anyone can spend" payments, and could immediately be claimed by anyone monitoring the blockchain, and therefore should have been expected to be "lost".

 2. Old nodes will consider transactions spending segwit outputs as non-standard, due to apparent violation of [BIP62][] CLEANSTACK rules, and thus won't be included in old nodes' mempools. Similarly, transactions with P2WPKH or P2WSH outputs (though not P2WPKH/P2WSH encoded via P2SH) will also be considered non-standard due to being a new output type.

    This makes it impossible to achieve double spends of segwit outputs by relaying one transaction through old nodes and a different transaction through segwit nodes.

    However, these differences may still be used to attempt a double spend, for example by combining a non-segwit output and a segwit output in a single transaction (that will only be relayed via the upgraded segwit nodes), then attempting to double-spend it via a higher fee transaction only using the non-segwit output, which may be successfully relayed via the old nodes.

    These concerns only affect unconfirmed transactions in the mempool; once a transaction is confirmed and mined in a block, double spending remains impossible. Existing methods for monitoring double spends should remain equally effective, provided the monitoring tools are able to track segwit spends at all.

 3. Ensuring miners mine valid blocks is obviously a high priority to everyone involved, and significant work has gone into guaranteeing this is the case with segwit. This includes both the direct work, under [BIP145][], as well as indirect work, such as Compact Blocks ([BIP152][]).

 4. If the segwit soft-fork were reverted after being activated, this could allow anyone who had made segwit transactions to lose funds -- for example, a malicious miner could replay the transaction on a chain without segwit enabled, at which point it would be anyone-can-spend, and the miner could then steal the funds by spending it to themselves. There are two ways in which a segwit soft-fork could be reverted after being activated while allowing theft of segwit-enabled transactions:

    - Miners could abandon the segwit enabled chain and start mining from prior to segwit's activation. Based on the [BIP9][] activation rules, this would require abandoning over 2016 blocks (the LOCKED IN period, plus enough blocks to ensure the 95% threshold wasn't reached). This would require miners to abandon over 25,200 BTC in block reward, which at current prices is over $15,000,000 USD.

    - Miners could simply use software that does not recognise segwit rules (such as earlier versions of Bitcoin Core) to mine blocks on top of a chain that has activated segwit. This would be a hard-fork as far as segwit-aware software is concerned, and those blocks would consequently be ignored by Bitcoin users using segwit-aware validating nodes. If there are sufficiently many users using segwit nodes, such a hard-fork would be no more effective than introducing a new alt coin.

    As a result, neither approach seems likely.

 5. Significant work has gone into ensuring that segwit enabled peers will form a strongly connected subgraph of the Bitcoin P2P network.  This includes providing a dedicated service bit for witness enabled nodes and preferentially connecting to such nodes.

## Risks due to larger blocks

Segwit updates the 1MB block size limit to a 4M unit block weight limit, counting serialised witness data as one unit, and core block data as four units. As transactions that use segwit features begin to be used, this change will allow more data to be included per block (with 100% of transactions using segwit features this is expected to be about 2MB of data per block, however in the worst case could be up to 4MB of data per block). In so far as it allows a greater transaction volume, it can be expected to increase the UTXO database more quickly (with 100% of transactions using segwit features, the rate of increase might be expected to approximately double; however because segwit is a soft fork, the worst case UTXO growth is unchanged).

These outcomes may have positive attributes (more volume allows more user uptake, for example), but also have possibly significant negatives:

 * Larger blocks may result in slower block transmission, resulting in higher orphan rates for miners -- this in turn may result in lower security (less hashpower required to take over the network), or higher centralisation (larger miners being more able to reduce their orphan rate).

 * Larger blocks will result in higher resource requirements for full nodes, potentially causing users to shut down their nodes, which would result in higher centralisation.

 * Larger UTXO sets will result in higher resource requirements for miners, potentially causing miners to share validation resources, which would result in higher centralisation.

### Avoidance

The negative impact of larger blocks is limited in a number of ways:

 * Validation times of blocks have been significantly reduced thanks to deployment of libsecp256k1.

 * Deployment of Compact Blocks via [BIP152][] helps limit the impact of larger blocks on block transmission, and hence orphan rates, and also reduces the bandwidth usage of full nodes.

 * Pruning support allows users to run full nodes without storing the entire history of the blockchain, which allows users who have constrained storage resources to continue running full nodes, even with a larger block size.

 * The changes to the signature hashing algorithm used by segwit signatures to avoid quadratic scaling, provides a significant reduction in cost for some large transactions.

The negative impact of increased UTXO growth is limited by:

 * The deployment of segwit as a soft-fork to ensure the worst-case UTXO growth does not get any worse.

 * The reduced weighting of witness data to rebalance the lifecycle cost of a UTXO, reducing the cost of introducing an additional input that spends a segwit output, and therefore (relatively) increasing the cost of introducing an additional output creating a new UTXO, changing the create/spend cost ratio from about 1:4.5 to about 1:2. This should moderately reduce incentives to increase the UTXO set by both discouraging UTXO creation, and encouraging spending of UTXOs.

### Mitigation

 * Since the maximum amount of data per block is capped at no more than four times the current rate, mitigation work to address problems that arise from large blocks should be within the bounds of relatively straightforward engineering work. Further, since the expected amount of data per block is only approximately double the current rate, this means any necessary mitigation efforts should be further eased.

 * There is ongoing work to improve on-disk and network serialisation of transactions and blocks, further reducing the storage and bandwdith requirements of running a full node.

## Risks due to lower fees

The security of the Bitcoin blockchain is provided by hashpower, which is rewarded by both a fixed block reward and by fees from individual transactions. As a result, decreases in fee income have the potential to reduce the hashpower used to mine Bitcoin, which in turn may lower the security of the Bitcoin blockchain.

In so far as the individual transaction fees are determined by market forces and supply and demand, the changes introduced by segwit may risk lowering prices by increasing supply (presuming that demand does not also rise, either because of or at least concurrent with segwit deployment), and lower individual prices may result in lower overall mining revenue (if the price elasticity of demand is in the inelastic range).

In addition, the changes made in segwit may make "layer two" solutions, such as the Lightning Network, more compelling. If this leads to users treating layer two solutions as a substitute for on-chain transactions, this may significantly decrease demand for on-chain transactions, which would put additional downward pressure on transaction fee levels.

### Avoidance

Fees are currently approximately 0.5 BTC per block versus 12.5 BTC per block from the block subsidy, or about 4% of miner income, so the potential impact on miner income and hence network security is likely small in the short term.

In addition, fees have been rising over the past twelve months both in BTC denominated value (from under 0.2 BTC per block a year ago) and in real terms (from under $300 USD per BTC a year ago, to over $600 USD per BTC today), so moderate falls in fee levels will only be equivalent to reverting to fee incomes from up to twelve months ago, which should not be a major impact.

### Mitigation

Miners are able to individually and collectively limit supply, either by individually setting a soft-limit on the maximum weight for blocks they produce ("blockmaxweight" setting, which defaults to 3M), or by collectively using a soft-fork to effectively lower the consensus limits by orphaning blocks above a particular weight. This approach has the potential to prevent any fee decreases due to increased supply (or indeed to increase individual fees by reducing supply, though that may not increase overall revenue), but cannot prevent decreases to fee income due to substitution effects (such as the adoption of layer two networks).

While layer two networks may act as a substitute for on-chain transactions, they cannot avoid on-chain transactions entirely, and in some scenarios, even these comparatively few on-chain transactions from layer two networks can easily saturate the on-chain capacity with segwit enabled. Even if only a very small amount of the value of these networks are captured via on-chain transaction fees, this would likely be substantially above the current fee value.

## Risks related to long term scaling

As described above, full adoption of segwit by all transactions is expected to approximately double capacity. This provides a significant one-time increase in capacity, in either the short or medium term, depending on the speed of adoption. In addition, by adding features to enable layer two networks, some additional medium and long term scaling may be achieved. By fixing the quadratic sighash scaling bug, segwit also reduces the risk of negative impacts due to future capacity increases.

Segwit does not, however, provide any direct mechanism for scaling on-chain transaction volume further other than that one-off doubling.

This runs this risk that approaches to long-term scaling may be prevented or delayed: stakeholders may consider segwit to be "enough" scaling and decline to work on or support further scaling efforts.

### Avoidance

Efforts to avoid this risk have included:

 * Inclusion of "moderate block size increase proposals" in the [2015-12-07 Capacity increases roadmap](https://lists.linuxfoundation.org/pipermail/bitcoin-dev/2015-December/011865.html)
 * Inclusion of "flex caps or incentive-aligned dynamic block size controls" in the same roadmap.
 * Inclusion of features in segwit to make later scaling less risky, particularly [Linear scaling of sighash operations](/en/2016/01/26/segwit-benefits/#linear-scaling-of-sighash-operations) and [Moving towards a single combined block limit](/en/2016/01/26/segwit-benefits/#moving-towards-a-single-combined-block-limit).
 * Work on techniques to use block space more efficiently, such as using [Schnorr signatures and signature aggregation](http://diyhpl.us/wiki/transcripts/scalingbitcoin/milan/schnorr-signatures/)
 * Research on alternative models to the blockchain, maintaining decentralisation and security, but with better scalability properties, for example [Mimblewimble](http://diyhpl.us/wiki/transcripts/scalingbitcoin/milan/mimblewimble/), [Braiding](http://diyhpl.us/wiki/transcripts/scalingbitcoin/milan/breaking-the-chain/), and [Jute Braiding](http://diyhpl.us/wiki/transcripts/scalingbitcoin/milan/jute-braiding/).

Additionally, work that has made the scale increases segwit allows achievable (such as libsecp256k1 and compact blocks) have also, obviously, made further potential scale increases more achievable.

### Mitigation

Segwit does not make further scaling any more difficult on any technical level -- the risk here is entirely social. As a consequence, the most effective mitigation efforts are likely also social in nature: such as by having companies who support long-term scaling commit development resources to making that happen.

That segwit enables transaction volume to increase to approximately double current levels also provides the opportunity to demonstrate the actual impact of scaling, such as on node performance, decentralisation, and transaction demand, as well as the speed with which ecosystem upgrades can be undertaken. This data could reasonably be collected and used to support future scaling efforts, either by showing that some feared outcomes are less likely than expected, or by confirming valid concerns and allowing work to be focused on addressing those concerns.

## Alternative approaches

This section provides a brief comparison with some alternative approaches to achieving some or all of the benefits of segwit, and how those different approaches might change the costs and risks involved.

### Hard-forked segwit

Any hard-fork implementation of segwit would add significant new costs and risks, due to requiring all nodes to upgrade to understand the new rules prior to activation, and risking a chain fork into "old Bitcoin" and "new Bitcoin" with consequent confusion and loss of value. Due to the comparative lack of experience with hard-forks in the Bitcoin community, unexpected risks and costs might also occur, though that is obviously hard to analyse by its very nature.

A hard-fork implementation of segwit could realistically be made in two ways:

 1. SPV-invisible: if the witness commitment was moved from the coinbase into the block transaction merkle tree, most non-validating clients and wallets would continue to work without needing an upgrade. This would save the 38-47 bytes from the coinbase transaction, but does not offer any other advantages.

 2. SPV-visible: if calculation of transaction hashes were changed to exclude the scriptSig, this might allow for a simpler implementation, and reduce the per-transaction overhead; however it would render all existing Bitcoin software unable to work with those transactions prior to be being updated. Additionally, separate code paths to manage old style transactions would need to be kept, increasing code complexity and the possibility of bugs.  [BIP 134, Flexible Transactions][BIP134] presents an alternative approach at gaining some of the benefits of segwit via an SPV-visible hard-fork.

Either approach to a hard-fork would make it possible to simultaneously drastically alter the consensus limits on blocks.

### Simpler segwit

Many of the benefits of segwit could logically be separated into independent changes, and evaluated and deployed separately. The implementation requirements for the various features are, however, closely related:

 * Linear scaling of sighash operations: the CHECKSIG and CHECKMULTISIG opcodes need replacement.
 * Signing of input values: the CHECKSIG and CHECKMULTISIG opcodes need replacement.
 * Increased security for multisig via P2SH: the P2SH payment format needs replacement.
 * Script versioning: the P2SH payment format needs replacement.

Doing these fixes independently would increase the complexity of the Bitcoin codebase due to the need to handle different features being active at different times on the blockchain; while deploying them concurrently removes this complexity.

