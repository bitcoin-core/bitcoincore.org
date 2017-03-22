---
type: posts
layout: post
lang: en
name: schnorr-sigagg
id: en-schnorr-sigagg
title: Technology roadmap: Schnorr signatures & signature aggregation
permalink: /en/2017/03/22/schnorr-sigagg/
version: 1
excerpt: Status and explanation of Schnorr signatures and signature aggregation
---

The following post aims to highlight development milestones that helped preserve a reliable experience for users of the Bitcoin software client over the years. We present several upgrades that were critical in maintaining the decentralized properties of the network and mitigate the resources burden of its participants. We describe how numerous orders-of-magnitude optimizations were made so that the Bitcoin network could support the growth in transaction activity without dramatically increasing the costs of participation for new and existing users. Finally, we note how those improvements all fall within a larger, systematic approach to protocol development that uses insights from Big-O complexity concepts and leverages smarter algorithms that make more efficient use of the network’s resources. 

The replacement of Bitcoin’s digital signature algorithm (ECSDA) for the more efficient Schnorr algorithm has long been at the top of the bucket list for many Bitcoin developers. A simple algorithm leveraging elliptic curves cryptography, Schnorr enables several improvements over the existing scheme all while preserving all of the features and security assumptions of ECSDA.

Notably, Schnorr signatures support “native multisig” which enables the aggregation of multiple signatures into a single one valid for the sum of the keys of their respective inputs. This functionality offers three important benefits: 

Constant-size signatures irrespective of the number of participants in the multisig setup. A 50-of-50 transaction is effectively the same size as one that uses a single public key and signature. For this reason, the performance of such schemes is significantly improved by removing the original requirement of validating every signature individually. Additionally, the verification of Schnorr signatures is slightly faster than that of ECSDA.


The diminished size of data to be validated and transmitted across the network also translates in interesting capacity gains. Considering the historical growth in the number of multisig transactions displayed below, the potential to reduce the size of these transactions is an enticing addition to existing scaling efforts. We should expect this trend to continue with the emergence and further adoption of payment channels.


From a privacy standpoint, Schnorr allows the entire policy of the multisig to be obscured and indistinguishable from a conventional single pubkey. In a threshold setup, it also becomes impossible for participants to reveal which of them authorized, or not, a transaction. 


<p align="center">
  <img src="https://cdn-images-1.medium.com/max/800/1*9_03N_A0OlLkcJpCGq--Qw.png">
</p>
<p align="center">
  Distribution of unspent P2SH outputs according to their multisig setup. Source: p2sh.info 
</p>

Unfortunately, unlike ECSDA, the Schnorr algorithm has not been standardized since its invention, likely because of the original patent enforced on it (which has since expired). While the general outlines of the system are mathematically sound, the lack of documentation and specification makes it more challenging to implement. Specifically, its application to the ephemeral keypairs design of Bitcoin involves security considerations that require further optimization. 

The main challenge is defined by Pieter Wuille in his Scaling Bitcoin Milan presentation of Schnorr signatures as the “cancellation” problem. The possibility for a group of users to create a signature valid for the sum of their keys opens the door for an adversarial participant to subtract from the whole another user’s key. It essentially works like this:

Assume a 2-of-2 multisig scheme using input public keys Q1 and Q2. Rather than announce their key as Q2 to be combined with Q1, a malicious participant could provide, during the interaction phase, Q2-Q1 and effectively cancel out the other user’s key. Any fund sent to the joint public key is now only spendable by the owner of the Q2 key without the owner of Q1 even being aware of what is going on. 

Fortunately, a solution is now available which involves multiplying every key used during the setup with a hash based on itself and all other keys involved before signing. This process is called delinearization. A proof of the security of this scheme is currently undergoing peer-review and will be formally described in an upcoming whitepaper. 

In the near term, Schnorr signatures are being considered as viable replacement for two important functions of the Bitcoin protocol: OP_CHECKSIG & OP_CHECKMULTISIG. 

The former is currently used to check ECSDA signatures against their respective public key according to the message in a transaction. By switching to an equivalent that checks for Schnorr signatures rather than ECSDA, the opcode can be used to authorize a spend requiring multiple signatures which would typically require calling OP_CHECKMULTISIG. Using a priori interaction not observable by the network, the collection of signers compute a combined public key along with a common signature which is verified by the new OP_CHECKSIG with the benefits of increased privacy and reduced costs. 

The latter involves threshold scenarios where only n-of-m signatures are necessary to authorize a transaction. The current implementation of OP_CHECKMULTISIG validates all of the public keys and associated signatures required by the threshold policy. Because the computation scales linearly with the number of participants, Schnorr propose a much more efficient scheme which replaces the list of signatures with a single combined one along with a subset of the required pubkeys. 

Until more evaluation of the delinearization scheme securing signers from malicious actors is performed, further applications of Schnorr signatures may be premature but the implementation of the features above can hopefully pave the way for a better understanding of the scheme in production. Contingent on additional peer-review, a BIP for the implementation of Schnorr Signatures could be proposed by the end of the year. 

<p align="center">
  <img src="http://i.imgur.com/TCR4Tf3.png">
</p>

