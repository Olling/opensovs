package main

import (
	"fmt"
	"net/http"
	"github.com/gorilla/mux"
)

func InitializeWeb (router *mux.Router) {
	r.PathPrefix("/").Handler(http.StripPrefix("/", http.FileServer(http.Dir("./public/"))))
}

func handlerWebRoot(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello world")
}
