// Collapsible result-schema trees on the RPC doc pages. Every list item that
// has a nested list gets an editor-style fold toggle; the markup is otherwise
// left untouched.
document.addEventListener('DOMContentLoaded', function () {
  var items = document.querySelectorAll('.rpc-doc ul.rpc-schema li');
  Array.prototype.forEach.call(items, function (li) {
    var hasChildren = false;
    for (var i = 0; i < li.children.length; i++) {
      if (li.children[i].tagName === 'UL') { hasChildren = true; break; }
    }
    if (!hasChildren) return;
    var btn = document.createElement('button');
    btn.className = 'rpc-fold';
    btn.type = 'button';
    btn.setAttribute('aria-expanded', 'true');
    btn.setAttribute('aria-label', 'Collapse');
    btn.addEventListener('click', function () {
      var collapsed = li.classList.toggle('rpc-collapsed');
      btn.setAttribute('aria-expanded', String(!collapsed));
      btn.setAttribute('aria-label', collapsed ? 'Expand' : 'Collapse');
    });
    li.insertBefore(btn, li.firstChild);
  });
});
