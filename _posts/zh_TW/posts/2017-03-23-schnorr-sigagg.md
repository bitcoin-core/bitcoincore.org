---
type: posts
layout: post
lang: zh_TW
name: schnorr-sigagg
id: zh_TW-schnorr-sigagg
title: 技術路線圖 - Schnorr 簽章和簽章聚合
permalink: /zh_TW/2017/03/23/schnorr-signature-aggregation/
version: 1
excerpt: Schnorr 簽章和簽章聚合的狀態和說明
---

## Schnorr 簽章

用更有效的 Schnorr 演算法替代 Bitcoin 的數位簽章演算法（ECDSA）長期以來一直位於許多 Bitcoin 開發者願望清單的首位。Schnorr 是一種利用橢圓曲線密碼學的簡單演算法，在保留其所有功能和安全假設的同時，實現了對現有方案的多項改進。

值得注意的是，Schnorr 簽章支援「原生多重簽章」，可將多個簽章聚合為一個對其各自輸入的金鑰總和有效的簽章。此功能提供三個重要好處：

* 無論多重簽章設定中的參與者數量如何，簽章大小都是恆定的。50-of-50 交易的大小實際上與使用單個公鑰和簽章的交易相同。因此，此類方案的效能透過消除單獨驗證每個簽章的原始要求而顯著提高。此外，Schnorr 簽章的驗證比 ECDSA 稍快。


* 要驗證和透過網路傳輸的資料大小減少也轉化為有趣的容量增益。考慮到下面顯示的多重簽章交易數量的歷史增長，減少這些交易大小的潛力是對現有擴展工作的誘人補充。我們應該預期隨著支付通道的出現和進一步採用，這一趨勢將繼續。


* 從私密性的角度來看，Schnorr 允許將多重簽章的整個策略隱藏起來，並且與傳統的單個公鑰無法區分。在閾值設定中，參與者也變得不可能透露他們中的哪些人授權或未授權交易。


<p align="center">
  <img src="/assets/images/posts/utxo-repartition.png" alt="UTXO repartition">
</p>
<p align="center">
  根據多重簽章設定分佈未花費 P2SH 輸出。來源：p2sh.info
</p>

不幸的是，與 ECDSA 不同，Schnorr 演算法自發明以來尚未被標準化，可能是因為對其強制執行的原始專利（現已到期）。雖然系統的一般輪廓在數學上是健全的，但缺乏文件和規範使其實作更具挑戰性。具體來說，它對 Bitcoin 的臨時金鑰對設計的應用涉及需要進一步最佳化的安全考慮。

Pieter Wuille 在他的 Scaling Bitcoin Milan 上關於 Schnorr 簽章的演講中將主要挑戰定義為「抵消」問題。一組使用者建立對其金鑰總和有效的簽章的可能性為對抗性參與者從整體中減去另一個使用者的金鑰開啟了大門。它本質上是這樣運作的：

假設使用輸入公鑰 Q1 和 Q2 的 2-of-2 多重簽章方案。惡意參與者在互動階段期間，不是宣布他們的金鑰為 Q2 以與 Q1 組合，而是可以提供 Q2-Q1 並有效地抵消其他使用者的金鑰。傳送到聯合公鑰的任何資金現在只能由 Q2 金鑰的所有者花費，而 Q1 的所有者甚至不知道發生了什麼。

幸運的是，現在有一個解決方案可用，它涉及在簽章之前將設定期間使用的每個金鑰與基於其自身和所有其他相關金鑰的雜湊相乘。此過程稱為去線性化。此方案的安全性證明目前正在進行同行評審，並將在即將發布的白皮書中正式描述。

在短期內，Schnorr 簽章被認為是 Bitcoin 協定兩個重要功能的可行替代品：OP_CHECKSIG 和 OP_CHECKMULTISIG。

前者目前用於根據交易中的訊息檢查 ECDSA 簽章與其各自的公鑰。透過切換到檢查 Schnorr 簽章而不是 ECDSA 的等效物，該操作碼可用於授權需要多個簽章的支出，這通常需要呼叫 OP_CHECKMULTISIG。使用網路無法觀察到的先驗互動，簽署者集合計算組合的公鑰以及共同簽章，由新的 OP_CHECKSIG 驗證，具有增加私密性和降低成本的好處。

