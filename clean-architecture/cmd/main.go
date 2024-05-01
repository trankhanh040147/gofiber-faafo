package main

import (
	"gofiber-faafo/cmd/launcher"
	"log"
)

func main() {
	db, cancel := launcher.NewDBConnection("mongodb://localhost:27017", "books")
	app := launcher.NewApp()
	launcher.SetupRoute(app, db)

	defer cancel()
	log.Fatal(app.Listen(":8080"))
}
