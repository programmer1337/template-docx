package handler

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func NewHandler(logger *log.Logger) http.Handler {
	serveMux := mux.NewRouter()

	HandleReplace(serveMux, logger)
	// TokenHandler(serveMux, logger)

	return serveMux
}
