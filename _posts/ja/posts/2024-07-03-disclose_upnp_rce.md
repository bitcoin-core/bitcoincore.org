---
title: CVE-2015-20111 - miniupnpcのバグによるリモートコード実行
name: blog-disclose-upnp-rce
id: ja-blog-disclose-upnp-rce
lang: ja
permalink: /ja/2024/07/03/disclose_upnp_rce/
type: advisory
layout: post

## If this is a new post, reset this counter to 1.
version: 1

## Only true if release announcement or security annoucement. English posts only
announcement: 0

excerpt: >
  miniupnpcライブラリのバグにより、Bitcoin Coreでリモートコードが実行される可能性がありました。
  修正は、2015年10月15日に、Bitcoin Core 0.11.1 でリリースされました。
---

`miniupnpc`で、重大なデータ漏洩を引き起こすバッファオーバーフローが発見されました。
当時公開されたばかりのCVE-2015-6031と組み合わせると、`miniupnpc`でRCE（リモートコード実行）が可能になり、
Bitcoin CoreでのRCEにつながる可能性がありました。これは2016年2月にリリースされた
[Bitcoin Core 0.12](https://bitcoincore.org/en/releases/0.12.0/)で修正されました。

この問題の重大度は**中**です。

## 詳細 {#details}

2015年9月に開示された[CVE-2015-6031](https://nvd.nist.gov/vuln/detail/CVE-2015-6031)により、
悪意あるUPnPサーバーが起動時にローカルネットワーク上のBitcoin Coreプロセスをリモートでクラッシュさせることが可能になりました。
詳細については[こちら](https://nvd.nist.gov/vuln/detail/CVE-2015-6031)をご覧ください。
この修正は、[Bitcoin Coreに取り込まれ](https://github.com/bitcoin/bitcoin/pull/6789)、
2015年10月にリリースされた[バージョン0.11.1](https://bitcoincore.org/en/releases/0.11.1/)でリリースされました。
その後、UPnPは[デフォルトでオフになりました](https://github.com/bitcoin/bitcoin/pull/6795)。

CVE-2015-6031でのバッファオーバーフローが発見され、リモートクラッシュを可能にすることに加えて、
被害者のマシンでリモートコード実行が可能でした。この可能性の調査中に、
Wladimir J. Van Der Laanは、`miniupnpc`で重大なデータ漏洩を招く別のバッファーオーバーフローを発見しました。
これは、[Wladimirによって`miniupnpc`の](https://github.com/miniupnp/miniupnp/pull/157)コミット
`4c90b87ce3d2517097880279e8c3daa7731100e6`で修正されました。この修正は、
[Bitcoin Coreに取り込まれ](https://github.com/bitcoin/bitcoin/pull/6980)、バージョン0.12の一部としてリリースされました。

このデータ漏洩は、（ウォレットの秘密鍵などの）秘密情報を直接開示するものではありませんでした。
しかし、別のスタックオーバーフロー（CVE-2015-6031で開示されたものなど）と組み合わせることで、
リモートコード実行を引き起こすことが可能でした。Wladimirは、
これをUbuntuの`miniupnpc`バージョン`1.6-precise`で実証しました。しかし、
このエクスプロイトで使用された特定のアプローチは、Bitcoin Coreには直接移植できませんでした。

## 貢献 {#attribution}

CVE-2015-0035を特定したAleksandar Nikolicとその影響を調査し、2つめのバッファーオーバーフローを発見した
Wladimir J. Van Der Laanに感謝します。

## タイムライン {#timeline}

- 2015-09-15 CVE-2015-0035が[修正され](https://github.com/miniupnp/miniupnp/commit/79cca974a4c2ab1199786732a67ff6d898051b78)、[開示される](https://talosintelligence.com/vulnerability_reports/TALOS-2015-0035/)
- 2015-10-09 [PR #6789](https://github.com/bitcoin/bitcoin/pull/6789)がBitcoin Coreにマージされる
- 2015-10-14 Wladimirによる2つめのバッファーオーバーフローを利用したリモートコード実行がUbuntuセキュリティとBitcoin開発者に開示される
- 2015-10-15 Bitcoin Core 0.11.1[がリリースされる](https://lists.linuxfoundation.org/pipermail/bitcoin-dev/2015-October/011545.html)
- 2015-10-26 2つめのバッファーオーバーフローの修正が`miniupnpc`に[マージされる](https://github.com/miniupnp/miniupnp/pull/157)
- 2015-12-18 修正が[Bitcoin Coreに取り込まれる](https://github.com/bitcoin/bitcoin/pull/6980)
- 2016-02-23 Bitcoin Core バージョン 0.12 [がリリースされる](https://lists.linuxfoundation.org/pipermail/bitcoin-dev/2016-February/012456.html).
- 2017-03-08 脆弱性のある最後のバージョンのBitcoin Core (0.11.x) がEOLになる
- 2024-07-03 公開

{% include references.md %}
