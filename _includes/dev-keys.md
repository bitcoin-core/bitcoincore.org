{% capture /dev/null %}
!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!
!!! NB: for this file to render correctly, !!!
!!! you must also include _references.md   !!!
!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!

{% case page.lang %}
  {% else %}{% comment %}fallback to English{% endcomment %}
    {% assign _includes_dev_keys-name = "Name" %}
    {% assign _includes_dev_keys-fingerprint = "Fingerprint" %}
    {% assign _includes_dev_keys-import = "You can import a key by running the following command with that individual's fingerprint: `gpg --recv-keys <fingerprint>`" %}
{% endcase %}
{% endcapture %}

| {{_includes_dev_keys-name}} | {{_includes_dev_keys-fingerprint}}         |
|-----------------------------|--------------------------------------------|
| Wladimir van der Laan       | `71A3B16735405025D447E8F274810B012346C9A6` |
| Jonas Schnelli              | `32EE5C4C3FA15CCADB46ABE529D4BCB6416F53EC` |
| Pieter Wuille               | `133EAC179436F14A5CF1B794860FEB804E669320` |

{{_includes_dev_keys-import}}
