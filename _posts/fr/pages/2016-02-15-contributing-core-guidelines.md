---
title: Comment contribuer à Bitcoin Core ?
name: contributing-guidelines
id: fr-contributing-guidelines
permalink: /fr/faq/contributing-code/
layout: page
type: pages
lang: fr
share: false
category: faqs
version: 1
---
Le projet Bitcoin fonctionne sur un modèle de contribution ouvert où chacun est invité à participer à son développement en effectuant des revues de code, des tests et des correctifs.  Ce document présente la procédure à suivre et les directives concernant la contribution.

Tout d'abord, en termes structurels, le concept de "core developer" comme personne privilégiée n'existe pas.  L'open source repose le plus souvent sur la notion de méritocratie où les contributeurs les plus anciens acquièrent une plus grande confiance de la part de la communauté des développeurs.  Cependant, une certaine hiérarchie est nécessaire pour des raisons pratiques.  Il y a donc des "responsables" du référentiel qui sont chargés de fusionner les pull-requests ainsi qu'un "responsable principal (lead maintainer)" qui est chargé de gérer les livraisons, d'effectuer la modération et de désigner les responsables.

Processus de contribution
--------------------

Le code source est maintenu en respectant le "processus de contribution" où chacun, sans exception, contribue ses propositions de correctif en utilisant des "pull-requests".  Cela facilite la participation de tous, les tests et l'examen du code.

Pour contribuer des correctifs, voici le processus à suivre :

  - Créer un fork du référentiel
  - Créer une branche spécifique
  - Valider les correctifs

