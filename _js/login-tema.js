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

goin.addEventListener("click", function handleClick() {
  const loginInput = document.getElementById("login");
  const passInput = document.getElementById("pass");
  const login = loginInput.value;
  const pass = passInput.value;

  if (login === "imthepontes@gmail.com" && pass === "123456789") {
    loginInput.value = "";
    passInput.value = "";
    window.location = "/Acervo_Adm"
  } else {
    alert("Usuário ou senha incorretos!");
    loginInput.value = "";
    passInput.value = "";
  }
});
