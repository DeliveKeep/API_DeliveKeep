package models

import "time"

// Encomenda representa uma encomenda no sistema
type Encomenda struct {
	ID                  uint64  `json:"id,omitempty"`
	Descricao           string  `json:"descricao,omitempty"`
	CodigoRastreamento  string  `json:"codigo_rastreamento,omitempty"`
	RemetenteNome       string  `json:"remetente_nome,omitempty"`
	RemetenteEndereco   string  `json:"remetente_endereco,omitempty"`
	DestinatarioNome    string  `json:"destinatario_nome,omitempty"`
	DestinatarioEndereco string  `json:"destinatario_endereco,omitempty"`
	AlturaCM            float64 `json:"altura_cm,omitempty"`
	LarguraCM           float64 `json:"largura_cm,omitempty"`
	ComprimentoCM       float64 `json:"comprimento_cm,omitempty"`
	PesoKG              float64 `json:"peso_kg,omitempty"`
	UsuarioID           uint64  `json:"usuario_id,omitempty"`
	CriadoEm            time.Time `json:"criado_em,omitempty"`
}