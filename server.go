package main

import (
	"bones/config"
	"bones/repositories"
	"bones/web/handlers"
	"github.com/gorilla/pat"
	"log"
	"net/http"
)

func main() {
	repositories.Connect(config.Database())
	defer repositories.Cleanup()

	repositories.EnableSessions()

	r := setupRouting()

	http.Handle("/", r)
	http.Handle("/assets/", http.StripPrefix("/assets/", http.FileServer(http.Dir("./assets"))))

	log.Println("Starting server on port 8080")

	log.Fatal(http.ListenAndServe(":8080", nil))
}

func setupRouting() *pat.Router {
	r := pat.New()

	handlers.SetRouter(r)

	r.Get("/users/{id:[0-9]+}/profile", handlers.LoadUserProfilePage).Name("userProfile")

	r.Get("/signup", handlers.LoadSignupPage)
	r.Post("/signup", handlers.CreateNewUser)

	r.Get("/login", handlers.LoadLoginPage)
	r.Post("/login", handlers.CreateNewSession)
	r.Get("/logout", handlers.Logout)

	r.Get("/", handlers.LoadHomePage)

	return r
}
