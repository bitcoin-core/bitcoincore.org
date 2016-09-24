---
type: posts
layout: post
lang: en
name: segwit-next-steps
id: en-segwit-next-steps
title: "Segregated witness: the next steps"
permalink: /en/2016/06/24/segwit-next-steps/
categories: [mining]
tags: [soft fork, soft forks, bip9, version bits, mining, segwit]
version: 1
excerpt: Segregated witness (segwit) is approaching release.  This post provides some background information, details about how segwit was tested, information about how the upgrade is expected to proceed, and a description of some future features that segwit makes easier to implement.
---
{% include _toc.html %}
{% include _references.md %}

{{page.excerpt}}

## Background

[Segwit][segwit faq] is a proposal to allow transaction-producing software to separate (segregate) transaction signatures (witnesses) from the rest of the data in a transaction, and to allow miners to place those witnesses outside of the traditional block structure.  This provides two immediate benefits:

1. **Elimination of malleability:** Segregating the witness allows both existing software and upgraded software that receives transactions to calculate the transaction identifier (txid) of segwit-using transactions without referencing the witness.  This solves all known cases of unwanted third-party transaction malleability, which is a problem that makes programming Bitcoin wallet software more difficult and which seriously complicates the design of smart contracts for Bitcoin.

2. **Capacity increase:** Moving witness data outside of the traditional block structure (but still inside a new-style block structure) means new-style blocks can hold more data than older-style blocks, allowing a modest increase to the amount of transaction data that can fit in a block.

Segwit also simplifies the ability to add new features to Bitcoin and improves the efficiency of full nodes, which provides long-term benefits that will be described in more detail later in this document.

For more information about segwit, please see [our FAQ][segwit faq] or BIPs [141][BIP141], [143][BIP143], and [144][BIP144].

[segwit faq]: /en/2016/01/26/segwit-benefits/

## How segwit was tested

Segwit changes several parts of the Bitcoin system, most notably the consensus rules that full validation nodes use to come to agreement on the state of the Bitcoin ledger.  If nodes cease to agree on the state of the ledger, then it becomes unsafe to receive new Bitcoin transactions, so Bitcoin developers should be careful about making any such changes and perform extensive testing on any such proposed changes.

Also important are changes to the peer-to-peer network code used to relay blocks and transactions.  As segwit transactions and blocks are organized differently than earlier transaction and block versions, it's important to ensure the network implementation can both relay segwit data to segwit nodes and older-style data to older nodes.

These combined changes to the consensus rules and the P2P networking code consist of 1,486 lines of added or modified code.  The segwit patch also includes an additional 3,338 lines of added or modified code in the unit and integration tests that help ensure segwit is functioning as expected on every full build of the Bitcoin Core program.

In addition to over 3,000 lines of added automated testing code, segwit has been extensively tested by Bitcoin developers.  This section describes just some of the rigorous testing they performed on different versions of segwit over the last year.

- Segwit was originally implemented by Pieter Wuille and several other Blockstream developers on the [Elements Project][] sidechain in April through June 2015 as a “from scratch” version that wasn’t intended to be compatible with previous Bitcoin software.  This version has been used for every single transaction on Elements-based sidechains.

- In late October 2015, Luke Dashjr [described][luke-jr sfsw] a method that would allow segwit to be implemented in Bitcoin as a soft fork and Wuille used his experience with the Elements version to begin working on this new implementation that is backwards compatible with all existing Bitcoin software (although programs do need to upgrade to fully understand segwit).

- The code became fully operational in late December 2015 on a special segwit-specific testnet (called segnet) that allowed implementers and testers to run the code in a multi-user environment, and which also allowed wallet authors to test their code for generating segwit transactions.  Segnet went through several iterations as problems were found and fixed, and as improvements were discovered and implemented.

