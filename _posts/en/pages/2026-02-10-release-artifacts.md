---
title: Release Artifacts
name: release-artifacts
type: pages
layout: page
lang: en
permalink: /en/release-artifacts/
version: 1
---

Bitcoin Core release files are hosted at [https://bitcoincore.org/bin/](https://bitcoincore.org/bin/) and distributed via torrent.

## Build Suffixes

- **Standard Binaries** (`.tar.gz`, `.zip`, `-setup.exe`): Production-ready, signed binaries for general use.
- **`-debug`**: Includes symbols for backtrace generation. Larger file size and reduced performance.
- **`-unsigned`**: Raw Guix build outputs. Used to audit reproducibility against signed releases.
- **`-codesigning`**: Intermediate artifacts for macOS/Windows signing ceremonies. Not for end-users.
- **Source**: `bitcoin-<version>.tar.gz` (no platform string). For manual compilation.

## Verification Files

| File | Description |
|------|-------------|
| `SHA256SUMS` | Manifest of all release file hashes. |
| `SHA256SUMS.asc` | Detached GPG signatures from trusted signers. |
| `SHA256SUMS.ots` | OpenTimestamps proof. |

For step-by-step verification instructions, see the [download page](/en/download/).
