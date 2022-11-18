package main

import (
	"log"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"

	"example.com/crud-user/database"
	"example.com/crud-user/router"
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
