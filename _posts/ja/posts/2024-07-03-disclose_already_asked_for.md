---
title: CVE-2024-52913 - トランザクションの再要求処理による検閲
name: blog-disclose-already-asked-for
id: ja-blog-disclose-already-asked-for
lang: ja
permalink: /ja/2024/07/03/disclose_already_asked_for/
type: advisory
layout: post

## If this is a new post, reset this counter to 1.
version: 2

## Only true if release announcement or security annoucement. English posts only
announcement: 0

excerpt: >
  悪意あるピアによって、ノードが特定の未承認トランザクションを参照できないようにすることができました。
  修正は、2021年1月14日に、Bitcoin Core 0.21.0 でリリースされました。
---

攻撃者によりノードが特定の未承認トランザクションを参照できないようになる可能性がありました。

この問題の重大度は**中**です。

## 詳細 {#details}

この問題がPR 19988で修正される前は、トランザクションの`GETDATA`要求をスケジュールするのに「g_already_asked_for」メカニズムが使用されていました。
`SendMessages()`関数は、ピアによって最近アナウンスされたトランザクションについて`GETDATA`を送信し、
その要求が`g_already_asked_for`でいつ送信されたかを記憶していました。
しかし、この`g_already_asked_for`は、「limitedmap」データ構造でサイズが制限されているため、
50000エントリーに達すると最も古いエントリーを忘れてしまいます。これにより次の攻撃が可能になります:

* 攻撃者は、正当なトランザクションTを被害者に最初にアナウンスします。
* 被害者は、`GETDATA`を使用して攻撃者にTを要求します。
* 攻撃者は、被害者が他のピアにTを要求する時間近く（約60秒）まで`GETDATA`に応答しません。
* 次に攻撃者は、被害者に偽のアナウンスを巧妙に送りつけ、被害者の`g_already_asked_for`がTを排除するよう試みます。
* 攻撃者は被害者に再度Tをアナウンスします（`m_tx_process_time`のキューイングの仕組みにより、これは特に正確に時間を計る必要はありません）。
* 被害者者は、`g_already_asked_for`内にTが見つからないため、それを新しいアナウンスとして扱い、新しい`GETDATA`を攻撃者に送信します。
* 攻撃者は、再び`GETDATA`に応答しません。
* 同じように続けます。

このようにして、攻撃者は被害者が攻撃者以外にトランザクションを要求することを防ぐことができます。

## 貢献 {#attribution}

John Newberyによって責任を持って開示され、Amiti Uttarwarと彼によって発見されたとされています。

## タイムライン {#timeline}

- 2020-04-03 John Newberyが、Suhas Daftuarらにメールでバグを報告
- 2020-05-08 John Newberyが、バグを修正するアプローチを提案
- 2020-09-21 Pieter Wuilleが、このバグやその他のバグを修正するための包括的なアプローチとして[PR #19988](https://github.com/bitcoin/bitcoin/pull/19988)を公開
- 2020-10-14 PieterのPRがマージされる
- 2021-01-14 Bitcoin Core バージョン 0.21.0が修正と共にリリースされる
- 2022-04-25 脆弱性のある最後のバージョンのBitcoin Core (0.20.x) がEOLになる
- 2024-07-03 公開

{% include references.md %}
