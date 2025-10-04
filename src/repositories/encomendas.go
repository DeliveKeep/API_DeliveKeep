package repositories

import (
	"API/src/models"
	"database/sql"
)

// CriarEncomenda insere uma nova encomenda no banco de dados
func CriarEncomenda(encomenda models.Encomenda, db *sql.DB) (uint64, error) {
	statement, erro := db.Prepare(
		`INSERT INTO encomendas (
			descricao, codigo_rastreamento, remetente_nome, remetente_endereco,
			destinatario_nome, destinatario_endereco, altura_cm, largura_cm,
			comprimento_cm, peso_kg, usuario_id
		) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11) RETURNING id`,
	)
	if erro != nil {
		return 0, erro
	}
	defer statement.Close()

	var ultimoID uint64
	erro = statement.QueryRow(
		encomenda.Descricao,
		encomenda.CodigoRastreamento,
		encomenda.RemetenteNome,
		encomenda.RemetenteEndereco,
		encomenda.DestinatarioNome,
		encomenda.DestinatarioEndereco,
		encomenda.AlturaCM,
		encomenda.LarguraCM,
		encomenda.ComprimentoCM,
		encomenda.PesoKG,
		encomenda.UsuarioID,
	).Scan(&ultimoID)

	if erro != nil {
		return 0, erro
	}

	return ultimoID, nil
}