package models

import(
	_"database/sql"

	_ "github.com/lib/pq"
	"golang/src/db"
)
type Produto struct {
	Id int;
	Isbn string;
	Nome string;
	Descricao string;
	Status string;
	Autor string;
	Editora string;
}


func CriarNovoProduto(nome,descricao,status,isbn,autor string, editora int) error{
	db := db.ConectacomBancoDeDados()
	defer db.Close()
	insereDadosNoBanco, err := db.Prepare("insert into livros (nome,descricao,status,isbn,autor,id_editora) Values ($1,$2,$3,$4,$5,$6)")
	if err != nil {
		return err
	}

	_, err = insereDadosNoBanco.Exec(nome,descricao,status,isbn,autor,editora)
	if err != nil {
		return err
	}

	return nil	
}

func DeletaProduto(id string){
	db := db.ConectacomBancoDeDados()

	deletarOProduto,err := db.Prepare("delete from livros where id=$1")
	 if err != nil {
		panic(err.Error())
	 }

	deletarOProduto.Exec(id)
	 defer db.Close()
}

func EditaProduto(id string) Produto{
	db := db.ConectacomBancoDeDados()

	produtoDoBanco, err := db.Query("select * from livros where id=$1",id)
	if err != nil {
		panic(err.Error())
	}

	produtoParaAtualizar := Produto{}

	for produtoDoBanco.Next() {
		var  id, id_editora int
		var nome,descricao,status,isbn,autor string

		err = produtoDoBanco.Scan(&id,&nome,&descricao,&isbn,&autor,&id_editora,&status)

		if err != nil{
			panic(err.Error())
		}
		produtoParaAtualizar.Id = id
		produtoParaAtualizar.Nome = nome
		produtoParaAtualizar.Descricao = descricao
		produtoParaAtualizar.Isbn = isbn
		produtoParaAtualizar.Status = status
		produtoParaAtualizar.Autor = autor
		if id_editora == 1 {
			produtoParaAtualizar.Editora = "Novatec"
		}else if (id_editora == 2){
			produtoParaAtualizar.Editora = "Alta Books"
		}else if (id_editora == 3){
			produtoParaAtualizar.Editora = "OReilly"
		}else if (id_editora == 4){
			produtoParaAtualizar.Editora = "Casa do Codigo"
		}else if (id_editora == 5){
			produtoParaAtualizar.Editora = "Sextante"
		}
	}
	defer db.Close()
	return produtoParaAtualizar
}
func AtualizaProduto(id int, nome,descricao,status,isbn,autor string,editora int){
	db := db.ConectacomBancoDeDados()

	AtualizaProduto, err := db.Prepare("Update livros set nome=$1,descricao=$2,status=$3,isbn=$4,autor=$5,id_editora=$6 where id=$7")
	if err != nil {
		panic(err.Error())
	}
	AtualizaProduto.Exec(nome,descricao,status,isbn,autor,editora,id)
	defer db.Close()
}
func BuscaProdutosPaginados(page int, itemsPerPage int) ([]Produto, int) {
    db := db.ConectacomBancoDeDados()
    defer db.Close()

   
    offset := (page - 1) * itemsPerPage

   
    query := "SELECT * FROM livros ORDER BY nome ASC LIMIT $1 OFFSET $2"
    rows, err := db.Query(query, itemsPerPage, offset)
    if err != nil {
        panic(err.Error())
    }
    defer rows.Close()
	p := Produto {}
    produtos := []Produto{}

    for rows.Next() {
		var id,id_editora  int
		var nome, descricao, status,isbn,autor string

		err = rows.Scan(&id, &nome, &descricao,&isbn,&autor,&id_editora,&status)
		 if err != nil {
			panic(err.Error())
		 }

		 p.Id = id
		 p.Nome = nome
		 p.Descricao = descricao
		 p.Status = status
		 produtos = append(produtos, p)

    }

    
    var totalItems int
    err = db.QueryRow("SELECT COUNT(*) FROM livros").Scan(&totalItems)
    if err != nil {
        panic(err.Error())
    }

    return produtos, totalItems
}