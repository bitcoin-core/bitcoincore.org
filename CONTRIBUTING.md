Contributing
============

 - [Fork it](https://github.com/bitcoin-core/bitcoincore.org)
 - Create a topic branch
 - Commit patches
 - Create pull request

## Process

Changes to the website are made by pull-request in order to facilitate review process.

All pull-requests must pass the continuous integration tests which test HTML validity, links and images.

Anyone may participate in the review process. All comments and ACK/NACKs will be taken into consideration but the decision about merging rests with the website maintainers.

## Content Policy

The purpose of the website is to be an official mouthpiece for the Bitcoin Core project as well as be a resource for technical information that has a direct impact on Bitcoin Core software. This can include research, presentations, and developer blogs. 

Where proposed content is about Bitcoin Core project policy, there should be rough consensus among the Bitcoin Core project maintainers.

## Translation Process

Translation of the website is done manually for now.
 
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
    ---

Notice the remaining fields are left untranslated.

### Important note

Please be careful when translating Markdown formatting, this must be exact (including spaces). Markdown titles require a space after the header tags, e.g. `## Title`.

Translations, as well as [languages.yml](/_data/languages.yml), [translations.yml](/_data/translations.yml) and [navigation.yml](/_data/navigation.yml) should be translated and submitted as normal pull-requests.

## License

The code and documentation of this website are licensed under [MIT license][MIT] and all contributors agree to irrevocably license their contributions under the same license.

Unless specified in the header of the file, the website content in `_posts/` of this website is licensed under the Creative Commons Attribution-ShareAlike 4.0 International License: [CC-BY-SA] and all contributors agree to irrevocably license their content under the same license.

[MIT]: https://github.com/bitcoin-core/bitcoincore.org/blob/master/LICENSE
[CC-BY-SA]: https://creativecommons.org/licenses/by-sa/4.0/
