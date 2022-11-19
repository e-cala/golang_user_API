package services

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"

	"example.com/crud-user/database"
	"example.com/crud-user/logs"
)

func CreateUser(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json")
	var user database.User
	err := json.NewDecoder(request.Body).Decode(&user)
	if err != nil {
		logs.Error.Println("There was an error encoding the request body into the struct")
	}
	database.DBConn.Create(&user)
	logs.Info.Println(user)
}

func ReadUser(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json")
	var userList []database.User
	database.DBConn.Find(&userList)
	json.NewEncoder(writer).Encode(userList)
}

func UpdateUser(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "applicaiton/json")
	var user database.User
	params := mux.Vars(request)
	err := json.NewDecoder(request.Body).Decode(&user)
	if err != nil {
		logs.Error.Println("There was an error encoding the request body into the struct")
	}
	database.DBConn.Model(&user).
		Where("email = ?", params["email"]).
		Updates(database.User{
			FirstName: user.FirstName,
			LastName:  user.LastName,
			Email:     user.Email})
}

func DeleteUser(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "appliation/json")
	var user []database.User
	params := mux.Vars(request)["email"]
	database.DBConn.Where("email = ?", params).Delete(&user)
}
