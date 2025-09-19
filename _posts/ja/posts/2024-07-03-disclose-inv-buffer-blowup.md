---
title: CVE-2024-52915 - 巨大なINVメッセージを使用したメモリDoS
name: blog-disclose-inv-buffer-blowup
id: ja-blog-disclose-inv-buffer-blowup
lang: ja
permalink: /ja/2024/07/03/disclose-inv-buffer-blowup/
type: advisory
layout: post

## If this is a new post, reset this counter to 1.
version: 1

## Only true if release announcement or security annoucement. English posts only
announcement: 0

excerpt: >
  ノードは、悪意ある`INV`メッセージを送信する攻撃者毎に最大50MBのメモリを割り当てます。
  修正は、2020年6月3日に、Bitcoin Core 0.20.0 でリリースされました。
---

特別に細工された`INV`メッセージを受信すると、ノードは大量のメモリを強制的に割り当てざるを得なくなる可能性がありました。
これは、使用可能なメモリが少ないノードや接続数の多いノードで特に問題になります。

この問題の重大度は**中**です。

## 詳細 {#details}

50,000個のブロック項目で満たされた`INV`メッセージは、1回の`ProcessMessages()`呼び出しで、
50,000個の`getheaders`応答を送信する可能性があります。各応答は、ロケーターを含む約1kBです。
このすべてが一度に送信バッファに入れられます。
攻撃者はデータの受信を拒否するだけで、50MBのバッファが空になるのを防ぐことができます。

John Newberyは、受信した`INV`毎に1つの`GETHEADERS`を送信することによる帯域幅の増加を口実に、
この問題を修正するための[PR #18962](https://github.com/bitcoin/bitcoin/pull/18962)を公開しました。

## 貢献 {#attribution}

このバグを発見し、責任を持って開示し修正したJohn Newberyに感謝します。

## タイムライン {#timeline}

- 2020-05-08 John Newberyが発見した内容をメールで報告
- 2020-05-12 John Newberyが#18962を公開
- 2020-05-14 #18962がマージされる
- 2020-06-03 Bitcoin Core バージョン 0.20.0が修正と共にリリースされる
- 2021-09-13 脆弱性のある最後のバージョンのBitcoin Core(0.19.x)がEOLになる
- 2024-07-03 公開

{% include references.md %}
