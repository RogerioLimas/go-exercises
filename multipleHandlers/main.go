package main

import (
	"fmt"
	"io"
	"log"
	"math/rand"
	"net/http"
)

func main() {
	newMux := http.NewServeMux()

	newMux.HandleFunc("/randomFloat", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println(r.RemoteAddr, r.RequestURI, r.Method)
		randomFloatNumber := rand.Float64()
		io.WriteString(w, fmt.Sprintf("%.2f", randomFloatNumber))
	})

	newMux.HandleFunc("/randomInt", func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, fmt.Sprintf("%d", rand.Intn(100)))
	})

	log.Fatal(http.ListenAndServe(":8080", newMux))

}