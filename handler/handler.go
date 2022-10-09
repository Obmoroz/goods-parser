package handler

import (
	"fmt"
	"github.com/gorilla/mux"
	"goods-parser/file_system"
	"net/http"
	"regexp"
)

type Handler struct {
}

func (h *Handler) InitRoutes() *mux.Router {

	router := mux.NewRouter()

	TokenSubRouter := router.PathPrefix("/").Subrouter()
	TokenSubRouter.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		read0, err := file_system.Read0("C:\\Users\\User\\go\\bin\\goods-parser\\html\\list1.html")
		if err != nil {
			fmt.Println("error " + err.Error())
			return
		}
		writer.Write(read0)

	}).Methods("GET")
	TokenSubRouter.HandleFunc("/search", func(writer http.ResponseWriter, request *http.Request) {
		err := request.ParseForm()
		if err != nil {
			fmt.Println("error " + err.Error())
			return
		}
		s2 := request.PostForm.Get("RadioModel")
		if s2 == "" {
			writer.Write([]byte("Empty String"))
			return
		}
		if len([]rune(s2)) > 50 {
			writer.Write([]byte("String must be less then 50 symbols"))
			return
		}
		matched, _ := regexp.MatchString(`^a.b$`, "aaxbb")
		fmt.Println(s2)
		writer.Write([]byte("hellow search"))
	}).Methods("POST")

	return router

}
