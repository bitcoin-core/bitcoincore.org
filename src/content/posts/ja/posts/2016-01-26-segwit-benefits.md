---
type: posts
lang: ja
name: segwit-benefits
id: ja-segwit-benefits
title: Segregated Witnessの恩恵
permalink: /ja/2016/01/26/segwit-benefits/
version: 1
excerpt: このページはSegregated Witnessの恩恵についてまとめています。
---
{% include toc.html %}

Segregated Witnessのソフトフォーク（segwit）には幅広い機能が含まれており、その多くは高度に技術的です。このページにはこれらの機能の恩恵の一部がまとめられています。

## Malleabilityの解決 {#malleability-fixes}

Bitcoinのトランザクションはトランザクション識別子(txid)と呼ばれる64桁の16進数のハッシュで識別されます。このハッシュはトランザクションで使用されるコインとそのトランザクションの結果使用可能になるコインの両方の情報に基いて生成されます。

残念ながら、txidの計算方法上、トランザクションの意味を変えない小さな変更を加えることでtxidは変更されます。これは第三者によるmalleabilityと呼ばれています。BIP 62("dealing with malleability")は、この問題に少しずつ取り組もうとしましたが、コンセンサスに実装するには複雑過ぎ取りやめました。

例えば、あなたがtxidがef74...c309のトランザクションをネットワークに送信したとして、あなたのトランザクションを中継するネットワーク上のノードや、あなたのトランザクションのブロックに入れるマイナーなどの第三者がトランザクションを少し変更しても、あなたのトランザクションは同じコインを同じアドレスに支払うことになりますが、そのトランザクションはあなたが送信したものとは全く違うtxidで承認されています。

より一般的には、トランザクションの署名者の１人もしくは複数人が署名を変更すると、トランザクションは有効なまま同じ金額を同じアドレスに支払いますが、txidには署名も含まれるためtxidは別の値に変更されます。（インプットやアウトプットを変更せず）署名データを変更することでトランザクションを変更する方法をscriptSig malleabilityと呼びます。

Segwitでは、Bitcoinユーザーがトランザクションの可搬性のある部分を*transaction witness*に移動させ、witnessへの変更がtxidの計算に影響を与えないようwitnessを分離することによって第三者およびscriptSigによるmalleabilityを防ぎます。

### 誰に恩恵が？ {#who-benefits-1}

- **使用済みのビットコインをトラッキングするウォレットの作成者:** 自分が送信したトランザクションの状態を監視する一番簡単な方法はtxidで検索する方法です。しかし第三者のmalleabilityを考慮するとウォレットはtxidが変更されることを考慮したコードを実装する必要があります。

- **未承認トランザクションを使った取引:** アリスがボブにトランザクション １で支払いをし、ボブがそれを使ってトランザクション 2でチャーリーに支払いをした場合、アリスの支払いトランザクションが改竄され別のtxidで承認されるとトランザクション 2は無効となりチャーリーには支払いが行われません。もしボブが信頼できるなら、彼はチャーリーに支払いを再発行するでしょう。そうでない場合、ボブはビットコインを手元にとどめておくことができます。

- **Lightning Network:** 第三者とscriptSigによるmalleabilityが解決すると、Lightning Networkの実装の複雑さが減り、ブロックチェーン上のスペースの使用が大幅に効率化されます。scriptSigのmalleabilityが解決することで、各LightningのクライアントはBitcoinのフルノードを必要とせず、ブロックチェーンの監視をアウトソースする軽量なLightningクライアントを実行することも可能になります。

- **ブロックチェーンを使用するすべての人:** マイクロペイメントチャネルや予想される新しいスマートコントラクトなど現在のスマートコントラクトの設計、理解、監視がより簡単になります。

注意： Segwitのトランザクションはそのインプットが全てSegwitを使用している（直接もしくは後方互換性のあるP2SHアドレスを使用している）場合にのみ、そのmalleabilityを回避します。

