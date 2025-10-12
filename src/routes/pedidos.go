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

	r.Get("/", controllers.BuscarPedidos)

	//r.Get("/operador", controllers.BuscarPedidosOperador)

	return r
}
