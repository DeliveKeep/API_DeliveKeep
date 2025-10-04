package controllers

import (
	"API/src/config"
	"API/src/database"
	"API/src/models"
	"API/src/repositories"
	"API/src/responses"
	"encoding/json"
	"io"
	"net/http"
)

// CriarEncomenda recebe os dados de uma encomenda e os salva no banco
func CriarEncomenda(w http.ResponseWriter, r *http.Request) {
	// Passo 1: Extrair o ID do usuário que fez a requisição (que está logado)
	idLogado := r.Context().Value(config.IdKey).(int)

	// Passo 2: Ler o corpo da requisição (o JSON com os dados da encomenda)
	corpoReq, erro := io.ReadAll(r.Body)
	if erro != nil {
		responses.RespostaDeErro(w, http.StatusUnprocessableEntity, erro)
		return
	}
	defer r.Body.Close()

	// Passo 3: Converter o JSON para a nossa struct de Encomenda
	var encomenda models.Encomenda
	if erro = json.Unmarshal(corpoReq, &encomenda); erro != nil {
		responses.RespostaDeErro(w, http.StatusBadRequest, erro)
		return
	}

	// Passo 4: Associar a encomenda ao usuário que a criou
	encomenda.UsuarioID = uint64(idLogado)

	// (Opcional) Aqui entraria a validação dos dados da encomenda,
	// para ver se campos obrigatórios não estão vazios, etc.

	// Passo 5: Abrir a conexão com o banco de dados
	db, erro := database.ConectarDB()
	if erro != nil {
		responses.RespostaDeErro(w, http.StatusInternalServerError, erro)
		return
	}
	defer db.Close()

	// Passo 6: Chamar o repository para salvar a encomenda no banco
	novoID, erro := repositories.CriarEncomenda(encomenda, db)
	if erro != nil {
		responses.RespostaDeErro(w, http.StatusInternalServerError, erro)
		return
	}

	// Passo 7: Retornar uma resposta de sucesso com os dados da encomenda criada
	encomenda.ID = novoID
	responses.RespostaDeSucesso(w, http.StatusCreated, encomenda)
}