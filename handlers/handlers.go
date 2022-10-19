package handlers

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/drc288/contable/routes"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

func Manage(version string) {
	router := mux.NewRouter()

	c := fmt.Sprintf("/api/contable/%s/create", version)

	router.HandleFunc(c, routes.Create).Methods("POST")

	handler := cors.AllowAll().Handler(router)

	PORT := os.Getenv("PORT")
	if PORT == "" {
		PORT = "8080"
	}

	log.Printf("using: http://localhost:%s/\n", PORT)
	log.Fatal(http.ListenAndServe(":"+PORT, handler))
}
