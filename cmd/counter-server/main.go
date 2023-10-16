package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/Niproblema/NU-Task/api/requests"
	"github.com/Niproblema/NU-Task/api/responses"
	"github.com/Niproblema/NU-Task/internal/counter"
)

func handleRequest(w http.ResponseWriter, r *http.Request) {
	// Check if the Content-Type header is set to application/json
	if r.Header.Get("Content-Type") != "application/json" {
		http.Error(w, "Content-Type header is not application/json", http.StatusUnsupportedMediaType)
		return
	}

	// Read the body of the request
	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()
	var req requests.Request
	err := decoder.Decode(&req)
	if err != nil {
		log.Printf("Error parsing the request; [%s]", err.Error())
		http.Error(w, "Error parsing the request.", http.StatusBadRequest)
		return
	}
	log.Printf("Received message with body; [%+v]", req)

	var totalCount = counter.CountRepository(req.Directory, req.Word, req.Case, req.Whole)

	// Prepare response
	resp := responses.Response{Count: totalCount}
	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(resp)
	if err != nil {
		log.Printf("Error preparing the response; [%s]", err.Error())
		http.Error(w, fmt.Sprintf("Error preparing the response: %v", err), http.StatusInternalServerError)
		return
	}
	log.Printf("Responded to request with response; [%+v]", resp)
}

func main() {
	var port = 8080
	var addressPattern = "/counter"

	log.Printf("Starting service at http://localhost:%d/%s", port, addressPattern)
	http.HandleFunc(addressPattern, handleRequest)
	http.ListenAndServe(fmt.Sprintf(":%d", port), nil)
}
