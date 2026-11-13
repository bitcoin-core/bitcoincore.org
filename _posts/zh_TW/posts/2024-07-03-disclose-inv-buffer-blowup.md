---
title: CVE-2024-52915 - 使用巨大 INV 訊息的記憶體 DoS
name: blog-disclose-inv-buffer-blowup
id: zh_TW-blog-disclose-inv-buffer-blowup
lang: zh_TW
type: advisory
layout: post

## If this is a new post, reset this counter to 1.
version: 1

## Only true if release announcement or security annoucement. English posts only
announcement: 1

excerpt: >
  節點會為每個傳送惡意 `INV` 訊息的攻擊者分配最多 50 MB 的記憶體。修復程式於 2020 年 6 月 3 日在 Bitcoin Core 0.20.0 中發布。
---

節點可能被迫在接收到特別建構的 `INV` 訊息時分配大量記憶體。對於可用記憶體很少或連線數量很大的節點來說，這是一個特別的問題。

此問題被認為是**中等**嚴重性。

## 細節

填充 50,000 個區塊項目的 `INV` 訊息可能導致在單個 `ProcessMessages()` 呼叫中傳送 50,000 個 `getheaders` 回應。每個回應包含一個定位器，約 1 kB。所有這些都會立即放入傳送緩衝區。攻擊者可以只是拒絕接收資料以防止 50 MB 緩衝區排空。

John Newbery 開啟 [PR #18962](https://github.com/bitcoin/bitcoin/pull/18962) 以修復此問題，假託從每個接收的 `INV` 傳送單個 `GETHEADERS` 可以獲得頻寬增益。

## 歸屬

感謝 John Newbery 發現此錯誤、負責任地披露它並修復它。

## 時間表

- 2020-05-08 John Newbery 透過電子郵件報告他的發現
- 2020-05-12 John Newbery 開啟 #18962
- 2020-05-14 #18962 被合併
- 2020-06-03 Bitcoin Core 版本 0.20.0 發布，包含修復程式
- 2021-09-13 最後一個易受影響的 Bitcoin Core 版本（0.19.x）終止支援
- 2024-07-03 公開披露。

{% include references.md %}
