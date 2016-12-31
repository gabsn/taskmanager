package routers

import (
    "github.com/codegangsta/negroni"
    "github.com/gorilla/mux"
    "github.com/gabsn/taskmanager/common"
    "github.com/gabsn/taskmanager/controllers"
)

func SetTaskRoutes(router *mux.Router) *mux.Router {
    taskRouter := mux.NewRouter()
    noteRouter.HandleFunc("/notes", controllers.CreateNote).Methods("POST")
    noteRouter.HandleFunc("/notes/{id}", controllers.UpdateNote).Methods("PUT")
    noteRouter.HandleFunc("/notes", controllers.GetNotes).Methods("GET")
    noteRouter.HandleFunc("/notes/{id}", controllers.GetNoteById).Methods("GET")
    noteRouter.HandleFunc("/notes/tasks/{id}", controllers.GetNoteByTask).Methods("GET")
    noteRouter.HandleFunc("/notes/{id}", controllers.DeleteNote).Methods("DELETE")
    router.PathPrefix("/notes").Handler(negroni.New(
        negroni.HandlerFunc(common.Authorize),
        negroni.Wrap(noteRouter),
    ))
    return router
}
