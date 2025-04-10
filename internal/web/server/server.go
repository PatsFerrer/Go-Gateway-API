// server é o servidor que vai rodar a aplicação
package server

import (
	"net/http"

	"github.com/go-chi/chi/v5" // chi é um router HTTP para Go, facilitando a criação de rotas e middlewares
	"github.com/patsferrer/go-gateway/internal/service"
	"github.com/patsferrer/go-gateway/internal/web/handlers"
	"github.com/patsferrer/go-gateway/internal/web/middleware"
)

// Server representa o servidor da aplicação
type Server struct {
	// chi é um router para o go
	router         *chi.Mux
	server         *http.Server
	accountService *service.AccountService
	invoiceService *service.InvoiceService
	port           string
}

func NewServer(accountService *service.AccountService, invoiceService *service.InvoiceService, port string) *Server {
	return &Server{
		router:         chi.NewRouter(),
		accountService: accountService,
		invoiceService: invoiceService,
		port:           port,
	}
}

func (s Server) ConfigureRoutes() {
	// cria o handler para as contas
	accountHandler := handlers.NewAccountHandler(s.accountService)
	invoiceHandler := handlers.NewInvoiceHandler(s.invoiceService)
	authMiddleware := middleware.NewAuthMiddleware(s.accountService)

	// define as rotas para operações de conta (criação e recuperação)
	s.router.Post("/accounts", accountHandler.Create)
	s.router.Get("/accounts", accountHandler.Get)

	// grupo de rotas protegidas por autenticação
	s.router.Group(func(r chi.Router) {
		r.Use(authMiddleware.Authenticate)
		s.router.Post("/invoice", invoiceHandler.Create)
		s.router.Get("/invoice/{id}", invoiceHandler.GetById)
		s.router.Get("/invoice", invoiceHandler.ListByAccount)
	})

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
