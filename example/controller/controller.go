package controller

import (
	"github.com/gofiber/fiber/v2"
)

type Controller struct {
	Data map[string]string
}

func RegisterRoutes(router fiber.Router, controller Controller) {
	router.Get("/animal", controller.GetAnimal)
	router.Post("/animal", controller.AddAnimal)
}
