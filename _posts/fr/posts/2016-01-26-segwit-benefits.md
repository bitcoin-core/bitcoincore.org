---
type: posts
layout: post
lang: fr
name: segwit-benefits
id: fr-segwit-benefits
title: Avantages de "Segregated Witness"
permalink: /fr/2016/01/26/segwit-benefits/
version: 1
excerpt: Cette page résume certains des avantages apportés par segregated witness.
---
{% include _toc.html %}

Le soft-fork Segregated Witness (segwit) comprend un large éventail de fonctionnalités, dont beaucoup sont d'un niveau technique élevé.  Cette page récapitule certains des avantages apportés par ces fonctionnalités.

## Correctifs de la malléabilité

Les transactions Bitcoin sont identifiés par un hash hexadécimal de 64 chiffres appelé identifiant de transaction (txid) lequel est basé sur les fonds dépensés d'une part et sur qui sera capable de dépenser les résultats de la transaction d'autre part.

Malheureusement, la façon dont le txid est calculé permet à quiconque de faire de petites modifications à la transaction qui ne changeront pas sa signification, mais qui vont changer le txid.  Ceci est appelé la malléabilité par un tiers.  BIP 62 ("traiter la malléabilité") a tenté de répondre à ces questions d'une manière fragmentaire, mais était trop compliqué à mettre en œuvre à l'aide de contrôles de consensus et a été retirée.

Par exemple, vous pourriez émettre une transaction avec le txid ef74...C309 sur le réseau, mais constater qu'un tiers, comme un nœud sur le réseau relayant votre transaction, ou le mineur qui inclut votre transaction dans un bloc, modifie légèrement la transaction, résultant en une transaction qui dépense les mêmes fonds et paye les mêmes adresses, mais est confirmée sous un txid totalement différent 683f...8bfa.

