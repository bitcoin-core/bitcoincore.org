---
type: posts
layout: post
lang: zh_CN
name: segwit-benefits
id: en-segwit-benefits
title: 隔离见证的好处
permalink: /zh_CN/2016/01/26/segwit-benefits/
version: 1
excerpt: 隔离见证软分叉（segwit）包括很多的特性，其中许多是技术性很强的。
---
{% include _toc.html %}

隔离见证软分叉（segwit）包括很多的特性，其中许多是技术性很强的。 此页面总结了这些特性的一些好处。

## 修复延展性问题

比特币交易是由一串 64 位的十六进制哈希交易ID（TxID）标识的，这是基于交易中币的来源和币的接收者确定的。

不幸的是,txid 的计算方法允许任何人可以对交易做小的修动，虽然不会改变交易的内容，但会改变 txid。这就是所谓的第三方延展性。
BIP62 （"处理可塑性"）试图以一些方式解决这些问题，但是它太复杂，以至于无法实现为共识检查,所以已经被放弃了。

例如，您可以发送 txid 为 ef74... C309 到比特币网络，当网络中的节点中继这笔交易，或者矿工打包交易到区块中的时候，它们可以轻微修改这笔交易，
导致您的交易仍然花一样的币，并支付到相同的地址，但是以完全不同的txid683f...8bfa出现。

更通俗地说，如果一笔交易的一个或更多个签名者修改他们的签名, 那么交易仍然有效并且支付相同的比特币，以相同的地址，但这笔交易的txid完全改变，
因为它们包含签名。 更改签名数据（但不改变 output 或 input）来修改交易的情况称为 scriptSig 延展性。

Segwit 可以防止第三方和 scriptSig 延展性, 通过把比特币交易中的的可修改部分移动到见证交易里, 并且分离后不影响txid的计算。

### 谁将从中受益?

- **钱包作者用来监控发出比特币：** 这是最简单的，只需要监控发出的txid的状态就可以。但是在存在第三方延展性问题的系统里，钱包必须添加额外的代码，
以便能够应对变化的txids。

- **花费未确认的交易：** 如果 Alice 在交易1支付 Bob一些币，Bob 在交易2 使用收到的付款支付给Charlie，然后Alice的付款发生延展性修改,并用不同的txid确认，
那么交易2现在是无效的，而Charlie就不会被支付。如果Bob是值得信赖的，他会重新发出一笔交易给查理；但如果他不是，他可以简单地把这些比特币留给自己。

- **闪电网络：** 第三方和 scriptSig 延展性问题修复后可以降低闪电网络实现的的复杂性，而且在使用 blockchain 的空间上将更加有效. scriptSig
延展性删除后，它也可能运行轻量级的lighting客户端服务去监测区块链,而不是每个lightning客户端都运行比特币完整节点。

- **任何使用区块链的人:** 目前的智能合约，比如小额支付通道，预期新的智能合同，将会变得不用那么复杂的设计，理解和监控。

注意： segwit 交易只能在当所有input都是segwit交易（直接或经由一个向后兼容 的 segwit P2SH 地址）下避免延展性问题。

