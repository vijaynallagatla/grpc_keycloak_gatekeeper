package features

import (
	"encoding/json"
	Model "golang_keycloak_REST_interface/model"
	"net/http"

	log "github.com/sirupsen/logrus"
)

// Sample POST Request
func samplePostRequestController(w http.ResponseWriter, r *http.Request) {
	var l Model.Logger
	if err := json.NewDecoder(r.Body).Decode(&l); err != nil {
		log.Println("Error: ", err)
	}
	defer r.Body.Close()
	if err := json.NewEncoder(w).Encode("Success"); err != nil {
		log.WithFields(log.Fields{
			"Error": "Error in sending response",
		})
	}
	log.WithFields(log.Fields{
		"AUDIT_LOG": l.AuditLog,
		"UUID":      l.UUID,
	}).Info(l.Message)
}

// Sample GET Request Controller
func sampleGETRequest(w http.ResponseWriter, r *http.Request) {
	log.WithFields(log.Fields{
		"tag":    "Request",
		"method": "GET",
	}).Info("Get request '/'")

	if err := json.NewEncoder(w).Encode("Sample Get Request!!!"); err != nil {
		log.WithFields(log.Fields{
			"Error": "Error in sending response",
		})
	}
}
