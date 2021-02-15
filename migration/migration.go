package migration

import (
	"github.com/dodysat/goweb/api/v1/products"
	"github.com/dodysat/goweb/database"
)

func Migrate() {
	db := database.DB
	db.AutoMigrate(&products.Product{})
}
