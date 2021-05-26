---
layout: none
---
"use strict";

{% comment %}
// Function to expand part of a section when clicked
{% endcomment %}
$(function() {
  $(".show_more").hide();
  $(".toggle_show_more_less").show();

  $(".toggle_show_more_less").click(function() {
    $(".show_more").fadeToggle();
    $(".toggle_show_more_less").hide();
  });
});
