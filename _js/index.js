const check = document.getElementById("check");

check.addEventListener("change", () => {
  document.body.classList.toggle("dark");
  const header = document.querySelector(".back");
  header.classList.toggle("dark");
  const headerEnd = document.querySelector(".line");
  headerEnd.classList.toggle("dark");
  const login = document.querySelector("#login");
  login.classList.toggle("dark");
  const topTxt = document.querySelector(".hover-trigger");
  topTxt.classList.toggle("dark");
  const blocoCap = document.querySelector(".acervi");
  blocoCap.classList.toggle("dark");
});
document.addEventListener('DOMContentLoaded', function() {
  var hoverTriggers = document.querySelectorAll('.hover-trigger');
  var casa = document.getElementById('casa');
  function handleMouseEnter() {
    casa.style.color = 'black';
  }
  function handleMouseLeave() {
    casa.style.color = '#0a7e90';
  }
  hoverTriggers.forEach(function(element) {
    element.addEventListener('mouseenter', handleMouseEnter);
    element.addEventListener('mouseleave', handleMouseLeave);
  });
});