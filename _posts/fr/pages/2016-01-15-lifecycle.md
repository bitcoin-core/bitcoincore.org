---
title: Cycle de vie du logiciel
name: software-life-cycle
id: fr-software-life-cycle
permalink: /fr/lifecycle/
layout: page
type: pages
lang: fr
share: false
version: 1
---
Ce document décrit le cycle de vie du logiciel Bitcoin Core publié par le projet Bitcoin Core.  Il est conforme aux politiques de maintenance standards de la plupart des logiciels commerciaux.  

Notre objectif est de publier une version majeure tous les 6-7 mois, celles-ci seront numérotées 0.11, 0.12 etc.

Nous produirons des "versions mineures» qui corrigerons des bugs détectés dans les versions majeures.  Si nécessaire, nous pourrons y ajouter ponctuellement quelques fonctionnalités mineures, comme par exemple les changements des règles de consensus lors d'un rétro-portage suite à un soft fork.   Les versions mineures seront numérotées 0.11.1, 0.11.2, 0.12.1, 0.12.2 etc.

Nous maintenons les versions majeures jusqu'à leur "fin de maintenance".  Nous maintenons généralement la version majeure actuelle et la précédente.  Donc, si la version actuelle est la 0.12, la 0.11 est également considérée comme maintenue.  Dès que la 0.13 sera publiée, la 0.11 sera considérée comme étant en "fin de maintenance".  Plus la version majeure est ancienne, plus il y a de correctifs critiques à rétro-porter, et plus cela peut justifier une nouvelle version mineure.  Une fois que le logiciel a atteint la période de "fin de maintenance", il ne recevra plus que les correctifs de sécurité critiques jusqu'à la date de fin de vie.  Après la "fin de vie", les utilisateurs doivent mettre effectuer une niveau vers une version ultérieure pour recevoir les mises à jour de sécurité.

Veuillez noter que les versions mineures embarquent des corrections de bugs, des mises à jour de traduction et des soft forks.  La traduction sur [Transifex][bitcoin-transifex-link] est ouverte uniquement aux deux dernières versions majeures.

Par exemple, la version majeure 0.9 a été publié le 2014-03-19  et nous avons fourni des versions de maintenance jusqu'au 2015-02-16.  Les problèmes de sécurité critiques continueront d'être corrigé jusqu'à la date de fin de vie "EOL" fixée au 2016-02-31.  Cependant, pour profiter des dernières corrections de bug, vous aurez à effectuer une mise à jour vers une version majeure plus récente.

Une fois que la fin de vie est atteinte, vous devrez effectuer une mise à jour vers une version plus récente.

| Version | Date de sortie | Fin de maintenance | Fin de vie |
|---------|--------------|-----------------|-------------|
{% include posts/_maintenance-table.md %}

[bitcoin-transifex-link]: https://www.transifex.com/bitcoin/bitcoin/