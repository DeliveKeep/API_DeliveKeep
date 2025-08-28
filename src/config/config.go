package config

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

// declarando váriaveis globais de ambiente
var (
	StringConexao string
	PortaAPI      int
	ChaveSecreta  []byte
)

type contextKey string

const MatriculaKey contextKey = "matricula"

// Carregar inicializa as variaveis de ambiente
func Carregar() {
	var erro error

	// ao usar docker não precisa usar godotenv
	if erro = godotenv.Load("../.env"); erro != nil {
		log.Fatal(erro)
	}

	PortaAPI, erro = strconv.Atoi(os.Getenv("API_PORT"))
	if erro != nil {
		PortaAPI = 5000
	}

	ChaveSecreta = []byte(os.Getenv("SECRET_KEY"))

	StringConexao = fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", os.Getenv("DB_HOST"), os.Getenv("DB_PORT"), os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_NAME"))
}
