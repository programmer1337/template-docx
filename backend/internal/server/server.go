package server

import (
	"context"
	"document-parser/internal/config"
	"log"
	"net/http"
	"time"

	"github.com/rs/cors"
)

type Server struct {
	httpServer *http.Server
	logger     *log.Logger
}

func NewServer(handler http.Handler, logger *log.Logger) *Server {
	return &Server{
		httpServer: &http.Server{
			Addr:    ":" + config.DefaultHTTPPort,
			Handler: configureCORSFor(handler),
			// Handler:      handler,
			IdleTimeout:  120 * time.Second,
			WriteTimeout: 2 * time.Second,
			ReadTimeout:  2 * time.Second,
		},
		logger: logger,
	}
}

func configureCORS() *cors.Cors {
	return cors.New(cors.Options{
		// # http://mywebsite-domain.com/ is configured in hosts (localhost:80 alias)
		AllowedOrigins: []string{"http://mywebsite-domain.com/"},
		AllowedMethods: []string{"POST", "GET", "PUT", "DELETE"},
		// AllowCredentials: true,

		Debug: true,
	})
}

func configureCORSFor(handler http.Handler) http.Handler {
	ch := cors.New(cors.Options{
		// # http://mywebsite-domain.com/ is configured in hosts (localhost:80 alias)
		AllowedOrigins: []string{"http://mywebsite-domain.com/", "http://192.168.1.107:8080", "http://localhost:8080"},
		AllowedMethods: []string{"POST", "GET", "PUT", "DELETE"},
		// AllowCredentials: true,

		Debug: true,
	})

	return ch.Handler(handler)
}

func (s *Server) Run() {
	s.logger.Printf("Server start at %v port", config.DefaultHTTPPort)

	if err := s.httpServer.ListenAndServe(); err != nil {
		s.logger.Fatalf("Cannot start. Error: %v", err)
	}
}

func (s *Server) Stop(ctx context.Context) error {
	return s.httpServer.Shutdown(ctx)
}
