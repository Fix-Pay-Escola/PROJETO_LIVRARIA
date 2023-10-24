package routes
import (
	"net/http"
)
func Carregar_arquivos(){
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


	 http.HandleFunc("/script/js/login-script",func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w,r,"_js/login-tema.js")
	 })


	 http.HandleFunc("/styles/alugueis",func(w http.ResponseWriter, r*http.Request){
		http.ServeFile(w,r,"_css/style-alug.css")
	 })
	 http.HandleFunc("/styles/image/home_image", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w,r,"_imagens/pilha-livro.jpg")
	 })
	 http.HandleFunc("/styles/adm_css", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w,r,"_css/style-alug.css")
	 })
	 http.HandleFunc("/style/index-icon",func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w,r,"_imagens/biblioteca-fixpay-website-favicon-color_03_.ico")
	 })

	 http.HandleFunc("/script/index",func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w,r,"_js/index.js")
	})

	http.HandleFunc("/script/js/dark-mode",func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w,r,"_js/login-tema.js")
	 })

	http.HandleFunc("/_js/acv.js",func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w,r,"_js/acv.js")
	})
	http.HandleFunc("/_css/style-ac-adm.css",func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w,r,"_css/style-ac-adm.css")
	})
	http.HandleFunc("/_imagens/FPL3.ico",func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w,r,"_imagens/FPL3.ico")
	})
	http.HandleFunc("/_css/style-new.css",func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w,r,"_css/style-new.css")
	})
	http.HandleFunc("/_imagens/amin-acv.svg",func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w,r,"_imagens/amin-acv.svg")
	})
	http.HandleFunc("/_js/new-add.js",func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w,r,"_js/new-add.js")
	})
	http.HandleFunc("/_css/style-abt.css",func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w,r,"_css/style-abt.css")
	})
}