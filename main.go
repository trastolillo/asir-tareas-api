package main

import (
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gorilla/mux"
)

func main() {
	// Ping database
	bd, err := getDB()
	if err != nil {
		log.Printf("Error con la base de datos: " + err.Error())
		return
	} else {
		err = bd.Ping()
		if err != nil {
			log.Printf("Error conectando a la BD. Revisa las credenciales. Error: " + err.Error())
			return
		}
	}
	// Define routes
	router := mux.NewRouter()
	setupRoutesForTareas(router)
	// Setup and start server
	port := os.Getenv("PORT")
	if port == "" {
		port = ":4000"
	}
	server := &http.Server{
		Handler: router,
		Addr:    port,
		// timeouts so the server nevers waits forever...
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
	log.Printf("Servidor conectado en puerto %s", port)
	log.Fatal(server.ListenAndServe())
}
