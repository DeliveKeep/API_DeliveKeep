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

// Busca dados de todos os operadores
func BuscarOperadores(db *sql.DB) ([]models.Operador, error) {
	sqlStatement := `SELECT id_operador, nome, telefone, email, galpao FROM operadores`
	rows, err := db.Query(sqlStatement)
	if err != nil {
		return []models.Operador{}, err
	}
	defer rows.Close()
	var usuarios []models.Operador
	// Itera sobre as linhas retornadas
	for rows.Next() {
		var usuario models.Operador
		if err := rows.Scan(&usuario.Id, &usuario.Nome, &usuario.Telefone, &usuario.Email, &usuario.Galpao); err != nil {
			return []models.Operador{}, err
		}
		usuarios = append(usuarios, usuario)
	}
	// Verifica se ocorreu algum erro durante a iteração
	if err = rows.Err(); err != nil {
		return []models.Operador{}, err
	}
	return usuarios, nil
}

// BuscarLogado busca dados exceto a senha de um usuário pela id
func BuscarOperadorLogado(id int, db *sql.DB) (models.Operador, error) {
	sqlStatement := `SELECT id_operador, nome, telefone, email, galpao FROM clientes WHERE id_operador=$1`
	var usuario models.Operador
	if erro := db.QueryRow(sqlStatement, id).Scan(&usuario.Id, &usuario.Nome, &usuario.Telefone, &usuario.Email, &usuario.Galpao); erro != nil {
		if erro == sql.ErrNoRows {
			return models.Operador{}, errors.New("Id nao encontrado")
		}
		return models.Operador{}, erro
	}
	return usuario, nil
}
