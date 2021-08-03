---
# This file is licensed under the MIT License (MIT) available on
# http://opensource.org/licenses/MIT.

layout: post
lang: en
columns: 1
id: details-chain-rewrites
title: Details about chain rewrites
breadcrumbs:
  - features
  - verification
---
> Mallory gives Alice 1,000 bitcoins. When Alice's wallet says the
> transaction is confirmed, Alice gives Mallory some cash. Later Alice
> discovers that Mallory has managed to steal back the bitcoins.

This attack applies to **all Bitcoin wallets.**{:.fgred}

The attack works because powerful miners have the ability to rewrite the
block chain and replace their own transactions, allowing them to take
back previous payments.

The cost of this attack depends on the percentage of total network hash
rate the attacking miner controls. The more centralized mining becomes,
the less expensive the attack is for a powerful miner.

Although all programs---including Bitcoin Core---are
vulnerable to chain rewrites, Bitcoin provides a defense mechanism: the
more confirmations your transactions have, the safer you are. *There is
no known decentralized defense better than that.*

### Real Example

In September 2013, someone used centralized mining pool GHash.io to
steal an estimated 1,000 bitcoins (worth $124,000 USD at the time) from the gambling
site BetCoin.

The attacker would spend bitcoins to make a bet.  If he won, he would
confirm the transaction.  If he lost, he would create a transaction
returning the bitcoins to himself and confirm that, invalidating the
transaction that lost the bet.

By doing so, he gained bitcoins from his winning bets without losing
bitcoins on his losing bets.

Although this attack was performed on unconfirmed transactions, the
attacker had enough hash rate (about 30%) to have profited from
attacking transactions with one, two, or even more confirmations.

**Learn more:** [GHash.IO and double-spending against BetCoin
Dice](https://bitcointalk.org/index.php?topic=327767.0)

{% include references.md %}
