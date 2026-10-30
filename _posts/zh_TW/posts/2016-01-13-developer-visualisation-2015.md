---
type: posts
layout: post
lang: zh_TW
id: zh_TW-development-visualisation-2015
name: development-visualisation-2015
title: 2015 年核心開發視覺化
permalink: /zh_TW/2016/01/13/development-visualisation-2015/
version: 1
excerpt: 以下影片顯示了 2015 年期間 Bitcoin Core 儲存庫中的提交活動。
---
以下影片顯示了 2015 年期間 [Bitcoin Core 儲存庫][repository]中的提交活動。此期間程式碼貢獻者的完整列表可在[此處][activity]找到。

https://www.youtube.com/embed/FIt7GLxxIpY

2015 年，Bitcoin Core 專案發布了 2 個主要版本的軟體以及 5 個進一步的維護版本。此外，部署並成功啟動了兩個軟分叉升級。第一個 [BIP66] 修復了 openssl 引入的潛在嚴重安全漏洞；第二個 [BIP65] 向 Bitcoin 指令碼語言新增了一個新的操作碼 CHECKLOCKTIMEVERIFY。

該專案還完成了下一個主要版本 [0.12] 的大部分工作，計劃於 2 月發布。[0.12] 將包含 [libsecp256k1]，該程式庫已開發了[兩年半][secp_contributors]，並將簽章驗證速度提高了 7 倍，這對於未來提高可擴展性至關重要。

請注意，提交活動僅代表整體開發者活動的一部分，並未記錄同行審查者、程式碼審查者、整合測試人員和翻譯人員的活動。它也沒有準確反映在被接受到程式庫之前進行研究、討論和開發所花費的時間。

我們也想藉此機會感謝迄今為止參與貢獻 Bitcoin Core 並幫助讓 Bitcoin 為每個人變得更好的每一個人。

[repository]: https://github.com/bitcoin/bitcoin
[activity]: https://github.com/bitcoin/bitcoin/graphs/contributors?from=2015-01-01&to=2016-01-01&type=c
[BIP65]: https://github.com/bitcoin/bips/blob/master/bip-0065.mediawiki
[BIP66]: https://github.com/bitcoin/bips/blob/master/bip-0066.mediawiki
[0.12]: https://github.com/bitcoin/bitcoin/blob/0.12/doc/release-notes.md
[libsecp256k1]: https://github.com/bitcoin/secp256k1
[secp_contributors]: https://github.com/bitcoin/secp256k1/graphs/contributors?from=2013-03-04&to=2015-12-01&type=c