後者涉及閾值情況，其中只需要 n-of-m 簽章來授權交易。OP_CHECKMULTISIG 的當前實作驗證閾值策略所需的所有公鑰和相關簽章。因為計算隨著參與者數量線性擴展，Schnorr 提出了一個更有效的方案，它用單個組合簽章以及所需公鑰的子集替換簽章列表。

在對保護簽署者免受惡意行為者攻擊的去線性化方案進行更多評估之前，Schnorr 簽章的進一步應用可能為時過早，但上述功能的實作有望為在生產中更好地理解該方案鋪平道路。取決於額外的同行評審，Schnorr 簽章實作的 BIP 可以在年底前提出。

## 簽章聚合

Schnorr 允許在單個輸入上組合多個簽章的屬性也適用於所有交易的多個輸入的聚合。Bitcoin 開發者 Gregory Maxwell 是第一個使用基於 BLS 簽章的先前提案的見解引入這個想法的人。

要正確理解此應用與上述應用之間的差異，有必要考慮在每個各自情況下如何聚合簽章。在原生多重簽章設定中，簽署者彼此合作以計算共同公鑰及其相關簽章。這種互動發生在協定之外，只涉及相關各方。簽章聚合背後的想法是讓系統驗證者（即 Bitcoin 節點）在協定層級為所有交易的每個輸入計算單個金鑰和簽章。

因為此方案將聚合範圍擴展到確定性參與者集合之外，它引入了惡意行為者利用「抵消」錯誤的新攻擊向量。因此，前一節中強調的去線性化修正對於此方法的健全性至關重要。

在實作方面，該提案相當直接：修改 OP_CHECKSIG 和 OP_CHECKMULTISIG，使它們可以堆疊公鑰，對它們進行去線性化，一旦驗證了所有相關輸入，就為它們各自的交易產生組合簽章。

評估如果自創世區塊以來實作簽章聚合可能節省的資源類型相當簡單。假設每個歷史簽章減少到 1 位元組，除了每筆交易一個，分析表明該方法將導致儲存和頻寬至少減少 25%。增加使用 n-of-n 閾值可能會轉化為更多節省，儘管它們未在此分析中考慮。

<p align="center">
  <img src="/assets/images/posts/signature-agg-chart.png" alt="Schnorr signature addregation savings chart">
</p>

### 關於 Schnorr 簽章的更多資訊
  * [Pieter Wuille 在 Scaling Bitcoin Milan 上關於 Schnorr 簽章的演講逐字稿](https://diyhpl.us/wiki/transcripts/scalingbitcoin/milan/schnorr-signatures/)
  * [Pieter Wuille 的演講影片](https://youtu.be/_Z0ID-0DOnc?t=2297)
  * [2016 年 7 月 Bitcoin 開發者和礦工會議關於 Schnorr 簽章的討論](http://diyhpl.us/wiki/transcripts/2016-july-bitcoin-developers-miners-meeting/dan-boneh/)
  * [Schnorr 文件](https://github.com/sipa/secp256k1/blob/968e2f415a5e764d159ee03e95815ea11460854e/src/modules/schnorr/schnorr.md)
  * [StackExchange - Schnorr 的含義是什麼？](http://bitcoin.stackexchange.com/questions/34288/what-are-the-implications-of-schnorr-signatures/35351#35351)
  * [Elements 專案：Schnorr 簽章驗證](https://elementsproject.org/features/schnorr-signatures)
  * [SF Bitcoin Devs 研討會：Gregory Maxwell 談 Schnorr 多重簽章](https://www.youtube.com/watch?v=TYQ-3VvNCHE)
  * [Bitcoin Core 開發者關於 Schnorr 簽章的會議記錄](https://bitcoincore.org/logs/2016-05-zurich-meeting-notes.html)

### 關於簽章聚合的更多資訊
  * [Gregory Maxwell 在 Bitcointalk.org 論壇上關於簽章聚合的貼文](https://bitcointalk.org/index.php?topic=1377298.0)
  * [Bitcoin Core 開發者關於 Schnorr 簽章的會議記錄](https://bitcoincore.org/logs/2016-05-zurich-meeting-notes.html)
  * [Pieter Wuille 在 Scaling Bitcoin Milan 上關於 Schnorr 簽章的演講逐字稿](https://diyhpl.us/wiki/transcripts/scalingbitcoin/milan/schnorr-signatures/)
  * [Pieter Wuille 的演講影片](https://youtu.be/_Z0ID-0DOnc?t=2297)


