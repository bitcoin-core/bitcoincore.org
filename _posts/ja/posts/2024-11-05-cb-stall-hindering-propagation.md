---
title: CVE-2024-52922 - ピアの停止によるブロック伝播の妨害
name: blog-disclose-stalling-peers-hindering-propagation
id: ja-blog-disclose-stalling-peers-hindering-propagation
lang: ja
permalink: /ja/2024/11/05/cb-stall-hindering-propagation/
type: advisory
layout: post

## If this is a new post, reset this counter to 1.
version: 1

## Only true if release announcement or security announcement. English posts only
announcement: 0

excerpt: >
    ピアは最初にブロックを通知し、その後ブロックの送信を単に保留することで、ブロックの伝播を妨げる可能性がありました。
---

Bitcoin Core v25.1より前のバージョンでは、攻撃者が最新のブロックをダウンロードさせないようにすることができます。

この問題の重大度は**中**です。

## 詳細 {#details}

headersまたはコンパクトブロックメッセージを介して新しいブロックの通知を受信すると、
配信ピアは受信ノードから完全なブロックか不足しているトランザクションの詳細のいずれかを要求されます。
通知したピアがP2Pプロトコルの要件どおりに応答しない場合、影響を受けるBitcoin Coreノードは、
最大10分間待ってからピアを切断し、別でブロックダウンロードを試行します。
攻撃者が複数の着信または送信接続を行うことができる場合、このプロセスが繰り返される可能性があります。

ブロック配信の遅延は、ネットワークの収束を遅くし、マイニング報酬の支払いを不公平にし、
応答性の問題が発生することで、ネットワークの劣化を引き起こす可能性があります。

この問題は、mempoolが比較的異質だった際の最近開示された他の問題（たとえば[inventoryの蓄積](https://bitcoincore.org/ja/2024/10/08/disclose-large-inv-to-send/)）によってさらに悪化し、
正直なピアによるコンパクトブロックの再構築ができませんでした。

Bitcoin Core v26.0で導入され、v25.1にバックポートされた[#27626](https://github.com/bitcoin/bitcoin/pull/27626)で
緩和策が導入されました。これは、最大3つの高帯域幅のコンパクトブロックピアに
同時にブロックを要求できるようにするもので、そのうち1つはアウトバウント接続である必要があります。

## 貢献 {#attribution}

Greg Sandersによって報告および修正されました。

## タイムライン {#timeline}

- 2023-05-08 - ユーザーが[#bitcoin-core-dev IRCチャンネル](https://bitcoin-irc.chaincode.com/bitcoin-core-dev/2023-05-08)でブロックのタイムアウトを報告
- 2023-05-09 - 問題を説明する最初のGithubの課題 https://github.com/bitcoin/bitcoin/issues/25258#issuecomment-1540028533
- 2023-05-11 - 緩和策のPRの公開 https://github.com/bitcoin/bitcoin/pull/27626
- 2023-05-24 - Bitcoin Core v26.0の前にPRがマージされる
- 2023-05-25 - Bitcoin Core v25.1へのバックポートがマージされる https://github.com/bitcoin/bitcoin/pull/27752
- 2023-10-19 - Bitcoin Core v25.1 リリース
- 2024-11-05 - 公開

{% include references.md %}
