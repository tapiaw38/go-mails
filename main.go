package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"

	"github.com/tapiaw38/go-mails/routers"
)

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/email", routers.HandleEmail).Methods(http.MethodPost)

	PORT := ":8080"

	log.Fatal(http.ListenAndServe(PORT, router))

}
