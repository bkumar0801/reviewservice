package main

import (
	"log"
	"net/http"

	"gopkg.in/mgo.v2"

	"github.com/reviewservice/app"
	"github.com/reviewservice/routers"
)

/*
AppContext ...
*/
type AppContext struct {
	session *mgo.Session
}

func main() {
	config := app.GetConf("app/config.json")
	router := routers.Init()
	server := &http.Server{
		Addr:    config.Server,
		Handler: router,
	}
	log.Fatal(server.ListenAndServe())
}
