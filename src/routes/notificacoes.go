package routes

import (
	"API/src/controllers"
	"API/src/middlewares"

	"github.com/go-chi/chi"
)

// PedidosRouters retorna roteador com rotas /notificacoes
func NotificacoesRouter() chi.Router {
	r := chi.NewRouter()

	r.Use(middlewares.Autenticar)

	r.Get("/", controllers.BuscarNotificacoes)

	return r
}
