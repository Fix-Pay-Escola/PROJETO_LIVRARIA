const check = document.getElementById("check");

check.addEventListener("change", () => {
  document.body.classList.toggle("dark");
  const main = document.querySelector(".main");
  main.classList.toggle("dark");
  const botao = document.querySelector(".boton");
  botao.classList.toggle("dark");
  const input = document.querySelector(".cont");
  input.classList.toggle("dark");
  const butCancel = document.querySelector(".cancelbtn");
  butCancel.classList.toggle("dark");
});

const goin = document.getElementById("goin");
