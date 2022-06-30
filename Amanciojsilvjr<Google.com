{% capture /dev/null %}<!-- suppress render of this part -->
<!-- Copyright 2013 - 2016 The Bitcoin.org Project.
     Copyright 2017 - 2020 The BitcoinCore.org Project
     This file is licensed under the MIT License (MIT) available on
     http://opensource.org/licenses/MIT. -->
{% assign VERSION_SORTED_RELEASES = site.releases | sort: 'release' | reverse %}
{% capture CURRENT_RELEASE %}{% for subver in VERSION_SORTED_RELEASES[0].release %}{{subver}}{% unless forloop.last %}.{% endunless %}{% endfor %}{% endcapture %}
{% assign magnet = VERSION_SORTED_RELEASES[0].optional_magnetlink %}
{% capture PATH_PREFIX %}/bin/bitcoin-core-{{CURRENT_RELEASE}}{% endcapture %}
{% capture FILE_PREFIX %}bitcoin-{{CURRENT_RELEASE}}{% endcapture %}
{% assign SIGNING_KEY_FINGERPRINT = "01EA5486DE18A882D4C2684590C8019E36C2E964" %}
{% capture SIGNING_KEY_FINGERPRINT_EXPLODED %}{% include fingerprint-split.html hex=SIGNING_KEY_FINGERPRINT %}{% endcapture %}
{% assign GPG_DOWNLOAD_URL = "https://www.gnupg.org/download/index.en.html#binary" %}
{% assign GPG_MACOS_DOWNLOAD_URL = "https://gpgtools.org/" %}
{% assign GPG_WINDOWS_DOWNLOAD_URL = "https://gpg4win.org/download.html" %}
{% assign GITIAN_REPOSITORY_URL = "https://github.com/bitcoin-core/gitian.sigs" %}
{% endcapture %}
<link rel="alternate" type="application/rss+xml" href="/en/releasesrss.xml" title="Bitcoin Core releases">
<div class="download">
  <h2>{{ page.latestversion }} {{CURRENT_RELEASE}} <a type="application/rss+xml" href="/en/releasesrss.xml"><img src="/assets/images/icons/icon_rss.svg" alt="rss" class="rssicon"></a></h2>
  <div class="mainbutton"><a id="downloadbutton" href="{{ PATH_PREFIX }}/{{ FILE_PREFIX }}-{{ site.data.binaries.win64exe }}"><img src="/assets/images/os/but_windows.svg" alt="icon"><span>{{ page.download }}</span></a></div>
  <div class="downloadbox">
    <p>{{ page.downloados }}</p>
    <div>
      <div>
        <img src="/assets/images/os/med_win.png" alt="windows">
        <span>
          <a href="{{ PATH_PREFIX }}/{{ FILE_PREFIX }}-{{ site.data.binaries.win64exe }}" class="dl" id="downloadwinexe">Windows</a>
          <span><a href="{{ PATH_PREFIX }}/{{ FILE_PREFIX }}-{{ site.data.binaries.win64exe }}" class="dl" id="win64exe">exe</a> -
            <a href="{{ PATH_PREFIX }}/{{ FILE_PREFIX }}-{{ site.data.binaries.win64zip }}" class="dl" id="win64zip">zip</a></span>
        </span>
      </div>
      <div>
        <img src="/assets/images/os/med_osx.png" alt="osx">
        <span>
          <a href="{{ PATH_PREFIX }}/{{ FILE_PREFIX }}{{ site.data.binaries.macdmg }}">Mac OS X</a>
          <span><a href="{{ PATH_PREFIX }}/{{ FILE_PREFIX }}{{ site.data.binaries.macdmg }}" class="dl" id="macdmg">dmg</a> -
            <a href="{{ PATH_PREFIX }}/{{ FILE_PREFIX }}-{{ site.data.binaries.mactar }}" class="dl" id="mactar">tar.gz</a></span>
        </span>
      </div>
      <div>
        <img src="/assets/images/os/med_linux.png" alt="linux">
        <span>
          <a href="{{ PATH_PREFIX }}/{{ FILE_PREFIX }}-{{ site.data.binaries.lin64 }}" class="dl" id="lin64">Linux (tgz)</a>
        </span>
      </div>
    </div>
    <div>
      <div>
        <img src="/assets/images/os/arm.png" alt="ARM Linux">
        <span>
          <a href="{{ PATH_PREFIX }}/{{ FILE_PREFIX }}-arm-linux-gnueabihf.tar.gz" class="dl">ARM Linux</a>
          	<span><a href="{{ PATH_PREFIX }}/{{ FILE_PREFIX }}-aarch64-linux-gnu.tar.gz" class="dl" id="lin64arm">64 bit</a> -
            	<a href="{{ PATH_PREFIX }}/{{ FILE_PREFIX }}-arm-linux-gnueabihf.tar.gz" class="dl" id="lin32arm">32 bit</a></span>
        </span>
      </div>
      <div>
        <img src="/assets/images/os/but_riscv.svg" alt="RISC-V Linux">
        <span>
          <a href="{{ PATH_PREFIX }}/{{ FILE_PREFIX }}-{{ site.data.binaries.riscv64 }}.tar.gz" class="dl">RISC-V Linux</a>
            <span><a href="{{ PATH_PREFIX }}/{{ FILE_PREFIX }}-{{ site.data.binaries.riscv64 }}.tar.gz" class="dl" id="lin64riscv">64 bit</a></span>
        </span>
      </div>
      <div>
        <img src="/assets/images/os/med_snap.svg" alt="Snap Store Linux">
        <span>
          <a href="https://snapcraft.io/bitcoin-core" class="dl">Snap Store Linux</a>
        </span>
      </div>
    </div>
    <p class="downloadmore">
      <a href="{{ PATH_PREFIX }}/SHA256SUMS.asc" class="dl">{{ page.downloadsig }}</a><br>
      <a href="{{ PATH_PREFIX }}/{{ FILE_PREFIX }}.torrent" class="dl">{{ page.downloadtorrent }}</a>
        {% if magnet %} <a href="{{ magnet | replace: '&', '\&amp;'}}" class="magnetlink" data-proofer-ignore></a>{% endif %}<br>
      <a href="{{ PATH_PREFIX }}/{{ FILE_PREFIX}}.tar.gz" class="dl">{{ page.source }}</a><br>
      <a href="/en/releases">{{ page.versionhistory }}</a>
    </p>
    <p class="downloadkeys">
      <span>{{ page.releasekeys }}</span>
      v0.11.0+ <code title="{{page.pgp_key_fingerprint}}">{{SIGNING_KEY_FINGERPRINT}}</code><br>
      {% if page.version > 2 %}<i>{{page.key_refresh}}</i><br><code>gpg{{site.strings.gpg_keyserver}} --refresh-keys</code>{% endif %}
    </p>
  </div>

