---
title: CVE-2024-52920 - 巨大なGETDATAメッセージを使用したDoS
name: blog-disclose-getdata-cpu
id: ja-blog-disclose-getdata-cpu
lang: ja
permalink: /ja/2024/07/03/disclose-getdata-cpu/
type: advisory
layout: post

## If this is a new post, reset this counter to 1.
version: 1

## Only true if release announcement or security annoucement. English posts only
announcement: 0

excerpt: >
  不正な`GETDATA`メッセージにより、受信ノードのCPU使用率が100%になる可能性がありました。
  修正は、2020年6月3日に、Bitcoin Core 0.20.0 でリリースされました。
---

不正な`GETDATA`が受信ノードで無限ループを引き起こし、このスレッドに割り当てられたCPUを100%使用し、
その接続でのこれ以上の進行ができなくなる可能性がありました。

この問題の重大度は**低**です。

## 詳細 {#details}

Bitcoin Core 0.20.0より前のバージョンでは、攻撃者（またはバグのあるクライアント）が
`GETDATA`メッセージを送信して、net_processingスレッドが100%で回り始め、
攻撃ピアのメッセージ処理が進捗しなくなる可能性がありました。
他のピアからのメッセージの処理は進むので、それ以上の影響は小さなCPU DoSに過ぎません（攻撃ピアの処理が進まないことは問題ではありません）。
また、ピア毎の長期メモリ使用量が1.5MB増加します。

John Newberyは、この問題を修正するために[PR #18808](https://github.com/bitcoin/bitcoin/pull/18808)を公開し、進捗がないことだけを開示しました。

## 貢献 {#attribution}

このバグを発見し、責任ある開示をし修正したJohn Newberyに感謝します。

## タイムライン {#timeline}

- 2020-04-29 John Newberyが#18808を公開
- 2020-05-08 John Newberyがメールで彼の発見を報告
- 2020-05-12 #18808がマージされる
- 2020-06-03 Bitcoin Core バージョン 0.20.0が修正と共にリリースされる
- 2021-09-13 脆弱性のある最後のバージョンのBitcoin Core(0.19.x)がEOLになる
- 2024-07-03 公開

{% include references.md %}
