{% capture /dev/null %}
!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!
!!! NB: for this file to render correctly, !!!
!!! you must also include references.md   !!!
!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!

{% case page.lang %}
  {% else %}{% comment %}fallback to English{% endcomment %}
    {% assign _includes_dev_keys-name = "Name" %}
    {% assign _includes_dev_keys-fingerprint = "Fingerprint" %}
    {% capture _includes_dev_keys-import %}You can import a key by running the following command with that individual's fingerprint: 

<pre class="highlight code-pad"><code>gpg{{ site.strings.gpg_keyserver }} --recv-keys "&lt;fingerprint&gt;"</code></pre>

Ensure that you put quotes around fingerprints containing spaces.{% endcapture %}
{% endcase %}
{% endcapture %}

| {{_includes_dev_keys-name}} | {{_includes_dev_keys-fingerprint}}         |
|-----------------------------|--------------------------------------------|
| Michael Ford                | <code>{% include fingerprint-split.html hex="E777299FC265DD04793070EB944D35F9AC3DB76A" %}</code> |
| Ava Chow                    | <code>{% include fingerprint-split.html hex="152812300785C96444D3334D17565732E08E5E41" %}</code> |
| Niklas Gögge                | <code>{% include fingerprint-split.html hex="2CBBF208E594BF439B5F276C7465CFFF6793242E" %}</code> |

{{_includes_dev_keys-import}}