package server

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/devfullcycle/imersao22/go-gateway/internal/service"
	"github.com/devfullcycle/imersao22/go-gateway/internal/web/handlers"
	"github.com/go-chi/chi/v5"
)

type Server struct {
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

func (s *Server) ConfigureRoutes() {
	accountHandler := handlers.NewAccountHandler(s.accountService)
	s.router.Post("/accounts", accountHandler.Create)
	s.router.Get("/accounts", accountHandler.Get)
}

func (s *Server) Start() error {
	s.server = &http.Server{
		Addr:    ":" + s.port,
		Handler: s.router,
	}

	// Canal para capturar erros do servidor
	serverErrors := make(chan error, 1)

	// Inicia o servidor em uma goroutine separada
	go func() {
		log.Printf("Servidor iniciado na porta %s", s.port)
		serverErrors <- s.server.ListenAndServe()
	}()

	// Canal para capturar sinais do sistema operacional
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)

	// Aguarda um sinal ou erro do servidor
	select {
	case err := <-serverErrors:
		return err
	case sig := <-sigChan:
		log.Printf("Sinal recebido: %v. Iniciando shutdown gracioso...", sig)

		// Cria um contexto com timeout para o shutdown
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		// Inicia o shutdown gracioso
		if err := s.server.Shutdown(ctx); err != nil {
			log.Printf("Erro durante o shutdown gracioso: %v", err)
			return err
		}

		log.Println("Shutdown gracioso concluído.")
	}

	return nil
}
