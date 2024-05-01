package databasehandler

import (
	"gofiber-faafo/pkg/index/model"
	"gofiber-faafo/pkg/index/usecase"

	"github.com/gofiber/fiber/v2"
)

type indexHandler struct {
	uc usecase.IndexUsecase
}

type IndexHandler interface {
	GetIndexes() fiber.Handler
	CreateIndex() fiber.Handler
}

func NewIndexHandler(uc usecase.IndexUsecase) IndexHandler {
	return &indexHandler{
		uc: uc,
	}
}

func (h *indexHandler) GetIndexes() fiber.Handler {
	return func(c *fiber.Ctx) error {
		dbName := c.Params("db_name")
		collName := c.Query("coll_name")

		indexes, err := h.uc.GetIndexes(c.Context(), dbName, collName)

		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": err.Error(),
			})
		}

		return c.JSON(fiber.Map{
			"data": indexes,
		})
	}
}

func (h *indexHandler) CreateIndex() fiber.Handler {
	return func(c *fiber.Ctx) error {
		dbName := c.Params("db_name")

		var indexData model.Index
		if err := c.BodyParser(&indexData); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": err.Error(),
			})
		}

		index, err := h.uc.CreateIndex(c.Context(), dbName, indexData)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": err.Error(),
			})
		}

		return c.JSON(fiber.Map{
			"data": index,
		})
	}
}
