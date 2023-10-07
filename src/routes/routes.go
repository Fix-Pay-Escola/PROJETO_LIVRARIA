package routes

import (
	"net/http"
	"html/template"
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
	temp.ExecuteTemplate(w,"Alugueis",nil)
}