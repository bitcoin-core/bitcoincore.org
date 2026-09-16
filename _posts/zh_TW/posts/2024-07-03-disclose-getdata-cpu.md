---
title: CVE-2024-52920 - 使用巨大 GETDATA 訊息的 DoS
name: blog-disclose-getdata-cpu
id: zh_TW-blog-disclose-getdata-cpu
lang: zh_TW
type: advisory
layout: post

## If this is a new post, reset this counter to 1.
version: 1

## Only true if release announcement or security annoucement. English posts only
announcement: 1

excerpt: >
  格式錯誤的 `GETDATA` 訊息可能在接收節點上觸發 100% CPU 使用率。修復程式於 2020 年 6 月 3 日在 Bitcoin Core 0.20.0 中發布。
---

格式錯誤的 `GETDATA` 訊息可能在接收節點上觸發無限迴圈，使用分配給此執行緒的 100% CPU，並且不在此連線上取得進一步進展。

此問題被認為是**低**嚴重性。

## 細節

在 Bitcoin Core 0.20.0 之前，攻擊者（甚至是有錯誤的客戶端）可以向我們發送一個 `GETDATA` 訊息，這將導致我們的 net_processing 執行緒開始以 100% 的速度旋轉，並且不再為攻擊者對等節點處理訊息取得進展。它仍然會為其他對等節點處理訊息取得進展，所以除此之外它只是一個 CPU DoS，影響很小（不為攻擊者對等節點取得進展不是問題）。它還將每個攻擊者對等節點的長期記憶體使用量增加了 1.5 MB。

John Newbery 開啟 [PR #18808](https://github.com/bitcoin/bitcoin/pull/18808) 以修復此問題，僅披露缺乏進展。

## 歸屬

感謝 John Newbery 發現此錯誤、負責任地披露它並修復它。

## 時間表

- 2020-04-29 John Newbery 開啟 #18808
- 2020-05-08 John Newbery 透過電子郵件報告他的發現
- 2020-05-12 #18808 被合併
- 2020-06-03 Bitcoin Core 版本 0.20.0 發布，包含修復程式
- 2021-09-13 最後一個易受影響的 Bitcoin Core 版本（0.19.x）終止支援
- 2024-07-03 公開披露。

{% include references.md %}
