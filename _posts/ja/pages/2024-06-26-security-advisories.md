---
title: セキュリティアドバイザリ
name: Security Advisories
id: ja-security-advisories
permalink: /ja/security-advisories/
layout: page
type: pages
lang: ja
version: 1
---
{% include toc.html %}

このページでは、Bitcoin Coreの脆弱性の開示に関するポリシーと、
過去のセキュリティアドバイザリの概要を掲載しています。

## ポリシー {#policy}

すべての脆弱性は、security@bitcoincore.org までご報告ください（詳細は、
[SECURITY.md](https://github.com/bitcoin/bitcoin/blob/master/SECURITY.md)をご覧ください）。
脆弱性が報告されると、重大度のカテゴリが割り当てられます。私たちは脆弱性を4つのクラスで区別します。

* **低**: 悪用が難しい、または影響が少ないバグ。たとえば、被害者のマシンへのアクセスを必要とするウォレットのバグなどです。

* **中**: 影響が限定的なバグ。たとえば、ローカルネットワークのリモートクラッシュなどです。

* **高**: 重大な影響を与えるバグ。たとえば、リモートクラッシュやローカルネットワークのリモートコード実行などです。

* **クリティカル**: ネットワーク全体の完全性を脅かすバグ。たとえば、インフレーションやコイン盗難のバグなどです。

重要度が**低**のバグは、現在のメジャーリリースブランチに修正版が登場してから2週間後に公開されます。リリースと同時に事前の発表も行われます。

重要度が**中**と**高**のバグは、[影響を受ける最後のリリースがEOLになって](/ja/lifecycle/)から2週間後に公開されます。
  これは修正バージョンが最初にリリースされてから1年後です。事前の発表は公開の2週間前に行われます。

重要度が**クリティカル**のバグは、アドホックな手順が必要になる可能性があるため、標準ポリシーでは考慮されません。
  また、バグは脆弱性とはみなされない場合もあります。報告された問題が深刻であるとみなされても、
  それが必ずしも情報公開の禁止措置が必要であるとは限りません。

## 過去のセキュリティアドバイザリ {#past-security-advisories}

{% assign advisories=site.posts | where:"lang", 'ja' | where:"type", 'advisory' | sort: "date" | reverse %}
{% for advisory in advisories %}
{% assign post=advisory %}
  <article>
    <h2><a href="{{ post.url }}" title="{{ post.title | xml_escape }}">{{ post.title }}</a></h2>
    <p>{{ post.excerpt | markdownify | strip_html | truncate: 200 }}</p>
  </article>
{% endfor %}
