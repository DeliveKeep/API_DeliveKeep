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

// Cria um administrador
func CriarAdministrador(w http.ResponseWriter, r *http.Request) {
	// Lendo corpo da requisição
	corpoReq, erro := io.ReadAll(r.Body)
	if erro != nil {
		responses.RespostaDeErro(w, http.StatusUnprocessableEntity, erro)
		return
	}
	defer r.Body.Close()
	// Passando para struct e validando
	var usuario models.Administrador
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
	if erro = repositories.CriarAdministrador(&usuario, db); erro != nil {
		responses.RespostaDeErro(w, http.StatusInternalServerError, erro)
		return
	}
	// Enviando resposta de sucesso
	responses.RespostaDeSucesso(w, http.StatusCreated, usuario)
}

// Login executa o login de um usuário
func LoginAdministrador(w http.ResponseWriter, r *http.Request) {
	// Lendo corpo da requisição
	corpoReq, erro := io.ReadAll(r.Body)
	if erro != nil {
		responses.RespostaDeErro(w, http.StatusUnprocessableEntity, erro)
		return
	}
	defer r.Body.Close()
	// Passando para struct
	var usuario models.Administrador
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
	IdESenha, erro := repositories.BuscarIdESenhaPorEmailAdministrador(usuario.Email, db)
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
	permissao := "a"
	token, erro := auth.GerarToken(IdESenha.Id, permissao)
	if erro != nil {
		responses.RespostaDeErro(w, http.StatusInternalServerError, erro)
		return
	}
	resposta := models.RespostaLogin{Id: IdESenha.Id, Token: token}
	// Enviando resposta de sucesso
	responses.RespostaDeSucesso(w, http.StatusOK, resposta)
}

// Busca dados de todos usuários
func BuscarAdministradores(w http.ResponseWriter, r *http.Request) {
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
func BuscarAdministradorLogado(w http.ResponseWriter, r *http.Request) {
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
	dados, erro := repositories.BuscarAdministradorLogado(idLogado, db)
	if erro != nil {
		responses.RespostaDeErro(w, http.StatusInternalServerError, erro)
		return
	}
	// Enviando resposta
	responses.RespostaDeSucesso(w, http.StatusOK, dados)
}

// Busca dados de um usuário pelo id
func BuscarAdministrador(w http.ResponseWriter, r *http.Request) {
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
	dados, erro := repositories.BuscarAdministradorLogado(id, db)
	if erro != nil {
		responses.RespostaDeErro(w, http.StatusInternalServerError, erro)
		return
	}
	// Enviando resposta
	responses.RespostaDeSucesso(w, http.StatusOK, dados)
}

// Deleta o usuario logado
func DeletarAdministrador(w http.ResponseWriter, r *http.Request) {
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
	if erro = repositories.DeletarAdministrador(idLogado, db); erro != nil {
		responses.RespostaDeErro(w, http.StatusInternalServerError, erro)
		return
	}
	// Enviando resposta de sucesso
	responses.RespostaDeSucesso(w, http.StatusNoContent, nil)
}

// Atualiza senha de um usuário
func AtualizarSenhaAdministrador(w http.ResponseWriter, r *http.Request) {
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
	senhaSalva, erro := repositories.BuscarSenhaPorIdAdministrador(idLogado, db)
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
	if erro = repositories.AtualizarSenhaAdministrador(string(senhaNovaHash), idLogado, db); erro != nil {
		responses.RespostaDeErro(w, http.StatusInternalServerError, erro)
		return
	}
	// Enviando resposta de sucesso
	responses.RespostaDeSucesso(w, http.StatusNoContent, nil)
}
