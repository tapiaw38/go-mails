package utils

import (
	"bytes"
	"crypto/tls"
	"html/template"
	"log"
	"net/mail"
	"net/smtp"

	"github.com/tapiaw38/go-mails/models"
)

// SendEmail sends an email
func SendEmail(name, to, subject, body string, c chan string) {

	fromEmail := mail.Address{
		Name:    "Go-mails",
		Address: "waltertapia153@gmail.com",
	}
	toEmail := mail.Address{
		Name:    "",
		Address: to,
	}
	subjectEmail := subject

	bodyEmail := models.Email{
		Name:    name,
		To:      to,
		Subject: subject,
		Body:    body,
	}

	headers := make(map[string]string)
	headers["From"] = fromEmail.String()
	headers["To"] = toEmail.String()
	headers["Subject"] = subjectEmail
	headers["Content-Type"] = "text/html; charset=UTF-8"

	message := ""

	for k, v := range headers {
		message += k + ": " + v + "\r\n"
	}

	t, err := template.ParseFiles("templates/email.html")

	if err != nil {
		log.Println(err)
		c <- "error"
		return
	}

	buf := new(bytes.Buffer)
	err = t.Execute(buf, bodyEmail)

	if err != nil {
		log.Println(err)
		c <- "error"
		return
	}

	message += "\r\n" + buf.String()

	servername := "smtp.gmail.com:465"
	host := "smtp.gmail.com"

	auth := smtp.PlainAuth("", "waltertapia153@gmail.com", "Walter153294", host)

	tlsConfig := &tls.Config{
		InsecureSkipVerify: true,
		ServerName:         host,
	}

	conn, err := tls.Dial("tcp", servername, tlsConfig)

	if err != nil {
		log.Println(err)
		c <- "error"
		return
	}

	client, err := smtp.NewClient(conn, host)

	if err != nil {
		log.Println(err)
		c <- "error"
		return
	}

	err = client.Auth(auth)

	if err != nil {
		log.Println(err)
		c <- "error"
	}

	err = client.Mail(fromEmail.Address)

	if err != nil {
		log.Println(err)
		c <- "error"
		return
	}

	err = client.Rcpt(toEmail.Address)

	if err != nil {
		log.Println(err)
		c <- "error"
		return
	}

	w, err := client.Data()

	if err != nil {
		log.Println(err)
		c <- "error"
		return
	}

	_, err = w.Write([]byte(message))

	if err != nil {
		log.Println(err)
		c <- "error"
		return
	}

	err = w.Close()

	if err != nil {
		log.Println(err)
		c <- "error"
		return
	}

	err = client.Quit()

	if err != nil {
		log.Println(err)
		c <- "error"
		return
	}

	c <- "ok"
}
