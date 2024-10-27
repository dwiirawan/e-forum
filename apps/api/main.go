package main

import (
	"libs/api-core/database"
	"libs/api-core/utils"
)

func main() {

	utils.LoadEnv(1)
	DB := database.MigratorNew()

	DB.DBUp()

}
