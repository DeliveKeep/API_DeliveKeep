package repositories

import (
	"API/src/models"
	"database/sql"
	"errors"
)

// Cria um operador no banco de dados
func CriarOperador(usuario *models.Operador, db *sql.DB) error {
	sqlStatement := `INSERT INTO operadores (nome, telefone, email, senha, galpao) VALUES ($1, $2, $3, $4, $5) RETURNING id_operador`
	if erro := db.QueryRow(sqlStatement, usuario.Nome, usuario.Telefone, usuario.Email, usuario.Senha, usuario.Galpao).Scan(&usuario.Id); erro != nil {
		return erro
	}
	return nil
}

// BuscarIdSenhaPorEmail usa um email para buscar Id e senha de um usuário
func BuscarIdESenhaPorEmailOperador(email string, db *sql.DB) (models.Operador, error) {
	sqlStatement := `SELECT id_operador, senha FROM operadores WHERE email=$1`
	var usuario models.Operador
	if erro := db.QueryRow(sqlStatement, email).Scan(&usuario.Id, &usuario.Senha); erro != nil {
		if erro == sql.ErrNoRows {
			return models.Operador{}, errors.New("usuario com esse email nao encontrado")
		}
		return models.Operador{}, erro
	}
	return usuario, nil
}
