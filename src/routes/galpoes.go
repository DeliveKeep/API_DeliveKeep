package routes

import (
	"API/src/controllers"
	"API/src/middlewares"

	"github.com/go-chi/chi"
)

// PedidosRouters retorna roteador com rotas /notificacoes
func GalpoesRouter() chi.Router {
	r := chi.NewRouter()

	r.Post("/", controllers.CriarGalpao)

	r.Get("/", controllers.BuscarGalpoes)

	r.Get("/{id}", controllers.BuscarGalpao)

	r.Group(func(r chi.Router) {
		r.Use(middlewares.Autenticar)

		r.Put("/{id}", controllers.AtualizarGalpao)

		r.Delete("/{id}", controllers.DeletarGalpao)
	})

	return r
}
