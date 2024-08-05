/*! Responsive Menu */
// http://tympanus.net/codrops/2013/05/08/responsive-retina-ready-menu/
//  The function to change the class

// change root element class from no-js to js
// copied from https://github.com/Modernizr/Modernizr/blob/5eea7e2a213edc9e83a47b6414d0250468d83471/src/setClasses.js
function addJs() {
  var docElement = document.documentElement;
  var className = docElement.className;

  var reJS = new RegExp('(^|\\s)no-js(\\s|$)');
  className = className.replace(reJS, '$1js$2');
  docElement.className = className;
}

addJs();

var changeClass = function (r,className1,className2) {
  var regex = new RegExp("(?:^|\\s+)" + className1 + "(?:\\s+|$)");
  if( regex.test(r.className) ) {
    r.className = r.className.replace(regex,' '+className2+' ');
    }
    else{
    r.className = r.className.replace(new RegExp("(?:^|\\s+)" + className2 + "(?:\\s+|$)"),' '+className1+' ');
    }
    return r.className;
};
//  Creating our button in JS for smaller screens
var menuElements = document.getElementById('site-nav');
menuElements.insertAdjacentHTML('afterBegin','<button type="button" role="button" id="menutoggle" class="navtoggle navicon-lines-button x" aria-hidden="true"><span class="navicon-lines"></span>menu</button>');

//  Toggle the class on click to show / hide the menu
document.getElementById('menutoggle').onclick = function() {
  changeClass(this, 'navtoggle active', 'navtoggle');
};
// http://tympanus.net/codrops/2013/05/08/responsive-retina-ready-menu/comment-page-2/#comment-438918
document.onclick = function(e) {
  var mobileButton = document.getElementById('menutoggle'),
    buttonStyle =  mobileButton.currentStyle ? mobileButton.currentStyle.display : getComputedStyle(mobileButton, null).display;

  if(buttonStyle === 'block' && e.target !== mobileButton && new RegExp(' ' + 'active' + ' ').test(' ' + mobileButton.className + ' ')) {
    changeClass(mobileButton, 'navtoggle active', 'navtoggle');
  }
};

/*! Plugin options and other jQuery stuff */

// Table of Contents toggle
$(function() {
  $(".toc header").each(function(){
    var header = $(this);
    var drawer = header.find(".toc-drawer");
    header.find(".toc-header").click(function(){
      drawer.toggleClass("js-hidden");
    });
  });
  $(".toc-drawer.js-hide-on-start").each(function(){
    var drawer = $(this);
    drawer.toggleClass("js-hidden");
  });
});

// Search functionality
var toggleDrawers = function(show){
  $(".toc header").each(function(){
    var header = $(this);
    var drawer = header.find(".toc-drawer");
    if (show) {
      drawer.removeClass("js-hidden");
    } else {
      drawer.addClass("js-hidden");
    }
  });
};

$(function() {
  $("#searchbar").on('input', function(){
    var query = $(this).val();
    if (query.length > 1) {
      toggleDrawers(true);
      $(".leaf-article").each(function(){
        $(this).show();
        if ($(this).text().indexOf(query) == -1) {
          $(this).hide();
        }
      });
    } else {
      $(".leaf-article").each(function(){ $(this).show(); });
      toggleDrawers(false);
    }
  })
});
