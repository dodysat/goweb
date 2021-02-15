package routes

import (
	"github.com/dodysat/goweb/api/v1/products"
	"github.com/gofiber/fiber/v2"
)

func ProductRoutes(app *fiber.App)  {
	v1 := app.Group("/v1")
	v1.Get("/product", products.List)
	v1.Get("/product/:id", products.One)
	v1.Post("/product", products.Create)
	v1.Put("/product/:id", products.Change)
	v1.Delete("/product/:id", products.Delete)
}