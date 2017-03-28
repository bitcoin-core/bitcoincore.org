---
type: posts
layout: post
lang: en
name: hybrid-spv
id: en-hybrid-spv
title: Technology roadmap: Prioritized block download with hybrid full block SPV mode
permalink: /en/2017/03/30/hybrid-spv
version: 1
excerpt: Prioritized block download with hybrid full block SPV mode
---

## Hybrid full block SPV mode

One of the major hurdle hindering further adoption of fully validating software by regular users is the inability to use the wallet features of the client until it has fully synced the entire blockchain. For users bootstrapping a new node, this means that they are unable to receive or send transactions until every block has been downloaded and validated up to the current tip of the chain. This behaviour is not by mistake: the Bitcoin Core reference software, by default, is built to offer the strongest security and privacy guarantees a Bitcoin user can expect and this necessarily implies full validation in order to confirm the integrity of historical blockchain data. 

On the other hand, existing features of the software such as headers-first validation provide an opportunity to improve the usability of the wallet provided users are willing to make a temporary security tradeoff. Using the hybrid full block SPV mode, the software will prioritize download of blocks according to the oldest key in the user’s wallet. Along with the previously downloaded block headers chain, which should meet expected Proof-Of-Work difficulty checks, the client can then immediately start processing relevant transactions. The entire blockchain is still downloaded and eventually validated in parallel but this feature enables users to see and spend UTXOs associated with their wallet while synchronization is happening in the background. 

Contrary to typical implementation of SPV wallets, this model does not suffer from the [privacy degradation](http://bitcoin.stackexchange.com/questions/37756/are-public-keys-and-their-corresponding-hash-values-both-added-to-a-bitcoinj-blo) imposed on schemes relying on bloom filters and public disclosure of public keys. This benefit comes with a tradeoff which is that it consumes more bandwidth. Another caveat: confirmations received under SPV mode are inherently less safe than those received under full validation. A user leveraging the hybrid SPV mode should wait for several confirmations (6+) until his payment can be considered secure. 

### Further information
  * [Experimental implementation PR](https://github.com/bitcoin/bitcoin/pull/9076)
  * [Complete patch-set PR](https://github.com/bitcoin/bitcoin/pull/9483)
  * [Perfect privacy for received transactions](https://bitcoin.org/en/bitcoin-core/features/privacy#perfect-privacy-for-received-transactions)
 
