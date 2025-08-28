package database

import (
	"API/src/config"
	"database/sql"

	_ "github.com/lib/pq"
)

// ConectarDB abre conexão com banco de dados e a retorna
func ConectarDB() (*sql.DB, error) {
	db, erro := sql.Open("postgres", config.StringConexao)
	if erro != nil {
		return nil, erro
	}

	if erro = db.Ping(); erro != nil {
		return nil, erro
	}

	return db, nil
}
