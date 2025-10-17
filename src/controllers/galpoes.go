package controllers

import (
	"API/src/database"
	"API/src/models"
	"API/src/repositories"
	"API/src/responses"
	"encoding/json"
	"io"
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
)

// Cria um galpão
func CriarGalpao(w http.ResponseWriter, r *http.Request) {
	// Lendo corpo da requisição
	corpoReq, erro := io.ReadAll(r.Body)
	if erro != nil {
		responses.RespostaDeErro(w, http.StatusUnprocessableEntity, erro)
		return
	}
	defer r.Body.Close()
	// Passando para struct e validando
	var galpao models.Galpao
	if erro = json.Unmarshal(corpoReq, &galpao); erro != nil {
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
	if erro = repositories.CriarGalpao(&galpao, db); erro != nil {
		responses.RespostaDeErro(w, http.StatusInternalServerError, erro)
		return
	}
	// Enviando resposta de sucesso
	responses.RespostaDeSucesso(w, http.StatusCreated, galpao)
}

// Busca dados de um galpão
func BuscarGalpoes(w http.ResponseWriter, r *http.Request) {
	// Abrindo conexão com banco de dados
	db, erro := database.ConectarDB()
	if erro != nil {
		responses.RespostaDeErro(w, http.StatusInternalServerError, erro)
		return
	}
	defer db.Close()
	// Chamando repositories para buscar dados do usuário logado
	dados, erro := repositories.BuscarGalpoes(db)
	if erro != nil {
		responses.RespostaDeErro(w, http.StatusInternalServerError, erro)
		return
	}
	// Enviando resposta
	responses.RespostaDeSucesso(w, http.StatusOK, dados)
}

// Busca dados de um galpão pelo id passado na url
func BuscarGalpao(w http.ResponseWriter, r *http.Request) {
	// Extraindo id
	parametro := chi.URLParam(r, "id")
	idGalpao, erro := strconv.Atoi(parametro)
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
	// Chamando repositories para buscar dados do galpao
	dados, erro := repositories.BuscarGalpao(idGalpao, db)
	if erro != nil {
		responses.RespostaDeErro(w, http.StatusInternalServerError, erro)
		return
	}
	// Enviando resposta
	responses.RespostaDeSucesso(w, http.StatusOK, dados)
}
