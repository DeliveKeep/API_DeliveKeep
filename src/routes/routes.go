package routes

import (
	"API/src/controllers"
	"API/src/middlewares"
	"github.com/go-chi/chi"
)

// Rotear adiciona as rotas da api ao roteador
func Rotear() chi.Router {
	r := chi.NewRouter()

	// Monta as rotas públicas de /usuarios (como criar usuário e login)
	r.Mount("/usuarios", UsuariosRouter())

	// Monta TODAS as rotas de /perfil, já protegidas por autenticação
	r.Mount("/perfil", PerfilRouter())

	r.Mount("/encomendas", EncomendasRouter())

	return r
}

// PerfilRouter cria as rotas para o perfil do usuário
// Todas as rotas aqui dentro já exigem autenticação
func PerfilRouter() chi.Router {
	r := chi.NewRouter()

	// O middleware de autenticação é aplicado a TODAS as rotas deste grupo
	r.Use(middlewares.Autenticar)

	// A rota de notificações, agora com a sintaxe correta do chi
	r.Put("/notificacoes", controllers.AtualizarConfigNotificacoes)

	// Você pode adicionar outras rotas de perfil aqui no futuro
	// r.Get("/", controllers.BuscarLogado)
	// r.Put("/atualizar-nome", controllers.AtualizarNome)

	return r
}

func EncomendasRouter() chi.Router {
	r := chi.NewRouter()

	// Protege todas as rotas de encomendas com autenticação
	r.Use(middlewares.Autenticar)

	// Define a rota para CRIAR uma encomenda (POST /encomendas)
	r.Post("/", controllers.CriarEncomenda)

	// No futuro, você pode adicionar outras rotas aqui:
	// r.Get("/", controllers.BuscarEncomendasDoUsuario)
	// r.Get("/{encomendaID}", controllers.BuscarEncomenda)

	return r
}