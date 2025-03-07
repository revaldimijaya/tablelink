package repository

import (
	"strings"
	"time"

	"github/revaldimijaya/tablelink/model"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

type IngredientRepository struct {
	DB *sqlx.DB
}

func NewIngredientRepository(db *sqlx.DB) *IngredientRepository {
	return &IngredientRepository{DB: db}
}

func (r *IngredientRepository) GetAll(filter model.Filter) ([]model.Ingredient, error) {
	var ingredients []model.Ingredient
	query := `SELECT uuid, name, cause_alergy, type, status, created_at, updated_at, deleted_at 
              FROM tm_ingredient WHERE deleted_at IS NULL`

	var conditions []string
	var args []interface{}

	if filter.Name != "" {
		conditions = append(conditions, "name = ?")
		args = append(args, filter.Name)
	}

	if filter.UUID != "" {
		conditions = append(conditions, "uuid = ?")
		args = append(args, filter.UUID)
	}

	if len(conditions) > 0 {
		query += " AND " + strings.Join(conditions, " AND ")
	}

	query += " LIMIT ? OFFSET ?"
	args = append(args, filter.Pagination, filter.Offset)

	err := r.DB.Select(&ingredients, query, args...)
	if err != nil {
		return nil, err
	}
	return ingredients, nil
}

func (r *IngredientRepository) Create(ingredient model.Ingredient) error {
	uuid := uuid.New()
	query := `INSERT INTO tm_ingredient (uuid, name, cause_alergy, type, status, created_at, updated_at) 
              VALUES (?, ?, ?, ?, ?, ?, ?)`
	_, err := r.DB.Exec(query, uuid, ingredient.Name, ingredient.CauseAlergy, ingredient.Type, ingredient.Status, time.Now(), time.Now())
	if err != nil {
		return err
	}
	return nil
}

func (r *IngredientRepository) Update(ingredient model.Ingredient) error {
	query := `UPDATE tm_ingredient SET name = ?, cause_alergy = ?, type = ?, status = ?, updated_at = ? 
              WHERE uuid = ? AND deleted_at IS NULL`
	_, err := r.DB.Exec(query, ingredient.Name, ingredient.CauseAlergy, ingredient.Type, ingredient.Status, time.Now(), ingredient.UUID)
	if err != nil {
		return err
	}
	return nil
}

func (r *IngredientRepository) Delete(uuid string) error {
	query := `UPDATE tm_ingredient SET deleted_at = ? WHERE uuid = ? AND deleted_at IS NULL`
	_, err := r.DB.Exec(query, time.Now(), uuid)
	return err
}
