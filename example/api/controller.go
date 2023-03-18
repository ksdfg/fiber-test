package api

import (
	"github.com/gofiber/fiber/v2"
)

type Controller struct {
	Data map[string]string
}

func GenApp(controller Controller) *fiber.App {
	api := fiber.New()

	router := api.Group("/api/v1")
	router.Get("/animal", controller.GetAnimal)
	router.Post("/animal", controller.AddAnimal)

	return api
}
