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
	http.HandleFunc("/about_us",controllers.About_us)
	http.HandleFunc("/new",controllers.New)
	http.HandleFunc("/insert",controllers.Insert)
	http.HandleFunc("/delete",controllers.Delete)
	http.HandleFunc("/edit",controllers.Edit)
	http.HandleFunc("/update",controllers.Update)
	http.HandleFunc("/Acervo_Adm",controllers.Alugueis)
	http.HandleFunc("/Acervo",controllers.Alugueis_user)
}
