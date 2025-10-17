package controllers

import (
	"API/src/database"
	"API/src/models"
	"API/src/repositories"
	"API/src/responses"
	"encoding/json"
	"io"
	"net/http"
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
