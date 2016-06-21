---
type: posts
layout: post
lang: en
name: csv-softfork-instructions
id: en-csv-softfork-instructions
title: CSV softfork - Important upgrade instructions for miners
permalink: /en/2016/06/21/csv-softfork-instructions/
categories: [mining]
tags: [soft fork, soft forks, bip9, version bits, mining, bip68, bip112, bip113]
version: 1
excerpt: There is an ongoing soft fork of the Bitcoin consensus rules. While everything appears to be proceeding well, this article contains important information and checklists for miners and pool operators which must not be ignored.
---
{% include _toc.html %}

There is an ongoing soft fork of the Bitcoin consensus rules. While everything appears to be proceeding well, this article contains important information and checklists for miners and pool operators which must not be ignored.

If there is any doubt, miners and pool operators are welcome to [contact us][].

TL;DR 

Miners must make the following changes by block #419328.  Any miners who fail to make these changes may generate invalid blocks, causing them lost revenue and possibly disrupting other users and miners on the network.

1. Check all your nodes have been upgraded to Bitcoin Core 0.12.1 or compatible software.

1. If possible, automatically set the block version from Bitcoin Core.  If not possible, set the block version to 0x20000000.  You may do this any time before block 419328.  ([Details](#hardcoded-version))

1. If you generate coinbase transactions whose nSequence is not 0, ensure that the nVersion of the coinbase transaction is 1.  ([Details](#nsequence))

1. If you generate coinbase transactions where the nLockTime is greater than 500000000, ensure that whatever value you set it to passes the [BIP113][] rules.  (The commonly-seen value of 0xffffffff always passes the BIP113 rules.) ([Details](#nlocktime))

## Status of CSV soft fork

The "CSV" soft fork has reached the "locked in" threshold required to proceed to activation. Out of the 2016 blocks from #415296 to #417311, 1946 (96.53%) signaled the readiness for the [BIP68][], [BIP112][] and [BIP113][] (“CSV”) softfork. As of block #417312 (2016-06-21 05:18:58 UTC), the CSV softfork is now in a “locked in” grace period for about 2 weeks up to block 419327. After that, the new [BIP68][], [BIP112][] and [BIP113][] rules will be activated and enforced by the network. It has passed the “point-of-no-return” and is irreversible without a massive rollback of the blockchain.

## For all miners

During the grace period, all miners must upgrade to Bitcoin Core 0.12.1 or any implementation which supports the CSV softfork. In practice, at the time of writing, Bitcoin Core 0.12.1 is the only version that supports the CSV softfork. Miners must double check to make sure all the mining nodes and backup nodes have been upgraded. Failing to do so may result in generation of invalid blocks, or cause your nodes to build upon any invalid blocks causing chain forks and monetary loss to the concerned miners and general Bitcoin users. 

## For miners who manually hardcoded the block version {#hardcoded-version}

By default, Bitcoin Core automatically set and unsets version bits as required, however, we are aware some miners hardcode the  block version numbers. We strongly advise against hardcoding the block version because it can introduce risk to the Bitcoin system because the version signals support for certain consensus rules.

If a miner inadvertently has any nodes that don't support the rules indicated by the block version, it could cause invalid blocks to be produced, and it could cause the miner to follow and build upon an invalid chain. In short, by not using the default value provided by bitcoind, it increases the risk of decoupling of block rules signaling and block rules enforcement.

Unlike the IsSuperMajority softfork used in BIP33/66/65, in the [BIP9][] softfork system, no blocks will be orphaned due to a wrong version number (as long as the version is >= 4, which is required by [BIP65][]). Therefore, there should be no incentive for miners to hardcode the block version, which would unnecessarily increase the burden of maintenance and risks of human error.

However, if you are manually setting the block version against this recommendation, you MUST take specific action. Now that the "point of no return" grace period has been reached for CSV, you must unset the CSV version bit, bit 0. This means if you were signaling 0x20000001 you should signal 0x20000000. This **MUST** be changed before block #419328 or you will trigger "unknown softfork" messages in the logs of all BIP9 compliant nodes. For more information please see the [Version Bits FAQ][].

Failing to follow this advice may trigger the upgrade warning system of all BIP9 compliant nodes on the network, which will be very disruptive.

For miners that allow bitcoind to set the block version automatically, no further action is required. Please note it will keep generating blocks with version 0x20000001 until block #419328 at which point is will automatically unset bit 0.

## For miners who are using the nSequence field of the coinbase transaction {#nsequence}

Some miners may use the nSequence field of the coinbase transaction as the mining nonce. With the activation of [BIP68][], such miners must pay extra attention.

If miners are interfering with the nSequence of the coinbase transaction in any manner, they must make sure that the nVersion of the coinbase transaction (*not* the block version) must be 1 or below. One way to ensure this is to hardcode the nVersion of the coinbase transaction as 1. Failing to do so may result in generation of invalid blocks, causing a chain fork and monetary loss of the concerned miners and general Bitcoin users.

No further action is required for miners who are using the default coinbase nSequence and nVersion values provided by bitcoind.


## For miners who are interfering with the nLockTime field of the coinbase transaction {#locktime}

This is less common since the Stratum protocol does not support the use of nLockTime field as the mining nonce. But miners doing so must pay extra attention due to the activation of [BIP113][].

If a miner is interfering with the nLockTime of the coinbase transaction in any manner, they must make sure that the value, if interpreted as an UNIX timestamp (i.e. >= 500000000), must be smaller than the median timestamp value of the past 11 blocks, unless the nSequence of the coinbase transaction is exactly 0xffffffff. Failing to do so may result in generation of invalid blocks, causing chain fork and monetary loss of the concerned miners and general Bitcoin users.

No further action is required for miners who are using the default coinbase nLockTime values provided by bitcoind.

[contact us]: /en/contact/
[version bits faq]: /en/2016/06/08/version-bits-miners-faq/#when-should-miners-set-bits

{% include _references.md %}

