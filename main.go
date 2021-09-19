package main

import (
	"context"
	"encoding/json"
	"log"
	"net/http"

	firebase "firebase.google.com/go/v4"
	"firebase.google.com/go/v4/auth"

	"github.com/gorilla/mux"
)

type User struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

var app *firebase.App
var err error
var client *auth.Client

func main() {
	// Init
	r := mux.NewRouter()

	app, err = firebase.NewApp(context.Background(), nil)
	if err != nil {
		log.Fatalf("error initializing app: %v\n", err)
	}

	// Get an auth client from the firebase.App
	client, err = app.Auth(context.Background())
	if err != nil {
		log.Fatalf("error getting Auth client: %v\n", err)
	}

	// Route Handlers
	r.HandleFunc("/api/login", login).Methods("POST")

	log.Fatal(http.ListenAndServe(":8000", r))
}

func login(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var user User
	_ = json.NewDecoder(r.Body).Decode(&user)

	params := (&auth.UserToCreate{}).
		Email(user.Email).
		EmailVerified(false).
		Password(user.Password).
		DisplayName(user.Name)

	u, err := client.CreateUser(r.Context(), params)
	if err != nil {
		log.Fatalf("error creating user: %v\n", err)
	}
	log.Printf("Successfully created user: %v\n", u)

	json.NewEncoder(w).Encode(u)
}
