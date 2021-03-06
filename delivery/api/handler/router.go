package api

import (
	"net/http"

	"github.com/gorilla/mux"
)

func InitRoutes(ah *AuthorHandler) *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/api/v1/author/{id}", ah.Get).Methods(http.MethodGet)
	r.HandleFunc("/api/v1/author/", ah.Get).Methods(http.MethodGet)
	r.HandleFunc("/api/v1/author/", ah.Create).Methods(http.MethodPost)
	r.HandleFunc("/api/v1/author/{id}", ah.Delete).Methods(http.MethodDelete)
	r.HandleFunc("/api/v1/author/{id}", ah.Update).Methods(http.MethodPut)
	return r
}
