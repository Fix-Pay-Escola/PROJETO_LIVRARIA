package db

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)
func ConectacomBancoDeDados() *sql.DB{
	conexao := "user=postgres dbname=Projeto_Livraria_FP password=postgres host=localhost sslmode=disable port=5439"
	db, err := sql.Open("postgres",conexao)	
	if err != nil {
		log.Fatal(err)
	}
	return db
}