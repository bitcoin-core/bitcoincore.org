---
title: CVE-2019-25220 - 由於標頭垃圾導致的記憶體 DoS
name: blog-disclose-headers-oom
id: zh_TW-blog-disclose-headers-oom
lang: zh_TW
type: advisory
layout: post

## If this is a new post, reset this counter to 1.
version: 1

## Only true if release announcement or security annoucement. English posts only
announcement: 1

excerpt: >
  攻擊者可以向 Bitcoin Core 節點垃圾發送低難度標頭鏈，這可以被用來遠端崩潰它。
---

在 Bitcoin Core v24.0.1 之前，攻擊者可以向節點垃圾發送低難度標頭鏈，這可以被用來遠端崩潰對等節點。

此問題被認為是**高**嚴重性。

## 細節

Bitcoin Core 將區塊鏈標頭儲存在記憶體中。這使其容易受到 DoS 攻擊，透過讓它下載和儲存極長的標頭鏈，即使它們是低難度的。重要的是要注意，一旦製作完成，攻擊鏈可以被重複使用來崩潰網路上的任何節點。

使用此方法攻擊節點的可能性早已為人所知，這是檢查點系統仍然存在的主要原因：讓攻擊者從最後一個檢查點開始攻擊，使其成本遠高於從創世區塊開始。然而，隨著時間推移，隨著算力成本的降低，即使這種緩解措施也變得不那麼有效。

此攻擊於 2019 年 1 月由 David Jaenson 獨立發現並向 Bitcoin Core 專案報告，他建議引入更新的檢查點作為實用的緩解措施。然而：
1. 這仍然使執行 IBD 的節點在收到檢查點區塊之前沒有保護。
2. 它依賴於生態系統半定期採用包含新檢查點的更新軟體，這是 Bitcoin Core 貢獻者長期以來不太願意接受的做法。

當 Braydon Fuller 於 2019 年 10 月在 bitcoin-dev 郵件列表上[發布他的「鏈寬擴展」報告](https://lists.linuxfoundation.org/pipermail/bitcoin-dev/2019-October/017354.html)時，它引起了更多關注。他之前已經負責任地向 Bitcoin Core 安全列表報告過。由於對限制平行鏈數量時網路收斂的擔憂，所建議的方法未被 Bitcoin Core 採用。

當時，建立一個巨大的低難度標頭鏈的計算成本約等於在頂端挖掘一個區塊的 32.28%。這是約 4.12 BTC 的成本，因為當時的區塊獎勵約為 12.77 BTC。

到 2022 年 2 月，攻擊成本進一步下降到挖掘一個區塊成本的約 14.73%，這促使對替代解決方案的調查。如果不加以解決，今天（2024 年 9 月）的成本將僅為一個區塊的 4.44%。考慮到這些日期的區塊獎勵，這些數字分別轉化為約 1.07 BTC 和 0.14 BTC 的成本。

針對此 DoS 的保護措施在 Bitcoin Core PR [#25717](https://github.com/bitcoin/bitcoin/pull/25717) 中實作，節點將首先驗證所呈現的鏈是否有足夠的工作量，然後再承諾儲存它。這樣，Bitcoin Core 不再依賴檢查點來防禦任何已知攻擊。

## 歸屬

感謝 David Jaenson 和 Braydon Fuller 獨立重新發現該攻擊，估計其成本並建議修改。

感謝 Suhas Daftuar 和 Pieter Wuille 研究令人滿意的修復程式並實作它。

## 時間表

* 2010-07-17 - Bitcoin 0.3.2 發布，引入檢查點。它們保護包括低難度區塊垃圾在內的其他威脅。
* 2011-11-21 - Bitcoin 0.5.0 發布，跳過最後一個檢查點之前的區塊的腳本驗證。這使得檢查點的作用更加安全關鍵。
* 2014-04-09 - 區塊 295000 被挖掘，成為最後一個 Bitcoin Core 檢查點。隨著算力成本降低，檢查點對區塊垃圾的保護從這一點開始侵蝕。
* 2015-02-16 - Bitcoin Core 0.10.0 發布，具有標頭優先同步。這將低難度區塊垃圾攻擊弱化為區塊*標頭*垃圾攻擊。
* 2017-03-08 - Bitcoin Core 0.14.0 發布，將跳過腳本驗證與檢查點分離，使它們僅與防禦區塊標頭垃圾相關。
* 2019-01-28 - David Jaenson 向 Bitcoin Core 安全郵件列表報告此問題。
* 2019-09-18 - Braydon Fuller 向 Bitcoin Core 安全列表傳送電子郵件，標題為「[Bitcoin 鏈寬擴展拒絕服務攻擊](https://bcoin.io/papers/bitcoin-chain-expansion.pdf)」，討論區塊和區塊標頭垃圾的危險、成本分析和建議的解決方案。
* 2019-09-26 - Suhas Daftuar 回覆 Braydon Fuller 這是一個已知問題，並邀請他將他的報告發布到 bitcoin-dev 郵件列表。
* 2019-10-04 - Braydon Fuller 將他的論文傳送到 bitcoin-dev 郵件列表。
* 2019-10-31 - 作為對上述事件的回應，Suhas Daftuar 開啟 PR [#17332](https://github.com/bitcoin/bitcoin/pull/17332)，其中包含他早期但不實用的概念驗證，希望引起更多關於該主題的討論。
* 2022-02 - Suhas Daftuar 和 Pieter Wuille 討論此問題，並估計此攻擊的成本現在實際上已經變得如此之低，以至於需要立即採取行動，並需要避免公開談論它。
* 2022-06-22 - Suhas Daftuar 開啟 PR [#25454](https://github.com/bitcoin/bitcoin/pull/25454) 作為實作修復程式的準備工作。
* 2022-06-22 - Suhas Daftuar 向一群長期貢獻者傳送訊息，詳細說明攻擊、其成本以及 Pieter Wuille 和他一直在研究的修復程式。
* 2022-07-26 - Suhas Daftuar 開啟 PR [#25717](https://github.com/bitcoin/bitcoin/pull/25717)，與 Pieter Wuille 共同撰寫，實作修復程式。
* 2022-08-30 - PR #25717 被合併。
* 2022-10-21 - Niklas Gögge 的 PR [#26355](https://github.com/bitcoin/bitcoin/pull/26355) 被合併，修復了 PR #25717 中引入的標頭預同步步驟中的錯誤。如果沒有這個，仍然可能垃圾發送區塊標頭。發現此錯誤以及可能存在未發現的錯誤的可能性，是舊檢查點尚未完全刪除的原因。
* 2022-12-12 - Bitcoin Core 24.0.1 發布，包含修復程式。
* 2023-12-07 - 最後一個易受影響的 Bitcoin Core 版本（23.2）終止支援。
* 2024-09-18 - 公開披露。

{% include references.md %}
