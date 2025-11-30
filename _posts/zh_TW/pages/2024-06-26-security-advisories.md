---
title: 安全公告
name: Security Advisories
id: zh_tw-security-advisories
permalink: /zh_TW/security-advisories/
layout: page
type: pages
lang: zh_TW
version: 1
---
{% include toc.html %}

本頁面概述了與 Bitcoin Core 中漏洞披露相關的政策，
以及歷史安全公告的摘要。

## 政策

所有漏洞都應報告至 security@bitcoincore.org（詳情請參閱
[SECURITY.md](https://github.com/bitcoin/bitcoin/blob/master/SECURITY.md)）。
報告後，漏洞將被分配一個嚴重性類別。
我們區分 4 類漏洞：

* **嚴重（Critical）**：威脅整個 Bitcoin 網路的基本安全性和完整性的錯誤。
  這些錯誤允許在協議層面上竊取幣、在指定發行時程之外創造幣，或造成永久性、
  全網路範圍的鏈分裂。
  <details markdown="1">

  <summary>
  範例
  </summary>

  * 允許在區塊內兩次花費同一交易輸出以增加貨幣供應的錯誤
    ([CVE-2018-17144](/en/2018/09/20/notice/))。
  * 一個共識失敗，其中執行較舊軟體的節點拒絕了較新軟體接受的區塊，
    原因是底層資料庫限制，導致全網路範圍的鏈分裂
    ([BIP 50](https://github.com/bitcoin/bips/blob/master/bip-0050.mediawiki))。

  </details>

* **高（High）**：對受影響節點或網路有重大影響的錯誤。這些錯誤通常可以在預設配置下
  遠端利用，並可能造成廣泛的中斷。
  <details markdown="1">

  <summary>
  範例
  </summary>

  * 可遠端觸發的當機，可能使許多節點離線
    ([CVE-2024-35202](/en/2024/10/08/disclose-blocktxn-crash/))。
  * 阻斷服務攻擊，導致節點長時間停滯，無法處理新的交易和區塊
    ([CVE-2024-52914](/en/2024/07/03/disclose-orphan-dos/))。
  * 記憶體耗盡漏洞，可透過使節點儲存過量的區塊標頭來遠端觸發當機
    ([CVE-2019-25220](/en/2024/09/18/disclose-headers-oom/))。

  </details>

* **中（Medium）**：可能明顯降低網路或節點性能或功能的錯誤，但其範圍或可利用性有限。
  這些錯誤可能需要特殊條件才能觸發，例如非預設設定，或導致服務降級而非完全節點故障。
  <details markdown="1">

  <summary>
  範例
  </summary>

  * 本地網路上的潛在遠端程式碼執行（RCE）漏洞，僅在啟用 UPnP 等非預設功能時才能利用
    ([CVE-2015-20111](/en/2024/07/03/disclose_upnp_rce/))。
  * 對等節點可透過發送變異區塊來阻礙區塊傳播，延遲節點接收新區塊的時間
    ([CVE-2024-52921](/en/2024/10/08/disclose-mutated-blocks-hindering-propagation/))。
  * 攻擊者向節點宣告區塊，然後無法提供該區塊，導致受害節點在能夠從其他對等節點獲取之前等待最多 10 分鐘
    ([CVE-2024-52922](/en/2024/11/05/cb-stall-hindering-propagation/))。

  </details>

* **低（Low）**：難以利用或對節點運作影響輕微的錯誤。這些錯誤可能僅在非預設配置下
  或從本地網路才能觸發，且不會構成立即或廣泛的威脅。
  <details markdown="1">

  <summary>
  範例
  </summary>

  * 格式錯誤的 `getdata` 訊息可能導致對等連接進入無限迴圈，消耗 CPU
    但不影響節點處理區塊或處理其他對等連接的能力
    ([CVE-2024-52920](/en/2024/07/03/disclose-getdata-cpu/))。
  * 依賴套件中的錯誤可能導致節點當機，但僅在啟用 UPnP 等非預設功能時
    ([CVE-2024-52917](/en/2024/07/31/disclose-upnp-oom/))。
  * 可能導致節點當機的錯誤，但極難利用
    ([CVE-2024-52919](/en/2025/04/28/disclose-cve-2024-52919/))。

  </details>

**低**嚴重性漏洞將在包含修復的主要版本發布後 2 週內披露。**中**和**高**嚴重性漏洞
將在最後一個受影響版本[終止生命週期](/zh_TW/lifecycle/)（包含修復的主要版本首次發布後
約一年）後 2 週內披露。

在發布漏洞詳情的兩週前將進行預告。此預告將與新主要版本的發布同時進行，
並包含已修復漏洞的數量及其嚴重性等級。

**嚴重**錯誤不在標準政策考慮範圍內，因為它們很可能需要臨時程序。
此外，錯誤可能根本不被視為漏洞。任何報告的問題也可能被認為是嚴重的，但不需要禁運。

## 過往安全公告

{% assign advisories=site.posts | where:"lang", 'en' | where:"type", 'advisory' | sort: "date" | reverse %}
{% for advisory in advisories %}
{% assign post=advisory %}
  <article>
    <h2><a href="{{ post.url }}" title="{{ post.title | xml_escape }}">{{ post.title }}</a></h2>
    <p>{{ post.excerpt | markdownify | strip_html | truncate: 200 }}</p>
  </article>
{% endfor %}
