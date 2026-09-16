---
type: posts
layout: post
lang: zh_TW
name: performance-optimizations-1
id: zh_TW-performance-optimizations-1
title: 鏈上擴容 - Bitcoin 參考軟體歷史性能優化回顧。第 1 部分
permalink: /zh_TW/2017/03/13/performance-optimizations-1/
version: 1
excerpt: 多年來幫助 Bitcoin 軟體客戶端使用者保持可靠體驗的開發里程碑。
---

以下文章旨在強調多年來幫助 Bitcoin 軟體客戶端使用者保持可靠體驗的開發里程碑。我們展示了幾個關鍵升級，這些升級對於維護網路的去中心化特性和減輕參與者的資源負擔至關重要。我們描述了如何進行數個數量級的優化，使 Bitcoin 網路能夠支援交易活動的增長，而不會大幅增加新使用者和現有使用者的參與成本。最後，我們注意到這些改進都屬於更大的、系統性的協定開發方法的一部分，該方法使用來自 Big-O 複雜度概念的見解，並利用更智慧的演算法來更有效地利用網路資源。

## 簽名快取
發布版本：Bitcoin-Qt 0.7.0

ECSDA 簽名驗證是在對等節點層級完成的計算量最大的操作之一。因為它們需要對每個交易進行驗證，任何多餘的驗證都會給終端使用者帶來顯著的資源開銷。早期版本的軟體就是這種情況，它們會在交易進入節點記憶池之前和接收到區塊後都驗證簽名。

為了提高效率，開發人員建立了一個快取，允許節點儲存先前驗證過的簽名，並在交易進入已接受區塊後跳過冗餘工作。

此外，簽名快取還減輕了由特別製作的交易可能導致 Bitcoin 客戶端停滯所引入的 DoS 向量。

