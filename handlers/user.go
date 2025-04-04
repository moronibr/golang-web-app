package handlers

import (
	"golang-web-app/models"
	"net/http"
)

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		r.ParseForm()
		email := r.FormValue("email")
		password := r.FormValue("password")

		user, err := models.AuthenticateUser(email, password)
		if err != nil {
			http.Error(w, "Credenciais inválidas", http.StatusUnauthorized)
			return
		}

		// autenticação bem-sucedida
		http.Redirect(w, r, "/", http.StatusSeeOther)
		_ = user // usar sessão depois
	} else {
		http.ServeFile(w, r, "templates/login.html")
	}
}
