package models

// Email is a struct that holds the email data
type Email struct {
	Name    string `json:"name"`
	Number  string `json:"number"`
	To      string `json:"to"`
	Subject string `json:"subject"`
	Body    string `json:"body"`
}
