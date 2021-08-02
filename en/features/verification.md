---
# This file is licensed under the MIT License (MIT) available on
# http://opensource.org/licenses/MIT.

layout: post
lang: en
columns: 1
id: bitcoin-core-verification
title: Verification - Bitcoin Core Features
breadcrumbs:
  - features
---
> Imagine a scientist reading about an experimental result and then
> repeating the experiment for herself. Doing so allows her to trust
> the result without having to trust the original scientists.

Bitcoin Core checks each block of transactions it receives to ensure
that everything in that block is fully valid---allowing it to trust the
block without trusting the miner who created it.

This prevents miners from tricking Bitcoin Core users into accepting
blocks that violate the 21 million bitcoin limit or other
important rules.

Users of other wallets don't get this level of security, so they can be
tricked into accepting fabricated transactions or alternative block chains.

Why take that risk if you don't have to? Bitcoin Core provides the best
possible security against dishonest miners, along with additional
security against other types of attacks (see below for details).

## How Verification Protects Your Bitcoins

Light wallets and wallet services (such as exchanges) put your bitcoins
at increased risk of being stolen. That risk may be acceptable for small
values of bitcoin on mobile wallets, but is it what you want for your
real wallet?

{:.center}
*Click the title of any attack below to learn more about it.*

<table>
<tr>
  <th>Attack</th>
  <th markdown="span">Wallet Service</th>
  <th markdown="span">Light Wallet</th>
  <th>Bitcoin Core</th>
</tr>

<tr>
  <td><a href="/{{page.lang}}/features/details/direct-theft/">Direct theft</a></td>
  <td class="bgred"></td>
  <td class="bggreen"></td>
  <td class="bggreen"></td>
</tr>

<tr>
  <td><a href="/{{page.lang}}/features/details/bait-and-switch/">Bait and switch</a></td>
  <td class="bgred"></td>
  <td class="bgyellow"></td>
  <td class="bggreen"></td>
</tr>

<tr>
<td><a href="/{{page.lang}}/features/details/fabricated-transactions/">Fabricated transactions</a></td>
  <td class="bgred"></td>
  <td class="bgred"></td>
  <td class="bggreen"></td>
</tr>

<tr>
<td><a href="/{{page.lang}}/features/details/transaction-withholding/">Transaction withholding</a></td>
  <td class="bgred"></td>
  <td class="bgred"></td>
  <td class="bggreen"></td>
</tr>

<tr>
<td><a href="/{{page.lang}}/features/details/chain-hijacking/">Chain hijacking</a></td>
  <td class="bgred"></td>
  <td class="bgred"></td>
  <td class="bggreen"></td>
</tr>

<tr>
<td><a href="/{{page.lang}}/features/details/chain-rewrites/">Chain rewrites</a></td>
  <td class="bgred"></td>
  <td class="bgred"></td>
  <td class="bgred"></td>
</tr>

</table>

Note that although all programs---including Bitcoin Core---are
vulnerable to chain rewrites, Bitcoin provides a defense mechanism: the
more confirmations your transactions have, the safer you are. *There is
no known decentralized defense better than that.*

## Help Protect Decentralization

{% comment %}<!-- 1231006505 is the time in block 0; 31558149 is the average number of seconds in a year -->{% endcomment %}
{% capture bitcoin_age %}{{ site.time | date: "%s" | minus: "1231006505" | divided_by: "31558149" }}{% endcapture %}

The bitcoin currency only works when people accept bitcoins in exchange
for other valuable things. That means it's the people accepting
bitcoins who give it value and who get to decide how Bitcoin should work.

When you accept bitcoins, you have the power to enforce Bitcoin's rules,
such as preventing confiscation of any person's bitcoins without access
to that person's private keys.

Unfortunately, many users outsource their enforcement power. This
leaves Bitcoin's decentralization in a weakened state where a handful of
miners can collude with a handful of service providers to change
Bitcoin's rules for all those non-verifying users who outsourced their power.

<table class="received_transactions center">
  <tr>
    <td class="center" markdown="span">*Users of wallet services*<br
      >**Trust service providers**{:.fgred}</td>

    <td class="center" markdown="span">*Users of P2P light wallets*<br
    >**Trust miners**{:.fgred}</td>
  </tr>

  <tr>
    <td class="center" markdown="span">*Users of client light wallets*<br
    > **Trust "free" servers**{:.fgred}</td>

    <td class="center" markdown="span">*Users of Bitcoin Core*<br
    >**Enforce the rules**{:.fggreen}</td>
  </tr>
</table>

Unlike other wallets, **Bitcoin Core *does* enforce the rules**---so
if the miners and service providers change the rules for their non-verifying
users, those users will be unable to pay full verification Bitcoin Core
users like you.

As long as there are many non-verifying users who want to be able to
pay Bitcoin Core users, miners and others know that they can't effectively
change Bitcoin's rules.

But what if not enough non-verifying users care about paying Bitcoin
Core users? Then it becomes easy for miners and service providers to
take control of Bitcoin, likely bringing to an end this {{bitcoin_age}}
year experiment in decentralized currency.

{:.center}
![History of Bitcoin](/assets/images/losing-decentralization.png)

If you think **Bitcoin should remain decentralized,** the best thing you
can do is [verify every payment you receive](#do-you-verify) using your own personal
full node such as Bitcoin Core.

We don't know how many full verification users and business are needed,
but it's possible that for each person or business who verifies their
own transactions, Bitcoin can remain decentralized even if there are ten
or a hundred other non-verifying users. If this is the case, your
small contribution can have a large impact towards keeping Bitcoin
decentralized.

## Do You Verify Your Transactions? {#do-you-verify}

Some people confuse just running a node with helping to [protect
Bitcoin's decentralization][features verification decentralization].

To [improve your security][features verification protection] and help
protect decentralization, you must use a wallet that fully verifies
received transactions. There are three ways to do that with Bitcoin
Core right now:

1. **Use the built-in wallet's graphical mode.** If you request payment
   using the following screen in Bitcoin Core, your received
   transactions will be fully verified.

    {:.center}
    ![Bitcoin Core request payment](/assets/images/2021-request-payment.png)

2. **Use Bitcoin Core as a trusted peer for certain light
   wallets and other programs.** Learn more on the user interface page.<!-- FIXME: add link when page available -->
   If you use a secure connection to your personal trusted node *every
   time* you use the wallet, your received transactions will be fully
   verified.  Many Lightning Network wallets also support connecting to
   a trusted node.

3. **Use the built-in wallet's CLI/API interface.** This is meant for
   power users, businesses, and programmers. The user interface page <!-- FIXME: add link --> 
   provides an overview, the [installation instructions][bitcoin core
   install guide] can help you get started, and the [RPC][rpc docs]
   documentation can help you find specific commands. If you're using
   the `getnewaddress` RPC to create receiving addresses, your received
   transactions will be fully verified.

{% include references.md %}
