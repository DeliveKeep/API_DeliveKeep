package repositories

import (
	"API/src/models"
	"database/sql"
)

// CriarUsuario insere um novo pedido no banco de dados
func CriarPedido(pedido *models.Pedido, db *sql.DB) error {
	sqlStatement := `INSERT INTO pedidos (nome_remetente, endereco_remetente, nome_destinatario, endereco_destinatario, codigo_rastreamento, altura, comprimento, peso, largura, descricao) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10) RETURNING id`
	if erro := db.QueryRow(sqlStatement, pedido.Nome_remetente, pedido.Endereco_remetente, pedido.Nome_destinatario, pedido.Endereco_destinatario, pedido.Codigo_rastreamento, pedido.Altura, pedido.Comprimento, pedido.Peso, pedido.Largura, pedido.Descricao).Scan(&pedido.Id); erro != nil {
		return erro
	}
	return nil
}
