{% capture links %}
{% capture home_localized %}{% if page.lang != 'en' %}{{page.lang}}/{% endif %}{% endcapture %}
[Bitcoin Core](/{{home_localized}}) /
{%- for crumb in page.breadcrumbs -%}
&nbsp;[{{site.data.translations[page.lang][crumb] | default: site.data.translations.en[crumb] | default: "TRANSLATION_FAILURE" }}][breadcrumb {{site.data.translations.en[crumb] | default: "TRANSLATION_FAILURE"}}]&nbsp;/
{%- endfor %}

[breadcrumb features]: /{{page.lang}}/features/
[breadcrumb verification]: /{{page.lang}}/features/verification/
{% endcapture %}
{{ links | markdownify }}

