package main

import (
	"github.com/gorilla/mux"

	"example.com/crud-user/database"
	"example.com/crud-user/router"
)

func main() {
	r := mux.NewRouter()

	database.InitDatabase()
	router.SetupRouter(r)
	router.SetupServer(r)
}
