package products

import (
	"errors"
	"github.com/dodysat/goweb/database"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
	"strconv"
)

func Create(c *fiber.Ctx) error {
	db := database.DB
	newProduct := new(Product)
	if err := c.BodyParser(newProduct); err != nil {
		return err
	}

	err := db.Create(&newProduct).Error
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"status":  false,
			"message": nil,
			"data":    nil,
			"error":   err,
		})
	} else {
		return c.Status(201).JSON(fiber.Map{
			"status":  true,
			"message": "Product created successfully!",
			"data":    newProduct,
			"error":   nil,
		})
	}
}

func List(c *fiber.Ctx) error {
	var product []Product
	db := database.DB
	db.Find(&product)

	var dataReturn = map[string]interface{}{
		"status":  true,
		"message": "success!",
		"data":    product,
		"error":   nil,
	}
	return c.Status(200).JSON(dataReturn)
}

func One(c *fiber.Ctx) error {
	var product Product
	db := database.DB

	id := c.Params("id")
	if id != "" {
		id, _ := strconv.ParseUint(id, 10, 64)
		result := db.Where(&Product{ID: uint(id)}).First(&product)

		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return c.Status(404).JSON(fiber.Map{
				"status":  false,
				"message": "record not found",
				"data":    nil,
				"error":   nil,
			})
		} else {
			return c.Status(200).JSON(fiber.Map{
				"status":  true,
				"message": "Data Updated",
				"data":    product,
				"error":   nil,
			})
		}
	} else {
		return c.Status(401).JSON(fiber.Map{
			"status":  false,
			"message": "id of product needed",
			"data":    nil,
			"error":   nil,
		})
	}
}

func Change(c *fiber.Ctx) error {
	var product Product
	db := database.DB

	id := c.Params("id")
	if id != "" {
		id, _ := strconv.ParseUint(id, 10, 64)
		result := db.Where(&Product{ID: uint(id)}).First(&product)

		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return c.Status(404).JSON(fiber.Map{
				"status":  false,
				"message": "record not found",
				"data":    nil,
				"error":   nil,
			})
		} else {
			newProduct := new(Product)
			if err := c.BodyParser(newProduct); err != nil {
				return err
			}
			product.Name = newProduct.Name
			product.Detail = newProduct.Detail
			db.Save(&product)

			return c.Status(200).JSON(fiber.Map{
				"status":  true,
				"message": "Data Updated",
				"data":    product,
				"error":   nil,
			})
		}
	} else {
		return c.Status(401).JSON(fiber.Map{
			"status":  false,
			"message": "id of product needed",
			"data":    nil,
			"error":   nil,
		})
	}
}

func Delete(c *fiber.Ctx) error {
	var product Product
	db := database.DB

	id := c.Params("id")
	if id != "" {
		id, _ := strconv.ParseUint(id, 10, 64)
		result := db.Where(&Product{ID: uint(id)}).First(&product)

		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return c.Status(404).JSON(fiber.Map{
				"status":  false,
				"message": "record not found",
				"data":    nil,
				"error":   nil,
			})
		} else {
			db.Delete(&product)

			return c.Status(200).JSON(fiber.Map{
				"status":  true,
				"message": "Data Deleted",
				"data":    product,
				"error":   nil,
			})
		}
	} else {
		return c.Status(401).JSON(fiber.Map{
			"status":  false,
			"message": "id of product needed",
			"data":    nil,
			"error":   nil,
		})
	}
}
