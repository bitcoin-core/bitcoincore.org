---
version: 3
title: Segregated Witnessウォレット開発ガイド
name: segwit-wallet-dev
type: pages
layout: page
lang: ja
permalink: /ja/segwit_wallet_dev/
version: 1
---
{% include toc.html %}
{% include references.md %}

このドキュメントのほとんどの内容は、[BIP141][]や [BIP143][]、 [BIP144][]、 [BIP145][]などのsegregated witness関連のBIPに記載されています。このドキュメントは、他の関連ドキュメントへの参照点、また何をすべきで何をすべきではないかのチェックリストと考えてください。

### 基本的なsegregated witnessのサポート {#basic-segregated-witness-support}

ウォレットはsegwit互換であると認識されるためには、このセクションの全機能を実装しなければなりません。

#### P2SHへの送信 {#sending-to-p2sh}

* segwit互換のウォレットは、pay-to-script-hashと（[BIP16][]）とそのアドレスフォーマット（[BIP13][]）をサポートしなければなりません。
* 支払いを行う際、ウォレットは与えられたP2SHアドレスを<code>scriptPubKey</code>に適切に変換し、トランザクションを作成しなければなりません。
* 支払いを受け取る際、ウォレットはP2WPKHスクリプト（後述）をベースにしたP2SHアドレスを作成し、そのようなアドレスへの支払いを認識できる必要があります。
* 単一署名の支払いのみを受け入れるウォレットの場合でも、これは必須要件です。

#### P2SH-P2WPKHアドレスの作成 {#creation-of-p2sh-p2wpkh-address}

* P2SH-P2WPKHは単一署名の（1から始まる）P2PKHアドレスと互換性があります。
* 他のP2SHアドレスと同様にP2SH-P2WPKHアドレスも3から始まるアドレスです。
* P2SH-P2WPKHのUTXOが使用され、その<code>redeemScript</code>が公開されるまで、P2SH-P2WPKHアドレスは（非segwitのマルチシグアドレスなどの）非segwitのP2SHアドレスと区別できません。
* P2SH-P2WPKHアドレスは、P2PKHのように支払いを受け取るために１つの公開鍵のみを使用する場合に使われます。
* P2SH-P2WPKHはP2PKHと同じ公開鍵フォーマットを使用しますが、重要な違いは、P2SH-P2WPKHで使用される公開鍵は必ず圧縮公開鍵である必要があります。33バイトで、<code>0x02</code>もしくは<code>0x03</code>で始まる公開鍵です。非圧縮の公開鍵など他の形式の公開鍵を使用すると、取り返しのつかない資金損失につながる可能性があります。
* P2SH-P2WPKHアドレスを作成するには：
    1. 公開鍵のSHA256のRIPEMD160(<code>keyhash</code>)を計算します。 <code>keyhash</code>の形式はP2PKHと同じですが、圧縮されていない鍵を誤って使用するのを防ぎ、プライバシー強化のためにも<code>keyhash</code>の再利用は避けるべきです。
    2. P2SHの<code>redeemScript</code>は常に22バイトです。それは<code>OP_0</code>から始まり、その後に<code>keyhash</code>がプッシュされます。（例：<code>0x0014{20-byte keyhash}</code>）
    3. 他のP2SHと同様、<code>scriptPubKey</code>は<code>OP_HASH160 hash160(redeemScript) OP_EQUAL</code>となり、そのアドレスは3から始まるP2SHアドレスです。

#### トランザクションのシリアライゼーション {#transaction-serialization}

