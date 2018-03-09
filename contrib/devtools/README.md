Contents
========
Tools for developers working on this repository.

github-merge.py
===============
This script is used to automate the merging of "pull requests" securely, and signing them with GPG.

For example:

    ./github-merge.py 3077

(in any git repository) will help you merge pull request #3077 for the bitcoin-core/bitcoincore.org repository.

What it does:
  * fetches master, and the pull request
  * locally constructs a merge commit
  * shows the diff that the merge results in
  * asks to verify the resulting source tree (so you can do a make check, etc.)
  * asks whether to sign the merge commit using GPG
  * asks whether to push the result upstream

This means that there are no potential race conditions (e.g., where a pull request gets updated while you're reviewing it, but prior to merging it), and when using GPG signatures so that even a compromised GitHub repository couldn't alter the sources.

Setup
=====
Configuring the github-merge tool for the bitcoin repository is done in the following way:

    git config githubmerge.repository bitcoin-core/bitcoincore.org
    git config githubmerge.testcmd "make -j4 check" (adapt to whatever you want to use for testing)
    git config --global user.signingkey mykeyid (if you want to GPG sign)
