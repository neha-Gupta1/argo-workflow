package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type salary struct {
	Salary int `json:"salary"`
}

// Handler is the entry point for this fission function
func PostSalaryHandler(w http.ResponseWriter, r *http.Request) { // nolint:unused,deadcode
	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Error reading request body",
			http.StatusInternalServerError)
	}
	salary := salary{}
	json.Unmarshal(body, &salary)
	fmt.Println(salary.Salary)
	_, err = w.Write([]byte(fmt.Sprintf("%d", salary.Salary)))
	if err != nil {
		http.Error(w, "Error writing response", http.StatusInternalServerError)
	}
}
