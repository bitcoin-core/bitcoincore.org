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
