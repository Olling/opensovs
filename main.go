package main

import (
	"log"
	"strconv"
	"net/http"
	"github.com/gorilla/mux"
)

var (
	router mux.Router
)


func InitializeHttp () {
	// Bind to a port and pass our router in
	log.Fatal(http.ListenAndServe(":"+strconv.Itoa(Conf.ApiPort), &router))
}



func main() {
	InitializeConfiguration()
	InitializeDBMigration()

	router = *mux.NewRouter()
	InitializeApiHandlers(&router)
	InitializeWeb(&router)
	InitializeHttp()
}
