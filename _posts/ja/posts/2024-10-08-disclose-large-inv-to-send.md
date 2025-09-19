---
title: inv-to-sendの肥大化によるDoSの開示
name: blog-disclose-large-inv-to-send
id: ja-blog-disclose-large-inv-to-send
lang: ja
permalink: /ja/2024/10/08/disclose-large-inv-to-send/
type: advisory
layout: post

## If this is a new post, reset this counter to 1.
version: 1

## Only true if release announcement or security announcement. English posts only
announcement: 0

excerpt: >
    inv-to-sendのセットが大きくなりすぎると、セットをソートする時間がノードとピアの通信能力に影響する可能性がありました。
---

Bitcoin Core v25.0より前のバージョンでは、ピア毎の`m_tx_inventory_to_send`セットが大きくなりすぎて、
inventoryメッセージを構築する際にこれらのセットをソートすると、ノードがピアと通信する能力に影響を与える可能性がありました。
2023年5月上旬のネットワーク状況が、このDoSを引き起こし、ブロックとトランザクションの伝播に影響を与えました。

この問題の重大度は **中** です。

## 詳細 {#details}

トランザクションリレーの一部として、Bitcoin Coreはピアに通知すべきトランザクションを含む
`m_tx_inventory_to_send`のセットをピア毎に保持しています。
ピアに対してinventoryメッセージを作成する際、このセットはトランザクションの依存関係と手数料率によってソートされます。
手数料率の高いトランザクションを優先し、トランザクションについて学習した順序が漏れるのを防ぎます。
Bitcoin Core v25.0より前のバージョンでは、inventoryメッセージを構築する際、
関連する（まだmempoolにあり、ピアからまだ通知されておらず、手数料フィルターより上）トランザクションが
1秒あたり7トランザクションの割合で排出されていました。

2023年5月上旬には、ネットワーク活動の増加により、排出されるよりも早くセットが増加し、
P2P通信スレッドでこのセットをソートするのに多大な時間が費やされました。さらに、
トランザクションの通知を待ち受けるだけで自身からは通知しないピア（一般に「スパイノード」と呼ばれる）は、
ソートに時間がかかる巨大なセット（既に知っているトランザクションを含む）を持つことで、これを増幅しました。
ソート処理にP2P通信スレッドのほぼすべての時間が費やされ、ブロックやトランザクションの伝播、
ピアとの接続の維持に大きな影響を与えることが確認されました。

この問題は、[#27610](https://github.com/bitcoin/bitcoin/pull/27610)で(1) mempoolに存在しないトランザクションを早期に削除し、
(2) セットのサイズに応じて排出率を動的に増加させることで修正されました。

## 貢献 {#attribution}

修正に取り組んだAnthony Townsと、最初に問題を報告し、問題をinv-to-sendのソートの遅さに絞り込んだb10cに感謝します。

## タイムライン {#timeline}

- 2023-05-02 - 問題が最初に発見され報告される
- 2023-05-11 - 修正がマージされる（[#27610](https://github.com/bitcoin/bitcoin/pull/27610)）
- 2023-05-25 - v25.0がリリースされる
- 2024-10-09 - 公開

{% include references.md %}
