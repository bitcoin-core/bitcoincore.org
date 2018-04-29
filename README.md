HanCoin integration/staging tree
http://www.hancoin.io

Copyright (c) 2009-2018 Bitcoin Developers 
Copyright (c) 2017-2018 HanCoin Developers

What is HanCoin?
HanCoin is a lite version of Bitcoin using scrypt as a proof-of-work algorithm.

2.5 minute block targets
subsidy halves in 840k blocks (~4 years)
540 million total coins
140 million premined
The rest is the similar as Bitcoin.

1000 coins per block reward for han miners
2016 blocks to retarget difficulty
For more information, as well as an immediately useable, binary version of the HanCoin client sofware, see http://www.hancoin.io.

License
HanCoin is released under the terms of the MIT license. See COPYING for more information or see http://opensource.org/licenses/MIT.

Development process
Developers work in their own trees, then submit pull requests when they think their feature or bug fix is ready.

If it is a simple/trivial/non-controversial change, then one of the HanCoin development team members simply pulls it.

If it is a more complicated or potentially controversial change, then the patch submitter will be asked to start a discussion with the devs and community.

The patch will be accepted if there is broad consensus that it is a good thing. Developers should expect to rework and resubmit patches if the code doesn't match the project's coding conventions (see doc/coding.txt) or are controversial.

The master branch is regularly built and tested, but is not guaranteed to be completely stable. Tags are created regularly to indicate new official, stable release versions of HanCoin.

Testing
Testing and code review is the bottleneck for development; we get more pull requests than we can review and test. Please be patient and help out, and remember this is a security-critical project where any mistake might cost people lots of money.

Automated Testing
Developers are strongly encouraged to write unit tests for new code, and to submit new unit tests for old code.

Unit tests for the core code are in src/test/. To compile and run them:

cd src; make -f makefile.unix test
Unit tests for the GUI code are in src/qt/test/. To compile and run them:

qmake BITCOIN_QT_TEST=1 -o Makefile.test bitcoin-qt.pro
make -f Makefile.test
./hancoin-qt_test

Our first client that is currently mineable and sustained.

Coinbase maturity | 20 blocks
-- | --
Target spacing | 5 minutes
Target timespan | 10 minutes
Transaction confirmations | 6 blocks
Seednode 1 | node.hancoin.com
Seednode 2 | node2.hancoin.com

Getting started
Use the following instructions to mine a block on your blockchain


Open your wallet, and make sure you are connected to another wallet. 
You are connected when you see the icon Wallet Connections in the lower right corner of your wallet.

The message "No block source available" will disappear once you mine your first block.


Close your wallet and create the file hancoin.conf in the folder "%APPDATA%\hancoin\".

Paste the following text into hancoin.conf and save the file.

addnode=node.hancoin.com
addnode=node2.hancoin.com
Open your wallet.

Go to Help. 
Click Debug Window. 
This is the console where you will execute all commands.

Type this command to start mining your first block:

setgenerate true 8

If you want to use a specific number of CPU cores, instead of -1, type the number of cores.

You can type the following command to see the status of generation.

getmininginfo

It will take about +/- 30 minutes to mine your first block, depending on your computer hardware.

