---
title: CVE-2024-52916 - 低難易度ヘッダーを使用したメモリDoS
name: blog-disclose-header-spam-checkpoint-bypass
id: ja-blog-disclose-header-spam-checkpoint-bypass
lang: ja
permalink: /ja/2024/07/03/disclose-header-spam/
type: advisory
layout: post

## If this is a new post, reset this counter to 1.
version: 1

## Only true if release announcement or security annoucement. English posts only
announcement: 0

excerpt: >
  ノードに低難易度のヘッダーがスパム送信され、それによりクラッシュする可能性がありました。
  修正は、2017年9月14日に、Bitcoin Core 0.15.0 でリリースされました。
---

Bitcoin Core 0.12.0からBitcoin Core 0.15.0未満で、ノードに最小難易度のヘッダーがスパム送信される可能性があり、これを利用してOOMによりクラッシュする可能性がありました。

この問題の重大度は**中**です。

## 詳細 {#details}

[ヘッダーの事前同期](https://github.com/bitcoin/bitcoin/pull/25717)が導入される前は、ノードは難易度の低いヘッダーによるスパムを避けるため、
チェックポイントのみに依存していました。

Bitcoin Core 0.12.0では、最後のチェックポイントの高さより前にヘッダーが分岐した場合のチェックが、
`mapBlockIndex`にヘッダーを格納した後に移動されました。これにより、攻撃者は、
親がジェネシスブロックであるヘッダー（作成に必要な難易度は1）をスパムすることで、
そのようなブロックがチェックポイントロジックをバイパスするため、マップを無制限に拡大させることができました。

## 貢献 {#attribution}

このバグを発見し責任をもって開示したCory Fieldsに感謝します。

## タイムライン {#timeline}

- 2017-08-08 Cory Fieldsがバグを非公開で報告
- 2017-08-11 Pieter Wuilleが修正のために[PR #11028](https://github.com/bitcoin/bitcoin/pull/11028)を公開
- 2017-08-14 PR #11028がマージされる
- 2017-09-14 Bitcoin Core バージョン 0.15.0が修正と共にリリースされる
- 2018-10-03 脆弱性のある最後のバージョンのBitcoin Core (0.14.3)がEOLになる
- 2024-07-03 公開

{% include references.md %}
