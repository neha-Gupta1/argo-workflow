package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type firstname struct {
	FirstName string `json:"first_name"`
}

// Handler is the entry point for this fission function
func PostFirstNameHandler(w http.ResponseWriter, r *http.Request) { // nolint:unused,deadcode
	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Error reading request body",
			http.StatusInternalServerError)
	}
	firstname := firstname{}
	json.Unmarshal(body, &firstname)
	fmt.Println(firstname.FirstName)
	_, err = w.Write([]byte(firstname.FirstName))
	if err != nil {
		http.Error(w, "Error writing response", http.StatusInternalServerError)
	}
}

// ErrorHandler is the entry point for this fission function
func ErrorHandler(w http.ResponseWriter, r *http.Request) { // nolint:unused,deadcode

	http.Error(w, "Error reading request body", http.StatusBadRequest)
}
