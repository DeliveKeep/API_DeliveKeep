package routes

import (
	"github.com/go-chi/chi"
)

// Rotear adciona as rotas da api ao roteador
func Rotear() chi.Router {
	r := chi.NewRouter()

	// /usuarios

	r.Mount("/usuarios", UsuariosRouter())

	return r
}
