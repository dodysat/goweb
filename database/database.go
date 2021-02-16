package database

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func OpenConnection() {
	var err error
	//dbHost := os.Getenv("DB_HOST")
	//dbPort := os.Getenv("DB_PORT")
	//dbDatabase := os.Getenv("DB_DATABASE")
	//dbUsername := os.Getenv("DB_USERNAME")
	//dbPassword := os.Getenv("DB_PASSWORD")

	dbHost := "101.50.1.75"
	dbPort := "3306"
	dbDatabase := "edusaid1_gogogo"
	dbUsername := "edusaid1_gogogo"
	dbPassword := "DY3542v87435adfgQY3478u34U5bm345z"

	dsn := dbUsername + ":" + dbPassword + "@tcp(" + dbHost + ":" + dbPort + ")/" + dbDatabase + "?charset=utf8mb4&parseTime=True&loc=Local"
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("failed to connect database")
	}
	fmt.Println("Connection Opened to Database")
}

//PORT=8080
//
//DB_HOST=101.50.1.75
//DB_PORT=3306
//DB_DATABASE=edusaid1_gogogo
//DB_USERNAME=edusaid1_gogogo
//DB_PASSWORD=DY3542v87435adfgQY3478u34U5bm345z