- In April 2016, after four months of active development and testing, Wuille submitted a [pull request][PR#7910] to the Bitcoin Core project for review.

- May 2016, Segwit was enabled on Bitcoin's regular testnet where it was tested not just against other software that was expected to interact with segwit but also all the other programs which are regularly tested on Bitcoin's testnet and which had not been upgraded for segwit.

- Also in May 2016, twenty Bitcoin Core developers [met in Switzerland][may2016 core meetup] for (among other things) an in-person review of the segwit code and ensuring that test coverage was adequate.

- In June 2016, after almost two months of very active review on the original pull request plus extended operation on both segnet and testnet, Wuille created a [second pull request][PR#8149] that contained all the improvements made to the original pull request, rebased on top of the most recent version of Bitcoin Core's development branch, and which was specially formatted to make final review easy as well as ensure all reviews made to the original pull request remained valid.

With the original sidechains implementation of segwit having been used by a number of reviewers over the past year and the Bitcoin soft fork implementation having received rigorous testing and review over six months, the Bitcoin Core developers believe it is now ready to move to production.

[PR#8149]: https://github.com/bitcoin/bitcoin/pull/8149
[PR#7910]: https://github.com/bitcoin/bitcoin/pull/7910
[may2016 core meetup]: /logs/2016-05-zurich-meeting-notes.html
[luke-jr sfsw]: https://botbot.me/freenode/bitcoin-core-dev/2015-10-21/
[elements project]: http://elementsproject.org/

## Compact Blocks

Segwit will allow Bitcoin miners to include more transaction data in the blocks they create than they can now.  This will increase the bandwidth demands on Bitcoin full nodes that relay all that data as well as increase the latency between when a new block is published and when nodes receive it (as larger amounts of data typically take longer to propagate).  To help reduce these negative side effects, Bitcoin Core developers plan to make [compact block relay](/en/2016/06/07/compact-blocks-faq/) available for Bitcoin Core 0.13 and above.

Bitcoin full nodes currently download many transactions twice: once when they receive the transaction by itself and a second time when the transaction is received as part of a block.  Compact block relay can eliminate most (and sometimes all) of this duplication by sending nodes just the information they need in order to reconstruct blocks using the transactions that the nodes have already received.

In the optimistic case this reduces the amount of bandwidth a node uses by almost half.  Since segwit is expected to increase maximum capacity to about double, these two improvements roughly balance each other out to keep node bandwidth usage at roughly the same level as now.

More importantly, compact block relay helps reduce peak bandwidth load. Currently a freshly-received block of approximately 1 megabyte has to be downloaded all at once; when segwit is deployed, that will mean 2 megabyte or larger blocks may need to be downloaded.  On all but the fastest connections, these bandwidth spikes hurt the performance of other services users are running at the same time, such as games or video streaming.  With compact blocks, the user can receive transactions in a steady stream and then reconstruct each block using a small description of the block, eliminating bandwidth spikes that inconvenience many users.

Lastly, by reducing the amount of data that needs to be sent when a new block is announced, compact block relay also achieves much better block propagation speeds.  Compact block relay takes special advantage of this by supporting two operating modes, a low-bandwidth mode that’s optimized to reduce bandwidth (although it also improves speed in most cases) and a high-bandwidth mode that significantly improves speed and still manages to greatly reduce peak bandwidth usage by only requiring an average of 20 kilobytes more than the low-bandwidth mode.

The high bandwidth mode is being used as the basis for further development into optimizations for peer-to-peer block relay.  It's especially intended for use by miners who need to receive blocks as fast as possible (especially if other non-peer-to-peer block relay methods fail), but users with extra bandwidth can also enable this mode to help relay blocks to miners faster and also help discourage denial-of-service attacks by making it less evident which high-bandwidth nodes are being run by miners and which are being run by ordinary peers.

## Deployment plan

The following plan describes how segwit is expected to be deployed.

**Merge to master (without mainnet activation code):** after Bitcoin Core developers "ACK" (approve) the final segwit pull request, it will be merged into the Bitcoin Core master Git repository branch.  The code that is being merged will include everything in segwit except for the activation code.  This will make it easy for developers to test other features on top of segwit, such as compact blocks.  **Activation on testnet** has already occurred so users and developers may experiment and test segwit on testnet.

**Backport to 0.12 branch:** the unactivated code will be backported to the 0.12 maintenance branch and the backport will receive its own testing.

**Choosing the BIP9 parameters:** [BIP9][] is a soft fork deployment mechanism that allows miners to signal their readiness to enforce new consensus rules.  Each soft fork made with BIP9 chooses when miners can begin signaling for the soft fork, when the soft fork is considered unsuccessful if not enough miners have signaled for it, and which bit in the block header version field will be used by miners to signal their readiness.  These parameters will be selected at this time and implemented along with the code to activate segwit on both the master and 0.12 branches.

**Release candidate phase:** after all developer testing is successfully concluded, a release candidate (probably named 0.12.2RC1) will be publicly provided to anyone willing to test the code.  Miners, merchants, and wallet vendors are especially encouraged to test.  If any problems are found, they will be fixed and a new release candidate will be issued.  This will be repeated as necessary until a release candidate is found with no known problems.

**Binary release:** the final release candidate will have its version changed to the final release version (expected to be 0.12.2) and will be released for all users to download and begin running at their leisure (segwit is a soft fork, so upgrading is only required if they plan to use segwit features).

**Miners upgrade:** miners who choose to upgrade to 0.12.2 will be able to start producing blocks that signal readiness to enforce segwit once the date defined as segwit's BIP9 started date is reached.

**Lock-in:** once 95% of blocks in a 2,016-block long retarget period have signaled that their miners are ready to enforce segwit, segwit will lock-in -- meaning that unless the blockchain is rolled back at that point, segwit will become active (see next point).

**Activation:** 2,016 blocks (about two weeks) after segwit is locked-in, it will activate.  That means all full nodes running segwit-aware code will begin requiring miners to enforce the new segwit consensus rules.

**Wallets upgrade:** similar to the P2SH soft fork in 2012, after segwit activates it will not immediately be safe for wallets to upgrade to support segwit.  That’s because spends from segwit transactions look like unsecured transactions to older nodes, so if the blockchain is forked soon after segwit activates, those spends could be placed in an earlier block that is not subject to segwit's rules.  For this reason, it is suggested that wallets avoid upgrading for a few weeks after segwit activates.  Allowing that extra time to pass provides extra security to wallet users, although anyone who wants to test with a small amount of money they can afford to lose can begin spending as soon as segwit activates. Users can also begin testing immediately using testnet or regtest with the proposed segwit code or (when available) any release containing segwit.

## How segwit will affect you

- **Miners** who choose to run segwit will have the responsibility of ensuring that they're ready to enforce it by upgrading their full validation nodes to segwit-enforcing code.  When they’ve done this and performed whatever testing they believe is prudent, they can begin signaling support for segwit.

- **Full node operators** can continue using their existing nodes but they are recommended to upgrade to a segwit-enforcing version.  If any miners end up producing blocks that are invalid according to the segwit rules after segwit activates, upgraded full nodes will automatically reject those blocks, ensuring that those upgraded full node users see accurate confirmation counts.

    This upgrade is especially important for anyone, such as a business, who accepts transactions with low numbers of confirmations.

- **Bitcoin Core wallet users** can continue using their existing nodes.  Even if you upgrade, you will see no changes beyond those described above.  The code expected to be released in 0.12.2 does not begin generating segwit receiving addresses by default.

- **Users of other wallets** can continue using their existing wallets.  It is recommended that lightweight wallet users always wait for several confirmations when receiving significant amounts of money, so no extra waiting is expected to be required here.

    When you have the opportunity to upgrade to a version of your wallet that supports segwit, you may find that the transaction fee you have to pay is slightly lower when you spend bitcoins you received after upgrading to segwit because the witness is external and therefore the transaction size is counted as smaller.

## Future upgrades made easier by segwit

Segwit is a major step towards improving the operation of the Bitcoin system.  In addition to the fixes for third-party malleability and the capacity increase described above, it also provides a mechanism that allows Bitcoin's Script language to be more easily upgraded.

Since Bitcoin 0.3.6, the scripting language has supported ten special NOP opcodes whose behavior could be redefined in certain ways in later versions. Two of these ten special NOP opcodes have already been used to add new features to the system: NOP2 was changed to `OP_CHECKLOCKTIMEVERIFY` (CLTV) per [BIP65][] and NOP3 was changed to `OP_CHECKSEQUENCEVERIFY` (CSV) per [BIP112][].

Such operations need to be implemented very carefully to preserve the NOP semantics used by old nodes, which limits what soft forks can do and can make adding features this way somewhat messy. For example, because both CLTV and CSV accept parameters, each time they are used an `OP_DROP` opcode must also be used in order to preserve compatibility with their original NOP behavior. This makes it harder to both write and read scripts using them.

Segwit eliminates all of these problems by allowing segwit users to specify what version of the Bitcoin Script language to use.  Each version can be either a minor improvement on an earlier version or an entirely new language -- and the multiple versions can coexist together, allowing individual users to specify which version they want used to protect their transactions.

For example, segwit ships with an improvement to the `OP_CHECKSIG`, `OP_CHECKMULTISIG`, and related signature-checking opcodes that removes a known denial-of-service vulnerability vector that can be exploited through those opcodes.  This isn’t a complete solution to that problem because the previous versions of the CHECKSIG opcodes are still available for non-segwit transactions, but it does help make segwit transactions harder to abuse than non-segwit transactions.

Some ideas for future upgrades using this mechanism are described below:

### Schnorr signatures

Bitcoin uses the Elliptic Curve Digital Signature Algorithm (ECDSA).  There’s a simpler digital signature algorithm that also uses elliptic curves but which was patented up until shortly before Bitcoin was originally released.  This algorithm is called [Schnorr][] and it supports all the features in ECDSA that Bitcoin uses, including the ability to create secure signatures as well as the ability to create HD wallets.

Schnorr signatures are already used outside of Bitcoin. For example, the well-known [Ed25519 signature scheme][] is based on Schnorr.

Verification of Schnorr signatures is slightly faster than ECDSA signatures (which makes running full nodes more convenient) and the signatures can be made smaller because the DER encoding currently used for ECDSA signatures doesn’t need to be used for Schnorr (providing a modest increase in capacity).

Schnorr also allows for "native multisig" in cases where all participants need to sign (such as 2-of-2, 3-of-3, or any n-of-n) that allows all *n* public keys to be combined into a single overall public key and all *n* signatures to be combined into a single overall signature.  The overall public keys and signatures are the same size as a single one of the original public keys and signatures, so it's possible to create a 100-of-100 multisig transaction that’s no larger than a 1-of-1 transaction.  This will be quite useful as it expected that the network will see increased use of n-of-n multisig transactions (for example 2-of-2 is used in many payment channel transactions).

Schnorr support is already available in the [libsecp256k1][] library that Bitcoin Core uses for creating and verifying signatures, although Bitcoin doesn’t currently use Schnorr in any way and there are some changes the developers would like to make to the Schnorr parts of the library before starting to use it.  This, combined with segwit's support for Bitcoin Script versioning should make adding the features described above fairly easy.

A more difficult feature to add that is supported by Schnorr is [signature aggregation][].  This would change how signature validation works and would allow multisig transactions requiring 1-of-2, 2-of-3, or any m-of-n signatures to create just one signature per transaction provided all the signers are online simultaneously.  This would also allow creating one signature for each transaction no matter how many inputs it has (again, if all signers are online simultaneously).

Since signatures are often the largest individual part of transactions and many transactions have two or more inputs each requiring at least one signature, this would significantly reduce the size of many transactions and would also speed up validation (as only one signature total would need to be validated rather than one signature per input).

Once implemented, signature aggregation could be combined with the coinjoin privacy-enhancing technique to create a significant financial incentive to use coinjoin for spending your bitcoins.  Currently, there is a slight financial incentive to use coinjoin as a coinjoin transaction that combines what would be multiple individual transactions from different people is slightly smaller than what would be the total size of all those individual transactions.  With signature aggregation, the combined transaction would be significantly smaller because it would only need one signature rather than many signatures.

Although signature aggregation is still being designed, it will be easy to add support for it to the protocol using segwit's support for different versions of the Bitcoin Script evaluation language.

[schnorr]: https://en.wikipedia.org/wiki/Schnorr_signature
[ed25519 signature scheme]: http://ed25519.cr.yp.to/
[libsecp256k1]: https://github.com/bitcoin-core/secp256k1
[signature aggregation]: https://bitcointalk.org/index.php?topic=1377298.0

### MAST

[Merkelized Abstract Syntax Trees][BIP114] (MAST) allow the [creation of Bitcoin scripts][todd mast] with many different conditionals but which may allow only a relatively small amount of data to be placed in transactions.

They work by taking a program and splitting each conditional part of it into a separate segment, and then placing each of those segments into a merkle tree.  Bitcoins are spent (encumbered) to the merkle root of the merkle tree.

A minimal set of conditions that need to be used for final validation can be revealed to all full validation nodes, but code that does not execute can be replaced by a simple hash as part of a partial merkle branch, allowing all parts of the script to be connected to the merkle root used in the encumbrance so that validation can be completed.

This not only saves space but it may also help improve privacy.  For example, if Alice wants to be able to spend her bitcoins normally every day but also wants her estate lawyer Bob to be able to spend her bitcoins if they’ve been inactive for a year, she can place both of those conditions in separate branches.  When Alice is making normal spends, nobody can see just by looking at the transaction what that second condition is.

Enabling MAST can be done using the versioning of Bitcoin Script enabled by segwit.

[todd mast]: https://bitcointalk.org/index.php?topic=255145.msg2757327#msg2757327
