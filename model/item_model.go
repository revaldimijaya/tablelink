package model

import "time"

type Item struct {
	UUID      string     `json:"uuid" db:"uuid"`
	Name      string     `json:"name" db:"name"`
	Price     float64    `json:"price" db:"price"`
	Status    int        `json:"status" db:"status"`
	CreatedAt time.Time  `json:"created_at" db:"created_at"`
	UpdatedAt *time.Time `json:"updated_at" db:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at,omitempty" db:"deleted_at"`
}

type ItemIngredient struct {
	ItemUUID       string `json:"item_uuid" db:"uuid_item"`
	IngredientUUID string `json:"ingredient_uuid" db:"uuid_ingredient"`
}

type CreateItemRequest struct {
	Item
	IngredientsUUID []string `json:"ingredients"`
}
