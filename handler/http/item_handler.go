package http

import (
	"strconv"

	"github/revaldimijaya/tablelink/model"
	"github/revaldimijaya/tablelink/usecase"

	"github.com/gofiber/fiber/v2"
)

type ItemHandler struct {
	usecase *usecase.ItemUsecase
}

func NewItemHandler(usecase *usecase.ItemUsecase) *ItemHandler {
	return &ItemHandler{usecase: usecase}
}

func (h *ItemHandler) GetAll(c *fiber.Ctx) error {
	pagination, _ := strconv.Atoi(c.Query("pagination"))
	offset, _ := strconv.Atoi(c.Query("offset"))
	items, err := h.usecase.GetAll(pagination, offset)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(items)
}

func (h *ItemHandler) Create(c *fiber.Ctx) error {
	var item model.CreateItemRequest
	if err := c.BodyParser(&item); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request body"})
	}
	if err := h.usecase.Create(item); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}
	return c.Status(fiber.StatusCreated).JSON(fiber.Map{"message": "Item created successfully"})
}

func (h *ItemHandler) Update(c *fiber.Ctx) error {
	var item model.Item
	if err := c.BodyParser(&item); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request body"})
	}
	if err := h.usecase.Update(item); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{"message": "Item updated successfully"})
}

func (h *ItemHandler) Delete(c *fiber.Ctx) error {
	uuid := c.Query("uuid")
	if err := h.usecase.Delete(uuid); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{"message": "Item deleted successfully"})
}
