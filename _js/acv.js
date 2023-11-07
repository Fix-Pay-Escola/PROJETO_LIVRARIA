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
})

function expandText() {
  document.addEventListener('click', function (e) {
    if (e.target.classList.contains('truncate-text')) {
      const td = e.target;
      const tr = td.parentElement;

      if (td.getAttribute('data-expanded') === 'true') {
        td.style.whiteSpace = 'nowrap'
        td.style.overflow = 'hidden'
        td.style.textOverflow = 'ellipsis'
        td.style.height = '1.2em'
        tr.style.verticalAlign = 'top'
        td.setAttribute('data-expanded', 'false')
      } else {
        td.style.whiteSpace = 'normal'
        td.style.overflow = 'visible'
        td.style.height = 'auto'
        tr.style.verticalAlign = 'middle'
        td.setAttribute('data-expanded', 'true')
      }
    }
  });
}
