---
title: CVE-2024-52921 - 由於變異區塊導致的區塊傳播受阻
name: blog-disclose-mutated-blocks-hindering-propagation
id: zh_TW-blog-disclose-mutated-blocks-hindering-propagation
lang: zh_TW
type: advisory
layout: post

## If this is a new post, reset this counter to 1.
version: 1

## Only true if release announcement or security announcement. English posts only
announcement: 1

excerpt: >
    對等節點可以透過傳送變異區塊來阻礙區塊傳播。
---

在 Bitcoin Core v25.0 之前，傳送變異區塊的對等節點可以清除同樣向我們公告該區塊的其他對等節點的下載狀態，這會阻礙區塊傳播。

此問題被認為是**中等**嚴重性。

## 細節

Bitcoin Core 將區塊視為變異，例如，當標頭中的默克爾根或幣基交易中的見證承諾與區塊中的交易不匹配時。

在 Bitcoin Core v25.0 之前，對等節點可以透過傳送未請求的變異區塊來清除其他對等節點的區塊下載狀態。例如，這對緊湊區塊中繼是一個問題。在接收到緊湊區塊並在等待 `getblocktxn` 請求的回應以重建完整區塊時，接收變異區塊會讓 Bitcoin Core 忘記緊湊區塊重建狀態。在變異區塊之後到達的 `blocktxn` 回應無法用於重建區塊。這阻礙了區塊傳播。

這在 [#27608](https://github.com/bitcoin/bitcoin/pull/27608) 中得到修復，透過確保對等節點只能影響其自己的區塊下載狀態，而不能影響其他對等節點的下載狀態。

## 歸屬

感謝 Suhas Daftuar 注意到該問題並修復它。

## 時間表

- 2023-05-08 - 在 [#bitcoin-core-dev IRC 頻道](https://bitcoin-irc.chaincode.com/bitcoin-core-dev/2023-05-08)中首次報告變異區塊問題。
- 2023-05-10 - 修復程式被合併（[#27608](https://github.com/bitcoin/bitcoin/pull/27608)）
- 2023-05-25 - v25.0 發布
- 2024-10-09 - 公開披露

{% include references.md %}
