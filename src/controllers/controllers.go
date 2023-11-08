package controllers

import (
	"fmt"
	"golang/src/db"
	"golang/src/models"
	"html/template"
	"log"
	"net/http"
	"strconv"
	"math"
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
			http.Redirect(w, r, "/acervo_adm", http.StatusSeeOther)
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
func Alugueis(w http.ResponseWriter, r *http.Request) {
    // Parâmetros para paginação
    page := 1 // Página atual
    itemsPerPage := 10 // Itens por página

    // Obtenha o valor da página da solicitação
    pageStr := r.URL.Query().Get("page")
    if pageStr != "" {
        page, _ = strconv.Atoi(pageStr)
    }

    // Obtenha a lista de produtos com base na página atual e itens por página
    produtos, totalItems := models.BuscaProdutosPaginados(page, itemsPerPage)

    // Calcule o número total de páginas
    totalPages := int(math.Ceil(float64(totalItems) / float64(itemsPerPage)))
	var previousPage int
	if page > 1 {
		previousPage = page - 1
	}
	var nextPage int
	if page < totalPages {
    nextPage = page + 1
	}
	
	templateData := struct {
		Produtos    []models.Produto
		TotalPages  int
		CurrentPage int
		PreviousPage int // Adicione o valor da página anterior
		NextPage int
	}{
		Produtos:    produtos,
		TotalPages:  totalPages,
		CurrentPage: page,
		PreviousPage: previousPage,
		NextPage: nextPage, // Passe o valor da página anterior
	}
    temp.ExecuteTemplate(w, "Alugueis", templateData)
}
func Alugueis_user(w http.ResponseWriter, r*http.Request){
	todososprodutos := models.BuscaTodosOsProdutos()
	temp.ExecuteTemplate(w,"Alugueis_User",todososprodutos)
}
func Edit(w http.ResponseWriter, r*http.Request){

	
	idDoProduto := r.URL.Query().Get("id")

	idint,err := strconv.Atoi(idDoProduto)
	if err != nil {
		fmt.Print("Erro na conversao do ID")
	}
	exists, err := Checando_Existencia(idint)
	if err != nil {
		http.Error(w, "Erro ao verificar o banco de dados", http.StatusInternalServerError)
		return
	}

	if !exists {
		handleNotFound(w, r)
		return
	}

	

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
		id_editora := 0
		if editora == "Novatec" {
			id_editora = 1
		}else if editora == "Alta Books"{
			id_editora = 2
		}else if editora == "O'Reilly"{
			id_editora = 3
		} else if editora == "Casa do Código"{
			id_editora = 4
		}

		err := models.CriarNovoProduto(nome,descricao,status,isbn,autor,id_editora)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
            return
		}
	}
	http.Redirect(w,r,"/acervo_adm",301)
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
		id_editora := 0

		if editora == "Novatec"{
			id_editora = 1
		}else if editora == "Alta Books"{
			id_editora = 2
		}else if editora == "OReilly"{
			id_editora = 3
		}else if editora == "Casa do Codigo"{
			id_editora = 4
		}
	

	idConvertido, err := strconv.Atoi(id)
	 if err != nil {
		log.Println("Erro na Conversão do ID")
	 }

	 models.AtualizaProduto(idConvertido,nome,descricao,status,isbn,autor,id_editora)

	 http.Redirect(w,r,"/acervo_adm",301)
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

func Checando_Existencia(id int) (bool,error){
	db := db.ConectacomBancoDeDados()

	defer db.Close()

	var exists bool
	err := db.QueryRow("SELECT EXISTS(SELECT 1 FROM livros WHERE id=$1)", id).Scan(&exists)
	if err != nil {
		return false, err
	}
	return exists, nil
}

func Error(w http.ResponseWriter, r *http.Request){

	temp.ExecuteTemplate(w,"error",nil)
}

func handleNotFound(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotFound)
	fmt.Fprint(w, "ERROR 404:Página não encontrada")
}