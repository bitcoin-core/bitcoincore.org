# [bitcoincore.org](https://bitcoincore.org) [![Build Status](https://travis-ci.org/bitcoin-core/bitcoincore.org.svg?branch=master)](https://travis-ci.org/bitcoin-core/bitcoincore.org)

This repository is the source code of the Bitcoin Core project website built with Jekyll.

## Directory structure

  - `_posts/<lang>/posts` for blog articles.
  - `_posts/<lang>/pages` for static pages.
  - `_posts/<lang>/releases` for Bitcoin Core release notes

File names *must* be in the format `Y-m-d-title.md`, e.g. `2015-12-31-title.md`. File names can be translated.

## Translations

See [CONTRIBUTING.md](/CONTRIBUTING.md#translation-process) for more information.

## Front Matter notes

The following `Front Matter` is required for the multilingual setup in all files. The required fields are as follows:

  - `name:`      group name for unique article. Each translation must share the same group name, e.g. `october15-report`
  - `id:`        each article translation must have a unique ID. Use the language code + `-name` field. e.g. `en-october15-report`
  - `permalink:` the must include the language code, end a trailing slash, e.g. `/en/2015/12/31/report/`.
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

This website is based on [Jekyll](https://jekyllrb.com/).  To build
locally, [install Ruby 2.5.5](https://gorails.com/setup) using system
packages, [rvm](https://rvm.io), [rbenv](https://github.com/rbenv/rbenv), or another method.
Then clone this repository and change directory into it:

    git clone https://github.com/bitcoin-core/bitcoincore.org.git
    cd bitcoincore.org

Install the `bundle` utility using the Ruby package manager, `gem`, and
then use `bundle` to install the rest of the Ruby packages needed to
build this site.  Note, depending on your system configuration, you may
need to run `gem` as the superuser by putting "sudo" followed by a space
before the `gem` command.  You shouldn't need to use `sudo` with the
`bundle` command.

    gem update --system
    gem install bundler
    bundle install

To preview the site (this will launch a tiny webserver on port 4000):

    bundle exec jekyll server --future

To simply build the site (output placed in the `_site` directory):

    bundle exec jekyll build --future

Note that the `--future` parameter is only required if you're adding any
pages dated in the future (such as prepared release announcements).

To test the site:

    bundle exec jekyll build --future --drafts --unpublished
    bundle exec htmlproofer --disable-external --url-ignore '/^\/bin/.*/' ./_site

The additional parameters to `jekyll build` ensure that all possible
pages are built and checked.

## Contributing

Contributions welcome. Please see [CONTRIBUTING.md](/CONTRIBUTING.md) for details.

## References

The website uses an old version of the [Minimal Mistakes theme][].  The
theme's website provides [documentation][mm docs], including information
about [configuration variables][mm config], creating [pages and posts][mm
content], adding [new Javascript][mm js], and more.  Note that
current documentation may describe features not available in the old
version of the theme used by the website.

[minimal mistakes theme]: https://mmistakes.github.io/minimal-mistakes/
[mm docs]: https://mmistakes.github.io/minimal-mistakes/docs/quick-start-guide/
[mm config]: https://mmistakes.github.io/minimal-mistakes/docs/configuration/
[mm content]: https://mmistakes.github.io/minimal-mistakes/docs/posts/
[mm js]: https://mmistakes.github.io/minimal-mistakes/docs/javascript/
