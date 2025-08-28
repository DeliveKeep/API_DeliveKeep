package responses

import (
	"encoding/json"
	"net/http"
)

type Erro struct {
	Erro string `json:"erro"`
}

// RespostaDeErro formata uma resposta de erro e a envia
func RespostaDeErro(w http.ResponseWriter, statusCode int, erro error) {
	erroStruct := Erro{erro.Error()}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(statusCode)
	if erro := json.NewEncoder(w).Encode(erroStruct); erro != nil {
		http.Error(w, erro.Error(), http.StatusInternalServerError)
	}
}

// RespostaDeSucesso formata uma resposta de sucesso e a envia
func RespostaDeSucesso(w http.ResponseWriter, statusCode int, dados interface{}) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(statusCode)
	if dados != nil {
		if erro := json.NewEncoder(w).Encode(dados); erro != nil {
			http.Error(w, erro.Error(), http.StatusInternalServerError)
		}
	}
}
