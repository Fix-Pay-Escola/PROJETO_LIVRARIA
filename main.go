package main

import (
	"golang/src/db"
	"golang/src/routes"
	"net/http"

	_ "github.com/lib/pq"
)



func main() {
	db.ConectacomBancoDeDados()
	routes.Carregar_arquivos()
	routes.Carregar_Rotas()

	http.ListenAndServe(":5500", nil)
}