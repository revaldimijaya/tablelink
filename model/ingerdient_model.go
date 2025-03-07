package model

import "time"

type Ingredient struct {
	UUID        string     `json:"uuid" db:"uuid"`
	Name        string     `json:"name" db:"name"`
	CauseAlergy bool       `json:"cause_alergy" db:"cause_alergy"`
	Type        int        `json:"type" db:"type"`
	Status      int        `json:"status" db:"status"`
	CreatedAt   time.Time  `json:"created_at" db:"created_at"`
	UpdatedAt   *time.Time `json:"updated_at" db:"updated_at"`
	DeletedAt   *time.Time `json:"deleted_at,omitempty" db:"deleted_at"`
}
