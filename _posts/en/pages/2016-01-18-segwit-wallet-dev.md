---
version: 1
title: Segregated Witness Wallet Development Guide
name: segwit-wallet-dev
type: page
permalink: /en/segwit_wallet_dev
---

### Witness program

There are 2 ways to use segregated witness:

#### Native witness program

Native witness program is defined as a <code>scriptPubKey</code> with a single byte push (<code>OP_0</code>, <code>OP_1</code>, ...,  <code>OP_16</code>) followed by a push of 2 to 32 bytes.

#### Witness program nested in P2SH

Witness program nested in P2SH is a P2SH output with a <code>redeemScript</code> of a single byte push (<code>OP_0</code>, <code>OP_1</code>, ...,  <code>OP_16</code>) followed by a push of 2 to 32 bytes.

### Transaction Serialization

If a transaction is transmitted without any witness data, the traditional serialization format MUST be used.

If a transaction is transmitted with witness data, a new serialization format MUST be used:

<pre><code>[nVersion][marker][flag][txins][txouts][witness][nLockTime]</code></pre>

Format of <code>nVersion</code>, <code>txins</code>, <code>txouts</code>, and <code>nLockTime</code> are same as traditional serialization.

The <code>marker</code> MUST be <code>0x00</code>.

The <code>flag</code> MUST be a 1-byte non-zero value. Currently, <code>0x01</code> MUST be used.

The <code>witness</code> is a serialization of all witness data of the transaction. Each txin is associated with a witness field. A witness field starts with a var_int to indicate the number of stack items for the txin. It is followed by stack items, with each item starts with a var_int to indicate the length. Witness data is NOT script and is not restricted by the 520-byte push limit.

A non-witness program txin MUST be associated with an empty witness field, represented by a <code>0x00</code>. If all txins are not witness program, a transaction MUST be serialized with the traditional format (exception: coinbase transaction).

### Transaction ID

Each transaction will have 2 IDs.

Definition of <code>txid</code> remains unchanged: the double SHA256 of the traditional serialization.

A new <code>wtxid</code> is defined: the double SHA256 of the new serialization with witness data. If a transaction is transmitted without any witness data, its <code>wtxid</code> is same as the <code>txid</code>.

The txid remains the primary identifier of a transaction. Particularly, it is used in the txin when referring to a previous output.

### Standard script types

#### Pay-to-Public-Key-Hash (P2PKH)
P2PKH is the most common <code>scriptPubKey</code> template defined by Satoshi Nakamoto, allows simple payment to a single public key. The format is:

<pre><code>scriptPubKey (25 bytes): OP_DUP OP_HASH160 < 20-byte-pubkey-hash > OP_EQUALVERIFY OP_CHECKSIG</code></pre>

To spend a P2PKH output, the <code>scriptSig</code> is

<pre><code>scriptSig: < sig > < pubkey ></code></pre>

where the <code>RIPEMD160(SHA256(pubkey))</code> is equal to the <code>20-byte-pubkey-hash</code> in scriptPubKey.

#### Pay-to-Script-Hash (P2SH)
P2SH is defined in BIP16. It allows payment to arbitrarily complex scripts with a fixed length <code>scriptPubKey</code>. The format is:

<pre><code>scriptPubKey (23 bytes): OP_HASH160 <20-byte-script-hash> OP_EQUAL</code></pre>

To spend a P2SH output, the <code>scriptSig</code> is

<pre><code>scriptSig: <...> <...> <...> < redeemScript ></code></pre>

where the <code>RIPEMD160(SHA256(redeemScript))</code> is equal to the <code>20-byte-script-hash</code> in <code>scriptPubKey</code>. The <code>redeemScript</code> is deserialized and evaluated with the remaining data in the <code>scriptSig</code>.

#### Pay-to-Witness-Public-Key-Hash (P2WPKH)
P2WPKH is a new defined in BIP141. Similar to P2PKH, it allows simple payment to a single public key. The format is:

<pre><code>scriptPubKey (22 bytes): OP_0 < 20-byte-pubkey-hash ></code></pre>

To spend a P2WPKH output, the <code>scriptSig</code> MUST be empty, and the <code>witness</code> is:
  
<pre><code>witness: < sig > < pubkey ></code></pre>

where the <code>RIPEMD160(SHA256(pubkey))</code> is equal to the <code>20-byte-pubkey-hash</code> in <code>scriptPubKey</code>.

#### P2WPKH in P2SH (P2SH-P2WPKH)
P2SH-P2WPKH is using a P2WPKH script as the <code>redeemScript</code> for P2SH. The <code>scriptPubKey</code> of P2SH-P2WPKH looks exactly the same as an ordinary P2SH:

<pre><code>scriptPubKey (23 bytes): OP_HASH160 < 20-byte-script-hash > OP_EQUAL</code></pre>

To spend a P2SH-P2WPKH output, the <code>scriptSig</code> MUST contain a push of the <code>redeemScript</code> and nothing else and the <code>witness</code> is same as P2WPKH:

<pre><code>scriptSig (23 bytes): < OP_0 < 20-byte-pubkey-hash > >
witness: < sig > < pubkey ></code></pre>
  
where the <code>RIPEMD160(SHA256(pubkey))</code> is equal to the <code>20-byte-pubkey-hash</code>, and <code>RIPEMD160(0x0014{20-byte-pubkey-hash})</code> is equal to the <code>20-byte-script-hash</code>.
  
#### Pay-to-Witness-Script-Hash (P2WSH)
P2WSH is another new standard script defined in BIP141. Similar to P2SH, it allows payment to arbitrarily complex scripts. The format is:

<pre><code>scriptPubKey (34 bytes): OP_0 < 32-byte-script-hash ></code></pre>

To spend a P2WSH output, the <code>scriptSig</code> MUST be empty, and the witness is:

<pre><code>witness: <...> <...> <...> < witnessScript ></code></pre>

where the <code>RIPEMD160(SHA256(witnessScript))</code> is equal to the <code>32-byte-script-hash</code> in scriptPubKey. The <code>witnessScript</code> is deserialized and evaluated with the remaining data in the <code>witness</code>.


#### P2WSH in P2SH (P2SH-P2WSH)
P2SH-P2WPKH is using a P2WSH script as the <code>redeemScript</code> for P2SH. The <code>scriptPubKey</code> of P2SH-P2WPKH looks exactly the same as an ordinary P2SH:

<pre><code>scriptPubKey (23 bytes): OP_HASH160 <20-byte-script-hash> OP_EQUAL</code></pre>

To spend a P2SH-P2WPKH output, the <code>scriptSig</code> MUST contain a push of the <code>redeemScript</code> and nothing else and the <code>witness</code> is same as P2WSH:

<pre><code>scriptSig (35 bytes): < OP_0 < 32-byte-script-hash > > \
witness: <...> <...> <...> < witnessScript ></code></pre>

where the <code>SHA256(witnessScript)</code> is equal to the <code>32-byte-script-hash</code>, and <code>RIPEMD160(0x0020{32-byte-pubkey-hash})</code> is equal to the <code>20-byte-script-hash</code>.

### New signing algorithm
To spend a witness program output, a new signing algorithm MUST be used when producing the ECDSA signature. A step-by-step example could be found in BIP143.
