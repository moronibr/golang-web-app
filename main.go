package main

import (
	"html/template"
	"log"
	"net/http"
	"os"

	"golang-web-app/config"
)

func main() {
	config.LoadEnv()
	config.ConnectDB()

	// Rota para arquivos estáticos
	fs := http.FileServer(http.Dir("./static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	// Página principal - renderiza o template
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		tmpl, err := template.ParseFiles("templates/login.html")
		if err != nil {
			http.Error(w, "Erro ao carregar o template", http.StatusInternalServerError)
			return
		}
		tmpl.Execute(w, nil)
	})

	port := os.Getenv("PORT")
	if port == "" {
		port = "8000"
	}
	log.Println("Servidor rodando na porta " + port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
