package handlers

import (
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
	"github.com/tapiaw38/go-mails/routers"
)

// HandleServer handles the server request
func HandlerServer() {
	router := mux.NewRouter()
	router.HandleFunc("/email", routers.HandleEmail).Methods(http.MethodPost)

	PORT := os.Getenv("PORT")

	if PORT == "" {
		PORT = "8080"
	}

	log.Println("Server is running in port: " + PORT)

	handler := cors.AllowAll().Handler(router)
	log.Fatal(http.ListenAndServe(":"+PORT, handler))
}
