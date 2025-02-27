package main

import (
	"log"

	"github/revaldimijaya/tablelink/db"
	"github/revaldimijaya/tablelink/handler/http"
	"github/revaldimijaya/tablelink/repository"
	"github/revaldimijaya/tablelink/usecase"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	// Initialize SQLite DB
	sqliteDB := db.InitDB("tablelink.db")

	// Repositories
	ingredientRepo := repository.NewIngredientRepository(sqliteDB)

	// Usecases
	ingredientUsecase := usecase.NewIngredientUsecase(ingredientRepo)

	// Handlers
	ingredientHandler := http.NewIngredientHandler(ingredientUsecase)

	// Routes
	app.Get("/ingredients", ingredientHandler.GetAll)
	app.Post("/ingredients", ingredientHandler.Create)
	app.Put("/ingredients", ingredientHandler.Update)
	app.Delete("/ingredients", ingredientHandler.Delete)

	log.Fatal(app.Listen(":3000"))
}
