let cookie = {
  sidebar: "true",
  search: "true",
  this_page: "true",
  pages: "true",
  backlinks: "true",
}


const write_cookie = (cookie) => {
  const d = new Date();
  d.setTime(d.getTime() + 10 * 3600 * 1000);
  const expires = "expires=" + d.toUTCString();

  for (entry in cookie) {
    if (entry == "expires" || entry == "path")
      next

    document.cookie = entry + "=" + cookie[entry] + ";" + expires + ";path=/";
  }
}

// function getCookie(cname) {
//   const name = cname + "=";
//   const ca = document.cookie.split(";");

//   for (let i = 0; i < ca.length; i++) {
//     let c = ca[i];

//     while (c.charAt(0) == " ") {
//       c = c.substring(1);
//     }

//     if (c.indexOf(name) == 0) {
//       return c.substring(name.length, c.length);
//     }
//   }

//   return "";
// }

const read_cookie = () => {
  let cookie = {};

  document.cookie.split(';').forEach(function (el) {
    let [k, v] = el.split('=');
    cookie[k.trim()] = v;
  })

  return cookie;
}
