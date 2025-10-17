package controllers

import (
	"API/src/auth"
	"API/src/config"
	"API/src/database"
	"API/src/models"
	"API/src/repositories"
	"API/src/responses"
	"API/src/security"
	"encoding/json"
	"io"
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
)

func CriarOperador(w http.ResponseWriter, r *http.Request) {
	// Lendo corpo da requisição
	corpoReq, erro := io.ReadAll(r.Body)
	if erro != nil {
		responses.RespostaDeErro(w, http.StatusUnprocessableEntity, erro)
		return
	}
	defer r.Body.Close()
	// Passando para struct e validando
	var usuario models.Operador
	if erro = json.Unmarshal(corpoReq, &usuario); erro != nil {
		responses.RespostaDeErro(w, http.StatusBadRequest, erro)
		return
	}
	// Aqui campos também são formatos e senhas criptogradas
	if erro = usuario.Validar(); erro != nil {
		responses.RespostaDeErro(w, http.StatusBadRequest, erro)
		return
	}
	// Abrindo conexão com banco de dados
	db, erro := database.ConectarDB()
	if erro != nil {
		responses.RespostaDeErro(w, http.StatusInternalServerError, erro)
		return
	}
	defer db.Close()
	// Chamando repositories para inserir dados no banco de dados
	if erro = repositories.CriarOperador(&usuario, db); erro != nil {
		responses.RespostaDeErro(w, http.StatusInternalServerError, erro)
		return
	}
	// Enviando resposta de sucesso
	responses.RespostaDeSucesso(w, http.StatusCreated, usuario)
}

// Login executa o login de um usuário
func LoginOperador(w http.ResponseWriter, r *http.Request) {
	// Lendo corpo da requisição
	corpoReq, erro := io.ReadAll(r.Body)
	if erro != nil {
		responses.RespostaDeErro(w, http.StatusUnprocessableEntity, erro)
		return
	}
	defer r.Body.Close()
	// Passando para struct
	var usuario models.Operador
	if erro = json.Unmarshal(corpoReq, &usuario); erro != nil {
		responses.RespostaDeErro(w, http.StatusBadRequest, erro)
		return
	}
	if erro = usuario.ValidarLogin(); erro != nil {
		responses.RespostaDeErro(w, http.StatusBadRequest, erro)
		return
	}
	// Abrindo conexão com banco de dados
	db, erro := database.ConectarDB()
	if erro != nil {
		responses.RespostaDeErro(w, http.StatusInternalServerError, erro)
		return
	}
	defer db.Close()
	// Chamando repositories para buscar senha para comparação
	IdESenha, erro := repositories.BuscarIdESenhaPorEmailOperador(usuario.Email, db)
	if erro != nil {
		responses.RespostaDeErro(w, http.StatusInternalServerError, erro)
		return
	}
	// Verificando se senha está correta
	if erro = security.VerificarSenha(IdESenha.Senha, usuario.Senha); erro != nil {
		responses.RespostaDeErro(w, http.StatusUnauthorized, erro)
		return
	}
	// Gerando token
	permissao := "o"
	token, erro := auth.GerarToken(IdESenha.Id, permissao)
	if erro != nil {
		responses.RespostaDeErro(w, http.StatusInternalServerError, erro)
		return
	}
	resposta := models.RespostaLogin{Id: IdESenha.Id, Token: token}
	// Enviando resposta de sucesso
	responses.RespostaDeSucesso(w, http.StatusOK, resposta)
}

// BuscarUsuarios busca dados de todos usuários
func BuscarOperadores(w http.ResponseWriter, r *http.Request) {
	// Abrindo conexão com banco de dados
	db, erro := database.ConectarDB()
	if erro != nil {
		responses.RespostaDeErro(w, http.StatusInternalServerError, erro)
		return
	}
	defer db.Close()
	// Chamando repositories para buscar dados do usuário logado
	dados, erro := repositories.BuscarOperadores(db)
	if erro != nil {
		responses.RespostaDeErro(w, http.StatusInternalServerError, erro)
		return
	}
	// Enviando resposta
	responses.RespostaDeSucesso(w, http.StatusOK, dados)
}

