package repository

import (
	"strings"
	"time"

	"github/revaldimijaya/tablelink/model"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

type ItemRepository struct {
	DB *sqlx.DB
}

func NewItemRepository(db *sqlx.DB) *ItemRepository {
	return &ItemRepository{DB: db}
}

func (r *ItemRepository) GetAll(filter model.Filter) ([]model.Item, error) {
	var items []model.Item

	query := `SELECT uuid, name, price, status, created_at, updated_at, deleted_at 
              FROM tm_item WHERE deleted_at IS NULL`

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

	err := r.DB.Select(&items, query, args...)
	if err != nil {
		return nil, err
	}

	return items, nil
}

func (r *ItemRepository) Create(item model.Item) (string, error) {
	uuid, _ := uuid.NewUUID()
	query := `INSERT INTO tm_item (uuid, name, price, status, created_at, updated_at) 
              VALUES (?, ?, ?, ?, ?, ?)`
	_, err := r.DB.Exec(query, uuid.String(), item.Name, item.Price, item.Status, time.Now(), time.Now())
	if err != nil {
		return "", err
	}
	return uuid.String(), nil
}

func (r *ItemRepository) Update(item model.Item) error {
	query := `UPDATE tm_item SET name = ?, price = ?, status = ?, updated_at = ? 
              WHERE uuid = ? AND deleted_at IS NULL`
	_, err := r.DB.Exec(query, item.Name, item.Price, item.Status, time.Now(), item.UUID)
	if err != nil {
		return err
	}
	return nil
}

func (r *ItemRepository) Delete(uuid string) error {
	query := `UPDATE tm_item SET deleted_at = ? WHERE uuid = ? AND deleted_at IS NULL`
	_, err := r.DB.Exec(query, time.Now(), uuid)
	return err
}
