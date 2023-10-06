package main

import (
	"fmt"
	"html/template"
	"net/http"

	_ "github.com/lib/pq"
)

var temp = template.Must((template.ParseGlob("_html/*.html")))

func main() {
	
	carregar_arquivos()
	
	http.HandleFunc("/", index)
	http.HandleFunc("/login", login)
	http.HandleFunc("/about-us",about_us)

//	fs := http.FileServer(http.Dir("_css/"))
//	http.HandleFunc("/_css/", func(w http.ResponseWriter, r *http.Request) {
//		if strings.HasSuffix(r.URL.String(), ".css") {
//			w.Header().Set("Content-Type", "text/css")
//		}
//		http.StripPrefix("/_css/", fs)
//	})

	http.ListenAndServe(":5500", nil)

	fmt.Println("aiaiaiai")
}

func index(w http.ResponseWriter, r *http.Request) {
	temp.ExecuteTemplate(w, "Index", nil)
}

func login(w http.ResponseWriter, r *http.Request){
	temp.ExecuteTemplate(w,"Login",nil)
}
func about_us(w http.ResponseWriter, r*http.Request){
	temp.ExecuteTemplate(w,"About Us",nil)
}
func carregar_arquivos(){
	http.HandleFunc("/styles/home", func (w http.ResponseWriter, r*http.Request)  {
		http.ServeFile(w,r,"_css/style-index.css")
	})

	http.HandleFunc("/styles/image/home_icon", func (w http.ResponseWriter, r*http.Request)  {
		http.ServeFile(w,r,"_imagens/biblioteca-fixpay-website-favicon-color_04_.ico")
	})

	http.HandleFunc("/styles/login", func (w http.ResponseWriter, r*http.Request)  {
		http.ServeFile(w,r,"_css/style-login-adm.css")
	})
	
	 http.HandleFunc("/styles/image/adm_image",func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w,r,"_imagens/admin-pc.svg")
	 })

	 http.HandleFunc("/script/js/dark-mode",func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w,r,"_js/login-tema.js")
	 })
}