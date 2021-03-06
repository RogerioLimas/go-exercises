package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)


func QueryHandler(w http.ResponseWriter, r *http.Request) {
	queryParams := r.URL.Query()
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Got parameter id: %s!\n", queryParams["id"][0])
	fmt.Fprintf(w, "Got parameter category: %s!", queryParams["category"][0])
}

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/articles", QueryHandler)
	r.Queries("id", "category")

	srv := &http.Server{
		Handler: r,
		Addr: ":8080",
		ReadTimeout: 10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	log.Fatal(srv.ListenAndServe())
}