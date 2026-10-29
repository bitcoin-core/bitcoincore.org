---
title: Linux以外のシステムにおけるTCP緊急データによる無制限なメモリ増加
name: blog-disclose-tcp-oom
id: ja-blog-disclose-tcp-oom
lang: ja
permalink: /ja/2026/10/28/disclose-tcp-oom/
type: advisory
layout: post

## If this is a new post, reset this counter to 1.
version: 1

## Only true if release announcement or security annoucement. English posts only
announcement: 0

excerpt: >
  Linux以外のプラットフォームでは、悪意のあるピアがTCPの緊急（帯域外）データを送信することで、ノードのメモリを枯渇させることができました。
---

Linux以外のプラットフォームでは、悪意のあるピアがTCPの緊急（帯域外）データを送信することで、
ノードのメモリを枯渇させることができました。
この問題の修正は、Bitcoin Coreバージョン29.5、30.4、31.2および32.1でリリースされました。

この問題の重大度は **中** と見なされており、**容易に発見可能**な脆弱性に分類されています。

## 詳細 {#details}

Bitcoin Coreは、Linux以外のプラットフォームにおいて、
ソケットイベントの監視に（`select()`ベースの）別のバックエンドを使用しています
（Linuxでは`poll()`を使用します）。この問題は`select()`を使用するコードパスにのみ影響します。
したがって、**Linux**上のBitcoin Coreユーザーはこの問題の**影響を受けません**。

Bitcoin Coreは、キューが一定のサイズを超えた場合に読み取りを一時的に停止することで、
ピアごとに処理待ちとしてキューに入れるデータ量を制限しています。しかし、
`select()`のソケット監視バックエンドは、読み取りが停止されているかどうかにかかわらず、
ソケットの例外的な状態を監視していました。この状態が発生するとソケットの読み取りが実行され、
そのピアの処理待ちデータが無制限にキューに追加される可能性がありました。

この状態はTCPの緊急データを送信することで発生させることができます。これにより、
一時停止中のピアはバックプレッシャー（負荷制御）の仕組みを迂回し、無制限の量のデータを処理待ちとしてキューに入れることができます。
これは実際には、悪意のあるピアが接続先のWindows、macOS、BSDのノードをいとも簡単にクラッシュさせられることを意味します。

## 貢献 {#attribution}

この問題は、SpiralにおけるElias Rohrerの[Project Loupe](https://github.com/project-loupe/loupe)の一環として実施された、
LLMによる自動スキャンによって発見されました。Loupeの発見はGregory Sandersによって検証され、Bitcoin Coreチームに報告されました。

## タイムライン {#timeline}

- 2026-10-23 Project Loupeが問題を特定しGreg Sandersに通知。Greg Sandersは同日中に報告を検証し、
  Bitcoin Coreセキュリティチームに報告
- 2026-10-26 Antoine Poinsotがこの問題を修正するPR [#36665](https://github.com/bitcoin/bitcoin/pull/36665)をオープン
- 2026-10-26 PR #36665がマージされ、バージョン29、30、31および32にバックポートされる
- 2026-10-27 Bitcoin Core 31.2および32.1が修正と共にリリースされる
- 2026-10-28 Bitcoin Core 29.5および30.4が修正と共にリリースされる
- 2026-10-28 開示

{% include references.md %}
