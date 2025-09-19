---
title: CVE-2024-52914 - オーファンの処理による重大なDoS
name: blog-disclose-orphan-dos
id: ja-blog-disclose-orphan-dos
lang: ja
permalink: /ja/2024/07/03/disclose-orphan-dos/
type: advisory
layout: post

## If this is a new post, reset this counter to 1.
version: 1

## Only true if release announcement or security annoucement. English posts only
announcement: 0

excerpt: >
  特別に細工された未承認トランザクションを受信すると、ノードが数時間停止する可能性がありました。
  修正は、2019年5月18日に、Bitcoin Core 0.18.0 でリリースされました。
---

特別に細工された未承認トランザクションのオーファンを処理する際に、ノードが数時間停止する可能性がありました。

この問題の重大度は**高**です。

## 詳細 {#details}

トランザクションをmempoolに受け入れた後、ノードはオーファントランザクションのキャッシュを調べて、
新しく受け入れられたトランザクションによりオーファントランザクションが受け入れ可能になるかどうか調べます。
この探索は2次関数的で、新しく受け入れれたトランザクションの各アウトプットに対して、
（100件に制限されている）キャッシュされたオーファントランザクションをすべて調べます。
オーファントランザクションを無効でありながら検証にコストがかかるよう特別に細工すると、
ノードが数時間停止する可能性があります。

この停止は、（オーファンが有効であるかどうかにかかわらず）オーファンとの一致が見つかった際に
新しいメッセージを処理するためにオーファンの解決を中断することで、
Pieter Wuilleによって[PR #15644](https://github.com/bitcoin/bitcoin/pull/15644)で修正されました。

## 貢献 {#attribution}

このバグを責任を持って開示し、修正に関するフィードバックを提供してくれたsec.eineに感謝します。

## タイムライン {#timeline}

- 2019-03-19 sec.eineがメールでGreg Maxwellにこの問題を報告
- 2019-03-21 Greg Maxwellが提案されたパッチに関する情報を返信
- 2019-03-22 sec.eineがパッチについてフィードバックを提供（「堅実で、[..]注意を引かない」）
- 2019-03-22 Pieter WuilleがPR #15644を公開
- 2019-04-01 PR #15644がマージされる
- 2019-05-18 Bitcoin Core バージョン 0.18.0が修正と共にリリースされる
- 2020-07-22 PR review clubで問題が[部分的に開示される](https://bitcoincore.reviews/15644#l-285)
- 2020-08-01 脆弱性のある最後のバージョンのBitcoin Core version (0.17.x)がEOLになる
- 2024-07-03 公開

{% include references.md %}
