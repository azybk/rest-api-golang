package main

import (
	"net/http"
	"rest-api-golang/dto"
	"rest-api-golang/internal/api"
	"rest-api-golang/internal/config"
	"rest-api-golang/internal/connection"
	"rest-api-golang/internal/repository"
	"rest-api-golang/internal/service"

	jwtMid "github.com/gofiber/contrib/jwt"

	"github.com/gofiber/fiber/v2"
)

func main() {
	cnf := config.Get()
	dbConnection := connection.GetDatabase(cnf.Database)

	app := fiber.New()

	jwtMidd := jwtMid.New(jwtMid.Config{
		SigningKey: jwtMid.SigningKey{Key: []byte(cnf.Jwt.Key)},
		ErrorHandler: func(ctx *fiber.Ctx, err error) error {
			return ctx.Status(http.StatusUnauthorized).JSON(dto.CreateResponseError("authentication gagal, silahkan login dulu"))
		},
	})

	customerRepository := repository.NewCustomer(dbConnection)
	userRepository := repository.NewUser(dbConnection)

	customerService := service.NewCustomer(customerRepository)
	authService := service.NewAuth(cnf, userRepository)

	api.NewCustomer(app, customerService, jwtMidd)
	api.NewAuth(app, authService)

	_ = app.Listen(cnf.Server.Host + ":" + cnf.Server.Port)

}
