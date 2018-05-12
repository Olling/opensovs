package main

import (
	"fmt"
)

func main() {
	InitializeConfiguration()
	InitializeDBMigration()

	recipes := []Recipe{}
	recipes = append(recipes, Recipe{1, "Test01", "20180512", "1", "1"})

	fmt.Println("START BULK")

	err := bulkInsertRecipes(recipes)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("END BULK")

	InitializeApi()
}
