package repository

import (
	"github/revaldimijaya/tablelink/model"

	"github.com/jmoiron/sqlx"
)

type ItemIngredientRepository struct {
	DB *sqlx.DB
}

func NewItemIngredientRepository(db *sqlx.DB) *ItemIngredientRepository {
	return &ItemIngredientRepository{DB: db}
}

func (r *ItemIngredientRepository) GetIngredientsByItemUUID(itemUUID string) ([]model.ItemIngredient, error) {
	var ingredients []model.ItemIngredient
	query := `SELECT uuid_item, uuid_ingredient FROM tm_item_ingredient WHERE uuid_item = ?`
	err := r.DB.Select(&ingredients, query, itemUUID)
	if err != nil {
		return nil, err
	}
	return ingredients, nil
}

func (r *ItemIngredientRepository) Create(itemIngredient model.ItemIngredient) error {
	query := `INSERT INTO tm_item_ingredient (uuid_item, uuid_ingredient) VALUES (?, ?)`
	_, err := r.DB.Exec(query, itemIngredient.ItemUUID, itemIngredient.IngredientUUID)
	return err
}

func (r *ItemIngredientRepository) Delete(itemUUID string, ingredientUUID string) error {
	query := `DELETE FROM tm_item_ingredient WHERE uuid_item = ? AND uuid_ingredient = ?`
	_, err := r.DB.Exec(query, itemUUID, ingredientUUID)
	return err
}
