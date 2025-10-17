package routes

import (
	"API/src/controllers"
	"API/src/middlewares"

	"github.com/go-chi/chi"
)

// retorna roteador com rotas /operadores
func OperadoresRouter() chi.Router {
	r := chi.NewRouter()

	r.Post("/", controllers.CriarOperador)

	r.Post("/login", controllers.LoginOperador)

	r.Group(func(r chi.Router) {
		r.Use(middlewares.Autenticar)

		r.Get("/", controllers.BuscarOperadores)

		r.Get("/me", controllers.BuscarOperadorLogado)

		r.Get("/{id}", controllers.BuscarOperador)

		r.Patch("/nome", controllers.AtualizarNomeOperador)

		r.Patch("/telefone", controllers.AtualizarTelefoneOperador)

		r.Patch("/email", controllers.AtualizarEmailOperador)

		r.Patch("/senha", controllers.AtualizarSenhaOperador)

		r.Delete("/me", controllers.DeletarOperador)
	})

	return r
}
