package main

import (
	"fmt"
	"strconv"
	"net/http"
	"github.com/gorilla/mux"
	"github.com/olling/slog"
)

func InitializeApiHandlers(router *mux.Router) {
	router.HandleFunc("/api", handler)
	router.HandleFunc("/api/recipes", handlerRecipes).Methods("GET", "POST")
	router.HandleFunc("/api/recipes/{id}", handlerRecipesID).Methods("GET", "DELETE")
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
		handleGetRecipe(w, r, id)

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
			handleGetRecipes(w, r)
		} else {
			handleGetRecipe(w, r, irecipe.ID)
		}

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

func handleGetRecipes(w http.ResponseWriter, r *http.Request) {
	recipes, err := getRecipeList()
	if err != nil {
		http.Error(w, "Something went wrong while getting recipes", 500)
		slog.PrintError("Error getting recipe list", err)
	}
	output, err := ToJson(recipes)

	if err != nil {
		http.Error(w, "Something went wrong while getting recipes", 500)
		slog.PrintError("Error converting recipes to json", err)
		return
	}

	fmt.Fprint(w, output)
}

func handleGetRecipe(w http.ResponseWriter, r *http.Request, ID int) {
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
