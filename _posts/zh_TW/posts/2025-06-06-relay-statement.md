---
title: Bitcoin Core 開發和交易中繼策略
permalink: /zh_TW/2025/06/06/relay-statement/
lang: zh_TW
id: zh_TW-relay-statement
name: relay-statement
type: posts
layout: post
version: 1
---

我們想分享我們對 Bitcoin Core 開發與網路上交易中繼策略之間關係的看法。

Bitcoin 是一個由其使用者定義的網路，他們在選擇使用什麼軟體（完全驗證或非完全驗證）和實作他們想要的任何策略方面擁有最終自由。Bitcoin Core 貢獻者無權規定這些策略是什麼。這反映在我們長期以來避免在軟體中自動更新的做法中。這意味著沒有任何實體可以單方面向 Bitcoin Core 使用者推送變更：變更必須由使用者自己選擇採用新的軟體版本，或者如果他們願意，採用不同的軟體。自由執行任何軟體是網路對抗強制的主要保障。

作為 Bitcoin Core 開發者，我們也認為我們有責任使我們的軟體盡可能有效和可靠地為其目的服務，即在 Bitcoin 點對點網路中驗證和中繼區塊和交易，以便 Bitcoin 作為去中心化數位貨幣取得成功。關於交易中繼，這可能包括新增阻斷服務（DoS）保護和手續費評估的策略，但不會阻止中繼有持續經濟需求且可靠地進入區塊的交易。交易中繼的目標包括：

* 預測將被挖掘的交易（例如用於手續費估算或手續費提升，但它也是節點軟體內許多 DoS 保護策略的基礎）；
* 加快我們預期被挖掘的交易的區塊傳播。減少延遲有助於防止大型礦工獲得不公平的優勢；
* 幫助礦工了解付費交易（這樣他們就不需要依賴破壞挖礦去中心化的頻外交易提交方案）。

**明知拒絕中繼礦工無論如何都會包含在區塊中的交易會迫使使用者進入替代通訊管道，破壞上述目標。**

交易接受規則過去曾被有效地用於阻止使用區塊空間效率低下的用例的開發，而這樣做非常便宜。然而，這只有在使用者和礦工都對存在的任何替代方案感到滿意時才能有效。當情況不再如此，並且開發出與策略規則衝突的經濟上可行的用例時，使用者和礦工可以直接合作以避免對其活動施加限制的任何外部嘗試。事實上，能夠做到這一點是 Bitcoin 抗審查性的一個重要方面，具有優先對等的其他節點軟體也表明，規避絕大多數節點的過濾器相對容易。鑑於此，我們認為 Bitcoin 節點軟體最好旨在對下一個區塊中最終會包含什麼有一個現實的想法，而不是試圖在同意的交易建立者和礦工之間進行干預，以阻止在技術層面上基本無害的活動。

**這不是認可或容忍非金融資料使用，而是接受作為抗審查系統，Bitcoin 可以而且將被用於不是每個人都同意的用例。**

雖然我們認識到這一觀點並非所有使用者和開發者普遍持有，但我們真誠地相信這符合 Bitcoin 及其使用者的最佳利益，我們希望我們的使用者同意。我們將繼續作為開發者應用我們的最佳判斷，使交易接受規則與 Bitcoin 的長期健康和礦工的理性自利保持一致，包括特定的技術原因，例如升級安全性、高效的區塊建構和節點 DoS 攻擊。

簽署，

（支援此信的貢獻者列表）

* Andrew Toth
* Antoine Poinsot
* Anthony Towns
* Ava Chow
* b10c
* Bruno Garcia
* David Gumberg
* fjahr
* Gloria Zhao
* Gregory Sanders
* hodlinator
* ismaelsadeeq
* Josie Baker
* kevkevinpal
* l0rinc
* Marco De Leon
* Martin Zumsande
* Matthew Zipkin
* Michael Ford
* Murch
* Niklas Gögge
* pablomartin4btc
* Pieter Wuille
* Pol Espinasa
* Sebastian Falbesoner
* Sergi Delgado
* Stephan Vuylsteke
* TheCharlatan
* Vasil Dimov
* Will Clark
* w0xlt
