---
title: Undo file overflow under extreme conditions
name: blog-disclose-undo-overflow
id: en-blog-disclose-undo-overflow
lang: en
type: advisory
layout: post
redirect_from:
  - /en/2026/10/27/disclose-undo-overflow

## If this is a new post, reset this counter to 1.
version: 1

## Only true if release announcement or security annoucement. English posts only
announcement: 1

excerpt: >
  After filling thousands of blocks in preparation, a miner able to trigger a precisely timed short reorg following a specific sequence of blocks could crash some nodes by corrupting their undo data. A fix for 64-bit systems was released in Bitcoin Core versions 29.5, 30.4, 31.2 and 32.1; 32-bit systems and nodes storing blocks on FAT32 remain affected.
---

Disclosure of the details of a bug that lets a miner corrupt the undo data of other nodes, causing
them to shut down if a reorg of precisely the right blocks follows, or on Windows and 32-bit systems
as soon as the affected blocks are connected. It is impractical to exploit. A fix for 64-bit systems
was released in Bitcoin Core versions 29.5, 30.4, 31.2 and 32.1. 32-bit systems and nodes storing blocks
on FAT32 remain affected.

This issue is considered **Low** severity and **easily discoverable**.

## Details

To disconnect a block during a reorg, Bitcoin Core reads the outputs the block spent from its undo
data. Undo data is stored in `rev*.dat` files, each paired with the `blk*.dat` file holding the
corresponding blocks. Block files are capped at 128 MiB, but undo files have no size limit, and
positions within them were stored as 32-bit integers.

A miner can produce valid blocks whose undo data is much larger than the blocks themselves, enough to
push a single undo file past 4 GiB. Past that point the 32-bit position wraps: newer undo data can
overwrite older undo data, and the file is truncated to the wrapped length when its block file is
finalized. If a reorg later disconnects an affected block, the node cannot read its undo data and
shuts down. Restarting does not help, and recovery requires a `-reindex` with a fixed version.

On Windows and 32-bit systems, file positioning calls are limited to 2 GiB, so the node instead fails
to write undo data once an undo file passes 2 GiB, and shuts down on every restart. This needs about
half as much undo data and no reorg.

Exploiting this on 64-bit systems is impractical. The transactions involved are nonstandard, and an
attacker would have to mine several thousand blocks of them, roughly a month of the entire network's
block production. For a reorg short enough to be practical, about 18 more of their blocks must then
land at the very end of one block file on the targeted nodes, and the attacker must win a 2-block
reorg replacing exactly the last of them. Block file layouts differ from node to node, so no
placement works for every node, and if the reorg fails the preparation is wasted. In theory, with
even more precise control over each node's files, an attacker could instead make a node restore the
wrong coins during a reorg, corrupting its UTXO set.

The issue was fixed by storing block and undo file positions as 64-bit integers, and by using 64-bit
file positioning calls on Windows. Positions are serialized as variable-length integers, so the
on-disk format is unchanged.

On 32-bit systems, these releases still use positioning calls limited to 2 GiB, so those systems
remain affected by the failure described above. Nodes whose blocks directory is on a FAT32 file
system also remain affected: FAT32 limits files to 4 GiB, so writing undo data fails once an undo
file reaches that size, and the node shuts down on every restart. This needs as much undo data as on
other 64-bit systems, but no reorg.

## Attribution

This issue was discovered through an automated LLM scan, as part of Elias Rohrer's [Project
Loupe](https://github.com/project-loupe/loupe) at Spiral. The Loupe finding was validated and
surfaced to the Bitcoin Core team by Gregory Sanders.

## Timeline

- 2026-07-08 - Project Loupe reports the issue
- 2026-08-21 - L0rinc independently reports the issue
- 2026-10-23 - Andrew Dorman independently reports the issue
- 2026-10-26 - Antoine Poinsot opens PR [#36669](https://github.com/bitcoin/bitcoin/pull/36669) with fix
- 2026-10-26 - PR #36669 is merged into master
- 2026-10-26 - The fix is backported to the 32.x branch in PR
  [#36658](https://github.com/bitcoin/bitcoin/pull/36658)
- 2026-10-26 - The fix is backported to the 31.x branch in PR
  [#36661](https://github.com/bitcoin/bitcoin/pull/36661)
- 2026-10-26 - The fix is backported to the 30.x branch in PR
  [#36664](https://github.com/bitcoin/bitcoin/pull/36664)
- 2026-10-26 - The fix is backported to the 29.x branch in PR
  [#36666](https://github.com/bitcoin/bitcoin/pull/36666)
- 2026-10-27 - Bitcoin Core 29.5, 30.4, 31.2 and 32.1 are released with a fix
- 2026-10-27 - Public disclosure

{% include references.md %}
