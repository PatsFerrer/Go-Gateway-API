// server é o servidor que vai rodar a aplicação
package server

import (
	"net/http"

	"github.com/go-chi/chi/v5" // chi é um router HTTP para Go, facilitando a criação de rotas e middlewares
	"github.com/patsferrer/go-gateway/internal/service"
	"github.com/patsferrer/go-gateway/internal/web/handlers"
)

// Server representa o servidor da aplicação
type Server struct {
	// chi é um router para o go
	router         *chi.Mux
	server         *http.Server
	accountService *service.AccountService
	port           string
}

func NewServer(accountService *service.AccountService, port string) *Server {
	return &Server{
		router:         chi.NewRouter(),
		accountService: accountService,
		port:           port,
	}
}

func (s Server) ConfigureRoutes() {
	// cria o handler para as contas
	accountHandler := handlers.NewAccountHandler(s.accountService)

	// define as rotas para operações de conta (criação e recuperação)
	s.router.Post("/accounts", accountHandler.Create)
	s.router.Get("/accounts", accountHandler.Get)

	// pode ser feito dessa forma também:
	// s.router.Route("/accounts", func(r chi.Router) {
	// 	r.Post("/", accountHandler.Create)
	// 	r.Get("/", accountHandler.Get)
	// })
}

func (s Server) Start() error {
	// cria o servidor
	s.server = &http.Server{
		Addr:    ":" + s.port,
		Handler: s.router,
	}

	// inicia o servidor
	return s.server.ListenAndServe()
}
