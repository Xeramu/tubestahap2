package main

import (
	"fmt"
	"net/http"
)

func main() {
	// Endpoint untuk Tracking Service
	http.HandleFunc("/tracking", getTrackingHandler)
	http.HandleFunc("/tracking/event", insertTrackingEventHandler)

	// Endpoint untuk Integrasi Peta (Map/Location)
	http.HandleFunc("/distance", calculateDistanceHandler)
	http.HandleFunc("/route", calculateRouteHandler)
	http.HandleFunc("/location", getCourierLocationHandler)

	fmt.Println("Tracking Service running on :8084")
	http.ListenAndServe(":8084", nil)
}