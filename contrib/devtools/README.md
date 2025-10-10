Contents
========
This directory contains tools for developers working on this repository.

people-present.sh
=================

Turn lines from meetbot "People present" meeting summary section into table
for published meeting notes.

example usage:

```
> cat pp.txt
sipa (62)
provoostenator (20)
gribble (18)
[...]
> cat pp.txt | people-present.sh
| IRC nick        | Name/Nym                  |
|-----------------|---------------------------|
| sipa            | [Pieter Wuille][]         |
| provoostenator  | [Sjors Provoost][]        |
| luke-jr         | [Luke Dashjr][]           |
[...]
```

Dockerfile
==========

Run the website locally without having to install the dependencies on your machine.

To build the image (from the root of the repository):
```
docker build -f ./contrib/devtools/Dockerfile -t corewebsite .
```

Then the container can be run like so:
```
docker run -it --rm -v "$PWD":/site -p 4000:4000 corewebsite
```

And the website can be accessed at https://localhost:4000.
