package main

import (
	"log"
	"net/http"

	"golang-web-app/config"
	"golang-web-app/handlers"
)

func main() {
	config.ConnectDB() // conecta com o banco

	// Arquivos estáticos
	fs := http.FileServer(http.Dir("public"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	// Rotas
	http.HandleFunc("/", handlers.HomeHandler)
	http.HandleFunc("/login", handlers.LoginHandler)

	log.Println("Servidor rodando em http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
