---
version: 4
title: 隔離見證錢包開發指南
name: segwit-wallet-dev
id: zh_tw-segwit-wallet-dev
type: pages
layout: page
lang: zh_TW
permalink: /zh_TW/segwit_wallet_dev/
---
{% include toc.html %}
{% include references.md %}

本文件的大部分內容可在與隔離見證相關的 BIP 中找到，包括 [BIP141][]、[BIP143][]、[BIP144][] 和 [BIP145][]。請將此視為其他相關文件的首要參考點，以及應該做和不應該做的檢查清單。

### 基本隔離見證支援

為了被視為基本層級的 segwit 相容，錢包必須實作本節中的所有功能：

#### 發送至 P2SH

* segwit 相容錢包必須支援 pay-to-script-hash（[BIP16][]）及其位址格式（[BIP13][]）。
* 對於進行支付，錢包必須能夠正確地將給定的 P2SH 位址轉換為 <code>scriptPubKey</code>，並建立交易。
* 對於接收支付，錢包必須能夠基於 P2WPKH 腳本（如下定義）建立 P2SH 位址，並能夠識別支付至此類位址的款項。
* 這是強制性要求，即使錢包僅接受單一簽名支付。

#### 建立 P2SH-P2WPKH 位址

* P2SH-P2WPKH 位址可與 Bitcoin 原始的單一簽名 P2PKH 位址（前綴為 1 的位址）相比較。
* 與任何其他 P2SH 位址一樣，P2SH-P2WPKH 位址的前綴為 3。
* 在 P2SH-P2WPKH UTXO 被花費且 <code>redeemScript</code> 被揭露之前，P2SH-P2WPKH 位址與非 segwit P2SH 位址（例如非 segwit 多重簽名位址）無法區分。
* 當僅使用 1 個公鑰接收支付時（類似 P2PKH），應使用 P2SH-P2WPKH 位址。
* P2SH-P2WPKH 使用與 P2PKH 相同的公鑰格式，但有一個非常重要的例外：P2SH-P2WPKH 中使用的公鑰必須是壓縮的，即大小為 33 位元組，並以 <code>0x02</code> 或 <code>0x03</code> 開頭。使用任何其他格式（如未壓縮的公鑰）可能導致資金永久損失。
* 建立 P2SH-P2WPKH 位址：
    1. 計算公鑰的 SHA256 的 RIPEMD160（<code>keyhash</code>）。儘管 <code>keyhash</code> 公式與 P2PKH 相同，但為了更好的隱私和防止意外使用未壓縮金鑰，應避免重複使用 <code>keyhash</code>。
    2. P2SH <code>redeemScript</code> 總是 22 位元組。它以 <code>OP_0</code> 開頭，後接 <code>keyhash</code> 的標準推送（即 <code>0x0014{20-byte keyhash}</code>）。
    3. 與任何其他 P2SH 相同，<code>scriptPubKey</code> 是 <code>OP_HASH160 hash160(redeemScript) OP_EQUAL</code>，位址是對應的前綴為 3 的 P2SH 位址。

#### 交易序列化

