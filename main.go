package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/satwikcoder007/goserver/handlers"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", handlers.ServeHome).Methods("GET")
	r.HandleFunc("/allcources", handlers.GetAllCourses).Methods("GET")
	r.HandleFunc("/onecource", handlers.GetOneCource)
	r.HandleFunc("/addcource", handlers.AddOneCource).Methods("POST")
	log.Fatal(http.ListenAndServe(`:4000`, r))
}
