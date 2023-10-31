package controllers

import (
	"golang/src/db"
	"golang/src/models"
	"html/template"
	"log"
	"net/http"
	"strconv"
)

var temp = template.Must((template.ParseGlob("_html/*.html")))




func Index(w http.ResponseWriter, r *http.Request) {
	temp.ExecuteTemplate(w, "Index", nil)
}

func Login(w http.ResponseWriter, r *http.Request){
	db := db.ConectacomBancoDeDados()
	if r.Method == http.MethodPost {
		// Obtenha os dados do formulário (e-mail e senha)
		email := r.PostFormValue("email")
		password := r.PostFormValue("password")

		// Consulte o banco de dados para verificar as credenciais
		if verifyCredentials(email, password) {
			http.Redirect(w, r, "/Acervo_Adm", http.StatusSeeOther)
			return
		}else{
			http.Redirect(w, r, "/login?error=true", http.StatusSeeOther)
		}

	defer db.Close()

	
}
temp.ExecuteTemplate(w,"Login",nil)
}
func About_us(w http.ResponseWriter, r*http.Request){
	temp.ExecuteTemplate(w,"AboutUs",nil)
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

	idDoProduto := r.URL.Query().Get("id")

	produto := models.EditaProduto(idDoProduto)
	temp.ExecuteTemplate(w,"Edit",produto)
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
		autor := r.FormValue("autor")
		editora := r.FormValue("editora")

		models.CriarNovoProduto(nome,descricao,status,isbn,autor,editora)
	}
	http.Redirect(w,r,"/Acervo_Adm",301)
}

func Delete(w http.ResponseWriter, r *http.Request){
	idDoProduto := r.URL.Query().Get("id")


	models.DeletaProduto(idDoProduto)
	http.Redirect(w,r,"/Acervo_Adm",301)
}
func Update(w http.ResponseWriter, r *http.Request){
	if r.Method == "POST"{
		id := r.FormValue("id")
		nome := r.FormValue("nome")
		descricao := r.FormValue("descricao")
		status := r.FormValue("status")
		isbn := r.FormValue("isbn")
		autor := r.FormValue("autor")
		editora := r.FormValue("editora")
	

	idConvertido, err := strconv.Atoi(id)
	 if err != nil {
		log.Println("Erro na Conversão do ID")
	 }

	 models.AtualizaProduto(idConvertido,nome,descricao,status,isbn,autor,editora)

	 http.Redirect(w,r,"/Acervo_Adm",301)
	}
}

func verifyCredentials(email, password string) bool {

	db := db.ConectacomBancoDeDados()
	defer db.Close()
    // Consulte o banco de dados para obter as credenciais do usuário
    var storedPassword string
    err := db.QueryRow("SELECT password FROM users WHERE email = $1", email).Scan(&storedPassword)
    if err != nil {
        // Usuário não encontrado
        return false
    }

    // Verifique se a senha fornecida corresponde à senha armazenada
    // Neste exemplo, as senhas não são salvas com hash, mas isso é altamente recomendado na prática.
    if password == storedPassword {
        // As credenciais são corretas
        return true
    }

    // Senha incorreta
    return false

}