package main

import (
	"encoding/json"
	"go-mysql/go/controller"
	"go-mysql/go/dao"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Message struct {
	Status string `json:status`
}

// HomeHandler method
func HomeHandler(w http.ResponseWriter, r *http.Request) {
	success := Message{Status: "ok"}

	response, err := json.Marshal(success)
	if err != nil {
		// Error
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(response)
}

func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Do stuff here
		log.Println(r.RequestURI)
		// Call the next handler, which can be another middleware in the chain, or the final handler.
		next.ServeHTTP(w, r)
	})
}

func main() {
	dao.InitializeMySQL()
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", HomeHandler)
	router.HandleFunc("/users", controller.GetUsers).Methods("GET")
	router.HandleFunc("/users/{id}", controller.GetUser).Methods("GET")
	router.HandleFunc("/users", controller.AddUser).Methods("POST")
	router.HandleFunc("/users/{id}", controller.UpdateUser).Methods("PUT")
	router.HandleFunc("/users/{id}", controller.DeleteUser).Methods("DELETE")
	http.Handle("/", router)
	router.Use(loggingMiddleware)
	log.Fatal(http.ListenAndServe(":8080", router))
}
