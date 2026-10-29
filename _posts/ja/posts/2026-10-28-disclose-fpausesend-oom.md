---
title: 送信一時停止中のBIP 31非対応ピアによる緩やかで無制限なメモリ増加
name: blog-disclose-fpausesend-oom
id: ja-blog-disclose-fpausesend-oom
lang: ja
permalink: /ja/2026/10/28/disclose-fpausesend-oom/
type: advisory
layout: post

## If this is a new post, reset this counter to 1.
version: 1

## Only true if release announcement or security annoucement. English posts only
announcement: 0

excerpt: >
  Bitcoin Coreは、一時停止中のピアを含め、すべてのピアのアナウンスをキューに入れます。BIP 31非対応のピアは永久に一時停止状態のままでいることができ、キューを無制限に増大させる可能性がありました。
---

Bitcoin Coreは、一時停止中のピアを含め、すべてのピアのアナウンスをキューに入れます。
BIP 31非対応のピアは永久に一時停止状態のままでいることができ、キューを無制限に増大させる可能性がありました。
この問題の修正は、Bitcoin Coreバージョン29.5、30.4、31.2および32.1でリリースされました。

この問題の重大度は **中** と見なされており、**容易に発見可能**な脆弱性に分類されています。

## 詳細 {#details}

Bitcoin Coreは、読み取りの遅いピアに対するバックプレッシャー（負荷制御）の仕組みを実装しています。
あるピアに送信するメッセージのキューが一定の上限を超えると、そのピアを送信に関して一時停止状態としてマークします。
しかし、その状態でも一部のアナウンス、特に頻繁に発生するトランザクションの`INV`はキューに追加され続けます。

通常これは問題になりません。（nonce付きの）pingメッセージに20分以内に応答しないピアは切断されるためです。
しかし、この切断処理はBIP 31のnonce付きpingメッセージをサポートしないピアに対しては有効になっていません。
そのため、悪意のあるピアはBIP 31をサポートしていないふりをし、メッセージの読み取りを戦略的に遅延させることで、
キューを際限なくゆっくりと増大させることができました。

メモリ使用量は、インバウンドピア1つあたり最大で毎時2MB、アウトバウンドピア1つあたり毎時5MBの割合で増加する可能性がありました。
キューに入れられるメッセージの大半はトランザクションのアナウンスであるため、
`-blocksonly`で動作しているノードへの影響ははるかに小さくなります。

## 貢献 {#attribution}

この問題は、SpiralにおけるElias Rohrerの[Project Loupe](https://github.com/project-loupe/loupe)の一環として実施された、
LLMによる自動スキャンによって発見されました。Loupeの発見はGregory Sandersによって検証され、Bitcoin Coreチームに報告されました。

また、L0rincも独自にこの問題を発見しています。

## タイムライン {#timeline}

- 2026-09-15 Project Loupeが問題を特定しGreg Sandersに通知。Greg Sandersは同日中に報告を検証し、Bitcoin Coreセキュリティチームに報告
- 2026-09-21 L0rincがGregory Sandersに非公開で通知し、既に報告済みの問題であるとの返答を受ける
- 2026-10-26 Antoine Poinsotが脆弱性を緩和するPR [#36667](https://github.com/bitcoin/bitcoin/pull/36667)をオープン
- 2026-10-26 PR #36667がマージされ、バージョン29、30、31および32にバックポートされる
- 2026-10-27 Bitcoin Core 31.2および32.1が修正と共にリリースされる
- 2026-10-28 Bitcoin Core 29.5および30.4が修正と共にリリースされる
- 2026-10-28 開示

{% include references.md %}
