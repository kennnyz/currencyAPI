package currency

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/jmoiron/sqlx"
	"time"
)

// функкия которая будет принимать сообщение от клиента, проверять имеет ли она поля currency_from и currency_to
// если нет, то возвращать invalid request
// если да, то проверять есть ли такая валютная пара в базе данных, если есть, то возвращать что такая валютная пара уже есть
// если нет, то добавлять в базу данных. курс валют брать с openexchangerates.org

func CreateCurrency(db *sqlx.DB) func(c *fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		var in Currency
		body := c.Body()
		if err := json.Unmarshal(body, &in); err != nil {
			return c.Status(400).Send([]byte("Failed to unmarshal request body"))
		}
		// проверка наличия записи в БД
		var existingCurrency Currency
		err := db.QueryRow("SELECT currency_from, currency_to FROM operations WHERE currency_from = $1 AND currency_to = $2", in.CurrencyFrom, in.CurrencyTo).Scan(&existingCurrency.CurrencyFrom, &existingCurrency.CurrencyTo)
		if err != sql.ErrNoRows {
			_ = c.Status(400).SendString("User already exists")
			return err
		}
		in.ExchangeRate = getRate(in)
		in.UpdatedAt = time.Now().Format("2006-01-02 15:04:05")

		fmt.Println(in)

		// добавление записи в БД
		_, err = db.Exec("INSERT INTO operations (currency_from, currency_to, exchange_rate, updated_at) VALUES ($1, $2,$3, $4)", in.CurrencyFrom, in.CurrencyTo, in.ExchangeRate, in.UpdatedAt)
		if err != nil {
			_ = c.Status(500).SendString("Failed to insert currency into the database")
			return err
		}
		fmt.Println("Запрос от пользователя на добавление пар: : ", in)
		return c.SendString("Posting a new  currency ")
	}
}

// Получение всех записей из БД

func GetCurrency(c *fiber.Ctx) error {

	return c.SendString("Getting currency ")
}

//Перевод валюты в другую валюту

func PutCurrency(c *fiber.Ctx) error {
	return c.SendString("Агрегация добавленных валютных пар")
}