---
title: Security Advisories
name: Security Advisories
id: en-security-advisories
permalink: /en/security-advisories/
layout: page
type: pages
lang: en
version: 1
---
{% include toc.html %}

This page summarizes policies in relation to disclosing vulnerabilities in Bitcoin Core,
as well as provides a summary of historical Security Advisories.

## Policy

When reported, a vulnerability will be assigned a severity category. We
differentiate between 4 classes of vulnerabilities:

* **Low**: bugs which are hard to exploit or have a low impact. For instance
  a wallet bug which requires access to the victim's machine.

* **Medium**: bugs with limited impact. For instance a local network remote crash.

* **High**: bugs with significant impact. For instance a remote crash, or a local network RCE.

* **Critical**: bugs which threaten the whole network's integrity. For instance an inflation or coin theft bug.

**Low** severity bugs will be disclosed 2 weeks after a fixed version is released.
  A pre-announcement will be made at the same time as the release.

**Medium** and **High** severity bugs will be disclosed 2 weeks after the [last
  affected release goes EOL](/en/lifecycle/). This is a year after a fixed version was first
  released. A pre-announcement will be made 2 weeks prior to disclosure.

**Critical** bugs are not considered in the standard policy, as they would most likely require an
  ad-hoc procedure. Also, a bug may not be considered a vulnerability at all. Any reported issue
  may also be considered serious, yet not require embargo.

## Past Security Advisories

{% assign advisories=site.posts | where:"lang", 'en' | where:"type", 'advisory' | sort: "date" | reverse %}
{% for advisory in advisories %}
{% assign post=advisory %}
  <article>
    <h2><a href="{{ post.url }}" title="{{ post.title | xml_escape }}">{{ post.title }}</a></h2>
    <p>{{ post.excerpt | markdownify | strip_html | truncate: 200 }}</p>
  </article>
{% endfor %}
