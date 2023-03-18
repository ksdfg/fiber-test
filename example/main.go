package main

import (
	"log"

	"github.com/ksdfg/fiber-test/example/api"
)

func main() {
	controller := api.Controller{}

	controller.Data = make(map[string]string)
	controller.Data["cat"] = "meow"
	controller.Data["dog"] = "woof"
	controller.Data["cow"] = "moo"

	app := api.GenApp(controller)
	log.Fatalln(app.Listen(":8000"))
}
