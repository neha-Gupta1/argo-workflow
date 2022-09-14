package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type lastname struct {
	LastName string `json:"last_name"`
}

// Handler is the entry point for this fission function
func PostLastNameHandler(w http.ResponseWriter, r *http.Request) { // nolint:unused,deadcode
	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Error reading request body",
			http.StatusInternalServerError)
	}
	lastname := lastname{}
	json.Unmarshal(body, &lastname)
	fmt.Println("hey", lastname.LastName)
	_, err = w.Write([]byte(lastname.LastName))
	if err != nil {
		http.Error(w, "Error writing response", http.StatusInternalServerError)
	}
}
