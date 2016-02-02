# [bitcoincore.org](https://bitcoincore.org) [![Build Status](https://travis-ci.org/bitcoin-core/website.svg?branch=gh-pages)](https://travis-ci.org/bitcoin-core/website)

This repository is the source code of the Bitcoin Core project website built with Jekyll.

## Directory structure

  - `_posts/<lang>/posts` for blog articles.
  - `_posts/<lang>/pages` for static pages.
  - `_posts/<lang>/releases` for Bitcoin Core release notes

File names *must* be in the format `Y-m-d-title.md`, e.g. `2015-12-31-title.md`. File names can be translated.

## Translations

Menu and miscellaneous translations can be found in:

  - `_data/lanaguges.yml`
  - `_data/navgation.yml`
  - `_data/translations.yml`

## Front Matter notes

The following `Front Matter` is required for the multilingual setup in all files. The required fields are as follows:

  - `name:`      group name for unique article. Each translation must share the same group name, e.g. `october15-report`
  - `id:`        each article translation must have a unique ID. Use the language code + `-name` field. e.g. `en-october15-report`
  - `permalink:` the ML permalink must include the language code, e.g. `/en/2015/12/31/report`.
  - `title:`     the translated title of the article
  - `type:`      the content type (`pages`, `posts`, `releases`, `meetings`, etc.)
  - `layout:`    the layout template (`page`, `post`, etc.)
  - `lang:`      the language code (`en`, `fr`, etc.)

```
---
title: Short Title
name: short-title
layout: post
type: posts
lang: en
id: en-short-title
permalink: /en/2016/01/01/short-title
---
```

## Building

This website is based on Jekyll. To build locally, [install Ruby 2.2.2](https://gorails.com/setup) or greater
and then run the following commands:

    gem install bundle
    bundle install

To preview the site (this will launch a tiny webserver on port 4000):

    bundle exec jekyll server

To simply build the site (output placed in the `_site` directory):

    bundle exec jekyll build

To test the site:

    bundle exec jekyll build && bundle exec htmlproof ./_site

## Contributing

Contributions welcome. Please see [CONTRIBUTING.md](/CONTRIBUTING.md) for details.
