// main.go
package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"log"

	"be-stepup/config"
	"be-stepup/routes"
)

func main() {
	app := fiber.New()

	// Connect ke database
	config.ConnectDB()

	// Setup rute
	routes.Setup(app)

	// Mengaktifkan middleware CORS
	app.Use(cors.New(cors.Config{
		AllowOrigins: "http://127.0.0.1:5500",        // Mengizinkan permintaan dari frontend
		AllowMethods: "GET,POST,PUT,DELETE",          // Metode HTTP yang diizinkan
		AllowHeaders: "Origin, Content-Type, Accept", // Header yang diizinkan
	}))

	app.Get("/products", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{"message": "This is a products endpoint"})
	})

	log.Fatal(app.Listen(":3000"))
}
