package main

import (
	"fmt"
	"net/http"

	"golang-web-app/config"
)

func main() {
	config.ConnectDB()

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Servidor funcionando na porta 8000 🚀")
	})

	fmt.Println("🌐 Servidor iniciado em http://localhost:8000")
	http.ListenAndServe(":8000", nil)
}
