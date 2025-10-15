package models

type Encomenda struct {
	Id                  int     `json:"id,omitempty"`
	Nome_remetente      string  `json:"nome_remetente,omitempty"`
	Endereco_remetente  string  `json:"endereco_remetente,omitempty"`
	Codigo_rastreamento string  `json:"codigo_rastreamento,omitempty"`
	Status_pedido       string  `json:"status_pedido,omitempty"`
	Altura              float32 `json:"altura,omitempty"`
	Comprimento         float32 `json:"comprimento,omitempty"`
	Peso                float32 `json:"peso,omitempty"`
	Largura             float32 `json:"largura,omitempty"`
	Descricao           string  `json:"descricao,omitempty"`
	Cpf_cliente         string  `json:"cpf_cliente,omitempty"`
	Id_galpao           int     `json:"id_galpao,omitempty"`
}
