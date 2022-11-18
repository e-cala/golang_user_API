package router

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"

	"example.com/crud-user/services"
)

func SetupRouter(router *mux.Router) {
	router.HandleFunc("/api/v1/create", services.CreateUser).Methods("POST")
	router.HandleFunc("/api/v1/update/{email}", services.UpdateUser).Methods("PUT")
	router.HandleFunc("/api/v1/read", services.ReadUser).Methods("GET")
	router.HandleFunc("/api/v1/delete/{email}", services.DeleteUser).Methods("DELETE")
}

func SetupServer(router *mux.Router) {
	http.ListenAndServe(":8080", router)
	fmt.Printf("Listening to port 8080")
}
