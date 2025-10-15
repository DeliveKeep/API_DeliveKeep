package repositories

import (
	"API/src/models"
	"database/sql"
	"errors"
)

// CriarPedido insere um novo pedido no banco de dados
func CriarPedido(pedido *models.Encomenda, db *sql.DB) error {
	sqlStatement := `INSERT INTO encomendas (cpf_cliente, nome_remetente, endereco_remetente, codigo_rastreamento, altura, comprimento, peso, largura, descricao, galpao) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10) RETURNING id_encomenda`
	if erro := db.QueryRow(sqlStatement, pedido.Cpf_cliente, pedido.Nome_remetente, pedido.Endereco_remetente, pedido.Codigo_rastreamento, pedido.Altura, pedido.Comprimento, pedido.Peso, pedido.Largura, pedido.Descricao, pedido.Id_galpao).Scan(&pedido.Id); erro != nil {
		return erro
	}
	return nil
}

// BuscarLogado busca dados de um pedido
func BuscarPedido(id int, db *sql.DB) (models.Encomenda, error) {
	sqlStatement := `SELECT id_encomenda, cpf_cliente, nome_remetente, endereco_remetente, codigo_rastreamento, status_pedido, altura, comprimento, peso, largura, descricao, galpao FROM encomendas WHERE id_encomenda=$1`
	var pedido models.Encomenda
	if erro := db.QueryRow(sqlStatement, id).Scan(&pedido.Id, &pedido.Cpf_cliente, &pedido.Nome_remetente, &pedido.Endereco_remetente, &pedido.Codigo_rastreamento, &pedido.Status_pedido, &pedido.Altura, &pedido.Comprimento, &pedido.Peso, &pedido.Largura, &pedido.Descricao, &pedido.Id_galpao); erro != nil {
		if erro == sql.ErrNoRows {
			return models.Encomenda{}, errors.New("Id nao encontrado")
		}
		return models.Encomenda{}, erro
	}
	return pedido, nil
}

// Busca pedidos de um usuário cliente específico
func BuscarPedidos(db *sql.DB, cpf_cliente string) ([]models.Encomenda, error) {
	sqlStatement := `SELECT id_encomenda, cpf_cliente, nome_remetente, endereco_remetente, codigo_rastreamento, status_pedido, altura, comprimento, peso, largura, descricao, galpao FROM encomendas WHERE cpf_cliente=$1`
	rows, err := db.Query(sqlStatement, cpf_cliente)
	if err != nil {
		return []models.Encomenda{}, err
	}
	defer rows.Close()
	var pedidos []models.Encomenda
	// Itera sobre as linhas retornadas
	for rows.Next() {
		var pedido models.Encomenda
		if err := rows.Scan(&pedido.Id, &pedido.Cpf_cliente, &pedido.Nome_remetente, &pedido.Endereco_remetente, &pedido.Codigo_rastreamento, &pedido.Status_pedido, &pedido.Altura, &pedido.Comprimento, &pedido.Peso, &pedido.Largura, &pedido.Descricao, &pedido.Id_galpao); err != nil {
			return []models.Encomenda{}, err
		}
		pedidos = append(pedidos, pedido)
	}
	// Verifica se ocorreu algum erro durante a iteração
	if err = rows.Err(); err != nil {
		return []models.Encomenda{}, err
	}
	return pedidos, nil
}

// Busca pedidos de um usuário operador específico
func BuscarPedidosGalpao(db *sql.DB, id int) ([]models.Encomenda, error) {
	sqlStatement := `SELECT id_encomenda, cpf_cliente, nome_remetente, endereco_remetente, codigo_rastreamento, status_pedido, altura, comprimento, peso, largura, descricao, galpao FROM encomendas WHERE galpao=$1`
	rows, err := db.Query(sqlStatement, id)
	if err != nil {
		return []models.Encomenda{}, err
	}
	defer rows.Close()
	var pedidos []models.Encomenda
	// Itera sobre as linhas retornadas
	for rows.Next() {
		var pedido models.Encomenda
		if err := rows.Scan(&pedido.Id, &pedido.Cpf_cliente, &pedido.Nome_remetente, &pedido.Endereco_remetente, &pedido.Codigo_rastreamento, &pedido.Status_pedido, &pedido.Altura, &pedido.Comprimento, &pedido.Peso, &pedido.Largura, &pedido.Descricao, &pedido.Id_galpao); err != nil {
			return []models.Encomenda{}, err
		}
		pedidos = append(pedidos, pedido)
	}
	// Verifica se ocorreu algum erro durante a iteração
	if err = rows.Err(); err != nil {
		return []models.Encomenda{}, err
	}
	return pedidos, nil
}

// Cria notificação de criação de pedido
func NotificacaoCriacao(id int, db *sql.DB) error {
	conteudo := "Sua encomenda foi criada, veja detalhes aqui"
	sqlStatement := `INSERT INTO notificacoes (id_encomenda, conteudo) VALUES ($1, $2)`
	_, err := db.Exec(sqlStatement, id, conteudo)
	if err != nil {
		return err
	}
	return nil
}

// Busca id do galpão onde o operador/administrador trabalha
func BuscarIdGalpao(id int, permissao string, db *sql.DB) (int, error) {
	var sqlStatement string
	if permissao == "o" {
		sqlStatement = `SELECT galpao FROM operadores WHERE id_operador=$1`
	} else if permissao == "a" {
		sqlStatement = `SELECT galpao FROM administradores WHERE id_administrador=$1`
	} else {
		return 0, errors.New("permissao invalida: use 'o' ou 'a'")
	}
	var galpao int
	if erro := db.QueryRow(sqlStatement, id).Scan(&galpao); erro != nil {
		if erro == sql.ErrNoRows {
			return 0, errors.New("Id nao encontrado")
		}
		return 0, erro
	}
	return galpao, nil
}
