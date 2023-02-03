package main

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/jmoiron/sqlx"
	"github.com/kennnyz/currencyAPI/currency"
	"github.com/kennnyz/currencyAPI/database"
	_ "github.com/lib/pq"
	"log"
)

func main() {
	app := fiber.New()
	initDatabase()

	defer func(DB *sqlx.DB) {
		err := DB.Close()
		if err != nil {
			log.Panicf("Panic to close database: %v", err)
		}
	}(database.DB)

	app.Use(Middleware)
	setupRoutes(app, database.DB)

	err := app.Listen(":3000")
	if err != nil {
		return
	}

}

func Middleware(c *fiber.Ctx) error {
	err := c.Next()
	if err != nil {
		return err
	}
	println("Request:", c.Method(), c.Path())
	return nil
}
func initDatabase() {
	var err error
	database.DB, err = sqlx.Connect("postgres", "")
	if err != nil {
		log.Panicf("Panic to connect to database: %v", err)
		return
	}
	fmt.Println("Connected to database")
}
func setupRoutes(app *fiber.App, db *sqlx.DB) {
	app.Get("/api/currency", currency.GetCurrency(db))
	app.Post("/api/currency", currency.CreateCurrency(db))
	app.Put("/api/currency", currency.PutCurrency(db))
}
