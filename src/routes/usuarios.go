package routes

import (
	"API/src/controllers"
	"API/src/middlewares"

	"github.com/go-chi/chi"
)

// UsuariosRouters retorna roteador com rotas /usuarios
func UsuariosRouter() chi.Router {
	r := chi.NewRouter()

	r.Post("/", controllers.CriarUsuario)

	r.Post("/login", controllers.Login)

	r.Group(func(r chi.Router) {
		r.Use(middlewares.Autenticar)

		r.Get("/me", controllers.BuscarLogado)

		r.Patch("/celular", controllers.AtualizarCelular)

		r.Patch("/email", controllers.AtualizarEmail)

		r.Patch("/senha", controllers.AtualizarSenha)
	})

	return r
}
