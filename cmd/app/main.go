package main

import (
	"os"

	"github.com/johanagus/simple-pos/database"
	"github.com/johanagus/simple-pos/models"
	"github.com/johanagus/simple-pos/router"
	"github.com/sirupsen/logrus"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()
	logger := logrus.New()

	file, err := os.OpenFile("./log/app.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		panic(err)
	}

	router.New(app)

	db := database.Init()
	db.AutoMigrate(
		&models.User{},
		&models.Product{},
		&models.Customer{},
		&models.Supplier{},
		&models.Inventory{},
		&models.SalesPayment{},
		&models.Sales{},
		&models.SalesItem{},
		&models.SalesOrder{},
		&models.SalesOrderItem{},
		&models.Warehouse{})

	logger.SetOutput(file)
	logger.Fatal(app.Listen(":8000"))
}
