package models

import (
	"API/src/security"
	"errors"
	"strings"

	"github.com/badoux/checkmail"
)

type Usuario struct {
	Id       int    `json:"id,omitempty"`
	Nome     string `json:"nome,omitempty"`
	Cpf      string `json:"cpf,omitempty"`
	Telefone string `json:"telefone,omitempty"`
	Email    string `json:"email,omitempty"`
	Senha    string `json:"senha,omitempty"`
	Endereco string `json:"endereco,omitempty"`
}

// Validar valida formato e tamanho dos dados, remove espaços em branco e criptografa a senha
func (u *Usuario) Validar() error {
	u.Nome = strings.TrimSpace(u.Nome)
	if len(u.Nome) < 2 {
		return errors.New("nome deve ter pelo menos 2 caracteres")
	}
	if erro := checkmail.ValidateFormat(u.Email); erro != nil {
		return errors.New("email inváido")
	}
	u.Senha = strings.TrimSpace(u.Senha)
	if len(u.Senha) < 2 {
		return errors.New("senha deve ter pelo menos 2 caracteres")
	}
	senhaHash, erro := security.GerarSenhaComHash(u.Senha)
	if erro != nil {
		return erro
	}
	u.Senha = string(senhaHash)
	return nil
}

// ValidarEmail valida email
func (u *Usuario) ValidarEmail() error {
	if erro := checkmail.ValidateFormat(u.Email); erro != nil {
		return errors.New("email invalido")
	}
	return nil
}

// ValidarLogin verifica se dados de login estão presentes
func (u *Usuario) ValidarLogin() error {
	if u.Email == "" || u.Senha == "" {
		return errors.New("campos faltando")
	}
	return nil
}
