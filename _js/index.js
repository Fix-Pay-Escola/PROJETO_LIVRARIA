const check = document.getElementById("check");

check.addEventListener("change", () => {
  document.body.classList.toggle("dark");
  const main = document.querySelector(".line");
  main.classList.toggle("dark");
  const botao = document.querySelector(".back");
  botao.classList.toggle("dark");
  const input = document.querySelector("#login");
  input.classList.toggle("dark");
  const butCancel = document.querySelector(".cancelbtn");
  butCancel.classList.toggle("dark");
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
