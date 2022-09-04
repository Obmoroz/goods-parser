package handler

import (
	"github.com/gorilla/mux"
	"net/http"
)

type Handler struct {
}

func (h *Handler) InitRoutes() *mux.Router {

	router := mux.NewRouter()

	TokenSubRouter := router.PathPrefix("/").Subrouter()
	TokenSubRouter.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		writer.Write([]byte("hellow world"))
	}).Methods("GET")
	TokenSubRouter.HandleFunc("/search", func(writer http.ResponseWriter, request *http.Request) {
		writer.Write([]byte("hellow search"))
	}).Methods("GET")

	return router

}
