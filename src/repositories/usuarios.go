package repositories

import (
	"API/src/models"
	"database/sql"
	"errors"
)

// CriarUsuario insere um novo usuario no banco de dados
func CriarUsuario(usuario *models.Cliente, db *sql.DB) error {
	sqlStatement := `INSERT INTO clientes (nome, cpf, endereco, telefone, email, senha) VALUES ($1, $2, $3, $4, $5, $6) RETURNING id_cliente`
	if erro := db.QueryRow(sqlStatement, usuario.Nome, usuario.Cpf, usuario.Endereco, usuario.Telefone, usuario.Email, usuario.Senha).Scan(&usuario.Id); erro != nil {
		return erro
	}
	return nil
}

// BuscarIdSenhaPorEmail usa um email para buscar Id e senha de um usuário
func BuscarIdESenhaPorEmail(email string, db *sql.DB) (models.Cliente, error) {
	sqlStatement := `SELECT id_cliente, senha FROM clientes WHERE email=$1`
	var usuario models.Cliente
	if erro := db.QueryRow(sqlStatement, email).Scan(&usuario.Id, &usuario.Senha); erro != nil {
		if erro == sql.ErrNoRows {
			return models.Cliente{}, errors.New("usuario com esse email nao encontrado")
		}
		return models.Cliente{}, erro
	}
	return usuario, nil
}

// BuscarLogado busca dados exceto a senha de um usuário pela id
func BuscarLogado(id int, db *sql.DB) (models.Cliente, error) {
	sqlStatement := `SELECT id_cliente, nome, cpf, endereco, telefone, email FROM clientes WHERE id_cliente=$1`
	var usuario models.Cliente
	if erro := db.QueryRow(sqlStatement, id).Scan(&usuario.Id, &usuario.Nome, &usuario.Cpf, &usuario.Endereco, &usuario.Telefone, &usuario.Email); erro != nil {
		if erro == sql.ErrNoRows {
			return models.Cliente{}, errors.New("Id nao encontrado")
		}
		return models.Cliente{}, erro
	}
	return usuario, nil
}

// BuscarLogado busca dados exceto a senha de um usuário pela id
func BuscarUsuarios(db *sql.DB) ([]models.Cliente, error) {
	sqlStatement := `SELECT id_cliente, nome, cpf, endereco, telefone, email FROM clientes`
	rows, err := db.Query(sqlStatement)
	if err != nil {
		return []models.Cliente{}, err
	}
	defer rows.Close()
	var usuarios []models.Cliente
	// Itera sobre as linhas retornadas
	for rows.Next() {
		var usuario models.Cliente
		if err := rows.Scan(&usuario.Id, &usuario.Nome, &usuario.Cpf, &usuario.Endereco, &usuario.Telefone, &usuario.Email); err != nil {
			return []models.Cliente{}, err
		}
		usuarios = append(usuarios, usuario)
	}
	// Verifica se ocorreu algum erro durante a iteração
	if err = rows.Err(); err != nil {
		return []models.Cliente{}, err
	}
	return usuarios, nil
}

// DeletarUsuario deleta um usuario
func DeletarUsuario(id int, db *sql.DB) error {
	sqlStatement := `DELETE FROM clientes WHERE id_cliente=$1`
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
func BuscarSenhaPorId(id int, db *sql.DB) (string, error) {
	sqlStatement := `SELECT senha FROM clientes WHERE id_cliente=$1`
	var senhaSalva string
	if erro := db.QueryRow(sqlStatement, id).Scan(&senhaSalva); erro != nil {
		if erro == sql.ErrNoRows {
			return "", errors.New("usuario com esse id nao encontrado")
		}
		return "", erro
	}
	return senhaSalva, nil
}

// AtualizarSenha atualiza senha na tabela usuários
func AtualizarSenha(senha string, id int, db *sql.DB) error {
	sqlStatement := `UPDATE clientes SET senha=$1 WHERE id_cliente=$2`
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

// AtualizarNome atualiza Nome na tabela usuários
func AtualizarNome(dados models.Cliente, db *sql.DB) error {
	sqlStatement := `UPDATE clientes SET nome=$1 WHERE id_cliente=$2`
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

// AtualizarEndereco atualiza endereco na tabela usuários
func AtualizarEndereco(dados models.Cliente, db *sql.DB) error {
	sqlStatement := `UPDATE clientes SET endereco=$1 WHERE id_cliente=$2`
	result, erro := db.Exec(sqlStatement, dados.Endereco, dados.Id)
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

// AtualizarTelefone atualiza telefone na tabela usuários
func AtualizarTelefone(dados models.Cliente, db *sql.DB) error {
	sqlStatement := `UPDATE clientes SET telefone=$1 WHERE id_cliente=$2`
	result, erro := db.Exec(sqlStatement, dados.Telefone, dados.Id)
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

// AtualizarEmail atualiza email na tabela usuários
func AtualizarEmail(dados models.Cliente, db *sql.DB) error {
	sqlStatement := `UPDATE clientes SET email=$1 WHERE id_cliente=$2`
	result, erro := db.Exec(sqlStatement, dados.Email, dados.Id)
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

// Busca cpf do usuário logado
func BuscarCPFLogado(db *sql.DB, id int) (string, error) {
	sqlStatement := `SELECT cpf FROM clientes WHERE id_cliente=$1`
	var cpf string
	if erro := db.QueryRow(sqlStatement, id).Scan(&cpf); erro != nil {
		if erro == sql.ErrNoRows {
			return "", errors.New("Id nao encontrado")
		}
		return "", erro
	}
	return cpf, nil
}
