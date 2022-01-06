package main

import (
	"io"
	"log"
	"net/http"
	"time"
)

func HealtCheck(w http.ResponseWriter, r *http.Request) {
	currentTime := time.Now()
	w.Write([]byte(currentTime.String()))
	io.WriteString(w, currentTime.String())
}

func main() {
	port := ":8080"
	http.HandleFunc("/health", HealtCheck)
	log.Fatal(http.ListenAndServe(port, nil))
}