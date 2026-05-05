---
title: CVE-2024-35202 披露
name: blog-disclose-blocktxn-crash
id: zh_TW-blog-disclose-blocktxn-crash
lang: zh_TW
type: advisory
layout: post

## If this is a new post, reset this counter to 1.
version: 1

## Only true if release announcement or security annoucement. English posts only
announcement: 1

excerpt: >
  攻擊者可以透過觸發 blocktxn 訊息處理邏輯中的斷言來遠端崩潰 Bitcoin Core 節點。
---

在 Bitcoin Core v25.0 之前，攻擊者可以透過觸發 blocktxn 訊息處理邏輯中的斷言來遠端崩潰 Bitcoin Core 節點。

此問題被認為是**高**嚴重性。

## 細節

當透過 cmpctblock 訊息接收到區塊公告時，Bitcoin Core 嘗試使用其自己記憶池中的交易以及其他可用的交易來重建公告的區塊。如果由於缺少交易而重建失敗，它將透過 getblocktxn 訊息從公告對等節點請求它們。作為回應，預期會收到 blocktxn 訊息，其中應包含所請求的交易。

緊湊區塊協定採用縮短的交易識別碼來減少頻寬。這些短 ID 大小為 6 位元組，在區塊重建時導致碰撞的小機率（即交易 A 與交易 B 具有相同的短 ID）。碰撞將被檢測到，因為從重建的交易集計算的默克爾根將與區塊公告中的默克爾根不匹配。對等節點不應因碰撞而受到懲罰，因為它們可能偶然發生，因此透過回退到請求完整區塊來處理它們。

Bitcoin Core 將在每次收到新的緊湊區塊時建立一個 <code>PartiallyDownloadedBlock</code> 實例。如果請求缺少的交易，該實例將保持到處理相應的 blocktxn 訊息為止。收到 blocktxn 訊息後，將呼叫 <code>PartiallyDownloadedBlock::FillBlock</code>，嘗試重建完整區塊。在上述碰撞情況下，請求完整區塊，但 <code>PartiallyDownloadedBlock</code> 實例以及與底層區塊請求相關的其他狀態保持不變。這為同一區塊的第二個 blocktxn 訊息留下了處理空間，並觸發 <code>FillBlock</code> 再次被呼叫。這違反了 <code>FillBlock</code> 只能被呼叫一次的假設（記錄為 <code>assert</code> 陳述），並導致節點崩潰。

攻擊者不需要透過觸發碰撞來獲得幸運，因為碰撞處理邏輯可以透過簡單地在 blocktxn 訊息中包含未提交到區塊默克爾根的交易來輕鬆觸發。

## 歸屬

感謝 Niklas Gögge 發現和披露該漏洞，以及在 https://github.com/bitcoin/bitcoin/pull/26898 中修復該問題。

## 時間表

* 2022-10-05 - Niklas Gögge 向 Bitcoin Core 安全郵件列表報告該問題。
* 2023-01-24 - 包含修復程式的 PR #26898 被合併。
* 2023-05-25 - Bitcoin Core 25.0 發布，包含修復程式。
* 2024-10-09 - 公開披露。

{% include references.md %}
