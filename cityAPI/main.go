package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type city struct {
	Name string
	Area uint64
}

func postHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		var tempCity city
		decoder := json.NewDecoder(r.Body)
		err := decoder.Decode(&tempCity)
		if err != nil {
			panic(err)
		}
		defer r.Body.Close()

		fmt.Printf("Got %s city with area of %d sq miles!\n", tempCity.Name, tempCity.Area)

		w.WriteHeader(http.StatusOK)
		io.WriteString(w, "201 - Created")
	} else {
		w.WriteHeader(http.StatusBadRequest)
		io.WriteString(w, "405 - Method Not Allowed")
	}
}

func main() {
	http.HandleFunc("/", postHandler)

	http.ListenAndServe(":8080", nil)

}
