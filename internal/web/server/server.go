package server

import (
	"net/http"

	"github.com/devfullcycle/imersao22/go-gateway/internal/service"
	"github.com/devfullcycle/imersao22/go-gateway/internal/web/handlers"
	"github.com/go-chi/chi/v5"
)

type Server struct {
	router         *chi.Mux
	server         *http.Server
	accountService *server.AccountService
	port           string
}

func NewServer(acconntService *service.AccountService, port string) *Server {
	return &Server{
		router:         chi.NewRouter(),
		accountService: acconntService,
		port:           port,
	}
	
}
func (s *Server) ConfigureRoutes() {
    accountHandler := handlers.NewAccountHandler(s.accountService)

	s.router.Post("/accounts", accountHandler.Create)
	s.router.Get("/accounts", accountHandler.Get)
}

func (s *Server) Start() error {
	s.server =  &http.Server{
		Addr:   ":" + s.port, 
		Handler: s.router,
	}
	returt s.server.ListendAndServe()
	
}
