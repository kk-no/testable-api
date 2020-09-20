package router

import "github.com/go-chi/chi"

type Server struct {
	Router *chi.Mux
}

func New() *Server {
	return &Server{Router: chi.NewRouter()}
}
