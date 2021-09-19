package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type User struct {
	ID       string `json:"id"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func main() {
	// Init
	r := mux.NewRouter()

	// Route Handlers
	r.HandleFunc("/api/login", login).Methods("POST")

	log.Fatal(http.ListenAndServe(":8000", r))
}

func login(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var user User
	_ = json.NewDecoder(r.Body).Decode(&user)

	json.NewEncoder(w).Encode(user)
}
