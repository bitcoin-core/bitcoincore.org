---
# This file is licensed under the MIT License (MIT) available on
# http://opensource.org/licenses/MIT.

layout: post
lang: en
columns: 1
id: details-transaction-withholding
title: Details about transaction withholding
breadcrumbs:
  - features
  - verification
---
> Mallory offers Alice $1,000 USD in exchange for some bitcoins.  Alice
> sends the bitcoins but the transaction never seems to confirm and so
> Alice is unable to insist that Mallory pay the cash.  It turns out the
> transaction did confirm, so Alice gave away her bitcoins for nothing.

- **Wallet service**{:.fgred} users only see the transactions the service
  choose to show them.

- **Light wallets**{:.fgred} users only see the transactions their full
  node peers choose to send them, even if those transactions were
  included in a block the light wallet knows about.

- **Bitcoin Core**{:.fggreen} users see all transactions included in
  received blocks. If Bitcoin Core hasn't received a block for too long,
  it displays a catching-up progress bar in the graphical user
  interface or a warning message in the CLI/API user interface.

<div class="callout" markdown="block">
Unless you use Bitcoin Core, you can never be sure that your bitcoin balance
is correct according to the block chain.
</div>

### Real Example

In March 2015, spy nodes run by the company Chainalysis accidentally
prevented some users of the light wallet BRD from connecting to
honest nodes. Since the spy nodes didn't relay transactions, BRD
users stopped receiving notification of new transactions.

**Learn more:** [Chainalysis CEO Denies 'Sybil Attack' on Bitcoin's Network](http://www.coindesk.com/chainalysis-ceo-denies-launching-sybil-attack-on-bitcoin-network/)

{% include references.md %}
