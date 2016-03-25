---
type: posts
layout: post
lang: fr
name: bitcoin-core-capacity-increases-faq
id: fr-bitcoin-core-capacity-increases-faq
title: FAQ sur l'augmentation de la capacité de Bitcoin
permalink: /fr/2015/12/23/capacity-increases-faq/
version: 1
---
{% include _toc.html %}

## Quelles technologies particulières sont inclues dans la feuille de route, et quand seront-elles disponbiles&nbsp;? {#roadmap-dates}

Les nouvelles technologies seront déployées quand elles seront prêtes et testées. Malgré tout, nous pensons que ce qui suit est un planning raisonnable pour les améliorations spécifiques décrites dans [la feuille de route][roadmap].

| Déc 2015 | | Deploiement du testnet segregated witness |
| Fév 2016 | 0.12.0 | [Vérification libsecp256k1][libsecp256k1 verification] |
| Fév 2016 | | Segregated witness finalisé (feature complete) & prêt pour une revue globale |
| Mar 2016\* | 0.12.x | Deploiement de OP_CHECKSEQUENCEVERIFY (BIPs [68][BIP68] & [112][BIP112]) + [BIP113][] comme premier soft fork [BIP9][] versionbits |
| Avril 2016\* |  0.12.x |  Déploiement de segregated witness (incluant l'augmentation de la taille des blocs) |
| 2016 | | Weak blocks et IBLT, Lightning, ou les deux |

\* Les dates avec une astérisque correspondent aux moments où nous serons en mesure de livrer du code prêt pour le soft fork.  Le code ne sera pas livré tant que qu'il n'aura pas été attentivement revu, et le fork réel mettra un certain temps à s'activer ([BIP66][] a été activé en Juillet 2015 après quelques mois;  [BIP65][]  a été activé en Décembre 2015 après seulement 5 semaines).

- **Segregated witness testnet:** un testnet distinct (ne faisant pas partie du testnet habituel) qui permet aux contributeurs de Bitcoin Core de tester segregated witness et aux éditeurs de portefeuilles de commencer à travailler dessus.

- ** Vérification avec [Libsecp256k1][] :** une vitesse de vérification améliorée de 500% à 700% sur un matériel x86\_64 afin de permettre à de nouveaux nœuds de rejoindre le réseau et d'alléger la charge des nœuds existants.

- **[OP\_CHECKSEQUENCEVERIFY][BIP112]:** une amélioration de 25000% de [l'efficacité des canaux de paiement][payment channel efficiency] bi-directionnels en permettant aux utilisateurs de garder ouverts leur canaux de paiement aussi longtemps qu'ils le souhaitent.

- **[VersionBits][BIP9]:** augmente de 1 à 29 le nombre maximum de soft forks pouvant être déployés simultanément, permettant de futures mises à jour du réseau plus rapides et plus décentralisées. 

- **[Segregated witness][bip-segwit]:** une augmentation directe de la capacité de 175% à 400%, une amélioration de l'efficacité des canaux de paiement de 66% en consolidant les opérations d'ouverture et de fermeture, la fin de la malléabilité causée par autrui qui empêche le déploiement de contrats intelligents, des preuves de fraude qui permettent aux clients légers de mieux participer à l'application des règles économiques, et la possibilité d'améliorer le langage de Script de Bitcoin, permettant d'envisager de nouveaux types de contrat qui ne nécessitent pas de tiers de confiance.

- **IBLTs and weak blocks:** 90% or more reduction in critical bandwidth to relay blocks created by miners who want their blocks to propagate quickly with a modest [increase in total bandwidth][], bringing many of the benefits of the [Bitcoin Relay Network][] to all full nodes. This improvement is accomplished by spreading bandwidth usage out over time for full nodes, which means IBLT and weak blocks may allow for safer future increases to the max block size.

## Le soft fork segregated witness est-il équivalent à une augmentation de la taille des blocs de 4 Mo, de 2 Mo, de 1.75 Mo ou bien autre chose encore ? Je n’arrête pas d’entendre des valeurs différentes. {#segwit-size}

La [proposition actuelle][current proposal] pour le soft fork segregated witness (segwit) remplace la taille limite d'un bloc par une nouvelle limite de *coût* d'un bloc, comptabilisant chaque octet des données "témoin" comme 1 unité de coût et les données de transaction UTXO comme 4 unités; ce qui a pour conséquence d'augmenter la taille maximale effective d’un bloc à presque 4 Mo.

Cependant, les blocs ne contiendront pas uniquement des données ‘témoin’, de sorte qu’il est peu probable d’obtenir des blocs s'approchant de 4 Mo.

Selon certains [calculs][calculations] effectués par Anthony Towns, un bloc rempli de transactions standards P2PKH à une seule signature ferait environ 1,6 Mo et un bloc rempli de transactions multi-signatures 2-de-2 ferait environ 2,0 Mo.
Il est en outre probable que de futures améliorations de la scalabité, comme Lightning, pourront améliorer légèrement le ratio de sorte que les blocs remplis soient supérieurs à 2 Mo.

[current proposal]: https://youtu.be/fst1IK_mrng?t=2234
[calculations]: http://lists.linuxfoundation.org/pipermail/bitcoin-dev/2015-December/011869.html

## Segregated witness semble compliqué; l'éco-système sera-t-il prêt à le déployer&nbsp;?  {#ecosystem-ready}

Certaines idées sont faciles à expliquer mais difficiles à implémenter. D'autres sont faciles à implémenter mais difficiles à expliquer. Segregated witness (segwit) rentre dans la seconde catégorie.

Segwit peut être déployé de manière incrémentale sans briser la compatibilité, ce qui signifie qu'aucune préparation de l'écosystème n'est nécessaire. Les développeurs souhaitant immédiatement se faire la main avec segwit ont commencé à tester leur logiciel sur le testnet segwit déployé en Décembre 2015.

Initialement, seuls les mineurs qui souhaitent soutenir segwit doivent se mettre à niveau pour l’activer et de l’appliquer sur le mainnet. Les applications existantes ne doivent être modifiées que si elles souhaitent profiter des nouvelles fonctionnalités et de l'espace additionnel des blocs.

Les transactions Segregated witness nécessiteront moins de frais, vont permettre des optimisations de performances bien plus importantes, et vont pouvoir couvrir des contrats intelligents en plusieurs étapes ainsi que des protocoles comme les canaux de paiements bi-directionnels qui peuvent accroître les capacités sans écrire de données supplémentaires dans la blockchain. Les portefeuilles sont fortement encouragés à se mettre à jour mais peuvent continuer à opérer sans modification puisque le déploiement préserve la rétro-compatibilité.

## Segregated witness parait vraiment compliqué.  Pourquoi ne pas simplement augmenter la taille maximum des blocs&nbsp;?  {#size-bump}

Il n’y a qu’une [seule ligne de code][max_block_size] dans Bitcoin Core qui spécifie que la taille maximale d’un bloc est de 1.000.000 octets (1 Mo). La chose la plus simple à faire serait de faire un hard fork qui modifierait cette ligne en mettant, par exemple, 2.000.000 octets (2 Mo).

Les hard forks sont tout sauf simple :

- **Nous n’avons pas d’expérience :** les mineurs, les marchands, les développeurs, et les utilisateurs n’ont jamais déployé de hard forks non-urgents, donc les mécanismes pour les déployer sans risques n’ont pas été testés.

    C'est ce qui les différencie des soft forks dont les déploiements ont été gérés au départ par Nakamoto, où nous avons acquis de l'expérience suite aux complications lors du déploiement de [BIP16][], où nous avons affiné nos techniques lors du déploiement de [BIP34][], et où nous avons acquis suffisamment d'expérience avec les BIPs [66][BIP66] et [65][BIP65] pour envisager de commencer à pouvoir déployer plusieurs soft forks en parallèle avec [BIP9][] versionBits.

- **Mise à niveau obligatoire&nbsp;:** les hard forks imposent que tous les nœuds soient mis à jour au risque de voir les utilisateurs de ces nœuds perdre de l'argent.  Les gestionnaires de nœuds sont concernés a partir du moment où leurs portefeuilles sont sécurisés par ces nœuds, mais aussi tous les clients légers qui seraient connectés sur ces mêmes nœuds.

- **D’autres changements sont requis :** même un changement aussi simple que la modification d’une seule ligne de code pour augmenter la taille maximale d’un bloc peut avoir des effets indésirables sur d’autres parties du code.  Il est par exemple actuellement possible de créer une transaction qui occupe presque 1 Mo d’espace et qui mette plus de 30 secondes pour être validée sur un ordinateur récent (des blocs contenant de telles transactions ont déjà été minés).  Avec des blocs de 2 Mo, une transaction de 2 Mo pourrait être créée et prendre plus de 10 minutes pour être validée. Ce qui ouvre de dangereux vecteurs d’attaque par déni de service.  D’autres parties du code devront donc être modifiées pour éviter ce genre de problèmes.

Malgré ces complications conséquentes, en prenant toutes les précautions nécessaires, aucune d'entre elles n'est fatale à un éventuel hard fork, et nous nous attendons à faire des hard forks dans le futur.  Mais segregated witness (segwit) est un soft fork, semblable aux autres soft forks que nous avons déjà effectués et pour lesquels nous avons acquis un certaine expérience.  Ce soft fork nous offrira de nombreux avantages, en plus d'augmenter la capacité en transaction de la blockchain.

Segwit nécessite plus de changements dans les couches hautes des piles logicielles qu'un simple augmentation de taille des blocs, mais si nous voulons vraiment voir Bitcoin s'adapter à l'augmentation des usages, des changements beaucoup plus invasifs seront nécessaires de toute façon, et segwit va permettre d'encourager les gens à passer rapidement à des modèles plus évolutifs sans les forcer.

Les développeurs, les mineurs et la communauté ont accumulé une expérience considérable dans le déploiement des soft forks et nous croyons que segwit peut être déployé au moins aussi rapidement et probablement de manière plus sûre qu'un hard fork d'augmentation de la taille des blocs.

##  Y a-t-il un hard fork de prévu avant ou pendant l'implémentation de segregated witness&nbsp;?  {#pre-segwit-fork}

Non.  Cela ne fait pas partie de la [feuille de route][roadmap].

[roadmap]: https://lists.linuxfoundation.org/pipermail/bitcoin-dev/2015-December/011865.html

## Si l'éventualité d'un hard fork n'est pas totalement écartée, pourquoi ne pas le faire maintenant&nbsp;? {#why-not-now}

Nous avons actuellement la possibilité d'augmenter la capacité du système à l'aide de soft forks qui bénéficient d'un large consensus sans avoir les complications d'un hard fork tel que décrit dans une [précédente question][q simple raise].   Donc la perspective d'un hard fork éventuel n'est pas une raison suffisante pour en tenter un maintenant.

En plus d'augmenter la capacité de transaction, les améliorations proposées dans la feuille de route (combinées avec d'autres technologies telles que les canaux de paiement bi-directionnels) donnent aux utilisateurs la possibilité de réduire la quantité moyenne d'espace de la blockchain qu'ils utilisent---augmentant efficacement la capacité du système Bitcoin sans augmenter la quantité de bande passante utilisée par un nœud complet (full node).

Par exemple,

- [BIP68][] and [BIP112][] permettent aux canaux de paiement bi-directionnels de rester ouvert indéfiniment, ce qui devrait réduire considérablement le nombre de transactions de canal de paiement qui doivent être confirmées dans la blockchain.

- Segregated witness permet à une transaction de fermeture de paiement d'être combinée avec une transaction d'ouverture de canal de paiement, réduisant ainsi l'espace de la blockchain utilisé pour changer de canaux d'environ 66%. 

- Segregated witness permet à des soft forks de changer le langage de script de Bitcoin de manière à pouvoir réduire la taille moyenne d'une transaction, comme l'utilisation de 'public-key-recovery-from-signatures' ou bien l'utilisation des signatures de Schnorr combinées.

- Segregated witness permet la création de preuves de fraude compactes qui pourraient augmenter la sécurité des clients légers Simplified Payment Verification (SPV) jusqu'à s'approcher de celle des nœuds complets (full nodes), ce qui pourrait permettre au réseau de fonctionner correctement avec moins de nœuds complets qu'il ne le peut avec la technologie actuellement déployée.

L'effet réel de ces technologies est inconnu, mais augmenter la capacité maintenant avec un soft fork qui disposerait d'un large consensus nous permettrait d'obtenir des gains immédiats, de tester et de mesurer les possibilités à moyen terme, et d'utiliser les données obtenues pour formuler des plans à plus long terme.

## Comment fonctionnera segregated witness pour les portefeuilles&nbsp;?  {#segwit-in-wallets}

les portefeuilles qui supportent actuellement P2SH peuvent migrer vers segregated witness en deux étapes :

- Phase 1 : les scripts sont hachés à deux reprises, d'abord à 256 bits et puis 160 bits.  Le hachage de 160 bits sera compatible avec les adresses P2SH existantes, afin que les portefeuilles mise à jour soient en mesure d'envoyer et de recevoir des bitcoins vers et depuis les portefeuilles existant actuellement.

- Phase 2: les scripts sont hachés une seule fois à 256 bits.  Ce format ne sera pas compatible avec les portefeuilles existants, mais permettra une utilisation plus efficace de l'espace de bloc et offrira une meilleure sécurité due à une plus grande résistance à la collision.

## Si la mise à jour n'est pas obligatoire, qui va s'embêter à la faire&nbsp;?  J'ai entendu dire que P2SH avait mis presque deux années avant d'être largement déployé. {#why-upgrade}

Chaque octet de la partie 'witness' d'une transaction segregated witness (segwit) ne compte que pour 0.25 octets par rapport à la taille globale de la transaction.  Puisque les frais de transaction sont basés sur la taille d'une transaction, cela correspond effectivement à une réduction de 75% des frais pour cette partie de la transaction---mais uniquement pour les personnes qui utilisent segwit.

David Harding a réalisé un tableau des [économies estimées][estimated savings] pour différents niveaux de frais/transaction.  Par exemple, si les frais pour une transaction de 250 octets s'élèvent à 0.01 $ US, en utilisant segwit on économisera environ 0.003 $ lors de la dépense d'un 'output' d'une transaction P2PK-in-P2SH.

| Transaction | Octets éco. | $0.01/250B | $0.05/250B | $0.25/250B | $1.00/250B |
|-------------|-------------|------------|------------|------------|------------|
| P2PK-in-P2SH |  79/107 | $0.003 | $0.015 | $0.079 | $0.316 |
| 1-of-1 P2SH multisig | 83/112 | $0.003 | $0.016 | $0.083 | $0.332 |
| 2-of-2 P2SH multisig | 163/219 | $0.006 | $0.032 | $0.163 | $0.652 |
| 2-of-3 P2SH multisig | 189/254 | $0.007 | $0.037 | $0.189 | $0.756 |

(Nous ne pensons pas que les frais puissent être aussi élevés que les montants les plus élevés présentés dans ce tableau; ils sont juste fournis à titre d'information.)

Les portefeuilles web et les échanges qui envoient un grand nombre de transactions chaque jour à des taux fixes (gratuitement ou pour 1% par dépenses par exemple) devraient être les premiers à adopter segwit---même les petites économies par dépense montrées dans le tableau ci-dessus peuvent s'accumuler et devenir d'importantes sommes d'argent si elles sont répétées des centaines ou des milliers de fois par jour.

## J'ai entendu dire que les transactions avec zéro-confirmation ne seraient plus possibles.  Quelle est la technologie de la feuille de route en cause&nbsp;?  {#rbf}

Aucune. Par défaut, les versions actuelles de Bitcoin Core ne remplaceront pas une transaction non confirmée par une autre transaction qui dépense les mêmes entrées (inputs). Certains pensent que cela signifie que la première transaction qu'ils voient dépenser une entrée particulière (input) est sûre, mais cela est faux. (Si cela était vrai, nous n'aurions pas besoin de la blockchain.)

La politique par défaut actuelle ne signifie pas que les personnes qui souhaitent être en mesure de mettre à jour leurs transactions non confirmées ne puissent pas le faire. La version originale de Bitcoin donnait la possibilité d’indiquer qu’une transaction pouvait être mise à jour, mais Nakamoto a dû désactiver la fonctionnalité en 2010 pour empêcher des attaques par déni de service.

Les développeurs de Bitcoin Core ont récemment réalisé qu’ils pouvaient empêcher cette attaque DOS en exigeant que les transactions mises à jour payent des frais supplémentaires, et ils ont réactivé le mécanisme de Nakamoto qui permet d'indiquer qu’une transaction peut être remplacée. Cette fonctionnalité est prévue pour Bitcoin Core 0.12.0 (attendue en janvier / février 2016), mais, comme la version originale de Nakamoto, cette fonctionnalité est ‘opt-in’, les personnes qui souhaitent pouvoir remplacer leurs transactions doivent donc utiliser un portefeuille qui prend en charge cette fonctionnalité.

Il n'existe actuellement aucun portefeuille qui propose cette fonctionnalité.  Mais les portefeuilles qui la proposeront à l'avenir seront en mesure de combiner plusieurs transactions ensemble afin de réduire l'espace de la blockchain utilisé ainsi que d'augmenter les frais payés sur les transactions dont la confirmation prend trop de temps, ceci afin d'éviter qu'elles ne restent «coincées» indéfiniment (un vrai problème d'utilisabilité).

## Weak blocks et IBLTs indiquent “2016” dans le calendrier de la feuille de route.  Cela signifie-t-il que vous ne savez pas quand ces fonctionnalités seront disponibles&nbsp;?  {#weak-blocks-iblts}

[Weak blocs et IBLTs][Weak blocks and IBLTs] sont deux technologies distinctes toujours [activement étudiées][actively studied] afin de déterminer les bons paramètres, mais le nombre de développeurs qui travaillent sur ces projets est limité et il est donc difficile de savoir quand ils seront déployés.

Weak blocs et IBLTs pourront aussi être déployés en tant que simples améliorations du réseau (pas de soft fork ni de hard fork requis) ce qui signifie qu'il n'y aura probablement qu'un court laps de temps entre le moment où les tests seront terminés et le moment où tous les nœuds mis à jour pourront bénéficier de ces nouvelles fonctionnalités.  Nous espérons que ce sera le cas au cours de l'année 2016.

Après le déploiement, weak blocs et IBLTs pourraient bénéficier d'un simple soft fork non controversé ([commande de transaction canonique][canonical transaction ordering]), qui devrait être facile à déployer en utilisant le système versionBits [BIP9][] décrit ailleurs dans cette FAQ.

[canonical transaction ordering]: https://gist.github.com/gavinandresen/e20c3b5a1d4b97f79ac2#canonical-ordering-of-transactions

## Pourquoi les mineurs adopteraient-ils le format segwit, sachant que cela ne procure aucune économie de bande passante, d'espace de stockage ou de temps de traitement ? {#why-mine-segwit}

La plupart des [soft forks précédents][previous soft forks] n'ont pas non plus fourni ces avantages aux mineurs. Par exemple,

| [BIP16][] (P2SH) | Nouveau type de transaction |
| [BIP30][] (txids dupliqués) | Contrôle requis des txids dupliqués |
| [BIP34][] (hauteur dans coinbase) | Réduction de l'espace de coinbase disponible pour les mineurs de 4 octets |
| [BIP65][] (OP_CLTV) | Nouvel opcode |

Le soft fork [BIP66][] (stricte DER) qui a été activé en Juillet 2015 va bientôt permettre de réduire le temps de traitement en permettant de passer à libsecp256k1 pour la validation comme décrit ailleurs dans cette FAQ.  Cette réduction du temps de validation rend ce soft fork plutôt original en ce sens qu'il offre des avantages directs pour les mineurs.

Segregated witness (segwit) fournit plusieurs avantages majeurs à toute personne qui l'utilise pour créer des transactions:

Une solution permanente pour la malléabilité par autrui, permettant à des contrats intelligents à plusieurs niveaux de prospérer.  Une réduction modeste des frais.  Des mises à niveau facilitées du langage de script de Bitcoin permettant aux portefeuilles d'accéder plus facilement à de nouvelles fonctionnalités.

Par le biais des précédents soft forks et grâce aux échanges qui ont eu lieu lors du [Panel de mineurs][Miners' Panel] à Scaling Bitcoin Hong Kong, les mineurs ont montré à maintes reprises qu'ils veulent que Bitcoin soit le système le plus utile possible, même si ils n'en reçoivent pas de bénéfices directs.  Segwit et les autres améliorations prévues dans la feuille de route fournissent des améliorations importantes à l'utilité de Bitcoin.

En outre, segwit permet aux mineurs de mettre plus de transactions dans leurs blocs, ce qui peut leur permettre d'accroître leur revenu par bloc.

## Comment puis-je aider&nbsp;?

Commencez par lire les pages [Contributeur Bitcoin Core][Bitcoin Core contributor]  sur Bitcoin.org.  En particulier la partie sur la [revue de code][code review] qui est un élément préalable essentiel au déploiement des soft forks.

Pour obtenir des conseils précis sur la meilleure façon dont vous pourriez aider, rejoignez-nous sur le canal IRC
[#bitcoin-dev][]

[#bitcoin-dev]: https://webchat.freenode.net/?channels=bitcoin-dev&amp;uio=d4
[actively studied]: http://diyhpl.us/wiki/transcripts/scalingbitcoin/bitcoin-block-propagation-iblt-rusty-russell/
[bip-segwit]: https://github.com/bitcoin/bips/blob/master/bip-0141.mediawiki
[BIP9]: https://github.com/bitcoin/bips/blob/master/bip-0009.mediawiki
[BIP16]: https://github.com/bitcoin/bips/blob/master/bip-0016.mediawiki
[BIP30]: https://github.com/bitcoin/bips/blob/master/bip-0030.mediawiki
[BIP34]: https://github.com/bitcoin/bips/blob/master/bip-0034.mediawiki
[BIP50]: https://github.com/bitcoin/bips/blob/master/bip-0050.mediawiki
[BIP65]: https://github.com/bitcoin/bips/blob/master/bip-0065.mediawiki
[BIP66]: https://github.com/bitcoin/bips/blob/master/bip-0066.mediawiki
[BIP68]: https://github.com/bitcoin/bips/blob/master/bip-0068.mediawiki
[BIP112]: https://github.com/bitcoin/bips/blob/master/bip-0112.mediawiki
[BIP113]: https://github.com/bitcoin/bips/blob/master/bip-0113.mediawiki
[bitcoin core contributor]: https://bitcoin.org/en/bitcoin-core/
[Bitcoin relay network]: http://bitcoinrelaynetwork.org/
[code review]: https://bitcoin.org/en/development#code-review
[estimated savings]: https://www.reddit.com/r/bitcoinxt/comments/3w1i6b/i_attended_scaling_bitcoin_hong_kong_these_are_my/cxtkaih
[increase in total bandwidth]: https://scalingbitcoin.org/hongkong2015/presentations/DAY1/3_block_propagation_1_rosenbaum.pdf
[libsecp256k1]: https://github.com/bitcoin/secp256k1
[libsecp256k1 verification]: https://github.com/bitcoin/bitcoin/pull/6954
[max_block_size]: https://github.com/bitcoin/bitcoin/blob/3038eb63e8a674b4818cb5d5e461f1ccf4b2932f/src/consensus/consensus.h#L10
[miners' panel]: https://youtu.be/H-ErmmDQRFs?t=1086
[payment channel efficiency]: https://scalingbitcoin.org/hongkong2015/presentations/DAY2/1_layer2_2_dryja.pdf
[previous soft forks]: https://github.com/bitcoin/bips/blob/master/bip-0123.mediawiki#classification-of-existing-bips
[weak blocks and iblts]: https://www.youtube.com/watch?v=ivgxcEOyWNs&t=1h40m20s
[q simple raise]: #size-bump