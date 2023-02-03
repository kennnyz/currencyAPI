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

func GetCurrency(db *sqlx.DB) func(c *fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		rows, err := db.Query("SELECT * FROM operations")
		if err != nil {
			fmt.Println("Error querying the database:", err)
			return err
		}
		defer rows.Close()

		var currencies []Currency
		for rows.Next() {
			var currency Currency
			err := rows.Scan(&currency.CurrencyFrom, &currency.CurrencyTo, &currency.ExchangeRate, &currency.UpdatedAt)
			if err != nil {
				fmt.Println("Error scanning the rows:", err)
				c.SendStatus(500)
				return err
			}
			currencies = append(currencies, currency)
		}
		currenciesJSON, err := json.Marshal(currencies)
		if err != nil {
			fmt.Println("Error marshalling the currencies:", err)
			c.SendStatus(500)
			return err
		}
		c.Send(currenciesJSON)
		return nil
	}

}

//Обновление записи в БД

func PutCurrency(db *sqlx.DB) func(c *fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		var updateData Currency
		err := c.BodyParser(&updateData)
		if err != nil {
			fmt.Println("Error parsing the request body:", err)
			c.SendStatus(400)
			return err
		}

		updateData.ExchangeRate = getRate(updateData)
		updateData.UpdatedAt = time.Now().Format("2006-01-02 15:04:05")
		query := "UPDATE operations SET exchange_rate = $1 WHERE currency_from = $2 AND currency_to = $3"
		_, err = db.Exec(query, updateData.ExchangeRate, updateData.CurrencyFrom, updateData.CurrencyTo)
		if err != nil {
			fmt.Println("Error executing the update query:", err)
			c.SendStatus(500)
			return err
		}

		return c.SendString("Exchange rate updated successfully")

	}
}
