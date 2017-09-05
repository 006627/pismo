package common

import (
	"encoding/json"
	"log"
	"os"
)

type configuration struct {
	Server,
	Host,
	MongoDBHost,
	DBUser,
	DBPwd,
	Database string
}

// AppConfig holds the configuration values from config.json file
var AppConfig configuration

// Initialize AppConfig
func initConfig() {
	file, err := os.Open("common/config.json")
	defer file.Close()

	if err != nil {
		log.Fatalf("[Config] -> Error on Open File: %s\n", err)
	}

	decoder := json.NewDecoder(file)
	AppConfig = configuration{}

	err = decoder.Decode(&AppConfig)

	if err != nil {
		log.Fatalf("[Config] -> Error create configuration: %s\n", err)
	}
}