Plus généralement, si un ou plusieurs des signataires de la transaction modifient leurs signatures alors la transaction reste valable et paie les mêmes montants aux mêmes adresses, mais le txid change complètement car il intègre les signatures.  Le cas général où des modifications sont apportées aux données de signature (mais pas des sorties ou une sélection d'entrées) et qui modifient donc la transaction est appelé la malléabilité scriptSig.

Segwit empêche la malléabilité par un tiers et la malléabilité scriptSig en permettant aux utilisateurs Bitcoin de déplacer les parties malléables de la transaction dans le *témoin de la transaction* (transaction witness), et de séparer ce témoin afin que les changements apportés au témoin n'affectent pas le calcul du txid.

### Qui en profite ?

- **Les éditeurs de portefeuille souhaitant surveiller les bitcoins dépensés :** il est plus facile de contrôler l'état de vos propres transactions sortantes en surveillant uniquement les txids.  Mais dans un système où la malléabilité par un tiers est possible, les portefeuilles doivent mettre en œuvre du code supplémentaire pour être en mesure de détecter les txids qui ont été modifiés.

- **Toute personne dépensant des transactions non confirmées :** si Alice paie Bob dans la transaction 1, Bob utilise ce paiement pour payer Charlie dans la transaction 2, puis le paiement d'Alice est modifié par malléabilité et se retrouve confirmé avec un txid différent, alors la transaction 2 est maintenant invalide et Charlie n'a pas été payé.  Si Bob est digne de confiance, il enverra de nouveau le paiement à Charlie; mais s'il ne l'est pas, il peut tout simplement garder les bitcoins pour lui-même.

- **Le Réseau Lightning :** avec le solutionnement de la malléabilité par un tiers et de la malléabilité scriptSig, le Réseau Lightning est plus simple à mettre en œuvre et significativement plus efficace dans son utilisation de l'espace sur la blockchain.  Avec la disparition de la malléabilité scriptSig, il devient également possible d'exécuter des clients légers Lightning qui externalisent la surveillance de la blockchain et permettent d'éviter que chaque client Lightning soit aussi un full-node Bitcoin.

- **Toute personne utilisant la blockchain :** contrats intelligents d'aujourd'hui, tels que les canaux de micropaiement et futurs contrats intelligents deviennent moins compliqués à concevoir, à comprendre et à surveiller.

Note: les transactions segwit n'échappent au risque de malléabilité que si toutes leurs entrées sont des dépenses segwit (soit directement, soit via une adresse P2SH segwit rétrocompatible).

### Informations complémentaires

 * [La malléabilité sur le wiki Bitcoin](https://en.bitcoin.it/wiki/Transaction_Malleability)
 * [Article du Coin Telegraph sur l'attaque 'Malléabilité' de 2015](http://cointelegraph.com/news/115374/the-ongoing-bitcoin-malleability-attack)
 * [Article de Bitcoin Magazine sur l'attaque 'Malléabilité' de 2015](https://bitcoinmagazine.com/articles/the-who-what-why-and-how-of-the-ongoing-transaction-malleability-attack-1444253640)
 * [Transcription de "Overview of BIPs necessary for Lightning"](http://diyhpl.us/wiki/transcripts/scalingbitcoin/hong-kong/overview-of-bips-necessary-for-lightning/)
 * [BIP 62](https://github.com/bitcoin/bips/blob/master/bip-0062.mediawiki)
 * [BIP 140 -- approche alternative aux correctifs de malléabilité](https://github.com/bitcoin/bips/blob/master/bip-0140.mediawiki)
 * [Réponse Stack exchange concernant la transaction 683f...8bfa](http://bitcoin.stackexchange.com/questions/22051/transaction-malleability-in-the-blockchain/22058#22058)

## Scaling linéaire des opérations de sighash

Les approches basiques d'augmentation de la taille des blocs de Bitcoin font face à un problème majeur, en effet pour certaines transactions, la procédure de hachage des signatures croit de manière quadratique plutôt que linéaire.

![Linear versus quadratic](/assets/images/linear-quad-scale.png)

En substance, doubler la taille d'une transaction augmente à la fois le nombre des opérations de signature, mais aussi la quantité de données qui doit être hachée pour la vérification de chacune de ces signatures.  Cela a déjà pu être constaté, un seul bloc peut nécessiter jusqu'à 25 secondes pour être validé, et des transactions malicieusement conçues pourraient prendre plus de 3 minutes.

Segwit résout ce problème en changeant la manière de calculer le hash de la transaction pour les signatures, afin que chaque octet d'une transaction ne soit haché que deux fois au plus. Cela permet d'offrir la même fonctionnalité plus efficacement, de sorte que même de grosses transactions peuvent être générés sans rencontrer de problème du au hachage des signatures, même si elles sont générées par malveillance ou que des blocs beaucoup plus gros (et donc des transactions plus importantes) sont pris en charge.

### Qui en profite ?

La suppression de la croissance quadratique des données hachées pour la vérification des signatures permet d'augmenter la taille des blocs de manière plus sûre.  Faire cela sans limiter pour autant la taille des transactions permet à Bitcoin de continuer à gérer des paiements à destination ou en provenance de groupes importants, tels que les paiements des récompenses minières ou des services de crowdfunding.

Ce hash modifié ne s'applique qu'aux opérations de signature initiées depuis les données witness, donc les opérations de signature concernant les données du bloc de base continueront d'exiger des limites inférieures.

### Informations complémentaires

 * [BIP 143](https://github.com/bitcoin/bips/blob/master/bip-0143.mediawiki)
 * [Blog de Rusty Russell sur la transaction de 25s](http://rusty.ozlabs.org/?p=522)
 * [CVE 2013-2292 sur le wiki Bitcoin](https://en.bitcoin.it/wiki/Common_Vulnerabilities_and_Exposures#CVE-2013-2292)
 * [Proposition pour limiter les transactions à 100ko](https://lists.linuxfoundation.org/pipermail/bitcoin-dev/2015-July/009494.html)
 * [Commit Bitcoin Classic sur la branche 0.11.2 ajoutant une nouvelle limite de consensus sur les octets sighash](https://github.com/bitcoinclassic/bitcoinclassic/commit/842dc24b23ad9551c67672660c4cba882c4c840a)

## Signature des montants d'input

Quand un portefeuille matériel signe une transaction, il peut facilement vérifier le montant total dépensé, mais il ne peut déterminer les frais de manière sûre qu'en ayant une copie complète de toutes les transactions d'entrée dépensées, et doit hacher chacune de celles-ci pour s'assurer qu'on ne lui a pas fourni de fausses données.  Chaque transaction pouvant faire individuellement jusqu'à 1 Mo en taille, cette opération peut s'avérer coûteuse, même si la transaction que l'on signe est elle-même plutôt petite.

Segwit résout ce problème en hachant explicitement le montant de l'entrée (input).  Cela signifie qu'un portefeuille matériel peut simplement recevoir le hash de la transaction, l'index, et le montant (et connaître la clé publique utilisée), et peut signer en toute sécurité la transaction qui dépense, peu importe la taille ou la complexité de la transaction dépensé.

### Qui en profite ?

Les fabricants et les utilisateurs de portefeuille matériel seront les premiers à en profiter; cependant il est fort probable que cela rendra beaucoup plus simple une utilisation sûre de Bitcoin dans tous les petits équipements embarqués pour des usages dans l'"Internet des objets". 

Cet avantage est disponible uniquement lorsque les transactions qui dépensent envoient les fonds à des adresses segwit (ou à des adresses segwit-via-P2SH).

### Informations complémentaires

 * [BIP 143](https://github.com/bitcoin/bips/blob/master/bip-0143.mediawiki)

## Amélioration de la sécurité pour le multisig via pay-to-script-hash (P2SH)

Les paiements multisig utilisent actuellement P2SH qui est sécurisé par l'algorithme HASH160 sur 160 bits (RIPEMD de SHA256). Toutefois, si un des signataires souhaite voler tous les fonds, il suffit de trouver une collision entre une adresse valide dans le cadre d'un script multisig et un script qui envoie simplement tous les fonds à lui-même avec seulement 80 bits (2<sup>80</sup>) de travail, ce qui est déjà dans le domaine du possible pour un attaquant extrêmement bien doté en ressources.  (A titre de comparaison, à une puissance soutenue de 1 exahash/seconde, le réseau minier Bitcoin produit 80 bits de travail toutes les deux semaines)

Segwit résout ce problème en utilisant l'algorithme HASH160 uniquement pour les paiements directs à une clé publique unique (où ce genre d'attaque est inutile), tout en utilisant des hashs SHA256 de 256 bits pour les paiements à un hash de script.

### Qui en profite ?

Toute personne qui effectue un paiement à une adresse multisig ou à un contrat intelligent (smart contract) via segwit profite du niveau de sécurité additionnel dont bénéficient les scripts.

### Informations complémentaires

 * [Gavin Andresen se demande si les attaques 80-bit valent la peine qu'on s'en inquiète](https://lists.linuxfoundation.org/pipermail/bitcoin-dev/2016-January/012198.html)
 * [Ethan Heilman décrit un algorithme pour détecter un cycle](https://lists.linuxfoundation.org/pipermail/bitcoin-dev/2016-January/012202.html)
 * [Rusty Russell calcule les coûts d'une attaque](https://lists.linuxfoundation.org/pipermail/bitcoin-dev/2016-January/012227.html)
 * [Anthony Towns applique l'algorithme de détection de cycle pour exploiter les transactions](https://lists.linuxfoundation.org/pipermail/bitcoin-dev/2016-January/012218.html)
 * [Gavin Andresen résume la discussion](https://lists.linuxfoundation.org/pipermail/bitcoin-dev/2016-January/012234.html)

## Versionnage des scripts

Il est possible de modifier le langage de script de Bitcoin pour permettre à la fois d'améliorer la sécurité et d'ajouter de nouvelles fonctionnalités. Cependant, tel qu'il a été conçu, le langage de script de Bitcoin n'autorise les modifications que si elles sont rétrocompatibles (soft-forking) et qu'elles sont mises en œuvre en remplaçant l'un des dix OP\_NOP opcodes encore disponibles avec un nouvel opcode qui peut de manière conditionnelle faire échouer le script, mais qui ne peut rien faire d'autre.  Cette technique est suffisante pour la plupart des modifications du langage de script -- telles que l'introduction d'une nouvelle méthode de signature ou d'une fonctionnalité comme OP\_CLTV, mais elle peut parfois ressembler à du bidouillage (par exemple, OP\_CLTV doit être accompagné d'un OP\_DROP) et ne peut pas être utilisée pour activer de nouvelles fonctionnalités même aussi simples que l'assemblage de deux chaînes.

Segwit résout ce problème en ajoutant aux scripts un numéro de version, de sorte qu'il est possible de créer de nouveaux opcodes en augmentant simplement le numéro de version alors qu'il aurait fallu auparavant effectuer un hard-fork pour en bénéficier dans les transactions non-segwit.

### Qui en profite ?

Le fait de pouvoir modifier simplement les opcodes du langage de script va faciliter la création de scripts encore plus avancés dans Bitcoin, comme par exemple l'ajout des signatures Schnorr, l'utilisation de 'key recovery' pour réduire la taille des signatures, le support des sidechains, ou bien encore la création de contrats encore plus intelligents en utilisant des Merklized Abstract Syntax Trees (MAST) et bien d'autres idées qui sont encore au stade de la recherche.

## Réduire la croissance de la base des UTXO

La base de données des sorties de transaction non dépensées (UTXO) est maintenu par chaque nœud Bitcoin validant les transactions, ceci afin de déterminer si les transactions qui arrivent sont valides ou frauduleuses.  Pour permettre un fonctionnement performant du réseau, cette base de données doit être très rapide à interroger et à modifier, et devrait idéalement pouvoir être entièrement stockée en mémoire vive (RAM), donc il est important de maintenir aussi petite que possible la taille en octets de cette base de données.

Plus Bitcoin grandit plus la tâche devient difficile, chaque nouvel utilisateur doit en effet avoir au moins une entrée UTXO propre et préférera avoir des entrées multiples pour améliorer la confidentialité et la flexibilité, ou pour qu'elles soient disponibles en tant que support pour des canaux de paiement ou des contrats intelligents.

Segwit améliore la situation en rendant les données de signature, qui n'ont pas d'impact sur la taille de la base des UTXO, 75% moins coûteuse que les données qui on un impact sur la taille de la base des UTXO.  Cette réduction de coût devrait encourager les utilisateurs à privilégier l'utilisation de transactions segwit qui minimisent l'impact sur la base des UTXO car ils réduiront ainsi leurs frais de transaction, et cela devrait aussi encourager les développeurs à concevoir des contrats intelligents et des nouvelles fonctionnalités qui minimisent l'impact sur la base des UTXO.

Segwit étant déployé par soft-fork qui ne modifie pas la taille de base des blocks, le taux de croissance de la base des UTXO dans le cas le plus défavorable reste identique.

### Qui en profite ?

Une réduction de la croissance de la base des UTXO bénéficiera aux mineurs, aux entreprises et aux utilisateurs qui font tourner des full-nodes, ce qui contribuera à préserver le niveau de sécurité actuel du réseau Bitcoin alors que de plus en plus d'utilisateurs rejoignent le système. Les utilisateurs et les développeurs qui aident à minimiser la croissance de la base des UTXO profiteront de frais de transaction moins élevés par rapport à ceux qui ignorent l'impact de leurs transactions sur la croissance de la base des UTXO.

### Informations complémentaires

 * [Tableau de bord des UTXO de Statoshi](http://statoshi.info/dashboard/db/unspent-transaction-output-set)

## Preuves de fraude compactes (compact fraud proofs)

Avec l'augmentation de la base des utilisateurs de Bitcoin, le processus de validation de l'ensemble de la blockchain devient naturellement plus coûteux. Pour conserver la nature décentralisée et sans tiers de confiance de Bitcoin, il est important de permettre à ceux qui ne sont pas en mesure de valider l'intégralité de la blockchain d'être au moins capable d'en valider autant que possible pour un coût modique.

Grâce à segwit un futur soft-fork pourra étendre la structure de witness pour y inclure des données de commitment, ce qui permettra à des clients légers (SPV) de vérifier le respect des règles de consensus telles que le nombre de bitcoins créés dans un bloc, la taille d'un bloc, et le nombre de sigops utilisées dans un bloc.

### Qui en profite ?

Les preuves de fraude (fraud proofs) permettent aux utilisateurs de clients SPV d'aider à faire respecter les règles de consensus de Bitcoin, ce qui aura potentiellement pour effet d'augmenter considérablement la sécurité du réseau Bitcoin dans son ensemble, mais aussi de réduire les vecteurs d'attaques contre les utilisateurs individuels.

Ces preuves de fraude (faud proofs) peuvent être ajoutées à la structure de données witness dans le cadre d'un futur soft-fork, elles vont permettre aux clients SPV de faire respecter les règles de consensus, même pour les transactions qui ne font pas usage des fonctionnalités segwit.

## Gains de performance en l'absence de vérification des signatures

Les signatures des anciennes transactions ont probablement moins d'interêt que les signatures des transactions à venir -- par exemple, Bitcoin Core ne vérifie pas, par défaut, les signatures des transactions antérieures au dernier point de contrôle (checkpoint), et certains clients SPV ne vérifient pas les signatures du tout, faisant confiance aux mineurs ou au autres nœuds pour l'avoir déjà fait.  Cependant, les données de signature font actuellement partie intégrante de la transaction et doivent être présentes lors du calcul de la valeur de hachage de la transaction.

Séparer les données de signature du reste du bloc permet aux nœuds qui ne sont pas intéressés par les données de signature de les supprimer du disque voire de ne pas les télécharger du tout, et faire ainsi des économies de ressource.

### Qui en profite ?

Au plus les transactions utiliseront des adresses segwit, au plus les personnes qui font tourner des noeuds SPV ou 'pruned' pourront fonctionner en consommant moins de bande passante et d'espace disque.

## Augmentation de la capacité/taille des blocs

Comme les anciens nœuds ne téléchargeront que le bloc de base sans les "témoins", ils appliqueront uniquement la règle limitant la taille maximum des blocs à 1 Mo sur ces données.
Les nouveaux nœuds sachant interpréter le bloc complet incluant  les données "témoin" peuvent donc remplacer cette limite par une nouvelle permettant d'obtenir des blocs de plus grande taille.  Segregated witness tire profit de cette possibilité pour augmenter la taille limite des blocs à presque 4 Mo, et ajoute une nouvelle limite de coût pour s'assurer que le traitement des blocs reste raisonnable dans l'utilisation des ressources (cela se traduit au final par une limite réelle plus proche des 1,6 à 2 Mo).

### Qui en profite ?

Les personnes faisant tourner des portefeuilles mis à jour pourront profiter d'une augmentation de la taille des blocs en déplaçant les données de signature dans la partie "témoin" de la transaction.

## Vers une seule limite unifée de la taille des blocs

Actuellement, le consensus impose deux règles qui concernent la taille des blocs : la taille du bloc ne peut pas être supérieure à 1 Mo et il ne peut y avoir plus de 20.000 vérifications de signatures effectuées pour l'ensemble des transactions du bloc.

Trouver la sélection de transactions la plus rentable à inclure dans un bloc avec une seule limite est un cas typique du [problème du sac à dos](https://fr.wikipedia.org/wiki/Probl%C3%A8me_du_sac_%C3%A0_dos), qui peut être facilement et presque parfaitement résolu avec un simple [algorithme glouton](https://fr.wikipedia.org/wiki/Algorithme_glouton).  Cependant, l'ajout d'une seconde contrainte rend la recherche d'une solution beaucoup plus difficile dans certains cas, et ce problème théorique a été exploitée dans la pratique pour forcer certains blocs à être minés avec une taille bien inférieure à la capacité maximale.

Il est impossible de résoudre ce problème sans un hardfork ou une diminution significative de la taille des blocs.  Segwit ne pouvant pas non plus résoudre ce problème, il se contente de ne pas l'aggraver : en particulier, plutôt que d'introduire une nouvelle limite spécifique pour les données witness, segwit applique une seule limite qui combine le poids des données UTXO et des données witness, permettant aux deux d'être limités simultanément comme une seule entité combinée.

### Qui en profite ?

En fin de compte, ce sont les mineurs qui en profiteront si à l'avenir un hardfork définit la taille limite des blocs comme une simple somme pondérée de paramètres. Par exemple:

    50*sigops + 4*basedata + 1*witnessdata < 10M

Ceci permettra aux mineurs de remplir les blocs facilement et avec précision tout en maximisant leurs revenus liés aux frais de transaction, et cela profitera aussi aux utilisateurs car ils pourront calculer de manière fiable le montant des frais à payer pour que leur transaction soit minée.

### Informations complémentaires

 * [Problème du sac à dos](https://en.wikipedia.org/wiki/Knapsack_problem)
 * [Discussion sur les attaques Sigop sur bitcointalk en août 2015](https://bitcointalk.org/index.php?topic=1166928.0;all)
 * [Gregory Maxwell sur bitcoin-dev au sujet des limites witness](https://lists.linuxfoundation.org/pipermail/bitcoin-dev/2015-December/011870.html)
 * [Transcription "Métriques du coût de validation"](http://diyhpl.us/wiki/transcripts/scalingbitcoin/hong-kong/validation-cost-metric/)
