package app

import (
	"encoding/json"
	"log"
	"os"
)

/*
Configuration ...
*/
type Configuration struct {
	Server      string
	MongoDBHost string
	DBUser      string
	DBPwd       string
	Database    string
}

/*
GetConf ...
*/
func GetConf(filename string) *Configuration {
	file, err := os.Open(filename)
	defer file.Close()
	if err != nil {
		log.Fatalf("[loading configuration]: %s\n", err)
	}
	decoder := json.NewDecoder(file)
	config := &Configuration{}
	err = decoder.Decode(&config)
	if err != nil {
		log.Fatalf("[loading application config]: %s\n", err)
	}
	return config
}
