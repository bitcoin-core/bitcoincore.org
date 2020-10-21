---
title: 如何貢獻代碼 Bitcoin 核心
name: contributing-guidelines
id: zh_TW-contributing-guidelines
permalink: /zh_TW/faq/contributing-code/
layout: page
type: pages
lang: zh_TW
category: faqs
version: 2
---

Bitcoin 核心項目按照開放式貢獻者模式運作，歡迎任何人以同行評價、測試和打補丁的形式為它的發展做出貢獻。本文件說明了做出貢獻的操作流程和指導。

首先，在結構方面，沒有所謂的特權 “核心開發人員”的特殊概念。開源通常自然而然地圍繞有才能的人展開，而長期貢獻者會獲得開發者社區更多的信任。
但是出於實用的角度某些層次是必不可少的。因此有知識庫“維護者們”負責合併拉取請求以及一個“牽頭維護者”負責發布週期、整體合併、模式化和任命維護人員。

貢獻者工作流程
-----------------

代碼庫通過“貢獻者工作流程”維護，在這裡所有人都無一例外地採用拉取請求貢獻補丁提案。這有利於促進社區貢獻、簡單測試和同行評審。

貢獻補丁的工作流程如下：

  - 對版本庫進行分叉
  - 創建主題分支
  - 提交補丁