// BuscarLogado busca dados de um usuário logado
func BuscarOperadorLogado(w http.ResponseWriter, r *http.Request) {
	// Extraindo id logado do contexto da requisição
	idLogado := r.Context().Value(config.IdKey).(int)
	// Abrindo conexão com banco de dados
	db, erro := database.ConectarDB()
	if erro != nil {
		responses.RespostaDeErro(w, http.StatusInternalServerError, erro)
		return
	}
	defer db.Close()
	// Chamando repositories para buscar dados do usuário logado
	dados, erro := repositories.BuscarOperadorLogado(idLogado, db)
	if erro != nil {
		responses.RespostaDeErro(w, http.StatusInternalServerError, erro)
		return
	}
	// Enviando resposta
	responses.RespostaDeSucesso(w, http.StatusOK, dados)
}

// Busca dados de um usuário pelo id
func BuscarOperador(w http.ResponseWriter, r *http.Request) {
	// Extraindo id logado do contexto da requisição
	parametro := chi.URLParam(r, "id")
	id, erro := strconv.Atoi(parametro)
	if erro != nil {
		responses.RespostaDeErro(w, http.StatusBadRequest, erro)
		return
	}
	// Abrindo conexão com banco de dados
	db, erro := database.ConectarDB()
	if erro != nil {
		responses.RespostaDeErro(w, http.StatusInternalServerError, erro)
		return
	}
	defer db.Close()
	// Chamando repositories para buscar dados do usuário logado
	dados, erro := repositories.BuscarOperadorLogado(id, db)
	if erro != nil {
		responses.RespostaDeErro(w, http.StatusInternalServerError, erro)
		return
	}
	// Enviando resposta
	responses.RespostaDeSucesso(w, http.StatusOK, dados)
}

// Deleta o usuario logado
func DeletarOperador(w http.ResponseWriter, r *http.Request) {
	// Extraindo id do logado do contexto da requisição
	idLogado := r.Context().Value(config.IdKey).(int)
	// Abrindo conexão com banco de dados
	db, erro := database.ConectarDB()
	if erro != nil {
		responses.RespostaDeErro(w, http.StatusInternalServerError, erro)
		return
	}
	defer db.Close()
	// Chamando repositories para bucar dados no banco de dados
	if erro = repositories.DeletarOperador(idLogado, db); erro != nil {
		responses.RespostaDeErro(w, http.StatusInternalServerError, erro)
		return
	}
	// Enviando resposta de sucesso
	responses.RespostaDeSucesso(w, http.StatusNoContent, nil)
}

// Atualiza senha de um usuário
func AtualizarSenhaOperador(w http.ResponseWriter, r *http.Request) {
	// Lendo corpo da requisição
	corpoReq, erro := io.ReadAll(r.Body)
	if erro != nil {
		responses.RespostaDeErro(w, http.StatusUnprocessableEntity, erro)
		return
	}
	defer r.Body.Close()
	// Passando para struct e validando dados
	var senhas models.Senhas
	if erro = json.Unmarshal(corpoReq, &senhas); erro != nil {
		responses.RespostaDeErro(w, http.StatusBadRequest, erro)
		return
	}
	if erro = senhas.ValidarSenhas(); erro != nil {
		responses.RespostaDeErro(w, http.StatusBadRequest, erro)
		return
	}
	// Extraindo id logado do contexto da requisição
	idLogado := r.Context().Value(config.IdKey).(int)
	// Abrindo conexão com banco de dados
	db, erro := database.ConectarDB()
	if erro != nil {
		responses.RespostaDeErro(w, http.StatusInternalServerError, erro)
		return
	}
	defer db.Close()
	// Chamando repositories para buscar senha no banco de dados
	senhaSalva, erro := repositories.BuscarSenhaPorIdOperador(idLogado, db)
	if erro != nil {
		responses.RespostaDeErro(w, http.StatusInternalServerError, erro)
		return
	}
	// Vendo se senha salva é igual a recebida
	if erro = security.VerificarSenha(senhaSalva, senhas.SenhaAtual); erro != nil {
		responses.RespostaDeErro(w, http.StatusUnauthorized, erro)
		return
	}
	// Criptografando senha nova para guardar no banco
	senhaNovaHash, erro := security.GerarSenhaComHash(senhas.SenhaNova)
	if erro != nil {
		responses.RespostaDeErro(w, http.StatusInternalServerError, erro)
		return
	}
	// Chamando repositories para atualizar senha no banco
	if erro = repositories.AtualizarSenhaOperador(string(senhaNovaHash), idLogado, db); erro != nil {
		responses.RespostaDeErro(w, http.StatusInternalServerError, erro)
		return
	}
	// Enviando resposta de sucesso
	responses.RespostaDeSucesso(w, http.StatusNoContent, nil)
}

