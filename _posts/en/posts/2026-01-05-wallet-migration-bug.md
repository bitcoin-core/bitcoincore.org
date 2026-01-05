---
title: Wallet Migration Failure May Delete Unrelated Wallet Files In Bitcoin Core 30.0 and 30.1
name: wallet-deletion-migration-bug
id: en-wallet-deletion-migration-bug
lang: en
type: posts
layout: post
excerpt: A wallet migration bug present in Bitcoin Core 30.0 and 30.1 may delete all files in the wallet directory, potentially resulting in a loss of funds.
---

We have become aware of a wallet migration bug introduced in Bitcoin Core 30.0 and 30.1. Under rare circumstances, when the migration of a wallet.dat file fails, all files in the wallet directory may be deleted in the process, potentially resulting in a loss of funds. A fix is forthcoming and will be released as 30.2, but out of an abundance of caution we have removed the binaries for affected releases from bitcoincore.org.

At this time, we ask users to not attempt wallet migrations using the GUI or RPC until v30.2 is released. All other users, including existing wallet users, are unaffected and can keep using existing installations.

Specifically, it requires the presence of a default (unnamed) wallet.dat file, which has not been created by default since 0.21 (released 5 years ago), that fails to be migrated or loaded. One condition that may trigger this is when pruning is enabled, and the wallet was unloaded while pruning happened.

{% include references.md %}
