function showSmallNav() {
  let nav = document.getElementsByClassName("small-nav-button");
  for (let i = 0; i < nav.length; i++) {
    nav[i].classList.toggle("hidden");
  }
}
document.getElementById("three-bars").onclick = showSmallNav;


new TypeIt("#type-carousel", {
  speed: 60,
  deleteSpeed: 30,
  waitUntilVisible: true,
  loop: true,
  cursor: true
})
  .type("UC Berkeley alumni.")
  .pause(1500)
  .delete()
  .type("San Francisco native.")
  .pause(1500)
  .delete()
  .type("twin.")
  .pause(1500)
  .delete()
  .type("boba enthusiast.")
  .pause(1500)
  .delete()
  .go();
