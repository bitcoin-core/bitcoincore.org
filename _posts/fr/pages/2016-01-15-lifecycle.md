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

We will provide maintenance "minor releases" that fix bugs within the major releases. We may add minor features to point releases where necessary, for example when back-porting consensus rule changes such as soft forks. Minor releases will be numbered 0.11.1, 0.11.2, 0.12.1, 0.12.2 etc.

We maintain the major versions until their "Maintenance End". We generally maintain the current and previous major release. So if the current release is 0.12, then 0.11 is also considered maintained. Once 0.13 is released, then 0.11 would be considered at it's "Maintenance End". The older the major release, the more critical issues have to be to get backported to it, and the more to warrant a new minor release. Once software has reached the "Maintenance End" period it will only receive critical security fixes until the EOL date. After EOL, users must upgrade to a later version to receive security updates.

Veuillez noter que les versions mineures embarquent des corrections de bugs, des mises à jour de traduction et des soft forks.  La traduction sur [Transifex][bitcoin-transifex-link] est ouverte uniquement aux deux dernières versions majeures.

Par exemple, la version majeure 0.9 a été publié le 2014-03-19  et nous avons fourni des versions de maintenance jusqu'au 2015-02-16.  Les problèmes de sécurité critiques continueront d'être corrigé jusqu'à la date de fin de vie "EOL" fixée au 2016-02-31.  Cependant, pour profiter des dernières corrections de bug, vous aurez à effectuer une mise à jour vers une version majeure plus récente.

Une fois que la fin de vie est atteinte, vous devrez effectuer une mise à jour vers une version plus récente.

| Version | Date de sortie | Fin de maintenance | Fin de vie |
|---------|--------------|-----------------|-------------|
{% include posts/_maintenance-table.md %}

[bitcoin-transifex-link]: https://www.transifex.com/bitcoin/bitcoin/