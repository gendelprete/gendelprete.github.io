function showSmallNav() {
  let nav = document.getElementsByClassName("small-nav-button");
  for (let i = 0; i < nav.length; i++) {
    nav[i].classList.toggle("hidden");
  }
}
document.getElementById("three-bars").onclick = showSmallNav;
