const check = document.getElementById("check");

check.addEventListener("change", () => {
  document.body.classList.toggle("dark");
  const main = document.querySelector(".card");
  main.classList.toggle("dark");
  const input = document.querySelector(".formulary");
  input.classList.toggle("dark");
  const save = document.querySelector(".btn-save");
  save.classList.toggle("dark");
  const cb = document.querySelector(".btn-back");
  cb.classList.toggle("dark");
});