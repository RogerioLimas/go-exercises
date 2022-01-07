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

	fmt.Fprintf(w, queryParams["id"][0])
}

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/articles", QueryHandler).Name()

	srv := &http.Server{
		Handler: r,
		Addr: ":8080",
		ReadTimeout: 10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	log.Fatal(srv.ListenAndServe())
}