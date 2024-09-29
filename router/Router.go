package router

import (
	"github.com/gorilla/mux"
	controller "github.com/ranjanvivesh/topmedia/Controller"
)

func Router() *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/api/titles",controller.GetAllTitles).Methods("GET")
	router.HandleFunc("/api/title",controller.CreateTitle).Methods("POST")
	router.HandleFunc("/api/title/{id}",controller.MarkAsWatched).Methods("PUT")
	router.HandleFunc("/api/title/{id}",controller.DeleteATitle).Methods("DELETE")
	router.HandleFunc("/api/deleteallmovies",controller.DeleteAllTitles).Methods("DELETE")	
	
	return router
}