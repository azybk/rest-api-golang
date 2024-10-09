package main

import (
	"rest-api-golang/internal/api"
	"rest-api-golang/internal/config"
	"rest-api-golang/internal/connection"
	"rest-api-golang/internal/repository"
	"rest-api-golang/internal/service"

	"github.com/gofiber/fiber/v2"
)

func main() {
	cnf := config.Get()
	dbConnection := connection.GetDatabase(cnf.Database)

	app := fiber.New()

	customerRepository := repository.NewCustomer(dbConnection)
	customerService := service.NewCustomer(customerRepository)

	api.NewCustomer(app, customerService)

	_ = app.Listen(cnf.Server.Host + ":" + cnf.Server.Port)

}
