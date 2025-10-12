package models

type Notificacao struct {
	Id        int    `json:"id_notificacao,omitempty"`
	Id_pedido int    `json:"id_pedido,omitempty"`
	Conteudo  string `json:"conteudo,omitempty"`
}
