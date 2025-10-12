package repositories

import (
	"API/src/models"
	"database/sql"
	"errors"
)

// CriarPedido insere um novo pedido no banco de dados
func CriarPedido(pedido *models.Pedido, db *sql.DB) error {
	sqlStatement := `INSERT INTO pedidos (id_operador, cpf_cliente, nome_remetente, endereco_remetente, nome_destinatario, endereco_destinatario, codigo_rastreamento, altura, comprimento, peso, largura, descricao) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12) RETURNING id`
	if erro := db.QueryRow(sqlStatement, pedido.Id_operador, pedido.Cpf_cliente, pedido.Nome_remetente, pedido.Endereco_remetente, pedido.Nome_destinatario, pedido.Endereco_destinatario, pedido.Codigo_rastreamento, pedido.Altura, pedido.Comprimento, pedido.Peso, pedido.Largura, pedido.Descricao).Scan(&pedido.Id); erro != nil {
		return erro
	}
	return nil
}

// BuscarLogado busca dados de um pedido
func BuscarPedido(id int, db *sql.DB) (models.Pedido, error) {
	sqlStatement := `SELECT id, cpf_cliente, nome_remetente, endereco_remetente, nome_destinatario, endereco_destinatario, codigo_rastreamento, status_pedido, altura, comprimento, peso, largura, descricao FROM pedidos WHERE id=$1`
	var pedido models.Pedido
	if erro := db.QueryRow(sqlStatement, id).Scan(&pedido.Id, &pedido.Cpf_cliente, &pedido.Nome_remetente, &pedido.Endereco_remetente, &pedido.Nome_destinatario, &pedido.Endereco_destinatario, &pedido.Codigo_rastreamento, &pedido.Status_pedido, &pedido.Altura, &pedido.Comprimento, &pedido.Peso, &pedido.Largura, &pedido.Descricao); erro != nil {
		if erro == sql.ErrNoRows {
			return models.Pedido{}, errors.New("Id nao encontrado")
		}
		return models.Pedido{}, erro
	}
	return pedido, nil
}

// Busca pedidos de um usuário cliente específico
func BuscarPedidos(db *sql.DB, cpf_cliente string) ([]models.Pedido, error) {
	sqlStatement := `SELECT id, cpf_cliente, nome_remetente, endereco_remetente, nome_destinatario, endereco_destinatario, codigo_rastreamento, status_pedido, altura, comprimento, peso, largura, descricao FROM pedidos WHERE cpf_cliente=$1`
	rows, err := db.Query(sqlStatement, cpf_cliente)
	if err != nil {
		return []models.Pedido{}, err
	}
	defer rows.Close()
	var pedidos []models.Pedido
	// Itera sobre as linhas retornadas
	for rows.Next() {
		var pedido models.Pedido
		if err := rows.Scan(&pedido.Id, &pedido.Cpf_cliente, &pedido.Nome_remetente, &pedido.Endereco_remetente, &pedido.Nome_destinatario, &pedido.Endereco_destinatario, &pedido.Codigo_rastreamento, &pedido.Status_pedido, &pedido.Altura, &pedido.Comprimento, &pedido.Peso, &pedido.Largura, &pedido.Descricao); err != nil {
			return []models.Pedido{}, err
		}
		pedidos = append(pedidos, pedido)
	}
	// Verifica se ocorreu algum erro durante a iteração
	if err = rows.Err(); err != nil {
		return []models.Pedido{}, err
	}
	return pedidos, nil
}

// Busca pedidos de um usuário operador específico
func BuscarPedidosOperador(db *sql.DB, id int) ([]models.Pedido, error) {
	sqlStatement := `SELECT id, cpf_cliente, nome_remetente, endereco_remetente, nome_destinatario, endereco_destinatario, codigo_rastreamento, status_pedido, altura, comprimento, peso, largura, descricao FROM pedidos WHERE id_operador=$1`
	rows, err := db.Query(sqlStatement, id)
	if err != nil {
		return []models.Pedido{}, err
	}
	defer rows.Close()
	var pedidos []models.Pedido
	// Itera sobre as linhas retornadas
	for rows.Next() {
		var pedido models.Pedido
		if err := rows.Scan(&pedido.Id, &pedido.Cpf_cliente, &pedido.Nome_remetente, &pedido.Endereco_remetente, &pedido.Nome_destinatario, &pedido.Endereco_destinatario, &pedido.Codigo_rastreamento, &pedido.Status_pedido, &pedido.Altura, &pedido.Comprimento, &pedido.Peso, &pedido.Largura, &pedido.Descricao); err != nil {
			return []models.Pedido{}, err
		}
		pedidos = append(pedidos, pedido)
	}
	// Verifica se ocorreu algum erro durante a iteração
	if err = rows.Err(); err != nil {
		return []models.Pedido{}, err
	}
	return pedidos, nil
}

// Cria notificação de criação de pedido
func NotificacaoCriacao(id int, db *sql.DB) error {
	conteudo := "Sua encomenda foi criada, veja detalhes aqui"
	sqlStatement := `INSERT INTO notificacoes (id_pedido, conteudo) VALUES ($1, $2)`
	_, err := db.Exec(sqlStatement, id, conteudo)
	if err != nil {
		return err
	}
	return nil
}
