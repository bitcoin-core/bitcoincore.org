Copyright (c) 2009-2013 Bitcoin Developers
Distributed under the MIT/X11 software license, see the accompanying
file COPYING or http://www.opensource.org/licenses/mit-license.php.
This product includes software developed by the OpenSSL Project for use in the [OpenSSL Toolkit](http://www.openssl.org/). This product includes
cryptographic software written by Eric Young ([eay@cryptsoft.com](mailto:eay@cryptsoft.com)), and UPnP software written by Thomas Bernard.


See readme-qt.rst for instructions on building HanCoin-Qt, the
graphical user interface.

WINDOWS BUILD NOTES
===================

Compilers Supported
-------------------
TODO: What works?
Note: releases are cross-compiled using mingw running on Linux.


Dependencies
------------
Libraries you need to download separately and build:

                default path               download
OpenSSL         \openssl-1.0.1c-mgw        http://www.openssl.org/source/
Berkeley DB     \db-4.8.30.NC-mgw          http://www.oracle.com/technology/software/products/berkeley-db/index.html
Boost           \boost-1.50.0-mgw          http://www.boost.org/users/download/
miniupnpc       \miniupnpc-1.6-mgw         http://miniupnp.tuxfamily.org/files/

Their licenses:

	OpenSSL        Old BSD license with the problematic advertising requirement
	Berkeley DB    New BSD license with additional requirement that linked software must be free open source
	Boost          MIT-like license
	miniupnpc      New (3-clause) BSD license

Versions used in this release:

	OpenSSL      1.0.1c
	Berkeley DB  4.8.30.NC
	Boost        1.50.0
	miniupnpc    1.6


OpenSSL
-------
MSYS shell:

un-tar sources with MSYS 'tar xfz' to avoid issue with symlinks (OpenSSL ticket 2377)
change 'MAKE' env. variable from 'C:\MinGW32\bin\mingw32-make.exe' to '/c/MinGW32/bin/mingw32-make.exe'

	cd /c/openssl-1.0.1c-mgw
	./config
	make

Berkeley DB
-----------
MSYS shell:

	cd /c/db-4.8.30.NC-mgw/build_unix
	sh ../dist/configure --enable-mingw --enable-cxx
	make

Boost
-----
MSYS shell:

	downloaded boost jam 3.1.18
	cd \boost-1.50.0-mgw
	bjam toolset=gcc --build-type=complete stage

MiniUPnPc
---------
UPnP support is optional, make with `USE_UPNP=` to disable it.

MSYS shell:

	cd /c/miniupnpc-1.6-mgw
	make -f Makefile.mingw
	mkdir miniupnpc
	cp *.h miniupnpc/

HanCoin
-------
MSYS shell:

	cd \hancoin\src
	mingw32-make -f makefile.mingw
	strip hancoind.exe
