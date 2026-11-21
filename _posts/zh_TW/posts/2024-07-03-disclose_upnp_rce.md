---
title: CVE-2015-20111 - 由於 miniupnpc 中的錯誤導致遠端程式碼執行
name: blog-disclose-upnp-rce
id: zh_TW-blog-disclose-upnp-rce
lang: zh_TW
type: advisory
layout: post

## If this is a new post, reset this counter to 1.
version: 1

## Only true if release announcement or security annoucement. English posts only
announcement: 1

excerpt: >
  miniupnpc 程式庫中的錯誤可能導致 Bitcoin Core 中的遠端程式碼執行。修復程式於 2015 年 10 月 15 日在 Bitcoin Core 0.11.1 中發布。
---

在 `miniupnpc` 中發現了一個啟用重大資料洩漏的緩衝區溢位。結合當時最近披露的 CVE-2015-6031，它在 `miniupnpc` 中啟用了 RCE，這可能導致 Bitcoin Core 中的 RCE。這在 2016 年 2 月發布的 [Bitcoin Core 0.12](https://bitcoincore.org/zh_TW/releases/0.12.0/) 中得到修復。

此問題被認為是**中等**嚴重性。

## 細節

[CVE-2015-6031](https://nvd.nist.gov/vuln/detail/CVE-2015-6031)，於 2015 年 9 月披露，使惡意 UPnP 伺服器有可能在啟動時遠端崩潰本機網路上的 Bitcoin Core 程序。詳情請參見[此處](https://nvd.nist.gov/vuln/detail/CVE-2015-6031)。該修復程式在 [Bitcoin Core 中拉取](https://github.com/bitcoin/bitcoin/pull/6789)，並在 2015 年 10 月發布的[版本 0.11.1](https://bitcoincore.org/zh_TW/releases/0.11.1/) 中發布。UPnP 隨後[預設關閉](https://github.com/bitcoin/bitcoin/pull/6795)。

CVE-2015-6031 披露了一個緩衝區溢位，除了啟用遠端崩潰之外，還可能使遠端在受害者的機器上執行程式碼成為可能。在調查這種可能性時，Wladimir J. Van Der Laan 在 `miniupnpc` 中發現了另一個啟用重大資料洩漏的緩衝區溢位。這在提交 `4c90b87ce3d2517097880279e8c3daa7731100e6` 中由 [Wladimir 在 `miniupnpc` 中修復](https://github.com/miniupnp/miniupnp/pull/157)。該修復程式隨後[拉取到 Bitcoin Core](https://github.com/bitcoin/bitcoin/pull/6980) 中，並作為版本 0.12 的一部分發布。

此資料洩漏不會直接披露秘密資訊（例如錢包的私鑰）。但與另一個堆疊溢位（例如 CVE-2015-6031 中披露的那個）結合，這使得觸發遠端程式碼執行成為可能。Wladimir 針對 Ubuntu 的 `miniupnpc` 版本 `1.6-precise` 演示了這一點。然而，此漏洞利用中使用的特定方法並不直接可移植到 Bitcoin Core。

## 歸屬

感謝 Aleksandar Nikolic 識別 CVE-2015-0035，以及 Wladimir J. Van Der Laan 調查其影響並發現第二個緩衝區溢位。

## 時間表

- 2015-09-15 CVE-2015-0035 被[修復](https://github.com/miniupnp/miniupnp/commit/79cca974a4c2ab1199786732a67ff6d898051b78)並[披露](https://talosintelligence.com/vulnerability_reports/TALOS-2015-0035/)。
- 2015-10-09 [PR #6789](https://github.com/bitcoin/bitcoin/pull/6789) 在 Bitcoin Core 中被合併
- 2015-10-14 Wladimir 透過利用第二個緩衝區溢位的遠端程式碼執行向 Ubuntu 安全和 Bitcoin 開發者披露。
- 2015-10-15 Bitcoin Core 0.11.1 [發布](https://lists.linuxfoundation.org/pipermail/bitcoin-dev/2015-October/011545.html)
- 2015-10-26 第二個緩衝區溢位的修復程式[被合併](https://github.com/miniupnp/miniupnp/pull/157)到 `miniupnpc` 中。
- 2015-12-18 該修復程式[被拉取到 Bitcoin Core](https://github.com/bitcoin/bitcoin/pull/6980)。
- 2016-02-23 Bitcoin Core 版本 0.12 [發布](https://lists.linuxfoundation.org/pipermail/bitcoin-dev/2016-February/012456.html)。
- 2017-03-08 最後一個易受影響的 Bitcoin Core 版本（0.11.x）終止支援
- 2024-07-03 公開披露

{% include references.md %}
