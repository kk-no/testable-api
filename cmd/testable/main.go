package main

import (
	"net/http"
	"os"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	_ "github.com/kk-no/testable-api/database/mysql"
	"github.com/kk-no/testable-api/router"
)

func main() {
	r := router.New()
	r.Router.Use(middleware.Logger)
	r.Router.Route("/api/v1", func(r chi.Router) {
		r.Route("/users", router.Users)
	})

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	http.ListenAndServe(":"+port, r.Router)
}
