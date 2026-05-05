---
type: posts
layout: post
lang: zh_TW
name: segwit-benefits
id: zh_TW-segwit-benefits
title: 隔離見證的好處
permalink: /zh_TW/2016/01/26/segwit-benefits/
version: 1
excerpt: 本頁總結了隔離見證的一些好處。
---
{% include toc.html %}

隔離見證軟分叉（segwit）包含了廣泛的功能，其中許多功能具有高度技術性。本頁總結了這些功能的一些好處。

## 可塑性修復

Bitcoin 交易透過一個 64 位十六進位雜湊來識別，稱為交易識別碼（txid），該識別碼基於正在花費的代幣和誰將能夠花費交易的結果。

不幸的是，txid 的計算方式允許任何人對交易進行小幅修改，這不會改變其含義，但會改變 txid。這稱為第三方可塑性。BIP 62（「處理可塑性」）試圖以零碎的方式解決這些問題，但實作起來過於複雜而無法作為共識檢查，已被撤回。

例如，您可以向網路提交一個 txid 為 ef74...c309 的交易，但發現第三方（例如網路上中繼您交易的節點，或將您的交易包含在區塊中的礦工）稍微修改了交易，導致您的交易仍然花費相同的代幣並支付相同的地址，但在完全不同的 txid 683f...8bfa 下確認。

更一般地說，如果交易的一個或多個簽名者修改他們的簽名，則交易仍然有效並向相同地址支付相同金額，但 txid 會完全改變，因為它包含了簽名。簽名資料（但不是輸出或輸入選擇）的變更修改交易的一般情況稱為 scriptSig 可塑性。

Segwit 透過允許 Bitcoin 使用者將交易的可塑部分移至*交易見證*中，並隔離該見證，使見證的變更不影響 txid 的計算，從而防止第三方和 scriptSig 可塑性。

### 誰受益？

- **追蹤已花費 bitcoin 的錢包作者**：最簡單的監控自己傳出交易狀態的方法是透過 txid 查詢它們。但在具有第三方可塑性的系統中，錢包必須實作額外的程式碼才能處理變更的 txid。

- **任何花費未確認交易的人**：如果 Alice 在交易 1 中向 Bob 付款，Bob 在交易 2 中使用該付款向 Charlie 付款，然後 Alice 的付款被可塑化並以不同的 txid 確認，那麼交易 2 現在無效，Charlie 沒有收到付款。如果 Bob 值得信任，他將重新發出向 Charlie 的付款；但如果他不值得信任，他可以簡單地為自己保留這些 bitcoin。

- **閃電網路**：隨著第三方和 scriptSig 可塑性被修復，閃電網路的實作變得不那麼複雜，並且在區塊鏈上的空間使用上顯著更有效率。移除 scriptSig 可塑性後，也可以執行輕量級閃電客戶端，將監控區塊鏈的工作外包出去，而不是每個閃電客戶端都需要成為完整的 Bitcoin 節點。

- **任何使用區塊鏈的人**：今天的智慧合約（例如微支付通道）和預期的新智慧合約，變得不那麼複雜設計、理解和監控。

注意：只有當隔離見證交易的所有輸入都是 segwit 花費（直接或透過向後相容的 segwit P2SH 地址）時，segwit 交易才能避免可塑性。

