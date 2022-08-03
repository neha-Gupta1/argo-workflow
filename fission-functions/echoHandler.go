package main

import (
	"log"
	"net/http"
)

// Handler is the entry point for this fission function
func Handler(w http.ResponseWriter, r *http.Request) {
	echoValue := r.URL.Query().Get("echoValue")
	// name := stringutil.Reverse(stringutil.Reverse("Vendor Example Test"))
	_, err := w.Write([]byte(echoValue))
	if err != nil {
		log.Printf("Error writing response: %v", err)
	}
}
