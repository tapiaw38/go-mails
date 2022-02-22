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

	HOST := os.Getenv("HOST")
	PORT := os.Getenv("PORT")

	if PORT == "" {
		PORT = "8080"
	}

	log.Println("Server is running in: ", HOST+":"+PORT)

	handler := cors.AllowAll().Handler(router)
	log.Fatal(http.ListenAndServe(HOST+":"+PORT, handler))
}
