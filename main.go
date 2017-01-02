package main

import (
	"log"
	"net/http"

	"github.com/codegangsta/negroni"
	"github.com/gabsn/taskmanager/common"
	"github.com/gabsn/taskmanager/routers"
)

func main() {
	common.StartUp()
	router := routers.InitRoutes()
	n := negroni.Classic()
	n.UseHandler(router)

	server := &http.Server{
		Addr:    common.AppConfig.Server,
		Handler: n,
	}
	log.Println("Listening...")
	server.ListenAndServe()
}
