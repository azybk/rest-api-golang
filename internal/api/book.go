package api

import (
	"context"
	"net/http"
	"rest-api-golang/domain"
	"rest-api-golang/dto"
	"rest-api-golang/internal/util"
	"time"

	"github.com/gofiber/fiber/v2"
)

type bookApi struct {
	bookService domain.BookService
}

func NewBook(app *fiber.App, bookService domain.BookService, autMid fiber.Handler) {
	b := bookApi{
		bookService: bookService,
	}

	books := app.Group("/books", autMid)
	books.Get("", b.Index)
	books.Post("", b.Create)
	books.Get(":id", b.Show)
	books.Put(":id", b.Update)
	books.Delete(":id", b.Delete)
}

func (b bookApi) Index(ctx *fiber.Ctx) error {
	c, cancel := context.WithTimeout(ctx.Context(), 10*time.Second)
	defer cancel()

	res, err := b.bookService.Index(c)
	if err != nil {
		return ctx.Status(http.StatusInternalServerError).JSON(err.Error())
	}

	return ctx.Status(http.StatusOK).JSON(dto.CreateResponseSuccess(res))
}

func (b bookApi) Create(ctx *fiber.Ctx) error {
	c, cancel := context.WithTimeout(ctx.Context(), 10*time.Second)
	defer cancel()

	var req dto.CreateBookRequest
	if err := ctx.BodyParser(&req); err != nil {
		ctx.SendStatus(http.StatusUnprocessableEntity)
	}

	fails := util.Validate(req)
	if len(fails) > 0 {
		return ctx.Status(http.StatusBadRequest).JSON(dto.CreateResponseErrorData("validasi gagal", fails))
	}

	err := b.bookService.Create(c, req)
	if err != nil {
		return ctx.Status(http.StatusInternalServerError).JSON(dto.CreateResponseError(err.Error()))
	}

	return ctx.Status(http.StatusCreated).JSON(dto.CreateResponseSuccess(""))
}

func (b bookApi) Show(ctx *fiber.Ctx) error {
	c, cancel := context.WithTimeout(ctx.Context(), 10*time.Second)
	defer cancel()

	id := ctx.Params("id")
	res, err := b.bookService.Show(c, id)
	if err != nil {
		return ctx.Status(http.StatusInternalServerError).JSON(dto.CreateResponseError(err.Error()))
	}

	return ctx.Status(http.StatusOK).JSON(dto.CreateResponseSuccess(res))
}

func (b bookApi) Update(ctx *fiber.Ctx) error {
	c, cancel := context.WithTimeout(ctx.Context(), 10*time.Second)
	defer cancel()

	var req dto.UpdateBookRequest
	if err := ctx.BodyParser(&req); err != nil {
		ctx.SendStatus(http.StatusUnprocessableEntity)
	}

	id := ctx.Params("id")
	req.Id = id

	err := b.bookService.Update(c, req)
	if err != nil {
		return ctx.Status(http.StatusInternalServerError).JSON(dto.CreateResponseError(err.Error()))
	}

	return ctx.Status(http.StatusOK).JSON(dto.CreateResponseSuccess(""))
}

func (b bookApi) Delete(ctx *fiber.Ctx) error {
	c, cancel := context.WithTimeout(ctx.Context(), 10*time.Second)
	defer cancel()

	id := ctx.Params("id")

	err := b.bookService.Delete(c, id)
	if err != nil {
		return ctx.Status(http.StatusInternalServerError).JSON(dto.CreateResponseError(err.Error()))
	}

	return ctx.Status(http.StatusOK).JSON(dto.CreateResponseSuccess(""))
}
