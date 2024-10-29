package main

import (
	"nodes_app/bootstrap"
	"nodes_app/database"
	"nodes_app/database/models"
	"nodes_app/router"
)

func main() {

	// Setup database
	DB := database.New()
	DB.AutoMigrate(
		&models.Node{},
	)

	// Setup application
	app := bootstrap.New()

	app.SetupRouter(router.New())

	app.LoadHTMLGlob("views/**")

	app.Run()
}
