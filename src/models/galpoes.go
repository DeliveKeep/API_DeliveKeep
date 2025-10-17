package models

type Galpao struct {
	Id       int    `json:"id,omitempty"`
	Nome     string `json:"nome,omitempty"`
	Endereco string `json:"endereco,omitempty"`
}
