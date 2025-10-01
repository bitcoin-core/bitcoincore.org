---
type: posts
lang: zh_CN
name: scaling-optimizations-zh_ch
id: zh_cn-scaling-optimizations-zh_ch
title: 链上扩容：比特币参考软件性能优化历史回顾
permalink: /zh_CN/2017/03/11/比特币参考软件性能优化历史回顾/
version: 1
---

最近几年为维持可靠的用户体验，比特币软件进行了很多里程碑式的发展，本文对这些重大发展进行了回顾。我们对这个软件进行的一些升级对维护比特币网络去中心化特征、缓解参与者负担来说是非常重要的。本文描述了我们对这个软件所做的大量成倍提升性能的优化，这些优化让比特币网络能够容纳更多的交易活动，同时没有给新参与者和现有参与者增加过多的参与成本。最后，我们解释了为什么说这些改善是一个更大且系统化协议开发方式的一部分，这种开发方式从Big-O复杂概念中获得的启发，并使用了更精妙的算法让这个网络的资源发挥出更大的效力。
 
## 签名缓冲
发布版本：Bitcoin-Qt 0.7.0
ECSDA签名验证是节点层面最耗计算资源的任务之一。因为每个交易都要进行这样的验证，所有任何多余的验证都会让终端用户消耗大量的资源。这样就造成这样一种情况：早期版本软件在交易进入一个节点内存池前和被接收到一个区块后都要对签名进行验证。
为了提高效率，开发者创建出一个缓存，可以存储之前验证过的签名，一旦这些交易被打包进区块，很多冗余的工作量就能省略。
另外，这个签名缓存还可以消除一个DoS攻击路径，通过构造一些特殊结构的交易来使比特币客户端的运行缓慢。
 
