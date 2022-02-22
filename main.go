package main

import (
	"github.com/tapiaw38/go-mails/db"
	"github.com/tapiaw38/go-mails/handlers"
)

func main() {
	db.LoadEnv()
	handlers.HandlerServer()
}
