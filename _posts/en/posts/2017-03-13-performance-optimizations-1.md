---
type: posts
layout: post
lang: en
name: performance-optimizations-1
id: en-performance-optimizations-1
title: On-chain scaling - a review of historical performance optimization made to Bitcoin’s reference software. Part 1
permalink: /en/2017/03/13/performance-optimizations-1/
version: 1
excerpt: Development milestones that helped preserve a reliable experience for users of the Bitcoin software client over the years. 
---

The following post aims to highlight development milestones that helped preserve a reliable experience for users of the Bitcoin software client over the years. We present several upgrades that were critical in maintaining the decentralized properties of the network and mitigate the resources burden of its participants. We describe how numerous orders-of-magnitude optimizations were made so that the Bitcoin network could support the growth in transaction activity without dramatically increasing the costs of participation for new and existing users. Finally, we note how those improvements all fall within a larger, systematic approach to protocol development that uses insights from Big-O complexity concepts and leverages smarter algorithms that make more efficient use of the network’s resources. 

## Signature Caching
Release: Bitcoin-Qt 0.7.0

Verification of ECSDA signatures are one of the most computationally heavy operations done at the peer level. Because they need to be verified for every transaction, any superfluous validation leads to significant resource overhead for the end user. That was the case for early versions of the software which would both verify signatures before they entered a node mempool and after they were received into a block. 

In order to gain in efficiency, developers created a cache allowing nodes to store previously validated signatures and skip redundant work once the transactions make it into an accepted block. 

Additionally, the signature cache also mitigates a DoS vector introduced by the potential for specifically crafted transactions to stall Bitcoin clients. 