### 更多相关信息请参考：
 * [Bitcoin-Qt 0.7.0 Release notes](https://bitcoin.org/en/release/v0.7.0#core-bitcoin-handling-and-blockchain-database)
 * [Fixed vulnerability explanation: Why the signature cache is a DoS protection](https://bitcointalk.org/index.php?topic=136422.0)

 
## Ultraprune + LevelDB
发布版本：Bitcoin-Qt 0.8.0
Ultraprune（区块链高度精简）是比特币软件最重要的升级之一，它解决了验证区块链上的交易时出现的资源开支问题。比特币参考软件之前版本维护着一个针对所有交易输出（已花的和未花的）的索引，Ultraprune大幅削减了这个索引的体积，因为你只需跟踪未花交易输出，一旦一个输出被花掉了，那么就可以把这个输出从这个索引中删掉。
为了实现这些，开发者实施了一个新的数据库布局，可以把未花交易输出压缩进一个紧凑的格式中，可以减少验证工作所需信息的体积。
为了进一步优化这个系统的性能，LevelDB和Ultraprune一起被引入，老的BDB数据库技术被弃用。整体上来看，影响是显著的：虽然不同硬件上体验到的改善程度会有所不同，但是验证区块数据时的性能至少提高了一个数量级。这个新数据库结构也为以后对比特币全节点的精简和轻量化铺平了道路。
 
### 更多相关信息请参考：
* [Bitcoin-Qt 0.8.0 Release notes](https://bitcoin.org/en/release/v0.8.0#improvements)
* [Ultraprune in plain english](https://archive.is/bUocJ)
* [Ultraprune merged in mainline](https://bitcointalk.org/index.php?topic=119525.0)
* [Pruning in the reference client: ultraprune mode](https://bitcointalk.org/index.php?topic=91954.0)

 
## 并行脚本验证
发布版本：Bitcoin-Qt 0.8
虽然并不是一个很明显的改变，让交易脚本验证过程更加并行化，可以使区块验证时间显著减少。早期版本软件在接收每个UTXO（未花的交易輸出）之间验证来自交易输入的数据，这造成了一种性能问题，因为所有这些动作都是线性地进行的。这有违高效计算过程的设计原则，高效计算原则规定：在可能的情况下，计算应该和I/O工作同时进行。因此，区块验证机制被重新设计，以便能并行处理脚本检验，这样一来，在客户端从区块接收到所有UTXO前验证就能开始进行。为了实现这功能，脚本检查动作在交易被处理后，被放置进一个队列中，与其它输入验证的相关操作分开进行。
这样优化的结果是，节点资源的使用更高效，区块同步更快。根据测试结果来看，对比以前版本，这个版本的速度提高了35%到100%。
 
### 更多相关信息请参考：
* [Parallel script verification #2060](https://github.com/bitcoin/bitcoin/pull/2060)

 
## 区块头先同步
发布版本：Bitcoin Core 0.10
为了进一步改善初始区块下载时间，Core项目在2014年末，重新设计了让节点与包含最大工作量的区块链进行同步的机制的架构。
起初，一个新比特币客户端的引导进程会和用户从一个单节点处接收区块这个过程牵扯到一起，这产生的结果是，连接的中断或连接质量下降会拖累这个进程。随着区块链体积的不断增加，这个问题有时会让同步时间显著增加，有很多用户报告完成这个启动过程的时间长达几天。
区块头先同步完全解决了这个问题，它让节点从一个单节点处先下载、验证区块头数据，然后并行地从其它多个节点处接收区块数据。
自比特币早期时期以来，人们就对初始区块下载时间有抱怨。采用了区块头先同步这种方式，这个软件对新用户来说实用性又向前跨越了一步。现在用户可以利用整个网络的节点，大幅削减引导时间，而不必在不稳定的同步中浪费大量时间。这方面对比特币的长期可持续性来说是非常重要的，开发者使用更精妙的算法渐进式地在这方面进行改善。
 
### 更多相关信息请参考：
* [Bitcoin-Qt 0.10.0 Release notes](https://bitcoin.org/en/release/v0.10.0#faster-synchronization)

* [Bitcoin.org Developer Guide](https://bitcoin.org/en/developer-guide#headers-first)

* [Pieter Wuille’s post to the Bitcoin-dev mailing list](https://lists.linuxfoundation.org/pipermail/bitcoin-dev/2014-October/006724.html)
 
## 区块文件修剪
发布版本：Bitcoin Core 0.11
 
旧数据修剪这个概念在中本聪的白皮书中首次被提出，可解决磁盘空间稀缺问题。不幸的是，这个原始设计是不充分的，无法按照比特币创始人的想象来实施。7年之后，区块链体积达到100多G，区块文件修剪功能对资源有限的用户来说帮助是巨大的。
区块文件修剪使用了上文提到的ultraprune功能；用户现在可以把288区块之前的旧区块数据删去，这些节点依然可以保护完整的UTXO集，也能够验证未花输出数据，这足以完整地验证新区块并对抗潜在的双花。
这样的区块文件修剪假设仍旧会有足够多的存档节点来维护完整区块链数据。另一方面，这种创新让维护节点的成本更低从而扩大了验证者的范围。总而言之，这种解决方案有助于我们维持网络多样化。
 
### 更多相关信息请参考：
* [Bitcoin-Qt 0.11.0 Release notes](https://bitcoin.org/en/release/v0.11.0#block-file-pruning)

 
## libsecp256k1
发布版本：Bitcoin Core 0.12
根据测算确定，在解决区块链下载的效率问题后要处理的问题是交易验证瓶颈及其繁重的计算负载。为了解决问题，Core项目启用了专门优化ECDSA操作的新代码库。ECDSA（椭圆曲线数字签名算法）是比特币公钥架构的基石，用户每次用私钥签名来发送比特币时都要用到它。这些签名需要被这个网络中的每个节点验证，以保证这个账簿的完整性。
早期开发者很早就考虑脱离原始的OpenSSL库，在经过5年时间的设计考量、测试、同级评审后，Pieter Wuille的libsecp256k1库被引入作为它的替代者。就像预期的那样，库被替换后每个比特币交易的签名验证速度明显加快。对比报告显示，比使用OpenSSL快了5～7倍。对于用户来说就是，引导时间最多能节省一半，而这些时间都是从ECSDA操作过程中省出来的，它是从零开始同步新节点最繁重的步骤之一。
比特币交易活动日益增多，这次更新为网络节点维持合理的用户体验至关重要。而且，降低了算法复杂性，让用户更高效地利用他们的资源，为新参与者降低了参与门槛。
 
### 更多相关信息请参考：
* [Bitcoin-Qt 0.12.0 Release notes](https://bitcoin.org/en/release/v0.12.0#signature-validation-using-libsecp256k1)
* [Andrew Poelstra (andytoshi) on security and testing of - libsecp256k1](https://bitcointalk.org/index.php?action=profile;u=80376)
* [Greg Maxwell on testing of libsecp256k1 revealing bug in OpenSSL](https://www.reddit.com/r/Bitcoin/comments/2rrxq7/on_why_010s_release_notes_say_we_have_reason_to/)
* [Greg Maxwell presentation at DevCore](https://www.youtube.com/watch?v=RguZ0_nmSPw&t=1297)
* [Hal Finney post on libsecp256k1](https://bitcointalk.org/index.php?topic=3238.0)

 
## 内存池限制
发布版本：Bitcoin Core 0.12
比特币长期存在的一个薄弱点是它不能恰当地应对节点内存池的泛洪（flooding），攻击者可以发起大量的低价值、低交易费的交易，这些交易在内存池中堆积，直到内存超载。这会导致内存资源不足的节点在非正常活动时期陷入崩溃。对抗这种攻击的唯一有效措施是提高软件的最小中继费，不过此时内存池的体积依然没有上限。
为了解决这个问题，Core项目开发者设置了内存池最大体积限制，使用具体的筛选政策，把交易按照交易费排序，并首先把费用较小的交易给筛选出去。为了避免相同费用的交易重新进入内存池，节点会提高它的有效最小中继费率，把初始的最小中继费率加上最后被筛选出去的交易的交易费。
用户可以自己设置最大体积，默认设置是300MB，这次更新为资源有限的节点用户带来更健壮的体验，整体上来说，让整个网络更加可靠。
 
### 更多相关信息请参考：
* [Bitcoin-Qt 0.12.0 Release notes](https://bitcoin.org/en/release/v0.12.0#memory-pool-limiting)

 
在第二部分，我们将讨论更多基于上述技术的优化和对网络健壮性、扩展潜力的改善。
