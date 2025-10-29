// main.go
package main

import (
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func main() {
	r := chi.NewRouter()
	r.Use(middleware.Logger)

	// Define route constants
	const userByID = "/users/{id}"

	// Setup routes
	r.Get("/users", getAllUsersHandler)
	r.Post("/users", createUserHandler)
	r.Get(userByID, getUserHandler)
	r.Put(userByID, updateUserHandler)
	r.Delete(userByID, deleteUserHandler)

	log.Println("Server starting on :3000")
	if err := http.ListenAndServe(":3000", r); err != nil {
		log.Fatalf("Could not start server: %s\n", err)
	}
}