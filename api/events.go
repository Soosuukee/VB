package api

import (
	"fmt"
	"net/http"
)

func GetEvents(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Liste des événements à venir (fictive pour l’instant)")
}
