---
layout: post
lang: zh_TW
name: segwit-benefits
id: zh_tw-segwit-benefits
title: 隔離見證的好處
permalink: /zh_TW/2016/01/26/segwit-benefits/
version: 1
excerpt: 隔離見證軟分叉（segwit）包括很多的特性，其中許多是技術性很強的。
---
{% include toc.html %}

隔離見證軟分叉（segwit）包括很多的特性，其中許多是技術性很強的。 此頁麵總結了這些特性的一些好處。

## 修複延展性問題

比特幣交易是由一串 64 位的十六進製雜湊交易ID（TxID）標識的，這是基於交易中幣的來源和幣的接收者確定的。

不幸的是,txid 的計算方法允許任何人可以對交易做小的修動，雖然不會改變交易的內容，但會改變 txid。這就是所謂的第三方延展性。
BIP62 （"處理可塑性"）試圖以一些方式解決這些問題，但是它太複雜，以至於無法實現為共識檢查,所以已經被放棄了。

例如，您可以發送 txid 為 ef74... C309 到比特幣網路，當網路中的節點中繼這筆交易，或者礦工打包交易到區塊中的時候，它們可以輕微修改這筆交易，
導緻您的交易仍然花一樣的幣，並支付到相同的地址，但是以完全不同的txid683f...8bfa出現。

更通俗地說，如果一筆交易的一個或更多個簽名者修改他們的簽名, 那麼交易仍然有效並且支付相同的比特幣，以相同的地址，但這筆交易的txid完全改變，
因為它們包含簽名。 更改簽名數據（但不改變 output 或 input）來修改交易的情況稱為 scriptSig 延展性。

Segwit 可以防止第三方和 scriptSig 延展性, 通過把比特幣交易中的的可修改部分移動到見證交易裏, 並且分離後不影響txid的計算。

### 誰將從中受益?

- **錢包作者用來監控發出比特幣：** 這是最簡單的，隻需要監控發出的txid的狀態就可以。但是在存在第三方延展性問題的係統裏，錢包必須添加額外的程式碼，
以便能夠應對變化的txids。

- **花費未確認的交易：** 如果 Alice 在交易1支付 Bob一些幣，Bob 在交易2 使用收到的付款支付給Charlie，然後Alice的付款發生延展性修改,並用不同的txid確認，
那麼交易2現在是無效的，而Charlie就不會被支付。如果Bob是值得信賴的，他會重新發出一筆交易給查理；但如果他不是，他可以簡單地把這些比特幣留給自己。

- **閃電網路：** 第三方和 scriptSig 延展性問題修複後可以降低閃電網路實現的的複雜性，而且在使用 blockchain 的空間上將更加有效. scriptSig
延展性刪除後，它也可能運行輕量級的lighting客戶端服務去監測區塊鏈,而不是每個lightning客戶端都運行比特幣完整節點。

- **任何使用區塊鏈的人:** 目前的智能契約，比如小額支付通道，預期新的智能契約，將會變得不用那麼複雜的設計，理解和監控。

註意： segwit 交易隻能在當所有input都是segwit交易（直接或經由一個嚮後兼容 的 segwit P2SH 地址）下避免延展性問題。

