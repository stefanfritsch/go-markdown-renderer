
addEventListener("DOMContentLoaded", (event) => {
  var toc = document.getElementById('md-toc');
  if (toc !== null)
    document.getElementById('this_page').appendChild(toc);

  if (window.innerWidth < 100 / 18 * document.getElementById("sidebar").offsetWidth)
    hide_sidebar();


  // ========================================================================== //
  //   apply cookie settings
  var cookie = read_cookie();

  for (key in cookie) {
    if (!["sidebar", "this_page", "pages", "backlinks", "search"].includes(key))
      continue
    if (key === "sidebar") {
      if (cookie.sidebar === "true") {
        show_sidebar();
      } else {
        hide_sidebar();
      }
    } else {
      if (cookie[key] === "true") {
        show_sidebar_entry(key);
      } else {
        hide_sidebar_entry(key);
      }
    }
  }

});

addEventListener('beforeunload', (_) => {
  write_cookie(cookie);
});

// ========================================================================== //
// ScrollSpy powers the page toc

try {
  var scrollSpy = new bootstrap.ScrollSpy(document.getElementById('content'), {
    target: '#md-toc'
  });
} catch (error) {
  console.error(error);
  var scrollSpy = null;
}

// ========================================================================== //
// Highlight the current page in the file tree

var acc = document.getElementsByClassName("subdir-button");
var i;

for (i = 0; i < acc.length; i++) {
  acc[i].addEventListener("click", function () {
    /* Toggle between adding and removing the "active" class,
    to highlight the button that controls the panel */
    this.classList.toggle("active");

    /* Toggle between hiding and showing the active panel */
    var panel = this.nextElementSibling;
    if (panel.style.display === "block") {
      panel.style.display = "none";
    } else {
      panel.style.display = "block";
    }
  });
}
