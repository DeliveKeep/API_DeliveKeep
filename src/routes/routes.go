package routes

import (
	"github.com/go-chi/chi"
	"github.com/go-chi/cors"
)

// Rotear adciona as rotas da api ao roteador
func Rotear() chi.Router {
	r := chi.NewRouter()

	r.Use(cors.Handler(cors.Options{
		// Permite qualquer porta do localhost
		AllowedOrigins:   []string{"http://localhost:*", "http://127.0.0.1:*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Location", "Link", "X-Total-Count"},
		AllowCredentials: true,
		MaxAge:           300,
	}))

	r.Mount("/clientes", UsuariosRouter())

	r.Mount("/operadores", OperadoresRouter())

	r.Mount("/administradores", AdministradoresRouter())

	r.Mount("/encomendas", PedidosRouter())

	r.Mount("/notificacoes", NotificacoesRouter())

	r.Mount("/galpoes", GalpoesRouter())

	return r
}
