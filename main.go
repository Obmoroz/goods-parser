package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"goods-parser/file_system"
	"goods-parser/handler"

	"gopkg.in/yaml.v3"
	"log"
	"net/http"
	"time"
)

type App struct {
	httpServer *http.Server
	Router     *mux.Router
}

type Config struct {
	SiteList []Site `yaml:"SiteList"`
}

type SiteList struct {
	Site []Site `yaml:"Site"`
}

type Site struct {
	Url      string `yaml:"Url"`
	DataPath string `yaml:"DataPath"`
}

func main() {
	/*
		поставить пакет yaml 3 +
		прочитать файл yaml  +
		сделать unmarshl в слайс структуру +
		результат переменная со срезом, содержащим настройки сайтов/ресурсов для поискать +
	*/
	readlist, err := file_system.Read0("C:\\Users\\User\\go\\bin\\goods-parser\\conf.yaml")
	if err != nil {
		fmt.Println("error " + err.Error())
		return
	}

	l := Config{}
	err = yaml.Unmarshal([]byte(readlist), &l)
	if err != nil {
		log.Fatalf("cannot unmarshal data: %v", err)
	}
	fmt.Printf("%v", l)
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
