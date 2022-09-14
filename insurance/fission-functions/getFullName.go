package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type fullname struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}

// Handler is the entry point for this fission function
func GetFullNameHandler(w http.ResponseWriter, r *http.Request) { // nolint:unused,deadcode
	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Error reading request body",
			http.StatusInternalServerError)
	}
	name := fullname{}
	json.Unmarshal(body, &name)
	fmt.Println("hey", name.FirstName+" "+name.LastName)
	_, err = w.Write([]byte(name.FirstName + " " + name.LastName))
	if err != nil {
		http.Error(w, "Error writing response", http.StatusInternalServerError)
	}
}