</div>

  <h2 style="text-align: center">{{ page.patient }}</h2>
  <p>{{ page.notesync | replace: '$(DATADIR_SIZE)', site.data.stats.datadir_gb | replace: '$(PRUNED_SIZE)', site.data.stats.pruned_gb | replace: '$(MONTHLY_RANGE_GB)', site.data.stats.monthly_storage_increase_range_gb }} {{ page.full_node_guide }}</p>

{% if page.version > 4 %}
  <h2 style="text-align: center" id="{{page.verify_download | slugify}}">{{page.verify_download}}</h2>
  <p>{{page.verification_recommended}}</p>
  <details>
  {% assign GPG = "C:\Program Files\Gnu\GnuPg\gpg.exe" %}
    <summary><strong>{{page.windows_instructions}}</strong></summary>
    <ol>
      <li><p>{{page.download_release}}</p></li>

      <li><p>{{page.download_checksums}} <a href="{{ PATH_PREFIX }}/SHA256SUMS.asc">SHA256SUMS.asc</a></p></li>

      <li><p>{{page.cd_to_downloads}}</p>

        <pre class="highlight"><code>{{page.cd_example_windows}}</code></pre>

        </li>

      {% capture windows_example_binary %}{{FILE_PREFIX}}-{{site.data.binaries.win64exe}}{% endcapture %}
      <li><p>{{page.generate_checksum | replace: "$(FILE)", windows_example_binary}}</p>

        <pre class="highlight"><code>certUtil -hashfile {{windows_example_binary}} SHA256</code></pre></li>

      <li><p>{{page.ensure_checksum_matches}}</p>

        <pre class="highlight"><code>type SHA256SUMS.asc</code></pre></li>

      <li><p>{{page.install_gpg}} <a
      href="{{GPG_WINDOWS_DOWNLOAD_URL}}">{{page.gpg_download_page}}</a>
      {{page.gpg_download_other}}
      <a href="{{GPG_DOWNLOAD_URL}}">{{page.gpg_download_options}}</a></p></li>

      <li><p>{{page.obtain_release_key}}</p>

        <pre class="highlight"><code>{{GPG}}{{site.strings.gpg_keyserver}} --recv-keys {{SIGNING_KEY_FINGERPRINT}}</code></pre>

        <p>{{page.release_key_obtained}}</p></li>

      <li>{{page.verify_checksums_file}}

        <pre class="highlight"><code>{{GPG}} --verify SHA256SUMS.asc</code></pre></li>

      <li><p>{{page.check_gpg_output}}</p>
        <ol><li><p>{{page.line_starts_with}} <code>gpg: {{page.localized_gpg_good_sig}}</code></p></li>
          <li><p>{{page.complete_line_saying}} <code>{{page.localized_gpg_primary_fingerprint}} {{SIGNING_KEY_FINGERPRINT_EXPLODED}}</code></p></li>
        </ol>

        <p>{{page.gpg_trust_warning}}</p></li>

    </ol>
  </details>
  <details>
    {% assign GPG = "gpg" %}
    <summary><strong>{{page.macos_instructions}}</strong></summary>
    <ol>
      <li><p>{{page.download_release}}</p></li>

      <li><p>{{page.download_checksums}} <a href="{{ PATH_PREFIX }}/SHA256SUMS.asc">SHA256SUMS.asc</a></p></li>

      <li><p>{{page.cd_to_downloads}}</p>

        <pre class="highlight"><code>{{page.cd_example_linux}}</code></pre>

        </li>

      <li><p>{{page.verify_download_checksum}}</p>

        <pre class="highlight"><code>shasum -a 256 --check SHA256SUMS.asc</code></pre>

        <p>{{page.checksum_warning_and_ok | replace, "$(SHASUMS_OK)", page.localized_checksum_ok}} <code>{{FILE_PREFIX}}{{site.data.binaries.macdmg}}: {{page.localized_checksum_ok}}</code></p></li>

      <li><p>{{page.install_gpg}} <a
      href="{{GPG_MACOS_DOWNLOAD_URL}}">{{page.gpg_download_page}}</a>
      {{page.gpg_download_other}}
      <a href="{{GPG_DOWNLOAD_URL}}">{{page.gpg_download_options}}</a></p></li>

      <li><p>{{page.obtain_release_key}}</p>

        <pre class="highlight"><code>gpg{{site.strings.gpg_keyserver}} --recv-keys {{SIGNING_KEY_FINGERPRINT}}</code></pre>

        <p>{{page.release_key_obtained}}</p></li>

      <li>{{page.verify_checksums_file}}

        <pre class="highlight"><code>gpg --verify SHA256SUMS.asc</code></pre></li>

      <li><p>{{page.check_gpg_output}}</p>
        <ol><li><p>{{page.line_starts_with}} <code>gpg: {{page.localized_gpg_good_sig}}</code></p></li>
          <li><p>{{page.complete_line_saying}} <code>{{page.localized_gpg_primary_fingerprint}} {{SIGNING_KEY_FINGERPRINT_EXPLODED}}</code></p></li>
        </ol>

        <p>{{page.gpg_trust_warning}}</p></li>
    </ol>
  </details>

  <details>
    <summary><strong>{{page.linux_instructions}}</strong></summary>
    <ol>
      <li><p>{{page.download_release}}</p></li>

      <li><p>{{page.download_checksums}} <a href="{{ PATH_PREFIX }}/SHA256SUMS.asc">SHA256SUMS.asc</a></p></li>

      <li><p>{{page.cd_to_downloads}}</p>

        <pre class="highlight"><code>{{page.cd_example_linux}}</code></pre>

        </li>

      <li><p>{{page.verify_download_checksum}}</p>

        <pre class="highlight"><code>sha256sum --ignore-missing --check SHA256SUMS.asc</code></pre>

        <p>{{page.checksum_warning_and_ok | replace, "$(SHASUMS_OK)", page.localized_checksum_ok}} <code>{{FILE_PREFIX}}-{{site.data.binaries.lin64}}: {{page.localized_checksum_ok}}</code></p></li>

      <li><p>{{page.obtain_release_key}}</p>

        <pre class="highlight"><code>gpg{{site.strings.gpg_keyserver}} --recv-keys {{SIGNING_KEY_FINGERPRINT}}</code></pre>

        <p>{{page.release_key_obtained}}</p></li>

      <li>{{page.verify_checksums_file}}

        <pre class="highlight"><code>gpg --verify SHA256SUMS.asc</code></pre></li>

      <li><p>{{page.check_gpg_output}}</p>
        <ol><li><p>{{page.line_starts_with}} <code>gpg: {{page.localized_gpg_good_sig}}</code></p></li>
          <li><p>{{page.complete_line_saying}} <code>{{page.localized_gpg_primary_fingerprint}} {{SIGNING_KEY_FINGERPRINT_EXPLODED}}</code></p></li>
        </ol>

        <p>{{page.gpg_trust_warning}}</p></li>

    </ol>
  </details>

  {% if page.version > 3 %}
  <details>
    <summary><strong>{{page.snap_instructions}}</strong></summary>

    <p>{{page.snap_warning}}</p>

  </details>
  {% endif %}

  <h2 style="text-align: center">{{page.build_reproduction}}</h2>

  <p>{{page.additional_steps}}</p>

  <ul>
    <li><p><em>{{page.reproducible_builds}}</em> {{page.build_identical_binaries}}</p></li>

    <li><p><em>{{page.verified_reproduction}}</em> {{page.independently_reproducing}}</p></li>
  </ul>

  <p>{{page.verifying_and_reproducing}} <a href="{{GITIAN_REPOSITORY_URL}}">{{page.gitian_repository}}</a>.</p>
{% endif %}{% comment %}END VERSION > 1 CONTENT{% endcomment %}

