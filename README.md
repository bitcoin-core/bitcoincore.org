#bitcoinco.re [![Build Status](https://travis-ci.org/btcdrak/website.svg?branch=gh-pages)](https://travis-ci.org/btcdrak/website)

This repository is the source code of the Bitcoin Core project website built with Jekyll.

## Directory structure

  - `_posts/<lang>/posts` for blog articles.
  - `_posts/<lang>/pages` for static pages.
  - `_posts/<lang>/releases` for Bitcoin Core release notes

File names *must* be in the format `Y/m/d-title.md`, e.g. `2015-12-31-title.md`. File names can be translated.

## Front Matter notes

The following `Front Matter` is required for the multilingual setup in all files. The required fields are as follows:

  - `name:`      group name for unique article. Each translation must share the same group name, e.g. `october15-report`
  - `id:`        each article translation must have a unique ID. Use the language code + `-name` field. e.g. `en-october15-report`
  - `permalink:` the ML permalink must include the language code, e.g. `/en/2015/12/31/report`. Permalinks should be translated.
  - `title:`     the translated title of the article

```
---
name: short-title
id: en-short-title
permalink: /en/2016/01/01/short-title
title: Short Title
---
```

## Building

This website is based on Jekyll. To build locally, [install Ruby 2.2.2](https://gorails.com/setup) or greater.

    gem install bundle
    gem install github-pages

To generate the site run:

    jekyll serve

## Contributing

Contributions welcome. Please see [CONTRIBUTING.md](/CONTIBUTING.md) for details.