### Further information

 * [Bitcoin-Qt 0.7.0 Release notes](https://bitcoin.org/en/release/v0.7.0#core-bitcoin-handling-and-blockchain-database)
 * [Fixed vulnerability explanation: Why the signature cache is a DoS protection](https://bitcointalk.org/index.php?topic=136422.0)

## Ultraprune + LevelDB
Release: Bitcoin-Qt 0.8.0

Ultraprune was one of the first major upgrades to the Bitcoin software aimed at solving the overhead associated with validating transaction data from the blockchain. Previous versions of the Bitcoin reference client maintained an index of all transaction outputs, spent or unspent. Ultraprune significantly reduced the size of that index with the insight that you only needed to keep track of unspent outputs, and an output - once spent - can be removed from the indexes entirely.

To achieve this, a new database layout was implemented which allocates unspent transaction outputs to a compact custom format in order to reduce the size of information required for validation work. 

To further optimize the performance of the system, Ultraprune was introduced in parallel with LevelDB, which deprecated the old BDB database technology. Overall, the impact was notable: depending on their hardware, users could experience at least an order of magnitude improvement when validating blockchain data. This new database structure would also pave the way for future work on pruning and lighter implementations of Bitcoin full nodes. 

### Further information

  * [Bitcoin-Qt 0.8.0 Release notes](https://bitcoin.org/en/release/v0.8.0#improvements)
  * [Ultraprune in plain english](https://archive.is/bUocJ)
  * [Ultraprune merged in mainline](https://bitcointalk.org/index.php?topic=119525.0)
  * [Pruning in the reference client: ultraprune mode](https://bitcointalk.org/index.php?topic=91954.0)

## Parallel script verification
Release: Bitcoin-Qt 0.8

While a more subtle change, transitioning script verification to a more parallelized process removed significant overhead from block validation times. Early versions of the software would validate script data from inputs in between every UTXO fetch, creating a performance issue because of the linear processing of all actions. This violates a basic principle for the design of efficient computing processes, which dictates that computation should happen concurrently with I/O jobs, where possible. 
With that in mind, the block validation mechanism was re-engineered in order to be able to allocate script checks to parallel threads so that their verification can happen even before the client is done fetching all of the UTXOs from the block. To achieve this, script check actions are stored in a queue after transaction are processed and are handled separately from other input validation jobs. 

As a consequence, synchronization to the tip of the chain happens much faster by making more efficient use of the peer’s resources. During testing of the implementation, developers noted 35% to 100% speed-up when benchmarking against previous versions of the software.

### Further information

  * [Parallel script verification #2060](https://github.com/bitcoin/bitcoin/pull/2060)

## Headers first synchronization
Release: Bitcoin Core 0.10

Striving to further improve initial block download time, the Core project introduced in late 2014 an important re-architecture of the mechanism used by nodes to synchronize with the most-work valid chain. 

Initially, the process of bootstrapping a new Bitcoin client would involve a user fetching block data from a single peer with the consequence that any interruption or decrease in connection quality would significantly stall the process. With an ever-increasing blockchain size, this would result in sometimes massive waiting time for the synchronization to complete, with a large percentage of users reporting up to multi-day periods depending on their hardware.

Headers first synchronization completely mitigates this issue using a new method that involves nodes first downloading and validating block headers from a single peer and then fetching block data in parallel from a multitude of others. 

Complaints about initial block download time have been prevalent since the early days of Bitcoin. With headers first synchronization, the software took a major step forward in terms of usability for new users. Rather than wasting many hours on unreliable synchronization, nodes could now leverage their entire network of peers and cut down the bootstrapping time significantly. With the use of smarter algorithms, another asymptotic improvement was made to this crucial aspect of Bitcoin’s long term sustainability. 

### Further information

  * [Bitcoin-Qt 0.10.0 Release notes](https://bitcoin.org/en/release/v0.10.0#faster-synchronization)
  * [Bitcoin.org Developer Guide](https://bitcoin.org/en/developer-guide#headers-first)
  * [Pieter Wuille's post to the Bitcoin-dev mailing list](https://lists.linuxfoundation.org/pipermail/bitcoin-dev/2014-October/006724.html)
  
## Block file pruning
Release: Bitcoin Core 0.11

Pruning of old data was a concept first described by Satoshi Nakamoto in his white paper as a potential solution to disk space scarcity. Unfortunately, the original design was inadequate and could not be implemented as imagined by its creator. Seven years later, with the blockchain reaching more than a hundred gigabytes, the introduction of block file pruning as we know it today presents a major boon to users with limited resources.

Block file pruning leverages previous work with ultraprune; users who have initially downloaded and validated the blockchain may now discard raw data older than 288 blocks. Because those nodes still preserve the full UTXO set, they remain able to validate unspent data, which is enough to fully validate new blocks and protect against potential double-spends. 

Of course, pruning implies that there remains a sufficient number of archival nodes around to serve full blockchain data. On the other hand, this innovation expands the range of validators by making it more cost-effective to remain one. As a whole, the solution is a welcome addition to the options available for us to bolster network diversity. 

### Further information

  * [Bitcoin-Qt 0.11.0 Release notes](https://bitcoin.org/en/release/v0.11.0#block-file-pruning)

## libsecp256k1
Release: Bitcoin Core 0.12

After measurements, it was determined that the next step after solving the inefficiencies of blockchain download was to tackle the bottleneck of transaction verification and its heavy computing load. The Core project set out to do this by using a new library designed for optimized performance of ECDSA operations. ECDSA (Elliptic Curve Digital Signature Algorithm) is the backbone of Bitcoin’s public key infrastructure and is used every time a user moves coins by signing a message with their private keys. These signatures need to be verified by every peer in the network in order to preserve the ledger’s integrity.

Early developers had long considered transitioning away from the original OpenSSL library and after 5 years of design considerations, testing and peer-review, Pieter Wuille’s libsecp256k1 library was introduced as its replacement. As expected, the implementation led to major speed-up of the signature validation process behind every Bitcoin transaction. Benchmarks reported 5–7x improvements over the OpenSSL implementation. For users this would translate to saving up to half the bootstrap time typically dedicated to ECSDA operations, one of the most laborious steps in synchronizing a new node from scratch. 

Considering the growth in Bitcoin transaction activity, this upgrade was essential to preserving a reasonable user experience for network peers. Once again, reduction of algorithmic complexity provided users with more efficient usage of their resources and drastically lowered the barrier of entry for new participants. 

### Further information

  * [Bitcoin-Qt 0.12.0 Release notes](https://bitcoin.org/en/release/v0.12.0#signature-validation-using-libsecp256k1)
  * [Andrew Poelstra (andytoshi) on security and testing of libsecp256k1](https://bitcointalk.org/index.php?action=profile;u=80376)
  * [Greg Maxwell on testing of libsecp256k1 revealing bug in OpenSSL](https://www.reddit.com/r/Bitcoin/comments/2rrxq7/on_why_010s_release_notes_say_we_have_reason_to/)
  * [Greg Maxwell presentation at DevCore](https://www.youtube.com/watch?v=RguZ0_nmSPw&t=1297)
  * [Hal Finney post on libsecp256k1](https://bitcointalk.org/index.php?topic=3238.0)

## Memory pool limiting
Release: Bitcoin Core 0.12

A long standing vulnerability of the Bitcoin software was its inability to properly deal with the flooding of a peer’s memory pool. Indeed, an attacker could send a high number of low value, low fee transactions that would accumulate in the memory pool until it would overload the memory available. This would cause nodes with relatively low RAM resources to crash during periods of unusual activity. The only effective measure against this was to increase the software’s minimum relay fee which still left no upper bound on the potential size of the mempool.

To remediate this, developers of the Core project implemented a strict memory pool maximum size with specific eviction policies sorting transactions by fees and evicting the lowest paying ones first. In order to prevent transactions with the same fee from re-entering the memory pool, the node will increase its effective minimum relay feerate to match the one of the last evicted transaction plus the initial minimum relay feerate. 

Configuration of the maximum size is left to the users with the default size being 300 megabytes. This update provides a much more robust experience for node users with limited resources and, in general, makes the entire network more reliable.

### Further information

* [Bitcoin-Qt 0.12.0 Release notes](https://bitcoin.org/en/release/v0.12.0#memory-pool-limiting)

In Part 2, we will discuss more recent improvements that build on the technologies presented above and further improve the robustness and scaling potential of the network. 