// AtualizarNome atualiza Nome de um usuário
func AtualizarNomeOperador(w http.ResponseWriter, r *http.Request) {
	// Lendo corpo da requisição
	corpoReq, erro := io.ReadAll(r.Body)
	if erro != nil {
		responses.RespostaDeErro(w, http.StatusUnprocessableEntity, erro)
		return
	}
	defer r.Body.Close()
	// Passando para struct e validando dados
	var nome models.Operador
	if erro = json.Unmarshal(corpoReq, &nome); erro != nil {
		responses.RespostaDeErro(w, http.StatusBadRequest, erro)
		return
	}
	if erro = nome.ValidarNome(); erro != nil {
		responses.RespostaDeErro(w, http.StatusBadRequest, erro)
		return
	}
	// Extraindo id logado do contexto da requisição
	idLogado := r.Context().Value(config.IdKey).(int)
	nome.Id = idLogado
	// Abrindo conexão com banco de dados
	db, erro := database.ConectarDB()
	if erro != nil {
		responses.RespostaDeErro(w, http.StatusInternalServerError, erro)
		return
	}
	defer db.Close()
	// Chamando repositories para atualizar dados no banco de dados
	if erro = repositories.AtualizarNomeOperador(nome, db); erro != nil {
		responses.RespostaDeErro(w, http.StatusInternalServerError, erro)
		return
	}
	// Enviando resposta de sucesso
	responses.RespostaDeSucesso(w, http.StatusNoContent, nil)
}

// AtualizarTelefoneOperador atualiza telefone de um usuário
func AtualizarTelefoneOperador(w http.ResponseWriter, r *http.Request) {
	// Lendo corpo da requisição
	corpoReq, erro := io.ReadAll(r.Body)
	if erro != nil {
		responses.RespostaDeErro(w, http.StatusUnprocessableEntity, erro)
		return
	}
	defer r.Body.Close()
	// Passando para struct e validando dados
	var telefone models.Operador
	if erro = json.Unmarshal(corpoReq, &telefone); erro != nil {
		responses.RespostaDeErro(w, http.StatusBadRequest, erro)
		return
	}
	// Extraindo id logado do contexto da requisição
	idLogado := r.Context().Value(config.IdKey).(int)
	telefone.Id = idLogado
	// Abrindo conexão com banco de dados
	db, erro := database.ConectarDB()
	if erro != nil {
		responses.RespostaDeErro(w, http.StatusInternalServerError, erro)
		return
	}
	defer db.Close()
	// Chamando repositories para atualizar dados no banco de dados
	if erro = repositories.AtualizarTelefoneOperador(telefone, db); erro != nil {
		responses.RespostaDeErro(w, http.StatusInternalServerError, erro)
		return
	}
	// Enviando resposta de sucesso
	responses.RespostaDeSucesso(w, http.StatusNoContent, nil)
}

// AtualizarEmailOperador atualiza email de um usuário
func AtualizarEmailOperador(w http.ResponseWriter, r *http.Request) {
	// Lendo corpo da requisição
	corpoReq, erro := io.ReadAll(r.Body)
	if erro != nil {
		responses.RespostaDeErro(w, http.StatusUnprocessableEntity, erro)
		return
	}
	defer r.Body.Close()
	// Passando para struct e validando dados
	var email models.Operador
	if erro = json.Unmarshal(corpoReq, &email); erro != nil {
		responses.RespostaDeErro(w, http.StatusBadRequest, erro)
		return
	}
	if erro = email.ValidarEmail(); erro != nil {
		responses.RespostaDeErro(w, http.StatusBadRequest, erro)
		return
	}
	// Extraindo id logado do contexto da requisição
	idLogado := r.Context().Value(config.IdKey).(int)
	email.Id = idLogado
	// Abrindo conexão com banco de dados
	db, erro := database.ConectarDB()
	if erro != nil {
		responses.RespostaDeErro(w, http.StatusInternalServerError, erro)
		return
	}
	defer db.Close()
	// Chamando repositories para atualizar dados no banco de dados
	if erro = repositories.AtualizarEmailOperador(email, db); erro != nil {
		responses.RespostaDeErro(w, http.StatusInternalServerError, erro)
		return
	}
	// Enviando resposta de sucesso
	responses.RespostaDeSucesso(w, http.StatusNoContent, nil)
}
