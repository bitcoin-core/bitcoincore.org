---
type: posts
layout: post
lang: zh_TW
id: zh_TW-segwit-testnet
name: segwit-testnet
title: 隔離見證測試網啟動
permalink: /zh_TW/2016/01/21/launch_segwit_testnet/
version: 1
excerpt: 我們非常高興和興奮地宣布隔離見證測試網的發布
---

我們非常高興和興奮地宣布[隔離見證測試網](https://github.com/sipa/bitcoin/commits/segwit)的發布！我們鼓勵所有開發者立即開始協助測試和整合工作。

這代表了迄今為止最受期待和令人興奮的發展之一，是許多未來改進和創新的重要基礎。隔離見證透過將交易簽名資料安全地移至交易區塊外的特別指定「隔離見證」資料結構中，釋放了 Bitcoin 區塊鏈上的空間。

這項重大創新使 Bitcoin 區塊鏈的使用效率大幅提升，同時為更廣泛的 Bitcoin 生態系統開啟了令人興奮的新可能性和機會，特別是在智慧合約應用和更快速的交易方面，同時也為未來更進階的功能和可能性奠定了基礎。

這項開發始於 Bitcoin Core 開發者 Dr. Pieter Wuille 的研究工作，最初專注於解決交易可塑性問題，這是一個長期存在且廣為人知的關注點和優先事項。然而，在研究過程中以及在縮小解決方案範圍時，發現了該解決方案的其他特性，這些特性不僅允許增加區塊大小，同時還開啟了一些令人難以置信的次要好處。

這項工作由 Dr. Pieter Wuille 發起，但也包含了許多其他人的貢獻，特別感謝 Eric Lombrozo、Johnson Lau、Alex Morcos、Nicolas Dorier、Bryan Bishop、Gregory Maxwell、Peter Todd、Cory Fields、Suhas Daftuar 和 Luke-Jr。

## Bitcoin Core 生態系統

提供者和其他交易所營運者將使用此版本中包含的基礎開發和創新來創造什麼，存在著廣泛的興奮和期待。到目前為止，最受歡迎的錢包和支援程式庫已表示將支援 segwit，包括 Ledger、Trezor、Electrum 和 Bitgo。此外，許多其他程式庫如 bitcoinj、bitcoinjs、pycoin 和 bitcore 的工作已經開始。

「segnet」代幣的水龍頭可在[這裡](https://segwit.greenaddress.it/faucet/)取得。

第三方錢包支援的早期預覽可在以下位置取得：

- [mSIGNA](https://github.com/ciphrex/mSIGNA/tree/segwit)（錢包原始碼）
- [Green Address](https://segwit.greenaddress.it/)（網頁錢包）

## 如何參與

請加入 irc.freenode.net 上的 `segwit-dev` IRC 頻道。

錢包提供者應閱讀[遷移指南](/zh_TW/segwit_wallet_dev)。

## 測試

最後也是最重要的，請協助測試隔離見證測試網！

原始碼可以在[這裡](https://github.com/sipa/bitcoin/tree/segwit)找到，請 checkout `segwit` 分支。

編譯完成後，在標準的 `bitcoind` 和 `bitcoin-cli` 命令列中加入 `-segnet`。

## 其他背景和歷史

- [Scaling Bitcoins 香港演講](https://prezi.com/lyghixkrguao/segregated-witness-and-deploying-it-for-bitcoin/)
- [擴展影片](https://bitcoincore.org/zh_TW/2015/12/14/segregated-witness)
- [逐字稿](http://diyhpl.us/wiki/transcripts/scalingbitcoin/hong-kong/segregated-witness-and-its-impact-on-scalability/)

## 技術參考

### SegWit BIP

- [BIP141](https://github.com/bitcoin/bips/blob/master/bip-0141.mediawiki)：隔離見證軟分叉的概述，以及對其好處、向後相容性、共識限制和部署的討論。
- [BIP143](https://github.com/bitcoin/bips/blob/master/bip-0143.mediawiki)：隔離見證交易中用於簽名驗證的新摘要演算法定義。
- [BIP144](https://github.com/bitcoin/bips/blob/master/bip-0144.mediawiki)：隔離見證交易的新訊息類型和序列化格式。

### 參考資料

- [隔離見證好處分析](http://lists.linuxfoundation.org/pipermail/bitcoin-dev/2016-January/012248.html)
- [硬分叉與軟分叉](https://petertodd.org/2016/soft-forks-are-safer-than-hard-forks)
- [「非爭議性」硬分叉研究的早期探索](https://scalingbitcoin.org/hongkong2015/presentations/DAY1/1_overview_1_timon.pdf)

[FAQ]: https://bitcoincore.org/zh_TW/2015/12/23/capacity-increases-faq
[roadmap]: http://lists.linuxfoundation.org/pipermail/bitcoin-dev/2015-December/011865.html
