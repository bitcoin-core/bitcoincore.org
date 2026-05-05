---
title: CVE-2024-52914 - 由於孤兒處理導致的重大 DoS
name: blog-disclose-orphan-dos
id: zh_TW-blog-disclose-orphan-dos
lang: zh_TW
type: advisory
layout: post

## If this is a new post, reset this counter to 1.
version: 1

## Only true if release announcement or security annoucement. English posts only
announcement: 1

excerpt: >
  節點在接收到特別建構的未確認交易時可能會停滯數小時。修復程式於 2019 年 5 月 18 日在 Bitcoin Core 0.18.0 中發布。
---

節點在處理特別建構的未確認交易的孤兒時可能會停滯數小時。

此問題被認為是**高**嚴重性。

## 細節

在將交易接受到其記憶池後，節點將瀏覽其孤兒交易快取，以找出這個新接受的交易是否使接受任何交易成為可能。此搜尋是二次方的：對於新接受的交易中的每個輸出，它將瀏覽所有快取的孤兒交易（限制為 100）。透過特別建構孤兒交易為無效但驗證昂貴，節點可能會停滯數小時。

停滯由 Pieter Wuille 在 [PR #15644](https://github.com/bitcoin/bitcoin/pull/15644) 中修復，透過在找到匹配時中斷孤兒解析以處理新訊息（無論孤兒最終是否有效）。

## 歸屬

感謝 sec.eine 負責任地披露該錯誤並提供關於修復程式的回饋。

## 時間表

- 2019-03-19 sec.eine 透過電子郵件向 Greg Maxwell 報告問題
- 2019-03-21 Greg Maxwell 回應有關建議修補程式的資訊
- 2019-03-22 sec.eine 對修補程式給出回饋（「看起來很穩固且 [..] 不會引起注意」）
- 2019-03-22 Pieter Wuille 開啟 PR #15644
- 2019-04-01 PR #15644 被合併
- 2019-05-18 Bitcoin Core 版本 0.18.0 發布，包含修復程式
- 2020-07-22 在 PR 審查俱樂部期間[部分披露](https://bitcoincore.reviews/15644#l-285)該問題
- 2020-08-01 最後一個易受影響的 Bitcoin Core 版本（0.17.x）終止支援
- 2024-07-03 公開披露。

{% include references.md %}