### 進一步資訊

 * [Bitcoin Wiki 上的可塑性](https://en.bitcoin.it/wiki/Transaction_Malleability)
 * [Coin Telegraph 關於 2015 年可塑性攻擊的文章](http://cointelegraph.com/news/115374/the-ongoing-bitcoin-malleability-attack)
 * [Bitcoin Magazine 關於 2015 年可塑性攻擊的文章](https://bitcoinmagazine.com/articles/the-who-what-why-and-how-of-the-ongoing-transaction-malleability-attack-1444253640)
 * [「閃電網路所需的 BIP 概述」逐字稿](http://diyhpl.us/wiki/transcripts/scalingbitcoin/hong-kong/overview-of-bips-necessary-for-lightning/)
 * [BIP 62](https://github.com/bitcoin/bips/blob/master/bip-0062.mediawiki)
 * [BIP 140 -- 可塑性修復的替代方法](https://github.com/bitcoin/bips/blob/master/bip-0140.mediawiki)
 * [關於 683f...8bfa 交易的 Stack exchange 回答](http://bitcoin.stackexchange.com/questions/22051/transaction-malleability-in-the-blockchain/22058#22058)

## 簽章操作的線性擴展

增加 Bitcoin 區塊大小的簡單方法的一個主要問題是，對於某些交易，簽名雜湊的擴展是二次方的而不是線性的。

![線性與二次方](/assets/images/linear-quad-scale.png)

本質上，將交易的大小加倍可以使簽名操作的數量加倍，以及為了驗證每個簽名而必須雜湊的資料量加倍。這在實際環境中已經出現，其中單個區塊需要 25 秒來驗證，惡意設計的交易可能需要超過 3 分鐘。

Segwit 透過變更簽名的交易雜湊計算來解決這個問題，使得交易的每個位元組最多只需要雜湊兩次。這更有效地提供了相同的功能，因此即使惡意生成或支援更大的區塊（因此更大的交易），也可以生成大型交易而不會遇到簽名雜湊問題。

### 誰受益？

移除驗證簽名的雜湊資料的二次方擴展使增加區塊大小更安全。在不限制交易大小的情況下執行此操作，允許 Bitcoin 繼續支援向大型群組付款或來自大型群組的付款，例如礦池獎勵的付款或群眾募資服務。

修改後的雜湊僅適用於從見證資料發起的簽名操作，因此來自基礎區塊的簽名操作將繼續需要較低的限制。

### 進一步資訊

 * [BIP 143](https://github.com/bitcoin/bips/blob/master/bip-0143.mediawiki)
 * [Rusty Russell 關於 25 秒交易的部落格文章](http://rusty.ozlabs.org/?p=522)
 * [Bitcoin wiki 上的 CVE 2013-2292](https://en.bitcoin.it/wiki/Common_Vulnerabilities_and_Exposures#CVE-2013-2292)
 * [將交易限制為 100kB 的提案](https://lists.linuxfoundation.org/pipermail/bitcoin-dev/2015-July/009494.html)
 * [Bitcoin Classic 在 0.11.2 分支上的提交，新增了對 sighash 位元組的額外共識限制](https://github.com/bitcoinclassic/bitcoinclassic/commit/842dc24b23ad9551c67672660c4cba882c4c840a)

## 透過 pay-to-script-hash (P2SH) 增加多重簽名的安全性

多重簽名付款目前使用 P2SH，它由 160 位元 HASH160 演算法（SHA256 的 RIPEMD）保護。然而，如果其中一個簽名者希望竊取所有資金，他們可以在作為多重簽名腳本一部分的有效地址與只向他們支付所有資金的腳本之間找到碰撞，只需 80 位元（2<sup>80</sup>）的工作量，這對於資源極其豐富的攻擊者來說已經在可能的範圍內。（作為比較，在持續 1 exahash/秒的情況下，Bitcoin 挖礦網路每兩週完成 80 位元的工作量）

Segwit 透過僅對直接向單個公鑰的付款使用 HASH160（在這種情況下這種攻擊是無用的），同時對向腳本雜湊的付款使用 256 位元 SHA256 雜湊來解決這個問題。

### 誰受益？

透過 segwit 向多重簽名或智慧合約付款的每個人都受益於為腳本提供的額外安全性。

### 進一步資訊

 * [Gavin Andresen 詢問 80 位元攻擊是否值得擔心](https://lists.linuxfoundation.org/pipermail/bitcoin-dev/2016-January/012198.html)
 * [Ethan Heilman 描述循環查找演算法](https://lists.linuxfoundation.org/pipermail/bitcoin-dev/2016-January/012202.html)
 * [Rusty Russell 計算執行攻擊的成本](https://lists.linuxfoundation.org/pipermail/bitcoin-dev/2016-January/012227.html)
 * [Anthony Towns 應用循環查找演算法來利用交易](https://lists.linuxfoundation.org/pipermail/bitcoin-dev/2016-January/012218.html)
 * [Gavin Andresen 總結討論串](https://lists.linuxfoundation.org/pipermail/bitcoin-dev/2016-January/012234.html)

## 腳本版本控制

對 Bitcoin 腳本的變更允許改進安全性和功能。然而，腳本的設計僅允許透過用新的操作碼替換十個額外的 OP\_NOP 操作碼之一來實作向後相容（軟分叉）的變更，該新操作碼可以有條件地使腳本失敗，但在其他情況下什麼都不做。這對於許多變更來說是足夠的 -- 例如引入新的簽名方法或像 OP\_CLTV 這樣的功能，但它既有點 hacky（例如，OP\_CLTV 通常必須伴隨 OP\_DROP），也無法用於啟用甚至像連接兩個字串這樣簡單的功能。

Segwit 透過為腳本包含版本號來解決這個問題，因此在非 segwit 交易中需要硬分叉才能使用的額外操作碼，可以透過簡單地增加腳本版本來支援。

### 誰受益？

更容易變更腳本操作碼將使 Bitcoin 中的進階腳本更容易。這包括引入 Schnorr 簽名、使用金鑰恢復來縮小簽名大小、支援側鏈，或使用 Merklized Abstract Syntax Trees (MAST) 和其他研究級想法來建立更智慧的合約等變更。

## 減少 UTXO 增長

每個驗證的 Bitcoin 節點都維護未花費交易輸出（UTXO）資料庫，以確定新交易是有效還是欺詐。為了網路的有效運作，此資料庫需要非常快速地查詢和修改，並且理想情況下應該能夠放入主記憶體（RAM），因此將資料庫的大小（以位元組為單位）保持在盡可能小的範圍內是有價值的。

隨著 Bitcoin 的成長，這變得更加困難，因為每個新使用者必須至少有一個自己的 UTXO 條目，並且將偏好擁有多個條目以幫助改善他們的隱私和靈活性，或提供作為支付通道或其他智慧合約的支援。

Segwit 透過使不影響 UTXO 集大小的簽名資料成本比影響 UTXO 集大小的資料低 75% 來改善這種情況。預期這將鼓勵使用者偏好使用對 UTXO 集影響最小的交易以最小化費用，並鼓勵開發者以同樣最小化對 UTXO 集影響的方式設計智慧合約和新功能。

因為 segwit 是軟分叉變更並且不增加基礎區塊大小，所以 UTXO 集的最壞情況增長率保持不變。

### 誰受益？

減少的 UTXO 增長將使礦工、企業和執行完整節點的使用者受益，這反過來又有助於在更多使用者進入系統時維持 Bitcoin 網路的當前安全性。幫助最小化 UTXO 集增長的使用者和開發者將受益於較低的費用，相對於那些忽視其交易對 UTXO 增長影響的人。

### 進一步資訊

 * [Statoshi UTXO 儀表板](http://statoshi.info/dashboard/db/unspent-transaction-output-set)

## 不驗證簽名時的效率提升

歷史交易的簽名可能不如未來交易的簽名有趣 -- 例如，Bitcoin Core 預設不檢查最近檢查點之前的交易簽名，而某些 SPV 客戶端根本不檢查簽名本身，信任這已經由礦工或其他節點完成。然而，目前簽名資料是交易的組成部分，必須存在才能計算交易雜湊。

隔離簽名資料允許對簽名資料不感興趣的節點將其從磁碟中剪除，或避免首先下載它，從而節省資源。

### 誰受益？

隨著更多交易使用 segwit 地址，執行剪除或 SPV 節點的人將能夠以更少的頻寬和磁碟空間運作。

## 區塊容量/大小增加

由於舊節點只會下載去除見證的區塊，它們只對該資料強制執行 1 MB 區塊大小限制規則。
因此，理解包含見證資料的完整區塊的新節點可以用新的限制取代這個限制，允許更大的區塊大小。隔離見證因此利用這個機會將區塊大小限制提高到接近 4 MB，並新增了新的成本限制以確保區塊在其資源使用上保持平衡（這有效地導致有效限制接近 1.6 到 2 MB）。

### 誰受益？

執行升級錢包的人將能夠透過將簽名移至交易的見證部分來利用增加的區塊大小。

## 邁向單一組合區塊限制

目前對區塊大小有兩個共識強制執行的限制：區塊不能大於 1MB，並且獨立地，不能在區塊中的交易之間執行超過 20,000 次簽名檢查。

在給定單一限制的情況下，找到要包含在區塊中的最有利可圖的交易集是背包問題的一個實例，可以用簡單的貪婪演算法輕鬆地幾乎完美地解決。然而，新增第二個約束使得在某些情況下找到好的解決方案變得非常困難，而且這個理論問題已在實踐中被利用，迫使區塊以遠低於容量的大小進行挖礦。

在不進行硬分叉或大幅減少區塊大小的情況下，無法解決這個問題。由於 segwit 無法修復這個問題，它選擇不使其惡化：特別是，與其為隔離見證資料引入獨立限制，不如將單一限制應用於 UTXO 資料和見證資料的加權總和，允許兩者作為組合實體同時受到限制。

### 誰受益？

如果未來的硬分叉將區塊容量限制變更為參數的單一加權總和，最終礦工將受益。例如：

    50*sigops + 4*basedata + 1*witnessdata < 10M

這讓礦工輕鬆準確地填充區塊，同時最大化費用收入，這將透過允許使用者更可靠地計算其交易被挖掘所需的適當費用來使使用者受益。

### 進一步資訊

 * [背包問題](https://en.wikipedia.org/wiki/Knapsack_problem)
 * [2015 年 8 月 bitcointalk 上的 Sigop 攻擊討論](https://bitcointalk.org/index.php?topic=1166928.0;all)
 * [Gregory Maxwell 在 bitcoin-dev 上關於見證限制](https://lists.linuxfoundation.org/pipermail/bitcoin-dev/2015-December/011870.html)
 * [「驗證成本指標」逐字稿](http://diyhpl.us/wiki/transcripts/scalingbitcoin/hong-kong/validation-cost-metric/)

## 2016-10-19 更新

本頁的早期版本將「緊湊欺詐證明」列為 segwit 的好處。然而，如實作的那樣，segwit 並沒有使這變得更容易：無論有沒有 segwit，啟用緊湊欺詐證明及其帶來的好處的未來軟分叉都需要包含自己的承諾（例如，在 coinbase 交易中），而不是能夠擴展 segwit 使用的承諾資料。

先前的文字是：

> **緊湊欺詐證明**
>
> 隨著 Bitcoin 使用者群的擴大，驗證整個區塊鏈自然變得更加昂貴。為了維持 Bitcoin 的去中心化、無需信任的特性，重要的是允許那些無法負擔驗證整個區塊鏈的人至少能夠以低成本驗證他們能夠負擔的範圍。
>
> Segwit 透過允許未來的軟分叉擴展見證結構以包含承諾資料來改善這種情況，這將允許輕量級（SPV）客戶端強制執行共識規則，例如區塊中引入的 bitcoin 數量、區塊的大小以及區塊中使用的 sigops 數量。
>
> **誰受益？**
>
> 欺詐證明允許 SPV 使用者幫助強制執行 Bitcoin 的共識規則，這將潛在地大大增加整個 Bitcoin 網路的安全性，並減少個別使用者可能受到攻擊的方式。
>
> 這些欺詐證明可以作為未來軟分叉的一部分新增到見證資料結構中，它們將幫助 SPV 客戶端強制執行規則，即使是那些不使用 segwit 功能的交易。

## 2020-06-23 更新

本頁的早期版本將「輸入值的簽名」列為 segwit 的好處。
然而，如實作的那樣，segwit 並沒有使這變得安全：
無論有沒有 segwit，都需要未來的軟分叉才能依賴簽名的輸入值。

由於每個輸入的值都是單獨簽名的，因此可以以欺騙性的方式操縱表面費用。
(CVE-2020-14199)

先前的文字是：

> **輸入值的簽名**
>
> 當硬體錢包簽署交易時，它可以輕鬆驗證總花費金額，但只能透過擁有所有正在花費的輸入交易的完整副本來安全地確定費用，並且必須雜湊每個交易以確保它沒有被餵以虛假資料。由於單個交易的大小最大可達 1MB，因此即使被簽署的交易本身很小，這也不一定是便宜的操作。
>
> Segwit 透過明確雜湊輸入值來解決這個問題。這意味著硬體錢包可以簡單地被給予交易雜湊、索引和值（並被告知使用了什麼公鑰），並且可以安全地簽署花費交易，無論被花費的交易有多大或多複雜。
>
> **誰受益？**
>
> 硬體錢包的製造商和使用者是明顯的受益者；然而，這也可能使在小型嵌入式裝置中安全使用 Bitcoin 進行「物聯網」應用變得更加容易。
>
> 此好處僅在花費發送到 segwit 啟用地址（或 segwit-via-P2SH 地址）的交易時可用。
>
> **進一步資訊**
>
>  * [BIP 143](https://github.com/bitcoin/bips/blob/master/bip-0143.mediawiki)
