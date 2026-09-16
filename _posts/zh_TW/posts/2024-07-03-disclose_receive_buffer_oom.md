---
title: CVE-2015-3641 披露
name: blog-disclose-receive-buffer-oom
id: zh_TW-blog-disclose-receive-buffer-oom
lang: zh_TW
type: advisory
layout: post

## If this is a new post, reset this counter to 1.
version: 1

## Only true if release announcement or security annoucement. English posts only
announcement: 1

excerpt: >
  攻擊者傳送大型不完整訊息會導致高記憶體使用量。修復程式於 2015 年 4 月 27 日在 Bitcoin Core 0.10.1 中發布。
---

節點可能被迫在接收訊息時分配大型緩衝區，這可能被利用透過 OOM 遠端崩潰它。

此問題被認為是**中等**嚴重性。

## 細節

在沒有更嚴格的界限的情況下，接收訊息的大小受限於最大序列化訊息大小 32 MiB。攻擊者可以強制節點為每個連線分配這麼多 RAM，這可能導致 OOM。

[PR #5843](https://github.com/bitcoin/bitcoin/pull/5843) 減少了 P2P 訊息在接收有效載荷之前可以擁有的大小。這減少了惡意對等節點可以導致的每個對等節點接收緩衝區記憶體大小。該 PR 將數字從 32 MiB 減少到 2 MiB，後來作為 Segwit BIP144 變更的一部分增加回 4 MB。

## 歸屬

由 bitcointalk 使用者 Evil-Knievel 向 Greg Maxwell 報告。由 Pieter Wuille 修復。

## 時間表

- 2015-02-05 Evil-Knievel 透過 bitcointalk 私人訊息向 Greg Maxwell 報告該漏洞。
- 2015-??-?? 為其註冊 `CVE-2015-3641`。
- 2015-03-01 開啟 [PR #5843](https://github.com/bitcoin/bitcoin/pull/5843) 以修復它。
- 2015-03-06 PR #5843 被合併。
- 2015-03-09 修復程式被回移植到版本 0.10.1。
- 2015-04-27 Bitcoin Core 版本 [0.10.1 發布](https://lists.linuxfoundation.org/pipermail/bitcoin-dev/2015-April/007828.html)，包含修復程式。
- 2015-06-25 披露被[預先宣布](https://lists.linuxfoundation.org/pipermail/bitcoin-dev/2015-June/009135.html)。
- 2015-07-07 披露被[推遲](https://lists.linuxfoundation.org/pipermail/bitcoin-dev/2015-July/009362.html)。
- 2016-08-23 最後一個易受影響的 Bitcoin Core 版本（0.10.x）終止支援
- 2024-07-03 公開披露。

{% include references.md %}
