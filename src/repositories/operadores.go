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

// Deleta um usuario
func DeletarOperador(id int, db *sql.DB) error {
	sqlStatement := `DELETE FROM operadores WHERE id_operador=$1`
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
func BuscarSenhaPorIdOperador(id int, db *sql.DB) (string, error) {
	sqlStatement := `SELECT senha FROM operadores WHERE id_operador=$1`
	var senhaSalva string
	if erro := db.QueryRow(sqlStatement, id).Scan(&senhaSalva); erro != nil {
		if erro == sql.ErrNoRows {
			return "", errors.New("usuario com esse id nao encontrado")
		}
		return "", erro
	}
	return senhaSalva, nil
}

// AtualizarSenhaOperador atualiza senha na tabela operadores
func AtualizarSenhaOperador(senha string, id int, db *sql.DB) error {
	sqlStatement := `UPDATE operadores SET senha=$1 WHERE id_operador=$2`
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
