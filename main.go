package main

/*
 	estef.link
    Una app web construida con Go + Astro que permite al usuario acortar URLs mediante una autenticacion con Google.
*/

import (
	"context"
	"log"
	"os"

	"estef.link/db"
	"estef.link/routes"
	"github.com/gofiber/fiber/v3"
	"github.com/subosito/gotenv"
)

func main() {
	if err := gotenv.Load(); err != nil {
		panic(err)
	}
	port := os.Getenv("PORT")
	uri := os.Getenv("MONGOURI")
	if err := db.ConnectDB(uri); err != nil {
		panic(err)
	}

	defer func() {
		if err := db.Client.Disconnect(context.TODO()); err != nil {
			log.Printf("error disconnecting from MongoDB %v", err)
		}
	}()
	app := fiber.New()

	// iniciando rutas
	routes.InitRoutes(app)

	app.Listen(port)
}
