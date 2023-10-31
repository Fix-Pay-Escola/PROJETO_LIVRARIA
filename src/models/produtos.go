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

func BuscaTodosOsProdutos() []Produto{
	db := db.ConectacomBancoDeDados()
	
	selectDeTodosOsProdutos, err := db.Query("select * from livros order by id asc")
	if err != nil {
		panic(err.Error())
	}
	
	p := Produto {}
	produtos := []Produto{}

	for selectDeTodosOsProdutos.Next(){
		var id int
		var nome, descricao, status, isbn,autor,editora string

		err = selectDeTodosOsProdutos.Scan(&id, &isbn, &nome, &descricao, &status,&autor,&editora)
		 if err != nil {
			panic(err.Error())
		 }

		 p.Id = id
		 p.Nome = nome
		 p.Descricao = descricao
		 p.Status = status
		 p.Autor = autor


		 produtos = append(produtos, p)

		 
	}
	defer db.Close()
	return produtos

}

func CriarNovoProduto(nome,descricao,status,isbn,autor,editora string){
	db := db.ConectacomBancoDeDados()

	insereDadosNoBanco, err := db.Prepare("insert into livros (nome,descricao,status,isbn,autor,editora) Values ($1,$2,$3,$4,$5,$6)")
	if err != nil {
		panic(err.Error())
	}

	insereDadosNoBanco.Exec(nome,descricao,status,isbn,autor,editora)



	defer db.Close()
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
		var  id int
		var nome,descricao,status,isbn,autor,editora string

		err = produtoDoBanco.Scan(&id,&isbn,&nome,&descricao,&status,&autor,&editora)

		if err != nil{
			panic(err.Error())
		}
		produtoParaAtualizar.Id = id
		produtoParaAtualizar.Nome = nome
		produtoParaAtualizar.Descricao = descricao
		produtoParaAtualizar.Isbn = isbn
		produtoParaAtualizar.Status = status
		produtoParaAtualizar.Autor = autor
		produtoParaAtualizar.Editora = editora
	}
	defer db.Close()
	return produtoParaAtualizar
}
func AtualizaProduto(id int, nome,descricao,status,isbn,autor,editora string){
	db := db.ConectacomBancoDeDados()

	AtualizaProduto, err := db.Prepare("Update livros set nome=$1,descricao=$2,status=$3,isbn=$4,autor=$5,editora=$6 where id=$7")
	if err != nil {
		panic(err.Error())
	}
	AtualizaProduto.Exec(nome,descricao,status,isbn,autor,editora,id)
	defer db.Close()
}