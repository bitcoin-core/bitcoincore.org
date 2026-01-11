---
title: CVE-2024-52918 - 使用惡意 BIP72 URI 導致崩潰
name: blog-disclose-bip70-crash
id: zh_TW-blog-disclose-bip70-crash
lang: zh_TW
type: advisory
layout: post

## If this is a new post, reset this counter to 1.
version: 2

## Only true if release announcement or security annoucement. English posts only
announcement: 1

excerpt: >
  Bitcoin-Qt 中的 BIP70 實作在開啟 BIP72 URI 時可能會靜默崩潰。修復程式於 2020 年 6 月 3 日在 Bitcoin Core 0.20.0 中發布。
---

Bitcoin-Qt 在開啟 [BIP72](https://github.com/bitcoin/bips/blob/master/bip-0072.mediawiki) URI 時可能會崩潰。

此問題被認為是**中等**嚴重性。

## 細節

[BIP72](https://github.com/bitcoin/bips/blob/master/bip-0072.mediawiki) 使用 `r` 參數擴展了 BIP21 URI 方案，以從中獲取付款請求。攻擊者可以簡單地將 `r` 參數中包含的 URL 指向一個非常大的檔案，Bitcoin-Qt 會嘗試為其分配足夠的記憶體並崩潰。

受害者可能會被誘騙開啟惡意付款請求。大型下載將在背景中進行，在 GUI 中幾乎沒有輸出，直到應用程式記憶體不足。

## 歸屬

感謝 Michael Ford (Fanquake) 負責任地披露問題並提供 PoC。

## 時間表

- 2019-08-12 Michael Ford 向 Cory Fields 和 Wladimir Van Der Laan 報告該錯誤
- 2019-10-16 Michael Ford 開啟 PR [#17165](https://github.com/bitcoin/bitcoin/pull/17165) 以完全移除 BIP70 支援
- 2019-10-26 Michael 的 PR 被合併到 Bitcoin Core
- 2020-06-03 Bitcoin Core 版本 0.20.0 發布
- 2021-09-13 最後一個易受影響的 Bitcoin Core 版本（0.19.x）終止支援
- 2024-07-03 公開披露

{% include references.md %}
