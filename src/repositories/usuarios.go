package repositories

import (
	"API/src/models"
	"database/sql"
)

// CriarUsuario insere um novo usuario no banco de dados
func CriarUsuario(usuario *models.Usuario, db *sql.DB) error {
	sqlStatement := `INSERT INTO usuarios (nome, cpf, endereco, telefone, email, senha) VALUES ($1, $2, $3, $4, $5, $6) RETURNING id`
	if erro := db.QueryRow(sqlStatement, usuario.Nome, usuario.Cpf, usuario.Endereco, usuario.Telefone, usuario.Email, usuario.Senha).Scan(&usuario.Id); erro != nil {
		return erro
	}
	return nil
}
