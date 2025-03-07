package model

type Filter struct {
	Name       string `json:"name" db:"name"`
	UUID       string `json:"uuid" db:"uuid"`
	Pagination int    `json:"pagination" db:"pagination"`
	Offset     int    `json:"offset" db:"offset"`
}
