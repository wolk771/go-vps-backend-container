package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	// Port aus Umgebungsvariable lesen, Standard ist 9001
	port := os.Getenv("APP_PORT")
	if port == "" {
		port = "9001"
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "EduBookingHub Pilot: Go-Backend ist bereit.\nStatus: Operationell via Jenkins CI/CD.\nUmgebung: %s", os.Getenv("APP_ENV"))
	})

	fmt.Printf("Server startet auf Port %s...\n", port)
	http.ListenAndServe(":"+port, nil)
}
