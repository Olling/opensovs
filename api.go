package main

import (
	"log"
	"strconv"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

func InitializeApi() {
	r := mux.NewRouter()

	r.HandleFunc("/api",handler)
	r.HandleFunc("/api/recipes",handlerRecipes).Methods("GET","POST")

	// Bind to a port and pass our router in
	log.Fatal(http.ListenAndServe(":" + strconv.Itoa(Conf.ApiPort), r))
}


func handlerRecipes(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		fmt.Fprint(w,"GET")
	case "POST":
		fmt.Fprint(w,"POST")
	}
}


func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w,"Running")
}

