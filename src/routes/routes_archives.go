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
}