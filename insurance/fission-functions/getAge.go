package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type age struct {
	Age int `json:"age"`
}

// Handler is the entry point for this fission function
func PostAgeHandler(w http.ResponseWriter, r *http.Request) { // nolint:unused,deadcode
	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Error reading request body",
			http.StatusInternalServerError)
	}
	age := age{}
	json.Unmarshal(body, &age)
	fmt.Println(age.Age)
	_, err = w.Write([]byte(fmt.Sprintf("%d", age.Age)))
	if err != nil {
		http.Error(w, "Error writing response", http.StatusInternalServerError)
	}
}

// ErrorHandler is the entry point for this fission function
func ErrorHandler(w http.ResponseWriter, r *http.Request) { // nolint:unused,deadcode

	http.Error(w, "Error reading request body", http.StatusBadRequest)
}
