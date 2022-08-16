package handlers

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/pr1s10n3r/microurl/pkg/url"
)

type urlHandler struct {
	repository url.Repository
}

func NewURLHandler(repo url.Repository) urlHandler {
	return urlHandler{repo}
}

func (u urlHandler) NewURL(ctx *fiber.Ctx) error {
	return ctx.SendStatus(http.StatusOK)
}

func (u urlHandler) StoreURL(ctx *fiber.Ctx) error {
	reqBody := struct {
		URL string `json:"url" validate:"required,url"`
	}{}

	if err := ctx.BodyParser(&reqBody); err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	if verrs := ValidateStruct(reqBody); verrs != nil {
		return ctx.Status(http.StatusBadRequest).JSON(verrs)
	}

	created, err := u.repository.Save(reqBody.URL)
	if err != nil {
		return ctx.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return ctx.Status(http.StatusCreated).JSON(created)
}

func (u urlHandler) GoTo(ctx *fiber.Ctx) error {
	code := ctx.Params("code")

	foundURL, err := u.repository.FindByCode(code)
	if err != nil {
		// TODO: Redirect to "not found" page
		return ctx.Status(http.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	return ctx.Redirect(foundURL.Value)
}
