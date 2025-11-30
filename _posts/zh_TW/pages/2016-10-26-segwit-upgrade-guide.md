---
title: 隔離見證升級指南
name: segwit-upgrade-guide
id: zh_tw-segwit-upgrade-guide
type: posts
layout: post
lang: zh_TW
permalink: /zh_TW/2016/10/27/segwit-upgrade-guide/
version: 1
excerpt: 在 Bitcoin Core 0.13.1 中發布的 segwit 版本經歷了近兩年的迭代設計、開發和測試，過去一年的大部分精力都集中在讓現有的 Bitcoin 使用者、企業、開發者和礦工盡可能輕鬆地升級到 segwit。
---
{% include toc.html %}
{% include references.md %}

在 Bitcoin Core 0.13.1 中發布的 segwit 版本經歷了近兩年的迭代設計、開發和測試，過去一年的大部分精力都集中在讓現有的 Bitcoin 使用者、企業、開發者和礦工盡可能輕鬆地升級到 segwit。

初始 segwit 採用需要兩個群體的參與：

- **[礦工][miners guide]**代表 95% 或更多的 Bitcoin 網路總算力必須發出支援 segwit 的信號，以便鎖定 segwit 的啟動。

- **[全節點][node guide]**由合理數量的使用者和企業執行以驗證其接收的支付，需要升級到 Bitcoin Core 0.13.1 或更高版本，或另一個 segwit 相容實作，以便在 segwit 啟動後激勵礦工遵循 segwit 的規則。（這是 Bitcoin 的正常激勵機制，礦工只有在遵循所有共識規則時才能獲得區塊收入，一旦 segwit 啟動，這將包括新的 segwit 共識規則。）

segwit 軟分叉的設計允許這兩個群體的個人自願決定是否要採用 segwit，下面為想要採用 segwit 和不想採用的人提供了指南。

如果足夠多的礦工確實決定採用 segwit，它最終將啟動，錢包使用者將能夠開始建立帶有隔離見證的交易。segwit 軟分叉也被設計為與所有常用錢包向後和向前相容，因此錢包開發者和使用者也可以獨立決定是否要採用 segwit 或繼續進行不帶隔離見證的交易。下面為採用和不採用的[開發者][dev guide]和[使用者][user guide]提供了指南。

除了說明之外，以下每個指南章節的末尾還提供了一個簡短的推薦列表，列出了您可能遇到的任何 segwit 相關問題的詢問地點。

## 礦工

*本節是為獨立礦工和礦池營運商撰寫的。礦池礦工應聯繫其礦池營運商，以了解他們需要做什麼（如果有的話）才能升級或不升級到 segwit。*

BIP9 軟分叉部署機制被用於 segwit——與 2016 年 7 月 4 日啟動的 BIP 68/112/113 軟分叉相同的機制。無論您是否希望升級，您都應該了解升級過程的重要階段：

- **已開始：** Segwit 將在 2016 年 11 月 15 日或之後的第一個重定目標期開始時處於*已開始*階段，直到它啟動或被認為失敗（根據 BIP9 的政策），在一年內未達到鎖定後。在此期間，願意並能夠強制執行 segwit 新共識規則的礦工將透過在區塊標頭 versionbits 欄位中放置位元 1 來發出他們這樣做的意圖信號。

- **已鎖定：** 如果在 2,016 個區塊的重定目標期間，95% 的區塊發出支援 segwit 的信號，segwit 軟分叉將被*鎖定*，並計劃在 2,016 個區塊後（約兩週）*啟動*。

- **已啟動：** 在*已鎖定*期間結束後，發出準備強制執行 segwit 信號的礦工將開始產生包含帶有隔離見證的交易的 segwit 樣式區塊。

### 升級

segwit 軟分叉的 BIP9 參數允許礦工在 2016 年 11 月 15 日或之後的第一個重定目標期開始時開始發出支援信號。要發出支援信號，您需要執行以下操作：

- 將您用於交易選擇和區塊建構的全節點升級到 Bitcoin Core 0.13.1+ 或另一個 segwit 相容的全節點。

