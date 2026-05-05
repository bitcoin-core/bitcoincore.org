---
title: CVE-2024-52916 - 使用低難度標頭的記憶體 DoS
name: blog-disclose-header-spam-checkpoint-bypass
id: zh_TW-blog-disclose-header-spam-checkpoint-bypass
lang: zh_TW
type: advisory
layout: post

## If this is a new post, reset this counter to 1.
version: 1

## Only true if release announcement or security annoucement. English posts only
announcement: 1

excerpt: >
  節點可能被低難度標頭垃圾訊息淹沒，這可能被用來崩潰它。修復程式於 2017 年 9 月 14 日在 Bitcoin Core 0.15.0 中發布。
---

在 Bitcoin Core 0.12.0 之後和 Bitcoin Core 0.15.0 之前，節點可能被最低難度標頭垃圾訊息淹沒，這可能被利用透過 OOM 崩潰它。

此問題被認為是**中等**嚴重性。

## 細節

在引入[標頭預同步](https://github.com/bitcoin/bitcoin/pull/25717)之前，節點完全依賴檢查點來避免被低難度標頭垃圾訊息淹沒。

在 Bitcoin Core 0.12.0 中，檢查在最後一個檢查點高度之前分叉的標頭的檢查被移動到將標頭儲存在 `mapBlockIndex` 之後。這允許攻擊者透過垃圾訊息發送其父節點是創世區塊（只需要難度 1 來建立）的標頭來無限制地增長映射，因為這些區塊繞過了檢查點邏輯。

## 歸屬

感謝 Cory Fields 發現並負責任地披露該錯誤。

## 時間表

- 2017-08-08 Cory Fields 私下報告該錯誤
- 2017-08-11 Pieter Wuille 開啟 [PR #11028](https://github.com/bitcoin/bitcoin/pull/11028) 以修復它
- 2017-08-14 PR #11028 被合併
- 2017-09-14 Bitcoin Core 版本 0.15.0 發布，包含修復程式
- 2018-10-03 最後一個易受影響的 Bitcoin Core 版本（0.14.3）終止支援
- 2024-07-03 公開披露。

{% include references.md %}
