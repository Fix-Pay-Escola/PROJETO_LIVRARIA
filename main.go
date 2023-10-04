package main

import (
	"net/http"
	_ "github.com/lib/pq"
	"html/template"

)
var temp = template.Must((template.ParseGlob("HTML/*.html")))

func main(){
	http.HandleFunc("/", index)
	http.ListenAndServe(":5500",nil)
	
}

func index(w http.ResponseWriter, r* http.Request){
 temp.ExecuteTemplate(w,"Index",nil)
}