### 進一步資訊

 * [Bitcoin-Qt 0.7.0 發布說明](https://bitcoin.org/en/release/v0.7.0#core-bitcoin-handling-and-blockchain-database)
 * [已修復的漏洞解釋：為什麼簽名快取是 DoS 保護](https://bitcointalk.org/index.php?topic=136422.0)

## Ultraprune + LevelDB
發布版本：Bitcoin-Qt 0.8.0

Ultraprune 是 Bitcoin 軟體的首批主要升級之一，旨在解決與驗證區塊鏈交易資料相關的開銷。Bitcoin 參考客戶端的早期版本維護了所有交易輸出（已花費或未花費）的索引。Ultraprune 透過洞察到您只需要追蹤未花費的輸出，並且輸出一旦被花費就可以完全從索引中刪除，從而顯著減少了該索引的大小。

為了實現這一點，實作了一種新的資料庫佈局，該佈局將未花費交易輸出分配到緊湊的自訂格式，以減少驗證工作所需的資訊大小。

為了進一步優化系統的性能，Ultraprune 與 LevelDB 同時引入，後者取代了舊的 BDB 資料庫技術。總體而言，影響是顯著的：根據其硬體，使用者在驗證區塊鏈資料時可以體驗到至少一個數量級的改進。這種新的資料庫結構也為未來關於修剪和更輕量級的 Bitcoin 全節點實作的工作鋪平了道路。

### 進一步資訊

  * [Bitcoin-Qt 0.8.0 發布說明](https://bitcoin.org/en/release/v0.8.0#improvements)
  * [用簡單英語解釋 Ultraprune](https://archive.is/bUocJ)
  * [Ultraprune 合併到主線](https://bitcointalk.org/index.php?topic=119525.0)
  * [參考客戶端中的修剪：ultraprune 模式](https://bitcointalk.org/index.php?topic=91954.0)

## 平行腳本驗證
發布版本：Bitcoin-Qt 0.8

雖然是一個更微妙的變化，但將腳本驗證轉變為更平行化的過程消除了區塊驗證時間的顯著開銷。早期版本的軟體會在每次 UTXO 提取之間驗證輸入的腳本資料，由於所有操作的線性處理而產生性能問題。這違反了設計高效計算過程的基本原則，該原則規定計算應盡可能與 I/O 作業同時發生。

考慮到這一點，區塊驗證機制被重新設計，以便能夠將腳本檢查分配給平行執行緒，以便即使在客戶端完成從區塊提取所有 UTXO 之前，它們的驗證也可以進行。為了實現這一點，腳本檢查操作在處理交易後儲存在佇列中，並與其他輸入驗證作業分開處理。

因此，透過更有效地利用對等節點的資源，同步到鏈的頂端發生得更快。在實作的測試期間，開發人員注意到與軟體早期版本相比，速度提升了 35% 到 100%。

### 進一步資訊

  * [平行腳本驗證 #2060](https://github.com/bitcoin/bitcoin/pull/2060)

## 標頭優先同步
發布版本：Bitcoin Core 0.10

為了進一步改善初始區塊下載時間，Core 專案在 2014 年底引入了節點用於與最多工作量有效鏈同步的機制的重要重新架構。

最初，引導新 Bitcoin 客戶端的過程涉及使用者從單個對等節點提取區塊資料，其結果是任何中斷或連接品質下降都會顯著停滯該過程。隨著區塊鏈大小的不斷增加，這將導致同步完成的等待時間有時會很長，很大比例的使用者報告根據其硬體最多需要數天時間。

標頭優先同步使用一種新方法完全緩解了這個問題，該方法涉及節點首先從單個對等節點下載和驗證區塊標頭，然後從眾多其他對等節點平行提取區塊資料。

自 Bitcoin 早期以來，關於初始區塊下載時間的抱怨一直很普遍。透過標頭優先同步，軟體在新使用者的可用性方面邁出了重要一步。節點現在可以利用其整個對等節點網路，而不是在不可靠的同步上浪費許多小時，並顯著縮短引導時間。透過使用更智慧的演算法，對 Bitcoin 長期永續性的這個關鍵方面進行了另一個漸近改進。

### 進一步資訊

  * [Bitcoin-Qt 0.10.0 發布說明](https://bitcoin.org/en/release/v0.10.0#faster-synchronization)
  * [Bitcoin.org 開發者指南](https://bitcoin.org/en/developer-guide#headers-first)
  * [Pieter Wuille 在 Bitcoin-dev 郵件列表上的文章](https://lists.linuxfoundation.org/pipermail/bitcoin-dev/2014-October/006724.html)

## 區塊檔案修剪
發布版本：Bitcoin Core 0.11

舊資料修剪是中本聰在其白皮書中首次描述的概念，作為磁碟空間稀缺的潛在解決方案。不幸的是，原始設計不充分，無法按照其創建者的想像實作。七年後，隨著區塊鏈達到超過一百 GB，我們今天所知的區塊檔案修剪的引入為資源有限的使用者帶來了重大福音。

區塊檔案修剪利用了 ultraprune 的先前工作；最初下載並驗證區塊鏈的使用者現在可以丟棄早於 288 個區塊的原始資料。因為這些節點仍然保留完整的 UTXO 集，它們仍然能夠驗證未花費的資料，這足以完全驗證新區塊並防止潛在的雙重花費。

當然，修剪意味著仍然需要足夠數量的存檔節點來提供完整的區塊鏈資料。另一方面，這項創新透過使保持驗證者身份更具成本效益來擴大了驗證者的範圍。總體而言，該解決方案是我們增強網路多樣性可用選項的受歡迎補充。

### 進一步資訊

  * [Bitcoin-Qt 0.11.0 發布說明](https://bitcoin.org/en/release/v0.11.0#block-file-pruning)

## libsecp256k1
發布版本：Bitcoin Core 0.12

經過測量，確定解決區塊鏈下載效率低下問題後的下一步是解決交易驗證及其大量計算負載的瓶頸。Core 專案著手透過使用為 ECDSA 操作的優化性能而設計的新程式庫來做到這一點。ECDSA（橢圓曲線數位簽章演算法）是 Bitcoin 公鑰基礎設施的骨幹，每當使用者透過使用其私鑰簽署訊息來移動幣時都會使用它。這些簽名需要由網路中的每個對等節點驗證，以保持帳本的完整性。

早期開發人員長期以來一直考慮從原始 OpenSSL 程式庫轉換，經過 5 年的設計考慮、測試和同行評審，Pieter Wuille 的 libsecp256k1 程式庫被引入作為其替代品。正如預期的那樣，該實作導致每個 Bitcoin 交易背後的簽名驗證過程的重大加速。基準測試報告比 OpenSSL 實作改進了 5-7 倍。對於使用者來說，這將轉化為節省多達一半通常用於 ECSDA 操作的引導時間，這是從頭開始同步新節點最費力的步驟之一。

考慮到 Bitcoin 交易活動的增長，此升級對於為網路對等節點保持合理的使用者體驗至關重要。再一次，演算法複雜度的降低為使用者提供了更有效的資源使用，並大幅降低了新參與者的進入門檻。

### 進一步資訊

  * [Bitcoin-Qt 0.12.0 發布說明](https://bitcoin.org/en/release/v0.12.0#signature-validation-using-libsecp256k1)
  * [Andrew Poelstra (andytoshi) 關於 libsecp256k1 的安全性和測試](https://bitcointalk.org/index.php?action=profile;u=80376)
  * [Greg Maxwell 關於 libsecp256k1 測試揭示 OpenSSL 中的錯誤](https://www.reddit.com/r/Bitcoin/comments/2rrxq7/on_why_010s_release_notes_say_we_have_reason_to/)
  * [Greg Maxwell 在 DevCore 的演講](https://www.youtube.com/watch?v=RguZ0_nmSPw&t=1297)
  * [Hal Finney 關於 libsecp256k1 的文章](https://bitcointalk.org/index.php?topic=3238.0)

## 記憶池限制
發布版本：Bitcoin Core 0.12

Bitcoin 軟體長期存在的漏洞是無法正確處理對等節點記憶池的氾濫。實際上，攻擊者可以傳送大量低價值、低費用的交易，這些交易會在記憶池中累積，直到它會使可用記憶體過載。這將導致具有相對較低 RAM 資源的節點在異常活動期間崩潰。對此唯一有效的措施是增加軟體的最低中繼費用，但這仍然沒有對記憶池的潛在大小設定上限。

為了補救這一點，Core 專案的開發人員實作了嚴格的記憶池最大大小，並具有特定的驅逐策略，按費用對交易進行排序，並首先驅逐支付最低的交易。為了防止具有相同費用的交易重新進入記憶池，節點將增加其有效的最低中繼費率，以匹配最後一個被驅逐交易的費率加上初始最低中繼費率。

最大大小的配置留給使用者，預設大小為 300 MB。此更新為資源有限的節點使用者提供了更強大的體驗，並且總體上使整個網路更加可靠。

### 進一步資訊

* [Bitcoin-Qt 0.12.0 發布說明](https://bitcoin.org/en/release/v0.12.0#memory-pool-limiting)

在第 2 部分中，我們將討論基於上述技術的最新改進，並進一步提高網路的穩健性和擴展潛力。
