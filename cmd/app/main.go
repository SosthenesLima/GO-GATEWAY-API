package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/devfullcycle/imersao22/go-gateway/internal/repository"
	"github.com/devfullcycle/imersao22/go-gateway/internal/service"
	"github.com/devfullcycle/imersao22/go-gateway/internal/web/server"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func getEnv(key, defautValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}

	return defautValue

}

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}

	// String de conexão com o banco
	// Constrói a string de conexão com o banco de dados
	connStr := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		getEnv("DB_HOST", "db"),
		getEnv("DB_PORT", "5432"),
		getEnv("DB_USER", "postgres"),
		getEnv("DB_PASSWORD", "postgres"),
		getEnv("DB_NAME", "gateway"),
		getEnv("DB_SSL_MODE", "disable"),
	)

	// Conecta ao banco de dados
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal("Erro connecting to database: ", err)
	}
	defer db.Close()

	// Verifica a conexão com o banco
	if err := db.Ping(); err != nil {
		log.Fatal("Erro ao verificar conexão com o banco de dados: ", err)
	}
	// Inicializa o repositório
	accountRepository := repository.NewAccountRepository(db)

	// Inicializa o serviço com o repositório como dependência
	accountService := service.NewAccountService(accountRepository)
	// Configura o servidor HTTP
	port := getEnv("HTTP_PORT", "8080")
	srv := server.NewServer(accountService, port)
	srv.ConfigureRoutes()

	// Inicia o servidor
	srv.Start()

	if err := srv.Start(); err != nil {
		log.Fatal("Error starting server: ", err)
	}
}
