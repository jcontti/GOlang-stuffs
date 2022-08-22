package main

import (
	"encoding/json"
	"fmt"
	"go-microservices/details"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

func healthHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("Checking application health")

	// serving the response as JSON
	response := map[string]string{
		"status":    "UP",
		"timestamp": time.Now().String(),
	}

	json.NewEncoder(w).Encode(response)
}

func rootHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("Serving the Homepage")
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Application is up and running")
}

func detailsHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("Fetching the details... ")
	hostname, err := details.GetHostName()
	if err != nil {
		panic(err)
	}

	IP, _ := details.GetIp()
	fmt.Println(hostname, IP)

	response := map[string]string{
		"Hostname":   hostname,
		"timestamp":  time.Now().String(),
		"IP address": IP.String(),
	}

	json.NewEncoder(w).Encode(response)

}

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/health", healthHandler)
	r.HandleFunc("/", rootHandler)
	r.HandleFunc("/details", detailsHandler)

	log.Println("Server has started :)")
	log.Fatal(http.ListenAndServe(":80", r))
}

/* NOTES:
- Using routing (gorilla/mux package ): the biggest strenght of the gorilla/mux Router is the ability to extract segments from the request URL.
*/
