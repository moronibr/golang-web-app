package config

import (
	"log"

	"github.com/joho/godotenv"
)

// LoadEnv carrega variáveis de ambiente do arquivo .env
func LoadEnv() {
	err := godotenv.Load()
	if err != nil {
		log.Println("Nenhum arquivo .env encontrado, usando variáveis do sistema.")
	}
}
