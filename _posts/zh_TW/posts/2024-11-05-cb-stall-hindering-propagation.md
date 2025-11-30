---
title: CVE-2024-52922 - 由於停滯對等節點導致的區塊傳播受阻
name: blog-disclose-stalling-peers-hindering-propagation
id: zh_TW-blog-disclose-stalling-peers-hindering-propagation
lang: zh_TW
type: advisory
layout: post

## If this is a new post, reset this counter to 1.
version: 1

## Only true if release announcement or security announcement. English posts only
announcement: 1

excerpt: >
    對等節點可以透過首先公告區塊然後簡單地扣留區塊來阻礙區塊傳播。
---

在 Bitcoin Core v25.1 之前，攻擊者可以導致節點不下載最新區塊。

此問題被認為是**中等**嚴重性。

## 細節

當透過標頭或緊湊區塊訊息接收到新的區塊公告時，接收節點會向傳遞的對等節點請求完整區塊或缺少的交易細節。如果公告對等節點隨後未按照對等節點對等節點協定的要求作出回應，受影響的 Bitcoin Core 節點將等待最多 10 分鐘，然後斷開對等節點連接並進行另一次區塊下載嘗試。如果攻擊者能夠建立多個傳入或傳出連接，則可以重複此過程。

延遲區塊傳遞可能會透過減慢網路收斂、使挖礦收益不太公平以及導致活性問題而導致網路降級。

此問題因最近披露的其他問題（例如[庫存堆積](https://bitcoincore.org/en/2024/10/08/disclose-large-inv-to-send/)）而進一步加劇，當記憶池相對異質時，不允許誠實對等節點機會性地重建緊湊區塊。

在 [#27626](https://github.com/bitcoin/bitcoin/pull/27626) 中引入了緩解措施，在 Bitcoin Core v26.0 中引入並向後移植到 v25.1。它確保可以同時從最多 3 個高頻寬緊湊區塊對等節點請求區塊，其中一個必須是傳出連接。

## 歸屬

由 Greg Sanders 報告和修復。

## 時間表

- 2023-05-08 - 使用者在 [#bitcoin-core-dev IRC 頻道](https://bitcoin-irc.chaincode.com/bitcoin-core-dev/2023-05-08)中報告區塊逾時
- 2023-05-09 - 第一個描述該問題的 github issue https://github.com/bitcoin/bitcoin/issues/25258#issuecomment-1540028533
- 2023-05-11 - 緩解 PR 開啟 https://github.com/bitcoin/bitcoin/pull/27626
- 2023-05-24 - PR 在 Bitcoin Core v26.0 之前合併
- 2023-05-25 - 向後移植到 Bitcoin Core v25.1 的 PR 合併 https://github.com/bitcoin/bitcoin/pull/27752
- 2023-10-19 - Bitcoin Core v25.1 發布
- 2024-11-05 - 公開披露

{% include references.md %}
