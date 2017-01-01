---
version: 3
title: Segregated Witness Wallet Development Guide
name: segwit-wallet-dev
type: pages
layout: page
lang: en
permalink: /en/segwit_wallet_dev/
version: 1
---
{% include _toc.html %}
{% include _references.md %}

Most contents of this document could be found in the BIPs related to segregated witness, including [BIP141][], [BIP143][], [BIP144][], and [BIP145][]. Please consider this as the first point of reference to other related documents, and as a checklist for what should and should not be done.

### Basic segregated witness support

A wallet MUST implement all the features in this section, in order to be considered as segwit-compatible at a basic level:

#### Sending to P2SH

* A segwit-compatible wallet MUST support pay-to-script-hash ([BIP16][]) and its address format ([BIP13][]).
* For making payments, the wallet must be able to correctly transform a given P2SH address to a <code>scriptPubKey</code>, and create a transaction.
* For receiving payments, the wallet must be able to create a P2SH address based on a P2WPKH script (defined hereinafter), and be able to recognize payment to such addresses.
* This is a mandatory requirement, even if the wallet accepts only single-signature payments.

#### Creation of P2SH-P2WPKH Address

* A P2SH-P2WPKH address is comparable to Bitcoin's original single-signature P2PKH address (address with prefix 1).
* Like any other P2SH address, P2SH-P2WPKH address has prefix 3.
* Until a P2SH-P2WPKH UTXO is spent and the <code>redeemScript</code> is exposed, a P2SH-P2WPKH address is indistinguishable from a non-segwit P2SH address (such as a non-segwit multi-signature address)
* P2SH-P2WPKH addresses should be used when only 1 public key is used to receive payment (like P2PKH)
* P2SH-P2WPKH uses the same public key format as P2PKH, with a very important exception: the public key used in P2SH-P2WPKH MUST be compressed, i.e. 33 bytes in size, and starting with a <code>0x02</code> or <code>0x03</code>. Using any other format such as uncompressed public key may lead to irrevocable fund loss.
* To create a P2SH-P2WPKH address:
    1. Calculate the RIPEMD160 of the SHA256 of a public key (<code>keyhash</code>). Despite the <code>keyhash</code> formula is same as P2PKH, reuse of <code>keyhash</code> should be avoided for better privacy and prevention of accidental use of uncompressed key
    2. The P2SH <code>redeemScript</code> is always 22 bytes. It starts with a <code>OP_0</code>, followed by a canonical push of the <code>keyhash</code> (i.e. <code>0x0014{20-byte keyhash}</code>)
    3. Same as any other P2SH, the <code>scriptPubKey</code> is <code>OP_HASH160 hash160(redeemScript) OP_EQUAL</code>, and the address is the corresponding P2SH address with prefix 3.

#### Transaction Serialization

