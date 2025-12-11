package routes

import (
	"API/src/controllers"
	"API/src/middlewares"

	"github.com/go-chi/chi"
)

// PedidosRouters retorna roteador com rotas /pedidos
func PedidosRouter() chi.Router {
	r := chi.NewRouter()

	r.Use(middlewares.Autenticar)

	r.Post("/", controllers.CriarPedido)

	r.Get("/{id}", controllers.BuscarPedido)

	r.Get("/cliente", controllers.BuscarPedidosCliente)

	r.Get("/galpao", controllers.BuscarPedidosGalpao)

	r.Put("/{id}", controllers.AtualizarEncomenda)

	r.Patch("/atualizar-status/{id}", controllers.AtualizarStatus)

	return r
}
