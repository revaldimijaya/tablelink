package main

import (
	"log"

	"github/revaldimijaya/tablelink/db"
	"github/revaldimijaya/tablelink/handler/http"
	"github/revaldimijaya/tablelink/repository"
	"github/revaldimijaya/tablelink/usecase"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {
	app := fiber.New()

	// CORS Configuration
	app.Use(cors.New(cors.Config{
		AllowOrigins:     "http://localhost:5173/", // Follow url FE
		AllowMethods:     "GET, POST, PUT, DELETE",
		AllowHeaders:     "Content-Type, Authorization",
		AllowCredentials: true,
	}))

	// Initialize SQLite DB
	sqliteDB := db.InitDB("database.db")

	// Repositories
	ingredientRepo := repository.NewIngredientRepository(sqliteDB)
	itemRepo := repository.NewItemRepository(sqliteDB)
	itemIngredientRepo := repository.NewItemIngredientRepository(sqliteDB)

	// Usecases
	ingredientUsecase := usecase.NewIngredientUsecase(ingredientRepo)
	itemUsecase := usecase.NewItemUsecase(itemRepo, itemIngredientRepo)
	itemIngredientUsecase := usecase.NewItemIngredientUsecase(itemIngredientRepo)

	// Handlers
	ingredientHandler := http.NewIngredientHandler(ingredientUsecase)
	itemHandler := http.NewItemHandler(itemUsecase)
	itemIngredientHandler := http.NewItemIngredientHandler(itemIngredientUsecase)

	// Routes
	app.Get("/ingredients", ingredientHandler.GetAll)
	app.Post("/ingredients", ingredientHandler.Create)
	app.Put("/ingredients", ingredientHandler.Update)
	app.Delete("/ingredients", ingredientHandler.Delete)

	app.Get("/items", itemHandler.GetAll)
	app.Post("/items", itemHandler.Create)
	app.Put("/items", itemHandler.Update)
	app.Delete("/items", itemHandler.Delete)

	app.Delete("/item-ingredients", itemIngredientHandler.Delete)

	log.Fatal(app.Listen(":8000"))
}
