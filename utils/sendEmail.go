package utils

import (
	"bytes"
	"crypto/tls"
	"html/template"
	"log"
	"net/mail"
	"net/smtp"
	"os"

	"github.com/tapiaw38/go-mails/models"
)

// SendEmail sends an email
func SendEmail(name, to, subject, body string, c chan string) {

	fromEmail := mail.Address{
		Name:    "Automatic Email",
		Address: os.Getenv("EMAIL_HOST_USER"),
	}
	toEmail := mail.Address{
		Name:    "",
		Address: os.Getenv("EMAIL_DESTINATION"),
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

	host := os.Getenv("EMAIL_HOST")
	servername := host + ":" + os.Getenv("EMAIL_PORT")

	auth := smtp.PlainAuth(
		"",
		os.Getenv("EMAIL_HOST_USER"),
		os.Getenv("EMAIL_HOST_PASSWORD"),
		host,
	)

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
