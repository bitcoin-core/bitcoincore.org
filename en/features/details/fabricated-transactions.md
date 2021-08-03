---
# This file is licensed under the MIT License (MIT) available on
# http://opensource.org/licenses/MIT.

layout: post
lang: en
columns: 1
id: details-fabricated-transactions
title: Details about fabricated transactions
breadcrumbs:
  - features
  - verification
---
> Mallory creates a transaction giving Alice 1,000 bitcoins, so Alice
> gives Mallory some cash. Later Alice discovers the transaction Mallory
> created was fake.

- **Wallet service**{:.fgred} users depend on the information reported
  by the service, so they can easily be fooled into accepting fabricated
  transactions.

- **Light wallet**{:.fgred} users depend on full nodes and
  miners to verify transactions for them. It costs nothing for
  dishonest full nodes to send unconfirmed fabricated transactions to a
  light wallet. Getting one or more confirmations on those fabricated
  transactions is also possible with help from a dishonest miner.

- **Bitcoin Core**{:.fggreen} users don't have to worry about fabricated
  transactions because Bitcoin Core verifies every transaction before
  displaying it.

<div class="callout" markdown="block">
Currently the best defense against fabricated transactions, besides
using Bitcoin Core, is to wait for as many confirmations as possible.
</div>

### Real Example

On 4 August 2015, web wallet BlockChain.info began indicating that a
transaction had spent the earliest mined 250 bitcoins, coins that some
people believed were owned by Bitcoin creator Satoshi Nakamoto.

It was soon discovered that the transaction was invalid. BlockChain.info
was not verifying transactions with Bitcoin Core and that transaction
had been [created by a security researcher](https://www.reddit.com/r/Bitcoin/comments/3fv42j/blockchaininfo_spoofed_transactions_problem_aug_4/)

**Learn more:** [BitcoinJ documentation about pending transaction
safety](https://bitcoinj.org/security-model#pending-transactions)

{% include references.md %}