* A segwit-compatible wallet MUST support the original transaction format, as <code>nVersion|txins|txouts|nLockTime</code>.
* A segwit-compatible wallet MUST also support the new serialization format, as <code>nVersion|marker|flag|txins|txouts|witness|nLockTime</code>
    * Format of <code>nVersion</code>, <code>txins</code>, <code>txouts</code>, and <code>nLockTime</code> are same as the original format
    * The <code>marker</code> MUST be <code>0x00</code>
    * The <code>flag</code> MUST be <code>0x01</code>
    * The <code>witness</code> is a serialization of all witness data of the transaction.
        * Each txin is associated with a witness field. As a result, there is no indication of number of witness fields, as it is implied by the number of <code>txins</code>
        * Each witness field starts with a <code>compactSize</code> [integer](https://bitcoin.org/en/developer-reference#compactsize-unsigned-integers) to indicate the number of stack items for the corresponding <code>txin</code>. It is then followed by witness stack item(s) for the corresponding <code>txin</code>, if any. 
        * Each witness stack item starts with a <code>compactSize</code> integer to indicate the number of bytes of the item.
        * If a <code>txin</code> is not associated with any witness data, its corresponding witness field is an exact <code>0x00</code>, indicating that the number of witness stack items is zero.
* If all <code>txins</code> in a transaction are not associated with any witness data, the transaction MUST be serialized in the original transaction format, without <code>marker</code>, <code>flag</code>, and <code>witness</code>. For example, if none of the <code>txins</code> are coming from segwit UTXO, it MUST be serialized in the original transaction format. (exception: coinbase transaction)
* Examples of the transaction serialization can be found under the example section of BIP143. Wallet developers may use the examples to test if their implementations correctly parse the new serialization format.

#### Transaction ID

* Under segwit, each transaction will have 2 IDs.
* The definition of <code>txid</code> remains unchanged: the double SHA256 of the original serialization format.
* A new <code>wtxid</code> is defined, which is the double SHA256 of the new serialization format with witness data.
* If a transaction does not have any witness data, its <code>wtxid</code> is same as the <code>txid</code>.
* The <code>txid</code> remains the primary identifier of a transaction:
    * It MUST be used in the <code>txin</code> when referring to a previous output.
    * If a wallet or service is currently using <code>txid</code> to identify transactions, it is expected to use the same after upgrade.

#### Signature Generation and Verification for P2SH-P2WPKH

* For spending of non-segwit UTXO, the signature generation algorithm is unchanged.
* For spending of P2SH-P2WPKH:
    * The <code>scriptSig</code> MUST ONLY contain a push of the <code>redeemScript</code>
    * The corresponding witness field MUST contain exactly 2 items, a signature followed by the public key
    * There is a new signature generation algorithm described in [BIP143][] for segwit scripts. Developers should follow the instructions carefully, and make use of the P2SH-P2WPKH example in [BIP143](https://github.com/bitcoin/bips/blob/master/bip-0143.mediawiki#P2SHP2WPKH) to make sure they are able to reproduce the <code>sighash</code>.
    * The [BIP143][] signature generating algorithm covers the value of the input being spent, which simplifies the design of air-gapped light-weight wallets and hardware wallets.
    * Please note that for a P2SH-P2WPKH, the <code>scriptCode</code> is always 26 bytes including the leading size byte, as <code>0x1976a914{20-byte keyhash}88ac</code>, NOT the <code>redeemScript</code> nor <code>scriptPubKey</code>
    * [Example](http://n.bitcoin.ninja/checktx?txid=8139979112e894a14f8370438a471d23984061ff83a9eba0bc7a34433327ec21)
    
#### Network Services (optional)

* Network services are needed if the wallet would send and receive transactions through the peer-to-peer network
* Segwit-capable nodes will advertise that they can provide witnesses through the service bit: <code>NODE_WITNESS = (1 << 3)</code>
* Transactions without any witness data (and therefore serialized in original format) may be sent to nodes with or without <code>NODE_WITNESS</code> support
* Transactions which spend segwit UTXOs (and therefore is serialized in the new format) MUST ONLY be sent to nodes with <code>NODE_WITNESS</code> support
* Transactions which spend segwit UTXOs but with witness data stripped (and therefore serialized in original format) may be sent to nodes without <code>NODE_WITNESS</code> support. Such transactions, however, are invalid after activation of segwit and would not be accepted in a block.
* Details of the network services could be found in [BIP144][].

#### User Privacy

* In the early days just after segwit is activated, there may be limited number of segwit transactions in the network.
* Using segwit transaction when it is uncommon may make Bitcoin tracking easier.
* Using P2SH-P2WPKH as default change output may also have an impact on privacy.

#### Transaction Fee Estimation

* Instead of transaction size, a new metric is defined, called "virtual size" (<code>vsize</code>)
* <code>vsize</code> of a transaction equals to 3 times of the size with original serialization, plus the size with new serialization, divide the result by 4 and round up to the next integer. For example, if a transaction is 200 bytes with new serialization, and becomes 99 bytes with <code>marker</code>, <code>flag</code>, and <code>witness</code> removed, the <code>vsize</code> is (99 * 3 + 200) / 4 = 125 with round up.
* <code>vsize</code> of a non-segwit transaction is simply its size
* Transaction fee should be estimated by comparing the <code>vsize</code> with other transactions, not the size.
* Developers should be careful not to make an off-by-4-times mistake in fee estimation.


#### Upgrade Safety

* End users MUST NOT be allowed to generate any P2SH-P2WPKH or other segwit addresses before segwit is fully activated on the network. Before activation, using P2SH-P2WPKH or other segwit addresses may lead to permanent fund loss
* Similarly, change MUST NOT be sent to a segwit output before activation
* Activation of segwit is defined by [BIP9][]. After 15 Nov 2016 and before 15 Nov 2017 UTC, if in a full retarget cycle at least 1916 out of 2016 blocks is signaling readiness, segwit will be activated in the retarget cycle after the next one
* If a wallet does not have the ability to follow the [BIP9][] signal, the upgraded version should not be released to end users until it is activated
* If there were any concerns that some miners may not correctly enforce the new rules, release of the upgraded wallet may be delayed until evidence has shown that vast majority (if not all) miners are following the new rules. Violations would be very obvious, shown as invalid orphaned blocks.

#### Backward Compatibility

* Sending and receiving P2PKH payment (address with prefix 1) should remain to be supported.


### Complex script support

If a wallet supports script types other than just single signature, such as multi-signature, it has to fulfill more than the basic requirements:

#### Creation of P2SH-P2WSH Address

* A P2SH-P2WSH address is comparable to Bitcoin's original P2SH address, which allows representation of arbritarily complex scripts with a fixed size address.
* Like any other P2SH and P2SH-P2WPKH address, P2SH-P2WSH address has prefix 3. They are indistinguishable until the UTXO is spent
* To create a P2SH-P2WSH address:
    1. Define a script, called (<code>witnessScript</code>)
    2. Calculate the SHA256 of the <code>witnessScript</code> (<code>scripthash</code>). Please pay attention that a single SHA256 is used, not double SHA256 nor RIPEMD160(SHA256)
    3. The P2SH <code>redeemScript</code> is always 34 bytes. It starts with a <code>OP_0</code>, followed by a canonical push of the <code>scripthash</code> (i.e. <code>0x0020{32-byte scripthash}</code>)
    4. Same as any other P2SH, the <code>scripPubKey</code> is <code>OP_HASH160 hash160(redeemScript) OP_EQUAL</code>, and the address is the corresponding P2SH address with prefix 3.
* Restrictions on the script
    * The script evaluation must not fail, and MUST leave one and only one TRUE stack item after evaluation. Otherwise, the evaluation is failed.
    * Any public key inside P2SH-P2WSH scripts MUST be compressed key, or fund may be lost permenantly.
    * If OP_IF or OP_NOTIF is used, it argument MUST be either an empty vector (for false) or <code>0x01</code> (for true). Use of other value may lead to permenant fund loss. ([BIP draft](https://lists.linuxfoundation.org/pipermail/bitcoin-dev/2016-August/013014.html))
    * If an OP_CHECKSIG or OP_CHECKMULTISIG is returning a fail, all signature(s) must be empty vector(s). Otherwise, fund may be lost permenantly. ([BIP146][])
    * There is a default policy limit for the <code>witnessScript</code> at 3600 bytes. Except the <code>witnessScript</code>, there could be at most 100 witness stack items, with at most 80 bytes each. Transactions excessing these limits may not be relayed nor included in a block
    * Many of the original scripts consensus limitations, such as 10000 bytes script size, 201 <code>nOpCount</code>, are still applied to P2SH-P2WSH
    * The 520 bytes script size limit for P2SH is not applicable to P2SH-P2WSH. It is replaced by the 3600 bytes policy limit and 10000 bytes consensus limit.
    
#### Signature Generation and Verification for P2SH-P2WSH

* For spending of P2SH-P2WSH:
    * The <code>scriptSig</code> MUST ONLY contain a push of the <code>redeemScript</code>
    * The last witness item of the corresponding witness field MUST be the <code>witnessScript</code>
    * The new [BIP143][] signature generation algorithm is applied:
        * Without using any OP_CODESEPARATOR, the <code>scriptCode</code> is <code>witnessScript</code> preceeded by a <code>compactSize</code> integer for the size of <code>witnessScript</code>. For example, if the script is OP_1 (<code>0x51</code>), the <code>scriptCode</code> being serialized is (<code>0x0151</code>)
        * For any unusal scripts containing OP_CODESEPARATOR, please refer to [BIP143][] for the exact semantics
    * Any witness stack items before the <code>witnessScript</code> are used as the input stack for script evaluation. The input stack is not interpreted as script. For example, there is no need to use a <code>0x4c</code> (OP_PUSHDATA1) to "push" a big item.
    * To verify the correctness of signature generation and stack serialization, please always test against the examples in [BIP143][]
    * [Example](http://n.bitcoin.ninja/checktx?txid=954f43dbb30ad8024981c07d1f5eb6c9fd461e2cf1760dd1283f052af746fc88)

### Advanced designs

The following functions are not required for initial segwit support.

#### Native Pay-to-Witness-Public-Key-Hash (P2WPKH)

* Native P2WPKH is a <code>scripPubKey</code> of 22 bytes. It starts with a <code>OP_0</code>, followed by a canonical push of the <code>keyhash</code> (i.e. <code>0x0014{20-byte keyhash}</code>)
* Same as P2SH-P2WPKH, <code>keyhash</code> is RIPEMD160(SHA256) of a compressed public key.
* When spending a native P2WPKH, the <code>scriptSig</code> MUST be empty, and the witness stack format and signature generating rules are same as P2SH-P2WPKH (including the requirement of using compressed public key)
* [Example](http://n.bitcoin.ninja/checktx?txid=d869f854e1f8788bcff294cc83b280942a8c728de71eb709a2c29d10bfe21b7c)

#### Native Pay-to-Witness-Script-Hash (P2WSH)

* Native P2WSH is a <code>scripPubKey</code> of 34 bytes. It starts with a <code>OP_0</code>, followed by a canonical push of the <code>scripthash</code> (i.e. <code>0x0020{32-byte keyhash}</code>)
* Same as P2SH-P2WSH, <code>scripthash</code> is SHA256 of the <code>witnessScript</code>.
* When spending a native P2WSH, the <code>scriptSig</code> MUST be empty, and the witness stack format and signature generating rules are same as P2SH-P2WSH (including the requirement of using compressed public key)
* [Example](http://n.bitcoin.ninja/checktx?txid=78457666f82c28aa37b74b506745a7c7684dc7842a52a457b09f09446721e11c)

#### Why and How to Use Native P2WPKH and P2WSH?

* There is no address format for native P2WPKH and P2WSH. BIP142 is deferred and is likely to be implemented in a completely different way
* Comparing with the P2SH versions, the transaction <code>vsize</code> of native versions is smaller in most cases, and therefore less fee may be required
* Native P2WPKH and P2WSH may be used with raw <code>scripPubKey</code> protocols, such as the Payment Protocol (BIP70). However, it may affect the privacy of the payer and recipient (see below).
* Native P2WPKH and P2WSH may be used as default change address, but this may allow other people identifying the change easily (see below)
* It is expected that the use of native P2WPKH and P2WSH would be uncommon at the beginning, which may cause privacy concerns among the users.


### Scripts and Transactions Examples

* [Examples of different witness transaction types and transaction validity checking tool](http://n.bitcoin.ninja/checktx)
* [BIP141][]
* [BIP143][]
* [Script tests](https://github.com/bitcoin/bitcoin/blob/master/src/test/data/script_tests.json)
* [Valid transaction tests](https://github.com/bitcoin/bitcoin/blob/master/src/test/data/tx_valid.json)
* [Invalid transaction tests](https://github.com/bitcoin/bitcoin/blob/master/src/test/data/tx_invalid.json)
