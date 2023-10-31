package routes

import(
	"net/http"
	"golang/src/controllers"
	_"golang/src/models"
	_ "github.com/lib/pq"
)
func Carregar_Rotas(){
	http.HandleFunc("/", controllers.Index)
	http.HandleFunc("/login", controllers.Login)
	http.HandleFunc("/sobre_nos",controllers.About_us)
	http.HandleFunc("/new",controllers.New)
	http.HandleFunc("/insert",controllers.Insert)
	http.HandleFunc("/delete",controllers.Delete)
	http.HandleFunc("/edit",controllers.Edit)
	http.HandleFunc("/update",controllers.Update)
	http.HandleFunc("/acervo_adm",controllers.Alugueis)
	http.HandleFunc("/acervo",controllers.Alugueis_user)
	http.HandleFunc("/404",controllers.Error)
}
