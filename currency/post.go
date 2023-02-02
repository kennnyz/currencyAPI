package currency

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
)

func CreateCurrency(c *fiber.Ctx) {
	var in Currency
	if err := c.BodyParser(&in); err != nil {
		c.Status(400).JSON(fiber.Map{
			"error": err.Error(),
		})
		return
	}
	fmt.Println(in)
}
