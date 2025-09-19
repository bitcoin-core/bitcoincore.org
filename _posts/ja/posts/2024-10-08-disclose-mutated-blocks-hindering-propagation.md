---
title: CVE-2024-52921 - 変異ブロックによるブロック伝播阻害
name: blog-disclose-mutated-blocks-hindering-propagation
id: ja-blog-disclose-mutated-blocks-hindering-propagation
lang: ja
permalink: /ja/2024/10/08/disclose-mutated-blocks-hindering-propagation/
type: advisory
layout: post

## If this is a new post, reset this counter to 1.
version: 1

## Only true if release announcement or security announcement. English posts only
announcement: 0

excerpt: >
    ピアは、変異ブロックを送信することで、ブロックの伝播を妨げる可能性がありました。
---

Bitcoin Core v25.0より前のバージョンでは、変異ブロックを送信するピアが、
同じブロックを通知した他のピアのダウンロード状態をクリアしてしまう可能性があり、
ブロックの伝播を妨げていました。

この問題の重大度は **中** です。

## 詳細 {#details}

Bitcoin Coreは、たとえばヘッダーのマークルルートや、
コインベーストランザクション内のwitness commitmentがブロック内のトランザクションと一致しない場合、
そのブロックを変異ブロックとして扱います。

Bitcoin Core v25.0より前のバージョンでは、ピアが要求されていない変異ブロックを送信することで、
他のピアのブロックダウンロード状態をクリアすることができました。
これは、たとえばコンパクトブロックのリレーでは問題でした。コンパクトブロックを受信した後、
完全なブロックを再構築するために`getblocktxn`要求に対する応答を待っている間に変異ブロックを受信すると、
Bitcoin Coreはコンパクトブロックの再構築状態を忘れてしまいます。変異ブロックの後に届いた`blocktxn`応答は、
ブロックの再構築には使用できませんでした。これはブロックの伝播を妨げました。

これは[#27608](https://github.com/bitcoin/bitcoin/pull/27608)で、
ピアが影響を与えるのは自分のブロックダウンロード状態のみで、
他のピアのダウンロード状態には影響を与えないようにすることで修正されました。

## 貢献 {#attribution}

この問題に気付き、修正に取り組んだSuhas Daftuarに感謝します。

## タイムライン {#timeline}

- 2023-05-08 - 変異ブロックに関する問題は[#bitcoin-core-dev IRC channel](https://bitcoin-irc.chaincode.com/bitcoin-core-dev/2023-05-08)で最初に報告される
- 2023-05-10 - 修正がマージされる（[#27608](https://github.com/bitcoin/bitcoin/pull/27608)）
- 2023-05-25 - v25.0がリリースされる
- 2024-10-09 - 公開

{% include references.md %}
