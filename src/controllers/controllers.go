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
	"time"
)

var temp = template.Must((template.ParseGlob("_html/*.html")))



func Index(w http.ResponseWriter, r *http.Request) {
	temp.ExecuteTemplate(w, "Index", nil)
}

func Login(w http.ResponseWriter, r *http.Request){
	db := db.ConectacomBancoDeDados()
	if r.Method == http.MethodPost {
		
		email := r.PostFormValue("email")
		password := r.PostFormValue("password")

		userID, authenticated := verifyCredentials(email, password)

    if authenticated {
		expiration := time.Now().Add(1 * time.Hour)
        cookie := http.Cookie{
            Name:  "userID",
            Value: strconv.Itoa(userID),
			Expires: expiration,
        }
        http.SetCookie(w, &cookie)

       
        http.Redirect(w, r, "/acervo_adm", http.StatusFound)
    } else {
        
        http.Redirect(w,r,"/login?error=true",http.StatusNotFound)
    }
		
	defer db.Close()

	
}
temp.ExecuteTemplate(w,"Login",nil)
}
func About_us(w http.ResponseWriter, r*http.Request){
	temp.ExecuteTemplate(w,"AboutUs",nil)
}
func Alugueis(w http.ResponseWriter, r *http.Request) {
	if !isAuthenticated(r) {
       
        http.Redirect(w, r, "/login", http.StatusFound)
        return
    }
   
    page := 1 
    itemsPerPage := 13 

   
    pageStr := r.URL.Query().Get("page")
    if pageStr != "" {
        page, _ = strconv.Atoi(pageStr)
    }

   
    produtos, totalItems := models.BuscaProdutosPaginados(page, itemsPerPage)

  
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
		PreviousPage int 
		NextPage int
	}{
		Produtos:    produtos,
		TotalPages:  totalPages,
		CurrentPage: page,
		PreviousPage: previousPage,
		NextPage: nextPage, 
	}
    temp.ExecuteTemplate(w, "Alugueis", templateData)
}
func Alugueis_user(w http.ResponseWriter, r*http.Request){
	
    page := 1 
    itemsPerPage := 13

   
    pageStr := r.URL.Query().Get("page")
    if pageStr != "" {
        page, _ = strconv.Atoi(pageStr)
    }

    
    produtos, totalItems := models.BuscaProdutosPaginados(page, itemsPerPage)


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
		PreviousPage int
		NextPage int
	}{
		Produtos:    produtos,
		TotalPages:  totalPages,
		CurrentPage: page,
		PreviousPage: previousPage,
		NextPage: nextPage, 
	}
	temp.ExecuteTemplate(w,"Alugueis_User",templateData)
}
func Edit(w http.ResponseWriter, r*http.Request){
	if !isAuthenticated(r) {
       
        http.Redirect(w, r, "/login", http.StatusFound)
        return
    }
	
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
	if !isAuthenticated(r) {
       
        http.Redirect(w, r, "/login", http.StatusFound)
        return
    }
	temp.ExecuteTemplate(w,"New",nil)
}

func Insert(w http.ResponseWriter, r*http.Request){
	if !isAuthenticated(r) {
       
        http.Redirect(w, r, "/login", http.StatusFound)
        return
    }
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
		} else if editora == "Sextante"{
			id_editora = 5
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
	if !isAuthenticated(r) {
       
        http.Redirect(w, r, "/login", http.StatusFound)
        return
    }
	idDoProduto := r.URL.Query().Get("id")


	models.DeletaProduto(idDoProduto)
	http.Redirect(w,r,"/acervo_adm",301)
}
func Update(w http.ResponseWriter, r *http.Request){
	if !isAuthenticated(r) {
       
        http.Redirect(w, r, "/login", http.StatusFound)
        return
    }
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
		}else if editora == "Sextante"{
			id_editora = 5
		}
	

	idConvertido, err := strconv.Atoi(id)
	 if err != nil {
		log.Println("Erro na Conversão do ID")
	 }

	 models.AtualizaProduto(idConvertido,nome,descricao,status,isbn,autor,id_editora)

	 http.Redirect(w,r,"/acervo_adm",301)
	}
}

func verifyCredentials(email, password string) (int,bool) {

	db := db.ConectacomBancoDeDados()
	defer db.Close()
    
    var storedPassword string
	var userID int
	err := db.QueryRow("SELECT id, password FROM users WHERE email = $1", email).Scan(&userID, &storedPassword)
    if err != nil {
       
        return 0, false
    }


   
    if password == storedPassword {
       
        return userID,true
    }

    // Senha incorreta
    return 0,false

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

func isAuthenticated(r *http.Request) bool {
   
    _, err := r.Cookie("userID")
    if err != nil {
        return false
    }
    return true
}