package models

import(
	"golang/src/db"
)
type Produto struct {
	Id int;
	Isbn string;
	Nome string;
	Descricao string;
	Status string
}

func BuscaTodosOsProdutos() []Produto{
	db := db.ConectacomBancoDeDados()
	
	selectDeTodosOsProdutos, err := db.Query("select * from livros")
	if err != nil {
		panic(err.Error())
	}
	
	p := Produto {}
	produtos := []Produto{}

	for selectDeTodosOsProdutos.Next(){
		var id int
		var nome, descricao, status, isbn string

		err = selectDeTodosOsProdutos.Scan(&id, &isbn, &nome, &descricao, &status)
		 if err != nil {
			panic(err.Error())
		 }

		 p.Id = id
		 p.Nome = nome
		 p.Descricao = descricao
		 p.Status = status

		 produtos = append(produtos, p)

		 
	}
	defer db.Close()
	return produtos

}

func CriarNovoProduto(nome,descricao,status,isbn string){
	db := db.ConectacomBancoDeDados()

	insereDadosNoBanco, err := db.Prepare("insert into livros (nome,descricao,status,isbn) Values ($1,$2,$3,$4)")
	if err != nil {
		panic(err.Error())
	}

	insereDadosNoBanco.Exec(nome,descricao,status,isbn)



	defer db.Close()
}

func DeletaProduto(id string){
	db := db.ConectacomBancoDeDados()

	deletarProduto,err := db.Prepare("delete from livros where id=$1")
	 if err != nil {
		panic(err.Error())
	 }

	 deletarProduto.Exec(id)
	 defer db.Close()
}