- 將您的挖礦軟體、礦池軟體或兩者都升級到 segwit 相容版本。

- 開始產生包含 segwit 的 BIP9 versionbit（即位元 1）的區塊。

當 segwit 啟動時，您將希望能夠挖掘和中繼 segwit 樣式區塊。以下挖礦軟體已升級以支援 segwit。

- 全節點：
  - [Bitcoin Core](https://bitcoin.org/en/download) 0.13.1 或更高版本
  - [Bitcoin Knots](http://bitcoinknots.org/) 0.13.1 或更高版本
  - [Btcd](https://github.com/btcsuite/btcd/pull/656)\*

- 挖礦軟體：
  - [BFGMiner](https://github.com/luke-jr/bfgminer)\*
  - [CGMiner](https://github.com/ckolivas/cgminer)
  - [libblkmaker](https://github.com/bitcoin/libblkmaker/pull/6)\*

- 礦池軟體：
  - [ckpool](https://bitbucket.org/ckolivas/ckpool)
  - [Eloipool](https://github.com/luke-jr/eloipool)
  - [Stratum-Mining](https://github.com/slush0/stratum-mining/pull/16)\*

- 中繼軟體：
   - [Bitcoin FIBRE](http://bitcoinfibre.org/)

請注意，支援 GetBlockTemplate (GBT) RPC 的軟體必須升級以支援對 GBT 的 BIP9 和 BIP145 變更。上面連結的所有支援 GBT 的程式都已升級。

Segwit 已經在測試網上啟動並強制執行，因此您可能會發現在測試網上使用少量算力進行挖礦來測試您的基礎設施升級很有用。或者，Bitcoin Core 0.13.1 的回歸測試模式（regtest）預設也支援 segwit。

**有疑問？** 歡迎獨立礦工和礦池營運商在 irc.freenode.net 的 #bitcoin-mining 中尋求幫助。礦池礦工應聯繫其礦池營運商以了解有關礦池關於 segwit 的政策的任何問題。

### 不升級

本節描述了如果您不想強制執行 segwit，作為礦工可以做什麼。

在*已開始*階段，如果您不想採用 segwit，您可以簡單地拒絕升級到 segwit 相容的全節點（例如 Bitcoin Core 0.13.1 或更高版本），以及避免使用任何假設您想要設定 segwit 的 versionbit 位元 1 的挖礦軟體。

如果 segwit 達到*已鎖定*，您仍然不需要升級，但強烈建議升級。segwit 軟分叉不要求您產生 segwit 樣式區塊，因此您可以無限期地繼續產生非 segwit 區塊。但是，一旦 segwit 啟動，其他礦工可能會產生您認為有效但每個強制執行 segwit 的節點都拒絕的區塊；如果您在這些無效區塊上建立任何區塊，您的區塊也將被視為無效。

因此，在 segwit 達到*已鎖定*後，建議您升級您的全節點到 Bitcoin Core 0.13.1 或更高版本（或相容的全節點），或者您遵循下面全節點章節中的「不升級」說明，使用 Bitcoin Core 0.13.1 或更高版本作為您 pre-segwit 軟體的過濾器。

## 全節點使用者

*本節是為任何操作全節點的人撰寫的，包括企業和個人。*

全節點防止其使用者接受違反任何 Bitcoin 共識規則的任何區塊。在像 segwit 這樣的軟分叉升級中，會新增新規則，任何不升級的節點都不會知道這些新規則。這不是問題：segwit 軟分叉的設計允許未升級的使用者像軟分叉之前一樣繼續使用 Bitcoin。

但是，任何想要使用 segwit 軟分叉啟用的功能的人都會希望知道，足夠數量的全節點使用者已經升級了他們的節點以拒絕違反 segwit 規則的區塊和交易，從而為礦工遵循 segwit 的更新共識規則提供強烈的激勵。

這個系統在過去運作良好，在之前幾個軟分叉啟動之前，至少有 25% 的可達節點（通常為 50% 或更多）已升級（不計算 BIP50 緊急和臨時軟分叉）。沒有理由期望 segwit 軟分叉會有任何不同，升級是支援 segwit 的人幫助鼓勵其採用的簡單方法。當然，對 segwit 不感興趣的人可以簡單地不升級。以下描述了兩種情況的詳細資訊。

### 升級

要升級到 segwit 相容版本，請下載您的全節點軟體的 segwit 相容版本（例如 [Bitcoin Core 0.13.1 版本](https://bitcoin.org/en/download)），確保您下載的檔案是合法的（使用 PGP 或其他方法），停止舊版本的節點軟體，然後啟動新版本的軟體。請注意，如果您在 segwit 啟動後升級，您的節點將需要從啟動點向前下載和重新同步區塊，因為舊版本沒有完全下載它們。

您可以使用 Bitcoin Core RPC `getblockchaininfo` 來追蹤 segwit 軟分叉的狀態（在 BIP9 樣式軟分叉清單中標記為 `segwit`）。此資訊包括有多少近期區塊是由發出打算強制執行 segwit 新共識規則的礦工產生的。`getblockchaininfo` RPC 的結果還將讓您確定 segwit 的軟分叉何時被鎖定（意味著它將在接下來的 2,016 個區塊內啟動）和啟動（意味著它現在由礦工強制執行）。

Bitcoin Core 0.13.1 提供的錢包預設將繼續僅產生非 segwit P2PKH 位址以接收支付。預計後續版本將允許使用者選擇接收支付到 segwit 位址。

如果您是想要產生位址進行測試的開發者或專家使用者，請參閱 [segwit dev guide][]。

**有疑問？** 如果您使用 Bitcoin Core 作為您的全節點，請參閱 Bitcoin.org 上的 [取得幫助](https://bitcoin.org/en/bitcoin-core/help) 頁面以獲取各種支援選項。如果您使用另一個全節點，最好的詢問地點是您的全節點軟體使用者尋求支援的地方。您的軟體維護者至少會熟悉 segwit 背後的想法，他們將能夠告訴您何時會實作它以及它將如何影響您。

### 不升級

如果您不想升級到 segwit 並且您不是礦工，您可以簡單地繼續使用您當前的全節點軟體。Segwit 實作為軟分叉，因此您不需要升級。您也不需要升級連接到您全節點的任何錢包；它們將像以前一樣繼續工作（有關詳細資訊，請參閱下面的[錢包章節][user guide]）。

但是，如果您接受確認區塊較少的交易（例如一個或兩個區塊），請注意，在軟分叉啟動後，未升級的全節點暫時接受無效區塊的風險會略微增加。隨著升級的礦工繼續強制執行新的 segwit 共識規則，情況將在幾個區塊內自行解決，但不能保證在無效區塊中顯示為已確認的交易將繼續在有效區塊中得到確認。

防止此問題的最簡單方法是升級到 Bitcoin Core 0.13.1 或更高版本，或另一個與 segwit 軟分叉相容的全節點版本。如果您仍然不想升級，可以使用較新的 Bitcoin Core 版本作為較舊 Bitcoin Core 版本的過濾器。

![Filtering by an upgraded node](/assets/images/filtering-by-upgraded-node.svg)

在此配置中，您將當前的 Bitcoin Core 節點（我們稱之為「較舊節點」）設定為僅連接到執行 Bitcoin Core 0.13.1 或更高版本的節點（我們稱之為「較新節點」）。較新節點像往常一樣連接到 Bitcoin P2P 網路。因為較新節點知道 segwit 對共識規則的變更，它不會將無效區塊或交易中繼到較舊節點——但它會中繼其他所有內容。

使用此配置時，請注意，如果較舊節點使用 Bitcoin Core 預設值，它將不會看到使用 segwit 功能的交易，直到這些交易被包含在區塊中。

配置：

對於較新節點，正常啟動它並讓它同步區塊鏈。目前，您不能將修剪節點用於此目的，因為修剪節點不會作為中繼節點。您可以選擇使用以下一個或兩個命令列參數啟動較新節點，以便它將較舊節點視為特殊（這些選項也可以放在 Bitcoin Core 的配置檔案中）：

~~~
  -whitebind=<addr>
       綁定到給定位址並將連接到它的對等點列入白名單。使用
       [host]:port 表示法表示 IPv6

  -whitelist=<netmask>
       將從給定網路遮罩或 IP 位址連接的對等點列入白名單。可以
       多次指定。列入白名單的對等點不能被 DoS 封禁
       並且它們的交易總是被中繼，即使它們已經
       在記憶池中，例如對閘道很有用
~~~

對於較舊節點，首先等待較新節點完成同步區塊鏈，然後使用以下命令列參數重新啟動較舊節點（這也可以放在 Bitcoin Core 配置檔案中）：

    -connect=<較新節點的 IP 位址或 DNS 名稱>

例如，

    -connect=192.168.8.8

這將導致較舊節點僅連接到較新節點，以便所有區塊和交易都由較新節點過濾。

## 錢包使用者

*本節是為使用輕量級錢包、網頁錢包、連接到個人全節點的錢包或任何其他錢包的任何人撰寫的。*

### 升級

如果您確實想升級到 segwit，您首先需要等待礦工啟動 segwit，然後您需要一個支援接收和花費 segwit 樣式支付的錢包。這適用於 Bitcoin Core 的錢包、輕量級錢包，以及第三方代表您發送和接收比特幣的錢包（有時稱為網頁錢包）。Bitcoin Core 或其他全節點的使用者也應該閱讀上面有關全節點的章節。

在您的錢包升級以支援 segwit 後，它將產生以「3」開頭的接收位址（P2SH 位址）。有些錢包多年來一直在產生 P2SH 位址，因此這對您來說可能不是變化。

所有常用錢包都能夠支付 P2SH 位址，因此您將能夠從任何常見錢包接收支付，無論它們是否已升級到 segwit。升級到 segwit 後花費您的比特幣時，您仍然能夠支付以「1」開頭的原始類型 Bitcoin 位址（P2PKH 位址），以及能夠支付其他 P2SH 位址的使用者。

使用 segwit 錢包花費您的比特幣時，您會注意到以下情況：

- 當僅花費您在升級前收到的比特幣時，您應該注意到與升級前建立的交易沒有差異。

- 當將升級到 segwit 後收到的比特幣花費給尚未升級到 segwit 的人時，他們可能要等到交易被包含在區塊中後才能看到您的交易。這是一個安全功能，可避免向他們顯示他們的錢包不完全理解的交易，直到這些交易得到確認。交易確認後，他們將能夠像正常一樣看到和花費您發送給他們的比特幣。

- 當花費您升級後收到的新 P2SH 位址的比特幣時，您可能會注意到您支付的交易手續費略低於花費您之前收到的非 segwit 支付時。這是因為包含您簽名的交易部分（「見證」）不需要被 Bitcoin 全節點快速存取，所以 segwit 允許礦工在區塊中儲存最多 4 倍的見證位元組，而不是非見證位元組。這更好地將建立區塊的成本（因此其交易手續費）與操作全節點的實際成本相一致。

**有疑問？** 如果您有任何問題，最好的詢問地點是您的錢包使用者尋求支援的地方。您的錢包維護者將熟悉 segwit 背後的想法，他們將能夠告訴您是否會為您的錢包實作 segwit、何時可能發生以及它將如何影響您對錢包的使用。

### 不升級

如果您不想升級到 segwit，您可以簡單地繼續使用任何尚未新增 segwit 支援的錢包。即使您沒有升級，您也將能夠與已升級到 segwit 的使用者以及像您一樣尚未升級到 segwit 的使用者進行交易。

如果您不升級，您可能會遇到一個差異：如果已升級到 segwit 的人向您付款，您的錢包可能要等到支付被包含在區塊中後才會向您顯示。這是一個安全功能，可防止您的錢包看到它不完全理解的交易，直到它們被礦工確認。

## Bitcoin 軟體開發者

*本節是為處理交易或區塊的任何 Bitcoin 軟體開發者撰寫的。*

所有 Bitcoin 軟體（包括挖礦軟體，前提是它僅選擇遵循預設中繼政策的交易）都應該像以前一樣繼續工作，並且只有在您想要利用 segwit 的新功能時才需要升級您的軟體。

Segwit 在以下文件中為開發者進行了描述：

- **[Segwit 錢包開發者指南][segwit dev guide]：** 您需要了解的所有內容的摘要描述，以升級您的錢包以支援 segwit。

- **[BIP141][] 隔離見證（共識層）：** 尋求實作 segwit 任何方面的開發者應閱讀本文件的規範章節。挖礦和全節點軟體的開發者將在部署章節中找到 segwit 的 BIP9 參數。

- **[BIP143][] 版本 0 見證程式的交易簽名驗證：** 希望建立或驗證 segwit 簽名的任何軟體開發者應閱讀本文件的規範章節，並建議使用範例章節進行測試。

- **[BIP144][] 隔離見證（對等服務）：** 尋求透過 Bitcoin P2P 網路發送或接收隔離見證的開發者應閱讀本文件的規範章節。

- **[BIP145][] 隔離見證的 getblocktemplate 更新：** 挖礦和其他產生或使用 GetBlockTemplate RPC 的軟體開發者應閱讀 BIP145 和 [BIP9][] 中的相關 GBT 變更。

- **[BIP147][] 處理虛擬堆疊元素可塑性：** 錢包開發者，尤其是新交易腳本開發者，應該了解這個新的共識規則，它反映了一個長期存在的預設網路中繼政策，禁止將除 `OP_0`「null」操作碼以外的任何內容作為「虛擬」參數傳遞給 checkmultisig 樣式操作碼。在 segwit 啟動後，這個新的共識規則將適用於使用 segwit 和不使用 segwit 的交易。

- **[VersionBits FAQ][]：** 礦工和挖礦軟體開發者應閱讀此常見問題以獲取有關設定其 versionbits 以發出支援軟分叉的信號的資訊。Segwit 使用位元 1 作為 versionbits。

請注意，[BIP142][]（隔離見證的位址格式）處於*延期*狀態（如 BIP1 所定義），並未作為標準提出。相反，邀請錢包開發者在 [bitcoin-dev 郵件列表][bitcoin-dev]上討論建立一個新的 Bitcoin 位址格式，該格式將比當前的 base58check 編碼位址更易於使用。

BIP 141、143、144 和 145 的大多數實作詳細資訊可在 [Bitcoin Core PR#8149](https://github.com/bitcoin/bitcoin/pull/8149) 中找到。BIP147 的實作可在 [Bitcoin Core PR#8636](https://github.com/bitcoin/bitcoin/pull/8636) 中找到。

為了在啟用 segwit 的網路上測試變更，測試網（testnet3）已經支援 segwit 數月，並包括大量 segwit 區塊，包括幾乎達到 segwit 允許的最大區塊大小的區塊。Bitcoin Core 的回歸測試（regtest）模式在 Bitcoin Core 0.13.0 和 0.13.1 中預設也支援 segwit。

除了 Bitcoin Core 之外，許多自由和開源軟體 Bitcoin 錢包和套件也已經新增了 segwit 相容性或準備好部署 segwit 相容程式碼，因此，如果它們的著作權授權與您的程式碼相容，您可以使用它們的程式碼變更作為更新您軟體的範例。

**有疑問？** Bitcoin 開發問題可以在 irc.freenode.net 的 #bitcoin-dev IRC 聊天室中詢問。問題也可以在 Bitcoin.StackExchange.com 和 BitcoinTalk.org [技術討論板](https://bitcointalk.org/index.php?board=6.0)上詢問。

[miners guide]: #miners
[node guide]: #full-node-users
[dev guide]: #bitcoin-software-developers
[user guide]: #wallet-users
[segwit dev guide]: /zh_TW/segwit_wallet_dev/
[versionbits faq]: /en/2016/06/08/version-bits-miners-faq/
