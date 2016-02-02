---
type: posts
layout: post
lang: fr
id: fr-bitcoin-core-2016-01-07-statement
name: bitcoin-core-2016-01-07-statement
title: Communiqué de Bitcoin Core -- 2016-01-07
permalink: /fr/2016/01/07/statement/
version: 1
---
Bitcoin est une “monnaie électronique peer-to-peer qui permet d’effectuer des paiements en ligne directs sans avoir à passer par l’intermédiaire d’un organisme financier”. Notre vision pour Bitcoin est d’étendre la flexibilité du système afin qu’il puisse fonctionner efficacement à très grande échelle, tout en maintenant la sécurité et les propriétés fondamentales de la décentralisation qui rendent Bitcoin unique.

Nous croyons que Bitcoin peut accomplir ceci en servant de base pour des couches additionnelles au dessus du protocole et des interfaces avec d’autres systèmes. En outre, nos objectifs à long terme comprennent la protection et l’amélioration de la vie privée des utilisateurs de Bitcoin.

“Bitcoin Core» fait référence à un projet open source descendant direct de l’implémentation originale de Bitcoin. Comme contributeurs au projet, nous maintenons et publions le logiciel pour la communauté Bitcoin. Nous nous efforçons d’apporter des améliorations au protocole du consensus en proposant des améliorations dont nous pensons qu’elles font sens au niveau technique selon la compréhension que nous avons des objectifs de Bitcoin, et dont nous croyons qu’elles auront une chance raisonnable de soutien et d’adoption généralisés.

Les modifications apportées aux règles du consensus Bitcoin peuvent être activéees soit par soft fork soit par hard fork (voir l’annexe A). les soft forks permettent des modifications compatibles. Avec les soft forks, ancienne et nouvelle versions du logiciel peuvent co-exister sur le réseau. Les soft forks peuvent introduire de nouvelles fonctionnalités sans perturber le système : les utilisateurs souhaitant utiliser les nouvelles fonctionnalités peuvent effectuer la mise à jour du logiciel tandis que ceux qui ne le souhaitent pas peuvent continuer d’utiliser leur ancienne version normalement.

Les hard forks créent un incompatibilité avec toutes les versions précédentes du logiciel Bitcoin et obligent chaque participant à mettre à niveau le logiciel afin d’utiliser les mêmes règles et ce dans un certain délai au risque de perdre de l’argent. De tels événements peuvent également nuire à l’effet de réseau en excluant les participants qui n’auraient pas pris les mesures nécessaires, et sont susceptibles de perturber les applications en aval.

Pour ces raisons, Bitcoin Core préconise fortement la compatibilité et estime que chaque utilisateur devrait pouvoir être libre de ne pas mettre à jour les règles de leur logiciel Bitcoin actuel. Il se trouve qu’il est possible d’ajouter quasiment tout type de nouvelle fonctionnalité avec un soft fork. Occasionnellement, les hard forks peuvent avoir certains avantages, et si il y a un accord quasi-universel, ces avantages peuvent l’emporter sur les inconvénients. En dehors de ces rares cas, les soft forks sont préférables. Nous croyons que cela est dans le meilleur intérêt des utilisateurs actuels et futurs du système.

Nous prévoyons également que, l’écosystème Bitcoin grandissant, le nombre d’implémentations alternatives du protocole Bitcoin puisse aussi augmenter, et il est inévitable que d’autres projets de logiciel puissent voir le jour avec des propositions radicalement différentes que l’écosystème Bitcoin devra considérer. En fin de compte, l’équipe de développement de Bitcoin Core ne décide pas des règles du consensus Bitcoin. Au lieu de cela, les utilisateurs participent au système Bitcoin en choisissant quel logiciel Bitcoin ils souhaitent exécuter. C’est pourquoi le logiciel Bitcoin Core n’est délibérément pas doté d’une fonction de mise à jour automatique. Son absence permet de s’assurer de la participation volontaire de l’utilisateur à chaque mise à jour, afin que les utilisateurs conservent toujours le choix du logiciel qu’ils exécutent.

### Annexe A

Un hard fork est une modification des règles de consensus qui a pour conséquence de pouvoir rendre valides selon les nouvelles règles des blocs qui auraient été invalides selon les anciennes règles.

Un soft fork est une modification des règles de consensus qui a pour conséquence de pouvoir rendre invalides selon les nouvelles règles des blocs qui auraient été valides selon les anciennes règles, mais qui garantit que tous les blocs qui auraient été invalides selon les anciennes règles restent invalides avec les nouvelles règles.
