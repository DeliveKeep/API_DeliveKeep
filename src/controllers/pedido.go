package controllers

import (
	"API/src/config"
	"API/src/database"
	"API/src/models"
	"API/src/repositories"
	"API/src/responses"
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
)

// Criarpedido cria um novo pedido
func CriarPedido(w http.ResponseWriter, r *http.Request) {
	// Lendo corpo da requisição
	corpoReq, erro := io.ReadAll(r.Body)
	if erro != nil {
		responses.RespostaDeErro(w, http.StatusUnprocessableEntity, erro)
		return
	}
	defer r.Body.Close()
	// Passando para struct e validando
	var pedido models.Encomenda
	if erro = json.Unmarshal(corpoReq, &pedido); erro != nil {
		responses.RespostaDeErro(w, http.StatusBadRequest, erro)
		return
	}
	// Extraindo permissao do logado do contexto da requisição
	permissao := r.Context().Value(config.PermissaoKey).(string)
	if permissao != "o" && permissao != "a" {
		responses.RespostaDeErro(w, http.StatusForbidden, errors.New("permissão negada"))
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
	if erro = repositories.CriarPedido(&pedido, db); erro != nil {
		responses.RespostaDeErro(w, http.StatusInternalServerError, erro)
		return
	}
	// Criar a notificação de criação do pedido
	if erro = repositories.NotificacaoCriacao(pedido.Id, db); erro != nil {
		responses.RespostaDeErro(w, http.StatusInternalServerError, erro)
		return
	}
	// Enviando resposta de sucesso
	responses.RespostaDeSucesso(w, http.StatusCreated, pedido)
}

// BuscarPedido busca um pedido
func BuscarPedido(w http.ResponseWriter, r *http.Request) {
	// Extraindo id logado do contexto da requisição
	parametro := chi.URLParam(r, "id")
	idPedido, erro := strconv.Atoi(parametro)
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
	dados, erro := repositories.BuscarPedido(idPedido, db)
	if erro != nil {
		responses.RespostaDeErro(w, http.StatusInternalServerError, erro)
		return
	}
	// Enviando resposta
	responses.RespostaDeSucesso(w, http.StatusOK, dados)
}

// Busca pedidos do cliente logado
func BuscarPedidosCliente(w http.ResponseWriter, r *http.Request) {
	// Abrindo conexão com banco de dados
	db, erro := database.ConectarDB()
	if erro != nil {
		responses.RespostaDeErro(w, http.StatusInternalServerError, erro)
		return
	}
	defer db.Close()
	// Extraindo id do logado do contexto da requisição
	idLogado := r.Context().Value(config.IdKey).(int)
	// Buscando cpf do cliente logado no banco de dados
	cpf_cliente, erro := repositories.BuscarCPFLogado(db, idLogado)
	if erro != nil {
		responses.RespostaDeErro(w, http.StatusInternalServerError, erro)
		return
	}
	// Chamando repositories para buscar dados do usuário logado
	dados, erro := repositories.BuscarPedidos(db, cpf_cliente)
	if erro != nil {
		responses.RespostaDeErro(w, http.StatusInternalServerError, erro)
		return
	}
	// Enviando resposta
	responses.RespostaDeSucesso(w, http.StatusOK, dados)
}

// Busca pedidos de um galpão
func BuscarPedidosGalpao(w http.ResponseWriter, r *http.Request) {
	// Extraindo permissao do logado do contexto da requisição
	permissao := r.Context().Value(config.PermissaoKey).(string)
	if permissao != "o" && permissao != "a" {
		responses.RespostaDeErro(w, http.StatusForbidden, errors.New("permissão negada"))
		return
	}
	// Abrindo conexão com banco de dados
	db, erro := database.ConectarDB()
	if erro != nil {
		responses.RespostaDeErro(w, http.StatusInternalServerError, erro)
		return
	}
	defer db.Close()
	// Extraindo id do logado do contexto da requisição
	idLogado := r.Context().Value(config.IdKey).(int)
	// Buscando galpão do operador
	idGalpao, erro := repositories.BuscarIdGalpao(idLogado, permissao, db)
	if erro != nil {
		responses.RespostaDeErro(w, http.StatusInternalServerError, erro)
		return
	}
	// Chamando repositories para buscar dados do usuário logado
	dados, erro := repositories.BuscarPedidosGalpao(db, idGalpao)
	if erro != nil {
		responses.RespostaDeErro(w, http.StatusInternalServerError, erro)
		return
	}
	// Enviando resposta
	responses.RespostaDeSucesso(w, http.StatusOK, dados)
}

// Atualiza dados de uma encomenda
func AtualizarEncomenda(w http.ResponseWriter, r *http.Request) {
	// Extraindo permissao do logado do contexto da requisição
	permissao := r.Context().Value(config.PermissaoKey).(string)
	if permissao != "a" {
		responses.RespostaDeErro(w, http.StatusForbidden, errors.New("permissão negada"))
		return
	}
	// Lendo corpo da requisição
	corpoReq, erro := io.ReadAll(r.Body)
	if erro != nil {
		responses.RespostaDeErro(w, http.StatusUnprocessableEntity, erro)
		return
	}
	defer r.Body.Close()
	// Passando para struct e validando
	var pedido models.Encomenda
	if erro = json.Unmarshal(corpoReq, &pedido); erro != nil {
		responses.RespostaDeErro(w, http.StatusBadRequest, erro)
		return
	}
	// Extraindo id da encomenda da requisição
	parametro := chi.URLParam(r, "id")
	idPedido, erro := strconv.Atoi(parametro)
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
	// Chamando repositories para inserir dados no banco de dados
	if erro = repositories.AtualizarEncomenda(pedido, idPedido, db); erro != nil {
		responses.RespostaDeErro(w, http.StatusInternalServerError, erro)
		return
	}
	// Enviando resposta de sucesso
	responses.RespostaDeSucesso(w, http.StatusNoContent, nil)
}
