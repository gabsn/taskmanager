package common

import (
	"encoding/json"
	"log"
	"os"
)

type (
    config struct {
        Server, MongoDBHost, DBUser, DBPwd, Database string
    }
    appError struct {
        Error string `json:"error"`
        Message string `json:"message"`
        HttpStatus int `json:"status"`
    }
    errorResource struct {
        Data appError `json:"data"`
    }
)

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

func DisplayAppError(w http.ResponseWriter, handlerError error, message string, code int) {
    errObj := appError{handlerError.Error(), message, code}
    log.Printf("[AppError]: %s\n", handlerError)
    w.Header.Set("Content-Type", "application/json; charset=utf-8")
    w.WriteHeader(code)
    if j, err := json.Marshal(errorResource{errObj}); err == nil {
        w.Write(j)
    }
}

