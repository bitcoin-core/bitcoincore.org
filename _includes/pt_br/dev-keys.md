{% capture /dev/null %}
!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!
!!! NB: for this file to render correctly, !!!
!!! you must also include references.md   !!!
!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!

{% case page.lang %}
  {% else %}{% comment %}fallback to English{% endcomment %}
    {% assign _includes_dev_keys-name = "Nome" %}
    {% assign _includes_dev_keys-fingerprint = "Fingerprint" %}
    {% capture _includes_dev_keys-import %}Você pode importar a chave executando o seguinte comando com uma único fingerprint : `gpg{{site.strings.gpg_keyserver}} --recv-keys "<fingerprint>"`  Certifique-se de colocar aspas em torno de fingerprint que contenham espaços.{% endcapture %}
{% endcase %}
{% endcapture %}

| {{_includes_dev_keys-name}} | {{_includes_dev_keys-fingerprint}}         |
|-----------------------------|--------------------------------------------|
| Wladimir van der Laan       | <code>{% include fingerprint-split.html hex="71A3B16735405025D447E8F274810B012346C9A6" %}</code> |
| Pieter Wuille               | <code>{% include fingerprint-split.html hex="133EAC179436F14A5CF1B794860FEB804E669320" %}</code> |
| Michael Ford                | <code>{% include fingerprint-split.html hex="E777299FC265DD04793070EB944D35F9AC3DB76A" %}</code> |

{{_includes_dev_keys-import}}
