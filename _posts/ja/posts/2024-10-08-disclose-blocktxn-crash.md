---
title: CVE-2024-35202の開示
name: blog-disclose-blocktxn-crash
id: ja-blog-disclose-blocktxn-crash
lang: ja
permalink: /ja/2024/10/08/disclose-blocktxn-crash/
type: advisory
layout: post

## If this is a new post, reset this counter to 1.
version: 1

## Only true if release announcement or security annoucement. English posts only
announcement: 0

excerpt: >
  攻撃者は、blocktxnメッセージ処理ロジックのアサーションのトリガーすることで、
  Bitcoin Coreノードをリモートでクラッシュさせる可能性がありました。
---

Bitcoin Core v25.0より前のバージョンでは、攻撃者はblocktxnメッセージ処理ロジックのアサーションをトリガーすることで、
Bitcoin Coreノードをリモートでクラッシュさせることができました。

この問題の重大度は**高**です。

## 詳細 {#details}

cmpctblockメッセージを介してブロックの通知を受信した際、Bitcoin Coreは、
自身のmempool内のトランザクションと、その他の利用可能なトランザクションを使用して、
通知されたブロックを再構築しようとします。トランザクションの欠如により再構築が失敗した場合は、
通知ピアに対してgetblocktxnメッセージでトランザクションを要求します。
その応答として、要求したトランザクションを含むblocktxnメッセージが返ってくるのを期待します。

コンパクトブロックプロトコルは、帯域幅を削減するために、短縮されたトランザクション識別子を使用します。
これらの短縮IDは、サイズが6 byteであるため、ブロックの再構築時に衝突が発生する可能性
（つまり、トランザクションAがトランザクションBと同じ短縮IDを持つこと）がわずかにあります。
衝突は、再構築されたトランザクションのセットから計算されたマークルルートが、
通知されたブロックのマークルルートと一致しないことで検出されます。
衝突は不意に発生する可能性があるため、ピアは罰せされるべきではなく、
衝突は完全なブロックを要求するようにフォールバックすることで処理されます。

Bitcoin Coreは、新しいコンパクトブロックを受信するたびに、
<code>PartiallyDownloadedBlock</code>のインスタンスを作成します。
欠如しているトランザクションを要求すると、このインスタンスは対応するblocktxnメッセージが処理されるまで保持されます。
blocktxnメッセージを受信すると、<code>PartiallyDownloadedBlock::FillBlock</code>が呼び出され、
完全なブロックの再構築を試行します。上記のような衝突が起きると完全なブロックを要求しますが、
<code>PartiallyDownloadedBlock</code>インスタンスと、
ベースとなるブロックの要求に関連するその他の状態はそのまま残されます。
これにより、同じブロックに対する２つめのblocktxnメッセージが処理され、
<code>FillBlock</code>が再び呼び出されるトリガーとなる余地が残ります。
これは、<code>FillBlock</code>が一度しか呼び出されないという仮定
（<code>assert</code>ステートメントとして記載されている）に違反し、ノードをクラッシュさせます。

衝突の処理ロジックは、ブロックのマークルルートにコミットされていないトランザクションをblocktxnメッセージに含めるだけで
簡単にトリガーできるため、攻撃者は衝突のトリガーに運は必要ありません。

## 貢献 {#attribution}

この脆弱性を発見、開示し、[#26898](https://github.com/bitcoin/bitcoin/pull/26898) で問題を修正した
Niklas Göggeに感謝します。

## タイムライン {#timeline}

* 2022-10-05 - Niklas GöggeがBitcoin Coreのセキュリティメーリングリストに問題を報告
* 2023-01-24 - 修正が含まれたPR #26898がマージされる
* 2023-05-25 - Bitcoin Core 25.0が修正と共にリリースされる
* 2024-10-09 - 公開

{% include references.md %}
