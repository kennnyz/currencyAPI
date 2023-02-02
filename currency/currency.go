package currency

import (
	"github.com/gofiber/fiber/v2"
)

type Currency struct {
	CurrencyFrom string  `json:"currency_from,omitempty"`
	CurrencyTo   string  `json:"currency_to,omitempty"`
	ExchangeRate float64 `json:"exchange_rate,omitempty"`
	UpdatedAt    string  `json:"updated_at,omitempty"`
}

func GetCurrency(c *fiber.Ctx) error {

	return c.SendString("Getting currency ")
}

func PostCurrency(c *fiber.Ctx) error {
	CreateCurrency(c)
	return c.SendString("Posting a new  currency ")

}

func PutCurrency(c *fiber.Ctx) error {
	return c.SendString("Агрегация добавленных валютных пар")
}
