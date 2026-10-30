---
title: CVE-2024-52913 - 由於交易重新請求處理導致的審查
name: blog-disclose-already-asked-for
id: zh_TW-blog-disclose-already-asked-for
lang: zh_TW
type: advisory
layout: post

## If this is a new post, reset this counter to 1.
version: 2

## Only true if release announcement or security annoucement. English posts only
announcement: 1

excerpt: >
  惡意對等節點可以阻止節點看到特定的未確認交易。修復程式於 2021 年 1 月 14 日在 Bitcoin Core 0.21.0 中發布。
---

攻擊者可以阻止節點看到特定的未確認交易。

此問題被認為是**中等**嚴重性。

## 細節

在 PR 19988 修復此問題之前，「g_already_asked_for」機制用於排程交易的 `GETDATA` 請求。`SendMessages()` 函式會為對等節點最近公告的交易發出 `GETDATA`，記住該請求在 `g_already_asked_for` 中的發送時間。然而，這個 `g_already_asked_for` 是一個「limitedmap」資料結構，具有有界大小，如果達到 50000 個條目，將忘記最舊的條目。這使得以下攻擊成為可能：
* 攻擊者是第一個向受害者公告合法交易 T 的人。
* 受害者使用 `GETDATA` 從攻擊者請求 T。
* 攻擊者在接近受害者將從其他對等節點請求 T 的時間（約 60 秒）之前不回應 `GETDATA`。
* 然後，攻擊者仔細地向受害者發送虛假公告垃圾訊息，導致受害者的 `g_already_asked_for` 驅逐 T。
* 攻擊者再次向受害者公告 T（由於 `m_tx_process_time` 中的佇列工作方式，這不需要特別精確的時間）。
* 受害者在 `g_already_asked_for` 中找不到 T，會將其視為新公告，向攻擊者發送新的 `GETDATA`。
* 攻擊者再次不回應 `GETDATA`。
* 等等。

這樣，攻擊者可以阻止受害者從除攻擊者之外的任何人請求交易。

## 歸屬

由 John Newbery 負責任地披露，聲稱由 Amiti Uttarwar 和他發現。

## 時間表

- 2020-04-03 John Newbery 在一封電子郵件中向 Suhas Daftuar 和其他人報告該錯誤
- 2020-05-08 John Newbery 建議修復該錯誤的方法
- 2020-09-21 Pieter Wuille 開啟 [PR #19988](https://github.com/bitcoin/bitcoin/pull/19988) 作為修復此錯誤和其他錯誤的綜合方法
- 2020-10-14 Pieter 的 PR 被合併
- 2021-01-14 Bitcoin Core 版本 0.21.0 發布，包含修復程式
- 2022-04-25 最後一個易受影響的 Bitcoin Core 版本（0.20.x）終止支援
- 2024-07-03 公開披露

{% include references.md %}
