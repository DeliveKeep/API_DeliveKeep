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

		r.Get("/", controllers.BuscarUsuarios)

		r.Get("/me", controllers.BuscarLogado)

		r.Get("/{id}", controllers.BuscarUsuario)

		r.Patch("/nome", controllers.AtualizarNome)

		r.Patch("/endereco", controllers.AtualizarEndereco)

		r.Patch("/telefone", controllers.AtualizarTelefone)

		r.Patch("/email", controllers.AtualizarEmail)

		r.Patch("/senha", controllers.AtualizarSenha)

		r.Delete("/me", controllers.DeletarUsuario)
	})

	return r
}
