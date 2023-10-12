package controllers

import (
	"html/template"
	"net/http"
	"golang/src/models"
)

var temp = template.Must((template.ParseGlob("_html/*.html")))




func Index(w http.ResponseWriter, r *http.Request) {
	temp.ExecuteTemplate(w, "Index", nil)
}

func Login(w http.ResponseWriter, r *http.Request){
	temp.ExecuteTemplate(w,"Login",nil)
}
func About_us(w http.ResponseWriter, r*http.Request){
	temp.ExecuteTemplate(w,"About Us",nil)
}
func Alugueis(w http.ResponseWriter, r*http.Request){
	todososprodutos := models.BuscaTodosOsProdutos()
	temp.ExecuteTemplate(w,"Alugueis",todososprodutos)

}
func Alugueis_user(w http.ResponseWriter, r*http.Request){
	todososprodutos := models.BuscaTodosOsProdutos()
	temp.ExecuteTemplate(w,"Alugueis_User",todososprodutos)
}
func Edit(w http.ResponseWriter, r*http.Request){
	temp.ExecuteTemplate(w,"Edit",nil)
}
func New(w http.ResponseWriter, r*http.Request){
	temp.ExecuteTemplate(w,"New",nil)
}

func Insert(w http.ResponseWriter, r*http.Request){
	if r.Method == "POST"{
		nome := r.FormValue("nome")
		descricao := r.FormValue("descricao")
		status := r.FormValue("status")
		isbn := r.FormValue("isbn")

		models.CriarNovoProduto(nome,descricao,status,isbn)
	}
	http.Redirect(w,r,"/Acervo_Adm",301)
}

func Delete(w http.ResponseWriter, r*http.Request){
	idDoProduto := r.URL.Query().Get("id")
	models.DeletaProduto(idDoProduto)
	http.Redirect(w,r,"/Acervo_Adm",301)
}