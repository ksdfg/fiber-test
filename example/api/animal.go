package api

import (
	"log"
	"net/http"

	"github.com/gofiber/fiber/v2"

	"github.com/ksdfg/fiber-test/example/schemas"
)

func (c Controller) GetAnimal(ctx *fiber.Ctx) error {
	var query schemas.AnimalSoundQuery
	err := ctx.QueryParser(&query)
	if err != nil {
		log.Println(err)
		return ctx.SendStatus(http.StatusInternalServerError)
	}

	noise, ok := c.Data[query.Name]
	if !ok {
		return ctx.Status(http.StatusNotFound).SendString("I don't know what sound this animal makes ;-;")
	}

	return ctx.Status(http.StatusOK).JSON(schemas.Animal{Name: query.Name, Sound: noise})
}

func (c Controller) AddAnimal(ctx *fiber.Ctx) error {
	var animal schemas.Animal
	err := ctx.BodyParser(&animal)
	if err != nil {
		log.Println(err)
		return ctx.SendStatus(http.StatusInternalServerError)
	}

	c.Data[animal.Name] = animal.Sound

	return ctx.SendStatus(http.StatusCreated)
}
