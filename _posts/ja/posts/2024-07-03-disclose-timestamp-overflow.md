---
title: CVE-2024-52912 - タイムスタンプの調整によるネットワーク分割
name: blog-disclose-timestamp-overflow
id: ja-blog-disclose-timestamp-overflow
lang: ja
permalink: /ja/2024/07/03/disclose-timestamp-overflow/
type: advisory
layout: post

## If this is a new post, reset this counter to 1.
version: 1

## Only true if release announcement or security annoucement. English posts only
announcement: 0

excerpt: >
  最初の200個のピアによってノードが攻撃されると、ネットワークから切り離される可能性がありました。
  修正は、2021年1月15日に、Bitcoin Core バージョン 0.21.0 でリリースされました。
---

ネットワーク分割を引き起こす可能性のあった整数オーバーフローバグの詳細が開示され、
その修正は、2021年1月15日にBitcoin Coreバージョン0.21.0でリリースされました。

この問題の重大度は**中**です。

## 技術的な詳細 {#technical-details}

ネットワーク分割の脆弱性は、`version`メッセージの処理コードにおける2つの別々のバグによって発生しました:

* 新しく接続するピアの時間オフセットを計算する際の符号付き整数のオーバーフロー
* abs64ロジックのバグ (`abs64(std::numeric_limits<int64_t>::min()) ==
  std::numeric_limits<int64_t>::min()`)により、最大時間の調整制限がバイパスされる

この2つのバグにより、攻撃者は被害者の調整時間（`システム時間 + ネットワーク時間オフセット`）を強制的に歪ませることができ、
新しいブロックのタイムスタンプの日付が遠い未来すぎるために拒否させるようにすることができます。
この攻撃では、攻撃者が被害者に接続する最初の200ピアの1つであると想定されることに注意してください。
この初期の接続からの時間オフセットのみが調整時間に考慮されるためです。

## 貢献 {#attribution}

脆弱性の発見と最初の修正を提供した[practicalswift](https://github.com/practicalswift)と、
修正とリスクのあるコードの一般的なクリーンアップを担当したPieter Wuilleに感謝します。

## タイムライン {#timeline}

* 2020-10-10 最初の報告がsecurity@bitcoincore.orgに送信される
* 2020-10-13 修正がBitcoin Coreにマージされる(https://github.com/bitcoin/bitcoin/pull/20141)
* 2021-01-15 v0.21.0がリリースされる
* 2022-04-25 脆弱性のある最後のバージョンのBitcoin Core(0.20.x)がEOLになる
* 2024-07-03 公開

{% include references.md %}
