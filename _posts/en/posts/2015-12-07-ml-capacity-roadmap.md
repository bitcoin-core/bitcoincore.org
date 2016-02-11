---
title: Capacity increases Roadmap for the Bitcoin system
permalink: /en/2015/12/07/roadmap/
lang: en
id: en-ml-roadmap-2015-12-07
name: ml-roadmap-2015-12-07
type: posts
layout: post
canonical: https://lists.linuxfoundation.org/pipermail/bitcoin-dev/2015-December/011865.html
version: 1
---
_The following roadmap was originally posted to the [bitcoin-dev mailing list](https://lists.linuxfoundation.org/pipermail/bitcoin-dev/2015-December/011865.html), by Gregory Maxwell on 2015-12-07._

The Scaling Bitcoin Workshop in HK is just wrapping up. Many fascinating proposals were presented.
I think this would be a good time to share my view of the near term arc for capacity increases in the Bitcoin system.
I believe we’re in a fantastic place right now and that the community is ready to deliver on a clear forward path with a shared vision that addresses the needs of the system while upholding its values.

I think it’s important to first clearly express some of the relevant principles that I think should guide the ongoing development of the Bitcoin system.

Bitcoin is P2P electronic cash that is valuable over legacy systems because of the monetary autonomy it brings to its users through decentralization. Bitcoin seeks to address the root problem with conventional currency: all the trust that's required to make it work--

-- Not that justified trust is a bad thing, but trust makes systems brittle, opaque, and costly to operate.
Trust failures result in systemic collapses, trust curation creates inequality and monopoly lock-in, and naturally arising trust choke-points can be abused to deny access to due process.
Through the use of cryptographic proof and decentralized networks Bitcoin minimizes and replaces these trust costs.

With the available technology, there are fundamental trade-offs between scale and decentralization.
If the system is too costly people will be forced to trust third parties rather than independently enforcing the system's rules.
If the Bitcoin blockchain’s resource usage, relative to the available technology, is too great, Bitcoin loses its competitive advantages compared to legacy systems because validation will be too costly (pricing out many users), forcing trust back into the system.
If capacity is too low and our methods of transacting too inefficient, access to the chain for dispute resolution will be too costly, again pushing trust back into the system.

Since Bitcoin is an electronic cash, it _isn't_ a generic database; the demand for cheap highly-replicated perpetual storage is unbounded, and Bitcoin cannot and will not satisfy that demand for non-ecash (non-Bitcoin) usage, and there is no shame in that.
Fortunately, Bitcoin can interoperate with other systems that address other applications, and--with luck and hard work--the Bitcoin system can and will satisfy the world's demand for electronic cash.

Fortunately, a lot of great technology is in the works that make navigating the trade-offs easier.

First up: after several years in the making Bitcoin Core has recently merged libsecp256k1, which results in a huge increase in signature validation performance.
Combined with other recent work we're now getting ConnectTip performance 7x higher in 0.12 than in prior versions. This
has been a long time coming, and without its anticipation and earlier work such as headers-first I probably would have been arguing for a block size decrease last year.
This improvement in the state of the art for widely available production Bitcoin software sets a stage for some capacity increases while still catching up on our decentralization deficit. This shifts the bottlenecks off of CPU and more strongly onto propagation latency and bandwidth.

Versionbits (BIP9) is approaching maturity and will allow the Bitcoin network to have multiple in-flight soft-forks. Up until now we’ve had to completely serialize soft-fork work, and also had no real way to handle a soft-fork that was merged in core but rejected by the network.
All that is solved in BIP9, which should allow us to pick up the pace of improvements in the network. It looks like versionbits will be ready for use in the next soft-fork performed on the network.

The next thing is that, at Scaling Bitcoin Hong Kong, Pieter Wuille presented on bringing Segregated Witness to Bitcoin.
What is proposed is a _soft-fork_ that increases Bitcoin's scalability and capacity by reorganizing data in blocks to handle the signatures separately, and in doing so takes them outside the scope of the current blocksize limit.

The particular proposal amounts to a 4MB blocksize increase at worst. The separation allows new security models, such as skipping downloading data you're not going to check and improved performance for lite clients (especially ones with high privacy).
The proposal also includes fraud proofs which make violations of the Bitcoin system provable with a compact proof.
This completes the vision of "alerts" described in the "Simplified Payment Verification" section of the Bitcoin whitepaper, and would make it possible for lite clients to enforce all the rules of the system (under a new strong assumption that they're not partitioned from someone who would generate the proofs).
The design has numerous other features like making further enhancements safer and eliminating signature malleability 
problems. If widely used this proposal gives a 2x capacity increase (more if multisig is widely used), but most importantly it makes that additional capacity--and future capacity beyond it--safer by increasing efficiency and allowing more trade-offs (in particular, you can use much less bandwidth in exchange for a strong non-partitioning assumption).

There is a working implementation (though it doesn't yet have the fraud proofs) at <https://github.com/sipa/bitcoin/commits/segwit>.

(Pieter's talk is at: [transcript](http://diyhpl.us/wiki/transcripts/scalingbitcoin/hong-kong/segregated-witness-and-its-impact-on-scalability/), [slides](https://prezi.com/lyghixkrguao/segregated-witness-and-deploying-it-for-bitcoin/), [Video](https://www.youtube.com/watch?v=fst1IK_mrng#t=36m)).

I had good success deploying an earlier (hard-fork) version of segwit in the Elements Alpha sidechain; the soft-fork segwit now proposed is a second-generation design. And I think it's quite reasonable to get this deployed in a relatively short time frame.
The segwit design calls for a future bitcoinj compatible hardfork to further increase its efficiency--but it's not necessary to reap most of the benefits,and that means it can happen on its own schedule and in a non-contentious manner.

Going beyond segwit, there has been some considerable activity brewing around more efficient block relay. There is a collection of proposals, some stemming from a p2pool-inspired informal sketch of mine and some independently invented, called "weak blocks", "thin blocks" or "soft blocks".
These proposals build on top of efficient relay techniques (like the relay network protocol or IBLT) and move virtually all the transmission time of a block to before the block is found, eliminating size from the orphan race calculation. We already desperately need this at the current block sizes. These have not yet been implemented, but fortunately the path appears clear.
I've seen at least one more or less complete specification, and I expect to see things running using this in a few months. This tool will remove propagation latency from being a problem in the absence of strategic behavior by miners.  Better understanding their behavior when miners behave strategically is an open question.

Concurrently, there is a lot of activity ongoing related to “non-bandwidth” scaling mechanisms.
Non-bandwidth scaling mechanisms are tools like transaction cut-through and bidirectional payment channels which increase Bitcoin’s capacity and speed using clever smart contracts rather than increased bandwidth.
Critically, these approaches strike right at the heart of the capacity vs autotomy trade-off, and may allow us to achieve very high capacity and very high decentralization. CLTV (BIP65), deployed a month ago and now active on the network, is very useful for these techniques (essential for making hold-up refunds work); CSV (BIP68 / BIP112) is in the pipeline for merge in core and making good progress (and will likely be ready ahead of segwit).
Further Bitcoin protocol improvements for non-bandwidth scaling are in the works: Many of these proposals really want anti-malleability fixes (which would be provided by segwit), and there are checksig flag improvements already tendered and more being worked on, which would be much easier to deploy with segwit.
I expect that within six months we could have considerably more features ready for deployment to enable these techniques. Even without them I believe we’ll be in an acceptable position with respect to capacity in the near term, but it’s important to enable them for the future.

<http://diyhpl.us/wiki/transcripts/scalingbitcoin/hong-kong/overview-of-bips-necessary-for-lightning> is a relevant talk for some of the wanted network features for Lightning, a bidirectional payment channel proposal which many parties are working on right now; other non-bandwidth improvements discussed in the past include transaction cut-through, which I consider a must-read for the basic intuition about how transaction capacity can be greater than blockchain capacity: <https://bitcointalk.org/index.php?topic=281848.0>, though there are many others.)

Further out, there are several proposals related to flex caps or incentive-aligned dynamic block size controls based on allowing miners to produce larger blocks at some cost.
These proposals help preserve the alignment of incentives between miners and general node operators, and prevent defection between the miners from undermining the fee market behavior that will eventually fund security.
I think that right now capacity is high enough and the needed capacity is low enough that we don't immediately need these proposals, but they will be critically important long term.
I'm planning to help out and drive towards a more concrete direction out of these proposals in the following months.

(Relevant talks include <http://diyhpl.us/wiki/transcripts/scalingbitcoin/hong-kong/a-flexible-limit-trading-subsidy-for-larger-blocks/>.

Finally--at some point the capacity increases from the above may not be enough.
Delivery on relay improvements, segwit fraud proofs, dynamic block size controls, and other advances in technology will reduce the risk and therefore controversy around moderate block size increase proposals (such as 2/4/8 rescaled to respect segwit's increase).
Bitcoin will be able to move forward with these increases when improvements and understanding render their risks widely acceptable relative to the risks of not deploying them.
In Bitcoin Core we should keep patches ready to implement them as the need and the will arises, to keep the basic software engineering from being the limiting factor.

Our recent and current progress has well positioned the Bitcoin ecosystem to handle its current capacity needs.
I think the above sets out some clear achievable milestones to continue to advance the art in Bitcoin capacity while putting us in a good position for further improvement and evolution.

TL;DR:  I propose we work immediately towards the segwit 4MB block soft-fork which increases capacity and scalability, and recent speedups and incoming relay improvements make segwit a reasonable risk. BIP9 and segwit will also make further improvements easier and faster to deploy.
We’ll continue to set the stage for non-bandwidth-increase-based scaling, while building additional tools that would make bandwidth increases safer long term.
Further work will prepare Bitcoin for further increases, which will become possible when justified, while also providing the groundwork to make them justifiable.

Thanks for your time,

_Originally posted to the bitcoin-dev mailing list at <https://lists.linuxfoundation.org/pipermail/bitcoin-dev/2015-December/011865.html>, by Gregory Maxwell on 2015-12-07_
