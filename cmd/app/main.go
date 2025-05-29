package main

import (
	"os"

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

	logger.SetOutput(file)
	logger.Fatal(app.Listen(":8000"))
}
