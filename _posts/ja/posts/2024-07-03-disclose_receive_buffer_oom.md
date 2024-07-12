---
title: CVE-2015-3641の開示
name: blog-disclose-receive-buffer-oom
id: ja-blog-disclose-receive-buffer-oom
lang: ja
permalink: /ja/2024/07/03/disclose_receive_buffer_oom/
type: advisory
layout: post

## If this is a new post, reset this counter to 1.
version: 1

## Only true if release announcement or security annoucement. English posts only
announcement: 0

excerpt: >
  攻撃者が大規模な不完全なメッセージを送信すると、メモリ使用率が高くなる可能性がありました。
  修正は、2015年4月27日に、Bitcoin Core 0.10.1 でリリースされました。
---

ノードはメッセージの受信時に大きなバッファの割り当てを強制される可能性があり、これを利用してOOMによりリモートでノードがクラッシュする可能性がありました。

この問題の重大度は**中**です。

## 詳細 {#details}

制限が厳しくない場合、受信メッセージのサイズはシリアライズされたメッセージの最大サイズである32MiBに制限されます。
攻撃者は、ノードに接続毎にこの量のRAMを割り当てるよう強制することができ、OOMにつながる可能性があります。

[PR #5843](https://github.com/bitcoin/bitcoin/pull/5843)では、ペイロードを受信する前のP2Pメッセージのサイズが削減されました。
これにより、悪意のあるピアによって発生するピア毎の受信バッファのメモリサイズが削減されます。
PRでは、この数値が32MiBから2MiBに削減されましたが、その後、Segwit BIP144の変更の一環として4MBまで戻されました。

## 貢献 {#attribution}

bitcointalkのユーザーEvil-KnievelによってGreg Maxwellに報告され、Pieter Wuilleによって修正されました。

## タイムライン {#timeline}

- 2015-02-05 Evil-Knievelがbitcointalkのプライベートメッセージを通じてGreg Maxwellに脆弱性を報告
- 2015-??-?? `CVE-2015-3641`が登録される
- 2015-03-01 修正のための[PR #5843](https://github.com/bitcoin/bitcoin/pull/5843)が公開される
- 2015-03-06 PR #5843がマージされる
- 2015-03-09 修正がバージョン0.10.1にバックポートされる
- 2015-04-27 Bitcoin Core バージョン [0.10.1](https://lists.linuxfoundation.org/pipermail/bitcoin-dev/2015-April/007828.html)が修正と共にリリースされる
- 2015-06-25 開示の[事前発表](https://lists.linuxfoundation.org/pipermail/bitcoin-dev/2015-June/009135.html).
- 2015-07-07 開示が[延期される](https://lists.linuxfoundation.org/pipermail/bitcoin-dev/2015-July/009362.html).
- 2016-08-23 脆弱性のある最後のバージョンのBitcoin Core (0.10.x) がEOLになる
- 2024-07-03 公開

{% include references.md %}
