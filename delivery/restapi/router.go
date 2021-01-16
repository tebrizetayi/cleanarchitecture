package api

import (
	"net/http"

	"github.com/gorilla/mux"
)

func InitRoutes(ah *AuthorHandler) *mux.Router {

	r := mux.NewRouter()
	r.HandleFunc("/author", ah.Get).Methods(http.MethodGet)
	r.HandleFunc("/author", ah.Create).Methods(http.MethodPost)
	r.HandleFunc("/author/{id}", ah.Delete).Methods(http.MethodDelete)
	return r
}
