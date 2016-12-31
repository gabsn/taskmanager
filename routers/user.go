package routers

import (
    "github.com/gorilla/mux"
    "github.com/gabsn/taskmanager/controllers"
)

func SetUserRoutes(router *mux.Router) *mux.Router {
    router.HandlerFunc("/users/register", controllers.Register).Methods("POST")
    router.HandlerFunc("/users/login", controllers.Login).Methods("POST")
    return router
}
