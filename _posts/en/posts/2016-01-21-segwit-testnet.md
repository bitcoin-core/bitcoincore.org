---
layout: post
id: en-segwit-testnet
name: segwit-testnet
title: Launch of Segregated Witness Testnet
permalink: /en/2016/01/21/launch_segwit_testnet
version: 1
excerpt: We are extremely pleased and excited to announce the release of the Segregated Witness Testnet
---

We are extremely pleased and excited to announce the [Segregated Witness Testnet](https://github.com/sipa/bitcoin/commits/segwit)! All developers are encouraged to help begin testing and integration right away.

This represents one of the most hotly anticipated and exciting developments to date that is an important foundation for many future improvements and innovations. Segregated Witness frees up space on the Bitcoin blockchain by securely moving transaction signature data to a specially delegated “Segregated Witness” data structure outside of the transaction block.

This significant innovation leads to dramatically more efficient use of the Bitcoin blockchain, while simultaneously opening up exciting new possibilities and opportunities for the broader Bitcoin ecosystem, particularly around smart contract applications and dramatically faster transactions, while also providing the groundwork for more advanced features and possibilities in the future.

This development began with a research effort by Bitcoin Core developer, Dr. Pieter Wuille, initially focused on addressing transaction malleability, a longstanding and well-known concern and priority. However, in the process of this research, and in the narrowing toward a solution, additional properties of the solution were discovered that allow for increasing the block size while also simultaneously opening up some incredibly exciting secondary benefits.

This effort was initiated by Dr. Pieter Wuille, but included contributions from many others, with particular thanks to Eric Lombrozo, Johnson Lau, Alex Morcos, Nicolas Dorier, Bryan Bishop, Gregory Maxwell, Peter Todd, Cory Fields, Suhas Daftuar, and Luke-Jr.

## Bitcoin Core Ecosystem

There is broad excitement and anticipation as far as what providers and other exchange operators will create with the fundamental developments and innovations included in this release. So far, the [most popular wallets][segwit_adoption] and supporting libraries have stated they will support segwit including Ledger, Trezor, Electrum, and Bitgo. Additionally, work on numerous other libraries such as bitcoinj, bitcoinjs, pycoin and bitcore has already begun.

A faucet is available for "segnet" coins [here](https://segwit.greenaddress.it/faucet/).

Early previews of third party wallet support are available at:

- [mSIGNA](https://github.com/ciphrex/mSIGNA/tree/segwit) (wallet source-code)
- [Green Address](https://segwit.greenaddress.it/) (web wallet)

## How to get involved

Please join the `segwit-dev` IRC channel on irc.freenode.net.

Wallet providers should read the [migration guide](/en/segwit_wallet_dev).

If you are a wallet user interested in following the status of your wallet’s implementation, please [follow their progress][segwit_adoption] which will be regularly updated. 

## Testing

Finally and most importantly, please help test the Segwit Testnet!

The source-code can be located [here](https://github.com/sipa/bitcoin/tree/segwit) checkout the `segwit` branch.

Once compiled, add `-segnet` to the standard `bitcoind` and `bitcoin-cli` command line.

## Additional Background and History

- [Scaling Bitcoins Hong Kong Presentation](https://prezi.com/lyghixkrguao/segregated-witness-and-deploying-it-for-bitcoin/)
- [Extended Video](https://bitcoincore.org/en/2015/12/14/segregated-witness)
- [Transcript](http://diyhpl.us/wiki/transcripts/scalingbitcoin/hong-kong/segregated-witness-and-its-impact-on-scalability/)

## Technical References

### SegWit BIPs

- [BIP141](https://github.com/bitcoin/bips/blob/master/bip-0141.mediawiki): An overview of the segregated witness soft fork, and a discussion on its benefits, backwards compatibility, consensus limits, and deployment.
- [BIP143](https://github.com/bitcoin/bips/blob/master/bip-0143.mediawiki): The definition of a new digest algorithm for signature verification used in segregated witness transactions.
- [BIP144](https://github.com/bitcoin/bips/blob/master/bip-0144.mediawiki): The new message types and serialization formats for segregated witness transactions.

### References

- [Analysis of segwit benefits](http://lists.linuxfoundation.org/pipermail/bitcoin-dev/2016-January/012248.html)
- [Hard Forks vs. Soft Forks](https://petertodd.org/2016/soft-forks-are-safer-than-hard-forks)
- [Early Exploration of “Non-contentious” Hard Fork Research](https://scalingbitcoin.org/hongkong2015/presentations/DAY1/1_overview_1_timon.pdf)

[FAQ]: https://bitcoincore.org/en/2015/12/23/capacity-increases-faq
[roadmap]: http://lists.linuxfoundation.org/pipermail/bitcoin-dev/2015-December/011865.html
[segwit_adoption]: /en/segwit_adoption

