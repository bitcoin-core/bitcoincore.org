---
title: CVE-2020-14198 披露
name: blog-disclose-unbounded-banlist
id: zh_TW-blog-disclose-unbounded-banlist
lang: zh_TW
type: advisory
layout: post

## If this is a new post, reset this counter to 1.
version: 1

## Only true if release announcement or security annoucement. English posts only
announcement: 1

excerpt: >
  節點在受到大量不同 IP 攻擊時可能會受到 CPU 和記憶體 DoS 的影響。修復程式於 2020 年 8 月 1 日在 Bitcoin Core 0.20.1 中發布。
---

Bitcoin Core 維護了一個無限制的禁止 IP 地址列表，並對其執行二次方操作。這可能導致 OOM 崩潰和 CPU DoS。

此問題被認為是**高**嚴重性。

## 細節

Bitcoin Core 維護了一個禁止 IP 地址列表。此列表沒有界限，可以被對手操縱。考慮到 IPV6，為攻擊者向此列表新增新條目特別便宜。此外，在接收到 `GETADDR` 訊息時，Bitcoin Core 會掃描整個禁止列表以檢查要返回的每個地址（最多 2500 個）。

## 歸屬

Calin Culianu 首先負責任地披露它。Calin 後來在 [PR 評論](https://github.com/bitcoin/bitcoin/pull/15617#issuecomment-640898523)中公開披露了該錯誤。

同一天，Bitcoin ABC 的 Jason Cox 透過電子郵件與 Bitcoin Core 專案分享了他們也收到的同一報告。

## 時間表

- 2020-06-08 Calin Culianu 私下向 Bitcoin Core 專案報告該錯誤
- 2020-06-08 Jason Cox 私下與 Bitcoin Core 分享傳送給 Bitcoin ABC 的（同一）報告
- 2020-06-08 Calin Culianu 在引入二次方行為的原始 PR 上公開披露該漏洞
- 2020-06-09 Pieter Wuille 開啟 PR [#19219](https://github.com/bitcoin/bitcoin/pull/19219)，修復了無界記憶體使用和二次方行為
- 2020-06-16 Luke Dashjr 在他的請求後被分配 CVE-2020-14198 用於此漏洞
- 2020-07-07 Pieter 的 PR 被合併
- 2020-08-01 Bitcoin Core 0.20.1 發布，包含修復程式
- 2021-01-14 Bitcoin Core 0.21.0 發布，包含修復程式
- 2022-04-25 最後一個易受影響的 Bitcoin Core 版本（0.20.0）終止支援
- 2024-07-03（官方）公開披露

{% include references.md %}
