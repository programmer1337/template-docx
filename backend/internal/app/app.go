package app

import (
	"document-parser/internal/handler"
	"document-parser/internal/server"
	"log"
	"os"

	"github.com/joho/godotenv"
)

func init() {
	// loads values from .env into the system
	envFile := "../.env"
	if err := godotenv.Load(envFile); err != nil {
		log.Printf("No %v file found", envFile)
	}
}

func Run() {
	logger := NewLogger()

	handler := handler.NewHandler(logger)
	server := server.NewServer(handler, logger)

	server.Run()
}

func NewLogger() *log.Logger {
	return log.New(os.Stdout, "develop: ", log.LstdFlags)
}
