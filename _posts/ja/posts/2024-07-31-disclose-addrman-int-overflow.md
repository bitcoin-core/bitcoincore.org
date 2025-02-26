---
title: CVE-2024-52919 - addrメッセージスパムによるリモートクラッシュ
name: blog-disclose-addrman-idcount-in-overflow
id: ja-blog-disclose-addrman-idcount-in-overflow
lang: ja
permalink: /ja/2024/07/31/disclose-addrman-int-overflow/
type: advisory
layout: post

## If this is a new post, reset this counter to 1.
version: 1

## Only true if release announcement or security annoucement. English posts only
announcement: 0

excerpt: >
  ノードにaddrメッセージスパムが送信され、ノードをクラッシュさせるために使用される可能性がありました。
  修正は、2021年9月14日に Bitcoin Core v22.0でリリースされました。
---

アサーションクラッシュを引き起こす整数オーバーフローバグの詳細を公開しました。
このバグの修正は、2021年9月14日に Bitcoin Core バージョン v22.0でリリースされました。

この問題の重大度は**高**です。

## 詳細 {#details}

`CAddrMan`には、32ビットの`nIdCount`フィールドがあり、addrmanへの挿入ごとに増加し、
新しいエントリーの識別子になります。（たとえば、addrメッセージをスパム送信するなど）
被害者に2<sup>32</sup>個のエントリーを挿入させると、この識別子がオーバーフローし、アサーションクラッシュが発生します。

## 貢献 {#attribution}

Eugene Siegelによる脆弱性の発見と開示と、Pieter Wuilleによる以下の修正に感謝します。
https://github.com/bitcoin/bitcoin/pull/22387

## タイムライン {#timeline}

* 2021-06-21 - Eugene Siegelによる最初のレポートがsecurity@bitcoincore.orgに送信される
* 2021-07-19 - 修正がマージされる(https://github.com/bitcoin/bitcoin/pull/22387)
* 2021-09-13 - v22.0がリリースされる
* 2024-07-31 - 公開

{% include references.md %}
