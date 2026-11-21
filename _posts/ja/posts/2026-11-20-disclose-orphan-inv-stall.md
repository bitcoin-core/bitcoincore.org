---
title: CVE-2026-117348 - 重複したオーファンのアナウンスによるノードの停止
name: blog-disclose-orphan-inv-stall
id: ja-blog-disclose-orphan-inv-stall
lang: ja
permalink: /ja/2026/11/20/disclose-orphan-inv-stall/
type: advisory
layout: post
redirect_from:
  - /ja/2026/11/20/disclose-cve-2026-117348

## If this is a new post, reset this counter to 1.
version: 1

## Only true if release announcement or security annoucement. English posts only
announcement: 0

excerpt: >
  ピアは、欠落している親を多数持つオーファントランザクションを繰り返しアナウンスすることで、ノードのブロックおよびトランザクションの処理を停止させることができました。修正はBitcoin Coreバージョン30.5、31.3および32.1でリリースされました。
---

ピアは、欠落している親を多数持つオーファントランザクションを繰り返しアナウンスすることで、
ノードのブロックおよびトランザクションの処理を一時的に停止させることができました。
修正はBitcoin Coreバージョン30.5、31.3および32.1でリリースされました。
影響を受けるバージョンは29.0〜29.5、30.0〜30.4および31.0〜31.2です。
29.xシリーズはEoLを迎えているため、修正は提供されません。

この問題の重大度は **中** と見なされており、**容易に発見可能**な脆弱性に分類されています。

## 詳細 {#details}

Bitcoin Coreは、親をまだ受信していないトランザクション（「オーファン」）を専用の領域（オーファネージ）に保持し、
オーファンをアナウンスしたピアに対して、欠落している親を要求します。バージョン29.0以降、
ピアがwtxidでトランザクションをアナウンスし、そのトランザクションがオーファネージ内にある場合、
Bitcoin Coreはそのオーファンの親のリストを作成し、それぞれをmempoolおよび最近確認したトランザクションのフィルターと照合して、
そのピアに親を要求するかどうかを判断します。

1つの`inv`メッセージには最大50,000個のエントリーを含めることができ、重複したエントリーはスキップされていませんでした。
攻撃者は、数千個もの親が欠けているオーファンをノードに送信し、その後、そのwtxidを50,000回繰り返した`inv`メッセージを送信することが可能でした。
このようなメッセージ1つにつき、ノードは十数秒以上停止し、その間、ブロックを含む他のピアからのメッセージを一切処理できなくなります。

この問題は、`inv`メッセージでアナウンスされた各トランザクションを1回だけ処理するようにすることで修正されました。

## 貢献 {#attribution}

この問題は、SpiralにおけるElias Rohrerの[Project Loupe](https://github.com/project-loupe/loupe)の一環として実施された、
LLMによる自動スキャンによって発見されました。Loupeの発見はGregory Sandersによって検証され、
Bitcoin Coreのセキュリティチームに報告されました。

## タイムライン {#timeline}

- 2026-07-08 - Project Loupeが問題を報告
- 2026-08-04 - Stutxoが独立して問題を報告
- 2026-08-04 - Gregory Sandersが修正を含むPR [#35878](https://github.com/bitcoin/bitcoin/pull/35878)をオープン
- 2026-08-06 - PR #35878がmasterにマージされる
- 2026-10-27 - Bitcoin Core 32.1が修正と共にリリースされる
- 2026-11-17 - 修正がPR [#36892](https://github.com/bitcoin/bitcoin/pull/36892)で31.xブランチにバックポートされる
- 2026-11-17 - 修正がPR [#36893](https://github.com/bitcoin/bitcoin/pull/36893)で30.xブランチにバックポートされる
- 2026-11-20 - Bitcoin Core 30.5および31.3が修正と共にリリースされる
- 2026-11-20 - 開示

{% include references.md %}
