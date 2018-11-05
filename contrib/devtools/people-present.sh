#!/bin/bash -eu

## people-present.sh
## Turn lines from meetbot "People present" meeting summary section into table
## for published meeting notes.  Reads from stdin

DB=$(mktemp)

cat <<EOF > $DB
| BlueMatt        | [Matt Corallo][]          |
| LukeJr          | [Luke Dashjr][]           |
| MarcoFalke      | [Marco Falke][]           |
| Murch           | [Mark Erhardt][]          |
| achow101        | [Andrew Chow][]           |
| aj              | [Anthony Towns][]         |
| cfields         | [Cory Fields][]           |
| echeveria       | Echeveria                 |
| fanquake        | [Michael Ford][]          |
| gmaxwell        | [Gregory Maxwell][]       |
| gwillen         | [Glenn Willen][]          |
| instagibbs      | [Gregory Sanders][]       |
| jamesob         | [James O'Beirne][]        |
| jarthur         | [Justin Arthur][]         |
| jcorgan         | [Johnathan Corgan][]      |
| jimpo           | [Jim Posen][]             |
| jnewbery        | [John Newbery][]          |
| jonasschenlli   | [Jonas Schnelli][]        |
| jonasschnelli   | [Jonas Schnelli][]        |
| jtimon          | [Jorge Timón][]           |
| kanzure         | [Bryan Bishop][]          |
| ken2812221      | [Chun Kuan Lee][]         |
| lclc            | [Lucas Betschart][]       |
| luke-jr         | [Luke Dashjr][]           |
| meshcollider    | [Samuel Dobson][]         |
| moneyball       | [Steve Lee][]             |
| morcos          | [Alex Morcos][]           |
| nmnkgl          | [Gleb Naumenko][]         |
| phantomcircuit  | [Patrick Strateman][]     |
| promag          | [Joao Barbosa][]          |
| provoostenator  | [Sjors Provoost][]        |
| randolf         | [Randolf Richardson][]    |
| roasbeef        | [Olaoluwa Osuntokun][]    |
| ryanofsky       | [Russell Yanofsky][]      |
| sdaftuar        | [Suhas Daftuar][]         |
| sipa            | [Pieter Wuille][]         |
| skeees          | [Jesse Cohen][]           |
| wumpus          | [Wladimir van der Laan][] |
EOF

echo '| IRC nick        | Name/Nym                  |'
echo '|-----------------|---------------------------|'

## Example: "     BlueMatt (96)"
cat | sed 's/ (.*//; s/ //g;' | while read line
do
    case "$line" in
      gribble|lightningbot|bitcoin-git)
        continue
      ;;
    esac

    grep "$line" $DB || echo $line
done

rm $DB
