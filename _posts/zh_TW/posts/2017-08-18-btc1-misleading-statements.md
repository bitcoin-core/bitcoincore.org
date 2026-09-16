---
type: posts
layout: post
lang: zh_TW
name: btc1-misleading-statements
id: zh_TW-btc1-misleading-statements
title: 更正關於 Segwit2x 和 btc1 的錯誤資訊
permalink: /zh_TW/2017/08/18/btc1-misleading-statements/
version: 1
excerpt: 更正關於 Segwit2x 和 btc1 的錯誤資訊
---

「Segwit2x」，一項對 Bitcoin 網路共識規則進行不相容變更的提案，最近受到越來越多的關注。有人試圖誤導人們相信 btc1 專案（Segwit2x 提案的實作）是現有軟體的必要更新──事實並非如此。相反，它是對現有網路規則的爭議性偏離，其使用者將很快發現自己與網路其他部分在區塊和交易的有效性方面存在分歧。

請注意：

  * 隔離見證（或 Segwit，一種將在未來幾天內啟動的軟分叉）與 Segwit2x 硬分叉無關。隔離見證與所有先前的 Bitcoin 軟體向後相容。對於絕大多數 Bitcoin 使用者來說，不需要採取任何行動。

  * [bitcoincore.org](https://bitcoincore.org) 是官方網站，[@bitcoincoreorg](https://twitter.com/bitcoincoreorg) 是 Bitcoin Core 專案的官方 Twitter 帳戶。任何其他聲稱代表該專案的網站或 Twitter 帳戶都是欺詐性的。Bitcoin Core 是一個開源專案，歡迎任何人透過其 GitHub 專案貢獻和審查。Bitcoin Core 二進位檔案可以從 [bitcoincore.org](/zh_TW/download) 取得，並且始終由發布經理的簽署金鑰進行數位簽章。撰寫本文時，Bitcoin Core 的最新版本是 0.14.2。

  * btc1 與 Bitcoin Core *沒有任何* 關聯。沒有常規的 Bitcoin Core 貢獻者支援 btc1 或與該專案有任何聯繫，也沒有參與其提議的硬分叉的設計。

  * 我們強烈建議使用者不要下載任何聲稱是對 Bitcoin 共識規則「升級」的 Bitcoin 全節點軟體，除非仔細考慮提議的變更對 Bitcoin 系統的影響以及社群對它的支援程度。這包括 Bitcoin Core 新版本中提議的共識變更。

  * 雖然很難確定更廣泛的 Bitcoin 社群支援什麼，但要警惕聲稱龐大而多樣化的 Bitcoin 社群在沒有獨立驗證的情況下完全轉向一個分叉或另一個分叉的說法。簽署信已被公司用來聲稱代表其客戶/使用者而未經其同意，並且經常使用不精確和誤導性的語言。過去，Bitcoin XT、Bitcoin Classic 和 Bitcoin Unlimited 以及其他的信件已被傳閱以表明對一個想法的普遍支援，同時被宣傳為不考慮社群考慮而執行軟體的承諾，只會在幾個月後被放棄。

  * Bitcoin Core 貢獻者和 Bitcoin 社群成員提出的關於 Segwit2x 提案的擔憂尚未得到其支持者的充分解決。該提案的細節是在 Bitcoin 的隔離見證啟動之前以及最近建立 BCH 貨幣之前確定的。在規劃未來時忽略這些事件的結果是不負責任的。例如，我們已經看到了當單個地址在兩條鏈上都有效時產生的混亂，但 Segwit2x 提案打算重複同樣的錯誤。此外，BCH 對強重放保護的實作為 BCH 和 Bitcoin 的使用者提供了顯著的保護，而 Segwit2x 不打算提供這種保護。

  * Bitcoin 的共識規則應該只在謹慎的情況下並在整個社群的廣泛同意下變更。Segwit2x 在其流程和實作中都受到了許多人的反對。Bitcoin Core 將繼續支援 Segwit 軟分叉，我們期待在未來幾年幫助 Bitcoin 擴展到新的高度。
