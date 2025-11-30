---
type: posts
layout: post
lang: zh_TW
name: hybrid-spv
id: zh_TW-hybrid-spv
title: 技術路線圖 - 使用完整區塊 SPV 模式的優先區塊下載
permalink: /zh_TW/2017/03/31//prioritized-block-download-using-hybrid-spv/
version: 1
excerpt: 使用完整區塊 SPV 模式的優先區塊下載
---

## 混合式完整區塊 SPV 模式

阻礙一般使用者進一步採用完全驗證軟體的主要障礙之一是，在完全同步整個區塊鏈之前無法使用客戶端的錢包功能。對於引導新節點的使用者來說，這意味著他們無法接收或傳送交易，直到每個區塊都已下載並驗證到鏈的當前尖端。這種行為並非錯誤：Bitcoin Core 參考軟體預設構建為提供 Bitcoin 使用者可以期望的最強安全性和私密性保證，這必然意味著完全驗證以確認歷史區塊鏈資料的完整性。

另一方面，軟體的現有功能（例如標頭優先驗證）提供了改進錢包可用性的機會，前提是使用者願意做出臨時的安全權衡。使用混合式完整區塊 SPV 模式，軟體將根據使用者錢包中最舊的金鑰優先下載區塊。與先前下載的區塊標頭鏈一起，它應該滿足預期的工作量證明難度檢查，客戶端可以立即開始處理相關交易。整個區塊鏈仍然並行下載並最終驗證，但此功能使使用者能夠在背景中進行同步時查看和花費與其錢包相關的 UTXO。

與 SPV 錢包的典型實作相反，此模型不會遭受依賴布隆過濾器和公鑰公開披露的方案所施加的[私密性降級](http://bitcoin.stackexchange.com/questions/37756/are-public-keys-and-their-corresponding-hash-values-both-added-to-a-bitcoinj-blo)。這種好處帶來的權衡是它消耗更多頻寬。另一個警告：在 SPV 模式下收到的確認本質上不如在完全驗證下收到的確認安全。利用混合式 SPV 模式的使用者應該等待多次確認（6+），直到他的付款可以被認為是安全的。

### 更多資訊
  * [完整修補程式集 PR](https://github.com/bitcoin/bitcoin/pull/9483)
  * [接收交易的完美私密性](https://bitcoin.org/en/bitcoin-core/features/privacy#perfect-privacy-for-received-transactions)

