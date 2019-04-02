package main

import (
	"net/http"
	"os"

	applicationRoute "golang_keycloak_REST_interface/router"

	"github.com/julienschmidt/httprouter"
	"github.com/rs/cors"
	log "github.com/sirupsen/logrus"
)

func init() {
	// Log as JSON instead of the default ASCII formatter.
	log.SetFormatter(&log.JSONFormatter{})
	log.SetOutput(os.Stdout)
}

func main() {
	router := httprouter.New()
	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:8081", "http://127.0.0.1:8081"},
		AllowCredentials: false,
		Debug:            false,
		AllowedMethods:   []string{"GET", "POST", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"*"},
	})
	handler := c.Handler(router)
	log.Printf("Starting service. \n Listening at 0.0.0.0:8081")
	applicationRoute.ApplicationRoutes(router)
	log.Fatal(http.ListenAndServe(":8081", handler))
}
