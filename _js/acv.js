const check = document.getElementById("check");

check.addEventListener("change", () => {
  document.body.classList.toggle("dark");
  const main = document.querySelector("#btn-add");
  main.classList.toggle("dark");
  const botao = document.querySelector("#head");
  botao.classList.toggle("dark");
  const input = document.querySelector("#btn-del");
  input.classList.toggle("dark");
  const table = document.querySelector(".table");
  table.classList.toggle("dark");
  const backgBtn = document.querySelector(".cell");
  backgBtn.classList.toggle("dark");
});
function openModal(){
  const clicked = document.getElementById('janela_modal')
  clicked.classList.add('open')
  clicked.addEventListener('click', (e) =>{
      if(e.target.id == 'close' || e.target.id == 'janela_modal'){
          clicked.classList.remove('open')
      }
  })
}