package main

import (
	"github.com/dodysat/goweb/routes"
	"os"

	"flag"
	"github.com/dodysat/goweb/api"
	"github.com/dodysat/goweb/database"
	"github.com/dodysat/goweb/migration"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

var (
	port = flag.String("port", "3000", "Port to listen on")
	prod = flag.Bool("prod", false, "Enable prefork in Production")
)

func main() {
	database.OpenConnection()
	migration.Migrate()

	// Create fiber app
	app := fiber.New(fiber.Config{
		Prefork: *prod, // go run main.go -prod
	})

	// Middleware
	app.Use(recover.New())
	app.Use(logger.New())

	app.Get("/", api.ServerStatus)

	routes.ProductRoutes(app)

	app.Use(api.NotFound)

	serverPort := os.Getenv("PORT")
	runningPort := *port
	if serverPort != "" {
		runningPort = serverPort
	}
	log.Fatal(app.Listen(":" + runningPort))
}
