package repositories

import (
	"API/src/models"
	"database/sql"
)

// Cria um administrador no banco de dados
func CriarAdministrador(usuario *models.Administrador, db *sql.DB) error {
	sqlStatement := `INSERT INTO administradores (nome, telefone, email, senha, galpao) VALUES ($1, $2, $3, $4, $5) RETURNING id_administrador`
	if erro := db.QueryRow(sqlStatement, usuario.Nome, usuario.Telefone, usuario.Email, usuario.Senha, usuario.Galpao).Scan(&usuario.Id); erro != nil {
		return erro
	}
	return nil
}
