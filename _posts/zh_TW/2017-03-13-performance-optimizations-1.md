---
layout: post
type: posts
lang: zh_TW
name: performance-optimizations-1
id: zh_tw-performance-optimizations-1
title: 鏈上擴容：比特幣參考軟體性能優化曆史回顧
permalink: /zh_TW/2017/03/11/performance-optimizations-1/
version: 1
excerpt: 多年來的發展里程碑有助於為比特幣軟體客戶端的用戶保留可靠的體驗。
---


最近幾年為維持可靠的用戶體驗，比特幣軟體進行了很多里程碑式的發展，本文對這些重大發展進行了回顧。我們對這個軟體進行的一些升級對維護比特幣網路去中心化特徵、紓解參與者負擔來說是非常重要的。本文描述了我們對這個軟體所做的大量成倍提升性能的最佳化，這些最佳化讓比特幣網路能夠容納更多的交易活動，同時沒有給新參與者和現有參與者增加過多的參與成本。最後，我們解釋了為什麼說這些改善是一個更大且系統化協議開發方式的一部分，這種開發方式從Big-O複雜概念中獲得的啟發，並使用了更精妙的算法讓這個網路的資源發揮出更大的效力。
 
## 簽名緩衝
發布版本：Bitcoin-Qt 0.7.0
ECSDA簽名驗證是節點層面最耗計算資源的任務之一。因為每個交易都要進行這樣的驗證，所有任何多餘的驗證都會讓終端用戶消耗大量的資源。這樣就造成這樣一種情況：早期版本軟體在交易進入一個節點記憶體池前和被接收到一個區塊後都要對簽名進行驗證。
為了提高效率，開發者創建出一個快取，可以存儲之前驗證過的簽名，一旦這些交易被打包進區塊，很多冗餘的工作量就能省略。
另外，這個簽名快取還可以消除一個DoS攻擊路徑，透過構造一些特殊結構的交易來使比特幣用戶端的運行緩慢。
 
