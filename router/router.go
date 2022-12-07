package router

import (
	"net/http"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"

	"example.com/crud-user/logs"
	"example.com/crud-user/services"
)

func SetupRouter(router *mux.Router) {
	router.HandleFunc("/api/v1/create", services.CreateUser).Methods("POST")
	router.HandleFunc("/api/v1/update/{id}", services.UpdateUser).Methods("PUT")
	router.HandleFunc("/api/v1/read", services.ReadUser).Methods("GET")
	router.HandleFunc("/api/v1/delete/{id}", services.DeleteUser).Methods("DELETE")
}

func SetupServer(router *mux.Router) {
	credentials := handlers.AllowCredentials()
	methods := handlers.AllowedMethods([]string{"GET"})
	ttl := handlers.MaxAge(3600)
	origins := handlers.AllowedOrigins([]string{"http://localhost:8080"})
	logs.Error.Println(http.ListenAndServe(":9080", handlers.CORS(credentials, methods, origins, ttl)(router)))
}
