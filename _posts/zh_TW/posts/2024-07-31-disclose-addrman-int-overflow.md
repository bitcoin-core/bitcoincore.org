---
title: CVE-2024-52919 - 由於 addr 訊息垃圾導致的遠端崩潰
name: blog-disclose-addrman-idcount-in-overflow
id: blog-disclose-addrman-idcount-in-overflow
lang: zh_TW
type: advisory
layout: post

## If this is a new post, reset this counter to 1.
version: 1

## Only true if release announcement or security annoucement. English posts only
announcement: 1

excerpt: >
  節點可能被 addr 訊息垃圾淹沒，這可能被用來崩潰它們。修復程式於 2021 年 9 月 14 日在 Bitcoin Core v22.0 中發布。
---

披露一個整數溢位錯誤的細節，該錯誤會導致斷言崩潰，修復程式於 2021 年 9 月 14 日在 Bitcoin Core 版本 v22.0 中發布。

此問題被認為是**高**嚴重性。

## 細節

`CAddrMan` 有一個 32 位元 `nIdCount` 欄位，每次插入 addrman 時都會遞增，然後成為新條目的識別碼。透過讓受害者插入 2<sup>32</sup> 個條目（例如透過垃圾訊息發送 addr 訊息），此識別碼會溢位，導致斷言崩潰。

## 歸屬

感謝 Eugene Siegel 發現並披露該漏洞，以及 Pieter Wuille 在 https://github.com/bitcoin/bitcoin/pull/22387 中修復該問題。

## 時間表

* 2021-06-21 - Eugene Siegel 向 security@bitcoincore.org 傳送初始報告
* 2021-07-19 - 修復程式被合併 (https://github.com/bitcoin/bitcoin/pull/22387)
* 2021-09-13 - v22.0 發布
* 2024-07-31 - 公開披露

{% include references.md %}
