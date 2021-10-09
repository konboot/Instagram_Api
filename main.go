package main

import (
	"fmt"
	route "instagram_api/handler"
	controller "instagram_api/key"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

var postCollection, userCollection = controller.ConnectDB()

func index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1>Instagram</h1>")
}
func main() {

	mux := mux.NewRouter()

	mux.HandleFunc("/", index)
	mux.HandleFunc("/posts", route.InsertPost).Methods("GET")
	mux.HandleFunc("/users", route.InsertUser).Methods("GET")
	mux.HandleFunc("/users/{id}", route.GetUser).Methods("GET")
	mux.HandleFunc("/posts/{id}", route.GetPost).Methods("GET")
	mux.HandleFunc("/posts/users/{id}", route.GetUserPost).Methods("GET")
	mux.HandleFunc("/posts", route.CreatePost).Methods("POST")
	mux.HandleFunc("/users", route.CreateUser).Methods("POST")

	fmt.Println("Server Started at port 9000")
	log.Fatal(http.ListenAndServe(":9000", mux))

}
