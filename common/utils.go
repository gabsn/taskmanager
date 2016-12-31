package common

import (
	"encoding/json"
	"log"
	"os"
)

type config struct {
	Server, MongoDBHost, DBUser, DBPwd, Database string
}

var AppConfig config

func initConfig() {
    file, err := os.Open("common/config.json")
    defer file.Close()
    if err != nil {
        log.Fatal("[initConfig]:", err)
    }
    err = json.NewDecoder(file).Decode(&AppConfig)
    if err != nil {
        log.Fatal("[initConfig]:", err)
    }
}

