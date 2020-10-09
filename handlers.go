package main

import (
	"fmt"
	"net/http"
)

func HandleRoot(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello world from root path!!")
}

func HandleApi(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "We are in the API")
}