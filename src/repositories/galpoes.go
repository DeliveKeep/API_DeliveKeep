package repositories

import (
	"API/src/models"
	"database/sql"
)

// Cria um galpão
func CriarGalpao(galpao *models.Galpao, db *sql.DB) error {
	sqlStatement := `INSERT INTO galpoes (nome, endereco) VALUES ($1, $2) RETURNING id_galpao`
	if erro := db.QueryRow(sqlStatement, galpao.Nome, galpao.Endereco).Scan(&galpao.Id); erro != nil {
		return erro
	}
	return nil
}
