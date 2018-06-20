package main

import (
	"log"
	"net/http"

	"github.com/reviewservice/app"
	"github.com/reviewservice/routers"
)

func main() {
	app.Boot()
	router := routers.Init()
	server := &http.Server{
		Addr:    app.Config.Server,
		Handler: router,
	}
	log.Fatal(server.ListenAndServe())
}
