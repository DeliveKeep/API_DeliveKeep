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
	NotificacoesAtivas bool   `json:"notificacoes_ativas,omitempty"`
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

// ValidarNome valida tamanho do nome
func (u *Usuario) ValidarNome() error {
	u.Nome = strings.TrimSpace(u.Nome)
	if len(u.Nome) < 2 {
		return errors.New("Nome deve ter pelo menos 2 caractéres")
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

type RespostaLogin struct {
	Id    int    `json:"id"`
	Token string `json:"token"`
}
type Senhas struct {
	SenhaAtual string `json:"senha_atual,omitempty"`
	SenhaNova  string `json:"senha_nova,omitempty"`
}

// ValidarSenhas valida tamanho das senhas
func (u *Senhas) ValidarSenhas() error {
	u.SenhaAtual = strings.TrimSpace(u.SenhaAtual)
	if len(u.SenhaAtual) < 2 {
		return errors.New("senha atual incorreta")
	}
	u.SenhaNova = strings.TrimSpace(u.SenhaNova)
	if len(u.SenhaNova) < 2 {
		return errors.New("senha nova deve ter pelo menos 2 caracteres")
	}
	return nil
}
