---
title: Disclosure of CVE-2024-35202
name: blog-disclose-blocktxn-crash
id: en-blog-disclose-blocktxn-crash
lang: en
type: advisory
layout: post

## If this is a new post, reset this counter to 1.
version: 1

## Only true if release announcement or security annoucement. English posts only
announcement: 1

excerpt: >
  An attacker could remotely crash a Bitcoin Core node by triggering an
  assertion in the blocktxn message handling logic.
---

Before Bitcoin Core v25.0, an attacker could remotely crash Bitcoin Core
nodes by triggering an assertion in the blocktxn message handling logic.

This issue is considered **High** severity.

## Details

When receiving a block announcement via a cmpctblock message, Bitcoin Core
attempts to reconstruct the announced block using the transactions in its own
mempool as well as other available transactions. If reconstruction fails due to
missing transactions it will request them from the announcing peer via a
getblocktxn message. In response a blocktxn message is expected, which should
contain the requested transactions.

The compact block protocol employs shortened transaction identifiers to reduce
bandwidth. These short-ids are 6 byte in size, resulting in a small chance for
collisions (i.e. transaction A has the same short-id as transaction B) upon
block reconstruction. Collisions will be detected as the merkle root computed
from the reconstructed set of transactions will not match the merkle root from
the block announcement. Peers should not be punished for collisions as they may
happen spuriously, therefore they are handled by falling back to requesting the
full block.

Bitcoin Core will create an instance of <code>PartiallyDownloadedBlock</code>
whenever a new compact block is received. If missing transactions are
requested, the instance is persisted until the corresponding blocktxn message
is processed. Upon receiving the blocktxn message,
<code>PartiallyDownloadedBlock::FillBlock</code> is called, attempting to
reconstruct the full block. In the collision case described above, the full
block is requested but the <code>PartiallyDownloadedBlock</code> instance as
well as the other state related to the underlying block request is left
untouched. This leaves room for a second blocktxn message for the same block to
be processed and trigger <code>FillBlock</code> to be called again. This
violates the assumption (documented as an <code>assert</code> statement) that
<code>FillBlock</code> can only be called once and causes the node to crash.

An attacker does not need to get lucky by triggering a collision, as the
collision handling logic can easily be triggered by simply including
transactions in the blocktxn message that are not committed to in the block's
merkle root.

## Attribution

Credit goes to Niklas Gögge for discovering and disclosing the vulnerability,
as well as fixing the issue in https://github.com/bitcoin/bitcoin/pull/26898.

## Timeline

* 2022-10-05 - Niklas Gögge reports the issue to the Bitcoin Core security mailing list.
* 2023-01-24 - PR #26898 containing the fix is merged.
* 2023-05-25 - Bitcoin Core 25.0 is released with the fix.
* 2024-10-09 - Public disclosure.

{% include references.md %}
