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

// Busca dados de todos os administradores
func BuscarAdministradores(db *sql.DB) ([]models.Administrador, error) {
	sqlStatement := `SELECT id_administrador, nome, telefone, email, galpao FROM administradores`
	rows, err := db.Query(sqlStatement)
	if err != nil {
		return []models.Administrador{}, err
	}
	defer rows.Close()
	var usuarios []models.Administrador
	// Itera sobre as linhas retornadas
	for rows.Next() {
		var usuario models.Administrador
		if err := rows.Scan(&usuario.Id, &usuario.Nome, &usuario.Telefone, &usuario.Email, &usuario.Galpao); err != nil {
			return []models.Administrador{}, err
		}
		usuarios = append(usuarios, usuario)
	}
	// Verifica se ocorreu algum erro durante a iteração
	if err = rows.Err(); err != nil {
		return []models.Administrador{}, err
	}
	return usuarios, nil
}

// BuscarLogado busca dados exceto a senha de um usuário pela id
func BuscarAdministradorLogado(id int, db *sql.DB) (models.Administrador, error) {
	sqlStatement := `SELECT id_administrador, nome, telefone, email, galpao FROM administradores WHERE id_administrador=$1`
	var usuario models.Administrador
	if erro := db.QueryRow(sqlStatement, id).Scan(&usuario.Id, &usuario.Nome, &usuario.Telefone, &usuario.Email, &usuario.Galpao); erro != nil {
		if erro == sql.ErrNoRows {
			return models.Administrador{}, errors.New("Id nao encontrado")
		}
		return models.Administrador{}, erro
	}
	return usuario, nil
}

// Deleta um usuario
func DeletarAdministrador(id int, db *sql.DB) error {
	sqlStatement := `DELETE FROM administradores WHERE id_administrador=$1`
	result, erro := db.Exec(sqlStatement, id)
	if erro != nil {
		return erro
	}
	rowsAffected, erro := result.RowsAffected()
	if erro != nil {
		return erro // Retorna erro se não foi possível verificar as linhas afetadas
	}
	if rowsAffected == 0 {
		return errors.New("usuario logado nao existe")
	}
	return nil
}

// BuscarSenhaPorId usa id para buscar senha de um usuário
func BuscarSenhaPorIdAdministrador(id int, db *sql.DB) (string, error) {
	sqlStatement := `SELECT senha FROM administradores WHERE id_administrador=$1`
	var senhaSalva string
	if erro := db.QueryRow(sqlStatement, id).Scan(&senhaSalva); erro != nil {
		if erro == sql.ErrNoRows {
			return "", errors.New("usuario com esse id nao encontrado")
		}
		return "", erro
	}
	return senhaSalva, nil
}

// AtualizarSenhaAdministradores atualiza senha na tabela administradores
func AtualizarSenhaAdministrador(senha string, id int, db *sql.DB) error {
	sqlStatement := `UPDATE administradores SET senha=$1 WHERE id_administrador=$2`
	result, erro := db.Exec(sqlStatement, senha, id)
	if erro != nil {
		return erro
	}
	// Verifica se alguma linha foi atualizada
	rowsAffected, erro := result.RowsAffected()
	if erro != nil {
		return erro // Retorna erro se não foi possível verificar as linhas afetadas
	}
	if rowsAffected == 0 {
		return errors.New("usuario nao encontrado para atualizar dados")
	}
	return nil
}

// AtualizarNome atualiza Nome na tabela administradores
func AtualizarNomeAdministrador(dados models.Administrador, db *sql.DB) error {
	sqlStatement := `UPDATE administradores SET nome=$1 WHERE id_administrador=$2`
	result, erro := db.Exec(sqlStatement, dados.Nome, dados.Id)
	if erro != nil {
		return erro
	}
	// Verifica se alguma linha foi atualizada
	rowsAffected, erro := result.RowsAffected()
	if erro != nil {
		return erro // Retorna erro se não foi possível verificar as linhas afetadas
	}
	if rowsAffected == 0 {
		return errors.New("usuario nao encontrado para atualizar dados")
	}
	return nil
}
