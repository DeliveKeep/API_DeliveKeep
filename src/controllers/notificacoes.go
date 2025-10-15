package controllers

import (
	"API/src/config"
	"API/src/database"
	"API/src/responses"
	"net/http"
)

// Busca notificacoes de um cliente usuário
func BuscarNotificacoes(w http.ResponseWriter, r *http.Request) {
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
	dados, erro := repositories.BuscarNotificacoes(db, cpf_cliente)
	if erro != nil {
		responses.RespostaDeErro(w, http.StatusInternalServerError, erro)
		return
	}
	// Enviando resposta
	responses.RespostaDeSucesso(w, http.StatusOK, dados)
}
