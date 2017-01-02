package routers

import (
    "github.com/gorilla/mux"
)

func InitRoutes() *mux.Router {
    router := mux.NewRouter()
    router = SetUserRoutes(router)
    router = SetTaskRoutes(router)
    router = SetNoteRoutes(router)
    return router
}
