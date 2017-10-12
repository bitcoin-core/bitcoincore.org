#!/bin/bash 

# Local builds and Travis CI builds of the BitcoinCore.org website do
# not include the /bin directory holding Bitcoin Core binaries, so this
# test extracts the local URLs from the download page and checks that
# they exist on the remote server.  Only the file's HTTP headers are
# downloaded, not the entire binaries.
#
# Often we prepare a release PR for the website before the binaries are
# uploaded.  It is expected that this test will fail in that case, and
# simply restarting the tests after the binaries have been uploaded
# should resolve the issue and allow a clean merge.
#
# If all URLs work: prints nothing, returns 0.
# If any URLs fail: prints an error for each problem URL, returns 1.

set -e ## Exit on unexpected false return
set -u ## Exit on call of unitintialized variable

grep 'class="dl"' _site/en/download/index.html | sed 's/.*href="//; s/".*//' | while read url
do
  ## Handle relative and absolute URLs differently
  if [ "${url##http*}" ]
  then
    curl -sI "https://bitcoincore.org${url}"
  else
    curl -sI "$url"
  fi | grep -q '200 OK' || echo "Error: Could not retrieve $url"
done | if grep . ; then sed 1iERROR ; false ; else true ; fi
