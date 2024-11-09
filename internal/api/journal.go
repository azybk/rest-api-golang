package api

import (
	"context"
	"net/http"
	"rest-api-golang/domain"
	"rest-api-golang/dto"
	"rest-api-golang/internal/util"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

type journalApi struct {
	journalService domain.JournalService
}

func NewJournal(app *fiber.App, jornalService domain.JournalService, authMid fiber.Handler) {
	ja := journalApi{
		journalService: jornalService,
	}

	journals := app.Group("/journals", authMid)
	journals.Get("", ja.Index)
	journals.Post("", ja.Create)
	journals.Put(":id", ja.Update)
}

func (ja journalApi) Index(ctx *fiber.Ctx) error {
	c, cancel := context.WithTimeout(ctx.Context(), 10*time.Second)
	defer cancel()

	customerId := ctx.Query("customer_id")
	status := ctx.Query("status")

	res, err := ja.journalService.Index(c, domain.JournalSearch{
		CustomerId: customerId,
		Status:     status,
	})
	if err != nil {
		return ctx.Status(http.StatusInternalServerError).JSON(dto.CreateResponseError(err.Error()))
	}

	return ctx.Status(http.StatusOK).JSON(dto.CreateResponseSuccess(res))
}

func (ja journalApi) Create(ctx *fiber.Ctx) error {
	c, cancel := context.WithTimeout(ctx.Context(), 10*time.Second)
	defer cancel()

	var req dto.CreateJournalRequest
	if err := ctx.BodyParser(&req); err != nil {
		return ctx.SendStatus(http.StatusUnprocessableEntity)
	}

	fails := util.Validate(req)
	if len(fails) > 0 {
		return ctx.Status(http.StatusBadRequest).JSON(dto.CreateResponseErrorData("validasi gagal", fails))
	}

	err := ja.journalService.Create(c, req)
	if err != nil {
		return ctx.Status(http.StatusInternalServerError).JSON(dto.CreateResponseError(err.Error()))
	}

	return ctx.Status(http.StatusCreated).JSON(dto.CreateResponseSuccess(""))
}

func (ja journalApi) Update(ctx *fiber.Ctx) error {
	c, cancel := context.WithTimeout(ctx.Context(), 10*time.Second)
	defer cancel()

	id := ctx.Params("id")
	claim := ctx.Locals("user").(*jwt.Token).Claims.(jwt.MapClaims)

	err := ja.journalService.Return(c, dto.ReturnedJournalRequest{
		JournalId: id,
		UserId:    claim["id"].(string),
	})

	if err != nil {
		return ctx.Status(http.StatusInternalServerError).JSON(dto.CreateResponseError(err.Error()))
	}

	return ctx.Status(http.StatusCreated).JSON(dto.CreateResponseSuccess(""))
}
