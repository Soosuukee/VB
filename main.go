package main

import (
	"fmt"
	"log"
	"net/http"

	"vb/api"
	"vb/internal/database"
)

func main() {
	// Connexion à PostgreSQL
	database.Connect()

	// Routes
	http.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "pong")
	})

	http.HandleFunc("/events", api.GetEvents)

	// Démarrage serveur
	fmt.Println("Serveur lancé sur http://localhost:8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}
