package handler

import (
	"gofiber-faafo/pkg/database/model"
	"gofiber-faafo/pkg/database/usecase"

	"github.com/gofiber/fiber/v2"
)

type databaseHandler struct {
	uc usecase.DatabaseUsecase
}

type DatabaseHandler interface {
	GetDatabases() fiber.Handler
	CreateDatabase() fiber.Handler
}

func NewDatabaseHandler(uc usecase.DatabaseUsecase) DatabaseHandler {
	return &databaseHandler{
		uc: uc,
	}
}

func (h *databaseHandler) GetDatabases() fiber.Handler {
	return func(c *fiber.Ctx) error {
		databases, err := h.uc.GetDatabases(c.Context())
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": err.Error(),
			})
		}

		return c.JSON(fiber.Map{
			"data": databases,
		})
	}
}

func (h *databaseHandler) CreateDatabase() fiber.Handler {
	return func(c *fiber.Ctx) error {
		var data model.Database
		if err := c.BodyParser(&data); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": err.Error(),
			})
		}

		id, err := h.uc.CreateDatabase(c.Context(), data)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": err.Error(),
			})
		}

		return c.JSON(fiber.Map{
			"id": id,
		})
	}
}
