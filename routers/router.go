package routers

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/yeasin2002/url-shortener/utils"
)

type UserUrl struct {
	URL string `json:"url"`
}

func GetUrl(c *fiber.Ctx) error {
	db := utils.ConnectDB("./database.db")
	defer db.Close()
	body := new(UserUrl)

	if err := c.BodyParser(body); err != nil {
		fmt.Printf("Error : %v", err)
		return c.Status(fiber.StatusBadRequest).SendString("Invalid JSON")
	}

	return c.JSON(map[string]string{
		"URL": body.URL,
	})

}

func RedirectToOriginalUrl(c *fiber.Ctx) error {
	return c.SendString("Hello, World!")
}
