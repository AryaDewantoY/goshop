package controllers

import (
	"net/http"

	"github.com/gorilla/mux"
)

func (routes *Server) initializeRoutes() {
	routes.Router = mux.NewRouter()
	routes.Router.HandleFunc("/", routes.Home).Methods("GET")
	routes.Router.HandleFunc("/Products", routes.Products).Methods("GET")



	staticFileDirectory := http.Dir("./assets/")
	staticFileHandler := http.StripPrefix("/public/", http.FileServer(staticFileDirectory))
	routes.Router.PathPrefix("/public/").Handler(staticFileHandler).Methods("GET")
}
