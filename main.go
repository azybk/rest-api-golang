package main

import (
	"rest-api-golang/internal/config"
	"rest-api-golang/internal/connection"
	"rest-api-golang/internal/repository"

	"github.com/gofiber/fiber/v2"
)

func main() {
	cnf := config.Get()
	dbConnection := connection.GetDatabase(cnf.Database)

	app := fiber.New()

	customerRepository := repository.NewCustomer(dbConnection)

	app.Get("/developers", developer)

	_ = app.Listen(cnf.Server.Host + ":" + cnf.Server.Port)

}

func developer(ctx *fiber.Ctx) error {
	return ctx.Status(200).JSON("data aink")
}
