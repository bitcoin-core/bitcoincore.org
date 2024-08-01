---
title: CVE-2020-14198の開示
name: blog-disclose-unbounded-banlist
id: ja-blog-disclose-unbounded-banlist
lang: ja
permalink: /ja/2024/07/03/disclose-unbounded-banlist/
type: advisory
layout: post

## If this is a new post, reset this counter to 1.
version: 1

## Only true if release announcement or security annoucement. English posts only
announcement: 0

excerpt: >
  ノードは、多数の個別のIPによって攻撃されると、CPUおよびメモリのDoSを受ける可能性がありました。
  修正は、2020年8月1日に、Bitcoin Core 0.20.1 でリリースされました。
---

Bitcoin Coreは、禁止されたIPアドレスの無制限のリストを保持し、それに対して2次関数的な操作を行います。これによりOOMクラッシュやCPU DoSが発生する可能性がありました。

この問題の重大度は**高**です。

## 詳細 {#details}

Bitcoin Coreは禁止されたIPアドレスのリストを保持していました。このリストには制限がなく、
敵対者によって操作される可能性がありました。このリストに新しいエントリーを追加することは、
IPV6を考慮すると、攻撃者にとって安価でした。さらに、`GETADDR`メッセージを受信すると、
Bitcoin Coreは返されたすべてのアドレス（最大2500）に対して禁止リスト全体をスキャンします。

## 貢献 {#attribution}

Calin Culianuが最初に責任を持って開示し、Calin laterがその後[PRのコメント](https://github.com/bitcoin/bitcoin/pull/15617#issuecomment-640898523)でこのバグを公表しました。

同じ日、Bitcoin ABCのJason Coxは、Bitcoin Coreプロジェクトにメールを送信し、彼らも受け取った同じレポートを共有しました。

## タイムライン {#timeline}

- 2020-06-08 Calin CulianuがバグをBitcoin Coreプロジェクトに非公開で報告
- 2020-06-08 Jason Coxは、Bitcoin ABCに送られた（同じ）レポートをBitcoin Coreと非公開で共有
- 2020-06-08 Calin Culianuは、2次関数的な動作を導入した元のPRの脆弱性を公表
- 2020-06-09 Pieter Wuilleが、無制限のメモリ使用と2次関数的な動作の両方を修正するPR [#19219](https://github.com/bitcoin/bitcoin/pull/19219)を公開
- 2020-06-16 Luke Dashjrの要求により、この脆弱性にCVE-2020-14198が割り当てられる
- 2020-07-07 PieterのPRがマージされる
- 2020-08-01 Bitcoin Core 0.20.1が修正と共にリリースされる
- 2021-01-14 Bitcoin Core 0.21.0が修正と共にリリースされる
- 2022-04-25 脆弱性のある最後のバージョンのBitcoin Core (0.20.0) がEOLになる
- 2024-07-03 (公式に)公表

{% include references.md %}