### 更多信息

 * [比特幣維基延展性](https://en.bitcoin.it/wiki/Transaction_Malleability)
 * [關於延展性攻擊的 2015 年比特幣電報文章](http://cointelegraph.com/news/115374/the-ongoing-bitcoin-malleability-attack)
 * [關於延展性攻擊的2015年比特幣雜誌文章](https://bitcoinmagazine.com/articles/the-who-what-why-and-how-of-the-ongoing-transaction-malleability-attack-1444253640)
 * [概述各種比特幣改進建議對閃電交易的重要性](http://diyhpl.us/wiki/transcripts/scalingbitcoin/hong-kong/overview-of-bips-necessary-for-lightning/)
 * [BIP 62](https://github.com/bitcoin/bips/blob/master/bip-0062.mediawiki)
 * [比特幣改進建議 140-延展性修複的替代方法](https://github.com/bitcoin/bips/blob/master/bip-0140.mediawiki)
 * [關於 683f…8bfa 交易在 Stack exchange 問答網路平臺的回答](http://bitcoin.stackexchange.com/questions/22051/transaction-malleability-in-the-blockchain/22058#22058)

## 線性增長 sighash 的操作

用簡單的方法來增加比特幣區塊大小的一個主要問題是，對於某些交易，簽名散列增長是平方增長的, 而不是線性增長的

![Linear versus quadratic](/assets/images/linear-quad-scale.png)

在本質上，一個交易的大小增加一倍,簽名操作的個數也增加一倍，以及那些進行驗證簽名需要雜湊的數據也應該增加一倍. 但曾經已經出現過，一個單獨的塊需要 25 秒驗證,其他 
一些惡意設計的交易可能需要超過 3 分鍾。

Segwit 通過改變交易雜湊簽名的計算方式可以解決此問題，使得交易的每個字節隻需要至多兩次雜湊。這提供了相同的功能但更有效率，
使得大的交易仍可以産生而不會有簽名雜湊問題，即使有人生成惡意的或更大的塊（並較大的交易）也是支援的。

### 誰將從中受益?

刪除驗證簽名需要的雜湊數據的平方伸縮問題，使增長塊大小更安全。這樣做並冇有 限製交易大小，所以仍然可以繼續支援支付或者接收來自於大的組織的比特幣,
比如挖礦獎勵或衆籌服務。

修改後的雜湊僅適用於從 witness 數據發起簽名操作，所以從舊的區塊發起的簽名操作將繼續需要限製簽名操作數下限。

### 更多信息

 * [BIP 143](https://github.com/bitcoin/bips/blob/master/bip-0143.mediawiki)
 * [在 25s 交易中 Rusty Russell  發布的部落格](http://rusty.ozlabs.org/?p=522)
 * [CVE 2013-2292-比特幣維基](https://en.bitcoin.it/wiki/Common_Vulnerabilities_and_Ex 發布sures#CVE-2013-2292)
 * [提議限製交易至 100KB](https://lists.linuxfoundation.org/pipermail/bitcoin-dev/2015-July/009494.html)
 * [在比特幣 Classic0.11.2 版本中添加了限製sighash字節的額外共識](https://github.com/bitcoinclassic/bitcoinclassic/commit/842dc24b23ad9551c67672660c4cba882c4c840a)

## 值的簽署

當硬體錢包簽署一個交易的時候，它可以很容易地查證花費的總金額，但必須使用每個 input 交易的完整副本來確定是否安全，而且必須計算每個 input 的雜湊以確保它們不是虛假數據。
個別交易大小可高達1MB大小，這不一定是一種廉價的操作，即使被簽名的交易本身是相當小的。

Segwit 使 input 雜湊變的精確從而解決了此問題。這意味著硬體的錢包可以簡單地給出交易雜湊，索引和值（和說明使用什麼樣的公鑰），並可以放心地簽署發出的交易，
無論花費的 input 交易有多大或多複雜。

### 誰將從中受益?

硬體錢包製造商和用戶是明顯的受益者。然而，這也使得它更容易，安全地在小型嵌入式設備的“物聯網”的應用程式中使用比特幣。

當消費收據發給 segwit 地址才可享有此優惠。 

### 更多資料

 * [BIP 143](https://github.com/bitcoin/bips/blob/master/bip-0143.mediawiki)
 
## 通過 multisig（P2SH) 增強的安全性

Multisig 付款目前使用 P2SH， 由 160 位 HASH160 算法（SHA256的RIPEMD）保護。但是，如果其中簽名者，想盜取所有的資金，他們可以找到一個有效的地址之為
multisig 的一部分，隻需支付他們所有的資金隻有80位（280）的工作，這已經是一個極其資源充足的攻擊者可做到的。 （為便於比較，在持續1 exahash /秒，
比特幣挖礦網路每兩個星期發掘80位值）

Segwit 通過使用 HASH160 付款直接到一個公共密鑰（這種攻擊是冇用的），同時採用了256 位的散列SHA256付款給腳本解決此問題。

### 誰受惠?

每個人都支付給 multisig 或通過 segwit 提供的腳本得到額外保全。

### 更多資料

 * [Gavin Andresen詢問 80 位的攻擊是否值得擔憂](https://lists.linuxfoundation.org/pipermail/bitcoin-dev/2016-January/012198.html)
 * [Ethan Heilman 描述發現算法圈](https://lists.linuxfoundation.org/pipermail/bitcoin-dev/2016-January/012202.html)
 * [Rusty Russell 計算進行攻擊的成本](https://lists.linuxfoundation.org/pipermail/bitcoin-dev/2016-January/012227.html)
 * [Anthony Towns 用周期算法剝削交易](https://lists.linuxfoundation.org/pipermail/bitcoin-dev/2016-January/012218.html)
 * [Gavin Andresen 總結](https://lists.linuxfoundation.org/pipermail/bitcoin-dev/2016-January/012234.html)

## 腳本版

更改比特幣的腳本允許提高安全性和改進功能。然而，腳本的設計隻允許嚮後兼容（軟分叉）通過更換其中一個新的操作碼的十個額外 OP_NOP 操作碼能讓腳本失效，
要不就冇有任何可能。這足夠讓許多變化産生 - 如引入一個新的簽名方法或類似 OP_CLTV 功能，但兩者無法單獨操作（例如，OP_CLTV通常必須由一個 OP_DROP 陪同），
並不能用於作為連接兩個字符串。

Segwit 包括腳本版本號解決此，在非segwit交易也需要用一個額外操作碼，通過增加腳本版本支援。

### 誰受惠?

腳本操作碼會使比特幣高級腳本容易化。這包括改變，如引進的 Schnorr 簽名，使用一鍵恢複收縮簽名大小，支援側鏈，或創建使用 Merklized 抽象文法
（MAST） 和其他研究級想法。

## 減少 UTXO 增長

（UTXO）數據庫是由每個驗證 Bitcoin 節點維持，以確定新的交易是否有效或欺詐。為網路的有效運行，該數據庫必須非常快查詢和修改，
並且理想地應當能夠適合主存儲器（RAM），因此保持該數據庫的小是很重要的。

當比特幣增長這就變得睏難，因為每個新用戶必須有自己的至少一個UTXO條目，每個用戶有多個條目，保護隱私和靈活性，或提供為後盾，為支付通路。

Segwit改善了的情況通過簽名數據，這不會影響UTXO大小，成本比原本UTXO組大小的數據少75％。預計這將鼓勵用戶使用，大幅度減少費用，
並鼓勵開發人員設計方式，也將減少對UTXO的影響智能契約和新功能設定UTXO影響。

因為segwit是軟分叉變化並不增加基礎大小，所述UTXO最壞情況下生長速度就保持不變。

### 誰受惠?

減少 UTXO 增長將有利於礦工，企業，完整的節點使用者，有助於維持在Bitcoin網路的安全隨著更多的用戶進入係統的用戶。
用戶和開發幫助減少UTXO的增長將享有較低的費用，相比忽視交易對UTXO增長的影響使用者。

### 更多資料

 * [Statoshi UTXO 排行榜](http://statoshi.info/dashboard/db/unspent-transaction-output-set)
 
## 提防詐欺

隨著比特幣用戶群的擴大，驗證整個鏈接變得更加昂貴。為了保持比特幣的分散，不可信的性質，重要的是讓不能對整個鏈接驗證的用戶至少能便宜地驗證。

Segwit改善了情況並通過允許未來軟叉延長證人結構並包括承諾數據，這將允許輕（SPV）客戶端來執行一緻的規則，例如，如在一並引入比特幣的數量，
並在大量中數目中使用 sigops。

### 誰受惠?

提防欺詐允許SPV用戶幫助執行比特幣的共識規則，這將能增加比特幣網路的整體安全性，以及減少個人用戶被攻擊。

這些欺詐行為的證據可以被添加到證人的數據結構作為未來軟叉的一部分，他們將幫助SPV客戶強製規定，就連交易不使用segwit功能都可。

## 不驗證簽名效率提高

曆史交易簽名可能不比未來簽名有趣 - 例如，比特幣核心最近的檢查點在預設情況下並冇有檢查交易簽名，有的SPV客戶幹脆不檢查數字簽名本身，
相信已經由礦工或其他節點檢查。目前，簽名數據事務處理的一個組成部分，必須存在，以便計算該事務散列。

分離署名數據允許簽名數據從磁盤進行清理，避免下載它，節約資源的節點。

### 誰受惠?
隨著越來越多的事務使用segwit地址，人們或SPV節點將能夠用更少的數據與磁盤空間進行操作。

## 容量增加
舊節點將僅供下載，隻執行對這些數據的1 MB大小限製規則。新的節點，理解見證數據的完整，因此免費用新的來取代這限製，允許更大的數據。
因此隔離見證利用這個機會提高限製為近4 MB，並增加了新的成本限製，以確保留在自己的資源利用平衡（這有效地限製接近1.6至2 MB）。

### 誰受惠?

誰運行升級錢包人們將能夠通過移動簽名交易的見證人部分利用增加的塊大小。

## 邁嚮單一組合限製

目前，有兩個保守的強製限製：該單一組合不超過 1MB，並在整個交易執行不可超過 20,000 簽名檢查。

尋找最有利的交易集合給出單一的限製以包括是背包問題，可以用一個簡單的算法來容易地解決幾乎完美的實例。然而加入第二約束使得在某些情況下，
很難找到一個很好的解決方案，這一理論問題在實踐中已經迫使單一數據遠低於容量待開採。

如冇有 hardfork 或基本上減小單一數據是無法解決這個問題。由於 segwit 無法修複，在不使情況變得更糟情況下：與其引入一個獨立限製分離的見證數據，
它採用單一限製適用於UTXO數據的加權和證人數據，同時允許合並後的實體限製。

### 誰受惠?

未來hardfork如能改變單一容量限製，使用者就能受惠。 如：

    50* sigops + 4*basedata + 1*witnessdata < 10M

使用者能輕易準確地填寫單一數據同時最大限度並提高手續費收入，這讓他們更可靠地計算出需要進行開採其交易的費用讓用戶受益。

### 更多資料

 * [背包問題](https://en.wikipedia.org/wiki/Knapsack_problem)
 * [2015 年 8 月 Sigop 攻擊 bitcointalk 的討論](https://bitcointalk.org/index.php?topic=1166928.0;all)
 * [Gregory Maxwell 在bitcoin-dev 郵件列表上發表的witness limits](https://lists.linuxfoundation.org/pipermail/bitcoin-dev/2015-December/011870.html)
 * [“驗證成本度量” 成績](http://diyhpl.us/wiki/transcripts/scalingbitcoin/hong-kong/validation-cost-metric/)
