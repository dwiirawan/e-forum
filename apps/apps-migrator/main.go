package main

import (
	"libs/api-core/database"
	"libs/api-core/utils"
	"os"
)

func main() {

	utils.LoadEnv(1)
	DB := database.MigratorNew()

	argsWithoutProg := os.Args[1:]

	if len(argsWithoutProg) == 0 {
		panic("Please provide a valid argument")
	}

	if argsWithoutProg[0] == "up" {
		DB.DBUp()
	} else if argsWithoutProg[0] == "down" {
		DB.DBDown()
	} else if argsWithoutProg[0] == "status" {
		DB.DBStatus()
	} else {
		panic("Please provide a valid argument")
	}

}
