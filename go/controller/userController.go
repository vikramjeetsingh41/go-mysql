package controller

import (
	"encoding/json"
	"fmt"
	"go-mysql/go/model"
	"go-mysql/go/service"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type Message struct {
	Status string `json:status`
}

func GetUsers(w http.ResponseWriter, r *http.Request) {
	users, err := service.GetUsers()
	if err != nil {
		fmt.Fprintf(w, "ERROR")
		return
	}

	jData, err := json.Marshal(users)
	if err != nil {
		// handle error
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(jData)
}

func GetUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	userId, err := strconv.ParseInt(vars["id"], 10, 64)
	if err != nil {
		panic(err)
	}

	user, err := service.GetUser(userId)
	if err != nil {
		fmt.Fprintf(w, "ERROR")
		return
	}

	jData, err := json.Marshal(user)
	if err != nil {
		// handle error
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(jData)
}

func AddUser(w http.ResponseWriter, r *http.Request) {

	var user model.User
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&user); err != nil {
		//respondWithError(w, http.StatusBadRequest, "Invalid payload")
	}
	defer r.Body.Close()

	user, err := service.AddUser(user.Name, user.Age)
	if err != nil {
		// error
	}

	jData, err := json.Marshal(user)
	if err != nil {
		// handle error
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(jData)
}

func UpdateUser(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	userId, _ := strconv.ParseInt(vars["id"], 10, 64)

	var user model.User
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&user); err != nil {
		//respondWithError(w, http.StatusBadRequest, "Invalid payload")
	}
	defer r.Body.Close()

	user, err := service.UpdateUser(userId, user.Name, user.Age, user.Status)
	if err != nil {
		// Error
	}
	jData, err := json.Marshal(user)
	if err != nil {
		// handle error
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(jData)
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userId, _ := strconv.ParseInt(vars["id"], 10, 64)

	isDeleted, err := service.DeleteUser(userId)
	if err != nil {
		// Handle error
	}
	if !isDeleted {
		// User is not deleted
	}

	success := Message{Status: "Deleted"}
	response, _ := json.Marshal(success)

	w.Header().Set("Content-Type", "application/json")
	w.Write(response)
}
