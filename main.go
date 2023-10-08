package main

import (
	"net/http"
	"golang/src/routes"
	_ "github.com/lib/pq"
)



func main() {
	
	routes.Carregar_arquivos()
	
	http.HandleFunc("/", routes.Index)
	http.HandleFunc("/login", routes.Login)
	http.HandleFunc("/about-us",routes.About_us)
	http.HandleFunc("/Acervo",routes.Alugueis)

	http.ListenAndServe(":5500", nil)
}