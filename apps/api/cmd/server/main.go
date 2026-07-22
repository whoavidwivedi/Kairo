package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
)



func main() {
	router := chi.NewRouter()

	leadRepository := &LeadRepository{}
	leadService := NewLeadService(leadRepository)
	
	router.Get("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		response := map[string]string{
			"status":  "healthy",
			"service": "kairo-api",
		}

		json.NewEncoder(w).Encode(response)
	})

	router.Post("/leads", func(w http.ResponseWriter, r *http.Request) {
		var req CreateLeadRequest

		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			http.Error(w, "Invalid request body", http.StatusBadRequest)
			return
		}

		response := leadService.CreateLead(req)
		w.Header().Set("Content-Type", "application/json")

		json.NewEncoder(w).Encode(response)

	})

	router.Get("/leads", func(w http.ResponseWriter, r *http.Request) {
		leads := leadService.ListLeads()

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(leads)
	})

	log.Println("Kairo API listening on http://localhost:8080")

	if err := http.ListenAndServe(":8080", router); err != nil {
		log.Fatal(err)
	}
}