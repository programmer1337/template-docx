package server

import (
	"context"
	"document-parser/internal/config"
	"log"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/rs/cors"
)

type Server struct {
	httpServer *http.Server
	logger     *log.Logger
}

func NewServer(handler http.Handler, logger *log.Logger) *Server {
	eOrigin := os.Getenv("ORIGIN")
	origins := []string{}
	for _, str := range strings.Split(eOrigin, ",") {
		origins = append(origins, strings.TrimSpace(str))
	}

	return &Server{
		httpServer: &http.Server{
			Addr:    ":" + config.DefaultHTTPPort,
			Handler: configureCORSFor(handler, origins),
			// Handler:      handler,
			IdleTimeout:  30 * time.Second,
			WriteTimeout: 30 * time.Second,
			ReadTimeout:  30 * time.Second,
		},
		logger: logger,
	}
}

// func configureCORS() *cors.Cors {
// 	return cors.New(cors.Options{
// 		// # http://mywebsite-domain.com/ is configured in hosts (localhost:80 alias)
// 		AllowedOrigins: []string{"http://mywebsite-domain.com/"},
// 		AllowedMethods: []string{"POST", "GET", "PUT", "DELETE", "OPTIONS"},
// 		// AllowCredentials: true,

// 		Debug: true,
// 	})
// }

func configureCORSFor(handler http.Handler, origins []string) http.Handler {
	ch := cors.New(cors.Options{
		// # http://mywebsite-domain.com/ is configured in hosts (localhost:80 alias)
		AllowedOrigins: origins,
		AllowedMethods: []string{"POST", "GET", "PUT", "DELETE", "OPTIONS"},
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
