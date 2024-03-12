---
title: Software-Lebenszyklus
name: software-life-cycle
id: en-software-life-cycle
permalink: /de/lifecycle/
layout: page
type: pages
lang: de
version: 2
---
{% include toc.html %}

Dieses Dokument beschreibt den Lebenszyklus des Bitcoin Core Softwarepakets, 
das vom Projekt veröffentlicht wurde. Es entspricht der Standardwartungsrichtlinie 
für kommerzielle Software.

## Versionierung

Bitcoin Core Releases sind wie folgt versioniert: 0.MAJOR.MINOR, und 
Release-Kandidaten haben das Suffix rc1, rc2 usw.

## Hauptversionen

Unser Ziel ist es, alle 6-7 Monate eine Hauptversion herauszubringen.

Diese werden mit 0.11.0, 0.12.0 usw. nummeriert.

## Wartungsversionen

Wir werden "Nebenversionen" für die Wartung bereitstellen, die Fehler in 
den Hauptversionen beheben. Als allgemeine Regel führen wir keine größeren 
neuen Funktionen in einer Wartungsversion ein (mit Ausnahme von Konsensregeln). 
Wir werdeb jedoch bei Bedarf kleinere Funktionen hinzufügen, und Änderungen 
der Konsensregeln wie Soft Forks.

Nebenversionen werden mit 0.11.1, 0.11.2, 0.12.1, 0.12.2 usw. nummeriert.

## Konsensregeln

Vorschläge zur Änderung von Konsensregeln werden immer zuerst in Wartungsversionen 
wie 0.11.2, 0.12.1 usw. ausgeliefert. Dies erleichtert es Unternehmen, den 
Vorschlag zu bewerten und zu testen, da er im Vergleich zu einer Hauptversion 
weniger Änderungen enthält. Es ermöglicht auch Benutzern, die einem konservativeren 
Upgrade-Pfad folgen, Änderungen der Konsensregeln zeitnaher zu übernehmen.

## Wartungszeitraum

Wir warten die Hauptversionen bis zu ihrem "Wartungsende". Wir pflegen im 
Allgemeinen die aktuelle und vorherige Hauptversion.
Wenn also die aktuelle Version 0.13 ist, dann wird auch 0.12 beibehalten.
Sobald 0.14 veröffentlicht wird, wird 0.12 als "Wartungsende" angesehen.
Je älter die Hauptversion, desto kritischer müssen die Probleme sein, um 
darauf zurückportiert zu werden, und um eine neue Nebenversion zu rechtfertigen.
Sobald die Software den „Wartungsende“-Zeitraum erreicht hat, erhält sie 
nur noch kritische Sicherheitsfixes bis zum End of Life „EOL“-Datum.
Nach EOL müssen Benutzer auf eine neuere Version aktualisieren, um Sicherheitsupdates 
zu erhalten, auch wenn die Community eventuell Updates für kritische Probleme bereitstellen könnte.
Im Allgemeinen wird empfohlen, den neuesten Wartungsrelease (Point-Release) der 
aktuellen oder vorherigen Hauptversion zu benutzen.

Bitte beachte, dass Nebenversionen Bugfixes, Übersetzungsaktualisierungen 
und Soft Forks erhalten. Die Übersetzung auf [Transifex][bitcoin-transifex-link] ist 
nur für die letzten beiden Hauptversionen offen.

Beispielsweise wurde die Hauptversion 0.9 am 19.03.2014 veröffentlicht, und wir 
stellten bis zum 16.06.2015 Wartungskorrekturen (Point-Releases) bereit.
Kritische Sicherheitsprobleme würden weiterhin bis zum EOL-Datum 2016-02-28 behoben.
Um jedoch von Fehlerkorrekturen profitieren zu können, muss man auf eine neuere 
Hauptversion aktualisieren.

## Zeitplan

Sobald EOL erreicht wurde, muss auf eine neuere Version aktualisieren.

| Version | Veröffentlichungsdatum | Wartungsende | End of Life |
|---------|------------------------|--------------|-------------|
{% include posts/maintenance-table.md %}

\* _Wir streben an, alle 6-7 Monate eine Hauptversion herauszubringen_

_TBA: wird noch bekannt gegeben_

## Protokollversionierung

Die obige Beschreibung beschreibt nur Bitcoin Core Softwareversionen. Viele 
andere Teile des Bitcoin-Systems enthalten ihre eigenen Versionen. Ein paar Beispiele:

- Jede **Transaktion** enthält eine Versionsnummer.
- Das **P2P-Netzwerkprotokoll** verwendet Versionsnummern, damit Nodes ankündigen 
können, welche Funktionen sie unterstützen.
- Die **integrierte Wallet** von Bitcoin Core hat eine eigene interne Versionsnummer.

Diese Versionsnummern sind bewusst von der Versionsnummer von Bitcoin Core entkoppelt, 
da das Bitcoin Core Projekt entweder keine direkte Kontrolle über sie hat 
(wie es bei Blöcken und Transaktionen der Fall ist) oder versucht, die Kompatibilität 
mit anderen Projekten aufrechtzuerhalten (wie es beim Netzwerk Protokoll der Fall ist) 
oder lässt die Möglichkeit zu, dass in einigen Versionen keine größeren Änderungen 
vorgenommen werden (wie dies manchmal bei der integrierten Wallet der Fall ist).

Das Konsensprotokoll selbst hat keine Versionsnummer.

## Beziehung zu SemVer

Die Versionierung der Bitcoin Core Software folgt nicht dem optionalen Versionierungsstandard 
[SemVer][], aber die Release-Versionierung ist oberflächlich ähnlich. SemVer wurde für die 
Verwendung in normalen Softwarebibliotheken entwickelt, in denen Einzelpersonen die 
Bibliothek in ihrem eigenen Tempo aktualisieren oder sogar auf einer älteren Version 
zurückbleiben können, wenn ihnen die Änderungen nicht gefallen.

Teile von Bitcoin, insbesondere die Konsensregeln, funktionieren so nicht. Damit 
eine neue Konsensregel in Kraft treten kann, muss sie von einer bestimmten Anzahl 
von Minern, Full Nodes oder beiden durchgesetzt werden und sobald sie in Kraft 
getreten ist, kann Software, die die neue Regel nicht kennt, ungültige Transaktionen 
generieren oder akzeptieren (obwohl Upgrades so konzipiert sind, dass dies nach 
Möglichkeit verhindert wird).

Aus diesem Grund weicht Bitcoin Core bei Änderungen an Konsensregeln und anderen 
Aktualisierungen, bei denen eine netzwerkweite Übernahme notwendig oder wünschenswert 
ist, von SemVer ab. Bitcoin Core veröffentlicht diese Änderungen als Wartungsversionen 
(`0.x.y`) statt als Hauptversionen (`0.x.0`); Dadurch wird die Größe des Patches 
minimiert, um es so vielen Personen wie möglich zu erleichtern, ihn zu inspizieren, 
zu testen und bereitzustellen. Es ermöglicht auch, denselben Patch auf mehrere frühere 
Hauptversionen zurückzuportieren, wodurch die Anzahl der Benutzer, die problemlos 
aktualisieren können, weiter erhöht wird, obwohl es nicht immer genügend Freiwillige 
gibt, um dies zu verwalten.

[SemVer]: https://semver.org/
[bitcoin-transifex-link]: https://www.transifex.com/bitcoin/bitcoin/
