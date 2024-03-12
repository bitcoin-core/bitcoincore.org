---
version: 4
title: Segregierter Witness Wallet-Entwicklungsleitfaden
name: segwit-wallet-dev
type: pages
layout: page
lang: de
permalink: /de/segwit_wallet_dev/
---
{% include toc.html %}
{% include references.md %}

Die meisten Inhalte dieses Dokuments sind in den BIPs zu Segregated 
Witness zu finden, einschließlich [BIP141][], [BIP143][], [BIP144][] 
und [BIP145][]. Bitte betrachte dies als ersten Bezugspunkt zu anderen 
verwandten Dokumenten und als Checkliste für das, was getan und was 
nicht getan werden sollte.

### Grundlegende Segregated Witness Unterstützung

Eine Wallet MUSS alle Funktionen in diesem Abschnitt implementieren, 
um auf einer grundlegenden Ebene als Segwit-kompatibel zu gelten:

#### Senden an P2SH

* Ein Segwit-kompatibles Wallet MUSS Pay-to-Script-Hash ([BIP16][]) 
und das Adressformat ([BIP13][]) unterstützen.
* Für Zahlungen muss die Wallet in der Lage sein, eine bestimmte 
P2SH-Adresse korrekt in einen <code>scriptPubKey</code> umzuwandeln 
und eine Transaktion zu erstellen.
* Für den Empfang von Zahlungen muss die Wallet in der Lage sein, 
eine P2SH-Adresse basierend auf einem P2WPKH-Skript (im Folgenden definiert) 
zu erstellen und Zahlungen an solche Adressen zu erkennen.
* Dies ist eine zwingende Voraussetzung, auch wenn die Wallet nur 
Single-Signatur-Zahlungen akzeptiert.

#### Erstellung der P2SH-P2WPKH-Adresse

* Eine P2SH-P2WPKH-Adresse ist vergleichbar mit der ursprünglichen 
  Single-Signature-P2PKH-Adresse von Bitcoin (Adresse mit Präfix 1).
* Wie jede andere P2SH-Adresse hat die P2SH-P2WPKH-Adresse das Präfix 3.
* Bis ein P2SH-P2WPKH-UTXO ausgegeben und das <code>redeemScript</code> 
  offengelegt wird, ist eine P2SH-P2WPKH-Adresse nicht von einer 
  Nicht-Segwit-P2SH-Adresse (z.B. einer Nicht-Segwit-Multi-Signatur-Adresse) 
  zu unterscheiden.
* P2SH-P2WPKH-Adressen sollten verwendet werden, wenn nur 1 öffentlicher 
  Schlüssel zum Empfangen von Zahlungen verwendet wird (wie P2PKH).
* P2SH-P2WPKH verwendet das gleiche öffentliche Schlüsselformat wie 
  P2PKH, mit einer sehr wichtigen Ausnahme: Der in P2SH-P2WPKH verwendete 
  öffentliche Schlüssel MUSS komprimiert sein, d.h. 33 Bytes groß sein 
  und mit einem <code>0x02</code> oder <code>0x03</code> beginnen. 
  Die Verwendung eines anderen Formats, wie z. B. eines unkomprimierten 
  öffentlichen Schlüssels, kann zu einem unwiderruflichen Geldverlust führen.
* So erstellt man eine P2SH-P2WPKH-Adresse:
    1. Berechne den RIPEMD160 des SHA256 eines öffentlichen Schlüssels 
    (<code>keyhash</code>). Obwohl die <code>keyhash</code>-Formel 
    dieselbe wie bei P2PKH ist, sollte die Wiederverwendung von 
    <code>keyhash</code> vermieden werden, um die Privatsphäre zu 
    verbessern und die versehentliche Verwendung eines unkomprimierten 
    Schlüssels zu verhindern.
    2. Das P2SH <code>redeemScript</code> ist immer 22 Bytes groß. 
    Es beginnt mit einem <code>OP_0</code>, gefolgt von einem canonical 
    push des <code>keyhash</code> (d.h. <code>0x0014{20-byte keyhash}</code>).
    3. Wie bei jedem anderen P2SH ist <code>scriptPubKey</code> der 
    <code>OP_HASH160 hash160(redeemScript) OP_EQUAL</code>, und die 
    Adresse ist die entsprechende P2SH-Adresse mit dem Präfix 3.

#### Transaktionsserialisierung

* Eine Segwit kompatible Wallet MUSS das ursprüngliche Transaktionsformat 
  als <code>nVersion|txins|txouts|nLockTime</code> unterstützen.
