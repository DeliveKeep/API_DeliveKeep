package routes

import (
	"github.com/go-chi/chi"
)

// Rotear adciona as rotas da api ao roteador
func Rotear() chi.Router {
	r := chi.NewRouter()

	r.Mount("/clientes", UsuariosRouter())

	r.Mount("/pedidos", PedidosRouter())

	r.Mount("/notificacoes", NotificacoesRouter())

	return r
}
