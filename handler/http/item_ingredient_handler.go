package http

import (
	"github/revaldimijaya/tablelink/usecase"

	"github.com/gofiber/fiber/v2"
)

type ItemIngredientHandler struct {
	usecase *usecase.ItemIngredientUsecase
}

func NewItemIngredientHandler(usecase *usecase.ItemIngredientUsecase) *ItemIngredientHandler {
	return &ItemIngredientHandler{usecase: usecase}
}

func (h *ItemIngredientHandler) Delete(c *fiber.Ctx) error {
	itemUUID := c.Query("item_uuid")
	ingredientUUID := c.Query("ingredient_uuid")
	if err := h.usecase.Delete(itemUUID, ingredientUUID); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{"message": "Item-Ingredient relationship deleted successfully"})
}
