---
title: 披露由於 inv-to-send 集合過大導致的 DoS
name: blog-disclose-large-inv-to-send
id: zh_TW-blog-disclose-large-inv-to-send
lang: zh_TW
type: advisory
layout: post

## If this is a new post, reset this counter to 1.
version: 1

## Only true if release announcement or security announcement. English posts only
announcement: 1

excerpt: >
    inv-to-send 集合可能會增長得太大，以至於排序這些集合所花費的時間會影響節點與其對等節點通訊的能力。
---

在 Bitcoin Core v25.0 之前，每個對等節點的 `m_tx_inventory_to_send` 集合可能會增長得太大，以至於在建構庫存訊息時排序這些集合會影響節點與其對等節點通訊的能力。2023 年 5 月初的網路狀況觸發了此 DoS，並影響了區塊和交易傳播。

此問題被認為是**中等**嚴重性。

## 細節

作為交易中繼的一部分，Bitcoin Core 維護每個對等節點的 `m_tx_inventory_to_send` 集合，其中包含應向對等節點公告的交易。在為對等節點建構庫存訊息時，該集合按交易依賴關係和費率排序，以優先處理高費率交易並避免洩漏節點了解交易的順序。在 Bitcoin Core v25.0 之前，在建構庫存訊息時，相關的（仍在記憶池中、尚未被對等節點向我們公告、高於費用篩選器）交易以每秒 7 個交易的速率被清空。

2023 年 5 月初，網路活動增加導致集合增長速度快於清空速度，導致在 P2P 通訊執行緒中花費大量時間排序集合。此外，只監聽交易公告但從不自己公告任何交易的對等節點（通常稱為「間諜節點」），透過擁有巨大的集合（包含它們已經知道的交易）而放大了這一點，這需要很長時間來排序。觀察到排序幾乎佔用了 P2P 通訊執行緒中花費的全部時間，這顯著影響了區塊和交易傳播以及保持與對等節點的連接。

這在 [#27610](https://github.com/bitcoin/bitcoin/pull/27610) 中得到修復，透過 1) 更早地刪除不再在記憶池中的交易，以及 2) 根據集合大小動態增加集合清空速率。

## 歸屬

感謝 Anthony Towns 修復該問題，以及 b10c 最初報告並將問題縮小到緩慢的 inv-to-send 排序。

## 時間表

- 2023-05-02 - 問題首次被觀察到並報告
- 2023-05-11 - 修復程式被合併（[#27610](https://github.com/bitcoin/bitcoin/pull/27610)）
- 2023-05-25 - v25.0 發布
- 2024-10-09 - 公開披露

{% include references.md %}
