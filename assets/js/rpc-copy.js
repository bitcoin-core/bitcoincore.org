// Copy-to-clipboard button on the plain-text and OpenRPC boxes of the RPC doc pages.
// Button stays at the top right and does not scroll with the box content.
document.addEventListener('DOMContentLoaded', function () {
  var boxes = document.querySelectorAll('.doc figure.highlight, pre.rpc-json');
  Array.prototype.forEach.call(boxes, function (box) {
    var pre = box.tagName === 'PRE' ? box : box.querySelector('pre');
    if (!pre) return;
    var text = pre.textContent;

    var host = document.createElement('div');
    host.className = 'rpc-copy-host';
    box.parentNode.insertBefore(host, box);
    host.appendChild(box);

    var btn = document.createElement('button');
    btn.className = 'rpc-copy';
    btn.type = 'button';
    btn.title = 'Copy to clipboard';
    btn.innerHTML = '<i class="fa fa-copy"></i>';
    host.appendChild(btn);

    function flash(icon) {
      btn.innerHTML = '<i class="fa fa-' + icon + '"></i>';
      setTimeout(function () { btn.innerHTML = '<i class="fa fa-copy"></i>'; }, 1200);
    }
    btn.addEventListener('click', function () {
      if (navigator.clipboard && navigator.clipboard.writeText) {
        navigator.clipboard.writeText(text).then(
          function () { flash('check'); },
          function () { flash('times'); }
        );
        return;
      }
      var ta = document.createElement('textarea');
      ta.value = text;
      document.body.appendChild(ta);
      ta.select();
      try { flash(document.execCommand('copy') ? 'check' : 'times'); }
      catch (e) { flash('times'); }
      document.body.removeChild(ta);
    });
  });
});
