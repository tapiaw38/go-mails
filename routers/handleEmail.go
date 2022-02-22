package routers

import (
	"encoding/json"
	"net/http"

	"github.com/tapiaw38/go-mails/models"
	"github.com/tapiaw38/go-mails/utils"
)

// HandleEmail handles the email request
func HandleEmail(w http.ResponseWriter, r *http.Request) {

	var email models.Email

	err := json.NewDecoder(r.Body).Decode(&email)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if email.To == "" || email.Subject == "" || email.Body == "" {
		http.Error(w, "Missing required fields", http.StatusBadRequest)
		return
	}

	response := models.SendEmailResponse{
		Status:  "ok",
		Message: "the email was sent successfully",
	}

	c := make(chan string)

	go utils.SendEmail(email.Name, email.To, email.Subject, email.Body, c)

	resp := <-c

	if resp != "ok" {
		response.Status = "error"
		response.Message = "the email was not sent"
		w.WriteHeader(http.StatusInternalServerError)
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(response)
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