### より詳しい情報 {#further-information-1}

 * [Bitcoin Wiki on Malleability](https://en.bitcoin.it/wiki/Transaction_Malleability)
 * [Coin Telegraph article on 2015 Malleability attack](http://cointelegraph.com/news/115374/the-ongoing-bitcoin-malleability-attack)
 * [Bitcoin Magazine article on 2015 Malleability attack](https://bitcoinmagazine.com/articles/the-who-what-why-and-how-of-the-ongoing-transaction-malleability-attack-1444253640)
 * ["Overview of BIPs necessary for Lightning" transcript](http://diyhpl.us/wiki/transcripts/scalingbitcoin/hong-kong/overview-of-bips-necessary-for-lightning/)
 * [BIP 62](https://github.com/bitcoin/bips/blob/master/bip-0062.mediawiki)
 * [BIP 140 -- alternative approach to malleability fixes](https://github.com/bitcoin/bips/blob/master/bip-0140.mediawiki)
 * [Stack exchange answer regarding 683f...8bfa transaction](http://bitcoin.stackexchange.com/questions/22051/transaction-malleability-in-the-blockchain/22058#22058)

## sighash操作の線形スケーリング {#linear-scaling-of-sighash-operations}

Bitcoinのブロックサイズを増やす簡単なアプローチの主な問題は、特定のトランザクションでは署名ハッシュの計算のスケールが線形ではなく二次関数的にスケールすることです。

![Linear versus quadratic](/assets/images/linear-quad-scale.png)

本質的に、トランザクションのサイズを2倍にすると署名操作の数と、検証されるそれぞれの署名についてハッシュする必要があるデータの量の両方が2倍になります。これは個々のブロックの検証に25秒を要し、悪意をもって作られたトランザクションの検証には3分以上かかるようになります。

Segwitは、トランザクションの各バイトを最大2回だけハッシュすればいいように署名のトランザクションハッシュの計算を変更することでこれを解決します。これにより同じ機能がより効率的になるため、悪意あるトランザクションが作成されたり、より大きなブロックがサポートされるようになったとしても、署名ハッシュの計算による問題が発生することなく大きなトランザクションを作ることができるようになります。

### 誰に恩恵が？ {#who-benefits-2}

署名検証のためハッシュデータの二次関数的なスケールが無くなると、より安全にブロックサイズを増やすことができます。またトランザクションサイズを制限することなくこれを行うことで、Bitcoinはマイニング報酬やクラウドファウンディングの支払いなどの大きなグループからの決済を引き続きサポートすることができます。

変更されたハッシュはwitnessデータにある署名操作にのみ適用されるため、ベースブロックの署名操作にはこれまで通り低い制限が必要です。

### より詳しい情報 {#further-information-2}

 * [BIP 143](https://github.com/bitcoin/bips/blob/master/bip-0143.mediawiki)
 * [Blog post by Rusty Russell on the 25s transaction](http://rusty.ozlabs.org/?p=522)
 * [CVE 2013-2292 on Bitcoin wiki](https://en.bitcoin.it/wiki/Common_Vulnerabilities_and_Exposures#CVE-2013-2292)
 * [Proposal to limit transactions to 100kB](https://lists.linuxfoundation.org/pipermail/bitcoin-dev/2015-July/009494.html)
 * [Bitcoin Classic commit on 0.11.2 branch adding additional consensus limit on sighash bytes](https://github.com/bitcoinclassic/bitcoinclassic/commit/842dc24b23ad9551c67672660c4cba882c4c840a)

## pay-to-script-hash (P2SH) によるマルチシグのセキュリティ強化 {#increased-security-for-multisig-via-pay-to-script-hash}

マルチシグの支払いは、現在160-bitのHASH160アルゴリズム（SHA256してRIPEMDしたもの）によって保護されているP2SHを使っています。しかし、署名者の1人が全ての資金を盗みたい場合、80bit(2<sup>80</sup>)の作業で全ての資金を入手することが可能なマルチシグスクリプトの一部として有効なアドレスとスクリプトの衝突を見つけることができます。これは非常に潤沢なりソースを持つ攻撃者にとって可能性のある範囲です。（比較のため、Bitcoinのマイニングネットワークは毎秒1エクサハッシュで2週間で80bit分の作業をしています。）

Segwitは単一の公開鍵への直接的な支払いには（この種の攻撃には役に立たないため）HASH160を使用し、スクリプトハッシュへの支払いは256-bitのSHA256ハッシュを使用すことでこれを解決します。

### 誰に恩恵が？ {#who-benefits-4}

Segwitを利用してマルチシグもしくはスマートコントラクトへの支払いを行う全てのユーザーが、スクリプトのための追加のセキュリティの恩恵を受けます。

### より詳しい情報 {#further-information-4}

 * [Gavin Andresen asking if 80-bit attacks are worth worrying about](https://lists.linuxfoundation.org/pipermail/bitcoin-dev/2016-January/012198.html)
 * [Ethan Heilman describing a cycle finding algorithm](https://lists.linuxfoundation.org/pipermail/bitcoin-dev/2016-January/012202.html)
 * [Rusty Russell calculating costs of performing an attack](https://lists.linuxfoundation.org/pipermail/bitcoin-dev/2016-January/012227.html)
 * [Anthony Towns applying the cycle finding algorithm to exploit transactions](https://lists.linuxfoundation.org/pipermail/bitcoin-dev/2016-January/012218.html)
 * [Gavin Andresen summarising the thread](https://lists.linuxfoundation.org/pipermail/bitcoin-dev/2016-January/012234.html)

## スクリプトのバージョンニング {#script-versioning}

Bitcoinのスクリプトの変更によりセキュリティの強化と機能性が向上しました。しかし、既存のスクリプトの設計では、予め確保されている10個のOP_NOP opcodeをある条件で失敗するような新しいopcodeに置き換えることで後方互換のある（ソフトフォーク）変更のみが可能です。これは新しい署名方法やOP_CLTVのような機能を導入するのには十分ですが、どちらもハッキーで（例えばOP_CLTVは通常OP_DROPと組み合わせて使用するなど）2つの文字列を結合するなどの単純な機能を有効にするために使用することはできません。

Segwitはスクリプトにバージョン番号を含めることでこれを解決しているため、スクリプトバージョンを増やすだけで、非Segwitのトランザクションではハードフォークをする必要があるような追加のopcodeをサポートすることができます。

### 誰に恩恵が？ {#who-benefits-5}

スクリプトのopcodeの変更が容易になると、Bitcoinの高度なスクリプトの作成が容易になります。これにはSchnorr署名の導入や鍵のリカバリを使用した署名サイズの縮小、サイドチェーンのサポート、Merklized Abstract Syntax Trees (MAST)を使用したよりスマートなコントラクトの作成やその他の研究レベルのアイディアが含まれます。

## UTXOの成長を縮小 {#reducing-utxo-growth}

未使用のトランザクションアウトプット（UTXO）のデータベースは、新しいトランザクションが有効かどうか判断するため各検証Bitcoinノードによって維持されます。ネットワークの効率的な運用のためには、このデータベースへのクエリ及び変更を速く行う必要があり、理想的にはインメモリに収めるべきです。そのためにもデータベースのサイズをできるだけ小さく保つことが重要になります。

ただこれはBitcoinが成長するにつれ困難になります。各新規ユーザーは少なくとも1つのUTXOエントリーを持ち、プライバシーや柔軟性を向上させるためや、ペイメントチャネルやその他のスマートコントラクトを使用するのに複数のエントリーを持つようになるでしょう。

Segwitは、UTXOセットのサイズに影響を与えない署名データに対して、UTXOセットのサイズに影響を与える署名データよりコストを75%安くすることでこれを改善します。これにより、ユーザーに手数料を最小限に抑えるためにUTXOセットへの影響を最小限に抑えるトランザクションを使用することを奨励し、開発者にはUTXOセットへの影響を最小限に抑える方法でスマートコントラクトや新しい機能の設計を促すことが期待されます。

Segwitはソフトフォークによる変更であるため、ベースのブロックサイズは増えず、UTXOセットのワーストケースの成長率は同じままです。

### 誰に恩恵が？ {#who-benefits-6}

UTXOの成長の縮小は、マイナーや企業およびフルノードを実行するユーザーにとって恩恵があり、より多くのユーザーがシステムに参加する中、Bitcoinネットワークの現在のセキュリティを維持するのに役立ちます。UTXOセットの成長を最小限にしようとするユーザーや開発者は、UTXOの成長に対するトランザクションの影響を無視する人と比較して、手数料が安いというメリットがあります。

### より詳しい情報 {#further-information-5}

 * [Statoshi UTXO dashboard](http://statoshi.info/dashboard/db/unspent-transaction-output-set)

## 署名を検証しない場合の効率が向上 {#efficiency-gains-when-not-verifying-signature}

過去のトランザクションの署名は今後のトランザクションの署名より重要でない場合があります。例えばBitcoin Coreはデフォルトで最新のチェックポイントより前のトランザクションの署名をチェックしません。またSPVクライアントは署名自体を一切チェックせず、既にマイナーや他のノードによって検証された結果を信頼します。ただし今のところ、署名データはトランザクションにとって不可欠なデータで、トランザクションハッシュを計算するのに必ず必要です。

署名データを分離することで、署名データに関心の無いノードはディスクから署名データを削除したり、最初からダウンロードするのを避けることでリソースを節約することができます。

### 誰に恩恵が？ {#who-benefits-7}

Segwitアドレスを使用するトランザクションが増えると、プルーニングノードやSPVノードを実行しているユーザーは、より少ない帯域幅とディスク量で動作するようになります。

## ブロック容量/サイズの増加 {#block-capacity-size-increase}

旧ノードはwitnessの含まれないブロックデータをダウンロードするので、そのデータに対して1MBのブロックサイズの制限ルールが適用されます。
このため、witnessデータを含む完全なブロックを理解する新しいノードは、この制限を新しいものに自由に置き換えることができ、より大きなブロックサイズが可能になります。したがってSegregated witnessはこの機会を利用してブロックサイズの上限を4MB近くに引き上げ、リソースの使用についてブロックのバランスを取るための新たなコスト制限を追加します（これにより事実上ブロックのサイズは1.6MB〜2MBになります）。

### 誰に恩恵が？ {#who-benefits-8}

アップグレードされたウォレットを使用するユーザーは、署名をトランザクションのwitness部分に移動することで、ブロックサイズの増加を利用できます。

## 単一の結合されたブロック制限に向けた取り組み {#moving-towards-single-combined-block-limit}

現在ブロックサイズには2つのコンセンサス制限があります。ブロックは1MBを超えることはできません。またブロック内の全トランザクションで20,000回以上の署名チェックを実行することはできません。

1つの制限が課せられたブロックに含めるトランザクションの最も収益性の高いセットを見つけるのは、ナップザック問題の応用で単純な貪欲法でほぼ完全に解決することができます。しかし2つめの制限が加わると、場合によっては良い解決策を見つけるのがとても困難になります。この理論上の問題は、ブロックを下回る大きさでマイニングされること強制するのに実際に使われています。

ハードフォークやブロックサイズを大幅に減らすこと無くこの問題を解決することはできません。Segwitでは問題を解決できないため、悪化させないようにします。Segregated witnessのデータに独立した制限を導入するのではなく、代わりにUTXOデータとwitnessデータの加重合計に1つの制限が適用され、両方を結合したエンティティとして同時に制限を適用します。

### 誰に恩恵が？ {#who-benefits-9}

ブロック容量の制限を変更する将来のハードフォークが1つのパラメータの加重合計になる場合、最終的にマイナーが恩恵を受けます。例えば、

    50*sigops + 4*basedata + 1*witnessdata < 10M

これによりマイナーは、手数料収入が最大になるようにブロックを簡単かつ正確に埋めることができ、ユーザーはトランザクションがマイニングされるのに必要な手数料をより確実に計算できるようになります。

### より詳しい情報 {#further-information-6}

 * [Knapsack problem](https://en.wikipedia.org/wiki/Knapsack_problem)
 * [Sigop attack discussion on bitcointalk in Aug 2015](https://bitcointalk.org/index.php?topic=1166928.0;all)
 * [Gregory Maxwell on bitcoin-dev on witness limits](https://lists.linuxfoundation.org/pipermail/bitcoin-dev/2015-December/011870.html)
 * ["Validation Cost Metric" transcript](http://diyhpl.us/wiki/transcripts/scalingbitcoin/hong-kong/validation-cost-metric/)

## 更新 2016-10-19 {#update-2016-10-19}

このページの前のバージョンでは、Segwitの恩恵として"Compact fraud proofs"が記載されていました。しかし、実装されているように、SegwitはSegwitの有無に関わらずこれをより簡単にするものではありません。将来のソフトフォークでCompact fraud proofsを有効にしその恩恵を受けるには、Segwitで使われているコミットメントデータを拡張するのではなく、（コインベーストランザクションなどに）独自のコミットメントを導入する必要があります。

以前のテキストは:

> **Compact fraud proofs**
>
> Bitcoinのユーザーベースが拡張されると、ブロックチェーン全体を自然に検証するのはより高価になります。分散化されたトラストレスなBitcoinの性質を維持するためには、ブロックチェーン全体の検証をする余裕がないユーザーも、安価に検証できるようにしてブロックチェーン全体の検証ができるようになることが重要です。
>
> Segwitは将来のソフトフォークでコミットメントデータを含められるようwitnessの構造を拡張することで、軽量（SPV）クライアントでも、ブロックで生成されたビットコインの数やブロックのサイズ、1ブロック内で実行されるsigopsの数などのコンセンサスルールの適用をできるようにします。
>
> **誰が恩恵を受ける？**
>
> Fraud proofsによりSPVのユーザーはBitcoinのコンセンサスルールを適用することができるようになり、個々のユーザーが攻撃される方法を削減するだけでなく、Bitcoinネットワーク全体のセキュリティを大幅に向上させる可能性があります。
>
> これらのFraud proofsは将来のソフトフォークの一部としてwitnessのデータ構造に追加することができ、Segwitを使用しないトランザクションであってもSPVクライアントがルールを適用するのに役立ちます。

## 更新 2020-06-23 {#update-2020-06-23}

このページの以前のバージョンではSegwitの恩恵として「インプットの量への署名」が挙げられていました。
しかし、現状の実装ではSegwitの有無に関わらず、これを安全にはできません。
署名されたインプットの値に依拠するためには将来のソフトフォークが必要になります。

各インプットの値は個別に署名されているため、見かけ上手数料を欺く操作が可能です(CVE-2020-14199)。

以前のテキストは：

> ## インプットの量への署名
> 
> ハードウェアウォレットがトランザクションに署名する際、使用される合計金額を簡単に確認できますが、インプットの全てのトランザクションの完全なコピーを持つ場合のみ安全に正確な手数料が決まり、偽のデータが提供されていないことを保証するためにそれらをハッシュする必要があります。署名されたトランザクション自体はとても小さくても、そのインプットが参照する個々のトランザクションは最大1MBになる可能性もあるため、これらは必ずしも安価な操作であるとは限りません。
> 
> Segwitではインプットの量を明示的にハッシュすることでこれを解決します。これはハードウェアウォレットにトランザクションのハッシュ、インデックスおよびコインの量を単純に与え（および使用される公開鍵が分かっていること）、使用されるトランザクションの規模や複雑さに関わらず安全にトランザクションに署名することができるようになることを意味します。
> 
> ### 誰に恩恵が？
> 
> ハードウェアウォレットの製造元とユーザーは明確に恩恵を受けます。しかしこれは"Internet of things"アプリケーションのための小さな組み込みデバイスでBitcoinを安全に使用する方がより簡単です。
> 
> この恩恵は、Segwit対応アドレス（もしくはP2SHでネストされたSegwitアドレス）に送信されたトランザクションを使用する場合にのみ使用できます。
> 
> ### より詳しい情報
> 
>  * [BIP 143](https://github.com/bitcoin/bips/blob/master/bip-0143.mediawiki)
