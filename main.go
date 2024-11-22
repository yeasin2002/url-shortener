package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/yeasin2002/url-shortener/routers"
)

func main() {

	app := fiber.New()

	app.Post("/", routers.GetUrl)
	app.Post("/redirect", routers.RedirectToOriginalUrl)

	app.Listen(":8080")
}
