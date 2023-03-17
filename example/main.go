package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"

	"github.com/ksdfg/fiber-test/example/controller"
)

func main() {
	c := controller.Controller{}

	c.Data = make(map[string]string)
	c.Data["cat"] = "meow"
	c.Data["dog"] = "woof"
	c.Data["cow"] = "moo"

	app := fiber.New()
	app.Use(logger.New())

	controller.RegisterRoutes(app.Group("/api/v1"), c)

	log.Fatalln(app.Listen(":8000"))
}
