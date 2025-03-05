package server

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html/v2"
)

func StartServer() {
	// Initialize Fiber app with HTML template engine
	engine := html.New("./template", ".html")
	app := fiber.New(fiber.Config{
		Views: engine,
	})

	// Setup Routes
	Routes(app)

	// Serve Static Files
	app.Static("/static", "./static")

	// Start Server on port 8080
	log.Fatal(app.Listen(":8080"))
}