Il faut avoir préalablement accepté les conventions de codage du projet disponible sur [doc/developer-notes.md](https://github.com/bitcoin/bitcoin/blob/master/doc/developer-notes.md).

En général les commits se doivent d'être [atomique](https://en.wikipedia.org/wiki/Atomic_commit#Atomic_commit_convention) et les diffs doivent être faciles à lire.  Pour cette raison, il ne faut pas mélanger de correctifs de mise en forme ou de réorganisation de code avec des changements de code réels.

Les messages de commit doivent être assez complet et se composent par défaut d'une ligne d'intitulé courte (50 caractères maximum), d'une ligne vide suivie du texte explicatif composé d'un ou plusieurs paragraphes; à moins que le titre seul soit suffisamment explicite (comme "typo corrigée dans main.cpp"), dans ce cas une seule ligne de titre suffit.  Les messages de commit doivent permettre aux personnes qui liront votre code plus tard de pouvoir comprendre la logique sous-jacente à vos décisions.  De plus amples explications [ici](http://chris.beams.io/posts/git-commit/).

Si un commit particulier fait référence à une autre fiche d'anomalie (issue), veuillez ajouter la référence, par exemple "refs #1234" ou bien "fixes #4321".  L'utilisation des mots clés "fixes ou closes" entraînera la fermeture automatique de la fiche d'anomalie correspondante dès que la pull-request sera fusionnée.

Veuillez consulter le [manuel Git](https://git-scm.com/doc) pour plus d'information.

  - Pousser vos modifications sur votre fork
  - Créer la pull request

Le titre de la pull-request doit être précédé par le nom du composant ou de la partie du code concerné par la pull-request.  Exemples:

    Consensus: Add new opcode for BIP-XXXX OP_CHECKAWESOMESIG
    Net: Automatically create hidden service, listen on Tor
    Qt: Add feed bump button
    Trivial: fix typo

Si une pull request ne doit pas être considérée comme prête à être fusionnée, veuillez préfixer son titre par [WIP] (work in progress) ou bien utilisez les [listes de tâches](https://github.com/blog/1375-task-lists-in-gfm-issues-pulls-comments) dans la description pour indiquer quelles sont les tâches encore en attente.

The body of the pull request should contain enough description about what the patch does together with any justification/reasoning. You should include references to any discussions (for example other tickets or mailing list discussions).

At this stage one should expect comments and review from other contributors. You can add more commits to your pull request by committing them locally and pushing to your fork until you have satisfied all feedback. If your pull request is accepted for merging, you may be asked by a maintainer to squash and or rebase your commits before it will be merged. The length of time required for peer review is unpredictable and will vary from patch to patch.


La philosophie d'une pull-request
-----------------------

Un patch doit reste focalisé sur un seul sujet.  Par exemple une pull-request peut ajouter une nouvelle fonctionnalité, corriger un bug ou effectuer une refactorisation de code mais pas tout faire à la fois. Veuillez également éviter les pull-requests qui cherche à trop en faire, sont trop volumineuses ou trop complexes car cela rend la revue de code difficile.

### Fonctionnalités

When adding a new feature, thought must be given to the long term technical debt and maintenance that feature may require after inclusion. Before proposing a new feature that will require maintenance, please consider if you are willing to maintain it (including bug fixing). If features get orphaned with no maintainer in the future, they may be removed by the Repository Maintainer.

### Refactorisation

Refactoring is a necessary part of any software project's evolution. The following guidelines cover refactoring pull requests for the project.

There are three categories of refactoring, code only moves, code style fixes, code refactoring. In general refactoring pull requests should not mix these three kinds of activity in order to make refactoring pull requests easy to review and uncontroversial. In all cases, refactoring PRs must not change the behaviour of code within the pull request (bugs must be preserved as is).

Project maintainers aim for a quick turnaround on refactoring pull requests, so where possible keep them short, uncomplex and easy to verify. 


Le processus de "prise de décision"
-------------------------

The following applies to code changes to the Bitcoin Core project (and related projects such as libsecp256k1), and is not to be confused with overall Bitcoin Network Protocol consensus changes.

Whether a pull request is merged into Bitcoin Core rests with the project merge maintainers and ultimately the project lead. 

Maintainers will take into consideration if a patch is in line with the general principles of the project; meets the minimum standards for inclusion; and will judge the general consensus of contributors.

In general, all pull requests must:

  - have a clear use case, fix a demonstrable bug or serve the greater good of the project (for example refactoring for modularisation);
  - be well peer reviewed;
  - have unit tests and functional tests where appropriate;
  - follow code style guidelines;
  - not break the existing test suite;
  - where bugs are fixed, where possible, there should be unit tests demonstrating the bug and also proving the fix. This helps prevent regression.

Patches that change Bitcoin consensus rules are considerably more involved than normal because they affect the entire ecosystem and so must be preceded by extensive mailing list discussions and have a numbered BIP. While each case will be different, one should be prepared to expend more time and effort than for other kinds of patches because of increased peer review and consensus building requirements.

### Revue de code

Toute le monde peut participer aux revues de code qui se concrétisent par des commentaires dans les pull-requests.  Typiquement les examinateurs vérifient le code à la recherche d'erreurs évidentes, effectuent des tests sur l'ensemble des correctifs et donnent leur avis sur les mérites techniques du patch.  Les responsables de la maintenance du projet tiennent compte des commentaires des examinateurs pour déterminer s'il y a consensus ou pas pour fusionner une pull-request (ne pas oublier que les discussions peuvent avoir eu lieu sur github, mais aussi sur la liste de diffusion et/ou sur IRC).  Voici ci-dessous la terminologie utilisée dans les commentaires des pull-requests :

  - ACK signifie "J'ai testé le code et je suis d'accord pour qu'il soit fusionné";
  - NACK signifie "Je ne suis pas d'accord pour que ce code soit fusionné", une justification technique détaillée doit accompagner ce rejet.  les NACKs non accompagnés d'une justification pourront ne pas être pris en compte;
  - utACK signifie "Je n'ai pas testé le code, mais je l'ai passé en revue et il m'a l'air OK, je suis d'accord pour qu'il soit fusionné";
  - Concept ACK signifie "Je suis d'accord sur le principe général de cette pull-request";
  - Nit fait référence à des anomalies triviales souvent non-blocantes.

Les examinateurs doivent inclure le hash du commit qu'ils ont passé en revue dans leurs commentaires.

Les responsables de la maintenance du projet se réservent le droit d'évaluer les opinions des examinateurs en faisant preuve de bon sens et en se basant sur la méritocratie : ceux qui ont fait preuve d'un engagement plus fort et d'une meilleure compréhension du projet (au fil du temps) ou qui ont une expertise évidente dans le domaine peuvent naturellement avoir plus de poids, comme on pourrait s'y attendre dans tous les aspects de la vie.

Lorsqu'un patch affecte du code critique relatif au consensus, la barre sera mise beaucoup plus haute en termes d'exigences de discussion et de revue de code, il faut garder à l'esprit que la moindre erreur peut coûter très cher à l'ensemble de la communauté.  Cela inclus la refactorisation de code critique relatif au consensus.

Lorsqu'un patch propose de modifier les règles du consensus Bitcoin, celui-ci doit avoir été longuement débattu sur la liste de diffusion et IRC, être accompagné d'un BIP discuté largement et avoir obtenu un consensus technique globalement perçu comme pouvant être bénéfique sur la base de l'avis des responsables de la maintenance.

Politique de livraison
--------------

Le chef de projet (project leader) est le responsable de la publication de chaque version de Bitcoin Core.