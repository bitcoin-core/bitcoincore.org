{% capture /dev/null %}<!-- suppress render of this part -->
<!-- Copyright 2013 - 2016 The Bitcoin.org Project.
     Copyright 2017 - 2021 The BitcoinCore.org Project
     This file is licensed under the MIT License (MIT) available on
     http://opensource.org/licenses/MIT. -->
{% assign VERSION_SORTED_RELEASES = site.releases | sort: 'release' | reverse %}
{% assign CURRENT_RELEASE = VERSION_SORTED_RELEASES[0].release | join: "." %}
{% endcapture %}
