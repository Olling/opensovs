package main

import (
	"log"
	"fmt"
	"strconv"
	"net/http"
	"github.com/gorilla/mux"
	"github.com/olling/slog"
)

func InitializeApi() {
	r := mux.NewRouter()

	r.HandleFunc("/api",handler)
	r.HandleFunc("/api/recipes",handlerRecipes).Methods("GET","POST")
	r.HandleFunc("/api/recipes/{id}", handlerRecipesID).Methods("GET")


	// Bind to a port and pass our router in
	log.Fatal(http.ListenAndServe(":" + strconv.Itoa(Conf.ApiPort), r))
}



func handlerRecipesID(w http.ResponseWriter, r *http.Request) {
	sid := mux.Vars(r)["id"]
	id,err := strconv.Atoi(sid)

	if err != nil {
		http.Error(w, "Could not read API input (needed ID)", 400)
		return
	}

	handleGetRecipes(w,r,id)
}

func handlerRecipes(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		var irecipe Recipe
		err := FromJsonReader(r.Body, &irecipe)
		defer r.Body.Close()

		if err != nil {
			http.Error(w, "Could not read API input (needed ID)", 400)
			return
		}

		handleGetRecipes(w,r,irecipe.ID)

	case "POST":
		fmt.Fprint(w,"POST")
		var recipe Recipe
		err := FromJsonReader(r.Body, &recipe)
		defer r.Body.Close()
		if err != nil {
			slog.PrintError("Error decoding API input",err)
		}
		slog.PrintDebug(recipe)
	}
}

func handleGetRecipes(w http.ResponseWriter, r *http.Request, ID int) {
	//TODO Get recipe
	var recipe Recipe
	recipe.ID=1338
	recipe.Title="Got me"
	output,err := ToJson(recipe)

	if err != nil {
		slog.PrintError("Error converting recipe to json", err)
		return
	}

	fmt.Fprint(w,output)
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w,"Running")
}

