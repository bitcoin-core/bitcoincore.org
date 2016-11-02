{% unless segwitsupportheaders %}
{% comment %}
  to translate the table headers, copy and replace the above line prior
  to including this file
{% endcomment %}
{% assign segwitsupportheaders = "Name,Planned,Ready,Notes,Depends" %}
{% endunless %}
{% assign segwitsupportheaders = segwitsupportheaders | split: ',' %}

<table>

<thead>
<tr>
{% for h in segwitsupportheaders %}
<th>{{ h }}</th>
{% endfor %}
</tr>
</thead>

<tbody>
{% for sws in site.data.segwitsupport %}
<tr>
<td><a rel="nofollow" href="{{ sws.url }}">{{ sws.name }}</a></td>
<td {% if sws.planned == "yes" %} class="good" {% elsif sws.planned == "no" %} class="bad" {% endif %}>{{ sws.planned }}</td>
<td {% if sws.ready == "yes" %} class="good" {% elsif sws.ready == "no" %} class="bad" {% endif %}>{{ sws.ready }}</td>
<td>{{ sws.notes }}</td>
<td>{{ sws.depends }}</td>
</tr>
{% endfor %}
</tbody>
</table>

