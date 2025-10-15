package repositories

import (
	"API/src/models"
	"database/sql"
)

// Busca notificacoes de um usuário cliente
func BuscarNotificacoes(db *sql.DB, cpf string) ([]models.Notificacao, error) {
	sqlStatement := `SELECT n.id_notificacao, n.id_encomenda, n.conteudo FROM notificacoes AS n INNER JOIN encomendas AS p ON p.id_encomenda = n.id_encomenda WHERE p.cpf_cliente = $1`
	rows, err := db.Query(sqlStatement, cpf)
	if err != nil {
		return []models.Notificacao{}, err
	}
	defer rows.Close()
	var notificacoes []models.Notificacao
	// Itera sobre as linhas retornadas
	for rows.Next() {
		var notificacao models.Notificacao
		if err := rows.Scan(&notificacao.Id_notificacao, &notificacao.Id_encomenda, &notificacao.Conteudo); err != nil {
			return []models.Notificacao{}, err
		}
		notificacoes = append(notificacoes, notificacao)
	}
	// Verifica se ocorreu algum erro durante a iteração
	if err = rows.Err(); err != nil {
		return []models.Notificacao{}, err
	}
	return notificacoes, nil
}
