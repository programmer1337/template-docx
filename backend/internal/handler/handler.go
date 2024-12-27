package handler

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func NewHandler(logger *log.Logger) http.Handler {
	serveMux := mux.NewRouter()

	HandleReplace(serveMux, logger)
	HandleCounterparty(serveMux, logger)

	return serveMux
}