### 更多信息

 * [比特币维基延展性](https://en.bitcoin.it/wiki/Transaction_Malleability)
 * [关于延展性攻击的 2015 年比特币电报文章](http://cointelegraph.com/news/115374/the-ongoing-bitcoin-malleability-attack)
 * [关于延展性攻击的2015年比特币杂志文章](https://bitcoinmagazine.com/articles/the-who-what-why-and-how-of-the-ongoing-transaction-malleability-attack-1444253640)
 * [概述各种比特币改进建议对闪电交易的重要性](http://diyhpl.us/wiki/transcripts/scalingbitcoin/hong-kong/overview-of-bips-necessary-for-lightning/)
 * [BIP 62](https://github.com/bitcoin/bips/blob/master/bip-0062.mediawiki)
 * [比特币改进建议 140-延展性修复的替代方法](https://github.com/bitcoin/bips/blob/master/bip-0140.mediawiki)
 * [关于 683f…8bfa 交易在 Stack exchange 问答网络平台的回答](http://bitcoin.stackexchange.com/questions/22051/transaction-malleability-in-the-blockchain/22058#22058)

## 线性增长 sighash 的操作

用简单的方法来增加比特币区块大小的一个主要问题是，对于某些交易，签名散列增长是平方增长的, 而不是线性增长的

![Linear versus quadratic](/assets/images/linear-quad-scale.png)

在本质上，一个交易的大小增加一倍,签名操作的个数也增加一倍，以及那些进行验证签名需要哈希的数据也应该增加一倍. 但曾经已经出现过，一个单独的块需要 25 秒验证,其他 
一些恶意设计的交易可能需要超过 3 分钟。

Segwit 通过改变交易哈希签名的计算方式可以解决此问题，使得交易的每个字节只需要至多两次哈希。这提供了相同的功能但更有效率，
使得大的交易仍可以产生而不会有签名哈希问题，即使有人生成恶意的或更大的块（并较大的交易）也是支持的。

### 谁将从中受益?

删除验证签名需要的哈希数据的平方伸缩问题，使增长块大小更安全。这样做并没有 限制交易大小，所以仍然可以继续支持支付或者接收来自于大的组织的比特币,
比如挖矿奖励或众筹服务。

修改后的哈希仅适用于从 witness 数据发起签名操作，所以从旧的区块发起的签名操作将继续需要限制签名操作数下限。

### 更多信息

 * [BIP 143](https://github.com/bitcoin/bips/blob/master/bip-0143.mediawiki)
 * [在 25s 交易中 Rusty Russell 发布的博客](http://rusty.ozlabs.org/?p=522)
 * [CVE 2013-2292-比特币维基](https://en.bitcoin.it/wiki/Common_Vulnerabilities_and_Exposures#CVE-2013-2292)
 * [提议限制交易至 100KB](https://lists.linuxfoundation.org/pipermail/bitcoin-dev/2015-July/009494.html)
 * [在比特币 Classic0.11.2 版本中添加了限制sighash字节的额外共识](https://github.com/bitcoinclassic/bitcoinclassic/commit/842dc24b23ad9551c67672660c4cba882c4c840a)

## 值的签署

当硬件钱包签署一个交易的时候，它可以很容易地查证花费的总金额，但必须使用每个 input 交易的完整副本来确定是否安全，而且必须计算每个 input 的哈希以确保它们不是虚假数据。
个别交易大小可高达1MB大小，这不一定是一种廉价的操作，即使被签名的交易本身是相当小的。

Segwit 使 input 哈希变的精确从而解决了此问题。这意味着硬件的钱包可以简单地给出交易哈希，索引和值（和说明使用什么样的公钥），并可以放心地签署发出的交易，
无论花费的 input 交易有多大或多复杂。

### 谁将从中受益?

硬件钱包制造商和用户是明显的受益者。然而，这也使得它更容易，安全地在小型嵌入式设备的“物联网”的应用程序中使用比特币。

当消费收据发给 segwit 地址才可享有此优惠。 

### 更多资料

 * [BIP 143](https://github.com/bitcoin/bips/blob/master/bip-0143.mediawiki)
 
## 通過 multisig（P2SH) 增強的安全性

Multisig 付款目前使用 P2SH， 由 160 位 HASH160 算法（SHA256的RIPEMD）保护。但是，如果其中签名者，想盗取所有的资金，他们可以找到一个有效的地址之为
multisig 的一部分，只需支付他们所有的资金只有80位（280）的工作，这已经是一个极其资源充足的攻击者可做到的。 （为便于比较，在持续1 exahash /秒，
比特币挖掘网络每两个星期发掘80位值）

Segwit 通过使用 HASH160 付款直接到一个公共密钥（这种攻击是没用的），同时采用了256 位的散列SHA256付款给脚本解决此问题。

### 谁受惠?

每个人都支付给 multisig 或通过 segwit 提供的脚本得到额外保全。

### 更多资料

 * [Gavin Andresen询问 80 位的攻击是否值得担忧](https://lists.linuxfoundation.org/pipermail/bitcoin-dev/2016-January/012198.html)
 * [Ethan Heilman 描述发现算法圈](https://lists.linuxfoundation.org/pipermail/bitcoin-dev/2016-January/012202.html)
 * [Rusty Russell 计算进行攻击的成本](https://lists.linuxfoundation.org/pipermail/bitcoin-dev/2016-January/012227.html)
 * [Anthony Towns 用周期算法剥削交易](https://lists.linuxfoundation.org/pipermail/bitcoin-dev/2016-January/012218.html)
 * [Gavin Andresen 总结](https://lists.linuxfoundation.org/pipermail/bitcoin-dev/2016-January/012234.html)

## 脚本版

更改比特币的脚本允许提高安全性和改进功能。然而，脚本的设计只允许向后兼容（软分叉）通过更换其中一个新的操作码的十个额外 OP_NOP 操作码能让脚本失效，
要不就没有任何可能。这足够让许多变化产生 - 如引入一个新的签名方法或类似 OP_CLTV 功能，但两者无法单独操作（例如，OP_CLTV通常必须由一个 OP_DROP 陪同），
并不能用于作为连接两个字符串。

Segwit 包括脚本版本号解决此，在非segwit交易也需要用一个额外操作码，通过增加脚本版本支持。

### 谁受惠?

脚本操作码会使比特币高级脚本容易化。这包括改变，如引进的 Schnorr 签名，使用一键恢复收缩签名大小，支持侧链，或创建使用 Merklized 抽象语法
（MAST） 和其他研究级想法。

## 减少 UTXO 增长

（UTXO）数据库是由每个验证 Bitcoin 节点维持，以确定新的交易是否有效或欺诈。为网络的有效运行，该数据库必须非常快查询和修改，
并且理想地应当能够适合主存储器（RAM），因此保持该数据库的小是很重要的。

当比特币增长这就变得困难，因为每个新用户必须有自己的至少一个UTXO条目，每个用户有多个条目，保护隐私和灵活性，或提供为后盾，为支付渠道。

Segwit改善了的情况通过签名数据，这不会影响UTXO大小，成本比原本UTXO组大小的数据少75％。预计这将鼓励用户使用，大幅度减少费用，
并鼓励开发人员设计方式，也将减少对UTXO的影响智能合同和新功能设置UTXO影响。

因为segwit是软分叉变化并不增加基础大小，所述UTXO最坏情况下生长速度就保持不变。

### 谁受惠?

减少 UTXO 增长将有利于矿工，企业，完整的节点使用者，有助于维持在Bitcoin网络的安全随着更多的用户进入系统的用户。
用户和开发帮助减少UTXO的增长将享有较低的费用，相比忽视交易对UTXO增长的影响使用者。

### 更多资料

 * [Statoshi UTXO 排行榜](http://statoshi.info/dashboard/db/unspent-transaction-output-set)
 
## 提防诈骗

随着比特币用户群的扩大，验证整个链接变得更加昂贵。为了保持比特币的分散，不可信的性质，重要的是让不能对整个链接验证的用户至少能便宜地验证。

Segwit改善了情况并通过允许未来软叉延长证人结构并包括承诺数据，这将允许轻（SPV）客户端来执行一致的规则，例如，如在一并引入比特币的数量，
并在大量中数目中使用 sigops。

### 谁受惠?

提防欺诈允许SPV用户帮助执行比特币的共识规则，这将能增加比特币网络的整体安全性，以及减少个人用户被攻击。

这些欺诈行为的证据可以被添加到证人的数据结构作为未来软叉的一部分，他们将帮助SPV客户强制规定，就连交易不使用segwit功能都可。

## 不验证签名效率提高

历史交易签名可能不比未来签名有趣 - 例如，比特币核心最近的检查点在默认情况下并没有检查交易签名，有的SPV客户干脆不检查数字签名本身，
相信已经由矿工或其他节点检查。目前，签名数据事务处理的一个组成部分，必须存在，以便计算该事务散列。

分离署名数据允许签名数据从磁盘进行清理，避免下载它，节约资源的节点。

### 谁受惠?
随着越来越多的事务使用segwit地址，人们或SPV节点将能够用更少的数据与磁盘空间进行操作。

## 容量增加
旧节点将仅供下载，只执行对这些数据的1 MB大小限制规则。新的节点，理解见证数据的完整，因此免费用新的来取代这限制，允许更大的数据。
因此隔离见证利用这个机会提高限制为近4 MB，并增加了新的成本限制，以确保留在自己的资源利用平衡（这有效地限制接近1.6至2 MB）。

### 谁受惠?

谁运行升级钱包人们将能够通过移动签名交易的见证人部分利用增加的块大小。

## 迈向单一组合限制

目前，有两个保守的强制限制：该单一组合不超过 1MB，并在整个交易执行不可超过 20,000 签名检查。

寻找最有利的交易集合给出单一的限制以包括是背包问题，可以用一个简单的算法来容易地解决几乎完美的实例。然而加入第二约束使得在某些情况下，
很难找到一个很好的解决方案，这一理论问题在实践中已经迫使单一数据远低于容量待开采。

如没有 hardfork 或基本上减小单一数据是无法解决这个问题。由于 segwit 无法修复，在不使情况变得更糟情况下：与其引入一个独立限制分离的见证数据，
它采用单一限制适用于UTXO数据的加权和证人数据，同时允许合并后的实体限制。

### 谁受惠?

未来hardfork如能改变单一容量限制，使用者就能受惠。 如：

    50* sigops + 4*basedata + 1*witnessdata < 10M

使用者能轻易准确地填写单一数据同时最大限度并提高手续费收入，这让他们更可靠地计算出需要进行开采其交易的费用让用户受益。

### 更多资料

 * [背包问题](https://en.wikipedia.org/wiki/Knapsack_problem)
 * [2015 年 8 月 Sigop 攻击 bitcoin 的讨论](https://bitcointalk.org/index.php?topic=1166928.0;all)
 * [Gregory Maxwell 对比特币-dev证人的限制 ](https://lists.linuxfoundation.org/pipermail/bitcoin-dev/2015-December/011870.html)
 * [“验证成本度量” 成绩](http://diyhpl.us/wiki/transcripts/scalingbitcoin/hong-kong/validation-cost-metric/)
