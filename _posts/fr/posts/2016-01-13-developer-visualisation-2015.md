---
type: posts
layout: post
lang: fr
id: fr-development-visualisation-2015
name: development-visualisation-2015
title: Visualisation du développement de Core pour 2015
permalink: /fr/2016/01/13/development-visualisation-2015/
version: 1
excerpt: La vidéo suivante présente l'activité de 'commit' dans le référentiel Bitcoin Core au cours de l'année 2015.
---
La vidéo suivante présente l'activité de 'commit' dans le [référentiel Bitcoin Core][repository] pour l'année 2015.  Une liste complète des contributeurs de code au cours de cette période est disponible [ici][activity].

<iframe width="560" height="315" src="https://www.youtube.com/embed/FIt7GLxxIpY" frameborder="0" allowfullscreen> </iframe>

En 2015, le projet Bitcoin Core a sorti 2 versions majeures de son logiciel ainsi que 5 versions de maintenance supplémentaires. 
En outre, deux mises à jour de type 'soft forks' ont été déployées et activées avec succès.  La première, [BIP66], a comblé une faille de sécurité sérieuse introduite par openssl; et la seconde, [BIP65], a ajouté un nouvel opcode CHECKLOCKTIMEVERIFY au langage de script de Bitcoin.

Le projet a aussi terminé le gros du travail de la prochaine version majeure, [0.12], dont la livraison est prévue pour le mois février.  La version [0.12] intégrera la bibliothèque [libsecp256k1] qui est en développement depuis [deux ans et demi][secp_contributors], et qui multiplie les performances de validation de signature par 7, élément indispensable à l'amélioration de la scalabilité à venir. 

Sachez que l'activité de 'commit' ne représente qu'une petite partie de l'activité globale d'un développeur, elle ne prend pas en compte l'activité des pair examinateurs, des évaluateurs de code, des testeurs/intégrateurs et des traducteurs.  Elle ne reflète pas non plus avec précision le temps passé en recherche, en discussion et en développement avant que le code ne soit versé dans le référentiel.

Nous souhaitons profiter de cette opportunité pour remercier tous ceux qui ont collaboré jusqu'à aujourd'hui au projet Bitcoin Core et nous ont ainsi permis de rendre Bitcoin encore meilleur au bénéfice de tous.

[repository]: https://github.com/bitcoin/bitcoin
[activity]: https://github.com/bitcoin/bitcoin/graphs/contributors?from=2015-01-01&to=2016-01-01&type=c
[BIP65]: https://github.com/bitcoin/bips/blob/master/bip-0065.mediawiki
[BIP66]: https://github.com/bitcoin/bips/blob/master/bip-0066.mediawiki
[0.12]: https://github.com/bitcoin/bitcoin/blob/0.12/doc/release-notes.md
[libsecp256k1]: https://github.com/bitcoin/secp256k1
[secp_contributors]: https://github.com/bitcoin/secp256k1/graphs/contributors?from=2013-03-04&to=2015-12-01&type=c