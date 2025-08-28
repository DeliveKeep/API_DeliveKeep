package repositories

import (
	"API/src/models"
	"database/sql"
	"errors"
)

// CriarUsuario insere um novo usuario no banco de dados
func CriarUsuario(usuario *models.Usuario, db *sql.DB) error {
	sqlStatement := `INSERT INTO usuarios (nome, cpf, endereco, telefone, email, senha) VALUES ($1, $2, $3, $4, $5, $6) RETURNING id`
	if erro := db.QueryRow(sqlStatement, usuario.Nome, usuario.Cpf, usuario.Endereco, usuario.Telefone, usuario.Email, usuario.Senha).Scan(&usuario.Id); erro != nil {
		return erro
	}
	return nil
}

// BuscarIdSenhaPorEmail usa um email para buscar Id e senha de um usuário
func BuscarIdESenhaPorEmail(email string, db *sql.DB) (models.Usuario, error) {
	sqlStatement := `SELECT id, senha FROM usuarios WHERE email=$1`
	var usuario models.Usuario
	if erro := db.QueryRow(sqlStatement, email).Scan(&usuario.Id, &usuario.Senha); erro != nil {
		if erro == sql.ErrNoRows {
			return models.Usuario{}, errors.New("usuario com esse email nao encontrado")
		}
		return models.Usuario{}, erro
	}
	return usuario, nil
}

// BuscarLogado busca dados exceto a senha de um usuário pela id
func BuscarLogado(id int, db *sql.DB) (models.Usuario, error) {
	sqlStatement := `SELECT id, nome, cpf, endereco, telefone, email FROM usuarios WHERE id=$1`
	var usuario models.Usuario
	if erro := db.QueryRow(sqlStatement, id).Scan(&usuario.Id, &usuario.Nome, &usuario.Cpf, &usuario.Endereco, &usuario.Telefone, &usuario.Email); erro != nil {
		if erro == sql.ErrNoRows {
			return models.Usuario{}, errors.New("Id nao encontrado")
		}
		return models.Usuario{}, erro
	}
	return usuario, nil
}
