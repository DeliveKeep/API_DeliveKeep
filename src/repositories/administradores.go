package repositories

import (
	"API/src/models"
	"database/sql"
	"errors"
)

// Cria um administrador no banco de dados
func CriarAdministrador(usuario *models.Administrador, db *sql.DB) error {
	sqlStatement := `INSERT INTO administradores (nome, telefone, email, senha, galpao) VALUES ($1, $2, $3, $4, $5) RETURNING id_administrador`
	if erro := db.QueryRow(sqlStatement, usuario.Nome, usuario.Telefone, usuario.Email, usuario.Senha, usuario.Galpao).Scan(&usuario.Id); erro != nil {
		return erro
	}
	return nil
}

// BuscarIdSenhaPorEmail usa um email para buscar Id e senha de um usuário
func BuscarIdESenhaPorEmailAdministrador(email string, db *sql.DB) (models.Administrador, error) {
	sqlStatement := `SELECT id_administrador, senha FROM administradores WHERE email=$1`
	var usuario models.Administrador
	if erro := db.QueryRow(sqlStatement, email).Scan(&usuario.Id, &usuario.Senha); erro != nil {
		if erro == sql.ErrNoRows {
			return models.Administrador{}, errors.New("usuario com esse email nao encontrado")
		}
		return models.Administrador{}, erro
	}
	return usuario, nil
}
