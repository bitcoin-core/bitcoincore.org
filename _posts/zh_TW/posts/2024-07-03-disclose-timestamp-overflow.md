---
title: CVE-2024-52912 - 由於時間戳記調整導致的網路分裂
name: blog-disclose-timestamp-overflow
id: zh_TW-blog-disclose-timestamp-overflow
lang: zh_TW
type: advisory
layout: post

## If this is a new post, reset this counter to 1.
version: 1

## Only true if release announcement or security annoucement. English posts only
announcement: 1

excerpt: >
  節點在被其前 200 個對等節點攻擊時可能會與網路分裂。修復程式於 2021 年 1 月 15 日在 Bitcoin Core 版本 0.21.0 中發布。
---

披露一個整數溢位錯誤的細節，該錯誤有引起網路分裂的風險，修復程式於 2021 年 1 月 15 日在 Bitcoin Core 版本 0.21.0 中發布。

此問題被認為是**中等**嚴重性。

## 技術細節

網路分裂漏洞源於 `version` 訊息處理程式碼中的兩個獨立錯誤：
* 計算新連線對等節點的時間偏移時的有號整數溢位。
* abs64 邏輯錯誤（`abs64(std::numeric_limits<int64_t>::min()) == std::numeric_limits<int64_t>::min()`），導致繞過最大時間調整限制。

這兩個錯誤允許攻擊者強制受害者的調整時間（`系統時間 + 網路時間偏移`）被偏斜，使得任何新區塊因時間戳記過於未來而被拒絕。應該注意的是，此攻擊假設攻擊者在連線到受害者的前 200 個對等節點中，因為只有來自這些初始連線的時間偏移才會被納入調整時間。

## 歸屬

感謝 [practicalswift](https://github.com/practicalswift) 發現並提供漏洞的初始修復程式，以及 Pieter Wuille 提供修復程式以及對風險程式碼的一般清理。

## 時間表

* 2020-10-10 初始報告傳送到 security@bitcoincore.org
* 2020-10-13 修復程式合併到 Bitcoin Core (https://github.com/bitcoin/bitcoin/pull/20141)
* 2021-01-15 v0.21.0 發布
* 2022-04-25 最後一個易受影響的 Bitcoin Core 版本（0.20.x）終止支援
* 2024-07-03 公開披露

{% include references.md %}
