package api

import (
	"context"
	"net/http"
	"path/filepath"
	"rest-api-golang/domain"
	"rest-api-golang/dto"
	"rest-api-golang/internal/config"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

type mediaApi struct {
	conf         *config.Config
	mediaService domain.MediaService
}

func NewMedia(app *fiber.App, conf *config.Config, mediaService domain.MediaService, authMid fiber.Handler) {
	ma := mediaApi{
		conf:         conf,
		mediaService: mediaService,
	}

	media := app.Group("/media", authMid)
	media.Post("", ma.Create)
	app.Static("/media", conf.Storage.BasePath)
}

func (ma mediaApi) Create(ctx *fiber.Ctx) error {
	c, cancel := context.WithTimeout(ctx.Context(), 10*time.Second)
	defer cancel()

	file, err := ctx.FormFile("media")
	if err != nil {
		return ctx.SendStatus(http.StatusBadRequest)
	}

	filename := uuid.NewString() + file.Filename
	path := filepath.Join(ma.conf.Storage.BasePath, filename)

	err = ctx.SaveFile(file, path)
	if err != nil {
		return ctx.Status(http.StatusInternalServerError).JSON(dto.CreateResponseError(err.Error()))
	}

	res, err := ma.mediaService.Create(c, dto.CreateMediaRequest{
		Path: filename,
	})
	if err != nil {
		return ctx.Status(http.StatusInternalServerError).JSON(dto.CreateResponseError(err.Error()))
	}

	return ctx.Status(http.StatusCreated).JSON(dto.CreateResponseSuccess(res))
}
