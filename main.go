package main

import (
	"log"

	"example.com/crud-user/database"
	"example.com/crud-user/router"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load("local.env")
	if err != nil {
		log.Fatalf("Some error occured. Er: %s", err)
	}
	r := mux.NewRouter()

	database.InitDatabase()
	router.SetupRouter(r)
	router.SetupServer(r)
}
