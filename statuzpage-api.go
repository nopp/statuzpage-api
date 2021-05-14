package main

import (
	"log"
	"net/http"
	"statuzpage-api/incidents"
	"statuzpage-api/urls"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

func main() {

	router := mux.NewRouter()
	// URL methods
	router.HandleFunc("/urls", urls.GetUrls).Methods("GET")
	router.HandleFunc("/url/{id}", urls.GetUrl).Methods("GET")
	router.HandleFunc("/url", urls.CreateUrl).Methods("POST")
	router.HandleFunc("/url/{id}", urls.DeleteUrl).Methods("DELETE")
	// Incident methods
	router.HandleFunc("/incident", incidents.CreateIncident).Methods("POST")
	router.HandleFunc("/incident/{id}", incidents.CloseIncident).Methods("POST")
	router.HandleFunc("/incidents", incidents.GetIncidents).Methods("GET")
	router.HandleFunc("/incidentsclosed", incidents.GetIncidentsClosed).Methods("GET")

	log.Fatal(http.ListenAndServe("0.0.0.0:8000", router))
}
