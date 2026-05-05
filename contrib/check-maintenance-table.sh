#!/bin/sh
# Copyright (c) The Bitcoin Core developers
# Distributed under the MIT software license, see the accompanying
# file COPYING or http://www.opensource.org/licenses/mit-license.php.
#
# Checks to ensure the maintenance table is up-to-date with release posts.

latest=$(
  printf '%s\n' _releases/*.md |
    sed -nE 's#^_releases/([0-9]+)\.0\.md$#\1#p' |
    sort -n |
    tail -n 1
)

[ -n "$latest" ] || exit 0

if ! grep -qE "^\|[[:space:]]*${latest}\.x[[:space:]]*\|[[:space:]]*[0-9]{4}-[0-9]{2}-[0-9]{2}[[:space:]]*\|" \
  _includes/posts/maintenance-table.md
then
  echo "error: _includes/posts/maintenance-table.md is missing a dated row for ${latest}.x but ${latest}.0 has shipped" >&2
  exit 1
fi
