---
title: CVE-2024-52918 - 悪意あるBIP72 URIを使用したクラッシュ
name: blog-disclose-bip70-crash
id: ja-blog-disclose-bip70-crash
lang: ja
permalink: /ja/2024/07/03/disclose-bip70-crash/
type: advisory
layout: post

## If this is a new post, reset this counter to 1.
version: 2

## Only true if release announcement or security annoucement. English posts only
announcement: 0

excerpt: >
  Bitcoin-QtのBIP70の実装は、BIP72 URIを開くと密かにクラッシュする可能性がありました。
  修正は、2020年6月3日に、Bitcoin Core 0.20.0でリリースされました。
---

Bitcoin-Qtは、[BIP72](https://github.com/bitcoin/bips/blob/master/bip-0072.mediawiki) URIを開く際にクラッシュする可能性がありました。

この問題の重大度は**中**です。

## 詳細 {#details}

[BIP72](https://github.com/bitcoin/bips/blob/master/bip-0072.mediawiki)は、BIP21 URIスキームを拡張し、支払い要求をフェッチするためのパラメーター`r`を保持します。
攻撃者は、パラメーター`r`に含まれるURLに単に非常に大きなファイルを指定するだけで、
Bitcoin-Qtは十分なメモリを割り当てようとし、クラッシュします。

被害者は、騙されて不正な支払い要求を開いてしまう可能性があります。大きなダウンロードはバックグラウンドで行われ、
アプリケーションのメモリが不足するまでGUIにはほとんど何も出力されません。

## 貢献 {#attribution}

責任を持って問題を開示し、PoCを提供したMichael Ford (Fanquake)に感謝します。

## タイムライン {#timeline}

- 2019-08-12 Michael FordがCory FieldsとWladimir Van Der Laanにバグを報告
- 2019-10-16 Michael FordはBIP70のサポートを完全に廃止するためにPR[#17165](https://github.com/bitcoin/bitcoin/pull/17165)を公開
- 2019-10-26 MichaelのPRがBitcoin Coreにマージされる
- 2020-06-03 Bitcoin Core バージョン 0.20.0がリリースされる
- 2021-09-13 脆弱性のある最後のBitcoin Coreのバージョン (0.19.x)がEOLになる
- 2024-07-03 公開

{% include references.md %}