### 更多相關訊息請參考：
 * [Bitcoin-Qt 0.7.0 Release notes](https://bitcoin.org/en/release/v0.7.0#core-bitcoin-handling-and-blockchain-database)
 * [Fixed vulnerability explanation: Why the signature cache is a DoS protection](https://bitcointalk.org/index.php?topic=136422.0)

 
## Ultraprune + LevelDB
發布版本：Bitcoin-Qt 0.8.0
Ultraprune（區塊鏈高度精簡）是比特幣軟體最重要的升級之一，它解決了驗證區塊鏈上的交易時出現的資源開支問題。比特幣參考軟體之前版本維護著一個針對所有交易輸出（已花的和未花的）的索引，Ultraprune大幅削減了這個索引的體積，因為你只需跟蹤未花交易輸出，一旦一個輸出被花掉了，那麼就可以把這個輸出從這個索引中刪掉。
為了實現這些，開發者實施了一個新的資料庫布局，可以把未花交易輸出壓縮進一個緊湊的格式中，可以減少驗證工作所需訊息的體積。
為了進一步最佳化這個系統的性能，LevelDB和Ultraprune一起被引入，老的BDB資料庫技術被棄用。整體上來看，影響是顯著的：雖然不同硬體上體驗到的改善程度會有所不同，但是驗證區塊數據時的性能至少提高了一個數量級。這個新資料庫結構也為以後對比特幣全節點的精簡和輕量化鋪平了道路。
 
### 更多相關訊息請參考：
* [Bitcoin-Qt 0.8.0 Release notes](https://bitcoin.org/en/release/v0.8.0#improvements)
* [Ultraprune in plain english](https://archive.is/bUocJ)
* [Ultraprune merged in mainline](https://bitcointalk.org/index.php?topic=119525.0)
* [Pruning in the reference client: ultraprune mode](https://bitcointalk.org/index.php?topic=91954.0)

 
## 並行腳本驗證
發布版本：Bitcoin-Qt 0.8
雖然並不是一個很明顯的改變，讓交易腳本驗證過程更加並行化，可以使區塊驗證時間顯著減少。早期版本軟體在接收每個UTXO（未花的交易輸出）之間驗證來自交易輸入的數據，這造成了一種性能問題，因為所有這些動作都是線性地進行的。這有違高效計算過程的設計原則，高效計算原則規定：在可能的情況下，計算應該和I/O工作同時進行。因此，區塊驗證機制被重新設計，以便能平行處理腳本檢驗，這樣一來，在用戶端從區塊接收到所有UTXO前驗證就能開始進行。為了實現這功能，腳本檢查動作在交易被處理後，被放置進一個隊列中，與其它輸入驗證的相關操作分開進行。
這樣最佳化的結果是，節點資源的使用更高效，區塊同步更快。根據測試結果來看，對比以前版本，這個版本的速度提高了35%到100%。
 
### 更多相關訊息請參考：
* [Parallel script verification #2060](https://github.com/bitcoin/bitcoin/pull/2060)

 
## 區塊頭先同步
發布版本：Bitcoin Core 0.10
為了進一步改善初始區塊下載時間，Core項目在2014年末，重新設計了讓節點與包含最大工作量的區塊鏈進行同步的機制的架構。
起初，一個新比特幣用戶端的引導進程會和用戶從一個單節點處接收區塊這個過程牽扯到一起，這產生的結果是，連接的中斷或連接質量下降會拖累這個進程。隨著區塊鏈體積的不斷增加，這個問題有時會讓同步時間顯著增加，有很多用戶報告完成這個啟動過程的時間長達幾天。
區塊頭先同步完全解決了這個問題，它讓節點從一個單節點處先下載、驗證區塊頭數據，然後並行地從其它多個節點處接收區塊數據。
自比特幣早期時期以來，人們就對初始區塊下載時間有抱怨。採用了區塊頭先同步這種方式，這個軟體對新用戶來說實用性又向前跨越了一步。現在用戶可以利用整個網路的節點，大幅削減引導時間，而不必在不穩定的同步中浪費大量時間。這方面對比特幣的長期可持續性來說是非常重要的，開發者使用更精妙的算法漸進式地在這方面進行改善。
 
### 更多相關訊息請參考：
* [Bitcoin-Qt 0.10.0 Release notes](https://bitcoin.org/en/release/v0.10.0#faster-synchronization)

* [Bitcoin.org Developer Guide](https://bitcoin.org/en/developer-guide#headers-first)

* [Pieter Wuille’s post to the Bitcoin-dev mailing list](https://lists.linuxfoundation.org/pipermail/bitcoin-dev/2014-October/006724.html)
 
## 區塊文件修剪
發布版本：Bitcoin Core 0.11
 
舊數據修剪這個概念在中本聰的白皮書中首次被提出，可解決磁碟空間稀缺問題。不幸的是，這個原始設計是不充分的，無法按照比特幣創始人的想像來實施。7年之後，區塊鏈體積達到100多G，區塊文件修剪功能對資源有限的用戶來說幫助是巨大的。
區塊文件修剪使用了上文提到的ultraprune功能；用戶現在可以把288區塊之前的舊區塊數據刪去，這些節點依然可以保護完整的UTXO集，也能夠驗證未花輸出數據，這足以完整地驗證新區塊並對抗潛在的雙花。
這樣的區塊文件修剪假設仍舊會有足夠多的存檔節點來維護完整區塊鏈數據。另一方面，這種創新讓維護節點的成本更低從而擴大了驗證者的範圍。總而言之，這種解決方案有助於我們維持網路多樣化。
 
### 更多相關訊息請參考：
* [Bitcoin-Qt 0.11.0 Release notes](https://bitcoin.org/en/release/v0.11.0#block-file-pruning)

 
## libsecp256k1
發布版本：Bitcoin Core 0.12
根據測算確定，在解決區塊鏈下載的效率問題後要處理的問題是交易驗證瓶頸及其繁重的計算負載。為了解決問題，Core項目啟用了專門最佳化ECDSA操作的新代碼庫。ECDSA（橢圓曲線數位簽章算法）是比特幣公鑰架構的基石，用戶每次用私鑰簽名來發送比特幣時都要用到它。這些簽名需要被這個網路中的每個節點驗證，以保證這個帳簿的完整性。
早期開發者很早就考慮脫離原始的OpenSSL庫，在經過5年時間的設計考量、測試、同級評審後，Pieter Wuille的libsecp256k1庫被引入作為它的替代者。就像預期的那樣，庫被替換後每個比特幣交易的簽名驗證速度明顯加快。對比報告顯示，比使用OpenSSL快了5～7倍。對於用戶來說就是，引導時間最多能節省一半，而這些時間都是從ECSDA操作過程中省出來的，它是從零開始同步新節點最繁重的步驟之一。
比特幣交易活動日益增多，這次更新為網路節點維持合理的用戶體驗至關重要。而且，降低了算法複雜性，讓用戶更高效地利用他們的資源，為新參與者降低了參與門檻。
 
### 更多相關訊息請參考：
* [Bitcoin-Qt 0.12.0 Release notes](https://bitcoin.org/en/release/v0.12.0#signature-validation-using-libsecp256k1)
* [Andrew Poelstra (andytoshi) on security and testing of - libsecp256k1](https://bitcointalk.org/index.php?action=profile;u=80376)
* [Greg Maxwell on testing of libsecp256k1 revealing bug in OpenSSL](https://www.reddit.com/r/Bitcoin/comments/2rrxq7/on_why_010s_release_notes_say_we_have_reason_to/)
* [Greg Maxwell presentation at DevCore](https://www.youtube.com/watch?v=RguZ0_nmSPw&t=1297)
* [Hal Finney post on libsecp256k1](https://bitcointalk.org/index.php?topic=3238.0)

 
## 記憶體池限制
發布版本：Bitcoin Core 0.12
比特幣長期存在的一個薄弱點是它不能恰當地應對節點記憶體池的泛洪（flooding），攻擊者可以發起大量的低價值、低交易費的交易，這些交易在記憶體池中堆積，直到記憶體超載。這會導致記憶體資源不足的節點在非正常活動時期陷入崩潰。對抗這種攻擊的唯一有效措施是提高軟體的最小中繼費，不過此時記憶體池的體積依然沒有上限。
為了解決這個問題，Core項目開發者設置了記憶體池最大體積限制，使用具體的篩選政策，把交易按照交易費排序，並首先把費用較小的交易給篩選出去。為了避免相同費用的交易重新進入記憶體池，節點會提高它的有效最小中繼費率，把初始的最小中繼費率加上最後被篩選出去的交易的交易費。
用戶可以自己設置最大體積，默認設置是300MB，這次更新為資源有限的節點用戶帶來更健壯的體驗，整體上來說，讓整個網路更加可靠。
 
### 更多相關訊息請參考：
* [Bitcoin-Qt 0.12.0 Release notes](https://bitcoin.org/en/release/v0.12.0#memory-pool-limiting)

 
在第二部分，我們將討論更多基於上述技術的最佳化和對網路健壯性、擴展潛力的改善。
