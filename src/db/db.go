package db

import (
	"database/sql"

	_ "github.com/lib/pq"
)
func ConectacomBancoDeDados() *sql.DB{
	conexao := "user=postgres dbname=Projeto_Livraria_FP password=postgres host=localhost sslmode=disable port=5433"
	db, err := sql.Open("postgres",conexao)	
	if err != nil {
		panic(err.Error())
	}
	return db
}