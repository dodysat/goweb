package main

import (
	"flag"
	"github.com/dodysat/goweb/api"
	"github.com/dodysat/goweb/database"
	"github.com/dodysat/goweb/migration"
	"github.com/dodysat/goweb/routes"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

var (
	port = flag.String("port", "8080", "Port to listen on")
	prod = flag.Bool("prod", false, "Enable prefork in Production")
)

func main() {
	//err := godotenv.Load()
	//if err != nil {
	//	log.Fatal("Error loading .env file")
	//}

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

	//serverPort := os.Getenv("PORT")
	runningPort := *port
	//if serverPort != "" {
	//	runningPort = serverPort
	//}
	log.Fatal(app.Listen(":" + runningPort))
}
