package models

type Notificacao struct {
	Id_notificacao int    `json:"id_notificacao,omitempty"`
	Id_encomenda   int    `json:"id_encomenda,omitempty"`
	Conteudo       string `json:"conteudo,omitempty"`
}
