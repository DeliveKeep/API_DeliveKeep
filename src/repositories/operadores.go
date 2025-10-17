package repositories

import (
	"API/src/models"
	"database/sql"
)

// Cria um operador no banco de dados
func CriarOperador(usuario *models.Operador, db *sql.DB) error {
	sqlStatement := `INSERT INTO operadores (nome, telefone, email, senha, galpao) VALUES ($1, $2, $3, $4, $5) RETURNING id_operador`
	if erro := db.QueryRow(sqlStatement, usuario.Nome, usuario.Telefone, usuario.Email, usuario.Senha, usuario.Galpao).Scan(&usuario.Id); erro != nil {
		return erro
	}
	return nil
}
