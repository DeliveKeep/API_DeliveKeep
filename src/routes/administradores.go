package routes

import (
	"API/src/controllers"
	"API/src/middlewares"

	"github.com/go-chi/chi"
)

// retorna roteador com rotas /operadores
func AdministradoresRouter() chi.Router {
	r := chi.NewRouter()

	r.Post("/", controllers.CriarAdministrador)

	r.Post("/login", controllers.LoginAdministrador)

	r.Group(func(r chi.Router) {
		r.Use(middlewares.Autenticar)

		r.Get("/", controllers.BuscarAdministradores)

		r.Get("/me", controllers.BuscarAdministradorLogado)

		r.Get("/{id}", controllers.BuscarAdministrador)

		r.Patch("/nome", controllers.AtualizarNomeAdministrador)

		r.Patch("/telefone", controllers.AtualizarTelefoneAdministrador)

		r.Patch("/email", controllers.AtualizarEmailAdministrador)

		r.Patch("/senha", controllers.AtualizarSenhaAdministrador)

		r.Delete("/me", controllers.DeletarAdministrador)
	})

	return r
}