必須遵守 [開發手冊](https://github.com/bitcoin/bitcoin/blob/master/doc/developer-notes.md) 規定的項目代碼規範。

通常情況下 [提交應該是原子化的](https://en.wikipedia.org/wiki/Atomic_commit#Atomic_commit_convention) 並且差異應該容易閱讀。
出於這個原因請不要將任何格式修正和代碼移除與真正在的代碼修改混在一起。

提交訊息應當有詳細描述， 預設是由一行簡短的主題行（最多50個字）、一行空行加詳細的說明文字組成的段落；除非標題已經是不言自明的
（如“更正 main.cpp 中的拼寫錯誤”），那麼一個標題行就足夠了。提交的訊息應該有助於將來其他人閱讀您的代碼，因此請您解釋做此決定的理由。
進一步解釋請 [點擊這裡](http://chris.beams.io/posts/git-commit/)。

如果某個提交提及另外一個問題，請添加參考訊息，例如`關於＃1234`，或`修復＃4321`。使用`修復`或`關閉`關鍵字將導致拉取請求合併過程中相應的問題被關閉。

關於 Git 的更多問題請參考 [Git 操作手冊](https://git-scm.com/doc)。

  - 將變更推送到你的分叉
  - 創建拉取請求

拉取請求的標題應該將拉取請求所影響的組件或領域加為前綴。 範例：

    共識: 為 BIP-XXXX OP_CHECKAWESOMESIG 添加新的 opcode
    淨: 自動生成隱藏服務，監聽 Tor
    Qt: 添加新的 bump 按鈕
    不重要的: 修改 main.cpp 的拼寫錯誤

如果一個專門的拉入請求（尚且）不應被合併請使用 [WIP] 做為標題前綴或使用拉取請求主體中的 [任務列表](https://help.github.com/articles/basic-writing-and-formatting-syntax/#task-lists) 來表示任務仍未完成。

拉取請求的主體應該就補丁的用途以及理由/原因給予足夠的描述。您應該引用任何相關的討論（例如其他問題單或郵件列表討論）。

在這個階段提交者應當預期收到其他貢獻者的意見和評審。您可以透過本地提交的方式多次提交代碼，然後將他們推到你的分叉上，直到滿足所有回饋意見。

改寫提交
-----------------
如果您的拉取請求被接受予以合併，維護者可能要求您在合併前改寫和／或 [重新基線](https://git-scm.com/docs/git-rebase) 您提交的代碼。
基本的 squash 工作流程如下所示。

    git 簽出 your_branch_name
    git 重新基線 -i HEAD~n
    # n 通常代表一個拉取請求裡提交的次數
    # 將提交從 'pick' 設置成 'squash'，保存並推出
    # 在下一個螢幕上，修改／潤色提交訊息
    #保存並推出
    git push -f # (強制推到 GitHub)

同行評審需要的時間不可預期並且每個拉取請求的評審時間都不盡相同。


拉取請求原則
-----------------------

補丁集應該有側重點。例如，一個拉取請求可能添加一個功能、修復一個缺陷或者重構代碼；但不能是是這些的混合。也請避免提交太多內容、
太大或太複雜的超級拉取請求，因為很難評審。

### 功能

在添加新功能時，必須要考慮到引入代碼後的長期技術債務和對該功能的維護。在提議一個需要維護的新功能前， 請考慮一下你是否願意對它進行維護（包括錯誤修復）。如果將來功能變成沒人維護的孤兒，他們可能會被知識庫維護者刪除。

### 重構

重構是任何軟體項目演進的必要組成部分。下列指導原則涵蓋了項目的重構拉取請求。


共有三個類別的重構，只移動代碼、代碼風格修復和代碼重構。一般來說，為了使重構拉入請求易於評審和無爭議的，不應混用這三種活動。 
在所有情況下，重構PR一定不能在拉取請求中改變代碼的行為（錯誤必須原樣保留）。


項目維護者的目標是重構拉取請求能快速周轉，所以儘可能保持簡短、不複雜和容易驗證。



決策流程
-------------------------

以下適用於 比特幣核心項目（以及相關的項目，如libsecp256k1）的代碼更改，不要與整個比特幣網路協議共識的變更相混淆。
一個拉取請求是否被合併到比特幣核心取決於項目合併維護者，並最終取決於項目牽頭人。

只要與項目的總體方針相符合的補丁維護者都會予以考慮；達到引入的最低條件；並判斷貢獻者的普遍共識。


總的來說，所有的拉取請求都必須：

  - 有清晰的用例、修復一個可證明的錯誤或提高程序整體能力（例如為了模組化而重構）;
  - 經過同行充分評審；
  - 在適當情況下進行單元測試和功能測試；
  - 遵守代碼風格指南；
  -不破壞現有的測試套件；
  - 只要進行了錯誤修復，在可能的情況下都應該進行單元測試來證明缺陷和修復。這有助於預防疊代。

改變比特幣共識規則的補丁參與的人會顯著增加，因為它們影響整個生態系統，因此事前必須進行大量的郵件列表討論，並有BIP編號。
雖然每一個案例都會有所不同，但由於增加的同行評審和共識條件的建立，大家應當準備好投入比其他各類補丁更多的時間。

### 同行評審

任何人都可以通過發布拉取請求評審意見參與同行評審。通常評審人會評審代碼中的明顯錯誤，同時測試補丁組並對補丁的技術優點給出意見。
在確定是否存在共識合併拉取請求時（記住，討論可能已經在github、郵件列表和IRC討論區展開了），項目維護者會考慮同行評審的意見。拉取請求評論中會使用以下語言：

  - ACK 指 "我已經測試過代碼並且同意將其合併 "；
  - NACK指"我不同意將其合併"，並必須提交合理的技術理由。不附帶理由的NACKs將被忽略。 
  - utACK指"我還沒有測試過代碼，但我檢查過代碼並認為可行，我同意將其合併"；
  - Concept ACK指"我原則上同意這個拉取請求 "；
  - Nit 指小問題，通常不會造成阻礙的為題。

評審人應當將他們在評審意見裡評審過的提交哈西表包括在哪 。

項目維護者保留使用常識判斷來衡量評審人意見的權利，並也可能基於才能進行權衡：那些（長期以來）
已經展示了對項目更大的承諾和認識或有明確的專業領域知識的人自然會有更多的權重，正如人們在各行各業所預期的一樣。

凡是影響關鍵共識代碼的補丁集，討論和同行評審要求的門檻就會高出很多，請牢記失誤對更大範圍的社區來說可能代價慘重。這也包括對核心共識代碼的重構。

凡是提議更改比特幣共識的補丁集，一定是已經在郵件列表和IRC上深入討論過的，伴有大規模討論過的BIP，並且是基於維護者的判斷達成了普遍技術認同的，
被認為值得的一個改變。


版本發布政策
--------------

項目帶頭人是每個比特幣核心版本發布的發布經理。