// Created By: Genevieve Del Prete
document.getElementById('back-to-top').onmouseover = function() {
  document.getElementById('top-label').classList.remove('hidden');
}

document.getElementById('back-to-top').onmouseout = function() {
  document.getElementById('top-label').classList.add('hidden');
}

new fullpage('#fullpage', {
  anchors: ['home', 'experience', 'projects', 'about-me'],
  sectionsColor: ['#FFFFFF', '#FFF5D7', '#FFFFFF', '#FFF5D7'],
  navigation: true,
  navigationPosition: 'left',
  navigationTooltips: ['Home', 'Experience', 'Projects', 'About Me']
});

new TypeIt("#type-carousel", {
  speed: 60,
  deleteSpeed: 30,
  waitUntilVisible: true,
  loop: true,
  cursor: true
})
  .type("Computer Scientist")
  .pause(1500)
  .delete()
  .type("UC Berkeley Student")
  .pause(1500)
  .delete()
  .type("San Francisco Native")
  .pause(1500)
  .delete()
  .type("Twin")
  .pause(1500)
  .delete()
  .type("Boba Enthusiast")
  .pause(1500)
  .delete()
  .go();
