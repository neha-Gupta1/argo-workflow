package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type insurance struct {
	Name   string `json:"name"`
	Age    int    `json:"age"`
	Salary int    `json:"salary"`
}

// Handler is the entry point for this fission function
func GetLowRiskInsuranceHandler(w http.ResponseWriter, r *http.Request) { // nolint:unused,deadcode
	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Error reading request body",
			http.StatusInternalServerError)
	}
	insurance := insurance{}
	json.Unmarshal(body, &insurance)
	_, err = w.Write([]byte(fmt.Sprintf("%v", insurance)))
	// _, err = w.Write([]byte(fmt.Sprintf("Hello %s !! as per your age: %d We can give you an insurance with assured money %d", insurance.Name, insurance.Age, insurance.Salary*100)))
	if err != nil {
		http.Error(w, "Error writing response", http.StatusInternalServerError)
	}
}

// Handler is the entry point for this fission function
func GetHighRiskInsuranceHandler(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Error reading request body",
			http.StatusInternalServerError)
	}
	insurance := insurance{}
	json.Unmarshal(body, &insurance)
	// _, err = w.Write([]byte(fmt.Sprintf("Hello %s !! as per your age: %d We can give you an insurance with assured money %d", insurance.Name, insurance.Age, insurance.Salary*50)))
	// if err != nil {
	// 	http.Error(w, "Error writing response", http.StatusInternalServerError)
	// }

	_, err = w.Write([]byte(fmt.Sprintf("%v", insurance)))
	// _, err = w.Write([]byte(fmt.Sprintf("Hello %s !! as per your age: %d We can give you an insurance with assured money %d", insurance.Name, insurance.Age, insurance.Salary*100)))
	if err != nil {
		http.Error(w, "Error writing response", http.StatusInternalServerError)
	}

}
