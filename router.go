package main

import (
	"net/http"
)

type Router struct {
	rules map[string]http.HandlerFunc
}

func NewRouter() *Router {
	return &Router{
		rules: make(map[string]http.HandlerFunc),
	}
}

func (r *Router) FindPath(path string) bool {
	_, exist := r.rules[path]
	return exist
}

func (r *Router) FindHandler(path string) (http.HandlerFunc, bool) {
	_, exist := r.rules[path]
	handler := r.rules[path]
	return handler, exist
}

func (r *Router) ServeHTTP(w http.ResponseWriter, request *http.Request) {
	handler, exist := r.FindHandler(request.URL.Path)
	if !exist {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	handler(w, request)
}