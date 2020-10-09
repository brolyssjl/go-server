package main

import (
	"log"
	"net/http"
)

type Server struct {
	port   string
	router *Router
}

func NewServer(port string) *Server {
	return &Server{
		port:   port,
		router: NewRouter(),
	}
}

func (s *Server) Listen() error {
	http.Handle("/", s.router)
	err := http.ListenAndServe(s.port, nil)
	if err != nil {
		log.Fatal("our server can run :(", err)
		return err
	}

	log.Println("running server :)")
	return nil
}

func (s *Server) Handle(path string, handler http.HandlerFunc) {
	if !s.router.FindPath(path) {
		s.router.rules[path] = nil
	}

	s.router.rules[path] = handler
}

func (s *Server) AddMiddleware(f http.HandlerFunc, middlewares ...Middleware) http.HandlerFunc {
	for _, m := range middlewares {
		f = m(f)
	}

	return f
}