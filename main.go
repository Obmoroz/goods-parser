package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"goods-parser/handler"
	"net/http"
	"time"
)

type App struct {
	httpServer *http.Server
	Router     *mux.Router
}

func main() {
	app := NewApp()
	if err := app.Run(); err != nil {
		fmt.Println(err)

	}

}

func NewApp() *App {

	h := handler.Handler{}
	router := h.InitRoutes()

	return &App{
		Router: router,
	}
}

func (a *App) Run() error {
	// HTTP Server

	a.httpServer = &http.Server{
		Handler:      a.Router,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
		Addr:         "127.0.0.1:8080",
	}

	err := a.httpServer.ListenAndServe()
	if err != nil {
		return err
	}

	return nil
}
