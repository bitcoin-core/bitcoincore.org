{% unless segwitsupportheaders %}
{% comment %}
  to translate the table headers, copy and replace the above line prior
  to including this file
{% endcomment %}
{% assign segwitsupportheaders = ", Name,Status,Notes,Depends" %}
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
{% assign rowid = '0' %}
{% assign ready = '0' %}
{% assign wip = '0' %}
{% assign planned = '0' %}
{% for sws in site.data.segwitsupport %}
{% assign rowid = rowid | plus:'1' %}
{% if sws.status == "ready" %}{% assign rowcolour = 'bg-green' %}{% assign ready = ready | plus:'1' %}
{% elsif sws.status == "wip" %}{% assign rowcolour = 'bg-yellow' %}{% assign wip = wip | plus:'1' %}
{% elsif sws.status == "planned" %}{% assign rowcolour = 'bg-white' %}{% assign planned = planned | plus:'1' %}
{% endif %}
<tr>
<td class="{{ rowcolour }}">#{{ rowid }}</td>
<td class="{{ rowcolour }}"><a rel="nofollow" href="{{ sws.url }}">{{ sws.name }}</a></td>
<td class="{{ rowcolour }}">{{ sws.status }}</td>
<td class="{{ rowcolour }}">{{ sws.notes }}</td>
<td class="{{ rowcolour }}">{{ sws.depends }}</td>
</tr>
{% endfor %}
</tbody>
</table>

Segwit Ready: {{ ready }}, work-in-progress: {{ wip }}, planned: {{ planned }}

