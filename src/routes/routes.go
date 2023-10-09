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
func Alugueis_user(w http.ResponseWriter, r*http.Request){
	temp.ExecuteTemplate(w,"Alugueis_User",nil)
}
func Edit(w http.ResponseWriter, r*http.Request){
	temp.ExecuteTemplate(w,"Edit",nil)
}
func New(w http.ResponseWriter, r*http.Request){
	temp.ExecuteTemplate(w,"New",nil)
}
func Delete(w http.ResponseWriter, r*http.Request){
	http.Redirect(w,r,"/",301)
}