* segwit 相容錢包必須支援原始交易格式，即 <code>nVersion|txins|txouts|nLockTime</code>。
* segwit 相容錢包也必須支援新的序列化格式，即 <code>nVersion|marker|flag|txins|txouts|witness|nLockTime</code>。
    * <code>nVersion</code>、<code>txins</code>、<code>txouts</code> 和 <code>nLockTime</code> 的格式與原始格式相同。
    * <code>marker</code> 必須為 <code>0x00</code>。
    * <code>flag</code> 必須為 <code>0x01</code>。
    * <code>witness</code> 是交易所有見證資料的序列化。
        * 每個 txin 都關聯一個見證欄位。因此，沒有見證欄位數量的指示，因為它由 <code>txins</code> 的數量暗示。
        * 每個見證欄位以一個 <code>compactSize</code> [整數](https://bitcoin.org/en/developer-reference#compactsize-unsigned-integers)開頭，以指示對應 <code>txin</code> 的堆疊項目數量。然後是對應 <code>txin</code> 的見證堆疊項目（如果有）。
        * 每個見證堆疊項目以一個 <code>compactSize</code> 整數開頭，以指示項目的位元組數。
        * 如果 <code>txin</code> 未與任何見證資料關聯，則其對應的見證欄位是確切的 <code>0x00</code>，表示見證堆疊項目數量為零。
* 如果交易中的所有 <code>txins</code> 都未與任何見證資料關聯，則交易必須以原始交易格式序列化，不包含 <code>marker</code>、<code>flag</code> 和 <code>witness</code>。例如，如果沒有任何 <code>txins</code> 來自 segwit UTXO，則必須以原始交易格式序列化。（例外：coinbase 交易）
* 交易序列化的範例可在 BIP143 的範例章節中找到。錢包開發者可使用這些範例來測試其實作是否正確解析新的序列化格式。

#### 交易 ID

* 在 segwit 下，每個交易將有 2 個 ID。
* <code>txid</code> 的定義保持不變：原始序列化格式的雙 SHA256。
* 定義了一個新的 <code>wtxid</code>，它是包含見證資料的新序列化格式的雙 SHA256。
* 如果交易沒有任何見證資料，則其 <code>wtxid</code> 與 <code>txid</code> 相同。
* <code>txid</code> 仍然是交易的主要識別碼：
    * 在引用先前輸出時，它必須在 <code>txin</code> 中使用。
    * 如果錢包或服務目前使用 <code>txid</code> 識別交易，預期在 segwit 中繼續使用。

#### P2SH-P2WPKH 的簽名產生和驗證

* 對於花費非 segwit UTXO，簽名產生演算法保持不變。
* 對於花費 P2SH-P2WPKH：
    * <code>scriptSig</code> 必須僅包含 <code>redeemScript</code> 的推送。
    * 對應的見證欄位必須恰好包含 2 個項目：簽名後接公鑰。
    * [BIP143][] 中描述了 segwit 腳本的新簽名產生演算法。開發者應仔細遵循說明，並利用 [BIP143](https://github.com/bitcoin/bips/blob/master/bip-0143.mediawiki#P2SHP2WPKH) 中的 P2SH-P2WPKH 範例來確保能夠重現 <code>sighash</code>。
    * [BIP143][] 簽名產生演算法涵蓋了被花費輸入的值，這簡化了氣隙輕量級錢包和硬體錢包的設計。
    * 請注意，對於 P2SH-P2WPKH，<code>scriptCode</code> 總是 26 位元組（包括前導大小位元組），即 <code>0x1976a914{20-byte keyhash}88ac</code>，而不是 <code>redeemScript</code> 或 <code>scriptPubKey</code>。
    * [範例](https://blockchainprogramming.azurewebsites.net/checktx?txid=8139979112e894a14f8370438a471d23984061ff83a9eba0bc7a34433327ec21)

#### 網路服務（可選）

* 如果錢包將透過點對點網路發送和接收交易，則需要網路服務。
* 具有 segwit 能力的節點將透過服務位元廣告它們可以提供見證：<code>NODE_WITNESS = (1 << 3)</code>。
* 沒有任何見證資料（因此以原始格式序列化）的交易可以發送到具有或不具有 <code>NODE_WITNESS</code> 支援的節點。
* 花費 segwit UTXO 的交易（因此以新格式序列化）只能發送到具有 <code>NODE_WITNESS</code> 支援的節點。
* 花費 segwit UTXO 但見證資料被剝離（因此以原始格式序列化）的交易可以發送到不具有 <code>NODE_WITNESS</code> 支援的節點。但是，此類交易在 segwit 啟動後無效，不會被區塊接受。
* 網路服務的詳細資訊可在 [BIP144][] 中找到。

#### 使用者隱私

* 在 segwit 交易變得普遍之前，此交易類型可能使 Bitcoin 追蹤更容易。
* 使用 P2SH-P2WPKH 作為預設找零輸出也可能對隱私產生影響。

#### 交易手續費估算

* 不使用交易大小，而是定義了一個新的指標，稱為「虛擬大小」（<code>vsize</code>）。
* 交易的 <code>vsize</code> 等於原始序列化大小的 3 倍，加上新序列化的大小，將結果除以 4 並向上取整到下一個整數。例如，如果交易使用新序列化為 200 位元組，並在移除 <code>marker</code>、<code>flag</code> 和 <code>witness</code> 後變為 99 位元組，則 <code>vsize</code> 為 (99 * 3 + 200) / 4 = 125（向上取整）。
* 非 segwit 交易的 <code>vsize</code> 就是其大小。
* 交易手續費應透過將 <code>vsize</code> 與其他交易比較來估算，而不是大小。
* 開發者應小心避免在手續費估算中出現 4 倍的錯誤。

#### 啟動 {#upgrade-safety}

* 從區塊高度 481824 開始，所有 SegWit 就緒節點開始強制執行新的 SegWit 共識規則。

#### 向後相容性

* 應繼續支援發送和接收舊版 P2PKH 支付（前綴為 1 的位址）。

### 複雜腳本支援

如果錢包支援單一簽名以外的腳本類型（例如多重簽名），則必須滿足基本要求以外的更多要求：

#### 建立 P2SH-P2WSH 位址

* P2SH-P2WSH 位址可與 Bitcoin 原始的 P2SH 位址相比較，它允許用固定大小的位址表示任意複雜的腳本。
* 與任何其他 P2SH 和 P2SH-P2WPKH 位址一樣，P2SH-P2WSH 位址的前綴為 3。在 UTXO 被花費之前，它們無法區分。
* 建立 P2SH-P2WSH 位址：
    1. 定義一個腳本，稱為（<code>witnessScript</code>）。
    2. 計算 <code>witnessScript</code> 的 SHA256（<code>scripthash</code>）。請注意使用單一 SHA256，而不是雙 SHA256 或 RIPEMD160(SHA256)。
    3. P2SH <code>redeemScript</code> 總是 34 位元組。它以 <code>OP_0</code> 開頭，後接 <code>scripthash</code> 的標準推送（即 <code>0x0020{32-byte scripthash}</code>）。
    4. 與任何其他 P2SH 相同，<code>scriptPubKey</code> 是 <code>OP_HASH160 hash160(redeemScript) OP_EQUAL</code>，位址是對應的前綴為 3 的 P2SH 位址。
* 腳本限制
    * 腳本評估不得失敗，並且必須在評估後僅留下一個且僅一個 TRUE 堆疊項目。否則，評估失敗。
    * P2SH-P2WSH 腳本內的任何公鑰必須是壓縮金鑰，否則資金可能永久損失。
    * 如果使用 OP_IF 或 OP_NOTIF，其參數必須是空向量（表示 false）或 <code>0x01</code>（表示 true）。使用其他值可能導致資金永久損失。（[BIP 草案](https://lists.linuxfoundation.org/pipermail/bitcoin-dev/2016-August/013014.html)）
    * 如果 OP_CHECKSIG 或 OP_CHECKMULTISIG 返回失敗，則所有簽名必須是空向量。否則，資金可能永久損失。（[BIP146][]）
    * <code>witnessScript</code> 有 3600 位元組的預設政策限制。除了 <code>witnessScript</code>，最多可以有 100 個見證堆疊項目，每個最多 80 位元組。超過這些限制的交易可能不會被中繼或包含在區塊中。
    * 許多原始腳本共識限制，例如 10000 位元組腳本大小、201 <code>nOpCount</code>，仍然適用於 P2SH-P2WSH。
    * P2SH 的 520 位元組腳本大小限制不適用於 P2SH-P2WSH。它被 3600 位元組政策限制和 10000 位元組共識限制取代。

#### P2SH-P2WSH 的簽名產生和驗證

* 對於花費 P2SH-P2WSH：
    * <code>scriptSig</code> 必須僅包含 <code>redeemScript</code> 的推送。
    * 對應見證欄位的最後一個見證項目必須是 <code>witnessScript</code>。
    * 應用新的 [BIP143][] 簽名產生演算法：
        * 在不使用任何 OP_CODESEPARATOR 的情況下，<code>scriptCode</code> 是前面帶有 <code>witnessScript</code> 大小的 <code>compactSize</code> 整數的 <code>witnessScript</code>。例如，如果腳本是 OP_1（<code>0x51</code>），則序列化的 <code>scriptCode</code> 是（<code>0x0151</code>）。
        * 對於任何包含 OP_CODESEPARATOR 的不尋常腳本，請參閱 [BIP143][] 以了解確切的語義。
    * <code>witnessScript</code> 之前的任何見證堆疊項目都用作腳本評估的輸入堆疊。輸入堆疊不會被解釋為腳本。例如，不需要使用 <code>0x4c</code>（OP_PUSHDATA1）來「推送」大項目。
    * 要驗證簽名產生和堆疊序列化的正確性，請始終針對 [BIP143][] 中的範例進行測試。
    * [範例](https://blockchainprogramming.azurewebsites.net/checktx?txid=954f43dbb30ad8024981c07d1f5eb6c9fd461e2cf1760dd1283f052af746fc88)

### Segwit 原生位址（可選）{#advanced-designs}

以下功能對於初始 segwit 支援不是必需的。

#### 原生 Pay-to-Witness-Public-Key-Hash (P2WPKH)

* 原生 P2WPKH 是 22 位元組的 <code>scriptPubKey</code>。它以 <code>OP_0</code> 開頭，後接 <code>keyhash</code> 的標準推送（即 <code>0x0014{20-byte keyhash}</code>）。
* 與 P2SH-P2WPKH 相同，<code>keyhash</code> 是壓縮公鑰的 RIPEMD160(SHA256)。
* 花費原生 P2WPKH 時，<code>scriptSig</code> 必須為空，見證堆疊格式和簽名產生規則與 P2SH-P2WPKH 相同（包括使用壓縮公鑰的要求）。
* [範例](https://blockchainprogramming.azurewebsites.net/checktx?txid=d869f854e1f8788bcff294cc83b280942a8c728de71eb709a2c29d10bfe21b7c)

#### 原生 Pay-to-Witness-Script-Hash (P2WSH)

* 原生 P2WSH 是 34 位元組的 <code>scriptPubKey</code>。它以 <code>OP_0</code> 開頭，後接 <code>scripthash</code> 的標準推送（即 <code>0x0020{32-byte scripthash}</code>）。
* 與 P2SH-P2WSH 相同，<code>scripthash</code> 是 <code>witnessScript</code> 的 SHA256。
* 花費原生 P2WSH 時，<code>scriptSig</code> 必須為空，見證堆疊格式和簽名產生規則與 P2SH-P2WSH 相同（包括使用壓縮公鑰的要求）。
* [範例](https://blockchainprogramming.azurewebsites.net/checktx?txid=78457666f82c28aa37b74b506745a7c7684dc7842a52a457b09f09446721e11c)

#### 為什麼以及如何使用原生（Bech32）P2WPKH 和 P2WSH？

* [BIP173][] 為原生 P2WPKH 和 P2WSH 位址提出了一個帶校驗和的 base32 格式（Bech32）。
* Bitcoin Core v0.16.0 中包含了對 Bech32 位址的支援。
* 與 P2SH 版本相比，原生版本的交易 <code>vsize</code> 在大多數情況下較小，因此可能需要較少的手續費。
* 原生 P2WPKH 和 P2WSH 可與原始 <code>scriptPubKey</code> 協議一起使用，例如支付協議（BIP70）。但是，這可能影響付款人和收款人的隱私（見下文）。
* 原生 P2WPKH 和 P2WSH 可用作預設找零位址，但這可能讓其他人輕易識別找零（見下文）。
* 在原生 P2WPKH 和 P2WSH 廣泛使用之前，這些位址類型可能在使用者中引起隱私問題。

### 腳本和交易範例

* [不同見證交易類型的範例和交易有效性檢查工具](https://blockchainprogramming.azurewebsites.net/checktx)
* [BIP141][]
* [BIP143][]
* [BIP173][]
* [腳本測試](https://github.com/bitcoin/bitcoin/blob/master/src/test/data/script_tests.json)
* [有效交易測試](https://github.com/bitcoin/bitcoin/blob/master/src/test/data/tx_valid.json)
* [無效交易測試](https://github.com/bitcoin/bitcoin/blob/master/src/test/data/tx_invalid.json)
