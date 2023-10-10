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

		 p.Nome = nome
		 p.Descricao = descricao
		 p.Status = status

		 produtos = append(produtos, p)

		 
	}
	defer db.Close()
	return produtos

}