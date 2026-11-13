# Bitcoin Core OpenRPC specs

OpenRPC 1.3.2 documents for the public JSON-RPC surface of Bitcoin Core
releases. Each file generated from the node itself.

For each release tag, the `getopenrpcinfo` RPC was cherry-picked onto the
release, the node was built, started on regtest, and `getopenrpcinfo` was
called to dump the RPC metadata. The output is the node's own view of its
RPC interface. Parameter names, types, defaults and result schemas are exact.

The dump uses `show_hidden=false`, so only the public RPC surface is included.
Don't include hidden/testing RPCs (generatetoaddress, invalidateblock, etc.).

Note that `getopenrpcinfo` is removed because it only exists because it was
cherry-picked in to produce these documents. Logically, it belongs in the set
only when a corresponding Core release includes it.

## Generating newer version specs

Backport PR [34683](https://github.com/bitcoin/bitcoin/pull/34683) into newer
version, then run command `bitcoin-cli getopenrpcinfo` to generate the spec.

For older versions before v31, the PR was backported. The backport get
progressively less trustworthy every version down. Down to version v26, no
workarounds were needed (the PR cherry-picked cleanly onto the release tag,
generally safe). From v25 down, progressively more problems to work
around were encountered and thus we cap at v26.

The Bitcoin Core branches that generated each version can be found on:

 Version 31.0: https://github.com/satsfy/bitcoin/tree/31.x-openrpc-backport
- 30.2: https://github.com/satsfy/bitcoin/tree/30.x-openrpc-backport
- 29.2: https://github.com/satsfy/bitcoin/tree/29.x-openrpc-backport
- 28.2: https://github.com/satsfy/bitcoin/tree/28.x-openrpc-backport
- 27.2: https://github.com/satsfy/bitcoin/tree/27.x-openrpc-backport
- 26.2: https://github.com/satsfy/bitcoin/tree/26.x-openrpc-backport
