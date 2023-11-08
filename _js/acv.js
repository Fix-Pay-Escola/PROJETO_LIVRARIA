const check = document.getElementById("check")
check.addEventListener("change", () => {
  document.body.classList.toggle("dark")
  const main = document.querySelector("#btn-add")
  main.classList.toggle("dark")
  const botao = document.querySelector("#head");
  botao.classList.toggle("dark")
  const input = document.querySelector("#btn-del")
  input.classList.toggle("dark")
  const table = document.querySelector(".table");
  table.classList.toggle("dark")
  const backgBtn = document.querySelector(".cell")
  backgBtn.classList.toggle("dark")
  const description = document.querySelector(".truncate-text");
  description.classList.toggle("dark");
})

function expandText() {
  document.addEventListener('click', function (e) {
    if (e.target.classList.contains('truncate-text')) {
      const td = e.target;
      if (td.getAttribute('data-expanded') === 'true') {
        td.style.whiteSpace = 'nowrap';
        td.style.overflow = 'hidden';
        td.style.height = '1.2em';
        td.setAttribute('data-expanded', 'false');
      } else {
        td.style.whiteSpace = 'normal';
        td.style.overflow = 'visible';
        td.style.height = 'auto';
        td.setAttribute('data-expanded', 'true');
      }
    }
  });
}

document.getElementById("search-button").addEventListener("click", function () {
  var criterio = document.getElementById("search-input").value;
  window.location = `/search?criterio=${criterio}`;
});
