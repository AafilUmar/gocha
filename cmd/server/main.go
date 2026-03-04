package main

import (
	"log"

	"github.com/AafilUmar/gocha/internal/server"
	"github.com/AafilUmar/gocha/internal/store"
)

func main(){

gocha := store.Gocha()

serv := server.New(":8080",gocha)
log.Print("Starting server at 8080")

if err := serv.Start();err != nil {
log.Fatal(err)
}

}
