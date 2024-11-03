package routes

import (
    "gestaotalentos/controllers"
    "gestaotalentos/middlewares"
    "net/http"

    "github.com/gorilla/mux"
)

func SetupRoutes() *mux.Router {
    router := mux.NewRouter()

    router.HandleFunc("/register", controllers.Register).Methods("POST")
    router.HandleFunc("/login", controllers.Login).Methods("POST")

    router.Handle("/candidatos", middlewares.AuthMiddleware(http.HandlerFunc(controllers.GetCandidatos))).Methods("GET")
    router.Handle("/candidatos", middlewares.AuthMiddleware(http.HandlerFunc(controllers.CreateCandidato))).Methods("POST")
    router.Handle("/candidatos/{id}", middlewares.AuthMiddleware(http.HandlerFunc(controllers.GetCandidato))).Methods("GET")
    router.Handle("/candidatos/{id}", middlewares.AuthMiddleware(http.HandlerFunc(controllers.UpdateCandidato))).Methods("PUT")
    router.Handle("/candidatos/{id}", middlewares.AuthMiddleware(middlewares.AdminMiddleware(http.HandlerFunc(controllers.DeleteCandidato)))).Methods("DELETE")

    return router
}