package repository

import (
	"errors"
	"time"

	"github/revaldimijaya/tablelink/model"

	"github.com/jmoiron/sqlx"
)

type IngredientRepository struct {
	DB *sqlx.DB
}

func NewIngredientRepository(db *sqlx.DB) *IngredientRepository {
	return &IngredientRepository{DB: db}
}

func (r *IngredientRepository) GetAll(pagination int, offset int) ([]model.Ingredient, error) {
	var ingredients []model.Ingredient
	query := `SELECT uuid, name, cause_alergy, type, status, created_at, updated_at, deleted_at 
              FROM tm_ingredient WHERE deleted_at IS NULL LIMIT ? OFFSET ?`
	err := r.DB.Select(&ingredients, query, pagination, offset)
	if err != nil {
		return nil, err
	}
	return ingredients, nil
}

func (r *IngredientRepository) Create(ingredient model.Ingredient) error {
	query := `INSERT INTO tm_ingredient (uuid, name, cause_alergy, type, status, created_at, updated_at) 
              VALUES (?, ?, ?, ?, ?, ?, ?)`
	_, err := r.DB.Exec(query, ingredient.UUID, ingredient.Name, ingredient.CauseAlergy, ingredient.Type, ingredient.Status, time.Now(), time.Now())
	if err != nil {
		return errors.New("ingredient name must be unique")
	}
	return nil
}

func (r *IngredientRepository) Update(ingredient model.Ingredient) error {
	query := `UPDATE tm_ingredient SET name = ?, cause_alergy = ?, type = ?, status = ?, updated_at = ? 
              WHERE uuid = ? AND deleted_at IS NULL`
	_, err := r.DB.Exec(query, ingredient.Name, ingredient.CauseAlergy, ingredient.Type, ingredient.Status, time.Now(), ingredient.UUID)
	if err != nil {
		return errors.New("ingredient name must be unique")
	}
	return nil
}

func (r *IngredientRepository) Delete(uuid string) error {
	query := `UPDATE tm_ingredient SET deleted_at = ? WHERE uuid = ? AND deleted_at IS NULL`
	_, err := r.DB.Exec(query, time.Now(), uuid)
	return err
}
