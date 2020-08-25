package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/", Index)
	log.Print("Running in 8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}

// Index -
func Index(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Eu sou Full Cycle"))
}
