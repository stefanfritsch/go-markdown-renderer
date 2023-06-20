const show_sidebar = () => {
  const tgt = document.getElementById("sidebar");
  const tgtbt = document.getElementById("toggle-hamburger-btn");

  if (tgt !== undefined) {
    tgt.style.display = "flex";
    cookie.sidebar = "true";
  }
  if (tgtbt !== undefined)
    tgtbt.classList.add("my-bg-lighterer");
}

const hide_sidebar = () => {
  const tgt = document.getElementById("sidebar");
  const tgtbt = document.getElementById("toggle-hamburger-btn");

  if (tgt !== undefined) {
    tgt.style.display = "none";
    cookie.sidebar = "false";
  }
  if (tgtbt !== undefined)
    tgtbt.classList.remove("my-bg-lighterer");
}

const toggle_sidebar = () => {
  const tgt = document.getElementById("sidebar");

  if (tgt.style.display === "none") {
    show_sidebar(cookie);
  } else {
    hide_sidebar(cookie);
  }
}


const show_sidebar_entry = (my_id) => {
  const tgtsub = document.getElementById(my_id);
  const tgtbt = document.getElementById("toggle-" + my_id + "-btn");

  if (tgtsub !== undefined) {
    tgtsub.style.display = "flex";
    cookie[my_id] = "true";
  }
  if (tgtbt !== undefined)
    tgtbt.classList.add("my-bg-lighterer");
}

const hide_sidebar_entry = (my_id) => {
  const tgtsub = document.getElementById(my_id);
  const tgtbt = document.getElementById("toggle-" + my_id + "-btn");

  if (tgtsub !== undefined) {
    tgtsub.style.display = "none";
    cookie[my_id] = "false";
  }

  if (tgtbt !== undefined)
    tgtbt.classList.remove("my-bg-lighterer");
}

const toggle_sidebar_entry = (my_id) => {
  const tgt = document.getElementById("sidebar");
  const tgtsub = document.getElementById(my_id);

  if (tgt.style.display === "none" || tgtsub.style.display === "none") {
    show_sidebar();
    show_sidebar_entry(my_id, cookie);
  } else {
    hide_sidebar_entry(my_id, cookie);
  }
}
