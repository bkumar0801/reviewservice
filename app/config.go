package app

import (
	"encoding/json"
	"log"
	"os"
)

/*
configuration ...
*/
type configuration struct {
	Server string
}

/*
Config ...
*/
var Config configuration

/*
LoadConfiguration ...
*/
func loadConfiguration() {
	file, err := os.Open("app/config.json")
	defer file.Close()
	if err != nil {
		log.Fatalf("[loadConfig]: %s\n", err)
	}
	decoder := json.NewDecoder(file)
	Config = configuration{}
	err = decoder.Decode(&Config)
	if err != nil {
		log.Fatalf("[logAppConfig]: %s\n", err)
	}
}
