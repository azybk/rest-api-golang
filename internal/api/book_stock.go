package api

import (
	"context"
	"net/http"
	"rest-api-golang/domain"
	"rest-api-golang/dto"
	"rest-api-golang/internal/util"
	"strings"
	"time"

	"github.com/gofiber/fiber/v2"
)

type bookStockApi struct {
	bookStockService domain.BookStockService
}

func NewBookStock(app *fiber.App, bookStockService domain.BookStockService, authMid fiber.Handler) {
	bsa := bookStockApi{
		bookStockService: bookStockService,
	}

	bookStock := app.Group("/book-stocks", authMid)
	bookStock.Post("", bsa.Create)
	bookStock.Delete("", bsa.Delete)
}

func (ba bookStockApi) Create(ctx *fiber.Ctx) error {
	c, cancel := context.WithTimeout(ctx.Context(), 10*time.Second)
	defer cancel()

	var req dto.CreateBookStockRequest
	if err := ctx.BodyParser(&req); err != nil {
		return ctx.SendStatus(http.StatusUnprocessableEntity)
	}

	fails := util.Validate(req)
	if len(fails) > 0 {
		return ctx.Status(http.StatusBadRequest).JSON(dto.CreateResponseErrorData("validasi gagal", fails))
	}

	err := ba.bookStockService.Create(c, req)
	if err != nil {
		return ctx.Status(http.StatusInternalServerError).JSON(dto.CreateResponseError(err.Error()))
	}

	return ctx.Status(http.StatusCreated).JSON(dto.CreateResponseSuccess(""))
}

func (ba bookStockApi) Delete(ctx *fiber.Ctx) error {
	c, cancel := context.WithTimeout(ctx.Context(), 10*time.Second)
	defer cancel()

	// ?code=
	codeStr := ctx.Query("code")
	if codeStr == "" {
		return ctx.Status(http.StatusBadRequest).JSON(dto.CreateResponseError("masukan code buku"))
	}

	codes := strings.Split(codeStr, ";")

	err := ba.bookStockService.Delete(c, dto.DeleteBookStockRequest{Codes: codes})
	if err != nil {
		return ctx.Status(http.StatusInternalServerError).JSON(dto.CreateResponseError(err.Error()))
	}

	return ctx.SendStatus(http.StatusNoContent)
}
