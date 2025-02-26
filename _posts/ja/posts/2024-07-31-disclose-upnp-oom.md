---
title: CVE-2024-52917 - miniupnp依存関係における無限ループバグの影響
name: blog-disclose-miniupnp-bug-impact
id: ja-blog-disclose-miniupnp-bug-impact
lang: ja
permalink: /ja/2024/07/31/disclose-upnp-oom/
type: advisory
layout: post

## If this is a new post, reset this counter to 1.
version: 1

## Only true if release announcement or security annoucement. English posts only
announcement: 0

excerpt: >
  ローカルネットワーク上の悪意あるUPnPデバイスによってノードがクラッシュする可能性がありました。
  修正は、2021年9月14日に Bitcoin Core v22.0でリリースされました。
---

Bitcoin Coreのminiupnp依存関係における無限ループバグの開示。
修正は、2021年9月14日に Bitcoin Core v22.0でリリースされました。

この問題の重大度は**低**です。

## 詳細 {#details}

Bitcoin Coreで使用されるUPnPライブラリであるMiniupnpは、ネットワーク上のデバイスからランダムなデータを受信するまで、
検出を待機します。さらに新しいデバイス情報ごとにメモリを割り当てます。ローカルネットワーク上の攻撃者は、
UPnPデバイスを装い、Bitcoin Coreノードがメモリを使い果たすまで、肥大化したM-SEARCH応答を送り続ける可能性があります。

Miniupnpはデフォルトでオフになっているため、
<code>-miniupnp</code>オプションを使用して実行しているユーザーのみがこのバグの影響を受けます。

## 貢献 {#attribution}

miniupnpプロジェクトに無限ループのバグを報告してくれたRonald Huveneersと、
OOMをトリガーするPoCエクスプロイトと（修正を含む）依存関係を更新するプルリクエストを
Bitcoin Coreプロジェクトに報告してくれたMichael Ford (Fanquake)に感謝します。

## タイムライン {#timeline}

* 2020-09-17 - Ronald Huveneersによる無限ループバグのminiupnpへの最初の報告
* 2020-10-13 - Michael Fordによる最初のレポートがsecurity@bitcoincore.orgに送信される
* 2021-03-23 - 修正がマージされる(https://github.com/bitcoin/bitcoin/pull/20421)
* 2021-09-13 - v22.0がリリースされる
* 2024-07-31 - 公開

{% include references.md %}
