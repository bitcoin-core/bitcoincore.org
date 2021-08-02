---
# This file is licensed under the MIT License (MIT) available on
# http://opensource.org/licenses/MIT.

layout: post
lang: en
columns: 1
id: details-bait-and-switch
title: Details about the bait-and-switch attack
breadcrumbs:
  - features
  - verification
---
> Alice installs Example Wallet, whose open source code has been
> audited.  The next day, the authors of Example Wallet push new code to
> Alice's device and steal all her bitcoins.

- **Wallet service**{:.fgred} users are vulnerable because they can only
  spend their bitcoins when they use the service's approved software.

- **Light wallet**{:.fgyellow} users are vulnerable with most software
  because auditors can't easily verify that the software you run (the
  executable) is the same as the program source code - this is called a
  deterministic build. However, some light wallets are moving to
  deterministic builds.

- **Bitcoin Core**{:.fggreen} is built deterministically.  Cryptographic
  signatures from build auditors---many of whom are well known to the
  community---are [released publicly][gitian sigs].

<div class="callout" markdown="block">
Bitcoin.org's [Choose Your Wallet][] page tells you whether or not
wallet builds are audited in the *Transparency* score for each wallet.
</div>

### Real Example

In April 2013, the OzCoin mining pool was hacked. The thief stole 923
bitcoins (worth $135,000 USD), but online wallet StrongCoin modified
their wallet code to 'steal back' 569 of those bitcoins ($83,000)
from one of their users who was suspected of the theft.

Although this attack was done with good intentions, it illustrated
that the operators of StrongCoin could steal bitcoins from their users
at any time even though the users supposedly controlled their own
private keys.

**Learn More:** [OzCoin Hacked, Stolen Funds Seized and Returned by StrongCoin](https://bitcoinmagazine.com/4273/ozcoin-hacked-stolen-funds-seized-and-returned-by-strongcoin/)

[gitian sigs]: https://github.com/bitcoin-core/gitian.sigs
[choose your wallet]: https://bitcoin.org/en/choose-your-wallet
