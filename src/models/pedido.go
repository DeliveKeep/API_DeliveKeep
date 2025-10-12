package models

type Pedido struct {
	Id                    int     `json:"id,omitempty"`
	Id_operador           int     `json:"id_operador,omitempty"`
	Cpf_cliente           int     `json:"cpf_cliente,omitempty"`
	Nome_remetente        string  `json:"nome_remetente,omitempty"`
	Endereco_remetente    string  `json:"endereco_remetente,omitempty"`
	Nome_destinatario     string  `json:"nome_destinatario,omitempty"`
	Endereco_destinatario string  `json:"endereco_destinatario,omitempty"`
	Codigo_rastreamento   string  `json:"codigo_rastreamento,omitempty"`
	Status_pedido         string  `json:"status_pedido,omitempty"`
	Altura                float32 `json:"altura,omitempty"`
	Comprimento           float32 `json:"comprimento,omitempty"`
	Peso                  float32 `json:"peso,omitempty"`
	Largura               float32 `json:"largura,omitempty"`
	Descricao             string  `json:"descricao,omitempty"`
}
