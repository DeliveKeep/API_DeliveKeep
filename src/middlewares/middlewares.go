package middlewares

import (
	"API/src/auth"
	"API/src/config"
	"API/src/responses"
	"context"
	"net/http"
)

// Autenticar verifica se exsite um token no cabeçalho da req e se ele é válido
func Autenticar(proximaFunc http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		//vendo se o token é válido
		token, erro := auth.ExtrairToken(r)
		if erro != nil {
			responses.RespostaDeErro(w, http.StatusUnauthorized, erro)
			return
		}
		if erro = auth.ValidarToken(token); erro != nil {
			responses.RespostaDeErro(w, http.StatusUnauthorized, erro)
			return
		}

		// Buscando id e permissao do usuário logado do token
		idLogado, permissaoLogado, erro := auth.ExtrairIDePermissao(token)
		if erro != nil {
			responses.RespostaDeErro(w, http.StatusUnauthorized, erro)
			return
		}
		// Salvando id e permissao no contexto da requisição
		ctx := context.WithValue(r.Context(), config.IdKey, idLogado)
		ctx = context.WithValue(ctx, config.PermissaoKey, permissaoLogado)
		proximaFunc.ServeHTTP(w, r.WithContext(ctx))
	})
}
