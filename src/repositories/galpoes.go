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

// Busca todos os galpões
func BuscarGalpoes(db *sql.DB) ([]models.Galpao, error) {
	sqlStatement := `SELECT id_galpao, nome, endereco FROM galpoes`
	rows, err := db.Query(sqlStatement)
	if err != nil {
		return []models.Galpao{}, err
	}
	defer rows.Close()
	var galpoes []models.Galpao
	// Itera sobre as linhas retornadas
	for rows.Next() {
		var galpao models.Galpao
		if err := rows.Scan(&galpao.Id, &galpao.Nome, &galpao.Endereco); err != nil {
			return []models.Galpao{}, err
		}
		galpoes = append(galpoes, galpao)
	}
	// Verifica se ocorreu algum erro durante a iteração
	if err = rows.Err(); err != nil {
		return []models.Galpao{}, err
	}
	return galpoes, nil
}
