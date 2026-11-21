---
title: CVE-2024-52917 - miniupnp 依賴項中的無限迴圈錯誤
name: blog-disclose-miniupnp-bug-impact
id: zh_TW-blog-disclose-miniupnp-bug-impact
lang: zh_TW
type: advisory
layout: post

## If this is a new post, reset this counter to 1.
version: 1

## Only true if release announcement or security annoucement. English posts only
announcement: 1

excerpt: >
  節點可能被本機網路上的惡意 UPnP 裝置崩潰。修復程式於 2021 年 9 月 14 日在 Bitcoin Core v22.0 中發布。
---

披露 miniupnp 依賴項中無限迴圈錯誤對 Bitcoin Core 的影響，修復程式於 2021 年 9 月 14 日在 Bitcoin Core 版本 v22.0 中發布。

此問題被認為是**低**嚴重性。

## 細節

Bitcoin Core 使用的 UPnP 程式庫 Miniupnp，只要從網路上的裝置接收到隨機資料，就會在發現時一直等待。此外，它會為每個新裝置資訊分配記憶體。本機網路上的攻擊者可以假裝是 UPnP 裝置，並不斷向 Bitcoin Core 節點傳送膨脹的 M-SEARCH 回覆，直到其記憶體耗盡。

只有使用 <code>-miniupnp</code> 選項執行的使用者才會受到此錯誤的影響，因為 Miniupnp 預設是關閉的。

## 歸屬

感謝 Ronald Huveneers 向 miniupnp 專案報告無限迴圈錯誤，以及 Michael Ford (Fanquake) 向 Bitcoin Core 專案報告並提供 PoC 漏洞利用以觸發 OOM 以及提升依賴項（包含修復程式）的拉取請求。

## 時間表

* 2020-09-17 - Ronald Huveneers 向 miniupnp 初始報告無限迴圈錯誤
* 2020-10-13 - Michael Ford 向 security@bitcoincore.org 傳送初始報告
* 2021-03-23 - 修復程式被合併 (https://github.com/bitcoin/bitcoin/pull/20421)
* 2021-09-13 - v22.0 發布
* 2024-07-31 - 公開披露

{% include references.md %}
