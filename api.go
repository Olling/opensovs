package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/olling/slog"
	"log"
	"net/http"
	"strconv"
)

func InitializeApi() {
	r := mux.NewRouter()

	r.HandleFunc("/api", handler)
	r.HandleFunc("/api/recipes", handlerRecipes).Methods("GET", "POST")
	r.HandleFunc("/api/recipes/{id}", handlerRecipesID).Methods("GET", "DELETE")

	// Bind to a port and pass our router in
	log.Fatal(http.ListenAndServe(":"+strconv.Itoa(Conf.ApiPort), r))
}

func handlerRecipesID(w http.ResponseWriter, r *http.Request) {
	sid := mux.Vars(r)["id"]
	id, err := strconv.Atoi(sid)

	if err != nil {
		http.Error(w, "Could not read API input (needed ID)", 400)
		return
	}

	switch r.Method {
	case "GET":
		handleGetRecipes(w, r, id)

	case "DELETE":
		err = deleteRecipeById(id)
		if err != nil {
			http.Error(w, "Error occured while deleting recipe", 500)
			slog.PrintError("Could not delete recipe with id:", id, ":", err)
		}
	}

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

		handleGetRecipes(w, r, irecipe.ID)

	case "POST":
		//fmt.Fprint(w, "POST")
		var recipe Recipe
		err := FromJsonReader(r.Body, &recipe)
		defer r.Body.Close()
		if err != nil {
			slog.PrintError("Error decoding API input", err)
		}
		err = bulkInsertRecipes([]Recipe{recipe})
		if err != nil {
			http.Error(w, "Error occured while adding recipe", 500)
			slog.PrintError("Error adding recipe to database:", err)
		}
	}
}

func handleGetRecipes(w http.ResponseWriter, r *http.Request, ID int) {
	recipe, err := getRecipeByID(ID)
	if err != nil {
		http.Error(w, "Recipe not found", 404)
		slog.PrintError("Error retrieving recipe with id="+strconv.Itoa(ID)+":", err)
	}
	output, err := ToJson(recipe)

	if err != nil {
		slog.PrintError("Error converting recipe to json", err)
		return
	}

	fmt.Fprint(w, output)
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Running")
}
