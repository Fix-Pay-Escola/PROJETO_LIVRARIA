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