package main

import (
	"encoding/json"
	"fmt"
	"github.com:vikramjeetsingh41/go-mysql/go/dao"
	"github.com:vikramjeetsingh41/go-mysql/go/service"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// Users struct
type Users []User

// User struct
type User struct {
	Name     string `json:"name"`
	LastName string `json:"lastName"`
	Genius   bool   `json:"genius"`
}

// Characters are
type Characters []User

// HomeHandler method
func HomeHandler(w http.ResponseWriter, r *http.Request) {
	// w.Header().Set("Content-Type", "application/json")
	// w.WriteHeader(http.StatusCreated)
	// json.NewEncoder(w).Encode(m)

	// users := Users{
	// 	User{firstName: "vikram", lastName: "singh"},
	// 	User{firstName: "rocky", lastName: "jeet"},
	// }

	//openDbConnection()

	student, err := service.AddUser("vikram", "test@test.ocm", 35)
	if err != nil {
		fmt.Println("Adding Student Failed With Error : ", err.Error())
	} else {
		fmt.Println("Added Student Successfully : ", student)
	}

	characters := Users{
		User{Name: "Jimmy Neutron", Genius: true, LastName: "singh"},
	}

	//url := "https://jsonplaceholder.typicode.com/todos/1"

	json.NewEncoder(w).Encode(characters)
}

func articlesHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome home articles page!")
}

// ArticlesCategoryHandler method
func ArticlesCategoryHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Category: %v%v\n", vars["category"], vars["id"])
}

func main() {

	dao.InitializeMySQL()

	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", HomeHandler)
	router.HandleFunc("/articles", articlesHandler).Methods("GET")
	router.HandleFunc("/articles/{category}/{id:[0-9]+}", ArticlesCategoryHandler).Methods("GET")
	http.Handle("/", router)
	log.Fatal(http.ListenAndServe(":8080", router))
}

// ExportedFuncSum is a function you can use from a different file
func ExportedFuncSum(x int, y int) int {
	return sum(x, y)
}

// while this one stays within the context of this file
func sum(x int, y int) int {
	return x + y
}
