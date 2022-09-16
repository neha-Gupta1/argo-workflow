package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
)

// Handler is the entry point for this fission function
func AddNumberHandler(w http.ResponseWriter, r *http.Request) {
	number := r.URL.Query().Get("number")
	// name := stringutil.Reverse(stringutil.Reverse("Vendor Example Test"))
	no, err := strconv.Atoi(number)
	if err != nil {
		log.Printf("Error writing response: %v", err)
	}
	no = no + 1
	_, err = w.Write([]byte(fmt.Sprint(no)))
	if err != nil {
		log.Printf("Error writing response: %v", err)
	}
}
