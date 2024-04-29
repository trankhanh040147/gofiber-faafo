package main

import (
	"gofiber-faafo/cmd/launcher"
	"log"
)

func main() {
	db, cancel := launcher.NewDBConnection()
	app := launcher.NewApp()
	launcher.SetupRoute(app, db)

	defer cancel()
	log.Fatal(app.Listen(":8080"))
}
