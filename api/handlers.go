package api

import (
	"github.com/gofiber/fiber/v2"
	"time"
)

func ServerStatus(c *fiber.Ctx) error {
	var serverData = map[string]interface{}{
		"version": "0.0.1",
		"time":    time.Now(),
	}
	var dataReturn = map[string]interface{}{
		"status":  true,
		"message": "Server up and running",
		"data":    serverData,
		"error":   nil,
	}
	return c.Status(200).JSON(dataReturn)
}

func NotFound(c *fiber.Ctx) error {
	return c.Status(404).JSON(fiber.Map{
		"status":  false,
		"message": "Endpoint Not Found!",
		"data":    nil,
		"error":   nil,
	})
}
