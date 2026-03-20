package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hallo! Dieser Go-Server nutzt Port 9001 und läuft stabil auf dem VPS! An internes Network angebunden. Jenkins Webhook Test...")
	})

	// Port 9001 für VPS
	port := ":9001"
	fmt.Println("Go-App startet auf Port", port)
	http.ListenAndServe(port, nil)
}
