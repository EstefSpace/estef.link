package main

/*
 	estef.link
    Una app web construida con Go + Astro que permite al usuario acortar URLs mediante una autenticacion con Google.
*/

import (
	"os"

	"estef.link/routes"
	"github.com/gofiber/fiber/v3"
	"github.com/subosito/gotenv"
)

func main() {
	app := fiber.New()
	err := gotenv.Load()
	if err != nil {
		panic(err)
	}
	port := os.Getenv("PORT")

	// iniciando rutas
	routes.InitRoutes(app)

	app.Listen(port)
}
