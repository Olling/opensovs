package main

import (
	"fmt"
	"net/http"
	"github.com/gorilla/mux"
)

func InitializeWeb (router *mux.Router) {
	router.HandleFunc("/", handlerWebRoot)
}

func handlerWebRoot(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello world")
}
