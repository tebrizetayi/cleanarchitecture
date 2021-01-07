package api

import (
	"github.com/gorilla/mux"
)

func InitRoute(authorHandler AuthorHandler) *mux.Router {

	r := mux.NewRouter()
	r.HandleFunc("/author", authorHandler.Get)

	return r
}