* segwit互換のウォレットは<code>nVersion|txins|txouts|nLockTime</code>形式のオリジナルのトランザクションフォーマットをサポートしなければなりません。
* segwit互換のウォレットは<code>nVersion|marker|flag|txins|txouts|witness|nLockTime</code>形式の新しいシリアライゼーションフォーマットをサポートしなければなりません。
    * <code>nVersion</code>と<code>txins</code>、<code>txouts</code>、<code>nLockTime</code>はオリジナルの形式と同じです。
    * <code>marker</code>は必ず<code>0x00</code>です。
    * <code>flag</code>は必ず<code>0x01</code>です。
    * <code>witness</code>はトランザクションの全witnessデータをシリアライズしたデータです。
        * 各txinはwitnessフィールドに関連付けられています。その結果、witnessフィールドの数を示すフィールドは無く、暗黙的に<code>txins</code>の数がその数を示します。
        * 各witnessフィールドは対応する<code>txin</code>のスタックアイテムの数を示す<code>compactSize</code> [integer](https://bitcoin.org/en/developer-reference#compactsize-unsigned-integers) で始まります。その後に対応する<code>txin</code>のwitnessスタックアイテムが続きます。
        * 各witnessスタックアイテムは、そのスタックアイテムのバイト数を示す<code>compactSize</code> integerで始まります。
        * <code>txin</code>がwitnessデータと関連付けられていない場合は、対応するwitnessフィールドの値は<code>0x00</code>となり、witnessスタックアイテムの数がゼロであることを示します。
* トランザクション内の全ての<code>txins</code>がwitnessデータに関連付けられていない場合は、トランザクションは<code>marker</code>や<code>flag</code>、<code>witness</code>が無いオリジナルのトランザクションフォーマットでシリアライズされなければなりません。例えば、トランザクションの<code>txins</code>が全てsegwitのUTXOではない場合、そのトランザクションはオリジナルのトランザクションフォーマットでシリアライズされなければなりません。（コインベーストランザクションを除く）
* トランザクションのシリアライゼーションの例はBIP143のexampleに記載されています。ウォレットの開発者は実装が新しいトランザクションフォーマットを正しくパースできているかそのexampleでテストすることができます。

#### Transaction ID {#transaction-id}

* segwitでは、トランザクションは2つのIDを持ちます。
* <code>txid</code>の定義は変わらず、オリジナルのシリアライゼーションフォーマットでダブルSHA256した値です。
* 新しく<code>wtxid</code>が定義され、これはwitnessデータを含む新しいトランザクションフォーマットをダブルSHA256した値になります。
* トランザクションがwitnessデータを持たない場合、<code>wtxid</code>と<code>txid</code>は同じ値になります。
* <code>txid</code>は引き続きトランザクションのプライマリ識別子として使われます：
    * <code>txin</code>内で前の出力を参照する際に使用されます。
    * ウォレットやサービスが現在トランザクションを識別するのに<code>txid</code>を使用している場合、アップグレード後も同様に使い続けることができます。

#### P2SH-P2WPKHの署名の作成と検証 {#signature-generation-and-verification-for-p2sh-p2wpkh}

* 非segwitのUTXOを使用する場合は、署名生成のアルゴリズムは今までと変わりません。
* P2SH-P2WPKHのUTXOを使用する場合：
    * <code>scriptSig</code>には<code>redeemScript</code>しかプッシュしてはいけません。
    * 対応するwitnessフィールドには署名がプッシュされ、その次に公開鍵がプッシュされなければなりません。
    * segwit scriptのための新しい署名生成アルゴリズムが[BIP143][]に定義されています。開発者は注意深くその指示に従い、[BIP143](https://github.com/bitcoin/bips/blob/master/bip-0143.mediawiki#P2SHP2WPKH)のexampleを参考に<code>sighash</code>を再現できるか確認する必要があります。
    * [BIP143][]の署名生成アルゴリズムは、使用される入力のコインの量をカバーしており、軽量ウォレットやハードウェアウォレットの設計のエアギャップを簡素化します。
    * P2SH-P2WPKHの場合、<code>scriptCode</code>は先頭のバイトを含めて常に26バイトで、<code>0x1976a914{20-byte keyhash}88ac</code>となります。<code>redeemScript</code>でも<code>scriptPubKey</code>でもないので注意してください。
    * [Example](http://n.bitcoin.ninja/checktx?txid=8139979112e894a14f8370438a471d23984061ff83a9eba0bc7a34433327ec21)
    
#### ネットワークサービス（オプション） {#network-services}

* ウォレットがピアツーピアネットワークを介してトランザクションを送受信する場合、ネットワークサービスが必要です。
* Segwit互換ノードは、service bit: <code>NODE_WITNESS = (1 << 3)</code>を使ってwitnessを提供できることを宣言します。
* witnessデータを持たないトランザクション（＝オリジナルフォーマットでシリアライズされたトランザクション）は、<code>NODE_WITNESS</code>のサポートの有無に関わらず別のノードに送信できます。
* segwitのUTXOを使用するトランザクション（＝新しいシリアライゼーションフォーマットでシリアライズされたトランザクション）は、<code>NODE_WITNESS</code>をサポートしているノードにのみ送信できます。
* segiwtのUTXOを使用しているがwitnessデータが取り除かれたトランザクション（＝オリジナルのフォーマットでシリアライズされたトランザクション）が、<code>NODE_WITNESS</code>をサポートしないノードに送信される可能性があります。ただし、segwitのアクティベート後はそのようなトランザクションは無効となるためブロックにも受け入れられなくなります。
* ネットワークサービスの詳細については[BIP144][]に記載されています。

#### ユーザーのプライバシー {#user-privacy}

* segwitがアクティベートされて数日は、ネットワーク内に流れるsegwitトランザクション数は限られたものでしょう。
* その間segwitのトランザクションを使用するとBitcoinの追跡が容易になる場合があります。
* デフォルトのお釣りの出力にP2SH-P2WPKHを使用すると、プライバシーに影響を与える可能性があります。

#### トランザクション手数料の見積もり {#transaction-fee-estimation}

* トランザクションサイズに代わって、"virtual size" (<code>vsize</code>)と呼ばれる新しいメトリックが定義されました。
* トランザクションの<code>vsize</code>は、オリジナルシリアライゼーションでシリアライズしたトランザクションデータを3倍し、新しいシリアライゼーションフォーマットでシリアライズしたトランザクションデータをプラスし、その値を4で割って次の整数に切り上げた数値です。例えば新しいシリアライゼーションフォーマットで200バイトのトランザクションで、 <code>marker</code>と<code>flag</code>と<code>witness</code>が全部で99バイトの場合、それらを除き、<code>vsize</code> は (99 * 3 + 200) / 4 = 125(切り上げ)となります。
* 非segwitトランザクションの<code>vsize</code>は単純にそのトランザクションサイズと同じになります。
* トランザクション手数料は他のトランザクションと<code>vsize</code>を比較して見積もられます。（今までのサイズではなく）
* 開発者は手数料の見積もりで、4倍の間違いをしないよう気をつける必要があります。

#### 安全なアップグレード {#ugrade-safety}

* エンドユーザーはsegwitがネットワーク上で完全にアクティベートされるまで、 P2SH-P2WPKHまたはその他のsegwitアドレスを生成してはいけません。アクティベートの前にP2SH-P2WPKHまたはその他のsegwitアドレスを使用すると、資金を失う可能性があります。
* 同様にお釣り用の出力もアクティベート前にsegwitの出力に送ってはいけません。
* segwitのアクティベーションは[BIP9][]で定義されています。2016年11月15日から2017年11月15日（UTC）までの間、retargetサイクルの2016ブロックのうち1916ブロックがシグナリングされ準備ができた場合、segwitは次のサイクルの後のretargetサイクルでアクティベートされます。
* [BIP9][]のシグナルを判断できないウォレットは、segwitがアクティベートされるまで、segwitのアップグレードバージョンのウォレットをエンドユーザーにリリースすべきではありません。
* 一部のマイナーが新しいルールを正しく実行できない懸念がある場合、アップグレードしたウォレットのリリースは（全てではないにしろ）大半のマイナーが新しいルールに従っていることが確認できるまで延期される可能性があります。ルールに違反している場合、無効な孤立ブロックとして明確に確認できます。

#### 後方互換性 {#backward-compatibility}

* （1で始まる）P2PKHの支払いの送受信は、引き続きサポートする必要があります。

### 複雑なスクリプトのサポート {#complex-script-support}

ウォレットが単一署名以外のマルチシグのようなスクリプトタイプをサポートしている場合、以下の要件を満たさなければなりません。

#### P2SH-P2WSHアドレスの作成 {#creation-fo-p2sh-p2wsh-address}

* P2SH-P2WSHアドレスは、P2SHアドレスと互換性があり、任意の複雑なスクリプトを固定サイズのアドレスで表現することができます。
* 他のP2SHやP2SH-P2WPKHアドレスと同様、P2SH-P2WSHアドレスも3から始まります。これらは実際にUTXOを使用し<code>redeemScript</code>が公開されるまで区別できません。
* P2SH-P2WSHアドレスを作成するには：
    1. <code>witnessScript</code>と呼ばれるスクリプトを定義します。
    2. <code>witnessScript</code>のSHA256（<code>scripthash</code>）を計算します。この時ダブルSHA256やRIPEMD160(SHA256)でなく、単体のSHA256であることに注意してください。
    3. P2SHの<code>redeemScript</code>は常に34バイトです。<code>OP_0</code>から始まり、その後に<code>scripthash</code>がプッシュされます。（例： <code>0x0020{32-byte scripthash}</code>）
    4. 他のP2SHと同様、<code>scriptPubKey</code>は<code>OP_HASH160 hash160(redeemScript) OP_EQUAL</code>で、対応するアドレスは3から始まるP2SHアドレスです。
* スクリプトの制限事項
    * スクリプトの評価に失敗してはならず、評価後のスタックにはただ１つTRUEとなるスタックアイテムが残っている状態でないといけません。それ以外の場合、評価は失敗します。
    * P2SH-P2WSHスクリプト内の公開鍵は圧縮公開鍵でなければならず、それ以外の場合資金が永久に失われる可能性があります。
    * OP_IFやOP_NOTIFが使われている場合、その引数は空のvector（false）か<code>0x01</code>（true）でなければなりません。他の値を使うと資金が永久に失われる可能性があります。 ([BIP draft](https://lists.linuxfoundation.org/pipermail/bitcoin-dev/2016-August/013014.html))
    * OP_CHECKSIGやOP_CHECKMULTISIGがfailを返す場合、全ての署名は空のvectorでなければなりません。それ以外の場合、資金が永久に失われる可能性があります。  ([BIP146][])
    * <code>witnessScript</code>のデフォルトポリシー制限は3600バイトです。<code>witnessScript</code>を除いて、witnessスタックアイテムは最大100個で、それぞれ最大で80バイトです。これらの制限を超えるトランザクションは中継されずブロックにも入れられません。
    * スクリプトサイズの10000バイト制限や、201 <code>nOpCount</code>などの元々のコンセンサスの制約の多くは、そのままP2SH-P2WSHにも適用されます。
    * P2SHの520バイトのスクリプトサイズ制限はP2SH-P2WSHには適用されません。これは3600バイトのポリシー制限と10000バイトのコンセンサス制限に置き換えられます。
    
#### P2SH-P2WSHの署名の生成と検証 {#signature-generation-and-verification-for-p2sh-p2wsh}

* P2SH-P2WSHを使用するには：
    * <code>scriptSig</code>には、<code>redeemScript</code>しかプッシュしてはいけません。
    * 対応するwitnessフィールドの最後のwitnessアイテムは<code>witnessScript</code>でなければなりません。
    * 新しい[BIP143][]の署名生成アルゴリズムが適用されます：
        * OP_CODESEPARATORを使わない場合、<code>scriptCode</code>は<code>witnessScript</code>のサイズを示す<code>compactSize</code> integerとそれに続く<code>witnessScript</code>になります。例えば、スクリプトがOP_1 (<code>0x51</code>)だとすると、シリアライズした<code>scriptCode</code>は(<code>0x0151</code>)になります。
        * あまり一般的ではありませんがOP_CODESEPARATORを含むスクリプトについては、[BIP143][]を参照ください。
    * <code>witnessScript</code>より前のwitnessスタックアイテムはスクリプトを評価する際の入力スタックとして使われます。入力スタックはスクリプトとしては解釈されません。そのため、例えば大きなデータをプッシュするのに<code>0x4c</code> (OP_PUSHDATA1) を使う必要はありません。
    * 署名の生成とスタック上のシリアライズされた署名が正しいか検証するには、[BIP143][]のexampleを参考にテストしてください。
    * [Example](http://n.bitcoin.ninja/checktx?txid=954f43dbb30ad8024981c07d1f5eb6c9fd461e2cf1760dd1283f052af746fc88)

### 高度な設計 {#advanced-designs}

最初のsegwitサポートには以下の機能は必要ありません。

#### Native Pay-to-Witness-Public-Key-Hash (P2WPKH) {#native-pay-to-witness-public-key-hash}

* Native P2WPKHは22バイトの<code>scriptPubKey</code>です。この<code>scriptPubKey</code>は<code>OP_0</code>で始まり、その後に<code>keyhash</code>がプッシュされます。（例：<code>0x0014{20-byte keyhash}</code>）
* P2SH-P2WPKHと同様、<code>keyhash</code>は圧縮公開鍵をRIPEMD160(SHA256)した値です。
* Native P2WPKHを使用する際は、<code>scriptSig</code>は必ず空で、witnessスタックの形式と署名の生成ルールはP2SH-P2WPKHと同じです（圧縮公開鍵を使用する要件含め）。
* [Example](http://n.bitcoin.ninja/checktx?txid=d869f854e1f8788bcff294cc83b280942a8c728de71eb709a2c29d10bfe21b7c)

#### Native Pay-to-Witness-Script-Hash (P2WSH) {#native-pay-to-witness-script-hash}

* Native P2WSHは34バイトの<code>scriptPubKey</code>です。この<code>scriptPubKey</code>は<code>OP_0</code>から始まり、その後に<code>scripthash</code>がプッシュされます。（例：<code>0x0020{32-byte scripthash}</code>）
* P2SH-P2WSHと同様、<code>scripthash</code>は<code>witnessScript</code>をSHA256した値です。
* Native P2WSHを使用する際は、<code>scriptSig</code>は必ず空で、witnessスタックの形式と署名の生成ルールはP2SH-P2WSHと同じです（圧縮公開鍵を使用する要件含め）。
* [Example](http://n.bitcoin.ninja/checktx?txid=78457666f82c28aa37b74b506745a7c7684dc7842a52a457b09f09446721e11c)

#### Native P2WPKHやP2WSHはどうやって使うのか? {#why-and-how-to-use-native-p2wpkh-and-p2wsh}

* Native P2WPKHおよびP2WSHのアドレス形式はありません。BIP142は延期され全く違う方法で実装される可能性が高いです。別の試みの[BIP173][] (Bech32)が2017年6月時点で'Draft'ステータスです。
* P2SH版と比べ、Native版のトランザクションの<code>vsize</code>はほとんどの場合小さくなります。そのため手数料はより少額になります。
* Native P2WPKHとP2WSHはBIP70のPayment ProtocolのようにRAW<code>scriptPubKey</code>プロトコルを使用します。ただ支払人と受取人のプライバシーに影響する可能性があります（下記参照）。
* Native P2WPKHとP2WSHをデフォルトのお釣り用のアドレスとして使用することはできますが、これにより他の人々が容易にお釣りを識別することができます（下記参照）。
* Native P2WPKHやP2WSHの使用は当初は珍しく、プライバシーの問題を引き起こす可能性があります。

### スクリプトとトランザクションの例 {#scripts-and-transactions-example}

* [Examples of different witness transaction types and transaction validity checking tool](http://n.bitcoin.ninja/checktx)
* [BIP141][]
* [BIP143][]
* [Script tests](https://github.com/bitcoin/bitcoin/blob/master/src/test/data/script_tests.json)
* [Valid transaction tests](https://github.com/bitcoin/bitcoin/blob/master/src/test/data/tx_valid.json)
* [Invalid transaction tests](https://github.com/bitcoin/bitcoin/blob/master/src/test/data/tx_invalid.json)
