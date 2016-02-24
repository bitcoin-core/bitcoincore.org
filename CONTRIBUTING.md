Contributing
============

 - [Fork it](https://github.com/bitcoin-core/website)
 - Create a topic branch
 - Commit patches
 - Create pull request

## Process

Changes to the website are made by pull-request in order to facilitate review process.

All pull-requests must pass the continuous integration tests which test HTML validity, links and images.

Anyone may participate in the review process. All comments and ACK/NACKs will be taken into consideration but the decision about merging rests with the website maintainers.

## Content Policy

The purpose of the website is to be an official mouthpiece for the Bitcoin Core project as well as be a resource for technical information that has a direct impact on Bitcoin Core software. This can include research, presentations, and developer blogs. 

Where proposed content is about Bitcoin Core project policy, there should rough consensus of the Bitcoin Core project maintainers as a minimum bar and final decision is with the Bitcoin Core project lead.

## Translation Process

Translation of the website is done using Transifex. To become a translator please create an account at [Transifex.com](https://transifex.com/), then apply to join the [bitcoincore.org project](https://www.transifex.com/bitcoincore/bitcoincoreorg).
 
Each document has a header called "Front Matter", which looks something like this:

    ---
    title: Clarifying Communications of the Bitcoin Core project
    name: clarifying-communications
    id: en-clarifying-communications
    lang: en
    permalink: /en/2016/01/28/clarification/
    layout: post
    type: post
    layout: post
    share: true
    ---

Translators should only translate the `title:` field, and change the language code in the `id:`, `lang:` and `permalink:` fields. For example the above would be translated for `zh_CN` as:
 
    ---
    title: 有关Bitcoin Core沟通渠道的澄清
    name: clarifying-communications
    id: zh_cn-clarifying-communications
    lang: zn_CN
    permalink: /zh_CN/2016/01/28/clarification/
    type: post
    layout: post
    share: true
    ---

Notice the remaining fields are left untranslated, however in Transifiex, you must copy these fields as translations (without actually translating them).

### Important note

Please be careful when translating Markdown formatting, this must be exact (including spaces). Markdown titles require a space after the header tags, e.g. `## Title`.

Please note that [translations.yml](https://github.com/bitcoin-core/website/blob/gh-pages/_data/translations.yml) and [navigation.yml](https://github.com/bitcoin-core/website/blob/gh-pages/_data/navigation.yml) should be translated and submitted as normal pull-requests because they are not compatible with Transifex at this time.

## License

The code and documentation of this website are licensed under [MIT license][MIT] and all contributors agree to irrevocably license their contributions under the same license.

Unless specified in the header of the file, the website content in `_posts/` of this website is licensed under the Creative Commons Attribution-ShareAlike 4.0 International License: [CC-BY-SA] and all contributors agree to irrevocably license their content under the same license.

[MIT]: https://github.com/bitcoin-core/website/blob/gh-pages/LICENSE
[CC-BY-SA]: http://creativecommons.org/licenses/by-sa/4.0/
