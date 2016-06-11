---
type: posts
layout: post
lang: en
name: version-bits-faq-miners
id: en-version-bits-faq-miners
title: Version bits FAQ for miners
permalink: /en/2016/06/08/version-bits-miners-faq/
categories: [FAQ, mining]
tags: [soft fork, soft forks, bip9, version bits, mining]
version: 1
excerpt: The "version bits" BIP9 system is a way to introduce backward compatible rule changes to the Bitcoin consensus rules, known as a soft fork.
---
{% include _toc.html %}

## What is _version bits_ BIP9?

The _version bits_ [BIP9][] system is a way to introduce backward compatible rule changes to the Bitcoin consensus rules, known as a soft fork.

_version bits_ allows miners to signal that they can validate the soft fork rules. It also allows for up to 29 soft forks to be proposed concurrently.

## How is _version bits_ activated?

_version bits_ does not require activation, it's simply a way for miners to signal readiness for a soft fork by setting bits in the block header nVersion field.

## What are soft fork timeouts?

Soft forks have a start time and an _timeout_ during which the proposal is active. The soft fork can only be activated between _start time_ and _timeout_. If the soft fork does not get activated by the _timeout_, the soft fork proposal fails and will not activate even if signalled.

## What is the activation workflow?

Under _version bits_, a soft fork proposal goes through a workflow: 

- `[DEFINED]` -> `[STARTED]` -> `[LOCKED_IN]` -> `[ACTIVE]`

or

- `[DEFINED]` -> `[STARTED]` -> `[FAILED]`

![version bits state diagram](https://raw.githubusercontent.com/bitcoin/bips/master/bip-0009/states.png)

The Bitcoin network retargets mining difficulty every 2016 blocks; at this time _version bits_ will look at the window of the previous 2016 blocks to see how many blocks signal for a given soft fork. If 95% of the blocks signal readiness for the soft fork, the state changes from `[STARTED]` to `[LOCKED_IN]`.

After `[LOCKED_IN]` the rules will activate after one more difficulty retarget, i.e. a further 2016 blocks. Nodes software will warn that an upgrade is pending.

## What is the version bit?

When no soft forks are being signalled, miners should set the block version field to `0x20000000`.

## When should miners set bits?

To signal readiness for soft fork(s), miners should set the relevant version bit(s) together with `0x20000000`. This should only be done after the soft fork's _start time_. 

The bits should be unset if either the soft fork activates, or reached `[FAILED]` state.

For example:

"alice" soft fork uses bit 0, i.e. `0x1` + `0x20000000`

|0|0|1|0|0| ... |0|0|0|0|0|0|0|0|0|1|

"bob" soft fork uses bit, 1, i.e. `0x2` + `0x20000000`

|0|0|1|0|0| ... |0|0|0|0|0|0|0|0|1|0|

To signal both soft forks at once, use `0x20000003` (i.e. `0x1` + `0x2` + `0x20000000`*)

|0|0|1|0|0| ... |0|0|0|0|0|0|0|0|1|1|

* note if one is activated before the other, you must unset the relevant bit and continue signalling the other. IF one fails to activate and the timeout expires, you should also unset the bit.

## How does it differ to an ISM soft fork?

IsSuperMajority() or ISM for short, is a legacy soft fork trigger that activates new rules once 950 out of 1000 blocks are mined which signal the new block version.

1. An IsSuperMajority() soft fork will orphan all blocks with previous version after activation. For example, if the current version is 4, and the next soft fork introduces version 5 blocks, then after activation is reached (950/1000 blocks), nodes will reject all version 4 blocks.

2. Once a _version bits_ soft fork reaches activation, nodes will simply begin enforcing the new rules, and will NOT orphan a non-signalling block _unless_ it violates the new rules.

3. ISM() looks at the previous 1000 blocks on a rolling basis; _version bits_ looks at the previous 2016 block once each time the mining difficulty retargets.

4. ISM() soft forks do not expire. _version bits_ soft forks can only activate between the _start time_ and the _timeout_.

## Do miners have to upgrade?

No. A [BIP9][] soft fork will not activate unless 95% of the miners signal readiness. If a soft fork reaches `[LOCKED_IN]` state, where the vast majority of the miners are ready for the change, the remaining miners should upgrade _before_ the next difficulty retarget (about 2 weeks).

Non-upgraded miners risk producing invalid blocks which would be orphaned if they are not able to validate the newly activated rules.

## Who assigns version bits to different upgrade proposals?

Soft forks are proposed through the [BIPs process][BIP1]. Active [BIP9][] soft fork proposals are listed on the [assignments page](https://github.com/bitcoin/bips/blob/master/bip-0009.mediawiki#deployments)

## Further reading

- <http://rusty.ozlabs.org/?p=576>
- <https://github.com/bitcoin/bips/blob/master/bip-0009.mediawiki>
- <https://bitcointalk.org/index.php?topic=1067693.msg11446462#msg11446462>

{% include _references.md %}