* Eine Segwit kompatibles Wallet MUSS auch das neue Serialisierungsformat 
  unterstützen, als <code>nVersion|marker|flag|txins|txouts|witness|nLockTime</code>.
* Das Format von <code>nVersion</code>, <code>txins</code>, 
  <code>txouts</code> und <code>nLockTime</code> entspricht dem Originalformat.
    * Die <code>Marker</code> MUSS <code>0x00</code> sein.
    * Das <code>Flag</code> MUSS <code>0x01</code> sein.
    * Der <code>witness</code> ist eine Serialisierung aller Witnessdaten der Transaktion.
        * Jedem TXIN ist ein Witness-Feld zugeordnet. Daher gibt es 
        keine Angabe zur Anzahl der Witness-Felder, da dies durch die 
        Anzahl der <code>txins</code> impliziert wird.
        * Jedes Witness-Feld beginnt mit einem <code>compactSize</code> 
        [integer](https://bitcoin.org/en/developer-reference#compactsize-unsigned-integers), 
        um die Anzahl der Stack-Elemente für das entsprechende <code>txin</code> 
        anzugeben. Es folgen dann ggf. Witness-Stack-Elemente für den 
        entsprechenden <code>txin</code>.
        * Jedes Witness-Stack-Element beginnt mit einer Ganzzahl 
        <code>compactSize</code>, um die Anzahl der Bytes des Elements 
        anzugeben.
        * Wenn ein <code>txin</code> mit keinen Witness-Daten verknüpft 
        ist, ist sein entsprechendes Witness-Feld exakt <code>0x00</code>, 
        das gibt an, dass die Anzahl der Witness-Stack-Elemente null ist.
* Wenn alle <code>txins</code> in einer Transaktion nicht mit Witnessdaten 
  verknüpft sind, MUSS die Transaktion im ursprünglichen Transaktionsformat 
  serialisiert werden, ohne <code>marker</code>, <code>flag</code > und 
  <code>witness</code>. Wenn beispielsweise keines der <code>txins</code> 
  von segwit UTXO stammt, MUSS es im ursprünglichen Transaktionsformat 
  serialisiert werden. (Ausnahme: Coinbase-Transaktion)
* Beispiele für die Transaktionsserialisierung findest du im Beispielabschnitt 
  von BIP143. Wallet-Entwickler können die Beispiele verwenden, um zu testen, 
  ob ihre Implementierungen das neue Serialisierungsformat korrekt parsen.

#### Transaktions-ID

* Unter Segwit hat jede Transaktion 2 IDs.
* Die Definition von <code>txid</code> bleibt unverändert: das doppelte SHA256 
  des ursprünglichen Serialisierungsformats.
* Eine neue <code>wtxid</code> wird definiert, die das doppelte SHA256 des neuen 
  Serialisierungsformats mit Witnessdaten ist.
* Wenn eine Transaktion keine Witnessdaten hat, ist ihre <code>wtxid</code> 
  dieselbe wie die <code>txid</code>.
* Die <code>txid</code> bleibt die primäre Kennung einer Transaktion:
     * Es MUSS im <code>txin</code> verwendet werden, wenn auf einen vorherigen 
       Output verwiesen wird.
     * Wenn eine Wallet oder ein Dienst derzeit <code>txid</code> verwendet, 
       um Transaktionen zu identifizieren, wird erwartet, dass du dies weiterhin mit segwit verwendest.

#### Signaturgenerierung und -verifizierung für P2SH-P2WPKH

* Für Outputs von Nicht-Segwit-UTXO bleibt der Signaturgenerierungsalgorithmus 
  unverändert.
* Für Outputs von P2SH-P2WPKH:
    * Die <code>scriptSig</code> DARF NUR einen Push des <code>redeemScript</code> 
      enthalten.
    * Das entsprechende Witness-Feld MUSS genau 2 Einträge enthalten, eine 
      Signatur gefolgt vom öffentlichen Schlüssel.
    * Es gibt einen neuen Signaturgenerierungsalgorithmus, der in [BIP143][] 
      für Segwit-Skripte beschrieben wird. Entwickler sollten die Anweisungen 
      sorgfältig befolgen und das P2SH-P2WPKH-Beispiel in 
      [BIP143](https://github.com/bitcoin/bips/blob/master/bip-0143.mediawiki#P2SHP2WPKH) 
      verwenden, um sicherzustellen, dass sie in der Lage sind, den 
      <code>sighash</code> zu reproduzieren.
    * Der [BIP143][]-Signaturgenerierungsalgorithmus deckt den eingegebenen 
      Ausgabewert ab, was das Design von Air-Gap-Light-Weight-Wallets und 
      Hardware-Wallets vereinfacht.
    * Bitte beachte, dass für ein P2SH-P2WPKH der <code>scriptCode</code> 
      immer 26 Bytes einschließlich des führenden Größenbytes ist, als 
      <code>0x1976a914{20-byte keyhash}88ac</code>, NICHT das <code>redeemScript</code> 
      noch der <code>scriptPubKey</code>.
    * [Beispiel](https://blockchainprogramming.azurewebsites.net/checktx?txid=8139979112e894a14f8370438a471d23984061ff83a9eba0bc7a34433327ec21).
    
#### Netzwerkdienste (optional)

* Netzwerkdienste werden benötigt, wenn die Wallet Transaktionen über das 
  Peer-to-Peer-Netzwerk senden und empfangen würde.
* Segwit fähige Nodes geben bekannt, dass sie Witnesses durch das Service 
  Bit bereitstellen können: <code>NODE_WITNESS = (1 << 3)</code>
* Transaktionen ohne Witnessdaten (und daher im Originalformat serialisiert) 
  können an Nodes mit oder ohne <code>NODE_WITNESS</code>-Unterstützung gesendet werden.
* Transaktionen, die Segwit-UTXOs ausgeben (und daher im neuen Format serialisiert werden), 
  DÜRFEN NUR an Nodes mit <code>NODE_WITNESS</code>-Unterstützung gesendet werden.
* Transaktionen, die Segwit-UTXOs ausgeben, aber mit vereinfachten Zeugendaten 
  (und daher im Originalformat serialisiert) werden möglicherweise an Nodes 
  ohne <code>NODE_WITNESS</code>-Unterstützung gesendet. Solche Transaktionen 
  sind jedoch nach der Aktivierung von segwit ungültig und würden nicht in einem 
  Block akzeptiert werden.
* Details zu den Netzwerkdiensten finden sich in [BIP144][].

#### Datenschutz der Benutzer

* Bis Segwit-Transaktionen alltäglich sind, kann dieser Transaktionstyp 
  das Bitcoin-Tracking erleichtern.
* Die Verwendung von P2SH-P2WPKH als Standardänderungsausgabe kann sich auch 
  auf den Datenschutz auswirken.

#### Schätzung der Transaktionsgebühr

* Anstelle der Transaktionsgröße wird eine neue Metrik namens "virtual size" 
  (<code>vsize</code>) definiert.
* <code>vsize</code> einer Transaktion entspricht dem 3-fachen der Größe mit 
  ursprünglicher Serialisierung plus der Größe mit neuer Serialisierung, das 
  Ergebnis durch 4 teilen und auf die nächste ganze Zahl aufrunden. Wenn 
  beispielsweise eine Transaktion mit neuer Serialisierung 200 Bytes groß 
  ist und durch Entfernen von <code>Marker</code>, <code>Flag</code> und 
  <code>Witness</code> zu 99 Bytes wird, wird die <code>vsize</code> zu 
  (99 * 3 + 200) / 4 = 125 mit Aufrundung.
* <code>vsize</code> einer Nicht-Segwit-Transaktion ist einfach ihre Größe.
* Die Transaktionsgebühr sollte geschätzt werden, indem die <code>vsize</code> 
  mit anderen Transaktionen verglichen wird, nicht die Größe.
* Entwickler sollten darauf achten, bei der Schätzung der Gebühren keinen 
  "off-by-4-times" Fehler zu machen.


#### Aktivierung {#upgrade-safety}

* Ab Blockhöhe 481824 begannen alle SegWit-fähigen Nodes mit der Einführung 
  der neuen SegWit-Konsensregeln.

#### Rückwärtskompatibilität

* Das Senden und Empfangen von Legacy-P2PKH-Zahlungen (Adresse mit Präfix 1) 
  sollte weiterhin unterstützt werden.


### Komplexe Skriptunterstützung

Wenn ein Wallet andere Skripttypen als nur eine einzelne Signatur unterstützt, 
wie z. B. Multi-Signatur, muss es mehr als die grundlegenden Anforderungen erfüllen:

#### Erstellung der P2SH-P2WSH-Adresse

* Eine P2SH-P2WSH-Adresse ist vergleichbar mit der ursprünglichen P2SH-Adresse 
  von Bitcoin, die die Repräsentation beliebig komplexer Skripte mit einer 
  Adresse mit fester Größe ermöglicht.
* Wie jede andere P2SH- und P2SH-P2WPKH-Adresse hat die P2SH-P2WSH-Adresse 
  das Präfix 3. Sie sind nicht zu unterscheiden, bis das UTXO ausgegeben wird.
* So erstellst du eine P2SH-P2WSH-Adresse:
    1. Definiere ein Skript namens (<code>witnessScript</code>).
    2. Berechne den SHA256 des <code>witnessScript</code> (<code>scripthash</code>). 
       Bitte beachte, dass ein einzelner SHA256 verwendet wird, kein doppelter 
       SHA256 oder RIPEMD160 (SHA256).
    3. Das P2SH <code>redeemScript</code> ist immer 34 Bytes groß. Es beginnt 
       mit einem <code>OP_0</code>, gefolgt von einem "canonical push" des 
       <code>scripthash</code> (d. h. <code>0x0020{32-byte scripthash}</code>).
    4. Wie bei jedem anderen P2SH, der <code>scriptPubKey</code> ist 
       <code>OP_HASH160 hash160(redeemScript) OP_EQUAL</code>, und die Adresse 
       ist die entsprechende P2SH-Adresse mit dem Präfix 3.
* Einschränkungen des Skripts:
    * Die Skriptauswertung darf nicht fehlschlagen und MUSS nach der Auswertung 
      ein und nur ein TRUE-Stack-Element hinterlassen. Andernfalls ist die 
      Auswertung fehlgeschlagen.
    * Jeder öffentliche Schlüssel in P2SH-P2WSH-Skripten MUSS ein komprimierter 
       Schlüssel sein, sonst kann das Geld dauerhaft verloren gehen.
    * Wenn OP_IF oder OP_NOTIF verwendet wird, MUSS das Argument entweder ein 
      leerer Vektor (für falsch) oder <code>0x01</code> (für wahr) sein. Die 
      Verwendung anderer Werte kann zu dauerhaften Geldverlusten führen. 
      ([BIP-Entwurf](https://lists.linuxfoundation.org/pipermail/bitcoin-dev/2016-August/013014.html))
    * Wenn ein OP_CHECKSIG oder OP_CHECKMULTISIG einen Fehler zurückgibt, 
      müssen alle Signaturen leere Vektoren sein. Andernfalls kann das Guthaben 
      dauerhaft verloren gehen. ([BIP146][])
    * Es gibt eine standardmäßige Begrenzung für das <code>witnessScript</code> 
      bei 3600 Bytes. Mit Ausnahme des <code>witnessScript</code> können höchstens 
      100 Witness-Stack-Elemente mit jeweils höchstens 80 Bytes vorhanden sein. 
      Transaktionen, die diese Grenzen überschreiten, dürfen nicht weitergeleitet 
      oder in einen Block aufgenommen werden.
    * Viele der ursprünglichen Konsensbeschränkungen für Skripte, wie z.B. 10000 
      Bytes Skriptgröße, 201 <code>nOpCount</code>, gelten immer noch für P2SH-P2WSH.
    * Die 520-Bytes-Skriptgrößenbeschränkung für P2SH gilt nicht für P2SH-P2WSH. 
      Es wird durch das Limit von 3600 Bytes und das Konsenslimit von 10000 Bytes ersetzt.
    
#### Signaturgenerierung und -überprüfung für P2SH-P2WSH

* Für P2SH-P2WSH Outputs:
    * Die <code>scriptSig</code> DARF NUR einen Push des <code>redeemScript</code> enthalten.
    * Das letzte Witness-Element des entsprechenden Witness-Feld MUSS das <code>witnessScript</code> sein.
    * Der neue Signaturgenerierungsalgorithmus [BIP143][] wird angewendet:
        * Ohne Verwendung von OP_CODESEPARATOR ist der <code>scriptCode</code> ein 
          <code>witnessScript</code>, dem eine <code>compactSize</code>-Ganzzahl für 
          die Größe von <code>witnessScript</code>. Wenn das Skript beispielsweise 
          OP_1 (<code>0x51</code>) ist, lautet der serialisierte <code>scriptCode</code> 
          (<code>0x0151</code>).
        * Für alle ungewöhnlichen Skripte, die OP_CODESEPARATOR enthalten, lese 
          bitte [BIP143][] für die genaue Semantik.
    * Alle Witness-Stack-Elemente vor dem <code>witnessScript</code> werden als 
      Eingabe-Stack für die Skriptauswertung verwendet. Der Eingabe-Stack wird 
      nicht als Skript interpretiert. Beispielsweise ist es nicht erforderlich, 
      einen <code>0x4c</code> (OP_PUSHDATA1) zu verwenden, um ein großes Element zu "pushen".
    * Um die Korrektheit der Signaturgenerierung und Stack-Serialisierung zu 
      überprüfen, teste bitte immer anhand der Beispiele in [BIP143][].
    * [Beispiel] (https://blockchainprogramming.azurewebsites.net/checktx?txid=954f43dbb30ad8024981c07d1f5eb6c9fd461e2cf1760dd1283f052af746fc88).

### Segwit native Adressen (optional) {#advanced-designs}

Die folgenden Funktionen sind für die anfängliche Segwit-Unterstützung nicht erforderlich.

#### Native Pay-to-Witness-Public-Key-Hash (P2WPKH)

* Natives P2WPKH ist ein <code>scriptPubKey</code> von 22 Bytes. Es beginnt 
  mit einem <code>OP_0</code>, gefolgt von einem "canonical push" des 
  <code>keyhash</code> (d. h. <code>0x0014{20-byte keyhash}</code>).
* Wie bei P2SH-P2WPKH, <code>keyhash</code> ist RIPEMD160(SHA256) eines 
  komprimierten öffentlichen Schlüssels.
* Bei dem Output eines nativen P2WPKH MUSS <code>scriptSig</code> leer sein, 
  und das Witness-Stack-Format und die Signaturgenerierungsregeln sind dieselben 
  wie bei P2SH-P2WPKH (einschließlich der Anforderung, einen komprimierten 
  öffentlichen Schlüssel zu verwenden).
* [Beispiel](https://blockchainprogramming.azurewebsites.net/checktx?txid=d869f854e1f8788bcff294cc83b280942a8c728de71eb709a2c29d10bfe21b7c).

#### Native Pay-to-Witness-Script-Hash (P2WSH)

* Natives P2WSH ist ein <code>scriptPubKey</code> von 34 Bytes. Es beginnt 
  mit einem <code>OP_0</code>, gefolgt von einem "canonical push" des <code>scripthash</code> (d. h. <code>0x0020{32-byte scripthash}</code>).
* Wie bei P2SH-P2WSH, <code>scripthash</code> ist SHA256 des <code>witnessScript</code>.
* Bei dem Output eines nativen P2WSH MUSS <code>scriptSig</code> leer sein, und 
  das Witness-Stack-Format und die Signaturgenerierungsregeln sind dieselben 
  wie bei P2SH-P2WSH (einschließlich der Anforderung, einen komprimierten 
  öffentlichen Schlüssel zu verwenden).
* [Beispiel](https://blockchainprogramming.azurewebsites.net/checktx?txid=78457666f82c28aa37b74b506745a7c7684dc7842a52a457b09f09446721e11c).

#### Warum und wie werden native (Bech32) P2WPKH und P2WSH verwendet?

* [BIP173][] schlägt ein Prüfsummen-Base32-Format (Bech32) für native P2WPKH- 
  und P2WSH-Adressen vor.
* Unterstützung für Bech32-Adressen war in Bitcoin Core v0.16.0 enthalten.
* Im Vergleich zu den P2SH-Versionen ist die Transaktions <code>vsize</code> der 
  nativen Versionen in den meisten Fällen kleiner, und daher kann eine geringere Gebühr anfallen.
* Natives P2WPKH und P2WSH können mit rohen <code>scriptPubKey</code>-Protokollen 
  wie dem Payment Protocol (BIP70) verwendet werden. Dies kann jedoch die 
  Privatsphäre des Zahlers und des Empfängers beeinträchtigen (siehe unten).
* Natives P2WPKH und P2WSH können als Standardänderungsadresse verwendet 
  werden, aber dies kann es anderen Personen ermöglichen, die Änderung leicht 
  zu identifizieren (siehe unten).
* Bis natives P2WPKH und P2WSH weit verbreitet sind, können diese Adresstypen 
  bei Benutzern Datenschutzbedenken hervorrufen.


### Beispiele für Skripts und Transaktionen

* [Beispiele für verschiedene Witness-Transaktionstypen und Tools zur 
  Überprüfung der Transaktionsgültigkeit](https://blockchainprogramming.azurewebsites.net/checktx)
* [BIP141][]
* [BIP143][]
* [BIP173][]
* [Skripttests](https://github.com/bitcoin/bitcoin/blob/master/src/test/data/script_tests.json)
* [Gültige Transaktionstests](https://github.com/bitcoin/bitcoin/blob/master/src/test/data/tx_valid.json)
* [Ungültige Transaktionstests](https://github.com/bitcoin/bitcoin/blob/master/src/test/data/tx_invalid.json)
