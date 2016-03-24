---
type: posts
layout: post
lang: fr
id: fr-segwit-testnet
name: segwit-testnet
title: Lancement du Testnet Segregated Witness
permalink: /fr/2016/01/21/launch_segwit_testnet/
version: 1
excerpt: C'est avec un grand plaisir et une certaine excitation que nous vous annonçons la sortie du Testnet Segregated Witness
---

C'est avec un grand plaisir et une certaine excitation que nous vous annonçons la sortie du [Testnet Segregated Witness](https://github.com/sipa/bitcoin/commits/segwit) !  Tous les développeurs sont invités à se joindre à nous dès maintenant en commençant à tester et intégrer.

Il s'agit d'un des développements les plus attendus et passionnants à ce jour qui servira de base à de nombreuses améliorations et innovations à venir. Segregated witness libère de l'espace sur la blockchain Bitcoin en déplaçant de manière sécurisée les données de signature des transactions vers une nouvelle structure de données conçue à cet effet dénommée "segregated witness" à l'extérieur du bloc de la transaction.

Cette innovation significative va permettre une utilisation beaucoup plus efficiente de la blockchain Bitcoin, tout en ouvrant la porte à de nouvelles possibilités et opportunités pour l'ensemble de l'écosystème Bitcoin, en particulier dans le domaine des applications de 'smart contract' et de la performance des transactions, tout en posant les bases pour des fonctionnalités encore plus avancées dans le futur.

Cette évolution a débuté par un travail de recherche du développeur Bitcoin Core, le Dr. Pieter Wuille, concentré au départ sur la correction de la malléabilité des transactions, un problème préoccupant et prioritaire depuis longtemps.  Cependant, au cours de ses recherches, tandis qu'il se rapprochait de plus en plus d'une solution, de nouvelles propriétés de la solution ont été découvertes comme l'augmentation de la taille des blocs et d'autres avantages annexes incroyables et passionnants. 

Cet effort initié par le Dr. Pieter Wuille a reçu les contributions de nombreuses autres personnes.  Nous souhaitons remercier en particulier Eric Lombrozo, Johnson Lau, Alex Morcos, Nicolas Dorier, Bryan Bishop, Gregory Maxwell, Peter Todd, Cory Fields, Suhas Daftuar, et Luc-Jr.

## L'écosystème Bitcoin Core

Il y a beaucoup d'excitation et d'anticipation autour de ce que vont pouvoir créer les fournisseurs et opérateurs d'échange avec les innovations et évolutions fondamentales apportées par segwit.  Jusqu'à présent, les [portefeuilles les plus populaires][segwit_adoption] et les bibliothèques sur lesquels ils s'appuient ont fait part de leur soutien à segwit y compris Ledger, Trezor, Electrum et Bitgo.  En outre, de nombreuses autres bibliothèques comme bitcoinj, bitcoinjs, pycoin et bitcore ont déjà commencé à implémenter segwit.

Un faucet est disponible pour "segnet" [ici](https://segwit.greenaddress.it/faucet/).

Premiers aperçus des portefeuilles tiers compatibles :

- [mSIGNA](https://github.com/ciphrex/mSIGNA/tree/segwit) (wallet source-code)
- [Green Address](https://segwit.greenaddress.it/) (web wallet)

## Comment participer

Rejoignez-nous sur le canal IRC `segwit-dev` sur irc.freenode.net.

Les fournisseurs de portefeuille devraient consulter le [guide de migration](/en/segwit_wallet_dev).

Si vous souhaitez connaître l'état d'avancement de l'implémentation de segregated witness pour le portefeuille bitcoin que vous utilisez, veuillez consulter [cette page][segwit_adoption], elle sera régulièrement mise à jour. 

## Tests

Enfin et surtout, s'il vous plaît, aider nous à tester le Testnet Segwit !

Le code source est disponible [ici](https://github.com/sipa/bitcoin/tree/segwit), checkout la branche `segwit`.

Une fois compilés, ajouter `-segnet` à la ligne de commande standard de `bitcoind` et de `bitcoin-cli`.

## Compléments de contexte et d'historique

- [Présentation réalisée à Scaling Bitcoins Hong Kong](https://prezi.com/lyghixkrguao/segregated-witness-and-deploying-it-for-bitcoin/)
- [Vidéo détaillée](https://bitcoincore.org/en/2015/12/14/segregated-witness)
- [Transcription](http://diyhpl.us/wiki/transcripts/scalingbitcoin/hong-kong/segregated-witness-and-its-impact-on-scalability/)

## Références techniques

### BIPs SegWit

- [BIP141](https://github.com/bitcoin/bips/blob/master/bip-0141.mediawiki): Aperçu du soft fork segregated witness, et discussion sur ses avantages, la rétrocompatibilité, les limites du consensus, et le déploiement.
- [BIP143](https://github.com/bitcoin/bips/blob/master/bip-0143.mediawiki): Définition d'un nouvel algorithme de 'digest' pour la vérification des signatures à destination des transactions segregated witness.
- [BIP144](https://github.com/bitcoin/bips/blob/master/bip-0144.mediawiki): Nouveaux types de message et formats de sérialisation pour les transactions segregated witness.

### Références

- [Analyse des avantages de segwit](http://lists.linuxfoundation.org/pipermail/bitcoin-dev/2016-January/012248.html)
- [Hard Forks vs. Soft Forks](https://petertodd.org/2016/soft-forks-are-safer-than-hard-forks)
- [Premières prospections concernant l'état de la recherche sur les hard forks “non-contentieux”](https://scalingbitcoin.org/hongkong2015/presentations/DAY1/1_overview_1_timon.pdf)

[FAQ]: https://bitcoincore.org/en/2015/12/23/capacity-increases-faq
[roadmap]: http://lists.linuxfoundation.org/pipermail/bitcoin-dev/2015-December/011865.html
[segwit_adoption]: /en/segwit_adoption
