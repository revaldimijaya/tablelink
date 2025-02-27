package http

import (
	"strconv"

	"github/revaldimijaya/tablelink/model"
	"github/revaldimijaya/tablelink/usecase"

	"github.com/gofiber/fiber/v2"
)

type IngredientHandler struct {
	usecase *usecase.IngredientUsecase
}

func NewIngredientHandler(usecase *usecase.IngredientUsecase) *IngredientHandler {
	return &IngredientHandler{usecase: usecase}
}

func (h *IngredientHandler) GetAll(c *fiber.Ctx) error {
	paginationStr := c.Query("pagination")
	offsetStr := c.Query("offset")

	pagination, err := strconv.Atoi(paginationStr)
	if err != nil || pagination <= 0 {
		pagination = 10
	}
	offset, err := strconv.Atoi(offsetStr)
	if err != nil || offset < 0 {
		offset = 0
	}

	ingredients, err := h.usecase.GetAll(pagination, offset)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.JSON(ingredients)
}

func (h *IngredientHandler) Create(c *fiber.Ctx) error {
	var ingredient model.Ingredient
	if err := c.BodyParser(&ingredient); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request body",
		})
	}

	if err := h.usecase.Create(ingredient); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message": "Ingredient created successfully",
	})
}

func (h *IngredientHandler) Update(c *fiber.Ctx) error {
	var ingredient model.Ingredient
	if err := c.BodyParser(&ingredient); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request body",
		})
	}

	if err := h.usecase.Update(ingredient); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Ingredient updated successfully",
	})
}

func (h *IngredientHandler) Delete(c *fiber.Ctx) error {
	uuid := c.Query("uuid")
	if uuid == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "UUID is required",
		})
	}

	if err := h.usecase.Delete(uuid); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Ingredient deleted successfully",
	})
}