<hr>

<p>{{ page.notelicense }}</p>

<script type="text/javascript">
var os = 'windows64';
if (navigator.userAgent.indexOf('Mac') != -1) var os = 'mac'
else if (navigator.userAgent.indexOf('Linux') != -1) var os = 'linux64'
else if (navigator.userAgent.indexOf('WOW64') != -1 || navigator.userAgent.indexOf('Win64') != -1) var os='windows64';
var but = document.getElementById('downloadbutton');
var linkwinexe = document.getElementById('downloadwinexe');
var linkwinzip = document.getElementById('downloadwinzip');
var hrefwin64exe = document.getElementById('win64exe').href;
var hrefwin64zip = document.getElementById('win64zip').href;
var hrefmacdmg = document.getElementById('macdmg').href;
var hrefmactar = document.getElementById('mactar').href;
var hreflin64 = document.getElementById('lin64').href;
switch (os) {
case 'windows64':
	but.getElementsByTagName('IMG')[0].src = '/assets/images/os/but_windows.svg';
	but.href = hrefwin64exe;
	linkwinexe.href = hrefwin64exe;
	linkwinzip.href = hrefwin64zip;
break;
case 'linux64':
	but.getElementsByTagName('IMG')[0].src = '/assets/images/os/but_linux.png';
	but.href = hreflin64;
	linklin.href = hreflin64;
break;
case 'mac':
	but.getElementsByTagName('IMG')[0].src = '/assets/images/os/but_mac.svg';
	but.href = hrefmacdmg;
break;
}

</script>orkut.com/index.html/a0c37d5d47c1403aeeea695aa0e13b92ad063851b7b3757a2a50b176e62fe93b
